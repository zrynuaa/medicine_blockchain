package server

import (
	"encoding/json"
	"fmt"
	"github.com/BlockchainLabFudan/SM-Collection/SM3"
	"github.com/zrynuaa/cpabe06_client/bswabe"
	"github.com/zrynuaa/medicine_blockchain/backend/based"
	"log"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
	"strings"
	"time"
)

var pub *bswabe.BswabePub
var prv *bswabe.BswabePrv

//获取ABE服务器上的主公钥
func GetABEKeys(attrs string) (*bswabe.BswabePub, *bswabe.BswabePrv) {

	client, err := rpc.DialHTTP("tcp", "10.141.211.220:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call同步方式调用
	var reply []byte
	err = client.Call("CPABE.Getpub", "", &reply)
	if err != nil {
		log.Fatal(" error:", err)
	}
	pub = bswabe.UnSerializeBswabePub(reply) //获取PublicKey

	err = client.Call("CPABE.Getsk", attrs, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	prv = bswabe.UnSerializeBswabePrv(pub, reply) //获取服务端返回的解密私钥

	return pub, prv
}

func AddDoses() {
	//设置化学名与药品对应关系
	d1 := based.Dose{Medicine_name: "mid1", Chemistry_name: "cid1", Medicine_amount: 2, Medicine_price: 1.2}
	d2 := based.Dose{Medicine_name: "mid2", Chemistry_name: "cid2", Medicine_amount: 1, Medicine_price: 2.2}
	d3 := based.Dose{Medicine_name: "mid3", Chemistry_name: "cid2", Medicine_amount: 3, Medicine_price: 0.9}
	d4 := based.Dose{Medicine_name: "mid4", Chemistry_name: "cid3", Medicine_amount: 2, Medicine_price: 1.5}
	d5 := based.Dose{Medicine_name: "mid5", Chemistry_name: "cid4", Medicine_amount: 4, Medicine_price: 0.78}
	d6 := based.Dose{Medicine_name: "mid6", Chemistry_name: "cid4", Medicine_amount: 2, Medicine_price: 1.2}
	d7 := based.Dose{Medicine_name: "mid7", Chemistry_name: "cid4", Medicine_amount: 7, Medicine_price: 1.6}
	d8 := based.Dose{Medicine_name: "mid8", Chemistry_name: "cid5", Medicine_amount: 10, Medicine_price: 1.8}
	d9 := based.Dose{Medicine_name: "mid9", Chemistry_name: "cid6", Medicine_amount: 5, Medicine_price: 2.5}
	d10 := based.Dose{Medicine_name: "mid10", Chemistry_name: "cid7", Medicine_amount: 6, Medicine_price: 3.9}
	d11 := based.Dose{Medicine_name: "mid11", Chemistry_name: "cid7", Medicine_amount: 2, Medicine_price: 0.9}
	d12 := based.Dose{Medicine_name: "mid12", Chemistry_name: "cid8", Medicine_amount: 1, Medicine_price: 3.7}
	d13 := based.Dose{Medicine_name: "mid13", Chemistry_name: "cid8", Medicine_amount: 1, Medicine_price: 3.2}
	d14 := based.Dose{Medicine_name: "mid14", Chemistry_name: "cid8", Medicine_amount: 2, Medicine_price: 2.1}
	d15 := based.Dose{Medicine_name: "mid15", Chemistry_name: "cid9", Medicine_amount: 3, Medicine_price: 3}
	d16 := based.Dose{Medicine_name: "mid16", Chemistry_name: "cid10", Medicine_amount: 4, Medicine_price: 2}
	d17 := based.Dose{Medicine_name: "mid17", Chemistry_name: "cid10", Medicine_amount: 1, Medicine_price: 3.7}

	var dose []based.Dose
	dose = append(dose, d1, d2, d3, d4, d5, d6, d7, d8, d9, d10, d11, d12, d13, d14, d15, d16, d17)
	for k, v := range dose {
		based.PutIntoDb("dose", string(k), v.Serialize())
	}
}

//医院发布处方信息
func PrescriptiontoTransaction(pre HospitalPrescription) bool {

	prePolicy := "cid rid1 2of2 hid1 1of2"

	var ptot based.Prescription
	ptot.Type = 0
	ptot.Hospital_id = pre.Hospital_id
	ptot.Patient_id = pre.Patient_id
	ptot.Ts = uint64(time.Now().Unix())

	num := len(pre.Chemistrys)
	ptot.Data = new(based.Data_pre)
	ptot.Data.Doctor_id = pre.Doctor_id
	ptot.Data.Disease = pre.Disease

	//生成处方ID
	buf, _ := json.Marshal(ptot)
	digest := SM3.SM3_256(buf)
	easypreid := fmt.Sprintf("%x", digest)

	for i := 0; i < num; i++ {
		ptot.Data.Chemistry_name = pre.Chemistrys[i].Chemistry_name
		ptot.Data.Amount = pre.Chemistrys[i].Amount

		//为每个化学名生成不同的policy用于ABE加解密
		policy := strings.Replace(prePolicy, "cid", pre.Chemistrys[i].Chemistry_name, -1)
		ptot.Policy = policy
		fmt.Println(ptot.Policy)

		if i > 0 {
			easypreid = easypreid[:len(easypreid)-2]
		}
		easypreid += "_" + strconv.Itoa(i+1)
		ptot.Prescription_id = easypreid
		fmt.Println(ptot.Prescription_id, ptot)

		//加密后存储到链上
		preencdata := bswabe.SerializeBswabeCphKey(bswabe.CP_Enc(pub, string(ptot.Serialize()), ptot.Policy))
		_, err := based.PutIntoFabric("0", ptot.Prescription_id, preencdata)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}
	return true
}

//药店获取能解密的处方信息,处理成药品信息
func StoregetMInfo(store Drugstore) []Transaction {
	based.QuickAccess()
	pres, _ := based.GetPreFromDbByFilter(nil) //获取药店能够解密的所有的处方信息

	num := len(pres)
	var trans []Transaction

	for i := 0; i < num; i++ {
		var tran Transaction
		tran.Patient_id = pres[i].Patient_id

		//todo 怎么存储本地的全局信息
		mname := GetMedicineName(store, pres[i].Data.Chemistry_name) //获取药品名称

		for _, name := range mname {
			tran.Data = new(based.Data_tran)
			tran.Data.Prescription_id = pres[i].Prescription_id
			tran.Data.Ts = pres[i].Ts
			tran.Data.Site = store.Location

			tran.Data.Medicine_name = name
			amount, totalprice, _ := based.GetDoseFromDb(name, pres[i].Data.Chemistry_name, pres[i].Data.Amount)
			tran.Data.Amount = amount
			tran.Data.Price = totalprice

			//药方已处理时,该药品信息不能操作.药方未处理时,查看链上是否存在该药品信息,若存在则不能操作
			if IsBuy(pres[i].Prescription_id, "", "") {
				//处方已经被处理
				if IsBuy(tran.Data.Prescription_id, store.Location, name) {
					tran.Ishandled = 3 //该药品是该药店卖的
				} else {
					tran.Ishandled = 2 //别的药店卖出去的
				}
			} else {
				//处方未被处理
				if IsPostdata(tran.Data.Prescription_id, store.Location, name) {
					tran.Ishandled = 1 //处方还没结束，但是已经接单
				} else {
					tran.Ishandled = 0 //还没有接单
				}
			}

			trans = append(trans, tran)
		}
	}
	return trans
}

//药店发布药品信息
func StoresendTransaction(tran Transaction, sid string) {
	var ttot based.Transaction

	ttot.Type = 1
	ttot.Patient_id = tran.Patient_id
	ttot.Data = tran.Data
	ttot.Data.Ts = uint64(time.Now().Unix())

	buf, _ := json.Marshal(ttot)
	digest := SM3.SM3_256(buf)
	tranid := fmt.Sprintf("%x", digest)
	ttot.Transaction_id = tranid

	//todo policy,药品信息只让本药店以及服务节点能看到，所以只用一个属性，药店id--sid
	tranencdata := bswabe.SerializeBswabeCphKey(bswabe.CP_Enc(pub, string(ttot.Serialize()), sid))
	_, err := based.PutIntoFabric("1", ttot.Transaction_id, tranencdata)
	if err != nil {
		fmt.Println(err)
	}
}

//获取链上数据
func GetreadyInfo(mark, username string) ([]*Prescription, []*Transaction, []*based.Buy) {
	based.QuickAccess()
	if mark == "prescription" {
		var pres []*Prescription
		var preget []*based.Prescription

		if username == "*" {
			preget, _ = based.GetPreFromDbByFilter(nil)
		} else {
			fil := make(map[string]string)
			fil["patid"] = username
			preget, _ = based.GetPreFromDbByFilter(fil)
		}

		for _, v := range preget {
			pre := new(Prescription)
			pre.Data = v
			if IsBuy(v.Prescription_id, "", "") {
				pre.Isbuy = 1
			}
			pres = append(pres, pre)
		}
		return pres, nil, nil
	} else if mark == "transaction" {
		var trans []*Transaction
		var tranget []*based.Transaction

		if username == "*" {
			tranget, _ = based.GetTraFromDbByFilter(nil)
		} else {
			fil := make(map[string]string)
			fil["patid"] = username
			tranget, _ = based.GetTraFromDbByFilter(fil)
		}

		for _, v := range tranget {
			tran := new(Transaction)
			tran.Data = v.Data
			tran.Patient_id = v.Patient_id
			tran.Transaction_id = v.Transaction_id
			if IsBuy(v.Data.Prescription_id, "", "") {
				tran.Ishandled = 2
			}
			trans = append(trans, tran)
		}
		return nil, trans, nil
	} else if mark == "buy" {
		var buysget []*based.Buy
		if username == "*" {
			buysget, _ = based.GetBuyFromDbByFilter(nil)
		} else {
			fil := make(map[string]string)
			fil["patid"] = username
			buysget, _ = based.GetBuyFromDbByFilter(fil)
		}

		return nil, nil, buysget
	}
	return nil, nil, nil
}

//用户通过监管节点买药,发布买药信息
func BuyMedicine(tran based.Transaction) {
	var buy based.Buy
	data := tran.Data

	buy.Type = 2
	buy.Patient_id = tran.Patient_id

	buy.Data = &based.Data_buy{Prescription_id: data.Prescription_id, Medicine_name: data.Medicine_name, Medicine_amount: data.Amount, Medicine_price: data.Price, Site: data.Site}
	buy.Data.Ts = uint64(time.Now().Unix())

	buf, _ := json.Marshal(buy)
	digest := SM3.SM3_256(buf)
	buyid := fmt.Sprintf("%x", digest)
	buy.Buy_id = buyid

	//todo 需要获得这个处方id对应的处方，来获得该处方的化学名，进而构造policy
	predata, _ := based.GetFromDbById("prescription", buy.Data.Prescription_id)
	pre := based.DeserializePrescription(predata)

	prePolicy := "cid rid1 2of2 hid1 1of2"
	policy := strings.Replace(prePolicy, "cid", pre.Data.Chemistry_name, -1)

	//todo policy
	buyencdata := bswabe.SerializeBswabeCphKey(bswabe.CP_Enc(pub, string(buy.Serialize()), policy))

	based.PutIntoFabric("2", buy.Buy_id, buyencdata)
}

func AddHandletoServer(server *http.ServeMux, filename string) {
	fss := http.FileServer(http.Dir(os.Getenv("GOPATH") + "/src/github.com/zrynuaa/medicine_blockchain/frontend/static"))
	fsh := http.FileServer(http.Dir(os.Getenv("GOPATH") + "/src/github.com/zrynuaa/medicine_blockchain/frontend/html"))
	server.Handle("/static/", http.StripPrefix("/static/", fss))
	server.Handle("/html/"+filename, http.StripPrefix("/html/", fsh))
}

//todo
func GetMedicineName(store Drugstore, cname string) []string {
	var dose []Dose
	dose = store.Doses
	num := len(dose)
	for i := 0; i < num; i++ {
		if dose[i].Cname == cname {
			return dose[i].Mname
		}
	}
	return nil
}

//判断处方是否已经被处理，已处理返回true，否则返回false
func IsBuy(Presciption_id, Location, name string) bool {
	fil := make(map[string]string)
	fil["preid"] = Presciption_id
	if Location != "" {
		fil["site"] = Location
	}
	if name != "" {
		fil["medicine"] = name
	}

	trans, _ := based.GetBuyFromDbByFilter(fil)
	if trans != nil {
		return true
	}
	return false
}

//判断处方是否已接单，接单返回true，否则返回false
func IsPostdata(Presciption_id, Location, name string) bool {
	fil := make(map[string]string)
	fil["preid"] = Presciption_id
	fil["site"] = Location
	fil["medicine"] = name
	trans, _ := based.GetTraFromDbByFilter(fil)
	if trans != nil {
		return true
	}
	return false
}

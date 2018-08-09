package server

import (
	"github.com/zrynuaa/medicine_blockchain/backend/based"
	"encoding/json"
	//"crypto/md5"
	"fmt"
	"time"
	"net/http"
	"os"
	"strconv"
	"strings"
	"github.com/Doresimon/SM-Collection/SM3"
)

func AddDoses()  {
	//设置化学名与药品对应关系
	d1 := based.Dose{Medicine_name:"mid1", Chemistry_name:"cid1", Medicine_amount:2, Medicine_price:1.2}
	d2 := based.Dose{Medicine_name:"mid2", Chemistry_name:"cid2", Medicine_amount:1, Medicine_price:2.2}
	d3 := based.Dose{Medicine_name:"mid3", Chemistry_name:"cid2", Medicine_amount:3, Medicine_price:0.9}
	d4 := based.Dose{Medicine_name:"mid4", Chemistry_name:"cid3", Medicine_amount:2, Medicine_price:1.5}
	d5 := based.Dose{Medicine_name:"mid5", Chemistry_name:"cid4", Medicine_amount:4, Medicine_price:0.78}
	d6 := based.Dose{Medicine_name:"mid6", Chemistry_name:"cid4", Medicine_amount:2, Medicine_price:1.2}
	d7 := based.Dose{Medicine_name:"mid7", Chemistry_name:"cid4", Medicine_amount:7, Medicine_price:1.6}
	d8 := based.Dose{Medicine_name:"mid8", Chemistry_name:"cid5", Medicine_amount:10, Medicine_price:1.8}
	d9 := based.Dose{Medicine_name:"mid9", Chemistry_name:"cid6", Medicine_amount:5, Medicine_price:2.5}
	d10 := based.Dose{Medicine_name:"mid10", Chemistry_name:"cid7", Medicine_amount:6, Medicine_price:3.9}
	d11 := based.Dose{Medicine_name:"mid11", Chemistry_name:"cid7", Medicine_amount:2, Medicine_price:0.9}
	d12 := based.Dose{Medicine_name:"mid12", Chemistry_name:"cid7", Medicine_amount:1, Medicine_price:3.7}

	var dose []based.Dose
	dose = append(dose, d1, d2, d3, d4, d5, d6,d7,d8,d9,d10,d11,d12)
	for _,v := range dose {
		based.PutDose(v)
	}
}

//医院发布处方信息
func PrescriptiontoTransaction(pre HospitalPrescription) bool {

	prePolicy := "hid1 OR (cid AND rid1)"

	var ptot based.Presciption
	ptot.Type = 0
	ptot.Hospital_id = pre.Hospital_id
	ptot.Patient_id = pre.Patient_id
	ptot.Ts = uint64(time.Now().Unix())

	num := len(pre.Chemistrys)
	ptot.Data = new(based.Data_pre)
	ptot.Data.Doctor_id = pre.Doctor_id
	ptot.Data.Disease = pre.Disease

	//生成处方ID
	buf,_ := json.Marshal(ptot)
	//digest := md5.Sum(buf)
	digest := SM3.SM3_256(buf)
	easypreid := fmt.Sprintf("%x", digest)

	for i:=0; i<num ;i++{
		ptot.Data.Chemistry_name = pre.Chemistrys[i].Chemistry_name
		ptot.Data.Amount = pre.Chemistrys[i].Amount
		//policy := pre.Policy
		//fmt.Println(pre.Policy)
		policy := strings.Replace(prePolicy,"cid",pre.Chemistrys[i].Chemistry_name, -1)
		//policy = strings.Replace(policy,"rid","rid", -1)
		ptot.Policy = policy
		//fmt.Println(policy)

		if i>0{
			easypreid = easypreid[:len(easypreid)-2]
		}
		easypreid += "_" + strconv.Itoa(i+1)
		ptot.Presciption_id = easypreid

		
		based.PutPrescription(ptot)//处方上链
	}
	return true
}

//药店获取能解密的处方信息,处理成药品信息
func StoregetMInfo(store Drugstore) []Transaction {

	attrs := store.Attrs
	pres := based.GetPrescriptionByattr(attrs)//获取药店能够解密的所有的处方信息

	num := len(pres)
	var trans []Transaction

	for i:=0;i<num;i++{
		var tran Transaction
		tran.Patient_id = pres[i].Patient_id

		mname := GetMedicineName(store, pres[i].Data.Chemistry_name)	//获取药品名称

		for _,name := range mname {
			tran.Data = new(based.Data_tran)
			tran.Data.Presciption_id = pres[i].Presciption_id
			tran.Data.Ts = pres[i].Ts
			tran.Data.Site = store.Location

			tran.Data.Medicine_name = name
			amount, totalprice := based.GetDosedata(name, pres[i].Data.Chemistry_name, pres[i].Data.Amount)
			tran.Data.Amount = amount
			tran.Data.Price = totalprice

			//药方已处理时,该药品信息不能操作.药方未处理时,查看链上是否存在该药品信息,若存在则不能操作
			if based.IsBuy(pres[i].Presciption_id, "*", "*") {
				if based.IsBuy(tran.Data.Presciption_id,store.Location,name) {
					tran.Ishandled = 3		//该药品是该药店卖的
				}else {
					tran.Ishandled = 2
				}
			}else {
				if based.IsPostdata(tran.Data.Presciption_id, store.Location, name) {
					tran.Ishandled = 1
				}else {
					tran.Ishandled = 0
				}
			}

			trans = append(trans, tran)
		}
	}
	return trans
}

//药店发布药品信息
func StoresendTransaction(tran Transaction)  {
	var ttot based.Transaction

	ttot.Type = 1
	ttot.Patient_id = tran.Patient_id
	ttot.Data = tran.Data

	based.PutTransaction(ttot)
}

//获取链上数据
func GetreadyInfo(mark, username string) ([]Presciption, []Transaction) {
	if mark == "Prescription"{
		var pres []Presciption
		for _,v := range based.GetPrescriptionByid(username){
			pre := new(Presciption)
			pre.Data = v
			if based.IsBuy(v.Presciption_id,"*","*"){
				pre.Isbuy = 1
			}
			pres = append(pres, *pre)
		}
		return pres,nil
	}else {
		var trans []Transaction
		for _,v := range based.GetTransactionByid(username){
			tran := new(Transaction)
			tran.Data = v.Data
			tran.Patient_id = v.Patient_id
			if based.IsBuy(v.Data.Presciption_id,"*","*"){
				tran.Ishandled = 2
			}
			trans = append(trans, *tran)
		}
		return nil,trans
	}
}
//用户通过监管节点买药,发布买药信息
func BuyMedicine(tran based.Transaction)  {
	var buy based.Buy
	data := tran.Data

	buy.Type = 2
	buy.Patient_id = tran.Patient_id

	buy.Data = &based.Data_buy{Presciption_id:data.Presciption_id, Medicine_name:data.Medicine_name, Medicine_amount:data.Amount, Medicine_price:data.Price,Site:data.Site}
	buy.Data.Ts = uint64(time.Now().Unix())
	based.PutBuy(buy)
}

func AddHandletoServer(server *http.ServeMux, filename string)  {
	fss := http.FileServer(http.Dir(os.Getenv("GOPATH")+"/src/github.com/scottocs/medicine_blockchain/frontend/static"))
	fsh := http.FileServer(http.Dir(os.Getenv("GOPATH")+"/src/github.com/scottocs/medicine_blockchain/frontend/html"))
	server.Handle("/static/", http.StripPrefix("/static/", fss))
	server.Handle("/html/"+ filename, http.StripPrefix("/html/", fsh))
}

func GetMedicineName(store Drugstore, cname string) []string {
	var dose []*Dose
	dose = store.Doses
	num := len(dose)
	for i:=0;i<num;i++{
		if dose[i].Cname == cname{
			return dose[i].Mname
		}
	}
	return nil
}
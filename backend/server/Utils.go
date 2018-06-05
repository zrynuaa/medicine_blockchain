package server

import (
	"github.com/scottocs/medicine_blockchain/backend/based"
	"encoding/json"
	"crypto/md5"
	"fmt"
	"time"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func PrescriptiontoTransaction(pre HospitalPrescription) bool {

	var ptot based.Presciption
	ptot.Type = 0
	ptot.Hospital_id = pre.Hospital_id
	ptot.Patient_id = pre.Patient_id
	ptot.Ts = uint64(time.Now().Unix())
	ptot.Ishandled = false

	num := len(pre.Chemistrys)
	ptot.Data = new(based.Data_pre)
	ptot.Data.Doctor_id = pre.Doctor_id
	ptot.Data.Disease = pre.Disease

	//生成处方ID
	buf,_ := json.Marshal(ptot)
	digest := md5.Sum(buf)
	easypreid := fmt.Sprintf("%x", digest)

	for i:=0; i<num ;i++{
		ptot.Data.Chemistry_name = pre.Chemistrys[i].Chemistry_name
		ptot.Data.Amount = pre.Chemistrys[i].Amount
		//policy := pre.Policy
		policy := strings.Replace(pre.Policy,"Cid",pre.Chemistrys[i].Chemistry_name, -1)
		ptot.Policy = policy

		if i>0{
			easypreid = easypreid[:len(easypreid)-1]
		}
		easypreid += strconv.Itoa(i+1)
		ptot.Presciption_id = easypreid

		based.PutPrescription(ptot)//处方上链
	}

	var ptots []*based.Presciption
	ptots = based.GetPrescriptionByid(ptot.Patient_id)
	for i:=0;i<len(ptots);i++{
		fmt.Println(ptots[i],ptots[i].Data.Chemistry_name, ptots[i].Data.Amount)
	}
	return true
}

func StoregetMInfo(store Drugstore) []Transaction {

	attrs := store.Attrs
	pres := based.GetPrescriptionByattr(attrs)//获取药店能够解密的所有的处方信息

	num := len(pres)
	//trans := make([]Transaction,num)
	var trans []Transaction
	tran := new(Transaction)

	for i:=0;i<num;i++{
		tran.Data = new(based.Data_tran)
		tran.Patient_id = pres[i].Patient_id

		tran.Data.Presciption_id = pres[i].Presciption_id
		tran.Data.Ts = pres[i].Ts
		tran.Data.Site = store.Location

		mname := GetMedicineName(store, pres[i].Data.Chemistry_name)	//获取药品名称

		for _,name := range mname {
			tran.Data.Medicine_name = name
			amount, totalprice := based.GetDosedata(name, pres[i].Data.Chemistry_name, pres[i].Data.Amount)
			tran.Data.Amount = amount
			tran.Data.Price = totalprice

			//药方以处理时,该药品信息不能操作.药方未处理时,查看链上是否存在该药品信息,若存在则不能操作
			if pres[i].Ishandled {
				tran.Ishandled = true
			}else {
				tran.Ishandled = based.IsPostdata(tran.Data.Presciption_id, store.Location, name)
			}

			//TODO 后一个Mid会将后一个Mid覆盖掉
			trans = append(trans, *tran)
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

func AddHandletoServer(server *http.ServeMux)  {
	fss := http.FileServer(http.Dir(os.Getenv("GOPATH")+"/src/github.com/scottocs/medicine_blockchain/frontend/static"))
	fsh := http.FileServer(http.Dir(os.Getenv("GOPATH")+"/src/github.com/scottocs/medicine_blockchain/frontend/html"))
	server.Handle("/static/", http.StripPrefix("/static/", fss))
	server.Handle("/html/", http.StripPrefix("/html/", fsh))
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
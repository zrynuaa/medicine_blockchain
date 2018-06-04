package server

import (
	"github.com/scottocs/medicine_blockchain/backend/based"
	"encoding/json"
	"crypto/md5"
	"fmt"
	"time"
	"net/http"
	"os"
)

func PrescriptiontoTransaction(pre HospitalPrescription) bool {

	var ptot based.Presciption
	ptot.Type = 0
	ptot.Hospital_id = pre.Hospital_id
	ptot.Patient_id = pre.Patient_id
	ptot.Ts = uint64(time.Now().Unix())
	ptot.Ishandled = false
	ptot.Policy = pre.Policy

	num := len(pre.Chemistrys)
	ptot.Data = new(based.Data_pre)
	ptot.Data.Doctor_id = pre.Doctor_id
	ptot.Data.Disease = pre.Disease


	buf,_ := json.Marshal(ptot)
	digest := md5.Sum(buf)
	ptot.Presciption_id = fmt.Sprintf("%x", digest)//生成处方ID

	for i:=0; i<num ;i++{
		ptot.Data.Chemistry_name = pre.Chemistrys[i].Chemistry_name
		ptot.Data.Amount = pre.Chemistrys[i].Amount

		based.PutPrescription(ptot)//上链
	}

	var ptots []*based.Presciption
	ptots = based.GetPrescriptionByid(ptot.Presciption_id)
	for i:=0;i<len(ptots);i++{
		fmt.Println(ptots[i])
	}
	return true
}

func StoregetMInfo(store Drugstore) []based.Transaction {

	attrs := store.Attrs
	pres := based.GetPrescriptionByattr(attrs)//获取药店能够解密的所有的处方信息

	num := len(pres)
	trans := make([]based.Transaction,num)

	for i:=0;i<num;i++{
		trans[i].Data = new(based.Data_tran)
		trans[i].Patient_id = pres[i].Patient_id
		trans[i].Data.Presciption_id = pres[i].Presciption_id
		trans[i].Data.Ts = pres[i].Ts
		trans[i].Data.Site = store.Location
		trans[i].Data.Ishandled = false

		//TODO 获取药品名称
		trans[i].Data.Medicine_name = " "
		amount,totalprice := based.GetDosedata(trans[i].Data.Medicine_name, pres[i].Data.Chemistry_name, 1)
		trans[i].Data.Amount = amount
		trans[i].Data.Price = totalprice
	}
	return trans
}

//药店发布药品信息
func StoresendTransaction(tran Transaction)  {
	var ttot based.Transaction

	ttot.Type = 1
	ttot.Patient_id = tran.Patient_id
	ttot.Data = ttot.Data

	based.PutTransaction(ttot)
}

func AddHandletoServer(server *http.ServeMux)  {
	fss := http.FileServer(http.Dir(os.Getenv("GOPATH")+"/src/github.com/scottocs/medicine_blockchain/frontend/static"))
	fsh := http.FileServer(http.Dir(os.Getenv("GOPATH")+"/src/github.com/scottocs/medicine_blockchain/frontend/html"))
	server.Handle("/static/", http.StripPrefix("/static/", fss))
	server.Handle("/html/", http.StripPrefix("/html/", fsh))
}
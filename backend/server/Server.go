package server

import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
	"github.com/scottocs/medicine_blockchain/backend/based"
)

type preDecodeBystore struct {
	pres []*based.Presciption `json:"pres"`
}
var drugstore1 = Drugstore{Name:"", location:" ", attrs:[]string{"Cid1","Cid8","Cid9","Rid1"}}
var drugstore2 = Drugstore{Name:"", location:" ", attrs:[]string{"Cid1","Cid2","Cid8","Cid9","Rid2"}}

func AllChainInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	str,_ := r.GetBody()
	fmt.Println(str)
	fmt.Fprintf(w, "AllChainInfo")
}

func HospitalSendPrescription(w http.ResponseWriter, r *http.Request)  {
	var pre HospitalPrescription
	json.NewDecoder(r.Body).Decode(&pre)
	fmt.Fprint(w,http.StatusOK)

	PrescriptiontoTransaction(pre)	//将处方信息分解成只含一个化学名的交易信息
}

func DrugstoregetMInfo(w http.ResponseWriter, r *http.Request)  {
	var presciption preDecodeBystore
	presciption.pres = based.GetPrescriptionByattr(drugstore1.attrs)
}

func Run() {
	//设置路由
	http.HandleFunc("/", AllChainInfo)
	http.HandleFunc("/HospitalSendPrescription", HospitalSendPrescription) //post
	http.HandleFunc("/drugstore", DrugstoregetMInfo)

	//监听端口
	err := http.ListenAndServe(":8880", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


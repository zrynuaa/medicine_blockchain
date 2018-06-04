package server

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/scottocs/medicine_blockchain/backend/based"
)

type preDecodeBystore struct {
	pres []*based.Presciption `json:"pres"`
}
var drugstore1 Drugstore
var drugstore2 Drugstore


func HospitalSendPrescription(w http.ResponseWriter, r *http.Request)  {
	var pre HospitalPrescription
	json.NewDecoder(r.Body).Decode(&pre)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if PrescriptiontoTransaction(pre){		//将处方信息分解成只含一个化学名的交易信息
		fmt.Fprint(w,http.StatusOK)
	}
}

func Store1getMInfo(w http.ResponseWriter, r *http.Request)  {
	var trans []based.Transaction
	trans = StoregetMInfo(drugstore1)
	json.NewEncoder(w).Encode(trans)
}

func Store2getMInfo(w http.ResponseWriter, r *http.Request)  {
	var trans []based.Transaction
	trans = StoregetMInfo(drugstore2)
	json.NewEncoder(w).Encode(trans)
}

func Sethandle(w http.ResponseWriter, r *http.Request)  {
	var tran Transaction
	json.NewDecoder(r.Body).Decode(&tran)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//TODO 应该通过处方ID和药品信息更新
	based.UpdatePrescription(tran.Data.Presciption_id)
	StoresendTransaction(tran)
}

func Supervision(w http.ResponseWriter, r *http.Request)  {

}

func Run()  {

	finish := make(chan bool)

	server8880 := http.NewServeMux()
	AddHandletoServer(server8880)
	server8880.HandleFunc("/HospitalSendPrescription", HospitalSendPrescription)

	server8881 := http.NewServeMux()
	AddHandletoServer(server8881)
	&drugstore1 = SetStore1Attrs()
	server8881.HandleFunc("/", Store1getMInfo)
	server8881.HandleFunc("/Sethandle1", Sethandle)

	server8882 := http.NewServeMux()
	AddHandletoServer(server8882)
	&drugstore2 = SetStore2Attrs()
	server8882.HandleFunc("/", Store2getMInfo)
	server8881.HandleFunc("/Sethandle2", Sethandle)

	server8883 := http.NewServeMux()
	AddHandletoServer(server8883)
	server8883.HandleFunc("/", Supervision)

	go func() {
		http.ListenAndServe(":8880", server8880)
	}()

	go func() {
		http.ListenAndServe(":8881", server8881)
	}()

	go func() {
		http.ListenAndServe(":8882", server8882)
	}()

	go func() {
		http.ListenAndServe(":8883", server8883)
	}()

	<-finish
}




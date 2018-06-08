package server

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/scottocs/medicine_blockchain/backend/based"
)

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
	var trans []Transaction
	trans = StoregetMInfo(drugstore1)
	json.NewEncoder(w).Encode(trans)
}

func Store2getMInfo(w http.ResponseWriter, r *http.Request)  {
	var trans []Transaction
	trans = StoregetMInfo(drugstore2)
	json.NewEncoder(w).Encode(trans)
}

func Sethandle(w http.ResponseWriter, r *http.Request)  {
	var tran Transaction
	json.NewDecoder(r.Body).Decode(&tran)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	StoresendTransaction(tran)
	fmt.Fprint(w,http.StatusOK)
}

func GetPrescriptions(w http.ResponseWriter, r *http.Request)  {
	pres,_ := GetreadyInfo("Prescription", r.FormValue("username"))
	json.NewEncoder(w).Encode(pres)
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	trans,_ := GetreadyInfo("Transaction", r.FormValue("username"))
	json.NewEncoder(w).Encode(trans)
}

func GetBuys(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(based.GetBuyByid(r.FormValue("username")))
}

func UserbuyMedicine(w http.ResponseWriter, r *http.Request) {
	var tran based.Transaction
	json.NewDecoder(r.Body).Decode(&tran)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	BuyMedicine(tran)
	fmt.Fprint(w,http.StatusOK)
}

func Run()  {
	finish := make(chan bool)

	server8880 := http.NewServeMux()
	AddHandletoServer(server8880, "hospital.html")
	server8880.HandleFunc("/hospitalsendprescription", HospitalSendPrescription)

	server8881 := http.NewServeMux()
	AddHandletoServer(server8881, "store.html")
	drugstore1 = SetStore1Attrs()
	server8881.HandleFunc("/", Store1getMInfo)
	server8881.HandleFunc("/sethandle1", Sethandle)

	server8882 := http.NewServeMux()
	AddHandletoServer(server8882, "store.html")
	drugstore2 = SetStore2Attrs()
	server8882.HandleFunc("/", Store2getMInfo)
	server8882.HandleFunc("/sethandle2", Sethandle)

	server8883 := http.NewServeMux()
	AddHandletoServer(server8883, "controller.html")
	//server8883.HandleFunc("/", Supervision)
	server8883.HandleFunc("/getprescriptions", GetPrescriptions)
	server8883.HandleFunc("/gettransactions", GetTransactions)
	server8883.HandleFunc("/getbuys", GetBuys)
	server8883.HandleFunc("/userbuymedicine", UserbuyMedicine)

	go http.ListenAndServe(":8880", server8880)
	go http.ListenAndServe(":8881", server8881)
	go http.ListenAndServe(":8882", server8882)
	go http.ListenAndServe(":8883", server8883)

	<-finish
}




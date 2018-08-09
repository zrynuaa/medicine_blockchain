package server

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/zrynuaa/medicine_blockchain/backend/based"
)

var drugstore1 Drugstore
var drugstore2 Drugstore
var drugstore3 Drugstore

func setAccess(w http.ResponseWriter)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func HospitalSendPrescription(w http.ResponseWriter, r *http.Request)  {
	setAccess(w)
	var pre HospitalPrescription
	json.NewDecoder(r.Body).Decode(&pre)

	if PrescriptiontoTransaction(pre){		//将处方信息分解成只含一个化学名的交易信息
		fmt.Fprint(w,http.StatusOK)
	}
}

func Store1getMInfo(w http.ResponseWriter, r *http.Request)  {
	setAccess(w)
	var trans []Transaction
	trans = StoregetMInfo(drugstore1)
	json.NewEncoder(w).Encode(trans)
}

func Store2getMInfo(w http.ResponseWriter, r *http.Request)  {
	setAccess(w)
	var trans []Transaction
	trans = StoregetMInfo(drugstore2)
	json.NewEncoder(w).Encode(trans)
}

func Store3getMInfo(w http.ResponseWriter, r *http.Request)  {
w.Header().Set("Access-Control-Allow-Origin", "*")
	var trans []Transaction
	trans = StoregetMInfo(drugstore3)
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	pres,_ := GetreadyInfo("Prescription", r.FormValue("username"))
	json.NewEncoder(w).Encode(pres)
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// trans,_ := GetreadyInfo("Transaction", r.FormValue("username"))
	_,trans := GetreadyInfo("Transaction", r.FormValue("username"))
	json.NewEncoder(w).Encode(trans)
}

func GetBuys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(based.GetBuyByid(r.FormValue("username")))
}

func UserbuyMedicine(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var tran based.Transaction
	json.NewDecoder(r.Body).Decode(&tran)

	BuyMedicine(tran)
	fmt.Fprint(w,http.StatusOK)
}


func GetBlockchain(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(based.GetBlock())
}

func Run()  {
	if amount,_ := based.GetDosedata("mid1", "cid1", 1); amount == 0 {
		AddDoses() //初始化化学名与药品对应关系以及药品价格
	}

	finish := make(chan bool)

	server8880 := http.NewServeMux()
	AddHandletoServer(server8880, "hospital.html")
	server8880.HandleFunc("/hospitalsendprescription", HospitalSendPrescription)

	server8881 := http.NewServeMux()
	AddHandletoServer(server8881, "store.html")
	drugstore1 = SetStore1Attrs()
	server8881.HandleFunc("/getprelist8881", Store1getMInfo)
	server8881.HandleFunc("/sethandle8881", Sethandle)

	server8882 := http.NewServeMux()
	AddHandletoServer(server8882, "store.html")
	drugstore2 = SetStore2Attrs()
	server8882.HandleFunc("/getprelist8882", Store2getMInfo)
	server8882.HandleFunc("/sethandle8882", Sethandle)

	server8883 := http.NewServeMux()
	AddHandletoServer(server8883, "store.html")
	drugstore3 = SetStore3Attrs()
	server8883.HandleFunc("/getprelist8883", Store3getMInfo)
	server8883.HandleFunc("/sethandle8883", Sethandle)

	server8884 := http.NewServeMux()
	AddHandletoServer(server8884, "controller.html")
	server8884.HandleFunc("/getprescriptions", GetPrescriptions)
	server8884.HandleFunc("/gettransactions", GetTransactions)
	server8884.HandleFunc("/getbuys", GetBuys)
	server8884.HandleFunc("/userbuymedicine", UserbuyMedicine)

	server8885 := http.NewServeMux()
	AddHandletoServer(server8885, "blockExplore.html")
	server8885.HandleFunc("/getblockchain", GetBlockchain)

	go http.ListenAndServe(":8880", server8880)
	go http.ListenAndServe(":8881", server8881)
	go http.ListenAndServe(":8882", server8882)
	go http.ListenAndServe(":8883", server8883)
	go http.ListenAndServe(":8884", server8884)
	go http.ListenAndServe(":8885", server8885)

	<-finish
}

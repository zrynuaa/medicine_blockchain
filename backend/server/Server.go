package server

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/zrynuaa/medicine_blockchain/backend/based"
	"github.com/zrynuaa/cpabe06_client/bswabe"
)

type Peer struct {
	Typ int				//医院1、药店2、服务节点3
	Store Drugstore
	Hospital Hospital
	Controller Controller

	Port string
}

var drugstore Drugstore

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

func StoregetMInfos(w http.ResponseWriter, r *http.Request)  {
	setAccess(w)
	var trans []Transaction
	trans = StoregetMInfo(drugstore)
	json.NewEncoder(w).Encode(trans)
}

func Sethandle(w http.ResponseWriter, r *http.Request)  {
	var tran Transaction
	setAccess(w)
	json.NewDecoder(r.Body).Decode(&tran)

	StoresendTransaction(tran, drugstore.ID)
	fmt.Fprint(w,http.StatusOK)
}

func GetPrescriptions(w http.ResponseWriter, r *http.Request)  {
	setAccess(w)
	pres,_,_ := GetreadyInfo("prescription", r.FormValue("username"))
	json.NewEncoder(w).Encode(pres)
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	setAccess(w)
	_,trans,_ := GetreadyInfo("transaction", r.FormValue("username"))
	json.NewEncoder(w).Encode(trans)
}

func GetBuys(w http.ResponseWriter, r *http.Request) {
	setAccess(w)
	_,_,buys := GetreadyInfo("buy", r.FormValue("username"))
	json.NewEncoder(w).Encode(buys)
}

func UserbuyMedicine(w http.ResponseWriter, r *http.Request) {
	setAccess(w)
	var tran based.Transaction
	json.NewDecoder(r.Body).Decode(&tran)

	BuyMedicine(tran)
	fmt.Fprint(w,http.StatusOK)
}

func Run(peer *Peer)  {
	server := http.NewServeMux()
	var pub *bswabe.BswabePub
	var prv *bswabe.BswabePrv

	if peer.Typ == 1 {			//医院
		pub,prv = GetABEKeys(peer.Hospital.Attrs) //获取ABE服务上的密钥对
		based.Init(peer.Hospital.Name, pub, prv)

		AddHandletoServer(server, "hospital.html")
		server.HandleFunc("/hospitalsendprescription", HospitalSendPrescription)
	}else if peer.Typ == 2 {	//药店
		drugstore = peer.Store

		pub,prv = GetABEKeys(peer.Store.Attrs) //获取ABE服务上的密钥对
		based.Init(peer.Store.Name, pub, prv)
		AddDoses()

		AddHandletoServer(server, "store.html")
		server.HandleFunc("/getprelist" + peer.Port, StoregetMInfos)
		server.HandleFunc("/sethandle" + peer.Port, Sethandle)
	}else if peer.Typ ==3 {		//服务节点
		pub,prv = GetABEKeys(peer.Controller.Attrs) //获取ABE服务上的密钥对
		based.Init("Controller", pub, prv)

		AddHandletoServer(server, "controller.html")
		server.HandleFunc("/getprescriptions", GetPrescriptions)
		server.HandleFunc("/gettransactions", GetTransactions)
		server.HandleFunc("/getbuys", GetBuys)
		server.HandleFunc("/userbuymedicine", UserbuyMedicine)
	}

	fmt.Println("peer initial done!")
	http.ListenAndServe(":" + peer.Port, server)
}

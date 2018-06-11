package main

import (
	"fmt"
	"github.com/scottocs/medicine_blockchain/backend/server"
	"github.com/scottocs/medicine_blockchain/backend/based"
)

func main()  {
	//设置化学名与药品对应关系
	server.AddDoses()

	//医院上传处方信息 上链
	hp := server.HospitalPrescription{Hospital_id:"1234", Patient_id:"350622199009086758", Doctor_id:"67534", Disease:"上呼吸道感染", Policy:"hid1 OR (cid AND rid1)"}
	var ch []server.Chemistry
	ch = append(ch, server.Chemistry{Chemistry_name:"cid1",Amount:2})
	ch = append(ch, server.Chemistry{Chemistry_name:"cid2",Amount:4})
	ch = append(ch, server.Chemistry{Chemistry_name:"cid7",Amount:1})
	hp.Chemistrys = ch

	fmt.Println("医院处方信息:", hp)
	server.PrescriptiontoTransaction(hp)

	fmt.Println("病人处方信息:")
	for _,v := range  based.GetPrescriptionByid(hp.Patient_id){
		fmt.Println(v,v.Data)
	}

	//药店获取药品信息
	fmt.Println("\n药店基本信息:")
	drugstore1 := server.SetStore1Attrs()
	fmt.Println(drugstore1, drugstore1.Doses)
	drugstore2 := server.SetStore2Attrs()
	fmt.Println(drugstore2, drugstore2.Doses)

	fmt.Println("\n药店1能解密的信息:")
	tr1 := server.StoregetMInfo(drugstore1)
	for _,v := range tr1{
		fmt.Println(v, v.Data)
	}

	fmt.Println("\n药店2能解密的信息:")
	tr2 := server.StoregetMInfo(drugstore2)
	for _,v := range tr2{
		fmt.Println(v, v.Data)
	}

	//药店上传药品信息
	server.StoresendTransaction(tr1[0])
	server.StoresendTransaction(tr1[2])

	server.StoresendTransaction(tr2[0])
	server.StoresendTransaction(tr2[1])

	fmt.Println("\n病人查看药品信息:")
	_,trans := server.GetreadyInfo("Transaction",hp.Patient_id)
	for _,t := range trans {
		fmt.Println(t, t.Data)
	}

	fmt.Println("\n病人查看处方信息:")
	pres,_ := server.GetreadyInfo("Prescription",hp.Patient_id)
	for _,t := range pres {
		fmt.Println(t, t.Data)
	}

	fmt.Println("\n病人查看买药信息:")
	for _,buy := range based.GetBuyByid(hp.Patient_id){
		fmt.Println(buy, buy.Data)
	}

	//病人买药
	fmt.Println("\n病人买药--------------")
	var tran based.Transaction
	tran.Data = tr1[0].Data
	tran.Patient_id = tr1[0].Patient_id
	tran.Type = 1
	server.BuyMedicine(tran)


	fmt.Println("\n病人查看买药信息:")
	for _,buy := range based.GetBuyByid(hp.Patient_id){
		fmt.Println(buy, buy.Data)
	}

}

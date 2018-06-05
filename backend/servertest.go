package main

import (
	"fmt"
	"github.com/scottocs/medicine_blockchain/backend/server"
	"github.com/scottocs/medicine_blockchain/backend/based"
)

func main()  {
	//医院上传处方信息 上链
	hp := server.HospitalPrescription{Hospital_id:"1234", Patient_id:"350622199009086758", Doctor_id:"67534", Disease:"上呼吸道感染", Policy:"Hid1 OR (Cid AND Rid1)"}
	var ch []server.Chemistry
	ch = append(ch, server.Chemistry{Chemistry_name:"Cid1",Amount:2})
	ch = append(ch, server.Chemistry{Chemistry_name:"Cid2",Amount:4})
	hp.Chemistrys = ch

	fmt.Println("医院处方信息:", hp)

	//server.PrescriptiontoTransaction(hp)

	//药店获取药品信息
	fmt.Println((based.GetPrescriptionByid(hp.Patient_id))[0])
	fmt.Println((based.GetPrescriptionByid(hp.Patient_id))[1])

	drugstore1 := server.SetStore1Attrs()
	fmt.Println(drugstore1, drugstore1.Doses[0].Mname)
	drugstore2 := server.SetStore2Attrs()
	fmt.Println(drugstore2, drugstore2.Doses[0].Mname)

	tr1 := server.StoregetMInfo(drugstore1)
	for _,v := range tr1{
		fmt.Println(v, v.Data)
	}

	tr2 := server.StoregetMInfo(drugstore2)
	for _,v := range tr2{
		fmt.Println(v, v.Data)
	}

	//药店上传药品信息
	server.StoresendTransaction(tr1[1]) //TODO openfile failed ,can't store tran
	for _,t := range based.GetTransactionByid(hp.Patient_id) {
		fmt.Println(t, t.Data)
	}

	based.Update(based.GetPrescriptionByid(hp.Patient_id)[0].Presciption_id)
	fmt.Println((based.GetPrescriptionByid(hp.Patient_id))[0])
}

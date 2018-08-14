package server

import (
	"testing"
	"github.com/zrynuaa/medicine_blockchain/backend/based"
	"fmt"
)

func TestPrescriptiontoTransaction(t *testing.T) {
	//based.Setup()
	GetABEKeys() //获取ABE服务上的密钥对
	hp := HospitalPrescription{
		Hospital_id:"huashan",
		Patient_id:"111",
		Doctor_id:"1",
		Disease:"fever",
		Chemistrys:[]Chemistry{
			{
				Chemistry_name:"cid3",
				Amount:5,
			},
			{
				Chemistry_name:"cid4",
				Amount:1,
			},
		},
	}

	PrescriptiontoTransaction(hp) //将处方信息存到链上
	based.Init("zry",pub,prv)
	based.QuickAccess() //马上获取新的链上信息

	all,_ := based.GetPreFromDbByFilter(nil)
	for _,v := range all{
		fmt.Println(v, v.Data)
	}
}

func TestStoregetMInfo(t *testing.T) {
	GetABEKeys()
	based.Init("zry",pub,prv)
	AddDoses()
	drugstore1 := SetStore1Attrs()

	trans := StoregetMInfo(drugstore1)

	for _,v := range trans{
		fmt.Println(v, v.Data)
	}
}

func TestAddDoses(t *testing.T) {
	GetABEKeys()
	based.Init("zry",pub,prv)
	AddDoses()
	fmt.Println(based.GetDoseFromDb("mid1","cid1",2))
}

func TestStoresendTransaction(t *testing.T) {
	GetABEKeys()
	based.Init("zry",pub,prv)
	AddDoses()

	drugstore1 := SetStore1Attrs()
	trans := StoregetMInfo(drugstore1)

	for _, v := range trans{
		if v.Ishandled == 0 {		//未处理时，上传药品信息
			StoresendTransaction(v)
		}
	}

	all,_ := based.GetTraFromDbByFilter(nil)
	for _,v := range all{
		fmt.Println(v, v.Data)
	}
}


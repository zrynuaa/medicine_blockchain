package server

import (
	"testing"
	"github.com/zrynuaa/medicine_blockchain/backend/based"
	"fmt"
)

func TestPrescriptiontoTransaction(t *testing.T) {
	//based.Setup()
	GetABEPub() //获取ABE服务上的公钥
	hp := HospitalPrescription{
		Hospital_id:"huashan",
		Patient_id:"111",
		Doctor_id:"1",
		Disease:"fever",
		Chemistrys:[]Chemistry{
			{
				Chemistry_name:"cid3",
				Amount:2,
			},
			{
				Chemistry_name:"cid4",
				Amount:3,
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
	drugstore1 := SetStore1Attrs()

	trans := StoregetMInfo(drugstore1)

	for _,v := range trans{
		fmt.Println(v, v.Data)
	}
}

func TestAddDoses(t *testing.T) {
	AddDoses()
	based.GetDoseFromDb("mid1","cid1",2)
}

func TestGetBuys(t *testing.T) {

}


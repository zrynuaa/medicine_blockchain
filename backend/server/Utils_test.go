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
		Hospital_id:"zhongshan",
		Patient_id:"111",
		Doctor_id:"1",
		Disease:"fever",
		Chemistrys:[]Chemistry{
			{
				Chemistry_name:"cid1",
				Amount:2,
			},
			{
				Chemistry_name:"cid2",
				Amount:3,
			},
		},
	}

	PrescriptiontoTransaction(hp) //将处方信息存到链上
	based.Init("zry",pub,prv)
	based.QuickAccess() //马上获取新的链上信息

	all,_ := based.GetPreFromDbByFilter(nil)
	for _,v := range all{
		fmt.Println(v)
	}
}

func TestGetBuys(t *testing.T) {

}

func TestGetPrescriptions(t *testing.T) {

}

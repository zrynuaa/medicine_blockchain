package server

import (
	"testing"
)

func TestPrescriptiontoTransaction(t *testing.T) {
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
}

func TestGetBuys(t *testing.T) {

}

func TestGetPrescriptions(t *testing.T) {

}

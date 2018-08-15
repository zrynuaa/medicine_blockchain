package server

import (
	"testing"
	"github.com/zrynuaa/medicine_blockchain/backend/based"
	"fmt"
)

func TestPrescriptiontoTransaction(t *testing.T) {
	//based.Setup() //整个系统启动只运行一次
	GetABEKeys() //获取ABE服务上的密钥对
	based.Init("zry",pub,prv)

	hp := HospitalPrescription{
		Hospital_id:"zhongshan",
		Patient_id:"111222",
		Doctor_id:"123456",
		Disease:"fever",
		Chemistrys:[]Chemistry{
			{
				Chemistry_name:"cid1",
				Amount:5,
			},
			{
				Chemistry_name:"cid2",
				Amount:1,
			},
		},
	}

	PrescriptiontoTransaction(hp) //将处方信息存到链上
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

	based.QuickAccess() //马上获取新的链上信息
	all,_ := based.GetTraFromDbByFilter(nil)
	for _,v := range all{
		fmt.Println(v, v.Data)
	}
}

func TestBuyMedicine(t *testing.T) {
	GetABEKeys() //获取ABE服务上的密钥对
	based.Init("zry",pub,prv)

	fmt.Println("所有能解密的药品信息（药店卖药信息）")
	trans,_ := based.GetTraFromDbByFilter(nil)
	for _,v := range trans{
		fmt.Println(v, v.Data)
	}

	BuyMedicine(*trans[1])
}

func TestGetreadyInfo(t *testing.T) {
	GetABEKeys() //获取ABE服务上的密钥对
	based.Init("zry",pub,prv)
	based.QuickAccess() //马上获取新的链上信息

	fmt.Println("所有能解密的处方信息")
	pres,_ := based.GetPreFromDbByFilter(nil)
	for _,v := range pres{
		fmt.Println(v, v.Data)
	}

	fmt.Println("所有能解密的药品信息（药店卖药信息）")
	trans,_ := based.GetTraFromDbByFilter(nil)
	for _,v := range trans{
		fmt.Println(v, v.Data)
	}

	fmt.Println("所有能解密的用户买药信息")
	buys,_ := based.GetBuyFromDbByFilter(nil)
	for _,v := range buys{
		fmt.Println(v, v.Data)
	}

}


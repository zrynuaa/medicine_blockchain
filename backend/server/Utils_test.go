package server

import (
	"testing"
	"github.com/zrynuaa/medicine_blockchain/backend/based"
	"fmt"
)

var attrs = "cid1 cid2 cid3 cid4 cid5 cid6 cid7 cid8 cid9 cid10 rid1 sid1"

func TestPrescriptiontoTransaction(t *testing.T) {
	//based.Setup() //整个系统启动只运行一次
	pub, prv = GetABEKeys(attrs) //获取ABE服务上的密钥对
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
	pub, prv = GetABEKeys(attrs) //获取ABE服务上的密钥对
	based.Init("zry",pub,prv)
	AddDoses()

	trans := StoregetMInfo(testStore1)

	for _,v := range trans{
		fmt.Println(v, v.Data)
	}
}

func TestStoresendTransaction(t *testing.T) {
	pub, prv = GetABEKeys(attrs) //获取ABE服务上的密钥对
	based.Init("zry",pub,prv)
	AddDoses()

	//drugstore1 := SetStore1Attrs()
	trans := StoregetMInfo(testStore1)

	for _, v := range trans{
		if v.Ishandled == 0 {		//未处理时，上传药品信息
			StoresendTransaction(v, "sid1")
		}
	}

	based.QuickAccess() //马上获取新的链上信息
	all,_ := based.GetTraFromDbByFilter(nil)
	for _,v := range all{
		fmt.Println(v, v.Data)
	}
}

func TestBuyMedicine(t *testing.T) {
	pub, prv = GetABEKeys(attrs) //获取ABE服务上的密钥对
	based.Init("zry",pub,prv)

	fmt.Println("所有能解密的药品信息（药店卖药信息）")
	trans,_ := based.GetTraFromDbByFilter(nil)
	for _,v := range trans{
		fmt.Println(v, v.Data)
	}

	BuyMedicine(*trans[1])
}

func TestGetreadyInfo(t *testing.T) {
	pub, prv = GetABEKeys(attrs) //获取ABE服务上的密钥对
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

	fmt.Println("用户111222的处方")
	press,_,_ := GetreadyInfo("prescription", "111222")
	for _,v := range press{
		fmt.Println(v, v.Data)
	}

	fmt.Println("用户111222的药品信息")
	_,transs,_ := GetreadyInfo("transaction", "111222")
	for _,v := range transs{
		fmt.Println(v, v.Data)
	}

	fmt.Println("用户111222的买药信息")
	_,_,buyss := GetreadyInfo("buy", "111222")
	for _,v := range buyss{
		fmt.Println(v, v.Data)
	}

}


var d1 = Dose{Cname:"cid1",Mname:[]string{"mid1"}}
var d2 = Dose{Cname:"cid2",Mname:[]string{"mid2", "mid3"}}
var d3 = Dose{Cname:"cid3",Mname:[]string{"mid4"}}
var d4 = Dose{Cname:"cid4",Mname:[]string{"mid5","mid6","mid7"}}
var d5 = Dose{Cname:"cid5",Mname:[]string{"mid8"}}
var d6 = Dose{Cname:"cid6",Mname:[]string{"mid9"}}
var d7 = Dose{Cname:"cid7",Mname:[]string{"mid10", "mid11"}}
var d8 = Dose{Cname:"cid8",Mname:[]string{"mid12", "mid13", "mid14"}}
var d9 = Dose{Cname:"cid9",Mname:[]string{"mid15"}}
var d10 = Dose{Cname:"cid10",Mname:[]string{"mid16", "mid17"}}


var doses1 = []Dose{d1,d2,d3,d4,d5,d6,d7,d8,d9,d10}
var testStore1 = Drugstore{
	Name: "管星大药房",
	ID: "sid1",
	Location: "上海市杨浦区邯郸路666号",
	Attrs:    "cid1 cid2 cid3 cid4 cid5 cid6 cid7 cid8 cid9 cid10 rid1 sid1",
	Doses:    doses1,
}


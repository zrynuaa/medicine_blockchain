package main

import (
	"github.com/zrynuaa/medicine_blockchain/backend/server"
)


func main()  {
	//todo 进入fabric目录 make restart

	//based.Setup()		//整个系统只运行一次

	finish := make(chan bool)

	//开始启动节点
	//hospital := &server.Peer{Typ:1, Hospital:Hospital1, Port:"8880"}
	//controller := &server.Peer{Typ:3, Controller:Controller, Port: "8884"}
	store1 := &server.Peer{Typ:2, Store:Store1, Port:"8881"}
	//store2 := &server.Peer{Typ:2, Store:Store2, Port:"8882"}
	//store3 := &server.Peer{Typ:2, Store:Store3, Port:"8883"}

	var peers []*server.Peer
	//peers = append(peers, hospital, controller, store1, store2, store3)
	peers = append(peers, store1)

	for _,v := range peers{
		server.Run(v)
	}

	<-finish

}

var d1 = server.Dose{Cname:"cid1",Mname:[]string{"mid1"}}
var d2 = server.Dose{Cname:"cid2",Mname:[]string{"mid2", "mid3"}}
var d3 = server.Dose{Cname:"cid3",Mname:[]string{"mid4"}}
var d4 = server.Dose{Cname:"cid4",Mname:[]string{"mid5","mid6","mid7"}}
var d5 = server.Dose{Cname:"cid5",Mname:[]string{"mid8"}}
var d6 = server.Dose{Cname:"cid6",Mname:[]string{"mid9"}}
var d7 = server.Dose{Cname:"cid7",Mname:[]string{"mid10", "mid11"}}
var d8 = server.Dose{Cname:"cid8",Mname:[]string{"mid12", "mid13", "mid14"}}
var d9 = server.Dose{Cname:"cid9",Mname:[]string{"mid15"}}
var d10 = server.Dose{Cname:"cid10",Mname:[]string{"mid16", "mid17"}}


var doses1 = []server.Dose{d1,d2,d3,d4,d5,d6,d7,d8,d9,d10}
var Store1 = server.Drugstore{
	Name: "管星大药房",
	ID: "sid1",
	Location: "上海市杨浦区邯郸路666号",
	Attrs:    "cid1 cid2 cid3 cid4 cid5 cid6 cid7 cid8 cid9 cid10 rid1 sid1",
	Doses:    doses1,
}

var doses2 = []server.Dose{d1,d3,d5,d7,d9}
var Store2 = server.Drugstore{
	Name: "如意小药店",
	ID: "sid2",
	Location: "上海市浦东新区张横路888号",
	Attrs: "cid1 cid3 cid5 cid7 cid9 cid11 cid12 rid1 sid2",
	Doses: doses2,
}

var doses3 = []server.Dose{d1,d2,d3,d4,d5,d6,d7,d9}
var Store3 = server.Drugstore{
	Name:     "泽宁大药店",
	ID: "sid3",
	Location: "上海市徐汇区枫林路188号",
	Attrs:    "cid1 cid2 cid3 cid4 cid5 cid6 cid7 cid9 rid1 sid3",
	Doses:    doses1,
}

var Hospital1 = server.Hospital{
	Name: "zhongshan",
	ID: "hid1",
	Location: "上海市徐汇区枫林路188号",
	Attrs: "hid1",
}

var Hospital2 = server.Hospital{
	Name: "huashan",
	ID: "hid2",
	Location: "上海市静安区乌鲁木齐中路12号",
	Attrs: "hid2",
}

var Controller = server.Controller{
	Attrs:"cid1 cid2 cid3 cid4 cid5 cid6 cid7 cid8 cid9 cid10 rid1 sid1 sid2 sid3 hid1 hid2",
}
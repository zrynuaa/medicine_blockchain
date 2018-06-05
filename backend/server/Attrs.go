package server

func SetStore1Attrs() Drugstore {
	var drugstore Drugstore

	attrs := []string{"Cid1","Cid2", "Cid8","Rid1"}
	d1 := &Dose{Cname:"Cid1",Mname:[]string{"Mid1", "Mid2"}}
	d2 := &Dose{Cname:"Cid2",Mname:[]string{"Mid8", "Mid9"}}
	d3 := &Dose{Cname:"Cid8",Mname:[]string{"Mid66", "Mid88"}}

	drugstore.Name = "如意大药房"
	drugstore.Location = "上海市浦东新区张横路888号"
	drugstore.Attrs = attrs
	drugstore.Doses = append(drugstore.Doses, d1, d2, d3)

	return drugstore
}


func SetStore2Attrs() Drugstore {
	var drugstore Drugstore

	attrs := []string{"Cid1","Cid8","Cid9","Rid1"}
	d1 := &Dose{Cname:"Cid1",Mname:[]string{"Mid1", "Mid2", "Mid3"}}
	d2 := &Dose{Cname:"Cid9",Mname:[]string{"Mid4", "Mid5"}}
	d3 := &Dose{Cname:"Cid8",Mname:[]string{"Mid66", "Mid88"}}

	drugstore.Name = "管大星药店"
	drugstore.Location = "上海市杨浦区邯郸路666号"
	drugstore.Attrs = attrs
	drugstore.Doses = append(drugstore.Doses, d1, d2, d3)

	return drugstore
}

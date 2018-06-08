package server

func SetStore1Attrs() Drugstore {
	var drugstore Drugstore

	attrs := []string{"cid1","cid2", "cid8","rid1"}
	d1 := &Dose{Cname:"cid1",Mname:[]string{"mid1", "mid2"}}
	d2 := &Dose{Cname:"cid2",Mname:[]string{"mid8", "mid9"}}
	d3 := &Dose{Cname:"cid8",Mname:[]string{"mid66", "mid88"}}

	drugstore.Name = "如意大药房"
	drugstore.Location = "上海市浦东新区张横路888号"
	drugstore.Attrs = attrs
	drugstore.Doses = append(drugstore.Doses, d1, d2, d3)

	return drugstore
}


func SetStore2Attrs() Drugstore {
	var drugstore Drugstore

	attrs := []string{"cid1","cid8","cid9","rid1"}
	d1 := &Dose{Cname:"cid1",Mname:[]string{"mid1", "mid2", "mid3"}}
	d2 := &Dose{Cname:"cid9",Mname:[]string{"mid4", "mid5"}}
	d3 := &Dose{Cname:"cid8",Mname:[]string{"mid66", "mid88"}}

	drugstore.Name = "管大星药店"
	drugstore.Location = "上海市杨浦区邯郸路666号"
	drugstore.Attrs = attrs
	drugstore.Doses = append(drugstore.Doses, d1, d2, d3)

	return drugstore
}

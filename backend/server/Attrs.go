package server

func SetStore1Attrs() Drugstore {
	var drugstore Drugstore

	attrs := []string{"cid1","cid2","cid3","cid4","cid5","cid6","cid7","cid8","cid9","cid10","rid1"}
	d1 := &Dose{Cname:"cid1",Mname:[]string{"mid1"}}
	d2 := &Dose{Cname:"cid2",Mname:[]string{"mid2", "mid3"}}
	d3 := &Dose{Cname:"cid3",Mname:[]string{"mid4"}}
	d4 := &Dose{Cname:"cid4",Mname:[]string{"mid5","mid6","mid7"}}
	d5 := &Dose{Cname:"cid5",Mname:[]string{"mid8"}}
	d6 := &Dose{Cname:"cid6",Mname:[]string{"mid9"}}
	d7 := &Dose{Cname:"cid7",Mname:[]string{"mid10", "mid11"}}

	drugstore.Name = "管星大药房"
	drugstore.Location = "上海市杨浦区邯郸路666号"
	drugstore.Attrs = attrs
	drugstore.Doses = append(drugstore.Doses, d1, d2, d3,d4,d5,d6,d7)

	return drugstore
}

//
//func SetStore2Attrs() Drugstore {
//	var drugstore Drugstore
//
//	attrs := []string{"cid1","cid3","cid5","cid7","cid9","cid11","cid12","rid1"}
//	d1 := &Dose{Cname:"cid1",Mname:[]string{"mid1"}}
//	d3 := &Dose{Cname:"cid3",Mname:[]string{"mid4"}}
//	d5 := &Dose{Cname:"cid5",Mname:[]string{"mid8"}}
//	d7 := &Dose{Cname:"cid7",Mname:[]string{"mid10","mid12"}}
//
//	drugstore.Name = "如意小药店"
//	drugstore.Location = "上海市浦东新区张横路888号"
//	drugstore.Attrs = attrs
//	drugstore.Doses = append(drugstore.Doses, d1, d3, d5,d7)
//
//	return drugstore
//}
//
//func SetStore3Attrs() Drugstore {
//	var drugstore Drugstore
//
//	attrs := []string{"cid1","cid2","cid3","cid4","cid5","cid6","cid7","cid8","cid9","cid10","rid1"}
//	d1 := &Dose{Cname:"cid1",Mname:[]string{"mid1"}}
//	d2 := &Dose{Cname:"cid2",Mname:[]string{"mid2", "mid3"}}
//	d3 := &Dose{Cname:"cid3",Mname:[]string{"mid4"}}
//	d4 := &Dose{Cname:"cid4",Mname:[]string{"mid5","mid6","mid7"}}
//	d5 := &Dose{Cname:"cid5",Mname:[]string{"mid8"}}
//	d6 := &Dose{Cname:"cid6",Mname:[]string{"mid9"}}
//	d7 := &Dose{Cname:"cid7",Mname:[]string{"mid10", "mid11"}}
//
//	drugstore.Name = "泽宁大药店"
//	drugstore.Location = "枫林路825号"
//	drugstore.Attrs = attrs
//	drugstore.Doses = append(drugstore.Doses, d1, d2, d3,d4,d5,d6,d7)
//
//	return drugstore
//}
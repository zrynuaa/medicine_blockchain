package server

func SetStore1Attrs() *Drugstore {
	drugstore := new(Drugstore)

	attrs := []string{"Cid1","Cid2", "Cid8","Cid9","Rid1"}
	d1 := &Dose{Cname:"cname1",Mname:[]string{"mname1", "mname2"}}
	d2 := &Dose{Cname:"cname2",Mname:[]string{"mname8", "mname9"}}

	drugstore.Name = " "
	drugstore.Location = " "
	drugstore.Attrs = attrs
	drugstore.Doses = append(drugstore.Doses, d1, d2)

	return drugstore
}


func SetStore2Attrs() *Drugstore {
	drugstore := new(Drugstore)

	attrs := []string{"Cid1","Cid8","Cid9","Rid1"}
	d1 := &Dose{Cname:"cname1",Mname:[]string{"mname1", "mname2", "mname3"}}
	d2 := &Dose{Cname:"cname3",Mname:[]string{"mname4", "mname5"}}

	drugstore.Name = " "
	drugstore.Location = " "
	drugstore.Attrs = attrs
	drugstore.Doses = append(drugstore.Doses, d1, d2)

	return drugstore
}

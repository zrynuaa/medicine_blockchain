package server

import (
	"github.com/scottocs/medicine_blockchain/backend/based"
	"encoding/json"
	"crypto/md5"
	"fmt"
)

func PrescriptiontoTransaction(pre HospitalPrescription)  {
	buf,_ := json.Marshal(pre)
	digest := md5.Sum(buf)
	buf = digest[:]

	var ptot based.Presciption
	ptot.Type = 0
	ptot.Presciption_id = string(buf)
	ptot.Hospital_id = pre.Hospital_id
	ptot.Patient_id = pre.Patient_id
	ptot.Ts = pre.Ts
	ptot.Ishandled = false
	ptot.Policy = pre.Policy

	num := len(pre.Chemistrys)
	ptot.Data = new(based.Data_pre)
	for i:=0; i<num ;i++{
		ptot.Data.Doctor_id = pre.Doctor_id
		ptot.Data.Disease = pre.Disease
		ptot.Data.Chemistry_name = pre.Chemistrys[i].Chemistry_name
		ptot.Data.Amount = pre.Chemistrys[i].Amount

		based.PutPrescription(ptot)
	}

	fmt.Println(based.GetPrescriptionByid(ptot.Patient_id))
}

func Get()  {
	
}




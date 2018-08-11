package based

import (
	"testing"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)


//////////////////////////////////////////////////////////////////////////////////////////////////////
//新增的test
func TestSplitBytesbyn(t *testing.T) {
	Name = "czn"
	db,_ = leveldb.OpenFile("./db/" + Name + ".db", nil)
	Setup()
	go TimingAccess()
	var one = new(Prescription)
	one.Patient_id = "pat1"
	one.Data = new(Data_pre)
	one.Data.Chemistry_name = "che1"
	one.Data.Amount = 100
	one.Data.Disease = "aizheng"
	one.Data.Doctor_id = "doc1"
	one.Prescription_id = "pre1"
	one.Hospital_id = "hos1"
	one.Ts = 4862
	one.Policy = "SDA ADN DA"
	one.Type = 0
	tid, err := PutIntoFabric("0", "123456", one.Serialize())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tid)
	QuickAccess()
	res,err :=GetFromDbById("prescription","123456")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(deserializePrescription(res))
	fil := make(map[string]string)
	fil["patid"] = "pat1"
	res2, err :=GetPreFromDbByFilter(fil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res2)
	fil2 := make(map[string]string)
	fil2["patid"] = "pat2"
	res3, err :=GetPreFromDbByFilter(fil2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res3)
}
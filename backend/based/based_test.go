package based

import (
	"testing"
	"fmt"
)

func TestMatch(t *testing.T) {
	if match([]string{"cname2","region1"}, "hid1 OR (cname1 AND region1)"){
		fmt.Println("done!")
	}else {
		fmt.Println("wrong!")
	}
}

func TestPut(t *testing.T) {
	var a = Presciption{0,"pre1_1","hid1","pat123456789",1,
	&Data_pre{"czn","feiyan","zyfxs",1},false,"hid1 OR (cname1 AND region1)"}
	PutPrescription(a)

	var b = Presciption{0,"pre1_2","hid1","pat123456789",1,
	&Data_pre{"czn","feiyan","sads",1},false,"hid1 OR (cname2 AND region1)"}
	PutPrescription(b)

	var c = Presciption{0,"pre2_1","hid1","pat234567891",1,
		&Data_pre{"czn","feiyan","zyfxs",1},false,"hid1 OR (cname2 AND region1)"}
	PutPrescription(c)

	var d = Transaction{1,"pat123456789",&Data_tran{"pre1_1","zyfxs",1,2,"sad",21.13,false}}
	PutTransaction(d)

	var e = Transaction{1,"pat123456789",&Data_tran{"pre1_1","zyfxs",1,2,"wqe",21.13,false}}
	PutTransaction(e)
}

func TestGet(t *testing.T) {
	a := GetPrescriptionByid("hid1")
	fmt.Println("HID1:::::::::::::::::::::::")
	for i,pre := range a {
		fmt.Println(i,":")
		fmt.Printf("%v\n", pre)
		fmt.Printf("%v\n", pre.Data)
	}

	b := GetPrescriptionByid("pat123456789")
	fmt.Println("PAT1:::::::::::::::::::::::")
	for i,pre := range b {
		fmt.Println(i,":")
		fmt.Printf("%v\n", pre)
		fmt.Printf("%v\n", pre.Data)
	}

	c := GetPrescriptionByattr([]string{"cname2","region1"})
	fmt.Println("YAODIAN:::::::::::::::::::::::")
	for i,pre := range c {
		fmt.Println(i,":")
		fmt.Printf("%v\n", pre)
		fmt.Printf("%v\n", pre.Data)
	}

	d := GetTransactionByid("pat123456789")
	fmt.Println("PAT1:::::::::::::::::::::::")
	for i,tran := range d {
		fmt.Println(i,":")
		fmt.Printf("%v\n", tran)
		fmt.Printf("%v\n", tran.Data)
	}

	Update("pre1_1")

	e := GetTransactionByid("pat123456789")
	fmt.Println("PAT2:::::::::::::::::::::::")
	for i,tran := range e {
		fmt.Println(i,":")
		fmt.Printf("%v\n", tran)
		fmt.Printf("%v\n", tran.Data)
	}

	f := GetPrescriptionByid("pat123456789")
	fmt.Println("PAT3:::::::::::::::::::::::")
	for i,pre := range f {
		fmt.Println(i,":")
		fmt.Printf("%v\n", pre)
		fmt.Printf("%v\n", pre.Data)
	}

	g := GetPrescriptionBypreid("pre1_1")
	fmt.Println("PAT4:::::::::::::::::::::::")
	fmt.Printf("%v\n", g)
	fmt.Printf("%v\n", g.Data)
}

func TestDose(t *testing.T) {
	PutDose(Dose{"qwee", "asdd", 1, 1.1})

	a, b := GetDosedata("qwee", "asdd", 5)
	fmt.Println(a)
	fmt.Println(b)
}
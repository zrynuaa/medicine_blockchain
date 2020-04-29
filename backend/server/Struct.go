package server

import "github.com/zrynuaa/medicine_blockchain/backend/based"

type Chemistry struct {
	Chemistry_name string `json:"chemistry_name"` //化学名
	Amount         int    `json:"amount"`         //剂量
}

//处方信息
type HospitalPrescription struct {
	Hospital_id string      `json:"hospital_id"` //医院id
	Patient_id  string      `json:"patient_id"`  //病人id
	Doctor_id   string      `json:"doctor_id"`   //医生id
	Disease     string      `json:"disease"`     //病
	Chemistrys  []Chemistry `json:"chemistrys"`  //开药
	Policy      string      `json:"policy"`      //加密policy
}

//化学名与药品关系
type Dose struct {
	Cname string
	Mname []string
}

//药店属性
type Drugstore struct {
	Name     string
	ID       string
	Location string
	Attrs    string
	Doses    []Dose
}

type Hospital struct {
	Name     string
	ID       string
	Location string
	Attrs    string
}

type Controller struct {
	Attrs string
}

//药品信息
type Transaction struct {
	Transaction_id string           `json:"transaction_id"`
	Patient_id     string           `json:"patient_id"`
	Data           *based.Data_tran `json:"data"`
	Ishandled      int              `json:"ishandled"` //0药店未发布, 1药店已发布, 2处方已完成, 3该药店卖的
}

type Prescription struct {
	Data  *based.Prescription `json:"data"`
	Isbuy int                 `json:"isbuy"` //0处方为处理, 1处方已处理
}

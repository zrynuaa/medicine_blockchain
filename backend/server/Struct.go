package server

type Chemistry struct {
	Chemistry_name string 	`json:"chemistry_name"` //化学名
	Amount int 				`json:"amount"`			//剂量
}

//处方信息
type HospitalPrescription struct {
	Hospital_id string 		`json:"hospital_id"`		//医院id
	Patient_id string 		`json:"patient_id"`		//病人id
	Doctor_id string 		`json:"doctor_id"`		//医生id
	Disease string			`json:"disease"`//病
	Chemistrys []Chemistry	`json:"chemistrys"`//开药
	Policy string			`json:"policy"`//加密policy
}

type Dose struct{
	Cname string
	Mname []string
}

type Drugstore struct {
	Name string
	Location string
	Attrs []string
	Doses []*Dose
}

type Data_tran struct {
	Presciption_id string `json:"presciption_id"`
	Medicine_name string `json:"medicine_name"`
	Amount int `json:"amount"`
	Ts uint64 `json:"ts"`
	Site string `json:"site"`
	Price float32 `json:"price"`
	Ishandled bool
}

type Transaction struct {
	Patient_id string
	Data *Data_tran
}
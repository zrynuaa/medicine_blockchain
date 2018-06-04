package server

type Chemistry struct {
	Chemistry_name string 	`json:"chemistry_name"` //化学名
	Amount int 				`json:"amount"`			//剂量
}

//处方信息
type HospitalPrescription struct {
	//Type int 				`json:"type"`				//信息类型
	//TODO 定义一个处方ID生成规则
	//Presciption_id string 	`json:"presciption_id"`	//处方id
	Hospital_id string 		`json:"hospital_id"`		//医院id
	Patient_id string 		`json:"patient_id"`		//病人id
	Ts uint64 				`json:"ts"`				//时间戳
	Doctor_id string 		`json:"doctor_id"`		//医生id
	Disease string			`json:"disease"`//病
	Chemistrys []Chemistry	`json:"chemistrys"`//开药
	//Ishandled bool 			`json:"ishandled"`//是否已经处理
	Policy string			`json:"policy"`//加密policy
}

type Drugstore struct {
	Name string
	location string
	attrs []string
}
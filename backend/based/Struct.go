package based

import (
	"encoding/gob"
	"bytes"
	"log"
)

type Data_pre struct {
	Doctor_id string `json:"doctor_id"`
	Disease string `json:"disease"`
	Chemistry_name string `json:"chemistry_name"`
	Amount int `json:"amount"`
}

type Data_tran struct {
	Presciption_id string `json:"presciption_id"`
	Medicine_name string `json:"medicine_name"`
	Amount int `json:"amount"`
	Ts uint64 `json:"ts"`
	Site string `json:"site"`
	Price float32 `json:"price"`
}

type Data_buy struct {
	Medicine_name string `json:"medicine_name"`
	Medicine_amount int `json:"medicine_amount"`
	Medicine_price float32 `json:"medicine_price"`
	Presciption_id string `json:"presciption_id"`
	Site string `json:"site"`
	Ts uint64 `json:"ts"`
}

type Presciption struct {
	Type int `json:"type"`
	Presciption_id string `json:"presciption_id"`
	Hospital_id string `json:"hospital_id"`
	Patient_id string `json:"patient_id"`
	Ts uint64 `json:"ts"`
	Data *Data_pre `json:"data"`
	Policy string `json:"policy"`
}

type Transaction struct {
	Type int `json:"type"`
	Transaction_id string `json:"transaction_id"`
	Patient_id string `json:"patient_id"`
	Data *Data_tran `json:"data"`
}

type Dose struct {
	Medicine_name string
	Chemistry_name string
	Medicine_amount int
	Medicine_price float32
}

type Buy struct {
	Type int `json:"type"`
	Buy_id string `json:"buy_id"`
	Data *Data_buy `json:"data"`
	Patient_id string `json:"patient_id"`
}

//存储的处方结构，data部分可能将来加密
//在serial的时候加密，一样返回[]byte
type presciption struct {
	Type int
	Presciption_id string
	Hospital_id string
	Patient_id string
	Ts uint64
	Data_pre []byte
	Policy string
}

type transaction struct {
	Type int
	Patient_id string
	Data_tran []byte
}

type buy struct {
	Type int
	Data_buy []byte
	Patient_id string
}

func (b *Data_pre)serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func deserializeDatapre(d []byte) *Data_pre {
	dp := new(Data_pre)

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&dp)
	if err != nil {
		log.Panic(err)
	}

	return dp
}

func (b *Data_tran)serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func deserializeDatatran(d []byte) *Data_tran {
	dt := new(Data_tran)

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&dt)
	if err != nil {
		log.Panic(err)
	}

	return dt
}

func (b *Data_buy)serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func deserializeDatabuy(d []byte) *Data_buy {
	dt := new(Data_buy)

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&dt)
	if err != nil {
		log.Panic(err)
	}

	return dt
}

func (b *Presciption)Serialize() []byte {
	var result bytes.Buffer
	temp := new(presciption)

	temp.Data_pre = b.Data.serialize()
	temp.Hospital_id = b.Hospital_id
	temp.Patient_id = b.Patient_id
	temp.Ts = b.Ts
	temp.Type = b.Type
	temp.Presciption_id = b.Presciption_id
	temp.Policy = b.Policy

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(temp)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func deserializePrescription(d []byte) *Presciption {
	dp := new(Presciption)
	dptemp := new(presciption)

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&dptemp)
	if err != nil {
		log.Panic(err)
	}

	dp.Data = deserializeDatapre(dptemp.Data_pre)
	dp.Hospital_id = dptemp.Hospital_id
	dp.Patient_id = dptemp.Patient_id
	dp.Ts = dptemp.Ts
	dp.Type = dptemp.Type
	dp.Presciption_id = dptemp.Presciption_id
	dp.Policy = dptemp.Policy

	return dp
}

func (b *Transaction)Serialize() []byte {
	var result bytes.Buffer
	temp := new(transaction)

	temp.Data_tran = b.Data.serialize()
	temp.Patient_id = b.Patient_id
	temp.Type = b.Type

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(temp)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func deserializeTransaction(d []byte) *Transaction {
	dp := new(Transaction)
	dptemp := new(transaction)

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&dptemp)
	if err != nil {
		log.Panic(err)
	}

	dp.Data = deserializeDatatran(dptemp.Data_tran)
	dp.Patient_id = dptemp.Patient_id
	dp.Type = dptemp.Type

	return dp
}

func (b *Dose)Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func deserializeDose(d []byte) *Dose {
	dt := new(Dose)

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&dt)
	if err != nil {
		log.Panic(err)
	}

	return dt
}

func (b *Buy)Serialize() []byte {
	var result bytes.Buffer
	temp := new(buy)

	temp.Data_buy = b.Data.serialize()
	temp.Patient_id = b.Patient_id
	temp.Type = b.Type

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(temp)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func deserializeBuy(d []byte) *Buy {
	dp := new(Buy)
	dptemp := new(buy)

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&dptemp)
	if err != nil {
		log.Panic(err)
	}

	dp.Data = deserializeDatabuy(dptemp.Data_buy)
	dp.Patient_id = dptemp.Patient_id
	dp.Type = dptemp.Type

	return dp
}
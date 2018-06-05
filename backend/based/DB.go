package based

import (
	"github.com/syndtr/goleveldb/leveldb"
	"strconv"
	"strings"
	//"fmt"
)

//存储处方，输入一个处方
func PutPrescription(a Presciption) {
	db, err := leveldb.OpenFile("./db/Prescription.db", nil)
	if err != nil {
		return
	}
	db2, err := leveldb.OpenFile("./db/Mapping.db", nil)
	if err != nil {
		return
	}

	aserial := a.serialize()
	err = db.Put([]byte(a.Presciption_id), aserial, nil)
	if err != nil {
		return
	}

	//病人
	data, err := db2.Get([]byte(a.Patient_id), nil)
	if err != nil {
		preids := a.Presciption_id
		err = db2.Put([]byte(a.Patient_id), []byte(preids), nil)
		if err != nil {
			return
		}
	} else {
		plus := "," + a.Presciption_id
		data = append(data, []byte(plus)...)
		err = db2.Put([]byte(a.Patient_id), data, nil)
		if err != nil {
			return
		}
	}

	//医院
	data, err = db2.Get([]byte(a.Hospital_id), nil)
	if err != nil {
		preids := a.Presciption_id
		err = db2.Put([]byte(a.Hospital_id), []byte(preids), nil)
		if err != nil {
			return
		}
	} else {
		plus := "," + a.Presciption_id
		data = append(data, []byte(plus)...)
		err = db2.Put([]byte(a.Hospital_id), data, nil)
		if err != nil {
			return
		}
	}
	defer db.Close()
	defer db2.Close()
}

//存储药品，输入药品
func PutTransaction(a Transaction) {
	db, err := leveldb.OpenFile("./db/Transaction.db", nil)
	if err != nil {
		return
	}

	aserial := a.serialize()

	last, err := db.Get([]byte("last"), nil)
	if err != nil {
		//病人id链接放置id
		err = db.Put([]byte(a.Patient_id), []byte("1"), nil)
		if err != nil {
			return
		}
		//放置id链接药方信息
		err = db.Put([]byte("1"), []byte(aserial), nil)
		if err != nil {
			return
		}
		//last链接最后的放置id
		err = db.Put([]byte("last"), []byte("1"), nil)
		if err != nil {
			return
		}
	} else {
		//last链接最后的放置id
		no, err := strconv.Atoi(string(last))
		if err != nil {
			return
		}
		plus := strconv.Itoa(no + 1)
		err = db.Put([]byte("last"), []byte(plus), nil)
		if err != nil {
			return
		}
		//放置id链接药方信息
		err = db.Put([]byte(plus), []byte(aserial), nil)
		if err != nil {
			return
		}
		//病人id链接放置id
		data, err := db.Get([]byte(a.Patient_id), nil)
		if err != nil {
			err = db.Put([]byte(a.Patient_id), []byte(plus), nil)
			if err != nil {
				return
			}
		} else {
			data = append(data, []byte(","+plus)...)
			err = db.Put([]byte(a.Patient_id), []byte(data), nil)
			if err != nil {
				return
			}
		}
	}
	defer db.Close()
}


//存储剂量信息
func PutDose(a Dose) {
	db, err := leveldb.OpenFile("./db/Dose.db", nil)
	if err != nil {
		return
	}

	aserial := a.serialize()
	last, err := db.Get([]byte("last"), nil)
	if err != nil {
		//放置id链接剂量信息
		err = db.Put([]byte("1"), []byte(aserial), nil)
		if err != nil {
			return
		}
		//last链接最后的放置id
		err = db.Put([]byte("last"), []byte("1"), nil)
		if err != nil {
			return
		}
	} else {
		//last链接最后的放置id
		no, err := strconv.Atoi(string(last))
		if err != nil {
			return
		}
		plus := strconv.Itoa(no + 1)
		err = db.Put([]byte("last"), []byte(plus), nil)
		if err != nil {
			return
		}
		//放置id链接剂量信息
		err = db.Put([]byte(plus), []byte(aserial), nil)
		if err != nil {
			return
		}
	}
	defer db.Close()
}

//判断自己是否已经发布tran
func IsPostdata(presciption_id string, site string, medicine_name string) bool {
	db, err := leveldb.OpenFile("./db/Transaction.db", nil)
	if err != nil {
		return false
	}

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		if string(iter.Key())=="last" || len(iter.Key()) > 8 {
			continue
		}
		//fmt.Printf("%s\n",iter.Key())
		value := deserializeTransaction(iter.Value())
		if value.Data.Presciption_id == presciption_id && value.Data.Site == site && value.Data.Medicine_name == medicine_name{
			return true
		}
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		return false
	}
	return false
}

//获取剂量信息
func GetDosedata(medicine_name string, chemistry_name string, chemistry_amount int) (int,float32) {
	//一个剂量化学名对应的药品剂量
	var mamount int = 0
	//一剂量药品单价
	var mprice float32 = 0.0
	db, err := leveldb.OpenFile("./db/Dose.db", nil)
	if err != nil {
		return 0,0.0
	}

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		if string(iter.Key()) == "last" {
			continue
		}
		value := deserializeDose(iter.Value())
		if value.Chemistry_name == chemistry_name && value.Medicine_name == medicine_name {
			mamount = value.Medicine_amount
			mprice = value.Medicine_price
			break
		}
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		return 0,0.0
	}

	defer db.Close()
	return chemistry_amount * mamount, float32(chemistry_amount) * float32(mamount) * mprice
}


//根据病人ID或者医院ID查相关药品信息
func GetPrescriptionByid(id string) []*Presciption {
	var result []*Presciption

	db, err := leveldb.OpenFile("./db/Prescription.db", nil)
	if err != nil {
		return nil
	}
	db2, err := leveldb.OpenFile("./db/Mapping.db", nil)
	if err != nil {
		return nil
	}

	data, err := db2.Get([]byte(id), nil)
	if err != nil {
		return nil
	}

	pres := strings.Split(string(data), ",")
	for _, pre := range pres {
		re, err := db.Get([]byte(pre), nil)
		//fmt.Printf("%s\n",re)
		if err != nil {
			return nil
		}

		result = append(result, deserializePrescription(re))
	}
	defer db.Close()
	defer db2.Close()
	return result
}

//根据详细处方ID查相关药品信息
func GetPrescriptionBypreid(preid string) *Presciption {
	db, err := leveldb.OpenFile("./db/Prescription.db", nil)
	if err != nil {
		return nil
	}

	data, err := db.Get([]byte(preid), nil)
	if err != nil {
		return nil
	}

	defer db.Close()
	return deserializePrescription(data)
}

//根据简略处方ID查相关药品信息
func GetPrescriptionByeasypreid(easypreid string) []*Transaction {
	var result []*Transaction
	var temp []string
	
	db, err := leveldb.OpenFile("./db/Prescription.db", nil)
	if err != nil {
		return nil
	}	
	
	db2, err := leveldb.OpenFile("./db/Transaction.db", nil)
	if err != nil {
		return nil
	}

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		value := deserializePrescription(iter.Value())
		if value.Presciption_id[:16] == easypreid {
			temp = append(temp, value.Presciption_id)
		}
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		return nil
	}
	
	iter = db2.NewIterator(nil, nil)
	for iter.Next() {
		if string(iter.Key())=="last" || len(iter.Key()) > 8 {
			continue
		}
		value := deserializeTransaction(iter.Value())
		if isexist(temp, value.Data.Presciption_id) {
			result = append(result, value)
		}
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		return nil
	}

	defer db.Close()
	return result
}

//根据病人ID查相关药品信息
func GetTransactionByid(id string) []*Transaction {
	var result []*Transaction

	db, err := leveldb.OpenFile("./db/Transaction.db", nil)
	if err != nil {
		return nil
	}

	data, err := db.Get([]byte(id), nil)
	if err != nil {
		return nil
	}

	pres := strings.Split(string(data), ",")
	for _, pre := range pres {
		re, err := db.Get([]byte(pre), nil)
		//fmt.Printf("%s\n",re)
		if err != nil {
			return nil
		}

		result = append(result, deserializeTransaction(re))
	}
	defer db.Close()
	return result
}

//输入处方id，将对应处方设置成已处理
func Update(presciption_id string) {
	db, err := leveldb.OpenFile("./db/Prescription.db", nil)
	if err != nil {
		return
	}

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		value := deserializePrescription(iter.Value())
		if value.Presciption_id == presciption_id {
			value.Ishandled = true
			err = db.Put([]byte(presciption_id), value.serialize(), nil)
			if err != nil {
				return
			}
			break
		}
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		return
	}

	defer db.Close()
}

//根据药店属性获取对应处方
func GetPrescriptionByattr(attr []string) []*Presciption {
	var result []*Presciption

	db, err := leveldb.OpenFile("./db/Prescription.db", nil)
	if err != nil {
		return nil
	}

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		value := deserializePrescription(iter.Value())
		if match(attr, value.Policy) {
			result = append(result, value)
		}
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		return nil
	}

	defer db.Close()
	return result
}


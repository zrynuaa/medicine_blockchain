package based

import (
	//"fmt"
	//"github.com/Doresimon/SM-Collection/SM3"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb/util"
)



//输入类型what，0代表pre，1代表tran，2代表buy，id和value
func PutIntoFabric(what string, id string, value []byte) (string,error){
	return putInfo(what, id, string(value))
}

//types有prescription，transaction，buy
func GetFromDbById(types string, id string) ([]byte, error){
	key := append(commandToBytes(types), []byte(id)...)
	data, err := db.Get(key, nil)
	if err != nil {
		return []byte(""), fmt.Errorf("getFromDbById error, type:%s, id:%s\n", types, id)
	}
	return data,nil
}

//从db中获取pre信息，filter为map，示例见test，下同
func GetPreFromDbByFilter(fil map[string]string) ([]*Prescription, error){
	var result []*Prescription
	var flag bool
	all, err := getAllFromDb("prescription")
	if err != nil {
		return nil, fmt.Errorf("getPreFromDbByFilter error! %s", err)
	}

	for _, one := range all {
		temp := deserializePrescription(one)
		flag = true
		if fil != nil {
			for k, v := range fil {
				if (k == "patid" && v == temp.Patient_id) || (k == "hosid" && v == temp.Hospital_id) {
					continue
				}
				flag = false
			}
		}

		if flag {
			result = append(result, temp)
		}
	}
	return result, nil
}

func GetTraFromDbByFilter(fil map[string]string) ([]*Transaction, error){
	var result []*Transaction
	var flag bool
	all, err := getAllFromDb("transaction")
	if err != nil {
		return nil, fmt.Errorf("getTraFromDbByFilter error! %s", err)
	}

	for _, one := range all {
		temp := deserializeTransaction(one)
		flag = true
		if fil != nil {
			for k, v := range fil {
				if (k == "preid" && v == temp.Data.Prescription_id) || (k == "patid" && v == temp.Patient_id) ||
					(k == "site" && v == temp.Data.Site) || (k == "medicine" && v == temp.Data.Medicine_name) {
					continue
				}
				flag = false
			}
		}

		if flag {
			result = append(result, temp)
		}
	}
	return result, nil
}

func GetBuyFromDbByFilter(fil map[string]string) ([]*Buy, error){
	var result []*Buy
	var flag bool
	all, err := getAllFromDb("buy")
	if err != nil {
		return nil, fmt.Errorf("getBuyFromDbByFilter error! %s", err)
	}

	for _, one := range all {
		temp := deserializeBuy(one)
		flag = true
		if fil != nil {
			for k, v := range fil {
				if (k == "preid" && v == temp.Data.Prescription_id) || (k == "patid" && v == temp.Patient_id) ||
					(k == "site" && v == temp.Data.Site) || (k == "medicine" && v == temp.Data.Medicine_name) {
					continue
				}
				flag = false
			}
		}

		if flag {
			result = append(result, temp)
		}
	}
	return result, nil
}

//获取剂量信息
func GetDoseFromDb(medicine_name string, chemistry_name string, chemistry_amount int) (int,float32,error) {
	//一个剂量化学名对应的药品剂量
	var mamount int = 0
	//一剂量药品单价
	var mprice float32 = 0.0

	all, err := getAllFromDb("dose")
	if err != nil {
		return 0, 0.0, fmt.Errorf("getDoseFromDbByFilter error! %s", err)
	}

	for _, one := range all {
		value := deserializeDose(one)
		if value.Chemistry_name == chemistry_name && value.Medicine_name == medicine_name {
			mamount = value.Medicine_amount
			mprice = value.Medicine_price
			break
		}
	}

	return chemistry_amount * mamount, float32(chemistry_amount) * float32(mamount) * mprice, nil
}

func getAllFromDb(types string) ([][]byte, error) {
	var result [][]byte
	iter := db.NewIterator(util.BytesPrefix(commandToBytes(types)), nil)
	for iter.Next() {
		result = append(result, iter.Value())
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return nil, fmt.Errorf("getAllFromDb error, types:%s", types)
	}
	return result, nil
}

func commandToBytes(command string) []byte {
	var bytess [commandLength]byte

	for i, c := range command {
		bytess[i] = byte(c)
	}

	return bytess[:]
}

//将解密后的信息存到db中，应在abe解密之后调用，types为presciption、transaction、buy等
func PutIntoDb(types string, id string, value []byte) error{
	var key []byte
	key = append(commandToBytes(types), []byte(id)...)
	err := db.Put(key, value, nil)
	if err != nil {
		return fmt.Errorf("PutIntoDb error!%s", err)
	}
	return nil
}

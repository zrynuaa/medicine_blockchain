package based

import (
	"fmt"
	"time"
	"strconv"
	"github.com/zrynuaa/cpabe06_client/bswabe"
)

//初始化，只在一开始调用一次
func Setup() {
	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		return
	}
	// Close SDK
	//defer fSetup.CloseSDK()

	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC([]string{"init"})
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
		return
	}
}

//对应cc中的putinfo
func putInfo(what string, id string, value string) (string, error){
	trcid, err := fSetup.Invoke([]string{"invoke", "putinfo", what, id, value})
	if err != nil {
		return "", fmt.Errorf("invoke putinfo error!:%v", err)
	}
	return trcid, nil
}

//对应cc中的getids
func getIds(what,id string) (string, error){
	payload, err := fSetup.Query([]string{"invoke", "getids", what, id})
	if err != nil {
		return "", fmt.Errorf("query getids error!:%v", err)
	}
	return payload, nil
}

//对应cc中的getpres
func getPres(what,id string) (string, error){
	payload, err := fSetup.Query([]string{"invoke", "getpres", what, id})
	if err != nil {
		return "", fmt.Errorf("query getpres error!:%v", err)
	}
	return payload, nil
}

//定时获取
func TimingAccess() {
	ticker := time.NewTicker(delta)
	for _ = range ticker.C {
		fmt.Println(time.Now())
		err := QuickAccess()
		if err != nil {
			fmt.Println(err)
		}
	}
}

//快速获取
func QuickAccess() error{
	fmt.Println("Access to Fabric")
	for i := 0; i < 3; i++ {
		err := synchronize(i)
		if err != nil {
			return fmt.Errorf("input type error! need 0, 1 or 2", err)
		}
	}
	return nil
}

//同步cc到db中
func synchronize(what int) error {
	if what < 0 || what > 2 {
		return fmt.Errorf("input type error! need 0, 1 or 2")
	}
	//可适用于lastid为0
	idspayload, err := getIds(strconv.Itoa(what), LastId[what])
	if err != nil {
		return fmt.Errorf("syschronizeLastId error:%s", err)
	}
	prespayload, err := getPres(strconv.Itoa(what), LastId[what])
	if err != nil {
		return fmt.Errorf("syschronizeLastId error:%s", err)
	}
	tempencids := splitStringbyn(idspayload)
	tempencpres := splitStringbyn(prespayload)
	//去掉最后的空字符串
	tempencids = tempencids[:len(tempencids)-1]
	tempencpres = tempencpres[:len(tempencpres)-1]
	//更新lastid
	if len(tempencids) == 0 {
		return nil
	}
	//dec
	var tempids []string
	var temppres [][]byte
	for i, j := range tempencpres {
		tmp := []byte(j)
		if len(tmp) == 0 {
			fmt.Println("=================len(j)=0")
			continue
		}
		result := bswabe.CP_Dec(pk, sk, bswabe.UnSerializeBswabeCphKey(pk, []byte(j)))
		if string(result) == "" {
			continue
		}
		tempids = append(tempids, tempencids[i])
		temppres = append(temppres, result)
	}
	if len(tempids) == 0 {
		return nil
	}
	//update db
	LastId[what] = tempids[len(tempids)-1]
	types := whatmap[what]
	for i, j := range tempids {
		err := PutIntoDb(types, j, temppres[i])
		if err != nil {
			return fmt.Errorf("syschronizeLastId error:%s", err)
		}
	}
	return nil
}

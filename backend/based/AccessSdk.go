package based

import (
	"os"
	"github.com/thorweiyan/fabric_go_sdk"
	"fmt"
	"time"
	"strconv"
)

const delta  = time.Minute
var LastId = []string{"","",""}
var whatmap = map[int]string{
	0: "presciption",
	1: "transaction",
	2: "buy",
}
// Definition of the Fabric SDK properties
var fSetup = fabric_go_sdk.FabricSetup{
	// Network parameters
	OrdererID: "orderer.fudan.edu.cn",
	OrgID:     "org1.fudan.edu.cn",

	// Channel parameters
	ChannelID:     "fudanfabric",
	ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/thorweiyan/fabric_go_sdk/fixtures/artifacts/fudanfabric.channel.tx",

	// Chaincode parameters
	ChainCodeID:      "fudancc",
	ChaincodeGoPath:  os.Getenv("GOPATH"),
	ChaincodePath:    "github.com/zrynuaa/medicine_blockchain/backend/chaincode/",
	ChaincodeVersion: "0",
	OrgAdmin:         "Admin",
	OrgName:          "org1",
	ConfigFile:       os.Getenv("GOPATH") + "/src/github.com/thorweiyan/fabric_go_sdk/config.yaml",

	// User parameters
	UserName: "User1",
}

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
	tempids := splitStringbyn(idspayload)
	temppres := splitStringbyn(prespayload)
	//去掉最后的空字符串
	tempids = tempids[:len(tempids)-1]
	temppres = temppres[:len(temppres)-1]
	//更新lastid
	if len(tempids) == 0 {
		return nil
	}
	LastId[what] = tempids[len(tempids)-1]
	types := whatmap[what]
	//update db
	for i, j := range tempids {
		err := PutIntoDb(types, j, []byte(temppres[i]))
		if err != nil {
			return fmt.Errorf("syschronizeLastId error:%s", err)
		}
	}
	return nil
}

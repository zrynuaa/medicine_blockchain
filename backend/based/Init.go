package based

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/thorweiyan/fabric_go_sdk"
	"github.com/zrynuaa/cpabe06_client/bswabe"
	"os"
	"time"
)

var Name string = "default"
var db *leveldb.DB

const commandLength = 12

var LastId = []string{"", "", ""}

const delta = time.Second

var whatmap = map[int]string{
	0: "prescription",
	1: "transaction",
	2: "buy",
}
var sk *bswabe.BswabePrv
var pk *bswabe.BswabePub

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

func Init(name string, pub *bswabe.BswabePub, prv *bswabe.BswabePrv) error {
	if name == "" {
		return fmt.Errorf("name is nothing")
	}
	Name = name
	db, _ = leveldb.OpenFile(os.Getenv("GOPATH")+"/src/github.com/zrynuaa/medicine_blockchain/backend/db/"+Name+".db", nil)
	sk = prv
	pk = pub
	return nil
}

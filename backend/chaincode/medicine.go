package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type System struct {
	prescriptionID []string
	medicineID []string
	buyID []string
}

func (t *System) Init(stub shim.ChaincodeStubInterface) peer.Response {
	function, _ := stub.GetFunctionAndParameters()
	if function != "init" {
		return shim.Error("Unknown function call")
	}
	return shim.Success(nil)
}

func (t *System) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	if fn != "invoke" {
		return shim.Error("Unknown function call")
	}

	// Check whether the number of arguments is sufficient
	if len(args) < 1 {
		return shim.Error("The number of arguments is insufficient.")
	}

	if args[0] == "putinfo" {
		return t.Putinfo(stub, args)
	} else if args[0] == "getids" {
		return t.GetIDs(stub, args)
	} else if args[0] == "getpres" {
		return t.Getpres(stub, args)
	}

	// If the arguments given don’t match any function, we return an error
	return shim.Error("Unknown action, check the first argument")
}

//医院、药店、患者上传处方、卖药、买药信息，args（func, type，ID，value）1处方信息，2卖药信息，3买药信息
func (t *System)Putinfo(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 4 {
		return shim.Error("The number of arguments is insufficient.")
	}

	if args[1] == "0" {
		t.prescriptionID = append(t.prescriptionID, args[2])
	} else if args[1] == "1" {
		t.medicineID = append(t.medicineID, args[2])
	} else if args[1] == "2"{
		t.buyID = append(t.buyID, args[2])
	}

	// t.prescriptionID = append(t.prescriptionID, args[0])
	err := stub.PutState(args[2], []byte(args[3]))
	if err != nil {
		return shim.Error("Failed to put information: " + args[2])
	}
	err = stub.SetEvent("eventInvoke", []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

//传进来本地最新的一个ID，获取这个ID后面所有的新的ID,返回
func (t *System)GetIDs(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	preIDlist, err := t.getids(args)
	if err != nil {
		return shim.Error(err.Error())
	}

	//将[]string转换成[]byte，每个string用"\n"隔开
	var rs []byte
	for _,v := range preIDlist{
		rs = append(rs, []byte(v)...)
		rs = append(rs, []byte("\n\n")...)
	}

	return shim.Success(rs)
}

//传进来本地最新的一个ID，获取这个ID后面所有的新的信息,返回
func (t *System)Getpres(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	preIDlist, err := t.getids(args)
	if err != nil {
		return shim.Error(err.Error())
	}
	var prelist []string

	for _, preid := range preIDlist{
		value, _ := stub.GetState(preid)
		prelist = append(prelist, string(value))
	}

	//将[]string转换成[]byte，每个string用"\n"隔开
	var rs []byte
	for _,v := range prelist{
		rs = append(rs, []byte(v)...)
		rs = append(rs, []byte("\n\n")...)
	}

	return shim.Success(rs)
}

func (t *System)getids(args []string) ([]string, error) {
	if len(args) != 3 {
		return []string{""}, fmt.Errorf("The number of arguments is insufficient.")
	}
	var preIDlist []string
	var list []string
	if args[1] == "0" {
		list = t.prescriptionID
	}else if args[1] == "1" {
		list = t.medicineID
	}else if args[1] == "2" {
		list = t.buyID
	}
	var po= -1
	//传入空id,返回所有id
	if args[2] == "" {
		return list, nil
	}
	for k, v := range list {
		if v == args[2] {
			po = k
			break
		}
	}
	if po == -1 {
		return []string{""}, fmt.Errorf("Do not have this id: %s", args[2])
	}
	preIDlist = list[po+1:]

	return preIDlist, nil
}

func main() {
	if err := shim.Start(new(System)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
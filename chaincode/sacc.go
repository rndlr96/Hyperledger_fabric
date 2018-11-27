package main

import (
  "fmt"

  "github.com/hyperledger/fabric/core/chaincode/shim"
  "github.com/hyperledger/fabric/protos/peer"
)

type SimpleAsset struct {

}

func (t *SimpleAsset) init(stub shim.ChaincodeStubInterface) peer.Response {

  args := stub.GetStringArgs()
  if len(args) != 2 {
    return shim.Error("Incorrect arguments. Expecting a key and a value")
  }

  err := stub.PutState(args[0], []byte(args1[1]))
  if err != nil {
    return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
  }
  return shim.Success(nil)
}

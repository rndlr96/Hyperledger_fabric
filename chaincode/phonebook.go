package main

import (
  "fmt"
  "encoding/json"
  "bytes"
  "strconv"

  "github.com/hyperledger/fabric/core/chaincode/shim"
  "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type Member struct {
  Name    string `json:"name"`
  Number  string `json:"number"`
  Mail    string `json:"mail"`
  Company string `json:"company"`
}

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
  return shim.Success(nil)
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
  function, args := stub.GetFunctionAndParameters()

  if function == "queryMember" {
    return s.queryMember(stub, args)
  } else if function == "initLedger" {
    return s.initLedger(stub)
  } else if function == "createMember" {
    return s.createMember(stub, args)
  } else if function == "queryAllMembers" {
    return s.queryAllMembers(stub)
  }
  // } else if function == "changeNumber" {
  //   return s,changeNumber(stub, args)
  // } else if function == "changeMail" {
  //   return s,changeNumber(stub, args)
  // } else if function == "changeCompamy" {
  //   return s,changeNumber(stub, args)
  // }

  return shim.Error("Invalid Smart Contract function name")
}

func (s *SmartContract) queryMember(stub shim.ChaincodeStubInterface, args []string) peer.Response {

  if len(args) != 1 {
    return shim.Error("Incorrect number of arguments. Expecting 1")
  }

  memberAsBytes, _ := stub.GetState(args[0])
  return shim.Success(memberAsBytes)
}

func (s *SmartContract) initLedger(stub shim.ChaincodeStubInterface) peer.Response {
  members := []Member {
    Member{Name: "Tomoko", Number:"01012341234", Mail:"red@gmail.com", Company:"Toyota"},
    Member{Name: "Brad", Number:"01056785678", Mail:"blue@naver.com", Company:"Ford"},
    Member{Name: "Jin Soo", Number:"01039482938", Mail:"green@gmail.com", Company:"Hyundai"},
    Member{Name: "Max", Number:"01038402844", Mail:"orange@gmail.com", Company:"Volkswagen"},
    Member{Name: "Adriana", Number:"01030683039", Mail:"yello@gmail.com", Company:"Tesla"},
    Member{Name: "Michel", Number:"01030528594", Mail:" black@gmail.com", Company:"Peugeot"},
    Member{Name: "Aarav", Number:"01030593040", Mail:"purple@gmail.com", Company:"Chery"},
    Member{Name: "Pari", Number:"01050294929", Mail:"indigo@gmail.com", Company:"Fiat"},
    Member{Name: "Valeria", Number:"01002049293", Mail:"brown@gmail.com", Company:"Tata"},
    Member{Name: "Shotaro", Number:"01039598294", Mail:"violet@gmail.com", Company:"KIA"},
    Member{Name: "Eron", Number:"01049482849", Mail:"white@gmail.com", Company:"Holden"},
  }

  for i := 0 ; i < len(members) ; i++ {
    fmt.Println("i is ", i)
    memberAsBytes, _ := json.Marshal(members[i])
    stub.PutState("MEMBER"+strconv.Itoa(i), memberAsBytes)
    fmt.Println("Added", members[i])
  }

  return shim.Success(nil)
}

func (s *SmartContract) createMember(stub shim.ChaincodeStubInterface, args []string) peer.Response {
  if len(args) != 5 {
    return shim.Error("Incorrect number of arguments. Expecting 5")
  }

  var member = Member{Name: args[1], Number: args[2], Mail: args[3], Company: args[4]}

  memberAsBytes, _ := json.Marshal(member)
  stub.PutState(args[0], memberAsBytes)

  return shim.Success(nil)
}

func (s *SmartContract) queryAllMembers (stub shim.ChaincodeStubInterface) peer.Response {
  startKey := "MEMBER0"
  endKey := "MEMBER999"

  resultsIterator, err := stub.GetStateByRange(startKey, endKey)

  if err != nil {
    return shim.Error(err.Error())
  }

  defer resultsIterator.Close()

  var buffer bytes.Buffer
  buffer.WriteString("[")

  bArrayMemberAlreadyWritten := false

  for resultsIterator.HasNext() {
    queryResponse, err := resultsIterator.Next()

    if err != nil {
      return shim.Error(err.Error())
    }

    if bArrayMemberAlreadyWritten == true {
      buffer.WriteString(",")
    }

    buffer.WriteString("{\"Key\":")
    buffer.WriteString("\"")
    buffer.WriteString(queryResponse.Key)
    buffer.WriteString("\"")
    buffer.WriteString(", \"Record\":")
    buffer.WriteString(string(queryResponse.Value))
    buffer.WriteString("}")
    bArrayMemberAlreadyWritten = true
  }
  buffer.WriteString("]")

  fmt.Printf("- queryAllMembers:\n%s\n", buffer.String())

  return shim.Success(buffer.Bytes())
}

func main() {
  err := shim.Start(new(SmartContract))
  if err != nil {
    fmt.Printf("Error creating new Smart Contract: %s", err)
  }
}

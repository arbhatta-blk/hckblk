package main

import (
	// store car details in json format in our ledger
	"encoding/json"
	"fmt"
	// shim package from github supplies APIs to access state variables. also allows us transaction related access
	"github.com/hyperledger/fabric/core/chaincode/shim"
	// Allow us to structure the response that we send back to the peers who invoke the chaincode
	"github.com/hyperledger/fabric/protos/peer"
)

// Init , Invoke would be attached to this struct
type CarRegistry struct{
}

type Car struct{
	Make string `json:"make"`
	Model string `json:"model"`
	Colour string `json:"colour"`
	Owner string `json:"owner"`
}

// all chaincodes need to implement standard interfaces Init
// stub will give access to state variables
func (s *CarRegistry) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

// is called for each transaction on the chaincode
func (s *CarRegistry) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// get the function name and the parameters passed. The first argument i.e. args[0] is the unique id of the object 
	// The args[1] is expected to be Make, args[2] Model and so on.
	function, args := stub.GetFunctionAndParameters()

	if function == "queryCar" {
		return s.queryCar(stub, args)
	} else if function == "createCar" {
		return s.createCar(stub, args)
	} else if function == "changeCarOwner" {
		return s.changeCarOwner(stub, args)
	}
	return shim.Error("Invalid function name")
}

func (s *CarRegistry) queryCar(stub shim.ChaincodeStubInterface, args []string) peer.Response { 
	// The State of the blockchain can be cosidered as a key value store and in this case we are retrieving the 0th value
	carAsBytes,_ := stub.GetState(args[0])
	return shim.Success(carAsBytes)
}

// Adds a new entry
func (s *CarRegistry) createCar(stub shim.ChaincodeStubInterface, args []string) peer.Response { 
	var car = Car{Make: args[1], Model: args[2], Colour: args[3], Owner: args[4]}
	carAsBytes,_ := json.Marshal(car)
	stub.PutState(args[0], carAsBytes)
	return shim.Success(nil)
}

// Updates entry
func (s *CarRegistry) changeCarOwner(stub shim.ChaincodeStubInterface, args []string) peer.Response { 
	carAsBytes,_ := stub.GetState(args[0])
	car := Car{}
	json.Unmarshal(carAsBytes, &car)
	car.Owner = args[1]
	carAsBytes,_ = json.Marshal(car)
	stub.PutState(args[0], carAsBytes)
	return shim.Success(nil)
}

// Main method
func main(){
	// Creating a new CarRegistry will run the Init()
	err := shim.Start(new(CarRegistry))
	if err != nil{
		fmt.Printf("Error creating new chaincode: %s", err)
	}
}
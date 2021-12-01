/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package simple

import (
	"fmt"
	"os"

	// "strconv"
	"encoding/json"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

// WorkEx example simple Chaincode implementation
type WorkEx struct{}
type OrgAsset struct {
	Id   string `json:"id"`
	Hash string `json:"hash"`
}

func (t *WorkEx) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Init invoked")
	assetData := OrgAsset{
		Id:   "VNIT_106",
		Hash: "asshdifjs",
	}
	assetBytes, _ := json.Marshal(assetData)
	assetErr := stub.PutState("VNIT_106", assetBytes)
	if assetErr != nil {
		return shim.Error(fmt.Sprintf("Failed to create asset:"))
	}

	return shim.Success(nil)
}

func (t *WorkEx) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	if os.Getenv("DEVMODE_ENABLED") != "" {
		fmt.Println("invoking in devmode")
	}
	function, args := stub.GetFunctionAndParameters()
	switch function {
	case "init":
		// Make payment of X units from A to B
		return t.Init(stub)
	case "invoke":
		// Make payment of X units from A to B
		return t.invoke(stub, args)
	case "delete":
		// Deletes an entity from its state
		return t.delete(stub, args)
	case "query":
		// the old "Query" is now implemented in invoke
		return t.query(stub, args)
	case "history":
		return t.getHistory(stub, args)
	case "update":
		// return with an error
		return t.updateState(stub, args)
	// case "mspid":
	// 	// Checks the shim's GetMSPID() API
	// 	return t.mspid(args)
	// case "event":
	// 	return t.event(stub, args)
	default:
		return shim.Error(`Invalid invoke function name. Expecting "invoke", "delete", "query", "respond", "mspid", or "event"`)
	}
}

// Transaction makes payment of X units from A to B
func (t *WorkEx) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var id, hash string // Entities

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	id = args[0]
	hash = args[1]
	assetData := OrgAsset{
		Id:   id,
		Hash: hash,
	}
	assetBytes, _ := json.Marshal(assetData)
	var err error
	err = stub.PutState(id, assetBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(assetBytes)
}

// Update Existing Chaincode
func (t *WorkEx) updateState(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var id, hash string // Entities

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	id = args[0]
	hash = args[1]
	assetData := OrgAsset{
		Id:   id,
		Hash: hash,
	}
	assetBytes, _ := json.Marshal(assetData)
	var err error

	Avalbytes, err := stub.GetState(id)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to update state for " + id + " as it does not exist\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + id + "\"}"
		return shim.Error(jsonResp)
	}

	err = stub.PutState(id, assetBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Deletes an entity from state
func (t *WorkEx) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

// query callback representing the query of a chaincode
func (t *WorkEx) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := string(Avalbytes)
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

// Get history of asset
//
// Shows Off GetHistoryForKey() - reading complete history of a key/value
//
// ============================================================================================================================
func (c *WorkEx) getHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	type DocHistory struct {
		TxId  string   `json:"txId"`
		Value OrgAsset `json:"value"`
	}
	var history []DocHistory
	var orgAsset OrgAsset

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	assetId := args[0]
	fmt.Printf("- start getHistoryForAsset: %s\n", assetId)

	// Get History
	resultsIterator, err := stub.GetHistoryForKey(assetId)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		historyData, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		var tx DocHistory
		tx.TxId = historyData.TxId
		json.Unmarshal(historyData.Value, &orgAsset)
		tx.Value = orgAsset           //copy orgAsset over
		history = append(history, tx) //add this tx to the list
	}
	fmt.Printf("- getHistoryForAsset returning:\n%s", history)

	//change to array of bytes
	historyAsBytes, _ := json.Marshal(history) //convert to array of bytes
	return shim.Success(historyAsBytes)
}

// // respond simply generates a response payload from the args
// func (t *WorkEx) respond(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	if len(args) != 3 {
// 		return shim.Error("expected three arguments")
// 	}

// 	status, err := strconv.ParseInt(args[0], 10, 32)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
// 	message := args[1]
// 	payload := []byte(args[2])

// 	return pb.Response{
// 		Status:  int32(status),
// 		Message: message,
// 		Payload: payload,
// 	}
// }

// // mspid simply calls shim.GetMSPID() to verify the mspid was properly passed from the peer
// // via the CORE_PEER_LOCALMSPID env var
// func (t *WorkEx) mspid(args []string) pb.Response {
// 	if len(args) != 0 {
// 		return shim.Error("expected no arguments")
// 	}

// 	// Get the mspid from the env var
// 	mspid, err := shim.GetMSPID()
// 	if err != nil {
// 		jsonResp := "{\"Error\":\"Failed to get mspid\"}"
// 		return shim.Error(jsonResp)
// 	}

// 	if mspid == "" {
// 		jsonResp := "{\"Error\":\"Empty mspid\"}"
// 		return shim.Error(jsonResp)
// 	}

// 	fmt.Printf("MSPID:%s\n", mspid)
// 	return shim.Success([]byte(mspid))
// }

// // event emits a chaincode event
// func (t *WorkEx) event(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	if len(args) != 2 {
// 		return shim.Error("Incorrect number of arguments. Expecting 2")
// 	}

// 	if err := stub.SetEvent(args[0], []byte(args[1])); err != nil {
// 		return shim.Error(err.Error())
// 	}

// 	return shim.Success(nil)
// }

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"

)

var logger = shim.NewLogger("PHARMA_TRACK")

type pharma_track struct {

}

type Pharma_Asset struct {
    Asset_ID		string	`json:"asset_id"`
	Asset_Name		string	`json:"Asset_Name"`
	Asset_Batch_no	string 	`json:"asset_batch_no"`
	Asset_Expiry	string	`json:"asset_expiry"`
	Asset_Owner	Owner_Asset	`json:"asset_owner"`
	Asset_Status	string	`json:"asset_status"`
}

type Owner_Asset struct {
	Owner_Name          string		`json:"owner_name"`
	Owner_Id 			string		`json:"owner_id"`
	Address             string      `json:"address"`
}

type CounterNO struct {
	Counter int `json:"counter"`
}

var states = []string{"MANUFACTURED","RAEDYTOSHIP","SHIPPED","RECEIVED"}

func isValidStatus(status string) bool {
    fmt.Println("PASSED status="+status)
    for i := 0; i < len(states); i++ {
        if states[i] == status {
            return true
        }
    }
    return false
}

func (asset *Pharma_Asset) changeOwner(newOwner Owner_Asset) {
	asset.Asset_Owner = newOwner
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(pharma_track))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}

// ============================================================================================================================
// Init - reset all the things
// ============================================================================================================================
func (t *pharma_track) Init(stub shim.ChaincodeStubInterface) pb.Response {
    fmt.Printf("Chain Code Initialzed...")
    logger.Info("Chain Code Initialzed...")
    return shim.Success(nil)
}

//getCounter to the latest value of the counter based on the Asset Type provided as input parameter
func getCounter(APIstub shim.ChaincodeStubInterface, AssetType string) int {
	counterAsBytes, _ := APIstub.GetState(AssetType)
	counterAsset := CounterNO{}

	json.Unmarshal(counterAsBytes, &counterAsset)
	fmt.Sprintf("Counter Current Value %d of Asset Type %s",counterAsset.Counter,AssetType)
        logger.Infof("Counter Current Value %d of Asset Type %s",counterAsset.Counter,AssetType)

	return counterAsset.Counter
}

//incrementCounter to the increase value of the counter based on the Asset Type provided as input parameter by 1
func incrementCounter(APIstub shim.ChaincodeStubInterface,  AssetType string) int {
	counterAsBytes, _ := APIstub.GetState(AssetType)
	counterAsset := CounterNO{}

	json.Unmarshal(counterAsBytes, &counterAsset)
	counterAsset.Counter++
	counterAsBytes, _ = json.Marshal(counterAsset)

	err := APIstub.PutState(AssetType, counterAsBytes)
	if err != nil {

		fmt.Sprintf("Failed to Increment Counter")
                logger.Infof("Failed to Increment Counter for Asset Type %s", AssetType)

	}
	return counterAsset.Counter
}

// queryAsset - retrieve asset from the ledger
func (t *pharma_track) queryAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments, Required 2")
	}

	fmt.Println("In Query Asset")
        logger.Info("In Query Asset...")

	AssetAsBytes, err := stub.GetState(args[1])
	
	if err != nil {
	    var msg string = "Error getting state for asset with id : " + args[1]
	    fmt.Printf(msg)
	    return shim.Error(msg)
	}

	if AssetAsBytes == nil {
		return shim.Error("Could not find Asset with id : " + args[1])

	}

        logger.Infof("Query Response:%s\n", string(AssetAsBytes))
	return shim.Success(AssetAsBytes)
}

// create asset
func (t *pharma_track) createAsset(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
    
	//To check number of arguments are 5
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments, Required 5 arguments")
	}
	
	//To check each argument is not null
	for i := 0; i < len(args); i++ {
		if len(args[i]) <= 0 {
			return shim.Error(string(i+1) + "st argument must be a non-empty string")
		}
	}
	
	// check if the owner exists
	ownerBytes, _ := APIstub.GetState(args[4])

	if ownerBytes == nil {
		return shim.Error("Cannot Find Owner Asset with id : " + args[4])
	}
	
	ownerAsset := Owner_Asset{}
	
	json.Unmarshal(ownerBytes, &ownerAsset)
	
	pharmaCounter := getCounter(APIstub, "pharmaCounter")
	pharmaCounter++
	
	var pharmaAsset = Pharma_Asset{Asset_ID: "pharma" + strconv.Itoa(pharmaCounter), Asset_Name: args[1],	Asset_Batch_no: args[2], Asset_Expiry: args[3], Asset_Owner: ownerAsset, Asset_Status: "Manufactured"}
	
	//convert to bytes
	pharmaAssetAsBytes, errMarshal := json.Marshal(pharmaAsset)
	
	if errMarshal != nil {
	    return shim.Error(fmt.Sprintf("Marshal Error in Owner: %s", errMarshal))
	}
	
	errPut := APIstub.PutState(pharmaAsset.Asset_ID, pharmaAssetAsBytes)

	if errPut != nil {
		return shim.Error(fmt.Sprintf("Failed to create Pharma Asset: %s", pharmaAsset.Asset_ID))
	}

	//Increment the pharma Counter
	incrementCounter(APIstub, "pharmaCounter")

	fmt.Println("Success in creating Pharma Asset %v", pharmaAsset)
        logger.Infof("Success in creating Pharma Asset %+v", pharmaAsset)

	return shim.Success(nil)
}

// create owner
func (t *pharma_track) createOwner(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	//To check number of arguments are 3
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments, Required 3 arguments")
	}
	
	//To check each argument is not null
	for i := 0; i < len(args); i++ {
		if len(args[i]) <= 0 {
			return shim.Error(string(i+1) + "st argument must be a non-empty string")
		}
	}
	
	ownerCounter := getCounter(APIstub, "OwnerCount")
	ownerCounter++
	
	var ownerAsset = Owner_Asset{Owner_Name: args[1], Owner_Id: "Owner" + strconv.Itoa(ownerCounter), Address: args[2]}
	
	//convert to bytes
	ownerAssetAsBytes, errMarshal := json.Marshal(ownerAsset)
	
	if errMarshal != nil {
	    return shim.Error(fmt.Sprintf("Marshal Error in Owner: %s", errMarshal))
	}
	
	errPut := APIstub.PutState(ownerAsset.Owner_Id, ownerAssetAsBytes)

	if errPut != nil {
		return shim.Error(fmt.Sprintf("Failed to create Owner Asset: %s", ownerAsset.Owner_Id))
	}

	//Increment the owner Counter
	incrementCounter(APIstub, "OwnerCount")

	fmt.Println("Success in creating Owner Asset %v", ownerAsset)
        logger.Infof("Success in creating Owner Asset %+v", ownerAsset)

	return shim.Success(nil)
}

// transfer ownership 
func (t *pharma_track) transferOwnership(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
    
	//To check number of arguments are 3
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments, Required 3 arguments")
	}
	
	//To check each argument is not null
	for i := 0; i < len(args); i++ {
		if len(args[i]) <= 0 {
			return shim.Error(string(i+1) + "st argument must be a non-empty string")
		}
	}
	
	// get the assests on ledger
	ownerAsset := Owner_Asset{}
	
	ownerBytes, _ := APIstub.GetState(args[2])

	if ownerBytes == nil {
		return shim.Error("Cannot Find Owner Asset with id : " + args[2])
	}
	
	json.Unmarshal(ownerBytes, &ownerAsset)
	
	pharmaAsset := Pharma_Asset{}
	
	pharmaBytes, _ := APIstub.GetState(args[1])

	if pharmaBytes == nil {
		return shim.Error("Cannot Find Pharma Asset with id : " + args[1])
	}
	
	json.Unmarshal(pharmaBytes, &pharmaAsset)
	
	pharmaAsset.changeOwner(ownerAsset)
	
	pharmaAssetAsBytes, errMarshal := json.Marshal(pharmaAsset)
	
	if errMarshal != nil {
		return shim.Error(fmt.Sprintf("Marshal Error: %s", errMarshal))
	}
	
	errPutOrder := APIstub.PutState(pharmaAsset.Asset_ID, pharmaAssetAsBytes)

	if errPutOrder != nil {
		return shim.Error(fmt.Sprintf("Failed to transfer Ownership: %s", pharmaAsset.Asset_ID))
	}

	return shim.Success(nil)
}

// update status
func (t *pharma_track) updatePharmaOrderStatus(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	// check number of arguments are 3
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments, Required 3 arguments")
	}
	
	// check each argument is not null
	for i := 0; i < len(args); i++ {
		if len(args[i]) <= 0 {
			return shim.Error(string(i+1) + "st argument must be a non-empty string")
		}
	}
	
	// validate the status to be update is correct
	isValid := isValidStatus(args[2])
	if !isValid {
	    fmt.Println("Status value provided is not valid")
		return shim.Error("Status value provided is not valid")
	}
	
	// get the pharma asset
	pharmaBytes, _ := APIstub.GetState(args[1])

	if pharmaBytes == nil {
		return shim.Error("Cannot Find Pharma Asset with id : " + args[1])
	}
	
	pharmaAsset := Pharma_Asset{}
	json.Unmarshal(pharmaBytes, &pharmaAsset)
	
	// change the status
	pharmaAsset.Asset_Status = args[2]
	
	//update the ledger
	pharmaAssetAsBytes, errMarshal := json.Marshal(pharmaAsset)
	
	if errMarshal != nil {
		return shim.Error(fmt.Sprintf("Marshal Error: %s", errMarshal))
	}
	
	errPutOrder := APIstub.PutState(pharmaAsset.Asset_ID, pharmaAssetAsBytes)

	if errPutOrder != nil {
		return shim.Error(fmt.Sprintf("Failed to change status: %s", pharmaAsset.Asset_ID))
	}

	return shim.Success(nil)
	
}


// ============================================================================================================================
// Invoke - Our entry point for Invocations
// ============================================================================================================================
func (t *pharma_track) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("function is ==> :" + function)
        logger.Infof("function is ==> :" + function)
	action := args[0]
	fmt.Println(" action is ==> :" + action)
        logger.Infof(" action is ==> :" + action)
	fmt.Println(args)
        logger.Infof("args : %q", args)
	
	if action == "queryAsset" {
		return t.queryAsset(stub, args)
	} else if action == "createAsset" {
		return t.createAsset(stub, args)
	} else if action == "createOwner" {
		return t.createOwner(stub, args)
	} else if action == "transferOwnership" {
	    return t.transferOwnership(stub, args)
	} else if action == "updatePharmaOrderStatus" {
		return t.updatePharmaOrderStatus(stub, args)
	} 
	
	fmt.Println("invoke did not find func: " + action) //error

	return shim.Error("Received unknown function")
}

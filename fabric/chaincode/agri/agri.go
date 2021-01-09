package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	//"strings"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

//SmartContract is the data structure which represents this contract and on which  various contract lifecycle functions are attached
type SmartContract struct {
}

type Expert struct {
	ObjectType string `json:"Type"`
	Expert_ID  string `json:"expert_ID"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type Farmer struct {
	ObjectType string `json:"Type"`
	Farmer_ID  string `json:"farmer_ID"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type Vendor struct {
	ObjectType string `json:"Type"`
	Vendor_ID  string `json:"vendor_ID"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type Product struct {
	ObjectType string `json:"Type"`
	Product_ID string `json:"product_ID"`
	Name       string `json:"name"`
	Amount     string `json:"amount"`
	Price      string `json:"price"`
	Farmer_ID  string `json:"farmer_ID"`
	Vendor_ID  string `json:"vendor_ID"`
	Expert_ID  string `json:"expert_ID"`
	Status     string `json:"status"`
}
type VerificationRequest struct {
	ObjectType            string `json:"Type"`
	VerificationRequestID string `json:"verificationRequestID"`
	DisorderType          string `json:"disorderType"`
	Disorder_degree       string `json:"disorder_degree"`
	Research_institute    string `json:"research_institute"`
	Farmer_ID             string `json:"farmer_ID"`
	Status                string `json:"status"`
}

type ExpertResponse struct {
	ObjectType string   `json:"Type"`
	ResponseID string   `json:"responseID"`
	Expert_ID  string   `json:"expert_ID"`
	Farmer_ID  string   `json:"farmer_ID"`
	Response   string   `json:"response"`
	Products   []string `json:"products"`
}

func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {

	fmt.Println("Init Firing!")
	return shim.Success(nil)
}

func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("Chaincode Invoke Is Running " + function)

	if function == "addExpert" {
		return t.addExpert(stub, args)
	}
	if function == "queryExpert" {
		return t.queryExpert(stub, args)
	}
	if function == "addFarmer" {
		return t.addFarmer(stub, args)
	}
	if function == "queryFarmer" {
		return t.queryFarmer(stub, args)
	}
	if function == "queryFarmerbyID" {
		return t.queryFarmerbyID(stub, args)
	}
	if function == "addVendor" {
		return t.addVendor(stub, args)
	}
	if function == "queryVendor" {
		return t.queryVendor(stub, args)
	}
	if function == "queryVendorbyID" {
		return t.queryVendorbyID(stub, args)
	}
	if function == "addProduct" {
		return t.addProduct(stub, args)
	}
	if function == "queryProduct" {
		return t.queryProduct(stub, args)
	}
	if function == "queryProducts" {
		return t.queryProducts(stub, args)
	}
	if function == "queryProductbyFarmer" {
		return t.queryProductbyFarmer(stub, args)
	}
	if function == "addVerificationRequest" {
		return t.addVerificationRequest(stub, args)
	}
	if function == "queryVerificationRequest" {
		return t.queryVerificationRequest(stub, args)
	}
	if function == "queryVerificationRequestbyFarmer" {
		return t.queryVerificationRequestbyFarmer(stub, args)
	}
	if function == "queryVerificationRequestbyID" {
		return t.queryVerificationRequestbyID(stub, args)
	}
	if function == "queryVerificationRequests" {
		return t.queryVerificationRequests(stub, args)
	}
	if function == "addExpertResponse" {
		return t.addExpertResponse(stub, args)
	}
	if function == "queryExpertResponsebyFarmer" {
		return t.queryExpertResponsebyFarmer(stub, args)
	}
	if function == "updateVerificationRequest" {
		return t.updateVerificationRequest(stub, args)
	}
	if function == "updateProduct" {
		return t.updateProduct(stub, args)
	}
	if function == "giveRecomendation" {
		return t.giveRecomendation(stub, args)
	}
	if function == "queryAllProducts" {
		return t.queryAllProducts(stub, args)
	}
	if function == "queryAllVerificationRequests" {
		return t.queryAllVerificationRequests(stub, args)
	}

	fmt.Println("Invoke did not find specified function " + function)
	return shim.Error("Invoke did not find specified function " + function)
}

func (t *SmartContract) addExpert(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect Number of Aruments. Expecting 3")
	}

	fmt.Println("Adding new Expert")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}

	expert_ID := args[0]
	username := args[1]
	password := args[2]

	// ======Check if expert Already exists

	expertAsBytes, err := stub.GetState(expert_ID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if expertAsBytes != nil {
		return shim.Error("The Inserted expert ID already Exists: " + expert_ID)
	}

	// ===== Create expert Object and Marshal to JSON

	objectType := "expert"
	expert := &Expert{objectType, expert_ID, username, password}
	expertJSONasBytes, err := json.Marshal(expert)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save expert to State

	err = stub.PutState(expert_ID, expertJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved expert")
	return shim.Success(nil)
}

func (t *SmartContract) queryExpert(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	username := args[0]
	password := args[1]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"expert\",\"username\":\"%s\",\"password\":\"%s\"}}", username, password)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addFarmer(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect Number of Aruments. Expecting 3")
	}

	fmt.Println("Adding new Farmer")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}

	farmer_ID := args[0]
	username := args[1]
	password := args[2]

	// ======Check if farmer Already exists

	farmerAsBytes, err := stub.GetState(farmer_ID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if farmerAsBytes != nil {
		return shim.Error("The Inserted farmer ID already Exists: " + farmer_ID)
	}

	// ===== Create farmer Object and Marshal to JSON

	objectType := "farmer"
	farmer := &Farmer{objectType, farmer_ID, username, password}
	farmerJSONasBytes, err := json.Marshal(farmer)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save farmer to State

	err = stub.PutState(farmer_ID, farmerJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved farmer")
	return shim.Success(nil)
}

func (t *SmartContract) addVendor(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect Number of Aruments. Expecting 3")
	}

	fmt.Println("Adding new Vendor")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}

	vendor_ID := args[0]
	username := args[1]
	password := args[2]

	// ======Check if vendor Already exists

	vendorAsBytes, err := stub.GetState(vendor_ID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if vendorAsBytes != nil {
		return shim.Error("The Inserted vendor ID already Exists: " + vendor_ID)
	}

	// ===== Create vendor Object and Marshal to JSON

	objectType := "vendor"
	vendor := &Vendor{objectType, vendor_ID, username, password}
	vendorJSONasBytes, err := json.Marshal(vendor)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save vendor to State

	err = stub.PutState(vendor_ID, vendorJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved vendor")
	return shim.Success(nil)
}

func (t *SmartContract) queryVendor(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	username := args[0]
	password := args[1]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"vendor\",\"username\":\"%s\",\"password\":\"%s\"}}", username, password)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryVendorbyID(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	vendor_ID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"vendor\",\"vendor_ID\":\"%s\"}}", vendor_ID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addProduct(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 8 {
		return shim.Error("Incorrect Number of Aruments. Expecting 6")
	}

	fmt.Println("Adding new Product")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th Argument Must be a Non-Empty String")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th Argument Must be a Non-Empty String")
	}
	if len(args[5]) <= 0 {
		return shim.Error("6th Argument Must be a Non-Empty String")
	}
	if len(args[6]) <= 0 {
		return shim.Error("7th Argument Must be a Non-Empty String")
	}
	if len(args[7]) <= 0 {
		return shim.Error("8th Argument Must be a Non-Empty String")
	}

	product_ID := args[0]
	name := args[1]
	amount := args[2]
	price := args[3]
	farmer_ID := args[4]
	vendor_ID := args[5]
	expert_ID := args[6]
	status := args[7]

	// ======Check if product Already exists

	productAsBytes, err := stub.GetState(product_ID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if productAsBytes != nil {
		return shim.Error("The Inserted product ID already Exists: " + product_ID)
	}

	// ===== Create product Object and Marshal to JSON

	objectType := "product"
	product := &Product{objectType, product_ID, name, amount, price, farmer_ID, vendor_ID, expert_ID, status}
	productJSONasBytes, err := json.Marshal(product)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save product to State

	err = stub.PutState(product_ID, productJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved product")
	return shim.Success(nil)
}

func (t *SmartContract) queryProduct(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	status := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"product\",\"status\":\"%s\"}}", status)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryProductbyFarmer(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	farmer_ID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"product\",\"farmer_ID\":\"%s\"}}", farmer_ID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryProducts(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	status := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"product\",\"status\":{\"$ne\":\"%s\"}}}", status)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addVerificationRequest(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 6 {
		return shim.Error("Incorrect Number of Aruments. Expecting 10")
	}

	fmt.Println("Adding new Nuttrient defficiency")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4rd Argument Must be a Non-Empty String")
	}
	if len(args[4]) <= 0 {
		return shim.Error("4rd Argument Must be a Non-Empty String")
	}
	if len(args[5]) <= 0 {
		return shim.Error("4rd Argument Must be a Non-Empty String")
	}

	verificationRequestID := args[0]
	disorderType := args[1]
	disorder_degree := args[2]
	research_institute := args[3]
	farmer_ID := args[4]
	status := args[5]
	// ======Check if defficeincy Already exists

	verificationRequestAsBytes, err := stub.GetState(verificationRequestID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if verificationRequestAsBytes != nil {
		return shim.Error("The Inserted verificationRequest ID already Exists: " + verificationRequestID)
	}

	// ===== Create defficiency Object and Marshal to JSON

	objectType := "verificationRequest"
	verificationRequest := &VerificationRequest{objectType, verificationRequestID, disorderType, disorder_degree, research_institute, farmer_ID, status}
	VerificationRequestJSONasBytes, err := json.Marshal(verificationRequest)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save deficiency to State

	err = stub.PutState(verificationRequestID, VerificationRequestJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved verificationRequest")
	return shim.Success(nil)
}

func (t *SmartContract) queryVerificationRequest(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	status := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"verificationRequest\",\"status\":\"%s\"}}", status)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryVerificationRequestbyFarmer(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	farmer_ID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"verificationRequest\",\"farmer_ID\":\"%s\"}}", farmer_ID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryVerificationRequestbyID(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	verificationRequestID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"verificationRequest\",\"verificationRequestID\":\"%s\"}}", verificationRequestID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryVerificationRequests(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	status := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"verificationRequest\",\"status\":{\"$ne\":\"%s\"}}}", status)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryAllProducts(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"product\"}}")

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryAllVerificationRequests(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"verificationRequest\"}}")

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) addExpertResponse(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	var err error

	if len(args) != 5 {
		return shim.Error("Incorrect Number of Aruments. Expecting 3")
	}

	fmt.Println("Adding new response")

	// ==== Input sanitation ====
	if len(args[0]) <= 0 {
		return shim.Error("1st Argument Must be a Non-Empty String")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd Argument Must be a Non-Empty String")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd Argument Must be a Non-Empty String")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4rd Argument Must be a Non-Empty String")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5rd Argument Must be a Non-Empty String")
	}

	responseID := args[0]
	expert_ID := args[1]
	farmer_ID := args[2]
	response := args[2]
	products := args[3]
	// ======Check if admin Already exists

	responseAsBytes, err := stub.GetState(responseID)
	if err != nil {
		return shim.Error("Transaction Failed with Error: " + err.Error())
	} else if responseAsBytes != nil {
		return shim.Error("The Inserted response ID already Exists: " + responseID)
	}

	// ===== Create admin Object and Marshal to JSON

	objectType := "response"
	expertResponse := &ExpertResponse{objectType, responseID, expert_ID, farmer_ID, response, append(ExpertResponse{}.Products, products)}
	responseJSONasBytes, err := json.Marshal(expertResponse)

	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Save admin to State

	err = stub.PutState(responseID, responseJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ======= Return Success

	fmt.Println("Successfully Saved admin")
	return shim.Success(nil)
}

func (t *SmartContract) queryExpertResponsebyFarmer(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	farmer_ID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"response\",\"farmer_ID\":\"%s\"}}", farmer_ID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) updateVerificationRequest(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	verificationRequestID := args[0]
	newStatus := args[1]
	fmt.Println("- start  ", verificationRequestID, newStatus)

	verificationRequestAsBytes, err := stub.GetState(verificationRequestID)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if verificationRequestAsBytes == nil {
		return shim.Error("Session does not exist")
	}

	verificationRequestToUpdate := VerificationRequest{}
	err = json.Unmarshal(verificationRequestAsBytes, &verificationRequestToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	verificationRequestToUpdate.Status = newStatus //change the status

	verificationRequestJSONasBytes, _ := json.Marshal(verificationRequestToUpdate)
	err = stub.PutState(verificationRequestID, verificationRequestJSONasBytes) //rewrite the marble
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) updateProduct(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 5 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	product_ID := args[0]
	amount := args[1]
	price := args[2]
	vendor_ID := args[3]
	status := args[4]
	fmt.Println("- start  ", product_ID, amount, price, vendor_ID, status)

	productAsBytes, err := stub.GetState(product_ID)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if productAsBytes == nil {
		return shim.Error("product does not exist")
	}

	productToUpdate := Product{}
	err = json.Unmarshal(productAsBytes, &productToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	productToUpdate.Amount = amount
	productToUpdate.Price = price
	productToUpdate.Vendor_ID = vendor_ID
	productToUpdate.Status = status //change the status

	productJSONasBytes, _ := json.Marshal(productToUpdate)
	err = stub.PutState(product_ID, productJSONasBytes) //rewrite
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) giveRecomendation(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	product_ID := args[0]
	expert_ID := args[1]
	status := args[2]
	fmt.Println("- start  ", product_ID, expert_ID, status)

	productAsBytes, err := stub.GetState(product_ID)
	if err != nil {
		return shim.Error("Failed to get status:" + err.Error())
	} else if productAsBytes == nil {
		return shim.Error("product does not exist")
	}

	productToUpdate := Product{}
	err = json.Unmarshal(productAsBytes, &productToUpdate) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}

	productToUpdate.Expert_ID = expert_ID
	productToUpdate.Status = status //change the status

	productJSONasBytes, _ := json.Marshal(productToUpdate)
	err = stub.PutState(product_ID, productJSONasBytes) //rewrite
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end  (success)")
	return shim.Success(nil)
}

func (t *SmartContract) queryFarmer(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	username := args[0]
	password := args[1]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"farmer\",\"username\":\"%s\",\"password\":\"%s\"}}", username, password)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func (t *SmartContract) queryFarmerbyID(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	farmer_ID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"Type\":\"farmer\",\"farmer_ID\":\"%s\"}}", farmer_ID)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

//Main Function starts up the Chaincode
func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Smart Contract could not be run. Error Occured: %s", err)
	} else {
		fmt.Println("Smart Contract successfully Initiated")
	}
}

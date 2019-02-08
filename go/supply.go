/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

 package main

 /* Imports
  * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
  * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
  */
 import (
	 "bytes"
	 "encoding/json"
	 "fmt"
	 "strconv"
 
	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 sc "github.com/hyperledger/fabric/protos/peer"
 )
 
 // Define the Smart Contract structure
 type SmartContract struct {
 }
 
 // Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
 type Invoice struct {
	 BilledTo  string `json:"billedTo"`
	 InvoiceDate string `json:"invoiceDate"`
	 InvoiceAmount  string `json:"invoiceAmount"`
	 ItemDescription  string `json:"itemDescription"`
	 GoodReceived  string `json:"goodReceived"`
	 IsPaid  string `json:"isPaid"`
	 PaidAmount  string `json:"paidAmount"`
	 Repaid  string `json:"repaid"`
	 RepaymentAmount  string `json:"repaymentAmount"`
 }
 
 /*
  * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
  * Best practice is to have any Ledger initialization in separate function -- see initLedger()
  */
 func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	 return shim.Success(nil)
 }
 
 /*
  * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
  * The calling application program has also specified the particular smart contract function to be called, with arguments
  */
 func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
 
	 // Retrieve the requested Smart Contract function and arguments
	 function, args := APIstub.GetFunctionAndParameters()
	 // Route to the appropriate handler function to interact with the ledger appropriately
	 
	if function == "initLedger" {
		 return s.initLedger(APIstub)
	}else if function == "raiseInvoice" {
		return s.raiseInvoice(APIstub, args)
	}else if function == "queryAllInvoices" {
		return s.queryAllInvoices(APIstub)
	}else if function == "goodReceived" {
		return s.goodReceived(APIstub, args) 
	}else if function == "repaymentAmount" {
		return s.repaymentAmount(APIstub, args) 
	}
 
	 return shim.Error("Invalid Smart Contract function name.")
 }

// FABCAR===================
//  func (s *SmartContract) queryCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
// 	 if len(args) != 1 {
// 		 return shim.Error("Incorrect number of arguments. Expecting 1")
// 	 }
 
// 	 carAsBytes, _ := APIstub.GetState(args[0])
// 	 return shim.Success(carAsBytes)
//  }
 
 func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	 invoices := []Invoice{
		 Invoice{BilledTo: "ASUS", InvoiceDate: "02/07/19",InvoiceAmount: "350000", ItemDescription: "Processor", GoodReceived: "False", IsPaid: "False", PaidAmount: "100000", Repaid: "True", RepaymentAmount: "250000"},
		 Invoice{BilledTo: "OEM", InvoiceDate: "02/08/19",InvoiceAmount: "200000", ItemDescription: "SSD", GoodReceived: "False", IsPaid: "False", PaidAmount: "100000", Repaid: "True", RepaymentAmount: "100000"},
		 Invoice{BilledTo: "OEM", InvoiceDate: "02/09/19",InvoiceAmount: "300000", ItemDescription: "RAM", GoodReceived: "False", IsPaid: "False", PaidAmount: "100000", Repaid: "True", RepaymentAmount: "200000"},
		 Invoice{BilledTo: "OEM", InvoiceDate: "02/10/19",InvoiceAmount: "100000", ItemDescription: "HDD", GoodReceived: "False", IsPaid: "False", PaidAmount: "10000" , Repaid: "True", RepaymentAmount: "90000"},
	 }
 
	 i := 0
	 for i < len(invoices) {
		 fmt.Println("i is ", i)
		 invoiceAsBytes, _ := json.Marshal(invoices[i])
		 APIstub.PutState("INVOICE"+strconv.Itoa(i), invoiceAsBytes)
		 fmt.Println("Added", invoices[i])
		 i = i + 1
	 }
 
	 return shim.Success(nil)
 }
 func (s *SmartContract) raiseInvoice(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	
	if len(args) != 10 {
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}

	var invoice = Invoice{BilledTo: args[1], InvoiceDate: args[2] ,InvoiceAmount: args[3], ItemDescription: args[4], GoodReceived: args[5], IsPaid: args[6], PaidAmount: args[7], Repaid: args[8], RepaymentAmount: args[9]}

	invoiceAsBytes, _:= json.Marshal(invoice)
	APIstub.PutState(args[0], invoiceAsBytes)

	return shim.Success(nil)
}
 
//  func (s *SmartContract) createCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
// 	 if len(args) != 5 {
// 		 return shim.Error("Incorrect number of arguments. Expecting 5")
// 	 }
 
// 	 var car = Car{Make: args[1], Model: args[2], Colour: args[3], Owner: args[4]}
 
// 	 carAsBytes, _ := json.Marshal(car)
// 	 APIstub.PutState(args[0], carAsBytes)
 
// 	 return shim.Success(nil)
//  }
 
 func (s *SmartContract) queryAllInvoices(APIstub shim.ChaincodeStubInterface) sc.Response {
 
	 startKey := "INVOICE0"
	 endKey := "INVOICE999"
 
	 resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	 if err != nil {
		 return shim.Error(err.Error())
	 }
	 defer resultsIterator.Close()
 
	 // buffer is a JSON array containing QueryResults
	 var buffer bytes.Buffer
	 buffer.WriteString("[")
 
	 bArrayMemberAlreadyWritten := false
	 for resultsIterator.HasNext() {
		 queryResponse, err := resultsIterator.Next()
		 if err != nil {
			 return shim.Error(err.Error())
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
 
	 fmt.Printf("- queryAllInvoices:\n%s\n", buffer.String())
 
	 return shim.Success(buffer.Bytes())
 }

 func (s *SmartContract) goodReceived(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	invoiceAsBytes, _ := APIstub.GetState(args[0])
	invoice := Invoice{}

	json.Unmarshal(invoiceAsBytes, &invoice)
	invoice.GoodReceived = args[1]

	invoiceAsBytes, _ = json.Marshal(invoice)
	APIstub.PutState(args[0], invoiceAsBytes)
	
	return shim.Success(nil)
}

func (s *SmartContract) repaymentAmount(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	invoiceAsBytes, _ := APIstub.GetState(args[0])
	invoice := Invoice{}

	json.Unmarshal(invoiceAsBytes, &invoice)
	// invoice.RepaymentAmount = args[1]

	amount, err := strconv.ParseFloat(args[1], 32)
	if err != nil {
		// do something sensible
	}
	paid, err := strconv.ParseFloat(invoice.PaidAmount, 32)
	if err != nil {
		// do something sensible
	}

	if(amount < paid) {
		return shim.Error("Repayment should be greater than paid amount!")
	}
	
	return shim.Success(nil)
}
 
//  func (s *SmartContract) changeCarOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
 
// 	 if len(args) != 2 {
// 		 return shim.Error("Incorrect number of arguments. Expecting 2")
// 	 }
 
// 	 carAsBytes, _ := APIstub.GetState(args[0])
// 	 car := Car{}
 
// 	 json.Unmarshal(carAsBytes, &car)
// 	 car.Owner = args[1]
 
// 	 carAsBytes, _ = json.Marshal(car)
// 	 APIstub.PutState(args[0], carAsBytes)
 
// 	 return shim.Success(nil)
//  }
 
 // The main function is only relevant in unit test mode. Only included here for completeness.
 func main() {
 
	 // Create a new Smart Contract
	 err := shim.Start(new(SmartContract))
	 if err != nil {
		 fmt.Printf("Error creating new Smart Contract: %s", err)
	 }
 }
 
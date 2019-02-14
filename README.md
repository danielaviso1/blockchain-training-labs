Daniel D. Aviso
Blockchain Cadet
Laptop Specifications
CPU: Intel® Core™ i5-7300HQ @ 2.50 GHz
GPU: NVIDIA GeForce GTX 1050 (4 GB DDR5)
RAM: 8 GB DDR4 @ 2400MHz
OS (Operating System): Windows 10 Home © 2018 Microsoft Corporation

Oracle VM Virtual Box Specifications
Processor(s): 4 CPUs
Base Memory: 4096 MB
Execution Cap: 100%
Video Memory: 128 MB
Installed OS (Operating System): Ubuntu 16.04 LTS (Xenial Xerus) 64 Bit

Environment Setup
Step 1: Type in terminal: “sudo apt install snapd snapd-xdg-open” Link
Description: This command line will install the snap package manager to easily download and install software and set your environment for that specific software.
Step 2: Type in terminal: “sudo snap install –classic go”  Link
Description: This command line will install the latest version of Go Language.
Step 3:   Type in terminal: “snap install postman”  Link
Description: This will install the Postman app that will be used later.
Step 4: Type in terminal: “curl -O https://hyperledger.github.io/composer/latest/prereqs-ubuntu.sh” Link
Description: This will download the prereqs-ubuntu.sh file. 
Step 5: Type in terminal: “chmod u+x prereqs-ubuntu.sh”
Description: This command will read and execute the file that you downloaded. (This file is made to downl-oad and install all the prerequisites of HyperLedger).
Step 6: Type in terminal: “curl -sSL http://bit.ly/2ysbOFE | bash -s -- 1.4.0” Link
Description: The command above downloads and executes a bash script that will download and extract all the platform-specific binaries you will need to set up your network.

HyperLedger Fabric First Network Setup
Step 6: Type in terminal: “cd fabric-samples/first-network”
Description: Open your terminal and type this command to change your directory to the first-network folder.
Step 7: Type in terminal: “./byfn.sh generate” (If you do not get an error here ignore Step X and Step XX and Step Y)
Step 8: Type in terminal: “./byfn.sh up” (If you do not get an error here ignore Step X and Step XX and Step Y)
Step Y: Type in terminal: “./byfn.sh down”
Step X: Type in terminal: “docker ps -qa | xargs docker kill”
Step XX: Type in terminal: “docker ps -qa | xargs docker rm”
Step 9:  Type in terminal: “docker exec -it peer0.org2.example.com bash”
Step 10: Type in terminal: “cd /var/hyperledger/production/ledgersData/chains/chains/mychannel”
Step 11: Type in terminal: “cat blockfile_000000”
Step 12: Type in terminal: “exit”
Description: These commands are made to test your first network.

Access for State Level Certificate
Step 13:  Type in terminal: “docker exec -it peer0.org2.example.com”
Step 14: Type in terminal: “cd /var/hyperledger/production1/ledgersData/stateLeveldb”
Step 15: Type in terminal: “cat 000001.log”
Step 16: Type in terminal: “exit”
Description: These commands will access the state level certificate.

Mandatory Assignment Setup
Step 17: Go to fabric-samples folder.
Step 18: Type in terminal: “git clone https://github.com/khrandm/blockchain-training-labs.git” Link
Step 19: Open the blockchain-training-labs folder copy the chaincode and suppy folder then paste it inside the fabric-samples folder.
Step 20: Go to the supply folder that you pasted inside the fabric-samples folder.
Step 21: Type in terminal: “go get github.com/golang/protobuf/proto”
Step 22: Type in terminal: “go get github.com/hyperledger/fabric/common/attrmgr”
Step 23: Type in terminal: “go get github.com/pkg/errors”
Step 24: Type in terminal: “go get github.com/hyperledger/fabric/core/chaincode/lib/cid”
Step 25: Open file manager go to Home/go/src/github.com
Step 26: Copy all the folders you see.
Step 27: Paste it inside the fabric-sample/chaincode
Description: These commands will Download the required library for our chaincode.
Step 28: Open terminal from that directory.
Step 29: Type in terminal: “./startFabric.sh”
Step 30: Type in terminal: “npm install” (If you got stuck in “node-pre-gyp” do -> CTRL + C then Continue to next Step)

Mandatory Assignment Use
Step 31: Type in terminal: “node enrollAdmin.js”
Step 32: Type in terminal: “node registerUser.js”
Step 33: Type in terminal: “node app.js”
Description: If everything is right, you should see in your terminal a command saying “Example app listening on port 3000!”.
Step 34: Open the Postman app you install earlier in Step 3.
Step 35: Make sure Request Method is using GET then type in Request URL localhost:3000/
Step 36: Click Send
Description: You should see at the bottom of your Postman app all the sample JSON Data.
Step 37: Change Request Method to POST then type in Request URL localhost:3000/invoice
Step 38: Go to Body Tab and choose x-www-form-urlencoded
Step 39: Click Bulk Edit from Middle Right corner in Postman
Step 40: Copy and Paste the following Data to the Input Box 
KEY 			VALUE
invoicenumber:		INVOICE6
billedto:		OEM
invoicedate:		02/08/19
invoiceamount:		10000
itemdescription:	KEYBOARD
goodreceived:		False
ispaid:			False
paidamount:		0
repaid:			False
repaymentamount:	0

Step 41: Click Send.
Description: If you made this right you should see a success written at the bottom of your Postman app after clicking the Send button. This means that you already posted or added a data. You could see it from your terminal.
Step 42: Change Request Method to PUT then type in Request URL localhost:3000/invoice
Step 43: Going back to Key-Value Edit uncheck all data Except invoicenumber and goodreceived
Step 44: Change the Value of goodreceived from False to True
Step 45: Click Send.
Description: If you made this right you should see a success written at the bottom of your Postman app
After clicking the Send button. This means that you already updated the goodreceived the from INVOICE6 data from False to True. You could see it from your terminal.

Optional Assignment Use
Description: If you would like to use the Optional you should close the running terminal and go back to Step 29 and 30 and continue with the following steps.

Step 31: Type in terminal: “./startFabric.sh”
Step 32: Type in terminal: “node enrollAdmin.js”
Step 33: Type in terminal: “node registerSupplier.js”
Step 34: Type in terminal: “node registerOEM.js”
Step 35: Type in terminal: “node registerBank.js”
Step 36: Type in terminal: “node app.js”
Description: If everything is right, you should see in your terminal a command saying, “Example app listening on port 3000!”.
Step 37: Open Postman App
Description: To raise an invoice.
Step 38: Change the default Request Method from Get to POST 
Step 39: Type in REQUEST URL: “localhost:3000/invoice”
Step 40: Click the Headers Tab.
Step 41: Input “user” key under Content-Type Column
Step 42: Input “supplier” under Value Column
Step 43: Click the Body Tab
Step 44: Click x-www-form-urlencoded
Step 45: Click Bulk Edit and Copy Paste these Data
KEY		       VALUE
invoicenumber          INVOICE001
billedto                	OEM
invoicedate             	02/08/19
invoiceamount          10000
itemdescription        KEYBOARD
goodreceived            False
ispaid                  	False
paidamount              	0
repaid                  	False
repaymentamount   0

Step 46: Click Send 
Description: If you made it right, you should see a Success in JSON at the bottom of the Postman App
(If you want to see the POST you made you can change the Request Method with GET and type localhost:3000 in REQUEST URL and Click Send)

Description: Bank pays the Supplier
Step 47: Change the Request Method to PUT
Step 48: Type in the REQUEST URL: “localhost:3000/invoice”
Step 49: Click Header Tab
Step 50: Input “user” key under Content-Type Column
Step 51: Input “bank” under Value Column
Step 52: Click the Body Tab
Step 53: Click x-www-form-urlencoded
Step 54: Click Bulk Edit and Copy Paste these Data

KEY		       VALUE
invoicenumber          INVOICE001
paidamount              9000

*NOTE: There are conditions that should be met, the paid amount should be less than invoice amount.
Step 55: Click Send
Description: If you made it right, you should see a Success in JSON at the bottom of the Postman App
(If you change the Request Method to GET and input localhost:3000 and Click Send you will see that the isPaid value became True).




Description: OEM pays the Bank
Step 56: Change the Request Method to PUT
Step 57: Type in the REQUEST URL: “localhost:3000/invoice”
Step 58: Click Header Tab
Step 59: Input “user” key under Content-Type Column
Step 60: Input “value” under Value Column
Step 61: Click the Body Tab
Step 62: Click x-www-form-urlencoded
Step 63: Click Bulk Edit and Copy Paste these Data

KEY		       VALUE
invoicenumber          INVOICE001
repaymentamount   11000
*NOTE: There are conditions that should be met, the repayment amount should be more than paidamount.

Step 64: Click Send

Description: If you made it right, you should see a Success in JSON at the bottom of the Postman App
(If you want to see the changes you made you can change the Request Method with GET and type localhost:3000 in REQUEST URL and Click Send to check if the data is updated).

Description: Invoice Audit Trail
Step 65: Change the Request Method to GET
Step 66: Type in the REQUEST URL: “localhost:3000”
Step 67: Click Header Tab
Step 68: Input “user key under Content-Type Column
Step 69: Input “value” under Value Column
Step 70: Click the Body Tab
Step 71: Click x-www-form-urlencoded
Step 72:  Insert these Data
KEY		       VALUE
invoicenumber 	INVOICE001 

Step 73: Click Send

Description: If you made it right, you should see the respond from the server change it from Html to Json to see a json format of the response.
REFERENCES
Source: https://github.com/khrandm/blockchain-training-labs
Authors: Jaymar Dingcong,
	  Michael Forte,
	  Rhenzo Pacho
Access: February 12, 2019
Year: 2019

Source: https://en.wikipedia.org/wiki/Snappy_(package_manager)
Author: Wikipedia
Access: February 12, 2019
Year: 2014

Source: https://help.ubuntu.com/community/FilePermissions
Author: craig-leat
Access: February 12, 2019
Year: 2013

Source: http://ubuntuhandbook.org/index.php/2018/09/install-postman-app-easily-via-snap-in-ubuntu-18-04/
Author: Ji m (ubuntuhandbook1@gmail.com)
Access: February 13, 2019
Year: 2018

Source: https://hyperledger-fabric.readthedocs.io/en/latest/install.html
Author: HyperLedger
Access: February 12, 2019
Year: 2019

Source: https://medium.com/@patdhlk/how-to-install-go-1-9-1-on-ubuntu-16-04-ee64c073cd79
Author: Patrick Dahlke
Access: February 12, 2019
Year: 2017

Source: https://medium.com/kago/tutorial-to-install-hyperledger-composer-on-windows-88d973094b5c
Author: Eddie Kago
Access: January 14, 2019
Year: 2018

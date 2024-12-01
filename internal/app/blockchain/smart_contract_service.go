package blockchain

import (
	"awesomeProject/internal/app/domain/dao"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	mspID        = "Org1MSP"
	cryptoPath   = "internal/app/blockchain/organizations/peerOrganizations/org1.example.com"
	certPath     = cryptoPath + "/users/Admin@org1.example.com/msp/signcerts"
	keyPath      = cryptoPath + "/users/Admin@org1.example.com/msp/keystore"
	tlsCertPath  = cryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt"
	peerEndpoint = "dns:///localhost:7051"
	gatewayPeer  = "peer0.org1.example.com"
)

func InvokeChaincode() *client.Contract {
	clientConnection := newGrpcConnection()

	id := newIdentity()
	sign := newSign()

	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}

	chaincodeName := "motelcontract"
	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	channelName := "mychannel"
	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}

	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	return contract
}

func newGrpcConnection() *grpc.ClientConn {
	certificatePEM, err := os.ReadFile(tlsCertPath)
	if err != nil {
		panic(fmt.Errorf("failed to read TLS certifcate file: %w", err))
	}

	certificate, err := identity.CertificateFromPEM(certificatePEM)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, gatewayPeer)

	connection, err := grpc.NewClient(peerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

func newIdentity() *identity.X509Identity {
	certificatePEM, err := readFirstFile(certPath)
	if err != nil {
		panic(fmt.Errorf("failed to read certificate file: %w", err))
	}

	certificate, err := identity.CertificateFromPEM(certificatePEM)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(mspID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

func newSign() identity.Sign {
	privateKeyPEM, err := readFirstFile(keyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

func readFirstFile(dirPath string) ([]byte, error) {
	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}

	fileNames, err := dir.Readdirnames(1)
	if err != nil {
		return nil, err
	}

	return os.ReadFile(path.Join(dirPath, fileNames[0]))
}

func ReadAllContracts(contract *client.Contract) ([]*dao.Contract, error) {
	fmt.Println("\n--> Evaluate Transaction: ReadAllContract, returns all contracts on the ledger")

	evaluateResult, err := contract.EvaluateTransaction("ReadAllContract")
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate ReadAllContract transaction: %w", err)
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(evaluateResult, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ReadAllContract JSON: %w", err)
	}

	var contracts []*dao.Contract
	for _, item := range result {
		roomID, _ := strconv.Atoi(item["room_id"].(string))
		lessorID, _ := strconv.Atoi(item["lessor_id"].(string))
		renterID, _ := strconv.Atoi(item["renter_id"].(string))
		status, _ := strconv.Atoi(item["status"].(string))
		payment, _ := strconv.ParseFloat(item["payment"].(string), 64)
		startDate, _ := time.Parse("2006-01-02", item["start_date"].(string))
		datePay, _ := time.Parse("2006-01-02", item["date_pay"].(string))
		payMode, _ := strconv.Atoi(item["pay_mode"].(string))

		contract := &dao.Contract{
			BaseModel: dao.BaseModel{
				ID: uint(item["id"].(float64)),
			},
			RoomID:    uint(roomID),
			LessorID:  uint(lessorID),
			RenterID:  uint(renterID),
			Status:    status,
			PayMode:   payMode,
			StartDate: startDate,
			DatePay:   datePay,
			Payment:   payment,
			IsEnable:  item["is_enable"].(bool),
			Title:     item["title"].(string),
		}
		contracts = append(contracts, contract)
	}

	return contracts, nil
}

func CreateContract(contract *client.Contract, contractDAO *dao.Contract) error {
	fmt.Println("\n--> Submit Transaction: CreateContract, creates a new rental contract")

	_, err := contract.SubmitTransaction("CreateContract",
		strconv.Itoa(int(contractDAO.RoomID)),
		strconv.Itoa(int(contractDAO.LessorID)),
		strconv.Itoa(int(contractDAO.RenterID)),
		strconv.Itoa(contractDAO.Status),
		strconv.Itoa(int(contractDAO.ID)),
		strconv.Itoa(contractDAO.PayMode),
		contractDAO.StartDate.Format("2006-01-02"),
		contractDAO.DatePay.Format("2006-01-02"),
		strconv.FormatFloat(contractDAO.Payment, 'f', -1, 64),
		strconv.FormatBool(contractDAO.IsEnable),
		contractDAO.Title,
		strconv.Itoa(contractDAO.RentalDuration),
		strconv.FormatFloat(contractDAO.Deposit, 'f', -1, 64),
	)
	if err != nil {
		return fmt.Errorf("failed to submit CreateContract transaction: %w", err)
	}
	return nil
}

func CreateHashContract(contract *client.Contract, contractDAO *dao.HashContract) error {
	fmt.Println("\n--> Submit Transaction: CreateHashContract, creates new asset with ID, color, owner, size, and appraisedValue arguments")

	_, err := contract.SubmitTransaction("CreateHashContract",
		strconv.Itoa(int(contractDAO.ID)),
		strconv.Itoa(int(contractDAO.ContractID)),
		contractDAO.Hash,
	)
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
		return err
	}

	return nil
}

func ReadContract(contract *client.Contract, id int64) (*dao.Contract, error) {
	fmt.Println("\n--> Evaluate Transaction: ReadContract, returns a contract by ID")

	evaluateResult, err := contract.EvaluateTransaction("ReadContract", strconv.Itoa(int(id)))
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate ReadContract transaction: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(evaluateResult, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ReadContract JSON: %w", err)
	}

	roomID, _ := strconv.Atoi(result["room_id"].(string))
	lessorID, _ := strconv.Atoi(result["lessor_id"].(string))
	renterID, _ := strconv.Atoi(result["renter_id"].(string))
	status, _ := strconv.Atoi(result["status"].(string))
	payment, _ := strconv.ParseFloat(result["payment"].(string), 64)
	startDate, _ := time.Parse("2006-01-02", result["start_date"].(string))
	datePay, _ := time.Parse("2006-01-02", result["date_pay"].(string))
	payMode, _ := strconv.Atoi(result["pay_mode"].(string))

	contractResp := &dao.Contract{
		BaseModel: dao.BaseModel{
			ID: uint(id),
		},
		RoomID:    uint(roomID),
		LessorID:  uint(lessorID),
		RenterID:  uint(renterID),
		Status:    status,
		PayMode:   payMode,
		StartDate: startDate,
		DatePay:   datePay,
		Payment:   payment,
		IsEnable:  result["is_enable"].(bool),
		Title:     result["title"].(string),
	}

	return contractResp, nil
}

func UpdateContractStatus(contract *client.Contract, contractID int, status int) error {
	fmt.Println("\n--> Submit Transaction: UpdateContractStatus, updates the status of a contract")

	_, err := contract.SubmitTransaction("UpdateContractStatus", strconv.Itoa(contractID), strconv.Itoa(status))
	if err != nil {
		return fmt.Errorf("failed to submit UpdateContractStatus transaction: %w", err)
	}
	return nil
}

func CreateTransaction(contract *client.Contract, transaction *dao.Transaction) error {
	fmt.Println("\n--> Submit Transaction: CreateTransaction, creating a new transaction on the ledger")

	_, err := contract.SubmitTransaction("CreateTransaction",
		strconv.Itoa(int(transaction.UserID)),
		strconv.Itoa(transaction.TransactionType),
		strconv.Itoa(transaction.Status),
		strconv.FormatFloat(transaction.Amount, 'f', -1, 64),
		strconv.FormatFloat(transaction.BalanceBefore, 'f', -1, 64),
		strconv.FormatFloat(transaction.BalanceAfter, 'f', -1, 64),
		transaction.Description,
		strconv.Itoa(transaction.PaymentMethod),
		transaction.TransactionNo,
	)
	if err != nil {
		return fmt.Errorf("failed to submit transaction: %w", err)
	}
	return nil
}

func ReadTransaction(contract *client.Contract, transactionID string) (*dao.Transaction, error) {
	fmt.Println("\n--> Evaluate Transaction: ReadTransaction, retrieving a transaction from the ledger")

	evaluateResult, err := contract.EvaluateTransaction("ReadTransaction", transactionID)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate transaction: %w", err)
	}

	var transaction dao.Transaction
	err = json.Unmarshal(evaluateResult, &transaction)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal transaction: %w", err)
	}

	return &transaction, nil
}

func UpdateTransaction(contract *client.Contract, transactionID string, status int, description string) error {
	fmt.Println("\n--> Submit Transaction: UpdateTransaction, updating an existing transaction on the ledger")

	_, err := contract.SubmitTransaction("UpdateTransaction",
		transactionID,
		strconv.Itoa(status),
		description,
	)
	if err != nil {
		return fmt.Errorf("failed to submit transaction: %w", err)
	}
	return nil
}

func GetAllTransactions(contract *client.Contract) ([]*dao.Transaction, error) {
	fmt.Println("\n--> Evaluate Transaction: GetAllTransactions, retrieving all transactions from the ledger")

	evaluateResult, err := contract.EvaluateTransaction("GetAllTransactions")
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate transaction: %w", err)
	}

	if len(evaluateResult) == 0 {
		return nil, nil
	}

	var transactions []*dao.Transaction
	err = json.Unmarshal(evaluateResult, &transactions)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal transactions: %w", err)
	}

	return transactions, nil
}

func ReadHashContract(contract *client.Contract, id int64) (*dao.HashContract, error) {
	fmt.Println("\n--> Evaluate Transaction: ReadHashContract, returns a HashContract by ID")

	// Evaluate the ReadHashContract transaction
	evaluateResult, err := contract.EvaluateTransaction("ReadHashContract", strconv.Itoa(int(id)))
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate ReadHashContract transaction: %w", err)
	}

	// Parse the result into a map
	var result map[string]interface{}
	if err := json.Unmarshal(evaluateResult, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ReadHashContract JSON: %w", err)
	}

	// Parse the fields
	hashContract := &dao.HashContract{
		BaseModel: dao.BaseModel{
			ID: uint(id),
		},
		ContractID: uint(result["contract_id"].(float64)),
		Hash:       result["hash"].(string),
	}

	return hashContract, nil
}

// ReadAllHashContracts retrieves all HashContracts from the ledger.
func ReadAllHashContracts(contract *client.Contract) ([]*dao.HashContract, error) {
	fmt.Println("\n--> Evaluate Transaction: ReadAllHashContracts, returns all HashContracts on the ledger")

	// Evaluate the ReadAllHashContracts transaction
	evaluateResult, err := contract.EvaluateTransaction("ReadAllHashContracts")
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate ReadAllHashContracts transaction: %w", err)
	}

	// Parse the result into a list of maps
	var result []map[string]interface{}
	if err := json.Unmarshal(evaluateResult, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ReadAllHashContracts JSON: %w", err)
	}

	var hashContracts []*dao.HashContract
	// Loop through the result and construct HashContract objects
	for _, item := range result {
		hashContract := &dao.HashContract{
			BaseModel: dao.BaseModel{
				ID: uint(item["id"].(float64)),
			},
			ContractID: uint(item["contract_id"].(float64)),
			Hash:       item["hash"].(string),
		}
		hashContracts = append(hashContracts, hashContract)
	}

	return hashContracts, nil
}

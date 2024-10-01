package blockchain

import (
	"awesomeProject/internal/app/domain/dao"
	"bytes"
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
	certPath     = cryptoPath + "/users/User1@org1.example.com/msp/signcerts"
	keyPath      = cryptoPath + "/users/User1@org1.example.com/msp/keystore"
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

func GetAllAssets(contract *client.Contract) {
	fmt.Println("\n--> Evaluate Transaction: GetAllAssets, function returns all the current assets on the ledger")

	evaluateResult, err := contract.EvaluateTransaction("ReadAll")
	if err != nil {
		panic(fmt.Errorf("failed to evaluate transaction: %w", err))
	}
	result := formatJSON(evaluateResult)

	fmt.Printf("*** Result:%s\n", result)
}

func CreateAssets(contract *client.Contract, contractDAO *dao.Contract) error {
	fmt.Println("\n--> Submit Transaction: CreateAsset, creates new asset with ID, color, owner, size, and appraisedValue arguments")

	_, err := contract.SubmitTransaction("Create",
		strconv.Itoa(int(contractDAO.RoomID)),
		strconv.Itoa(int(contractDAO.LessorID)),
		strconv.Itoa(int(contractDAO.RenterID)),
		strconv.Itoa(contractDAO.Status),
		strconv.Itoa(int(contractDAO.ID)),
		strconv.Itoa(contractDAO.PayMode),
		contractDAO.DateRent.Format("2006-01-02"),
		contractDAO.DatePay.Format("2006-01-02"),
		strconv.FormatFloat(contractDAO.Payment, 'f', -1, 64),
		strconv.FormatBool(contractDAO.IsEnable),
	)
	if err != nil {
		panic(fmt.Errorf("failed to submit transaction: %w", err))
		return err
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

func ReadAsset(contract *client.Contract, id int64) (string, error) {
	fmt.Println("\n--> Evaluate Transaction: ReadAsset, function returns an asset with a given assetID")

	evaluateResult, err := contract.EvaluateTransaction("Read", strconv.Itoa(int(id)))
	if err != nil {
		panic(fmt.Errorf("failed to evaluate transaction: %w", err))
	}
	result := formatJSON(evaluateResult)

	return result, nil
}

func ReadHashContract(contract *client.Contract, id int64) (*dao.HashContract, error) {
	fmt.Println("\n--> Evaluate Transaction: ReadHashContract, function returns an asset with a given assetID")

	evaluateResult, err := contract.EvaluateTransaction("ReadHashContract", strconv.Itoa(int(id)))
	if err != nil {
		panic(fmt.Errorf("failed to evaluate transaction: %w", err))
	}

	fmt.Println(formatJSON(evaluateResult))

	var result map[string]interface{}
	if err := json.Unmarshal(evaluateResult, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	idresp, _ := strconv.Atoi(result["id"].(string))
	contractID, _ := strconv.Atoi(result["contract_id"].(string))
	hashContract := &dao.HashContract{
		BaseModel: dao.BaseModel{
			ID: uint(idresp),
		},
		ContractID: uint(contractID),
		Hash:       result["hash"].(string),
	}

	return hashContract, nil
}

func formatJSON(data []byte) string {
	var prettyJSON bytes.Buffer
	fmt.Println(string(data))
	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
		panic(fmt.Errorf("failed to parse JSON: %w", err))
	}
	return prettyJSON.String()
}

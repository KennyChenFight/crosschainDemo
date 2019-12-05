package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/KennyChenFight/crosschainDemo/clinic"
	"github.com/KennyChenFight/crosschainDemo/coordinate"
	"github.com/KennyChenFight/crosschainDemo/hospital"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"math/big"
	"net/http"
)

var clinicClient *ethclient.Client
var coordinateClient *ethclient.Client
var hospitalClient *ethclient.Client

var clinicInstance *clinic.Clinic
var coordinateInstance *coordinate.Coordinate
var hospitalInstance *hospital.Hospital

func main() {
	initDependency()
	
	router := gin.Default()
	
	router.POST("/ClinicSetInfo", func(c *gin.Context) {
		var para struct{
			Pk string `json:"pk"`
			Name string `json:"name"`
			Value string `json:"value"`
		}
		c.ShouldBindJSON(&para)

		privateKey, err := crypto.HexToECDSA(para.Pk)
		if err != nil {
			panic(err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			panic(err)
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

		nonce, err := clinicClient.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			panic(err)
		}

		gasPrice, err := clinicClient.SuggestGasPrice(context.Background())
		if err != nil {
			panic(err)
		}

		auth := bind.NewKeyedTransactor(privateKey)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(300000)
		auth.GasPrice = gasPrice

		key := [32]byte{}
		value := [32]byte{}
		copy(key[:], para.Name)
		copy(value[:], para.Value)

		tx, err := clinicInstance.SetInfo(auth, key, value)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("tx sent: %s", tx.Hash().Hex()),
		})
	})

	router.GET("/ClinicGetInfo", func(c *gin.Context) {
		name := c.Query("name")
		key := [32]byte{}
		copy(key[:], name)

		result, err := clinicInstance.Infos(nil, key)
		if err != nil {
			panic(err)
		}

		var str string
		for _, value := range result {
			if value != 0 {
				str += string(value)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"result": str,
		})
	})

	router.POST("/HospitalSetInfo", func(c *gin.Context) {
		var para struct{
			Pk string `json:"pk"`
			Name string `json:"name"`
			Value string `json:"value"`
		}
		c.ShouldBindJSON(&para)

		privateKey, err := crypto.HexToECDSA(para.Pk)
		if err != nil {
			panic(err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			panic(err)
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

		nonce, err := hospitalClient.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			panic(err)
		}

		gasPrice, err := hospitalClient.SuggestGasPrice(context.Background())
		if err != nil {
			panic(err)
		}

		auth := bind.NewKeyedTransactor(privateKey)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(300000)
		auth.GasPrice = gasPrice

		tx, err := hospitalInstance.SetInfo(auth, para.Name, para.Value)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("tx sent: %s", tx.Hash().Hex()),
		})
	})

	router.GET("/HospitalGetInfo", func(c *gin.Context) {
		name := c.Query("name")

		result, err := hospitalInstance.Infos(nil, name)
		if err != nil {
			panic(err)
		}

		var str string
		for _, value := range result {
			if value != 0 {
				str += string(value)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"result": str,
		})
	})

	router.POST("/ClinicSendTransaction", func(c *gin.Context) {

		go func() {
			contractAddress := common.HexToAddress("0x2031faC48bB3a9Ba72489c471B5294Ad3AD5e33B")
			query := ethereum.FilterQuery{
				Addresses: []common.Address{contractAddress},
			}

			logs := make(chan types.Log)
			sub, err := coordinateClient.SubscribeFilterLogs(context.Background(), query, logs)
			if err != nil {
				panic(err)
			}

			for {
				select {
				case err := <-sub.Err():
					log.Fatal(err)
				case vLog := <-logs:
					fmt.Println("coordinate blockchain:", vLog.TxHash.Hex())
				}
			}
		}()

		var message struct{
			Message string
		}
		c.ShouldBindJSON(&message)

		result, err := hospitalInstance.VerifyMessage(nil, message.Message)
		if err != nil {
			panic(err)
		}

		if result {
			id, err := uuid.NewUUID()
			if err != nil {
				log.Fatal(err)
			}

			privateKey, err := crypto.HexToECDSA("92891ae4dc8b974bcda5ca84f962a4ad8671c17d32d8e9c889c692bf3f2d35d6")
			if err != nil {
				panic(err)
			}

			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				panic(err)
			}

			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

			nonce, err := coordinateClient.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				panic(err)
			}

			gasPrice, err := coordinateClient.SuggestGasPrice(context.Background())
			if err != nil {
				panic(err)
			}

			auth := bind.NewKeyedTransactor(privateKey)
			auth.Nonce = big.NewInt(int64(nonce))
			auth.Value = big.NewInt(0)
			auth.GasLimit = uint64(300000)
			auth.GasPrice = gasPrice

			_, err = coordinateInstance.AddTransaction(auth, id.String(), message.Message)
			if err != nil {
				panic(err)
			}

			c.JSON(http.StatusOK, gin.H{
				"verify": result,
				"uuid": id,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"verify": result,
		})
	})

	router.POST("/ClinicGetTransactionValue", func(c *gin.Context) {
		var body struct{
			Id string
		}
		c.ShouldBindJSON(&body)

		result, err := coordinateInstance.Transactions(nil, body.Id)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"value": result,
		})
	})
	
	router.Run()
}

func initDependency() {
	var err error
	clinicClient, err = ethclient.Dial("ws://localhost:7545")
	coordinateClient, err = ethclient.Dial("ws://localhost:8545")
	hospitalClient, err = ethclient.Dial("ws://localhost:9545")

	if err != nil {
		log.Fatal("connect blockchain error:", err)
	}

	clinicContractAddress := common.HexToAddress("0x7E70d99D8b710bd71095d24bbE8eB9f7F2d4fD7A")
	clinicContract, err := clinic.NewClinic(clinicContractAddress, clinicClient)

	coordinateContractAddress := common.HexToAddress("0x2031faC48bB3a9Ba72489c471B5294Ad3AD5e33B")
	coordinateContract, err := coordinate.NewCoordinate(coordinateContractAddress, coordinateClient)

	hospitalContractAddress := common.HexToAddress("0x213f7B9900ff144D6Ac50722dC0f720a6a19546E")
	hospitalContract, err := hospital.NewHospital(hospitalContractAddress, hospitalClient)

	if err != nil {
		log.Fatal("load contract error:", err)
	}

	clinicInstance = clinicContract
	hospitalInstance = hospitalContract
	coordinateInstance = coordinateContract
}

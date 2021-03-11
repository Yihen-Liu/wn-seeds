/**
 * Description:
 * Author: Yihen.Liu
 * Create: 2021-03-09
 */
package contracts

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"kangaroo-net/common/log"
	"kangaroo-net/ethereum/contracts/seeds"
	"math/big"
	"strings"
)
const(
ETH_MAIN_NET 		= "https://mainnet.infura.io/v3/340d0a0efac940eca9bd94662887f940"
ETH_KOVAN_NET 		= "https://kovan.infura.io/v3/340d0a0efac940eca9bd94662887f940"
ETH_RINKEBY_NET 	="https://rinkeby.infura.io/v3/340d0a0efac940eca9bd94662887f940"
ETH_ROPSTEN_NET 	= "https://ropsten.infura.io/v3/340d0a0efac940eca9bd94662887f940"
ETH_LOCAL_NET		="http://127.0.0.1:7545"
ACCOUNT_KEY 		= "UTC--2021-03-10T09-01-45.480199000Z--f43a2f026909544d9703ddfcb3129140faaef6b1"
ACCOUNT_PWD			= "12345"
SEED_CONTRACT_ADDR	= "0x14269a6C3d81591Fa0eEbfCC6b2B1bED3acf18Fb"
)

func AddSeed(addr string) {
	conn, err := ethclient.Dial(ETH_LOCAL_NET)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	contract, err:=seeds.NewSeeds(common.HexToAddress(SEED_CONTRACT_ADDR),conn)
	if err!=nil{
		log.Fatal("new seed contract init err:", err.Error())
	}

	data, _ := ioutil.ReadFile(ACCOUNT_KEY)

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(string(data)), ACCOUNT_PWD, big.NewInt(3))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor:%v \n", err)
		return
	}

	tx, err:=contract.AddSeed(&bind.TransactOpts{
		From: auth.From,
		Signer: auth.Signer,
		Value: nil },
		common.HexToAddress(addr))
	if err!=nil{
		log.Fatalf("add seed tx err:%v \n", err.Error())
		return
	}
	fmt.Println("tx:",tx.Hash().String())
	log.Debug("add seed tx success, tx-id:", tx.Hash().Hex())
}

func Seeds(index int64) {
	conn, err := ethclient.Dial(ETH_LOCAL_NET)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	contract, err:=seeds.NewSeeds(common.HexToAddress(SEED_CONTRACT_ADDR),conn)
	if err!=nil{
		log.Fatal("new seed contract init err:", err.Error())
		return
	}
	oneSeed, err:=contract.Seeds(nil,big.NewInt(index))
	if err!=nil{
		log.Fatalf("query seed tx err:%v \n", err.Error())
		return
	}

	fmt.Println("query seed tx success, seeds:",oneSeed.String())
}

func SeedNums() uint{
	conn, err := ethclient.Dial(ETH_LOCAL_NET)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	contract, err:=seeds.NewSeeds(common.HexToAddress(SEED_CONTRACT_ADDR),conn)
	if err!=nil{
		log.Fatal("new seed contract init err:", err.Error())
		return uint(0)
	}
	cnt, err:=contract.Count(nil)
	if err!=nil{
		log.Fatalf("count seed tx err:%v \n", err.Error())
		return uint(0)
	}

	fmt.Println("count seed tx success, nums:",cnt.String())
	return cnt.Bit(0)
}

func UpdateSeed(index int64, addr string) {
	conn, err := ethclient.Dial(ETH_LOCAL_NET)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	contract, err:=seeds.NewSeeds(common.HexToAddress(SEED_CONTRACT_ADDR),conn)
	if err!=nil{
		log.Fatal("new seed contract init err:", err.Error())
		return
	}

		data, _ := ioutil.ReadFile(ACCOUNT_KEY)

		auth, err := bind.NewTransactorWithChainID(strings.NewReader(string(data)), ACCOUNT_PWD, big.NewInt(3))
		if err != nil {
			log.Fatalf("Failed to create authorized transactor:%v \n", err)
		}

	tx, err:=contract.UpdateSeed(&bind.TransactOpts{
		From: auth.From,
		Signer: auth.Signer,
		Value: nil },
		big.NewInt(index),
		common.HexToAddress(addr))
	if err!=nil{
		log.Fatalf("update seed tx err:%v \n", err.Error())
		return
	}

	fmt.Println("update seed tx success, seeds:",tx.Hash().String())
}


func DeleteSeed(index int64) {
	conn, err := ethclient.Dial(ETH_LOCAL_NET)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	contract, err:=seeds.NewSeeds(common.HexToAddress(SEED_CONTRACT_ADDR),conn)
	if err!=nil{
		log.Fatal("new seed contract init err:", err.Error())
		return
	}

	data, _ := ioutil.ReadFile(ACCOUNT_KEY)

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(string(data)), ACCOUNT_PWD, big.NewInt(3))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor:%v \n", err)
	}

	tx, err:=contract.DeleteSeed(&bind.TransactOpts{
		From: auth.From,
		Signer: auth.Signer,
		Value: nil },
		big.NewInt(index))
	if err!=nil{
		log.Fatalf("delete seed tx err:%v \n", err.Error())
		return
	}

	fmt.Println("delete seed tx success, seeds:",tx.Hash().String())
}
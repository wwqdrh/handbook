package geth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 与账户相关

// 以太坊地址类似: 0x71c7656ec7ab88b098defb751b7401b5f6d8976f
// 非常重要的地址
func Address(addr string) common.Address {
	address := common.HexToAddress(addr)
	return address
}

// 判断地址是否合法
// 1、使用正则
// 2、判断是账户还是智能合约，当地址上没有字节码时，
func AddressIsValid(addr string) {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	fmt.Printf("is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid: true
	fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid: false

	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol Token (ZRX) smart contract address
	address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract) // is contract: true

	// a random user account address
	address = common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
	bytecode, err = client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract = len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract)
}

// 获取账户的余额
// balanceat传递账户地址与可选的区块号，nil则返回最新的余额
// PengdingBalanceAt则是获取待确认的账户余额是多少
// 账户代币(ECR20)余额，涉及到智能合约(智能合约的创建只能由solidity处理)
// solc --abi erc20.sol 编译为jsonabi
// abigen --abi=erc20_sol_ERC20.abi --pkg=token --out=erc20.go // 使用abigen从abi创建go包
// 导入之后就可以调用任何ERC20方法，查询代币月，公共变量等
func Balance(addr string) (string, error) {
	client, err := Connection("https://cloudflare-eth.com")
	if err != nil {
		return "", err
	}

	account := common.HexToAddress(addr)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	return balance.String(), nil
}

// wallet
// 1、生成随机私钥
// 2、通过私钥派生出公钥
// 3、go-ethereum加密包中的PublicKeyToAddress方法接受公钥并返回公共地址
// 4、sha3.NewLegacyKeccak256加了publickeys之后取12位之后的编码等价于第三条
func Wallet() (string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", err
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return address, nil
	// 等价
	// hash := sha3.NewLegacyKeccak256()
	// hash.Write(publicKeyBytes[1:])
	// fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))
}

// 经过加密了的钱包秘钥，每个文件只能包含一个钱包秘钥对
// 如果要导入keystore，则使用NewKeyStore之后调用Import方法
func CreateKeyStore(p string) (string, error) {
	ks := keystore.NewKeyStore(p, keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		return "", err
	}

	return account.Address.Hex(), nil
}

func ImportKeyStore(f string) (string, error) {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(f)
	if err != nil {
		return "", err
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		return "", err
	}
	if err := os.Remove(f); err != nil {
		return "", err
	}

	return account.Address.Hex(), nil
}

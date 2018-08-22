/***************************************************************************
 *
 * Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
 * @author duanbing(duanbing@baidu.com)
 *
 **************************************************************************/

/**
 * @filename main.go
 * @desc
 * @create time 2018-04-19 15:49:26
**/
package main

import (
	"bufio"
	"errors"
	"fmt"
	"go-ethereum/rlp"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	ec "AmbrVM/core"
	"AmbrVM/server"
	"AmbrVM/state"
	"AmbrVM/types"
	"AmbrVM/utils"
	"AmbrVM/vm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	//common.BytesToHash
	testHash    = common.BytesToHash([]byte("kkk"))
	fromAddress = common.BytesToAddress([]byte("kan"))
	toAddress   = common.BytesToAddress([]byte("kimi"))
	amount      = big.NewInt(0)
	nonce       = uint64(0)
	gasLimit    = big.NewInt(100000)
	coinbase    = fromAddress
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func loadBin(filename string) []byte {
	fmt.Println(testHash, fromAddress, toAddress)
	fmt.Println()

	code, err := ioutil.ReadFile(filename)
	must(err)
	return hexutil.MustDecode("0x" + string(code))
	//return []byte("0x" + string(code))
}

func loadAbi(filename string) abi.ABI {
	abiFile, err := os.Open(filename)
	must(err)
	defer abiFile.Close()
	abiObj, err := abi.JSON(abiFile)
	must(err)
	return abiObj
}

//contract env
type ContractEnv struct {
	abiObj          *abi.ABI
	evm             *vm.EVM
	contractAddress *common.Address
	db              *state.StateDB
	binData         []byte

	//gasLeft uint64
} //end struct

func initEvm(statedb *state.StateDB, data []byte) *vm.EVM {
	msg := ec.NewMessage(fromAddress, &toAddress, nonce,
		amount, gasLimit, big.NewInt(0), data, false)
	cc := ChainContext{}
	ctx := ec.NewEVMContext(msg, cc.GetHeader(testHash, 0), cc, &fromAddress)

	//set balance
	statedb.GetOrNewStateObject(fromAddress)
	statedb.GetOrNewStateObject(toAddress)
	statedb.AddBalance(fromAddress, big.NewInt(1e18))
	balance := statedb.GetBalance(fromAddress)
	fmt.Println("Init total Balance: ", balance)

	config := params.MainnetChainConfig
	logConfig := vm.LogConfig{}
	structLogger := vm.NewStructLogger(&logConfig)
	vmConfig := vm.Config{
		Debug:  true,
		Tracer: structLogger,
		/*, JumpTable: vm.NewByzantiumInstructionSet()*/
	}
	return vm.NewEVM(ctx, statedb, config, vmConfig)
}

func newStateDB() (*state.StateDB, error) {
	config := utils.GetConfig()
	if config == nil {
		return nil, errors.New("invalid config")
	}
	//everytime, we start a new env
	os.Remove(config.General.DBFile)
	edb, err := ethdb.NewLDBDatabase(config.General.DBFile, 1000, 1000)
	if err != nil {
		return nil, err
	}
	db := state.NewDatabase(edb)
	root := common.Hash{}
	return state.New(root, db)
}

func newContractEnv(abiFile string, binFile string) (*ContractEnv, error) {
	abiObj := loadAbi(abiFile)
	data := loadBin(binFile)
	statedb, err := newStateDB()
	if err != nil {
		return nil, err
	}
	evm := initEvm(statedb, data)
	if err != nil {
		return nil, err
	}

	return &ContractEnv{
		abiObj:  &abiObj,
		evm:     evm,
		db:      statedb,
		binData: data,
	}, nil
}

//try to create a non-exist contract
func (p *ContractEnv) createContract() error {
	from2 := vm.AccountRef(fromAddress)
	/*contractCode*/ _, contractAddr, gasleft, vmerr := p.evm.Create(from2,
		p.binData, p.db.GetBalance(fromAddress).Uint64(), big.NewInt(0))
	p.contractAddress = &contractAddr
	//p.gasLeft = gasLeftover
	p.db.SetBalance(fromAddress, big.NewInt(0).SetUint64(gasleft))
	initedBalance := p.db.GetBalance(fromAddress)
	fmt.Println("After created contract Gas: ", initedBalance)
	//check contract code & data
	/*fmt.Println("----------------------------")
	fmt.Println(contractCode)
	fmt.Println(p.binData) */
	return vmerr
}

//mint => receiver.balance += amount
func (p *ContractEnv) mint(sender, addr common.Address, amount int64) error {
	if p.evm == nil {
		return errors.New("invalid evm instance")
	}
	if amount == 0 {
		return errors.New("invlalid amount")
	}
	from2 := vm.AccountRef(sender)
	/*hash := sender.Hash()
	hex := sender.Hex()
	fmt.Println("xxxxxxxxxxxxxxxx ", hash, hex)
	*/
	input, err := p.abiObj.Pack("mint", sender, big.NewInt(amount))
	if err != nil {
		return err
	}
	balance := p.db.GetBalance(sender)
	/*outputs*/ _, gasleft, vmerr := p.evm.Call(from2,
		*p.contractAddress, input,
		balance.Uint64(), big.NewInt(0))
	p.db.SetBalance(sender, big.NewInt(0).SetUint64(gasleft))
	fmt.Println("after mint gasleft: ", gasleft, p.db.GetBalance(sender))
	//here outputs should be caller, so check it
	//but what the fuck, why design as this
	if vmerr != nil {
		return vmerr
	}
	mintbalance, ex := p.getBalance(sender, sender)
	if ex != nil {
		return ex
	}
	fmt.Println("after mint: mint balance = ", mintbalance)
	/*
		sender2 := common.BytesToAddress(outputs)
		if !bytes.Equal(sender.Bytes(), sender2.Bytes()) {
			return errors.New("caller are not equal to minter!!")
		}*/
	return nil
}

//send function=>..
func (p *ContractEnv) send(from, receiver common.Address, amount int64) error {
	if p.evm == nil {
		return errors.New("invalid evm instance")
	}
	if amount == 0 {
		return errors.New("invlalid amount")
	}
	from2 := vm.AccountRef(from)
	input, err := p.abiObj.Pack("send", receiver, big.NewInt(amount))
	if err != nil {
		return err
	}
	/*outputs*/ _, gasleft, vmerr := p.evm.Call(from2,
		*p.contractAddress, input,
		p.db.GetBalance(from).Uint64(), big.NewInt(0))

	fmt.Println("after send gasleft: ", gasleft)
	//here outputs should be caller, so check it
	//but what the fuck, why design as this
	if vmerr != nil {
		return vmerr
	}
	p.db.SetBalance(receiver, big.NewInt(amount))
	tobalance, ex := p.getBalance(from, receiver)
	if ex != nil {
		return ex
	}
	fmt.Println("after send to balance: ", tobalance)
	/*
		sender := common.BytesToAddress(outputs)
		if !bytes.Equal(sender.Bytes(), from.Bytes()) {
			return errors.New("caller are not equal to minter!!")
		} */
	return nil
}

//getbalance() returns the real balance @chainnode
func (p *ContractEnv) getBalance(sender, addr common.Address) (int64, error) {
	if p.evm == nil {
		return 0, errors.New("invalid evm instance")
	}
	senderRef := vm.AccountRef(sender)
	input, err := p.abiObj.Pack("balances", addr)
	if err != nil {
		return 0, err
	}
	outputs, gasleft, vmerr := p.evm.Call(senderRef, *p.contractAddress, input,
		p.db.GetBalance(sender).Uint64(), big.NewInt(0))
	if vmerr != nil {
		return 0, vmerr
	}
	fmt.Println("getbalance gasleft: ", gasleft)
	//try to print outputs
	return big.NewInt(0).SetBytes(outputs).Int64(), nil
}

func (p *ContractEnv) print() {
	/*
		logs := p.db.Logs()
		for _, log := range logs {
			fmt.Printf("%#v\n", log)
			for _, topic := range log.Topics {
				fmt.Printf("topic: %#v\n", topic)
			}
			fmt.Printf("data: %#v\n", log.Data)
		}
	*/

	fmt.Println()
	p.db.ForEachStorage(*p.contractAddress, func(key, value common.Hash) bool {
		fmt.Println(key, value)
		return true
	})
}

func (p *ContractEnv) sync() error {
	root, err := p.db.Commit(true)
	if err != nil {
		return err
	}

	fmt.Println("Root Hash", root.Hex())
	return p.db.Database().TrieDB().Commit(root, true)
	//return nil
}

func parse(strs []string) (string, []interface{}, error) {
	var cmd string
	count := len(strs)
	if count >= 1 {
		cmd = strs[0]
	}

	cmds := []interface{}{}
	for i := 1; i < count; i++ {
		str := strs[i]
		switch str[0] {
		case 'a':
			newstr := strings.TrimPrefix(str, "a")
			addr := common.BytesToAddress([]byte(newstr))
			cmds = append(cmds, addr)
		case 'i':
			newstr := strings.TrimPrefix(str, "i")
			i64, err := strconv.ParseInt(newstr, 10, 64)
			if err != nil {
				return "", nil, err
			}
			cmds = append(cmds, big.NewInt(i64))
		case 's':
			newstr := strings.TrimPrefix(str, "s")
			cmds = append(cmds, newstr)
		default:
			return "", nil, errors.New("unsupported arg type")
		}
	}

	return cmd, cmds, nil
}

func call(p *ContractEnv, strs []string) error {
	cmd, args, e := parse(strs)
	if e != nil {
		fmt.Println("Error args: ", e)
		return e
	}
	//fmt.Println(cmd, args)
	input, err := p.abiObj.Pack(cmd, args...)
	if err != nil {
		return err
	}
	//fmt.Println(p)

	from2 := vm.AccountRef(fromAddress)
	/*outputs*/
	outputs, gasleft, vmerr := p.evm.Call(from2,
		*p.contractAddress, input,
		1e16, big.NewInt(0))
	//here outputs should be caller, so check it
	//but what the fuck, why design as this
	fmt.Println("gasleft:", gasleft)
	fmt.Println("outputs: ", outputs)
	return vmerr
}

func handleInput(env *ContractEnv) {
	inputReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Please enter function call: (mint kan, kimi)")
		input, err := inputReader.ReadString('\n')
		if len(input) == 0 {
			break
		}
		input = strings.TrimRight(input, "\n")
		input = strings.TrimRight(input, "\r")

		if err == nil {
			//fmt.Printf("The input was: %s", input)
			strs := strings.Split(input, " ")
			fmt.Println("input cmd: ", strs)
			if len(strs) <= 0 {
				fmt.Println("no args")
				break
			}

			if strs[0] == "quit" || strs[0] == "exit" {
				break
			}
			ex := call(env, strs)
			if ex != nil {
				fmt.Println("call failed: ", ex)
			}
		} else {
			fmt.Println("error: ", err)
		}
	}
}

func main() {
	abiFileName := "./coin_sol_Coin.abi"
	binFileName := "./coin_sol_Coin.bin"
	config := utils.GetConfig()
	if config == nil {
		log.Fatal("invalid config")
		return
	}
	abiFileName = config.General.ABI
	binFileName = config.General.BIN

	contractEnv, err := newContractEnv(abiFileName, binFileName)
	must(err)
	err = contractEnv.createContract()
	must(err)
	handleInput(contractEnv)
	/*
		err = contractEnv.mint(fromAddress, toAddress, 9999)
		must(err)

		err = contractEnv.send(fromAddress, toAddress, 2222)
		must(err)
		// get balance
		contractEnv.getBalance(fromAddress, toAddress)
		contractEnv.getBalance(fromAddress, fromAddress)
	*/
	// get event
	err = contractEnv.sync()
	must(err)

	contractEnv.print()

	toAddress2 := common.BytesToAddress([]byte("cao"))
	bg := toAddress2.Big()
	bs, _ := rlp.EncodeToBytes(toAddress2)
	fmt.Println("--------", bg, bs)
	err = contractEnv.send(fromAddress, toAddress2, 2222)
	//err = contractEnv.sync()
	//must(err)
	contractEnv.print()

	/*
		mdb2, err := ethdb.NewLDBDatabase(dataPath, 100, 100)
		defer mdb2.Close()
		must(err)
		db2 := state.NewDatabase(mdb2)
		statedb2, err := state.New(root, db2)
		must(err)
		testBalance = statedb2.GetBalance(fromAddress)
		fmt.Println("get testBalance =", testBalance)
		if !bytes.Equal(contractCode, statedb2.GetCode(contractAddr)) {
			fmt.Println("BUG!,the code was changed!")
			os.Exit(-1)
		}
		getVariables(statedb2, contractAddr)
	*/

	server.RunServer()
}

func getVariables(statedb *state.StateDB, hash common.Address) {
	cb := func(key, value common.Hash) bool {
		fmt.Printf("key=%x,value=%x\n", key, value)
		return true
	}

	statedb.ForEachStorage(hash, cb)
}

func Print(outputs []byte, name string) {
	fmt.Printf("method=%s, output=%x\n", name, outputs)
}

type ChainContext struct{}

func (cc ChainContext) GetHeader(hash common.Hash, number uint64) *types.Header {

	return &types.Header{
		// ParentHash: common.Hash{},
		// UncleHash:  common.Hash{},
		Coinbase: fromAddress,
		//	Root:        common.Hash{},
		//	TxHash:      common.Hash{},
		//	ReceiptHash: common.Hash{},
		//	Bloom:      types.BytesToBloom([]byte("duanbing")),
		Difficulty: big.NewInt(1),
		Number:     big.NewInt(1),
		GasLimit:   1000000,
		GasUsed:    0,
		Time:       big.NewInt(time.Now().Unix()),
		Extra:      nil,
		//MixDigest:  testHash,
		//Nonce:      types.EncodeNonce(1),
	}
}

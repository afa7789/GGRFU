package main

import (
	"fmt"
	"math/big"

	"github.com/afa7789/GGRFU/internal/mycontract"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

const (
	infura = "wss://mainnet.infura.io/ws/v3/da27eda8e9194fdfbf85cd072e98064a"
	pair   = "0xbb2b8038a1640196fbe3e38816f3e67cba72d940"
	BTC    = 1e8
)

func weiToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(params.Ether))
}

func satoshiToBitcoin(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(BTC))
}

func main() {
	println("The Challenge")
	var test common.Address = common.HexToAddress(pair)
	type Output struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	}
	var reserves Output

	// backend bind.ContractBackend
	client, _ := ethclient.Dial(infura)

	pair, _ := mycontract.NewMycontract(test, client)
	pair.BalanceOf(nil, test)

	reserves, err := pair.GetReserves(nil)
	if err != nil {
		println(err)
	}

	token0Address, err := pair.Token0(nil)
	if err != nil {
		println(err)
	}

	token0Contract, _ := mycontract.NewMycontract(token0Address, client)
	token0Name, _ := token0Contract.Name(nil)
	token0Symbol, _ := token0Contract.Symbol(nil)
	token0Decimals, _ := token0Contract.Decimals(nil)

	token1Address, err := pair.Token1(nil)
	if err != nil {
		println(err)
	}
	token1Contract, _ := mycontract.NewMycontract(token1Address, client)
	token1Name, _ := token1Contract.Name(nil)
	token1Symbol, _ := token1Contract.Symbol(nil)
	token1Decimals, _ := token1Contract.Decimals(nil)

	// fmt.Printf("%+v\n", token1Name)

	k := new(big.Int).Mul(reserves.Reserve1, reserves.Reserve0)

	ammount := big.NewInt(100000000000000)

	// EtherDecimal
	// >> 1wbtc
	// >> 1eth
	// y - (ammount) *x = k
	// x = k / y - (ammount)
	//5107248059731487396846
	// 		 100000000000000
	new1 := new(big.Int).Div(
		k,
		new(big.Int).Sub(
			reserves.Reserve1,
			ammount,
		),
	)
	output0 := new(big.Int).Sub(new1, reserves.Reserve0)

	fmt.Printf("%s reserve: \t %s %s; token decimals:%d;\n", token0Name, token0Symbol, satoshiToBitcoin(reserves.Reserve0).String(), token0Decimals)
	fmt.Printf("%s reserve: \t %s %s; token decimals:%d;\n", token1Name, token1Symbol, weiToEther(reserves.Reserve1).String(), token1Decimals)
	fmt.Println("-----")                // constant
	fmt.Printf("k: %+v;\n", k.String()) // constant
	fmt.Println("-----")                // constant
	fmt.Printf("Ammount to simulate swap: %s %+v;\n", token1Symbol, weiToEther(ammount).String())
	fmt.Printf("new reserve0 = %+v;\nold reserve0 = %+v;\nammount out = %v;\n",
		new1.String(),
		reserves.Reserve0,
		satoshiToBitcoin(output0).Text('f', 8),
	)

}

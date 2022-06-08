![banner image of a bazer on newsprint style](/resources/bazar11_4.png)
# GGRFU - Golang get Reserves from Uniswap ( Mock swap ) 
This is a simple repo, trying out the GETH interface to talk to Ethereum EVM.

## Challenge
Create a simple service in Golang(preferred) or Typescript that interacts(just reads) with UniswapV2 on Ethereum.   

There should be 2 inputs:

- An Ethereum address of a Uniswap V2 pair(for e.g "0xd3d2e2692501a5c9ca623199d38826e513033a17")
- An amount

The program should fetch from the chain the liquidity(reserves) of the pair's tokens and calculate(offchain) the price of the first(base/from) asset in the pool with respect to the second(quote/to), using Uniswap's formula x * y =k.
It should output a message that includes the token names of the pair along with the input/output amounts(e.g "2 ETH <-> 808.121 UNI")

__Links:__

- Pairs link: https://v2.info.uniswap.org/pairs
- You can check the pairs' ABI/contract [here](https://etherscan.io/address/0xd3d2e2692501a5c9ca623199d38826e513033a17#code). It should be identical for every pair.
- You can cross-check the final result [here](https://app.uniswap.org/#/swap?chain=mainnet).
- Some extra reading on how Uniswap V2 works: https://docs.uniswap.org/protocol/V2/concepts/protocol-overview/how-uniswap-works

__More about Uniswap:__

https://docs.uniswap.org/protocol/V2/reference/smart-contracts/pair

### Running it

#### Prerequisites

    - Install [GoLang](https://golang.org/dl/)
    - Install [Docker](https://www.docker.com/community-edition)

#### Setting up and running it.

```bash
    git clone github.com/afa7789/GGRFU && cd GGRFU
    go mod tidy
    cp .env.example .env
    # add your infura or other WSS to it
    #  run this to compile the contract
    make compile
    # run  it
    make run
```
### Output

I used this pair: __0xbb2b8038a1640196fbe3e38816f3e67cba72d940__

timestamp: 1652804597

```BASH
The Challenge
Wrapped BTC reserve:     WBTC 350.2066285; token decimals:8;
Wrapped Ether reserve:   WETH 5117.883978; token decimals:18;
-----
k: 179231689313875619843013296611368;
-----
Amount to simulate swap: WETH 0.0001;
new reserve0 = 35020663536;
old reserve0 = 35020662852;
amount out = 0.00000684;
```

Compare to the value in the uniswap interface at same timestamp:

![uniswap printscreen of swap with same value](/resources/swap.png)

## Tree

```sh
.
├── contract
│   ├── IUniswapV2Pair.sol
│   └── output # directory where the abi gen is gonna output the abi and bin of the contract.
├── go.mod
├── go.sum
├── internal
│   └── mycontract # you will not find the contract here when you download the repo, because you gotta run make compile to generate the go file
│       └── IUniswapV2Pair.go
├── LICENSE
├── main.go
├── Makefile
├── README.md
└── .env.example # copy this to add your keys.
```


UniswapModule | 0x9ac91386b4c6300d600c54882145d8eabdd1b5fd
The Contract Address 0x9ac91386b4c6300d600c54882145d8eabdd1b5fd 
SushiSwapModule | 0xbf1e8e287579afad68cc82fc801c3c07b8c5ae59
CurveModule | 0xbfd0d27fb6a6d3f509143c31a390a53069760ca7
EIP173ProxyWithReceive | 0x8a71704f0f8891f92417faf840568ecee0be0918
Slingshot | 0x71727089c52efc4ad7c3fb6565c865463418d7af

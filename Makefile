current_dir = $(shell pwd)

compile:
	@echo "compiling contract"
	@mkdir -p internal/mycontract
	@docker run -v ${current_dir}/contract/:/sources \
		ethereum/solc:0.5.16 \
		-o /sources/output \
		--abi \
		--bin /sources/IUniswapV2Pair.sol
	@abigen --bin=${current_dir}/contract/output/IUniswapV2Pair.bin \
		--abi=${current_dir}/contract/output/IUniswapV2Pair.abi \
		--pkg=mycontract \
		--out ${current_dir}/internal/mycontract/IUniswapV2Pair.go

run:
	@echo "running contract"
	@go run .
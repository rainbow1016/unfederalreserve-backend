#!/bin/sh

echo "Generating Comptroller..."
mkdir -p abi/comptroller
docker run -v `pwd`/abi:/abi -i ethereum/client-go:alltools-latest abigen --bin /abi/Comptroller.bin --abi - --pkg abicomptroller --alias _mintGuardianPaused=VarMintGuardianPaused,_borrowGuardianPaused=VarBorrowGuardianPaused < abi/Comptroller.abi > abi/comptroller/abi.go

echo "Generating Oracle..."
mkdir -p abi/oracle
docker run -v `pwd`/abi:/abi -i ethereum/client-go:alltools-latest abigen --bin /abi/Oracle.bin --abi - --pkg abioracle < abi/Oracle.abi > abi/oracle/abi.go

echo "Generating Timelock..."
mkdir -p abi/timelock
docker run -v `pwd`/abi:/abi -i ethereum/client-go:alltools-latest abigen --bin /abi/Timelock.bin --abi - --pkg abitimelock < abi/Timelock.abi > abi/timelock/abi.go

echo "Generating CToken..."
mkdir -p abi/ctoken
docker run -v `pwd`/abi:/abi -i ethereum/client-go:alltools-latest abigen --bin /abi/CToken.bin --abi - --pkg abictoken < abi/CToken.abi > abi/ctoken/abi.go

echo "Generating CErc20Delegator..."
mkdir -p abi/cerc20delegator
docker run -v `pwd`/abi:/abi -i ethereum/client-go:alltools-latest abigen --bin /abi/CErc20Delegator.bin --abi - --pkg abicerc20delegator < abi/CErc20Delegator.abi > abi/cerc20delegator/abi.go

echo "Generating CErc20Delegate..."
mkdir -p abi/cerc20delegate
docker run -v `pwd`/abi:/abi -i ethereum/client-go:alltools-latest abigen --bin /abi/CErc20Delegate.bin --abi - --pkg abicerc20delegate < abi/CErc20Delegate.abi > abi/cerc20delegate/abi.go

echo "Generating FaucetToken..."
mkdir -p abi/faucettoken
docker run -v `pwd`/abi:/abi -i ethereum/client-go:alltools-latest abigen --bin /abi/FaucetToken.bin --abi - --pkg abictoken < abi/FaucetToken.abi > abi/faucettoken/abi.go

echo "Generating CEther..."
mkdir -p abi/cether
docker run -v `pwd`/abi:/abi -i ethereum/client-go:alltools-latest abigen --bin /abi/CEther.bin --abi - --pkg abicether < abi/CEther.abi > abi/cether/abi.go

echo "Generating Lens..."
mkdir -p abi/lens
docker run -v `pwd`/abi:/abi -i ethereum/client-go:alltools-latest abigen --bin /abi/Lens.bin --abi - --pkg abilens < abi/Lens.abi > abi/lens/abi.go

echo "Generating LiquidityLens..."
mkdir -p abi/liquiditylens
docker run -v `pwd`/abi:/abi -i ethereum/client-go:alltools-latest abigen --bin /abi/LiquidityLens.bin --abi - --pkg abiliquiditylens < abi/LiquidityLens.abi > abi/liquiditylens/abi.go

echo "Generating InterestRateModel..."
mkdir -p abi/interestratemodel
docker run -v `pwd`/abi:/abi -i ethereum/client-go:alltools-latest abigen --bin /abi/InterestRateModel.bin --abi - --pkg abiinterestratemodel < abi/InterestRateModel.abi > abi/interestratemodel/abi.go

echo "Generating JumpRateModelV2..."
mkdir -p abi/jumpratemodelv2
docker run -v `pwd`/abi:/abi -i ethereum/client-go:alltools-latest abigen --bin /abi/JumpRateModelV2.bin --abi - --pkg abijumpratemodelv2 < abi/JumpRateModelV2.abi > abi/jumpratemodelv2/abi.go

echo "Generating ChainlinkInterface..."
mkdir -p abi/chainlink
docker run -v `pwd`/abi:/abi -i ethereum/client-go:alltools-latest abigen --abi - --pkg abichainlink < abi/ChainlinkInterface.abi > abi/chainlink/abi.go

echo "Generating Unitroller..."
mkdir -p abi/unitroller
docker run -v `pwd`/abi:/abi -i ethereum/client-go:alltools-latest abigen --bin /abi/Unitroller.bin --abi - --pkg abiunitroller < abi/Unitroller.abi > abi/unitroller/abi.go

echo "Generating GnosisSafe..."
mkdir -p abi/gnosissafe
docker run -v `pwd`/abi:/abi -i ethereum/client-go:alltools-latest abigen --abi - --pkg abignosissafe < abi/GnosisSafe.abi > abi/gnosissafe/abi.go

echo "Generating UniswapRouter..."
mkdir -p abi/uniswaprouter
docker run -v `pwd`/abi:/abi -i ethereum/client-go:alltools-latest abigen --abi - --pkg abiuniswaprouter < abi/UniswapRouter.abi > abi/uniswaprouter/abi.go



package token

import (
	"context"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	*ethclient.Client
}

func (client *Client) GetEthBal(address string) *big.Float {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	value := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return value
}

func (client *Client) GetErc20Bal(contract_address, wallet_address string) *big.Float {
	tokenAddress := common.HexToAddress(contract_address)
	instance, err := NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(wallet_address)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fbal := new(big.Float)
	fbal.SetString(bal.String())

	return new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

}

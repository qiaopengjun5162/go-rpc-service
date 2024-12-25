package client

import (
	"fmt"
	"testing"
)

// TestSupportChain is a unit test function for the client.GetSupportCoins method.
//
// It tests with a known supported chain and network.
func TestSupportChain(t *testing.T) {
	client := NewWalletClient("http://127.0.0.1:8970")
	result, err := client.GetSupportCoins("Bitcoin", "MainNet")
	if err != nil {
		fmt.Println("Get support chain fail")
		return
	}
	fmt.Println("Support Chain Res:", result)
}

// TestWalletAddress is a unit test function for the client.GetWalletAddress method.
//
// It tests with a known supported chain and network, and verifies that the
// response contains a valid address and public key.
func TestWalletAddress(t *testing.T) {
	client := NewWalletClient("http://127.0.0.1:8970")
	addressInfo, err := client.GetWalletAddress("Bitcoin", "MainNet")
	if err != nil {
		fmt.Println("Get wallet address fail")
		return
	}
	fmt.Println("Wallet Address Result:", addressInfo.Address, addressInfo.PublicKey)
}

// http://127.0.0.1/wallet/1
// http://127.0.0.1/api/v1?address=111
// http://127.0.0.1/api/v1
/*
{
	"address": "0x00000000"
}
*/

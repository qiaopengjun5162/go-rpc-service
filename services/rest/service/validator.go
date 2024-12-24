package service

type Validator struct{}

// VerifyWalletAddress verify the chain and network is valid or not
//
// only support Bitcoin and Ethereum for now
// only support MainNet and TestNet for now
func (v *Validator) VerifyWalletAddress(chain, network string) bool {
	if chain != "Bitcoin" && chain != "Ethereum" {
		return false
	}
	if network != "MainNet" && network != "TestNet" {
		return false
	}
	return true
}

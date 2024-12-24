package service

import (
	"github.com/qiaopengjun5162/go-rpc-service/database"
	models "github.com/qiaopengjun5162/go-rpc-service/services/rest/model"
)

type Service interface {
	// GetSupportCoins returns whether the service supports the given chain and network.
	// A "support" response indicates that the service supports the given chain and network.
	// A "not support" response indicates that the service does not support the given chain and network.
	GetSupportCoins(*models.ChainRequest) (*models.SupportChainResponse, error)
	// GetWalletAddress returns the wallet address for the given chain and network.
	// The response object contains the public key and address.
	GetWalletAddress(*models.ChainRequest) (*models.WalletAddressResponse, error)
}

type HandleSrv struct {
	v        *Validator
	keysView database.KeysView
}

func NewHandleSrv(v *Validator, ksv database.KeysView) Service {
	return &HandleSrv{
		v:        v,
		keysView: ksv,
	}
}

// GetSupportCoins checks if the specified blockchain and network are supported.
// It utilizes the Validator to verify the chain and network parameters.
// Returns a SupportChainResponse indicating whether the chain and network are supported.
// If supported, the Support field in the response will be true, otherwise false.
//
// Parameters:
//   - req: A pointer to a ChainRequest object containing the chain and network to be checked.
//
// Returns:
//   - A pointer to a SupportChainResponse object with the Support field set accordingly.
//   - An error if any occurs during the process.
func (h HandleSrv) GetSupportCoins(req *models.ChainRequest) (*models.SupportChainResponse, error) {
	ok := h.v.VerifyWalletAddress(req.Chain, req.Network)
	if ok {
		return &models.SupportChainResponse{
			Support: true,
		}, nil
	} else {
		return &models.SupportChainResponse{
			Support: false,
		}, nil
	}
}

// GetWalletAddress generates a wallet address and associated public key for the
// specified blockchain and network. The generated address and public key are
// returned in a WalletAddressResponse object.
//
// Parameters:
//   - req: A pointer to a ChainRequest object containing the chain and network
//     for which the address and public key should be generated.
//
// Returns:
//   - A pointer to a WalletAddressResponse object containing the generated
//     address and public key.
//   - An error if any occurs during the process.
func (h HandleSrv) GetWalletAddress(*models.ChainRequest) (*models.WalletAddressResponse, error) {
	return &models.WalletAddressResponse{
		PublicKey: "public key",
		Address:   "0x00000000000000000000000",
	}, nil
}

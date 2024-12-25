package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

var errWalletHTTPError = errors.New("wallet http error")

type Address struct {
	PublicKey string
	Address   string
}

type WalletClient interface {
	// GetSupportCoins checks if the specified blockchain and network are supported.
	// It utilizes the service's GetSupportCoins method. The boolean value in the
	// response indicates whether the chain and network are supported.
	//
	// Parameters:
	//   - chain: The name of the blockchain to be checked.
	//   - network: The name of the network to be checked.
	//
	// Returns:
	//   - A boolean indicating whether the chain and network are supported.
	//   - An error if any occurs during the process.
	GetSupportCoins(chain, network string) (bool, error)

	// GetWalletAddress returns the wallet address for the given chain and network.
	// The response object contains the public key and address.
	GetWalletAddress(chain, network string) (*Address, error)
}

type Client struct {
	client *resty.Client
}

// NewWalletClient creates a new WalletClient that connects to the specified URL.
//
// Parameters:
//   - url: The URL of the wallet service to connect to.
//
// Returns:
//   - A pointer to a WalletClient object.
func NewWalletClient(url string) *Client {
	client := resty.New()
	client.SetBaseURL(url)
	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		statusCode := r.StatusCode()
		if statusCode >= 400 {
			method := r.Request.Method
			baseUrl := r.Request.URL
			return fmt.Errorf("%d cannot %s %s: %w", statusCode, method, baseUrl, errWalletHTTPError)
		}
		fmt.Println("baseUrl::", r.Request.Method)
		fmt.Println("method::", r.Request.URL)
		fmt.Println("method::", r.Request.QueryParam)
		return nil
	})
	return &Client{
		client: client,
	}
}

// GetSupportCoins checks if the specified blockchain and network are supported.
// It utilizes the service's GetSupportCoins method. The boolean value in the
// response indicates whether the chain and network are supported.
//
// Parameters:
//   - chain: The name of the blockchain to be checked.
//   - network: The name of the network to be checked.
//
// Returns:
//   - A boolean indicating whether the chain and network are supported.
//   - An error if any occurs during the process.
func (c *Client) GetSupportCoins(chain, network string) (bool, error) {
	res, err := c.client.R().SetQueryParams(map[string]string{
		"chain":   chain,
		"network": network,
	}).SetResult(&SupportChainResponse{}).Get("/api/v1/support_chain")
	if err != nil {
		return false, errors.New("support chain request fail")
	}
	spt, ok := res.Result().(*SupportChainResponse)
	if !ok {
		return false, errors.New("support chain transfer type fail")
	}
	return spt.Support, nil
}

// GetWalletAddress returns the wallet address for the given chain and network.
// The response object contains the public key and address.
//
// Parameters:
//   - chain: The name of the blockchain to be checked.
//   - network: The name of the network to be checked.
//
// Returns:
//   - A pointer to an Address object containing the generated address and public key.
//   - An error if any occurs during the process.
func (c *Client) GetWalletAddress(chain, network string) (*Address, error) {
	res, err := c.client.R().SetQueryParams(map[string]string{
		"chain":   chain,
		"network": network,
	}).SetResult(&WalletAddressResponse{}).Get("/api/v1/wallet_address")
	if err != nil {
		return nil, errors.New("wallet address request fail")
	}
	wap, ok := res.Result().(*WalletAddressResponse)
	if !ok {
		return nil, errors.New("wallet address transfer type fail")
	}
	return &Address{
		PublicKey: wap.PublicKey,
		Address:   wap.Address,
	}, nil
}

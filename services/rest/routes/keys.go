package routes

import (
	"fmt"
	"net/http"

	models "github.com/qiaopengjun5162/go-rpc-service/services/rest/model"
)

// GetSupportCoins handles the HTTP request to check if a specific blockchain and network
// are supported by the service. It extracts the 'chain' and 'network' parameters from the
// query string, constructs a ChainRequest, and calls the service's GetSupportCoins method.
// If the service supports the given chain and network, it responds with a JSON indicating support.
// If an error occurs during processing or in writing the response, it logs the error.
//
// Parameters:
//   - w: The http.ResponseWriter used to write the JSON response.
//   - r: The http.Request containing the chain and network query parameters.
func (h Routes) GetSupportCoins(w http.ResponseWriter, r *http.Request) {
	chain := r.URL.Query().Get("chain")
	network := r.URL.Query().Get("network")
	cr := &models.ChainRequest{
		Chain:   chain,
		Network: network,
	}
	supRet, err := h.svc.GetSupportCoins(cr)
	if err != nil {
		return
	}
	err = jsonResponse(w, supRet, http.StatusOK)
	if err != nil {
		fmt.Println("Error writing response", "err", err.Error())
	}
}

// GetWalletAddress handles the HTTP request to retrieve a wallet address for a specific
// blockchain and network. It extracts the 'chain' and 'network' parameters from the query string,
// constructs a ChainRequest, and calls the service's GetWalletAddress method. The wallet address
// and public key are returned in a JSON response. If an error occurs during processing or in writing
// the response, it logs the error.
//
// Parameters:
//   - w: The http.ResponseWriter used to write the JSON response.
//   - r: The http.Request containing the chain and network query parameters.
func (h Routes) GetWalletAddress(w http.ResponseWriter, r *http.Request) {
	chain := r.URL.Query().Get("chain")
	network := r.URL.Query().Get("network")
	cr := &models.ChainRequest{
		Chain:   chain,
		Network: network,
	}

	addrRet, err := h.svc.GetWalletAddress(cr)
	if err != nil {
		return
	}

	err = jsonResponse(w, addrRet, http.StatusOK)
	if err != nil {
		fmt.Println("Error writing response", "err", err.Error())
	}
}

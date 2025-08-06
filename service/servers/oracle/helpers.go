package oracle

import (
	"github.com/warden-protocol/connect/oracle/types"
)

func ToReqPrices(prices types.Prices) map[string]string {
	reqPrices := make(map[string]string, len(prices))

	for cp, price := range prices {
		intPrice, _ := price.Int(nil)
		reqPrices[cp] = intPrice.String()
	}

	return reqPrices
}

package jobs

import (
	"github.com/buyandship/bso-order-cron/internal/rpc"
)

func CreateShipment() error {
	rpcH := rpc.H{}
	return rpcH.SystemCreateShipmentOrder()
}

package jobs

import "github.com/buyandship/bso-order-cron/internal/rpc"

func CloseOrderJob() error {
	rpcH := rpc.H{}
	return rpcH.SystemCloseExpiredOrder()
}

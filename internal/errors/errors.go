package errors

const (
	ErrCallSystemCreateShipmentOrder = iota + 19000
	ErrCallSystemCloseExpiredOrder
)

var ErrMsg = map[int]string{
	ErrCallSystemCreateShipmentOrder: "call system create shipment order error",
	ErrCallSystemCloseExpiredOrder:   "call system close expired order error",
}

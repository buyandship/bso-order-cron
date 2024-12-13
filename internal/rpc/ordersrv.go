package rpc

import (
	"context"
	"github.com/buyandship/bns-golib/config"
	"github.com/buyandship/bns-golib/errno"
	"github.com/buyandship/bso-order-cron/internal/errors"
	"github.com/buyandship/bso-order/kitex_gen/buyandship/one/order"
	"github.com/buyandship/bso-order/kitex_gen/buyandship/one/order/orderservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

func (h *H) OrderClient() (orderservice.Client, error) {
	cli, err := orderservice.NewClient(
		"bns.one.order", client.WithHostPorts(config.GlobalAppConfig.Client.Services["ordersrv"].Addr))
	if err != nil {
		klog.Errorf("connect to notificationservice error: %v", err)
		return nil, errno.NewErrNo(errno.ErrSystem, errno.ErrMsg[errno.ErrSystem])
	}
	return cli, nil
}

func (h *H) SystemCreateShipmentOrder() error {
	klog.Infof("call SystemCreateShipmentOrder at [%+v]", time.Now())
	cli, err := h.OrderClient()
	if err != nil {
		return err
	}
	resp, err := cli.SystemCreateShipmentOrderService(context.Background(), &order.SystemCreateShipmentOrderReq{})
	if err != nil {
		klog.Errorf("SystemCreateShipmentOrderService failed: %v", err)
		return errno.NewErrNo(errno.ErrRpc, errno.ErrMsg[errno.ErrRpc])
	}
	if resp.GetCode() != 0 {
		klog.Errorf("SystemCreateShipmentOrderService failed with code: %d", resp.GetCode())
		return errno.NewErrNo(errors.ErrCallSystemCreateShipmentOrder, errors.ErrMsg[errors.ErrCallSystemCreateShipmentOrder])
	}
	return nil
}

func (h *H) SystemCloseExpiredOrder() error {
	klog.Infof("call SystemCloseExpiredOrder at [%+v]", time.Now())
	cli, err := h.OrderClient()
	if err != nil {
		return err
	}
	resp, err := cli.SystemCloseExpiredOrderService(context.Background(), &order.SystemCloseExpiredOrderReq{})
	if err != nil {
		klog.Errorf("SystemCloseExpiredOrder failed: %v", err)
		return errno.NewErrNo(errno.ErrRpc, errno.ErrMsg[errno.ErrRpc])
	}
	if resp.GetCode() != 0 {
		klog.Errorf("SystemCloseExpiredOrder failed with code: %d", resp.GetCode())
		return errno.NewErrNo(errors.ErrCallSystemCloseExpiredOrder, errors.ErrMsg[errors.ErrCallSystemCloseExpiredOrder])
	}
	return nil
}

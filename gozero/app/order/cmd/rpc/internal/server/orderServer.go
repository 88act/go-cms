// Code generated by goctl. DO NOT EDIT!
// Source: order.proto

package server

import (
	"context"

	"looklook/app/order/cmd/rpc/internal/logic"
	"looklook/app/order/cmd/rpc/internal/svc"
	"looklook/app/order/cmd/rpc/pb"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedOrderServer
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}

// 下订单
func (s *OrderServer) AddOrder(ctx context.Context, in *pb.AddOrderReq) (*pb.AddOrderResp, error) {
	l := logic.NewAddOrderLogic(ctx, s.svcCtx)
	return l.AddOrder(in)
}

// 订单详情
func (s *OrderServer) OrderDetail(ctx context.Context, in *pb.OrderDetailReq) (*pb.OrderDetailResp, error) {
	l := logic.NewOrderDetailLogic(ctx, s.svcCtx)
	return l.OrderDetail(in)
}

// 更新订单
func (s *OrderServer) UpdateOrder(ctx context.Context, in *pb.UpdateOrderReq) (*pb.UpdateOrderResp, error) {
	l := logic.NewUpdateOrderLogic(ctx, s.svcCtx)
	return l.UpdateOrder(in)
}

// 更新支付状态
func (s *OrderServer) UpdatePayStatus(ctx context.Context, in *pb.UpdatePayStatusReq) (*pb.UpdatePayStatusResp, error) {
	l := logic.NewUpdatePayStatusLogic(ctx, s.svcCtx)
	return l.UpdatePayStatus(in)
}

// 更新订单状态
func (s *OrderServer) UpdateOrderStatus(ctx context.Context, in *pb.UpdateOrderStatusReq) (*pb.UpdateOrderStatusResp, error) {
	l := logic.NewUpdateOrderStatusLogic(ctx, s.svcCtx)
	return l.UpdateOrderStatus(in)
}

// 订单列表
func (s *OrderServer) OrderList(ctx context.Context, in *pb.OrderListReq) (*pb.OrderListResp, error) {
	l := logic.NewOrderListLogic(ctx, s.svcCtx)
	return l.OrderList(in)
}

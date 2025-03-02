package service

import (
	"context"
	"fmt"
	"sync"
	"github.com/google/uuid"
	test "order_service/pkg/api/test/api"
)

type Service struct {
    test.UnimplementedOrderServiceServer
    mu     sync.Mutex
    orders map[string]*test.Order
}

func New() *Service {
	return &Service{
		orders: make(map[string]*test.Order),
	}
}

func generateOrderID() string {
	return uuid.New().String()
}

func (s *Service) CreateOrder(ctx context.Context, req *test.CreateOrderRequest) (*test.CreateOrderResponse, error) {
	s.mu.Lock() 
	defer s.mu.Unlock()

	order := &test.Order{
		Id:     generateOrderID(),
		Item:   req.Item,
		Quantity: req.Quantity,
	}

	s.orders[order.Id] = order

	return &test.CreateOrderResponse{Id: order.Id}, nil
}

func (s *Service) GetOrder(ctx context.Context, req *test.GetOrderRequest) (*test.GetOrderResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, ok := s.orders[req.Id]
	if !ok {
		return nil, fmt.Errorf("order %s not found", req.Id)
	}

	return &test.GetOrderResponse{Order: order}, nil
}

func (s *Service) UpdateOrder(ctx context.Context, req *test.UpdateOrderRequest) (*test.UpdateOrderResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, ok := s.orders[req.Id]
	if !ok {
		return nil, fmt.Errorf("order %s not found", req.Id)
	}

	order.Item = req.Item
	order.Quantity = req.Quantity

	return &test.UpdateOrderResponse{Order: order}, nil
}

func (s *Service) DeleteOrder(ctx context.Context, req *test.DeleteOrderRequest) (*test.DeleteOrderResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.orders[req.Id]; !ok {
		return nil, fmt.Errorf("order %s not found", req.Id)
	}

	delete(s.orders, req.Id)

	return &test.DeleteOrderResponse{}, nil
}

func (s *Service) ListOrders(ctx context.Context, req *test.ListOrdersRequest) (*test.ListOrdersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var orders []*test.Order
	for _, order := range s.orders {
		orders = append(orders, order)
	}

	return &test.ListOrdersResponse{Orders: orders}, nil
}
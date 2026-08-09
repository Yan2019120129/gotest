package zhifubao_t

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// OrderStatus 订单状态
type OrderStatus string

// 订单状态枚举
const (
	OrderStatusCreated  OrderStatus = "CREATED"  // 已创建，等待支付
	OrderStatusPaid     OrderStatus = "PAID"     // 已支付
	OrderStatusRefunded OrderStatus = "REFUNDED" // 已退款
)

// Order 商户订单，模拟数据库中的订单记录
type Order struct {
	OutTradeNo  string      // 商户订单号
	Subject     string      // 订单标题
	TotalAmount string      // 订单金额（元）
	Status      OrderStatus // 订单状态
	CreatedAt   time.Time   // 创建时间
	PaidAt      time.Time   // 支付时间
}

// OrderStore 内存订单存储，模拟商户订单库，支持并发访问
type OrderStore struct {
	mu     sync.RWMutex
	orders map[string]*Order
	seq    int64 // 订单号自增序号
}

// NewOrderStore 创建订单存储
func NewOrderStore() *OrderStore {
	return &OrderStore{orders: make(map[string]*Order)}
}

// Create 创建新订单并生成唯一商户订单号，返回订单对象
func (s *OrderStore) Create(subject, totalAmount string) *Order {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.seq++
	order := &Order{
		OutTradeNo:  fmt.Sprintf("%s%06d", time.Now().Format("20060102150405"), s.seq),
		Subject:     subject,
		TotalAmount: totalAmount,
		Status:      OrderStatusCreated,
		CreatedAt:   time.Now(),
	}
	s.orders[order.OutTradeNo] = order
	return order
}

// Get 按商户订单号查询订单
func (s *OrderStore) Get(outTradeNo string) (*Order, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	order, ok := s.orders[outTradeNo]
	return order, ok
}

// MarkPaid 将订单标记为已支付，重复通知时保持幂等
func (s *OrderStore) MarkPaid(outTradeNo string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	order, ok := s.orders[outTradeNo]
	if !ok {
		return errors.New("订单不存在: " + outTradeNo)
	}
	if order.Status == OrderStatusCreated {
		order.Status = OrderStatusPaid
		order.PaidAt = time.Now()
	}
	return nil
}

// MarkRefunded 将订单标记为已退款
func (s *OrderStore) MarkRefunded(outTradeNo string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	order, ok := s.orders[outTradeNo]
	if !ok {
		return errors.New("订单不存在: " + outTradeNo)
	}
	order.Status = OrderStatusRefunded
	return nil
}

package ui

import "sync"

// AppState 管理应用全局状态
type AppState struct {
	mu              sync.RWMutex
	pricePerMillion float64
	onPriceChanged  []func(float64)
}

// NewAppState 创建新的应用状态
func NewAppState() *AppState {
	return &AppState{
		pricePerMillion: 0,
		onPriceChanged:  make([]func(float64), 0),
	}
}

// SetPrice 设置单价并通知所有订阅者
func (s *AppState) SetPrice(price float64) {
	s.mu.Lock()
	s.pricePerMillion = price
	s.mu.Unlock()

	// 通知所有订阅者
	for _, callback := range s.onPriceChanged {
		callback(price)
	}
}

// GetPrice 获取当前单价
func (s *AppState) GetPrice() float64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.pricePerMillion
}

// Subscribe 订阅单价变化事件
func (s *AppState) Subscribe(callback func(float64)) {
	s.onPriceChanged = append(s.onPriceChanged, callback)
}

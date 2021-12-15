package navigation

import (
	"fmt"
)

type OrderType uint64

const (
	OrderTypeUnknown = iota
	OrderTypeForward
	OrderTypeDown
	OrderTypeUp
)

type Order struct {
	Type  OrderType
	Delta int
}

type Summary struct {
	HPosition int
	Depth     int
	Aim       int
}

func (s Summary) String() string {
	return fmt.Sprintf("(%v) HPos: %v, Depth: %v, Aim: %v",
		s.HPosition*s.Depth,
		s.HPosition,
		s.Depth,
		s.Aim)
}

type Navigation interface {
	ProcessOrders(orders ...Order) (Summary, error)
}

func NewSimpleNavigation() Navigation {
	return &simpleNavigation{}
}

type simpleNavigation struct {
}

func (s *simpleNavigation) ProcessOrders(orders ...Order) (Summary, error) {
	summary := Summary{}

	for _, order := range orders {
		switch order.Type {
		case OrderTypeForward:
			summary.HPosition += order.Delta
		case OrderTypeDown:
			summary.Depth += order.Delta
		case OrderTypeUp:
			summary.Depth -= order.Delta
		default:
			return summary, fmt.Errorf("unknown command input: %v", order.Type)
		}
	}

	return summary, nil
}

func NewAimNavigation() Navigation {
	return &aimNavigation{}
}

type aimNavigation struct {
}

func (a *aimNavigation) ProcessOrders(orders ...Order) (Summary, error) {
	summary := Summary{}

	for _, order := range orders {
		switch order.Type {
		case OrderTypeForward:
			summary.HPosition += order.Delta
			summary.Depth += summary.Aim * order.Delta
		case OrderTypeDown:
			summary.Aim += order.Delta
		case OrderTypeUp:
			summary.Aim -= order.Delta
		default:
			return summary, fmt.Errorf("unknown command input: %v", order.Type)
		}
	}

	return summary, nil
}

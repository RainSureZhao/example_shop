package main

import (
	"context"
	stock "example_shop/kitex_gen/example/shop/stock"
	"example_shop/rpc/stock/service"
)

// StockServiceImpl implements the last service interface defined in the IDL.
type StockServiceImpl struct{}

// GetItemStock implements the StockServiceImpl interface.
func (s *StockServiceImpl) GetItemStock(ctx context.Context, req *stock.GetItemStockRequest) (resp *stock.GetItemStockResponse, err error) {
	return service.GetItemStock(ctx, req)
}

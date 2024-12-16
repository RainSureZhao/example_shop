package service

import (
	"context"
	"example_shop/kitex_gen/example/shop/stock"
)

func GetItemStock(ctx context.Context, req *stock.GetItemStockRequest) (resp *stock.GetItemStockResponse, err error) {
	resp = &stock.GetItemStockResponse{
		Stock: req.GetItemId(),
	}
	return resp, nil
}

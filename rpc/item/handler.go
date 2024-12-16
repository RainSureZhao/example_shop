package main

import (
	"context"
	item "example_shop/kitex_gen/example/shop/item"
	"example_shop/kitex_gen/example/shop/stock"
	"example_shop/kitex_gen/example/shop/stock/stockservice"
	"example_shop/rpc/item/service"
	"github.com/cloudwego/kitex/client"
	"log"
)

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct {
	stockCli stockservice.Client
}

func NewStockClient(addr string) (stockservice.Client, error) {
	return stockservice.NewClient("example.shop.stock", client.WithHostPorts(addr))
}

// GetItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItem(ctx context.Context, req *item.GetItemRequest) (resp *item.GetItemResponse, err error) {
	resp, err = service.GetItem(ctx, req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	stockResp, err := s.stockCli.GetItemStock(ctx, &stock.GetItemStockRequest{
		ItemId: req.GetId(),
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println(stockResp)
	resp.Item.SetStock(stockResp.GetStock())
	return resp, nil
}

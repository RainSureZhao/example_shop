package service

import (
	"context"
	"example_shop/kitex_gen/example/shop/item"
)

func GetItem(ctx context.Context, req *item.GetItemRequest) (resp *item.GetItemResponse, err error) {
	resp = item.NewGetItemResponse()
	resp.Item = item.NewItem()
	resp.Item.Id = req.GetId()
	resp.Item.Title = "Kitex"
	resp.Item.Description = "Kitex is an excellent framework!"
	resp.Item.Stock = 10000
	return resp, nil
}

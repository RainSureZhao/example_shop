package main

import (
	"context"
	item "example_shop/kitex_gen/example/shop/item"
)

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct{}

// GetItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItem(ctx context.Context, req *item.GetItemRequest) (resp *item.GetItemResponse, err error) {
	// TODO: Your code here...
	return
}

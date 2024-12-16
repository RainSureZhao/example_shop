// Code generated by Kitex v0.12.0. DO NOT EDIT.

package stockservice

import (
	"context"
	stock "example_shop/kitex_gen/example/shop/stock"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	GetItemStock(ctx context.Context, req *stock.GetItemStockRequest, callOptions ...callopt.Option) (r *stock.GetItemStockResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kStockServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kStockServiceClient struct {
	*kClient
}

func (p *kStockServiceClient) GetItemStock(ctx context.Context, req *stock.GetItemStockRequest, callOptions ...callopt.Option) (r *stock.GetItemStockResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetItemStock(ctx, req)
}

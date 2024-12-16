namespace go example.shop.stock

include "base.thrift"

struct GetItemStockRequest {
    1: required i64 ItemId
}

struct GetItemStockResponse {
    1: i64 stock
    255: base.BaseResp BaseResp
}

service StockService {
    GetItemStockResponse GetItemStock(1: GetItemStockRequest req)
}
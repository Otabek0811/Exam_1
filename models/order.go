package models

type OrderPrimaryKey struct {
	Id string `json:"id"`
}

type Order struct {
	Id        string `json:"id"`
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Sum       int    `json:"sum"`
	Count     int    `json:"count"`
	DateTime  string `json:"date_time"`
	Status    bool `json:"status"`
}

type OrderGetList struct {
	Count  int
	Orders []*Order
}
type CreateOrder struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
	Status    bool `json:"status"`
}

type OrderGetListRequest struct {
	Offset int
	Limit  int
}

type UpdateOrder struct {
	Id        string `json:"id"`
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Sum       int    `json:"sum"`
	Count     int    `json:"count"`
	DateTime  string `json:"date_time"`
	Status    bool `json:"status"`
}


type OrderPayment struct {
	OrderId string `json:"order_id"`
}

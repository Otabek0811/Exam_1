package jsondb

import (
	"app/config"
	"app/controller"
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"time"

	"github.com/google/uuid"
)

type OrderRepo struct {
	fileName string
	file     *os.File
}

func NewOrderRepo(fileName string, file *os.File) *OrderRepo {
	return &OrderRepo{
		fileName: fileName,
		file:     file,
	}
}
func (o *OrderRepo) Create(ord *models.CreateOrder) (*models.Order, error) {
	orders, err := o.read()
	if err != nil {
		return nil, err
	}
	cfg := config.Load()
	strg, err := NewConnectionJSON(&cfg)
	if err != nil {
		panic("Failed connect to json:" + err.Error())
	}
	con := controller.NewController(&cfg, strg)

	byid, err := con.ProductGetById(&models.ProductPrimaryKey{Id: ord.ProductId})
	if err != nil {
		log.Printf("Error while GetById: %+v\n", err)
	}

	var (
		id    = uuid.New().String()
		order = models.Order{
			Id:        id,
			ProductId: ord.ProductId,
			UserId:    ord.UserId,
			Count:     ord.Count,
			Sum:       ord.Count * byid.Price,
			DateTime:  time.Now().Format("2006-01-02 15:04:05"),
			Status:    ord.Status,
		}
	)
	if ord.Count>9{
		order = models.Order{
			Id:        id,
			ProductId: ord.ProductId,
			UserId:    ord.UserId,
			Count:     ord.Count,
			Sum:       (ord.Count-1) * byid.Price,
			DateTime:  time.Now().Format("2006-01-02 15:04:05"),
			Status:    ord.Status,
	}
}

	orders[id] = order
	err = o.write(orders)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *OrderRepo) GetById(ord *models.OrderPrimaryKey) (*models.Order, error) {

	orders, err := o.read()
	if err != nil {
		return nil, err
	}

	if _, ok := orders[ord.Id]; !ok {
		return nil, errors.New("not found")
	}

	order := orders[ord.Id]

	return &order, nil
}

func (o *OrderRepo) GetList(ord *models.OrderGetListRequest) (*models.OrderGetList, error) {
	var resp = &models.OrderGetList{}
	resp.Orders = []*models.Order{}

	orderMap, err := o.read()
	if err != nil {
		return nil, err
	}

	resp.Count = len(orderMap)
	for _, val := range orderMap {
		orders := val
		resp.Orders = append(resp.Orders, &orders)
	}

	

	

	return resp, nil
}

func (o *OrderRepo) Update(ord *models.UpdateOrder) (*models.Order, error) {

	orders, err := o.read()
	if err != nil {
		return nil, err
	}

	if _, ok := orders[ord.Id]; !ok {
		return nil, errors.New("Order not found!")
	}

	orders[ord.Id] = models.Order{
		Id:        ord.Id,
		ProductId: ord.ProductId,
		UserId:    ord.UserId,
		Sum:       ord.Sum,
		Count:     ord.Count,
		DateTime:  ord.DateTime,
		Status:    ord.Status,
	}

	err = o.write(orders)
	if err != nil {
		return nil, err
	}

	order := orders[ord.Id]

	return &order, nil
}

func (o *OrderRepo) Delete(ord *models.OrderPrimaryKey) error {
	orders, err := o.read()
	if err != nil {
		return err
	}
	delete(orders, ord.Id)
	err = o.write(orders)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderRepo) read() (map[string]models.Order, error) {
	var (
		orders   []*models.Order
		orderMap = make(map[string]models.Order)
	)

	data, err := ioutil.ReadFile(o.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &orders)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	for _, order := range orders {
		orderMap[order.Id] = *order
	}

	return orderMap, nil
}

func (u *OrderRepo) write(orderMap map[string]models.Order) error {

	var orders []models.Order

	for _, val := range orderMap {
		orders = append(orders, val)
	}

	body, err := json.MarshalIndent(orders, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

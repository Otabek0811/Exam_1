package controller

import (
	"errors"
	"fmt"

	"log"

	"app/config"
	"app/models"
	"app/pkg/convert"
)

func (c *Controller) OrderCreate(req *models.CreateOrder) (*models.Order, error) {

	log.Printf("Order create req: %+v\n", req)

	req.Status = config.OrderStatus["0"]
	resp, err := c.Strg.Order().Create(req)
	if err != nil {
		log.Printf("error while order Create: %+v\n", err)
		return nil, errors.New("invalid data")
	}

	return resp, nil
}

func (c *Controller) GetByIdOrder(req *models.OrderPrimaryKey) (*models.Order, error) {

	resp, err := c.Strg.Order().GetById(req)
	if err != nil {
		log.Printf("error while order GetById: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) OrderGetList(req *models.OrderGetListRequest) (*models.OrderGetList, error) {

	resp, err := c.Strg.Order().GetList(req)
	if err != nil {
		log.Printf("error while order GetList: %+v\n", err)
		return nil, err
	}
	

	
	return resp, nil
}

func (c *Controller) OrderUpdate(req *models.UpdateOrder) (*models.Order, error) {

	resp, err := c.Strg.Order().Update(req)
	if err != nil {
		log.Printf("error while order Update: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) OrderDelete(req *models.OrderPrimaryKey) error {

	err := c.Strg.Order().Delete(req)
	if err != nil {
		log.Printf("error while order Delete: %+v\n", err)
		return err
	}

	return nil
}

func (c *Controller) OrderPayment(req *models.OrderPayment) error {

	order, err := c.Strg.Order().GetById(&models.OrderPrimaryKey{Id: req.OrderId})
	if err != nil {
		log.Printf("error while Order => GetById: %+v\n", err)
		return err
	}

	user, err := c.Strg.User().GetById(&models.UserPrimaryKey{Id: order.UserId})
	if err != nil {
		log.Printf("error while Order => GetById: %+v\n", err)
		return err
	}

	if order.Sum > user.Balance {
		return errors.New("Not enough balance " + user.FirstName + " " + user.LastName)
	}

	order.Status = config.OrderStatus["1"]
	fmt.Println(order.Status)
	user.Balance -= order.Sum

	var updateOrder models.UpdateOrder
	err = convert.ConvertStructToStruct(order, &updateOrder)
	if err != nil {
		log.Printf("error while convertStructToStruct: %+v\n", err)
		return err
	}

	_, err = c.Strg.Order().Update(&updateOrder)
	if err != nil {
		log.Printf("error while order => Update: %+v\n", err)
		return err
	}

	var updateUser models.UpdateUser
	err = convert.ConvertStructToStruct(user, &updateUser)
	if err != nil {
		log.Printf("error while convertStructToStruct: %+v\n", err)
		return err
	}

	_, err = c.Strg.User().Update(&updateUser)
	if err != nil {
		log.Printf("error while Order => Update: %+v\n", err)
		return err
	}

	return nil
}

package controller

import (
	"errors"
	"log"

	"app/models"
)

func (c *Controller) BranchCreate(req *models.CreateBranch) (*models.Branch, error) {

	log.Printf("Branch create req: %+v\n", req)

	resp, err := c.Strg.Branch().Create(req)
	if err != nil {
		log.Printf("error while Branch Create: %+v\n", err)
		return nil, errors.New("invalid data")
	}

	return resp, nil
}

func (c *Controller) BranchGetById(req *models.BranchPrimaryKey) (*models.Branch, error) {

	resp, err := c.Strg.Branch().GetById(req)
	if err != nil {
		log.Printf("error while Branch GetById: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) BranchGetList(req *models.BranchGetListRequest) (*models.BranchGetListResponse, error) {

	resp, err := c.Strg.Branch().GetList(req)
	if err != nil {
		log.Printf("error while Branch GetList: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) BranchUpdate(req *models.UpdateBranch) (*models.Branch, error) {

	resp, err := c.Strg.Branch().Update(req)
	if err != nil {
		log.Printf("error while Branch Update: %+v\n", err)
		return nil, err
	}

	return resp, nil
}

func (c *Controller) BranchDelete(req *models.BranchPrimaryKey) error {

	err := c.Strg.Branch().Delete(req)
	if err != nil {
		log.Printf("error while Branch Delete: %+v\n", err)
		return err
	}

	return nil
}

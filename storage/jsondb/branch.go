package jsondb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"
)

type BranchRepo struct {
	fileName string
	file     *os.File
}

func NewBranchRepo(fileName string, file *os.File) *BranchRepo {
	return &BranchRepo{
		fileName: fileName,
		file:     file,
	}
}

func (b *BranchRepo) Create(req *models.CreateBranch) (*models.Branch, error) {

	branchs, err := b.read()
	if err != nil {
		return nil, err
	}

	var (
		id       = uuid.New().String()
		Branch = models.Branch{
			Id:   id,
			Name: req.Name,
		}
	)
	branchs[id] = Branch

	err = b.write(branchs)
	if err != nil {
		return nil, err
	}

	return &Branch, nil
}

func (b *BranchRepo) GetById(req *models.BranchPrimaryKey) (*models.Branch, error) {

	Branchs, err := b.read()
	if err != nil {
		return nil, err
	}

	if _, ok := Branchs[req.Id]; !ok {
		return nil, errors.New("Branch not found")
	}
	Branch := Branchs[req.Id]

	return &Branch, nil
}

func (b *BranchRepo) GetList(req *models.BranchGetListRequest) (*models.BranchGetListResponse, error) {

	var resp = &models.BranchGetListResponse{}
	resp.Branchs = []*models.Branch{}

	BranchMap, err := b.read()
	if err != nil {
		return nil, err
	}

	resp.Count = len(BranchMap)
	for _, val := range BranchMap {
		Branchs := val
		resp.Branchs = append(resp.Branchs, &Branchs)
	}

	return resp, nil
}

func (b *BranchRepo) Update(req *models.UpdateBranch) (*models.Branch, error) {

	Branchs, err := b.read()
	if err != nil {
		return nil, err
	}

	if _, ok := Branchs[req.Id]; !ok {
		return nil, errors.New("Branch not found")
	}

	Branchs[req.Id] = models.Branch{
		Id:   req.Id,
		Name: req.Name,
	}

	err = b.write(Branchs)
	if err != nil {
		return nil, err
	}
	Branch := Branchs[req.Id]

	return &Branch, nil
}

func (b *BranchRepo) Delete(req *models.BranchPrimaryKey) error {

	Branchs, err := b.read()
	if err != nil {
		return err
	}

	delete(Branchs, req.Id)

	err = b.write(Branchs)
	if err != nil {
		return err
	}

	return nil
}

func (b *BranchRepo) read() (map[string]models.Branch, error) {
	var (
		Branchs   []*models.Branch
		BranchMap = make(map[string]models.Branch)
	)

	data, err := ioutil.ReadFile(b.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &Branchs)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	for _, Category := range Branchs {
		BranchMap[Category.Id] = *Category
	}

	return BranchMap, nil
}

func (b *BranchRepo) write(BranchMap map[string]models.Branch) error {

	var Branchs []models.Branch

	for _, val := range BranchMap {
		Branchs = append(Branchs, val)
	}

	body, err := json.MarshalIndent(Branchs, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(b.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

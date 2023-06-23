package storage

import "app/models"

type StorageI interface {
	User() UserRepoI
	Product() ProductRepoI
	Category() CategoryRepoI
	Order() OrderRepoI
	Branch() BranchRepoI
	
}

type UserRepoI interface {
	Create(*models.CreateUser) (*models.User, error)
	GetById(*models.UserPrimaryKey) (*models.User, error)
	GetList(*models.UserGetListRequest) (*models.UserGetListResponse, error)
	Update(*models.UpdateUser) (*models.User, error)
	Delete(*models.UserPrimaryKey) error
}

type ProductRepoI interface {
	Create(*models.CreateProduct) (*models.Product, error)
	GetById(*models.ProductPrimaryKey) (*models.Product, error)
	GetList(*models.ProductGetListRequest) (*models.ProductGetListResponse, error)
	Update(*models.UpdateProduct) (*models.Product, error)
	Delete(*models.ProductPrimaryKey) error
}

type CategoryRepoI interface {
	Create(*models.CreateCategory) (*models.Category, error)
	GetById(*models.CategoryPrimaryKey) (*models.Category, error)
	GetList(*models.CategoryGetListRequest) (*models.CategoryGetListResponse, error)
	Update(*models.UpdateCategory) (*models.Category, error)
	Delete(*models.CategoryPrimaryKey) error
}

type OrderRepoI interface {
	Create(*models.CreateOrder) (*models.Order, error)
	GetById(*models.OrderPrimaryKey) (*models.Order, error)
	GetList(*models.OrderGetListRequest) (*models.OrderGetList, error)
	Update(*models.UpdateOrder) (*models.Order, error)
	Delete(*models.OrderPrimaryKey) error
}

type BranchRepoI interface {
	Create(*models.CreateBranch) (*models.Branch, error)
	GetById(*models.BranchPrimaryKey) (*models.Branch, error)
	GetList(*models.BranchGetListRequest) (*models.BranchGetListResponse, error)
	Update(*models.UpdateBranch) (*models.Branch, error)
	Delete(*models.BranchPrimaryKey) error
}



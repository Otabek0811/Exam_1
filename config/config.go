package config

type Config struct {
	DefaultOffset int
	DefaultLimit  int

	Path             string
	UserFileName     string
	CategoryFileName string
	ProductFileName  string
	OrderFileName    string
	BranchFileName   string
}

var OrderStatus = map[string]bool{
	"0": false,
	"1": true,
}

func Load() Config {

	cfg := Config{}

	cfg.DefaultOffset = 0
	cfg.DefaultLimit = 10

	cfg.Path = "./data"
	cfg.UserFileName = "/user.json"
	cfg.CategoryFileName = "/category.json"
	cfg.ProductFileName = "/product.json"
	cfg.OrderFileName = "/shop_chart.json"
	cfg.BranchFileName="/branch.json"

	return cfg
}

package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsondb"
	"fmt"
	"sort"
	"time"
)

func User(con *controller.Controller) {
	con.UserCreate(&models.CreateUser{
		FirstName: "Jony",
		LastName:  "Tim",
		Balance:   200_000,
	})
}
func main() {
	cfg := config.Load()
	strg, err := jsondb.NewConnectionJSON(&cfg)
	if err != nil {
		panic("Failed connect to json:" + err.Error())
	}
	con := controller.NewController(&cfg, strg)

	/*	Task-1	*/
	Task_1(con)

	/*	Task-2	*/
	// Task_2(con)

	/*	Task-3	*/
	// user_id:="76d58bdf-bca7-4484-846c-36383327f03c"
	// Task_3(con,user_id)

	/*	Task-4	*/
	// user_id:="76d58bdf-bca7-4484-846c-36383327f03c"
	// Task_4(con,user_id)

	/*	Task-5	*/
	// Task_5(con)

	/*	Task-6/7	*/
	// Task_6_7(con)

	/*	Task-8	*/
	// Task_8(con)

	/*    Task-9      */
	// Task_9(con)

	/*		Task - 10*/
	// Task_10(con)

	/*   Task -11   */
	// Task_11(con)

}

func Task_11(con *controller.Controller) {
	con.OrderCreate(&models.CreateOrder{
		UserId:    "204ff9b0-3f4e-41b3-a436-3a1fce028fa6",
		ProductId: "46a65fb0-650b-4b4a-a878-ce0d1539adb2",
		Count:     12,
	})

}

func Task_1(con *controller.Controller) {
	respOrder, err := con.OrderGetList(&models.OrderGetListRequest{
		Offset: 0,
		Limit:  10,
	})

	if err != nil {
		fmt.Println(err)
	}

	sort.Slice(respOrder.Orders, func(i, j int) bool {
		return respOrder.Orders[i].DateTime < respOrder.Orders[j].DateTime
	})

	for _, order := range respOrder.Orders {
		fmt.Printf("%+v", order)
		fmt.Println()
	}
}

func Task_2(con *controller.Controller){
	respOrder, err := con.OrderGetList(&models.OrderGetListRequest{
		Offset: 0,
		Limit:  10,
	})

	if err != nil {
		fmt.Println(err)
	}

	a := "2023-06-22 11:21:05"

	b := "2023-06-22 11:22:05"

	time1, _ := time.Parse("2006-01-02 15:04:05", a)
	time2, _ := time.Parse("2006-01-02 15:04:05", b)

	for _, order := range respOrder.Orders {
		x, _ := time.Parse("2006-01-02 15:04:05", order.DateTime)

		if time1.Before(x) && time2.After(x) {
			fmt.Printf("%+v", order)
			fmt.Println()

		}
	}
}

func Task_3(con *controller.Controller,id string){
	respOrder, err := con.OrderGetList(&models.OrderGetListRequest{
		Offset: 0,
		Limit:  10,
	})

	if err != nil {
		fmt.Println(err)
	}
	usrID := id
	fmt.Println(GetUserByID(con,usrID))
	for _, order := range respOrder.Orders {
		if order.UserId == usrID && order.Status {
			fmt.Printf("%+v", order)
			fmt.Println()
		}

	}

}

func Task_4(con *controller.Controller,id string){
	respOrder, err := con.OrderGetList(&models.OrderGetListRequest{
		Offset: 0,
		Limit:  10,
	})

	if err != nil {
		fmt.Println(err)
	}

	usrID:=id
	sum:=0
	for _, order := range respOrder.Orders {
		if order.UserId==usrID && order.Status{
			sum+=order.Sum
		}

	}

	fmt.Printf("%+v")
	fmt.Println(sum)

}

func Task_5(con *controller.Controller){
	respOrder, err := con.OrderGetList(&models.OrderGetListRequest{
		Offset: 0,
		Limit:  10,
	})

	if err != nil {
		fmt.Println(err)
	}
	respProducts,err:=con.ProductGetList(&models.ProductGetListRequest{})

	if err!=nil{
		fmt.Println(err)
	}

	for _,val:= range respProducts.Products{
		count:=0
		for _,val2:=range respOrder.Orders{
			if val.Id==val2.ProductId{
				count+=val2.Count
			}
		}
		fmt.Println(val.Name,count)
	}

}

func Task_6_7(con *controller.Controller){
	respOrder, err := con.OrderGetList(&models.OrderGetListRequest{
		Offset: 0,
		Limit:  10,
	})

	if err != nil {
		fmt.Println(err)
	}

	m := make(map[string]int)

	respProducts, err := con.ProductGetList(&models.ProductGetListRequest{})

	if err != nil {
		fmt.Println(err)
	}

	for _, val := range respProducts.Products {
		count := 0
		for _, val2 := range respOrder.Orders {
			if val.Id == val2.ProductId {
				count += val2.Count
			}
		}
		m[val.Name] = count
	}

	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	for ind, k := range keys {

		fmt.Println(k, m[k])
		if ind==9{
			break
		}
	}

}

func Task_8(con *controller.Controller){
	respOrder, err := con.OrderGetList(&models.OrderGetListRequest{
		Offset: 0,
		Limit:  10,
	})

	if err != nil {
		fmt.Println(err)
	}

	m := make(map[string]int)

	for _, val := range respOrder.Orders {
		count := 0

		for _, val2 := range respOrder.Orders {

			if val.DateTime[:10] == val2.DateTime[:10] {
				count += val2.Count

			}

		}

		val.DateTime = val.DateTime[:10]

		m[val.DateTime] = count
		
	}

	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	for _, k := range keys {

		fmt.Println(k, m[k])
	}

	
}


func Task_10(con *controller.Controller) {
	respOrder, err := con.OrderGetList(&models.OrderGetListRequest{
		Offset: 0,
		Limit:  10,
	})

	if err != nil {
		fmt.Println(err)
	}


	var list []string

	for _, val := range respOrder.Orders {

		if val.Status {
			list = append(list, val.UserId)
		}
	}
	m := make(map[string]int)

	for _, value := range list {
		k := 0
		for _, val := range respOrder.Orders {
			if value == val.UserId && val.Status {
				k += val.Count
			}
		}
		m[value] = k
	}

	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	for _, k := range keys {

		fmt.Println(GetUserByID(con,k),"==> Count Items:",m[k])

		break

	}
}

// /*			Task-----9			*/

func Task_9(con *controller.Controller) {
	respCtg, err := con.CategoryGetList(&models.CategoryGetListRequest{
		Offset: 0,
		Limit:  10,
	})

	if err != nil {
		fmt.Println(err)
	}

	var (
		list []string
	)

	for _, val := range respCtg.Categorys {
		list = Task_9_2(con, val.Id)

		count := Task_9_1(con, list)
		fmt.Println(val.Name, count)
	}
}

func Task_9_1(con *controller.Controller, list []string) int {

	respOrder, err := con.OrderGetList(&models.OrderGetListRequest{
		Offset: 0,
		Limit:  10,
	})

	if err != nil {
		fmt.Println(err)
	}

	m := make(map[string]int)

	for _, productId := range list {
		count := 0
		for _, val := range respOrder.Orders {
			if val.ProductId == productId && val.Status {

				count += val.Count
			}
		}
		m[productId] = count

	}
	sum := 0

	for _, val := range m {
		sum += val
	}
	return sum

}

func Task_9_2(con *controller.Controller, categoryId string) []string {
	respProducts, err := con.ProductGetList(&models.ProductGetListRequest{
		Offset: 0,
		Limit:  10,
	})
	var ls []string
	if err != nil {
		fmt.Println(err)
	}
	for _, val := range respProducts.Products {
		if val.CategoryID == categoryId {

			ls = append(ls, val.Id)

		}
	}
	return ls
}

func GetUserByID(con *controller.Controller, id string) string {
	respUsr, err := con.UserGetById(&models.UserPrimaryKey{
		Id: id,
	})
	if err != nil {
		fmt.Println(err)
	}

	return respUsr.FirstName
}

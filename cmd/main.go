package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsondb"
	"fmt"
	"sort"
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
	OrderGetList(con)

	/*	Task-2	*/
	// OrderGetList(con)

	/*	Task-3	*/
	// OrderGetList(con)

	/*	Task-4	*/
	// OrderGetList(con)

	/*	Task-5	*/
	// OrderGetList(con)

	/*	Task-6/7	*/
	// OrderGetList(con)

	/*	Task-8	*/
	// OrderGetList(con)

	/*    Task-9      */
	// CategoryGetList(con)

	/*		Task - 10*/
	// OrderGetList(con)

	/*   Task -11   */
	// Order(con)

	
}

func Order(con *controller.Controller) {
	con.OrderCreate(&models.CreateOrder{
		UserId:    "204ff9b0-3f4e-41b3-a436-3a1fce028fa6",
		ProductId: "46a65fb0-650b-4b4a-a878-ce0d1539adb2",
		Count:     12,
	})

}

func OrderGetList(con *controller.Controller) {
	respOrder, err := con.OrderGetList(&models.OrderGetListRequest{
		Offset: 0,
		Limit:  10,
	})

	if err != nil {
		fmt.Println(err)
	}

	/*        Task-1                      */
	sort.Slice(respOrder.Orders, func(i, j int) bool {
		return respOrder.Orders[i].DateTime < respOrder.Orders[j].DateTime
	})

	for _, order := range respOrder.Orders {
		fmt.Printf("%+v", order)
		fmt.Println()
	}

	/*        Task_2     */

	// a := "2023-06-22 11:21:05"

	// b := "2023-06-22 11:22:05"

	// time1, _ := time.Parse("2006-01-02 15:04:05", a)
	// time2, _ := time.Parse("2006-01-02 15:04:05", b)

	// for _, order := range respOrder.Orders {
	// 	x, _ := time.Parse("2006-01-02 15:04:05", order.DateTime)

	// 	if time1.Before(x) && time2.After(x) {
	// 		fmt.Printf("%+v", order)
	// 		fmt.Println()

	// 	}
	// }

	/*			Task-3					*/

	// usrID := "76d58bdf-bca7-4484-846c-36383327f03c"
	// fmt.Println(GetUserByID(con,usrID))
	// for _, order := range respOrder.Orders {
	// 	if order.UserId == usrID && order.Status {
	// 		fmt.Printf("%+v", order)
	// 		fmt.Println()
	// 	}

	// }

	/*			Task-4				*/

	// usrID:="76d58bdf-bca7-4484-846c-36383327f03c"
	// sum:=0
	// for _, order := range respOrder.Orders {
	// 	if order.UserId==usrID && order.Status{
	// 		sum+=order.Sum
	// 	}

	// }

	// fmt.Printf("%+v")
	// fmt.Println(sum)

	/*			Task-5				*/

	// respProducts,err:=con.ProductGetList(&models.ProductGetListRequest{})

	// if err!=nil{
	// 	fmt.Println(err)
	// }

	// for _,val:= range respProducts.Products{
	// 	count:=0
	// 	for _,val2:=range respOrder.Orders{
	// 		if val.Id==val2.ProductId{
	// 			count+=val2.Count
	// 		}
	// 	}
	// 	fmt.Println(val.Name,count)
	// }

	/*			Task-6/7(>,<)		*/

	// m := make(map[string]int)

	// respProducts, err := con.ProductGetList(&models.ProductGetListRequest{})

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// for _, val := range respProducts.Products {
	// 	count := 0
	// 	for _, val2 := range respOrder.Orders {
	// 		if val.Id == val2.ProductId {
	// 			count += val2.Count
	// 		}
	// 	}
	// 	m[val.Name] = count
	// }

	// keys := make([]string, 0, len(m))

	// for key := range m {
	// 	keys = append(keys, key)
	// }
	// sort.SliceStable(keys, func(i, j int) bool {
	// 	return m[keys[i]] > m[keys[j]]
	// })

	// for ind, k := range keys {

	// 	fmt.Println(k, m[k])
	// 	if ind==9{
	// 		break
	// 	}
	// }

	/*		Task---8			*/

	// m := make(map[string]int)

	// for _, val := range respOrder.Orders {
	// 	count := 0

	// 	for _, val2 := range respOrder.Orders {

	// 		if val.DateTime[:10] == val2.DateTime[:10] {
	// 			count += val2.Count

	// 		}

	// 	}

	// 	val.DateTime = val.DateTime[:10]

	// 	m[val.DateTime] = count
	// 	if ind==9{
	// 		break
	// 	}
	// }

	// keys := make([]string, 0, len(m))

	// for key := range m {
	// 	keys = append(keys, key)
	// }
	// sort.SliceStable(keys, func(i, j int) bool {
	// 	return m[keys[i]] > m[keys[j]]
	// })

	// for ind, k := range keys {

	// 	fmt.Println(k, m[k])
	// }

	/*				Task---10          */
	// var list []string

	// for _, val := range respOrder.Orders {

	// 	if val.Status {
	// 		list = append(list, val.UserId)
	// 	}
	// }
	// m := make(map[string]int)

	// for _, value := range list {
	// 	k := 0
	// 	for _, val := range respOrder.Orders {
	// 		if value == val.UserId && val.Status {
	// 			k += val.Count
	// 		}
	// 	}
	// 	m[value] = k
	// }

	// keys := make([]string, 0, len(m))

	// for key := range m {
	// 	keys = append(keys, key)
	// }
	// sort.SliceStable(keys, func(i, j int) bool {
	// 	return m[keys[i]] > m[keys[j]]
	// })

	// for _, k := range keys {

	// 	fmt.Println(GetUserByID(con,k),"==> Count Items:",m[k])

	// 	break

	// }
}

// /*			Task-----9			*/

// func CategoryGetList(con *controller.Controller){
// 	respCtg,err:=con.CategoryGetList(&models.CategoryGetListRequest{
// 		Offset: 0,
// 		Limit: 10,
// 	})

// 	if err!=nil{
// 		fmt.Println(err)
// 	}

// 	var (
// 		list []string

// 	)

// 	for _,val:=range respCtg.Categorys{
// 		list = ProductGetList(con, val.Id)

// 		count  :=OrderGetList(con, list)
// 		fmt.Println(val.Name,count)
// 	}
// }

// func OrderGetList(con *controller.Controller, list []string) int{

// 	respOrder, err := con.OrderGetList(&models.OrderGetListRequest{
// 		Offset: 0,
// 		Limit:  10,
// 	})

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	m := make(map[string]int)

// 	for _, productId := range list {
// 		count := 0
// 		for _, val := range respOrder.Orders {
// 			if val.ProductId == productId && val.Status {

// 				count += val.Count
// 			}
// 		}
// 		m[productId] = count

// 	}
// 	sum := 0

// 	for _, val := range m {
// 		sum += val
// 	}
// 	return sum

// }

// func ProductGetList(con *controller.Controller, categoryId string) []string {
// 	respProducts, err := con.ProductGetList(&models.ProductGetListRequest{
// 		Offset: 0,
// 		Limit:  10,
// 	})
// 	var ls []string
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	for _, val := range respProducts.Products {
// 		if val.CategoryID == categoryId {

// 			ls = append(ls, val.Id)

// 		}
// 	}
// 	return ls
// }

func GetUserByID(con *controller.Controller, id string) string {
	respUsr, err := con.UserGetById(&models.UserPrimaryKey{
		Id: id,
	})
	if err != nil {
		fmt.Println(err)
	}

	return respUsr.FirstName
}

package pattern

import (
	"fmt"
	"sync"
	"time"
)

/*
Фасад - это паттерн, позволяющий упростить работу с некой системой для клиента.
Она подразумевает реализацию некого механизма, который все сам сделает за клиента.

Плюсом считается возможность клиента не вдавться в подробности работы системы и просто запустить механизм без задней мысли.
Минусом - риск появления слишком нагруженного всяким функционалом, божественного объекта.

Например, вспоминая прошлые задания, у нас был stan-server, который возможно было запустить прямо из коробки, не думая о том, что вообще происходит под капотом
*/

type McDonalds_Cashier struct{}

func (employee *McDonalds_Cashier) createOrder(orderID int64, client *Client) {
	var order map[string]interface{} = map[string]interface{}{"id": orderID, "Bic Mac": false, "Franch Fries": false, "Soda": false}
	var wg sync.WaitGroup
	wg.Add(2)
	client.kitchen1.makeBigMac(&wg)
	client.kitchen2.makeFranchFries(&wg)
	client.collector.collectOrder(order, &wg)
}

type McDonalds_kitchen_worker struct{}

func (employee *McDonalds_kitchen_worker) makeBigMac(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	wg.Done()
}

func (employee *McDonalds_kitchen_worker) makeFranchFries(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	wg.Done()
}

type McDonalds_order_collector struct{}

func (employee *McDonalds_order_collector) collectOrder(order map[string]interface{}, wg *sync.WaitGroup) {
	employee.addSoda(order)
	wg.Wait()
	employee.addBigMac(order)
	employee.addFranchFry(order)
	fmt.Printf("Заказ %v  готов\n", order["id"])
}

func (employee *McDonalds_order_collector) addSoda(order map[string]interface{}) {
	time.Sleep(2 * time.Second)
	order["Soda"] = true
}

func (employee *McDonalds_order_collector) addBigMac(order map[string]interface{}) {
	order["Bic Mac"] = true
}

func (employee *McDonalds_order_collector) addFranchFry(order map[string]interface{}) {
	order["Franch Fries"] = true
}

// Сам фасад
type Client struct {
	cashier   McDonalds_Cashier
	kitchen1  McDonalds_kitchen_worker
	kitchen2  McDonalds_kitchen_worker
	collector McDonalds_order_collector
}

func CreateClient() *Client {
	client := Client{
		cashier:   McDonalds_Cashier{},
		kitchen1:  McDonalds_kitchen_worker{},
		kitchen2:  McDonalds_kitchen_worker{},
		collector: McDonalds_order_collector{},
	}

	return &client
}

func (client *Client) makeOrder() {
	client.cashier.createOrder(time.Now().Unix(), client)
}

func main() {
	var client = CreateClient()
	client.makeOrder()
}

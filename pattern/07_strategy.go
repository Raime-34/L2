package pattern

import "fmt"

/*
	Паттерн стратегия позваоляет релизовать возможность смены исполняемых классов,
	реализующих различное поведение.

	+ горячая смена алгоритмов на лету
	+ уход от наследования к делегированию
	- усложнение кода за счет новых классов
	- необходимость в понимании клиентом разницы между стратегиями

	В качетве примера реализации можно привести навигационную систему,
	которая может прокладывать маршрут опираясь на траспорт клиента, на
	загруженность различных путей и т.д.
*/

type Transport interface {
	drive()
}

type Bicycle struct{}

func (b *Bicycle) drive() {
	fmt.Println("Using bicycle")
}

type PublicTransport struct{}

func (t *PublicTransport) drive() {
	fmt.Println("Using public transport")
}

type PersonalTransport struct{}

func (t *PersonalTransport) drive() {
	fmt.Println("Using motoblock")
}

// Структура Human содержит в себе поле Transport
type Human struct {
	name      string
	age       int
	transport Transport
}

func (h *Human) letsDrive() {
	h.transport.drive()
}

func (h *Human) setTransport(transport Transport) {
	h.transport = transport
}

func main() {
	h := Human{
		name:      "Жмышенко Валерий Альбертович",
		age:       54,
		transport: new(Bicycle),
	}

	h.letsDrive()
	h.setTransport(new(PublicTransport))
	h.letsDrive()
	h.setTransport(new(PersonalTransport))
	h.letsDrive()

}

package main

import "fmt"

/*
Паттерн Команда позволяет работать с различынми запросами как с объектами

+ позволяет реализовывать отмену операций
+ позволяет собирать сложную операцию из простых
- нагромождает код

Данный поаттерн понадобится если имеется допустим какие-нибудь кнопки в элементах UI, которые при нажатии должны
что-то сделать. С помощью команды предлагается дать хэндлинг разнообразных команд стороннему объекту, который заранее
знает, что от него требуется
*/

// Отправитель
type DeliveryCompany struct {
	courier Courier
}

func (club *DeliveryCompany) execute(recipient Recipient) {
	club.courier.sendPackage(recipient)
}

// Интерфейс реализующий паттерн команда
type Courier interface {
	sendPackage(recipient Recipient)
}

type LegendaryCourier struct {
	name       string
	thePackage string
}

func (c LegendaryCourier) sendPackage(recipient Recipient) {
	recipient.getPackage(c.thePackage)
}

// Получатель
type Recipient interface {
	getPackage(thePackage string)
}

type SomeRecipient struct {
	name string
}

func (r *SomeRecipient) getPackage(thePackage string) {
	fmt.Printf("%v получил \"%v\" 👍\n", r.name, thePackage)
}

func main() {

	bridges := DeliveryCompany{
		courier: LegendaryCourier{
			name:       "Cэм Портер Бриджес",
			thePackage: "Пицца",
		},
	}

	mokhaveExpress := DeliveryCompany{
		courier: LegendaryCourier{
			name:       "Илья Мэдисон",
			thePackage: "Платиновая фишка",
		},
	}

	dom := SomeRecipient{
		name: "Мр. Хаус",
	}

	baker := SomeRecipient{
		name: "Хиггс Монаган",
	}

	bridges.execute(&baker)

	mokhaveExpress.execute(&dom)

}

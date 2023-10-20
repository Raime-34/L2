package pattern

import (
	"fmt"
	"time"
)

/*
Паттерн цепи обязаностей позволяет разделить какой-либо процесс на подзадачи,
которые в дальнейшем можно вызывать в различном порядка по мере необходимости.
Можно использовать так же в случае, когда подзадача обрабатывает определенный случай

+ позволяет присвоить конкретному объекту определенную функцию, за выполнение которой он ответственен
- если каждый обработчик направлен на работу над опредленными классами, может быть случай того, что для некоторого класса отсутствует обработчик

В качестве примера можно привести процедуру обращения к  техподдержке,
когда в самом начале клиента обслуживает чат-бот и если его функионала недостаточно,
то за дело берется сотрудник тех поддержки, тот в свою очередь может перебросить клиента
еще на другого специалиста и т.д.
*/

// Структура для обраотки
type SomeHuman struct {
	name             string
	checkByEyeDoctor bool
	checkByUrlogist  bool
	checkByEnt       bool
}

// Интерфейс для реализации обработчиков
type Doctor interface {
	execute(human *SomeHuman)
	setNext(d Doctor)
}

// Обработчик 1
type EyeDoctor struct {
	next Doctor
}

func (d *EyeDoctor) execute(human *SomeHuman) {
	time.Sleep(time.Second)
	human.checkByEyeDoctor = true
	fmt.Println("Глазной пройден")
	if d.next == nil {
		return
	}
	d.next.execute(human)
}

func (d *EyeDoctor) setNext(n Doctor) {
	d.next = n
}

// Обработчик 2
type Urologist struct {
	next Doctor
}

func (d *Urologist) execute(human *SomeHuman) {
	time.Sleep(time.Second)
	human.checkByUrlogist = true
	fmt.Println("Уролог пройден")
	if d.next == nil {
		return
	}
	d.next.execute(human)
}

func (d *Urologist) setNext(n Doctor) {
	d.next = n
}

// Обработчик 3
type Ent struct {
	next Doctor
}

func (d *Ent) execute(human *SomeHuman) {
	time.Sleep(time.Second)
	human.checkByEnt = true
	fmt.Println("Лор пройден")
	if d.next == nil {
		return
	}
	d.next.execute(human)
}

func (d *Ent) setNext(n Doctor) {
	d.next = n
}

func main() {

	human := SomeHuman{
		name: "Вася Пупкин",
	}
	urologist := new(Urologist)
	ent := new(Ent)
	eyeDoctor := new(EyeDoctor)

	urologist.setNext(ent)
	ent.setNext(eyeDoctor)

	urologist.execute(&human)
	fmt.Println(human)
}

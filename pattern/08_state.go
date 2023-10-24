package pattern

import "fmt"

/*
	Паттерн позволяет работать со стэйт машиной, без использования кучи условных опреаторов

	+ упрощение кода
	+ распредление различных модлей поведения в зависимости от состояиния
	+ избавляет от кучи условных операторов

	В данном примере имеется структура телефон, у которой имеется 2 различных состояния.
	В зависимости от заряда батареии, она условно проделать какую-либо операцию или просто выкинуть сообщение с ошибкой
*/

type Phone struct {
	charge int

	currentState State

	isLowCharge    State
	isNormalCharge State
}

func (p *Phone) setCharge(c int) {
	p.charge = c

	if c < 25 {
		p.setState(p.isNormalCharge)
	} else {
		p.setState(p.isLowCharge)
	}
}

func (p *Phone) setState(s State) {
	p.currentState = s
}

type State interface {
	takePicture() error
	makeCall() error
	browseInternet() error
}

type isLowCharge struct {
	phone *Phone
}

func (s *isLowCharge) takePicture() error {
	return fmt.Errorf("Низкий заряд батереи, невозможно сделать фото")
}

func (s *isLowCharge) makeCall() error {
	return fmt.Errorf("Низкий заряд батереи, невозможно сделать звонок")
}

func (s *isLowCharge) browseInternet() error {
	return fmt.Errorf("Низкий заряд батереи, невозможно открыть браузер")
}

type isNormalCharge struct {
	phone *Phone
}

func (s *isNormalCharge) takePicture() error {
	fmt.Println("Делаю фото")
	return nil
}

func (s *isNormalCharge) makeCall() error {
	fmt.Println("Делаю вызов")
	return nil
}

func (s *isNormalCharge) browseInternet() error {
	fmt.Println("Открываю браузер")
	return nil
}

func main() {

	phone := Phone{
		isLowCharge:    new(isLowCharge),
		isNormalCharge: new(isNormalCharge),
	}

	phone.setCharge(100)
	if err := phone.currentState.takePicture(); err != nil {
		fmt.Println(err)
	}

	phone.setCharge(10)
	if err := phone.currentState.takePicture(); err != nil {
		fmt.Println(err)
	}

}

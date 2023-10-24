package pattern

import "fmt"

/*
	Фабричный метод необходим для создания объектов разных классов
	Предлагается перенести ответственность по созданию объекта на определенную функцию,
	которая уже сама решит что необходимо создать клиенту

	+ упрощает добавление новых классов
*/

// Базовый интерфейс
type Soda interface {
	setName(name string)
	setManufacturer(manufacturer string)
	setVolume(v int)
	getName() string
	getManufacturer() string
	getVolume() int
}

// Базовая стрктура
type BaseSoda struct {
	name         string
	manufacturer string
	volume       int
}

func (s *BaseSoda) setName(name string) {
	s.name = name
}

func (s *BaseSoda) setManufacturer(manufacturer string) {
	s.manufacturer = manufacturer
}

func (s *BaseSoda) setVolume(volume int) {
	s.volume = volume
}

func (s *BaseSoda) getName() string {
	return s.name
}

func (s *BaseSoda) getManufacturer() string {
	return s.manufacturer
}

func (s *BaseSoda) getVolume() int {
	return s.volume
}

// Структура определенного объекта
type CocaCola struct {
	BaseSoda
}

// Структура определенного объекта
type PepsiCola struct {
	BaseSoda
}

// Конструктор для колы
func newCocaCola() (soda Soda) {
	soda = new(CocaCola)
	soda.setName("Coca cola")
	soda.setManufacturer("ООО «Coca‑Cola Ichimligi Uzbekiston, LTD», Узбекистан")
	soda.setVolume(1)
	return
}

// Конструктор для пепси
func newPepsiCola() (soda Soda) {
	soda = new(PepsiCola)
	soda.setName("Pepsi cola")
	soda.setManufacturer("ООО «Coca‑Cola Ichimligi Uzbekiston, LTD», Узбекистан")
	soda.setVolume(1)
	return
}

// Сам Фабричный метод
func SodaFacroty(brand string) Soda {
	switch brand {
	case "Coca cola":
		return newCocaCola()
	case "Pepsi":
		return newPepsiCola()
	default:
		return nil
	}
}

func main() {
	cc := SodaFacroty("Coca cola")
	pepsi := SodaFacroty("Pepsi")
	// other := SodaFacroty("Блейзер")

	fmt.Println(cc.getName(), cc.getManufacturer())
	fmt.Println(pepsi.getName(), pepsi.getManufacturer())
	// fmt.Println(other.getName(), other.getManufacturer())
}

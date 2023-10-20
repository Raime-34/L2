package pattern

/*
Билдер позвояет реализовывать создание различных варинтов объекта,
при этом не заставляя при вызове заполнять все строки,
например в сниппете ниже билдер позволяет создать как серверный компьютер, так и персональный,
причем (в данном контексте) серверному не понадоиться графический адаптер, звуковая карта и т.п.
Билдер подразумевает декомпозицию процесса создания объекта,
для каждого этапа у нас есть своя функция,
и если кокретный объект подразумевает наличие того или иного поля,
нам неободимо релизовать функцию её инициализации.

+ позвояет контролировать процесс создания объекта пошагово
- нагромождает код

Понадобится в случае необходимости создания множества вариантов одного и того же объекта
*/

type Computer struct {
	GPU    string
	CPU    string
	RAM    string
	SC     string
	MB     string
	Memory string
}

type Builder interface {
	insertGPU()
	insertCPU()
	insertRAM()
	insertSC()
	insertMB()
	insertMemory()
}

type ServerBuilder struct{}

func (b *ServerBuilder) insertCPU(c *Computer, component string) { c.CPU = component }

func (b *ServerBuilder) insertGPU(c *Computer, component string) { c.GPU = component }

func (b *ServerBuilder) insertRAM(c *Computer, component string) { c.RAM = component }

func (b *ServerBuilder) insertMB(c *Computer, component string) { c.MB = component }

func (b *ServerBuilder) insertMemory(c *Computer, component string) { c.Memory = component }

type PCBuilder struct{}

func (b *PCBuilder) insertCPU(c *Computer, component string) { c.CPU = component }

func (b *PCBuilder) insertGPU(c *Computer, component string) { c.GPU = component }

func (b *PCBuilder) insertRAM(c *Computer, component string) { c.RAM = component }

func (b *PCBuilder) insertMB(c *Computer, component string) { c.MB = component }

func (b *PCBuilder) insertMemory(c *Computer, component string) { c.Memory = component }

func (b *PCBuilder) insertSC(c *Computer, component string) { c.SC = component }

type Director struct{}

func (d *Director) sdelayCorito() (computer *Computer) {
	var builder = new(PCBuilder)
	builder.insertCPU(computer, "Intel Core 2 Duo")
	builder.insertGPU(computer, "Игровая видеокарта")
	builder.insertRAM(computer, "2 GB")
	builder.insertMB(computer, "Motherboard")
	builder.insertSC(computer, "Samsung UBFUE1")
	builder.insertMemory(computer, "Samsung WBBRWGWE")
	return
}

func (d *Director) makeServer() (computer *Computer) {
	var builder = new(ServerBuilder)
	builder.insertCPU(computer, "Intel Core 2 Duo")
	builder.insertRAM(computer, "1024 GB")
	builder.insertMemory(computer, "21321312 GB")
	builder.insertMB(computer, "Motherboard")
	return
}

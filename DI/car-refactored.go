
package Engine

type Engine interface {
	Start() error
}

type engine struct {
}

func (e engine) Start() error {
	// start engine
	return nil
}

//////////////////////////////////////////////////////
package Car
type Car struct {
	engine Engine
}

func NewCar() *Car {
	car := Car{}
	car.engine = engine{}
	return &car

}

func (car Car) Start() {
	car.engine.Start()
}


//////////////////////////////////////////////////////
package Car
type Car struct {
	engine Engine
}

func NewCar(eng Engine) *Car {
	car := Car{}
	car.engine = eng
	return &car

}

func (car Car) Start() {
	car.engine.Start()
}

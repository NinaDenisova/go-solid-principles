package Car

type Engine struct {
}

func (e *Engine) Start() error {
	// start engine
	return nil
}

type Car struct {
	engine Engine
}

func NewCar() *Car {
	car := Car{}
	car.engine = Engine{}
	return &car

}

func (car *Car) Start() {
	car.engine.Start()
}

//-----------------------------------------------------------

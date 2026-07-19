package composition

type engine struct {
	hp int
}

func (e engine) HP() int {
	return e.hp
}

type wheel struct {
	wheelDimensions int
}

func (w wheel) WheelDimensions() int {
	return w.wheelDimensions
}

type Car struct {
	CarName string
	Engine  engine
	Wheel   wheel
}

type Bike struct {
	BikeName string
	engine
	wheel
}

func Newcar(carName string, hp, wd int) Car {
	return Car{
		CarName: carName,
		Engine: engine{
			hp: hp,
		},
		Wheel: wheel{
			wheelDimensions: wd,
		},
	}
}

func Newbike(bikeName string, hp, wd int) Bike {
	return Bike{
		BikeName: bikeName,
		engine: engine{
			hp: hp,
		},
		wheel: wheel{
			wheelDimensions: wd,
		},
	}
}

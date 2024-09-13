package objects

type Vehicle interface {
	Startup() string
}

type Car struct {}
type Bike struct {}

func (this * Car) Startup() {
	println("Broom Broom")
}

func (this * Bike) Startup(){
	println("Wee")
}

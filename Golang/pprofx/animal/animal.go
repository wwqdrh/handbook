package animal

import (
	"pprofx/animal/canidae/dog"
	"pprofx/animal/canidae/wolf"
	"pprofx/animal/felidae/cat"
	"pprofx/animal/felidae/tiger"
	"pprofx/animal/muridae/mouse"
)

var (
	AllAnimals = []Animal{
		&tiger.Tiger{},
		&mouse.Mouse{},
		&dog.Dog{},
		&wolf.Wolf{},
		&cat.Cat{},
	}

	CPUAnimal = []Animal{
		&tiger.Tiger{},
	}

	MemoryAnimal = []Animal{
		&mouse.Mouse{},
	}

	AllocsAnimal = []Animal{
		&dog.Dog{},
	}

	GroutineAnimal = []Animal{
		&wolf.Wolf{},
	}

	BlockAnimal = []Animal{
		&cat.Cat{},
	}
)

type Animal interface {
	Name() string
	Live()

	Eat()
	Drink()
	Shit()
	Pee()
}

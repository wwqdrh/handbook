package felidae

import "pprofx/animal"

type Felidae interface {
	animal.Animal
	Climb()
	Sneak()
}

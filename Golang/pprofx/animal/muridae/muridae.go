package muridae

import "pprofx/animal"

type Muridae interface {
	animal.Animal
	Hole()
	Steal()
}

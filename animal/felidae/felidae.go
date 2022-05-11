package felidae

import "github.com/yhmain/go-pprof-practice/animal"

type Felidae interface {
	animal.Animal
	Climb()
	Sneak()
}

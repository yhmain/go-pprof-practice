package canidae

import "github.com/yhmain/go-pprof-practice/animal"

type Canidae interface {
	animal.Animal
	Run()
	Howl()
}

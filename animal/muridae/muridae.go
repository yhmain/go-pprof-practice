package muridae

import "github.com/yhmain/go-pprof-practice/animal"

type Muridae interface {
	animal.Animal
	Hole()
	Steal()
}

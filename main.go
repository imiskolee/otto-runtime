package otto_runtime

import "github.com/robertkrimen/otto"

func Inject(vm *otto.Otto) {
	for _,v := range plugins {
		inject(vm,v)
	}
}


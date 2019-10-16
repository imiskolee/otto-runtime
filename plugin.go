package otto_runtime

import "github.com/robertkrimen/otto"

type Plugin struct {
	Name string
	Value interface{}
}

var plugins []Plugin

func Register(name string,val interface{}) {
	plugins = append(plugins,Plugin{
		Name:  name,
		Value: val,
	})
}

func inject(vm *otto.Otto,plugin Plugin) {
	_ = vm.Set(plugin.Name,plugin.Value)
}



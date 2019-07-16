package main

import (
	"fmt"

	builtin "github.com/hedarikun/neko/builtin"
	evaluator "github.com/hedarikun/neko/evaluator"
)

func main() {
	eva := evaluator.New()
	eva.Global.RegisterFun("print", func(args []builtin.Object) builtin.Object {
		for _, arg := range args {
			fmt.Println(arg.CallMethod("toString", nil).(builtin.StringObject).Value)
		}
		return nil
	})
	eva.StartEvaluate(`
		let x = [25, "hello", true]
		print(1 + x[2])
	`)
	//val, _ := eva.Global.GetVariable("number").(builtin.NumberObject)
	//fmt.Println(val.Value)
}

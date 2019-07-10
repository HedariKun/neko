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
		if 3 == 4 {
			print("hello world")
		} else {
			print("no one is here")
		}
	`)
	// val, _ := eva.Global.GetVariable("x").(builtin.NumberObject)
	// fmt.Print(val.Value)
}

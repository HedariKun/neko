package main

import (
	"fmt"

	evaluator "github.com/hedarikun/neko/evaluator"
)

func main() {
	eva := evaluator.New()
	eva.Global.RegisterFun("print", func(args []evaluator.Object) evaluator.Object {
		for _, arg := range args {
			fmt.Println(arg.CallMethod("toString", nil).(evaluator.StringObject).Value)
		}
		return nil
	})
	eva.StartEvaluate("print('hello world', 12, 52 * 42)")
	// val, _ := eva.Global.GetVariable("x").(evaluator.NumberObject)
	// fmt.Print(val.Value)
}

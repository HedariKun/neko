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
			fmt.Println(arg.CallMethod("toString", nil).(*builtin.StringObject).Value)
		}
		return nil
	})
	eva.StartEvaluate(`
		struct point {
			x,
			y
		}
		let p1 = point.new(200, 300)
		let p2 = point.new(p1.x + 400, 0)
		print(p2.x + p1.y)
	`)
	//val, _ := eva.Global.GetVariable("number").(builtin.NumberObject)
	//fmt.Println(val.Value)
}

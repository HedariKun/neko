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
		impl point {
			fun change(self, x) {
				print(x)
			}
		}
		let p3 = point.new(400)
		p3.change(100)
	`)
	//val, _ := eva.Global.GetVariable("number").(builtin.NumberObject)
	//fmt.Println(val.Value)
}

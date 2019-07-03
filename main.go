package main

import (
	"fmt"

	evaluator "github.com/hedarikun/neko/evaluator"
)

func main() {
	eva := evaluator.New()
	eva.StartEvaluate("let x = 1 + 1")
	val, _ := eva.Global.GetVariable("x").(evaluator.NumberObject)
	fmt.Print(val.Value)
}

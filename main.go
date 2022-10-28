package main

import (
	"algo/equation"
	"algo/graph"
	"algo/integrate"
	m "algo/math"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"

	"github.com/dominikbraun/graph/draw"
)

func integrals_example() {
	var f = m.NewFunctionWithDerivative(
		func(x float64) float64 {
			return 7 / (x*x + 1)
		},
		func(x float64) float64 {
			return -14 * x / math.Pow(x*x+1, 2)
		},
	)

	var h = m.NewFunctionNoDerivatives(
		func(x float64) float64 {
			return x * math.Exp(x)
		},
	)

	fmt.Println(integrate.TrapezoidsN(f, 0, 5, 10))
	fmt.Println(integrate.TrapezoidsDelta(f, 0, 5, 0.00001))
	fmt.Println(integrate.TrapezoidsDelta(h, 0, 2, 0.001))
}

func graph_example() {
	x := graph.NewNamedNode(2.0, "x")
	y := graph.NewNamedNode(2.5, "y")
	c := 1.0
	f := graph.AddNode(x, graph.NegNode(graph.MulNode(y, c)))
	fmt.Println(f)
	fmt.Println(f.GetValue())
}

func graph_example2() {
	x := graph.NewNamedNode(2, "x")
	y := graph.NewNamedNode(2, "y")
	z := graph.NewNamedNode(2, "z")
	f := graph.MulNode(z, graph.AddNode(x, y))
	fmt.Println(f)
	fmt.Println("Результат вычислений: ", f.GetValue())
}

func drawExample(fileName string) {
	// based on https://dominikbraun.io/blog/visualizing-graph-structures-using-go-and-graphviz/
	x := graph.NewNamedNode(2.0, "x")
	y := graph.NewNamedNode(2.5, "y")
	f := graph.AddNode(x, graph.NegNode(graph.MulNode(y, 3.0)))
	g := f.ToGraph()
	file, _ := os.Create(fmt.Sprintf("graph_examples/%v.gv", fileName))
	_ = draw.DOT(g, file)

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	cmd := exec.Command(
		"cmd",
		"/C",
		fmt.Sprintf(
			"cd %v & cd graph_examples & dot -Tpng -O %v.gv & %v.gv.png",
			path,
			fileName,
			fileName,
		),
	)
	err = cmd.Run()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func fExample() {
	var f = m.NewFunctionWithDerivative(
		func(x float64) float64 {
			return x * x / 2
		},
		func(x float64) float64 {
			return x
		},
	)
	fmt.Printf("f(%v) = %v, f'(%v) = %v\n", 3, f.Y(3), 3, f.Dy(3))

	var g = m.NewFunctionNoDerivatives(
		func(x float64) float64 {
			return x * x / 2
		},
	)
	fmt.Printf("g(%v) = %v, g'(%v) = %v\n", 3, g.Y(3), 3, g.Dy(3))
}

func bisectionExample() {
	// https://studfile.net/preview/6736203/page:2/
	var f = m.NewFunctionNoDerivatives(
		func(x float64) float64 {
			// return x*x - math.Sin(x) - 1
			return math.Exp(-x) - math.Sin(x)
		},
	)
	// res := equation.BisectionEps(f, 1, 2, 1e-2)
	res := equation.BisectionEps(f, 0, 1, 1e-5)
	fmt.Println(res)
}

func chordsExample() {
	// https://elar.urfu.ru/bitstream/10995/1054/1/umk_2004_015.pdf
	var f = m.NewFunctionNoDerivatives(
		func(x float64) float64 {
			// return x*x - math.Sin(x) - 1
			return math.Tan(0.93*x+0.43) - math.Pow(x, 2)
		},
	)
	res := equation.ChordsEps(f, -0.4, -0.2, 1e-4)
	fmt.Println(res)
}

func succApproxExample() {
	var f = m.NewFunctionNoDerivatives(
		func(x float64) float64 {
			return math.Cos(x)
		},
	)
	res := equation.SuccessiveApprox(f, 0, 1, 0.1)
	fmt.Println(res)
}

func main() {
	_ = integrals_example
	_ = drawExample
	_ = graph_example
	_ = graph_example2
	_ = integrals_example
	_ = fExample
	_ = bisectionExample
	_ = chordsExample
	_ = succApproxExample
	// drawExample("graph2")
	// graph_example()
	// graph_example2()
	// integrals_example()
	// fExample()
	// bisectionExample()
	// chordsExample()
	succApproxExample()
}

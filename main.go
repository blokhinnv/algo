package main

import (
	"algo/equation"
	"algo/graph"
	"algo/integrate"
	m "algo/math"
	"algo/systems"
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
	res := equation.SuccessiveApprox(f, 0, 1, 0.001)
	fmt.Println(res)
}

func succApproxAdvExample() {
	var f = m.NewFunctionNoDerivatives(
		func(x float64) float64 {
			return math.Cos(x)
		},
	)
	res := equation.SuccessiveApproxAdv(f, 0, 1, 0.001)
	fmt.Println(res)
}

func newtonExample() {
	// https://elar.urfu.ru/bitstream/10995/1054/1/umk_2004_015.pdf
	var f = m.NewFunctionNoDerivatives(
		func(x float64) float64 {
			return math.Tan(0.93*x+0.43) - math.Pow(x, 2)
		},
	)
	res := equation.Newton(f, -0.4, -0.2, 1e-4)
	fmt.Println(res)
}

func gaussExample() {
	A := [][]float64{{4, 2, -1}, {5, 3, -2}, {3, 2, -3}}
	b := []float64{1, 2, 0}

	fmt.Println("Классический метод Гаусса")
	X := systems.SolveGauss(A, b, true)
	fmt.Println("X = ", X)
	fmt.Println()
	fmt.Println("Метод Жордана-Гаусса")
	X = systems.SolveGaussBackward(A, b, true)
	fmt.Println("X = ", X)
}

func gaussExample2() {
	A := [][]float64{{2, 5, 4, 1}, {1, 3, 2, 1}, {2, 10, 9, 7}, {3, 8, 9, 2}}
	b := []float64{20, 11, 40, 37}

	fmt.Println("Классический метод Гаусса")
	X := systems.SolveGauss(A, b, true)
	fmt.Println("X = ", X)
	fmt.Println()
	fmt.Println("Метод Жордана-Гаусса")
	X = systems.SolveGaussBackward(A, b, true)
	fmt.Println("X = ", X)
}

func invExample() {
	// A := [][]float64{{3, -2, 4}, {3, 4, -2}, {2, -1, -1}}
	// A := [][]float64{{1, 2, 4}, {5, 8, 5}, {2, 4, 4}}
	A := [][]float64{{2, 5, 4, 1}, {1, 3, 2, 1}, {2, 10, 9, 7}, {3, 8, 9, 2}}
	// b := []float64{21, 9, 10}
	X := systems.InverseMatrix(A, true)
	fmt.Println("X = ", X)
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
	_ = newtonExample
	_ = succApproxAdvExample
	_ = gaussExample
	_ = gaussExample2
	// drawExample("graph2")
	// graph_example()
	// graph_example2()
	// integrals_example()
	// fExample()
	// bisectionExample()
	// chordsExample()
	// succApproxExample()
	// succApproxAdvExample()
	// newtonExample()
	// gaussExample()
	// gaussExample2()
	invExample()
}

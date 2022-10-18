package main

import (
	"algo/graph"
	"algo/integrate"
	m "algo/math"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"

	"github.com/dominikbraun/graph/draw"
	"gonum.org/v1/gonum/diff/fd"
)

func integrals_example() {
	var f = m.NewFunction(
		func(x float64) float64 {
			return 7 / (x*x + 1)
		},
		func(x float64) float64 {
			return -14 * x / math.Pow(x*x+1, 2)
		},
	)

	var h = m.NewFunction(
		func(x float64) float64 {
			return x * math.Exp(x)
		},
		func(x float64) float64 {
			return 0
		},
	)

	fmt.Println(integrate.TrapezoidsN(f, 0, 5, 10))
	fmt.Println(integrate.TrapezoidsDelta(f, 0, 5, 0.00001))
	fmt.Println(integrate.TrapezoidsDelta(h, 0, 2, 0.001))

	dy := func(x float64) float64 {
		return fd.Derivative(f.Y, x, &fd.Settings{
			Formula: fd.Forward,
			Step:    1e-3,
		})
	}
	fmt.Println(f.Dy(5), dy(5))
}

func graph_example() {
	x := graph.NewNode(2.0, nil, nil, "x", nil)
	y := graph.NewNode(2.5, nil, nil, "y", nil)
	c := 1.0
	f := graph.AddNode(x, graph.NegNode(graph.MulNode(y, c)))
	fmt.Println(f)
	fmt.Println(f.GetValue())

}

func graph_example2() {
	x := graph.NewNode(2, nil, nil, "x", nil)
	y := graph.NewNode(2, nil, nil, "y", nil)
	z := graph.NewNode(2, nil, nil, "z", nil)
	f := graph.MulNode(z, graph.AddNode(x, y))
	fmt.Println(f)
	fmt.Println(f.GetValue())

}

func drawExample(fileName string) {
	// based on https://dominikbraun.io/blog/visualizing-graph-structures-using-go-and-graphviz/
	x := graph.NewNode(2.0, nil, nil, "x", nil)
	y := graph.NewNode(2.5, nil, nil, "y", nil)
	// c := graph.NewNode(1.0, nil, nil, "c", nil)
	// z := graph.NewNode(0.0, nil, nil, "z", nil)
	f := graph.AddNode(x, graph.NegNode(graph.MulNode(y, 1.0)))
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

func main() {
	// drawExample2("graph2")
	// graph_example()
	graph_example2()
}

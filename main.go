package main

import (
	"algo/differential"
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
	A := m.Matrix{{4, 2, -1}, {5, 3, -2}, {3, 2, -3}}
	b := m.Vector{1, 2, 0}

	fmt.Println("Классический метод Гаусса")
	X := systems.SolveGauss(A, b, true)
	fmt.Println("X = ", X)
	fmt.Println()
	fmt.Println("Метод Жордана-Гаусса")
	X = systems.SolveGaussBackward(A, b, true)
	fmt.Println("X = ", X)
}

func gaussExample2() {
	A := m.Matrix{{2, 5, 4, 1}, {1, 3, 2, 1}, {2, 10, 9, 7}, {3, 8, 9, 2}}
	b := m.Vector{20, 11, 40, 37}

	fmt.Println("Классический метод Гаусса")
	X := systems.SolveGauss(A, b, true)
	fmt.Println("X = ", X)
	fmt.Println()
	fmt.Println("Метод Жордана-Гаусса")
	X = systems.SolveGaussBackward(A, b, true)
	fmt.Println("X = ", X)
}

func invExample() {
	// A := [][]float64{{1, 2, 4}, {5, 8, 5}, {2, 4, 4}}
	A := m.Matrix{{2, 5, 4, 1}, {1, 3, 2, 1}, {2, 10, 9, 7}, {3, 8, 9, 2}}
	A_inv := m.InverseMatrix(A, true)
	fmt.Println("A_inv = ", A_inv)
	fmt.Println("A = ", A)
	fmt.Println("A @ A_inv = ", m.MatMul(A, A_inv))
}

func solveMMulExample() {
	A := m.Matrix{{2, 5, 4, 1}, {1, 3, 2, 1}, {2, 10, 9, 7}, {3, 8, 9, 2}}
	b := m.Vector{20, 11, 40, 37}
	X := systems.SolveMmul(A, b, true)
	fmt.Println("X = ", X)
}

func jacobianExample() {
	F := m.VariadicFunction{
		Y: func(fs, x []float64) {
			fs[0] = x[0] + x[1] - 3
			fs[1] = x[0]*x[0] + x[1]*x[1] - 9
		},
		M: 2,
		N: 2,
	}
	// JF = [
	// 	1 1
	// 	2x1 2x2
	// ]
	x := m.Vector{6, 2}
	fmt.Println(F.Jacobian(x))
}

func newtonSystemExample1() {
	F := m.VariadicFunction{
		Y: func(fs, x []float64) {
			fs[0] = x[0] + x[1] - 3
			fs[1] = x[0]*x[0] + x[1]*x[1] - 9
		},
		M: 2,
		N: 2,
	}
	x := systems.SolveNewton(F, m.Vector{1, 5}, 0.001)
	fmt.Println("Ответ ", x)
}

func newtonSystemExample2() {
	F := m.VariadicFunction{
		Y: func(fs, x []float64) {
			fs[0] = math.Pow(x[0], 2) + math.Pow(x[1], 2) + math.Pow(x[2], 2) - 1
			fs[1] = 2*math.Pow(x[0], 2) + math.Pow(x[1], 2) - 4*x[2]
			fs[2] = 3*math.Pow(x[0], 2) - 4*x[1] + math.Pow(x[2], 2)
		},
		M: 3,
		N: 3,
	}
	x := systems.SolveNewton(F, m.Vector{0.5, 0.5, 0.5}, 0.0001)
	fmt.Println("Ответ ", x)
}

func solveRungeKuttaExample() {
	f := func(x, y float64) float64 {
		return x*x - 2*y
	}
	res_euler := differential.SolveEuler(f, 0, 1, 0.1, 1)
	fmt.Printf("Метод Эйлера: \n%v", res_euler)
	res_rk2 := differential.SolveRK2(f, 0, 1, 0.1, 1)
	fmt.Printf("Метод Рунге-Кутта II: \n%v", res_rk2)
	res_rk4 := differential.SolveRK4(f, 0, 1, 0.1, 1)
	fmt.Printf("Метод Рунге-Кутта IV: \n%v", res_rk4)
	true_f := differential.NewTabularFunctionFromF(
		func(x float64) float64 {
			return 3.0/4*math.Exp(-2*x) + 1.0/2*x*x - 1.0/2*x + 1.0/4
		},
		0, 1.1, 0.1,
	)
	fmt.Printf(
		"Ошибка метода Эйлера: %.5f\n",
		true_f.ComputeDissimilarity(res_euler),
	)
	fmt.Printf(
		"Ошибка метода Рунге-Кутта II: %.5f\n",
		true_f.ComputeDissimilarity(res_rk2),
	)
	fmt.Printf(
		"Ошибка метода Рунге-Кутта IV %.5f\n",
		true_f.ComputeDissimilarity(res_rk4),
	)
	differential.DrawTabularFunctions(
		[]differential.TabularFunction{res_euler, res_rk2, res_rk4, true_f},
		[]string{"Euler method", "Runge-Kutta II", "Runge-Kutta IV", "True function"},
		"Runge-Kutta methods",
		24,
		"graph_examples/runge-kutta.png",
	)
}

func solvePredictorCorrectorExample() {
	f := func(x, y float64) float64 {
		return x*x - 2*y
	}
	res_pc := differential.SolvePredictorCorrector(f, 0, 1, 0.05, 1, 0.1)
	_ = res_pc
	fmt.Printf("Метод прогноза и коррекции: \n%v", res_pc)
	res_rk4 := differential.SolveRK4(f, 0, 1, 0.05, 1)
	fmt.Printf("Метод Рунге-Кутта IV: \n%v", res_rk4)
	true_f := differential.NewTabularFunctionFromF(
		func(x float64) float64 {
			return 3.0/4*math.Exp(-2*x) + 1.0/2*x*x - 1.0/2*x + 1.0/4
		},
		0, 1, 0.05,
	)
	fmt.Printf(
		"Ошибка метода прогноза и коррекции: %.5f\n",
		true_f.ComputeDissimilarity(res_pc),
	)
	fmt.Printf(
		"Ошибка метода Рунге-Кутта IV %.5f\n",
		true_f.ComputeDissimilarity(res_rk4),
	)
	res_euler := differential.SolveEuler(f, 0, 1, 0.1, 1)
	differential.DrawTabularFunctions(
		[]differential.TabularFunction{res_euler, res_pc, res_rk4, true_f},
		[]string{
			"Euler method",
			"Prediction-correction method",
			"Runge-Kutta IV",
			"True function",
		},
		"Numerical methods",
		8,
		"graph_examples/prediction-correction.png",
	)
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
	_ = invExample
	_ = solveMMulExample
	_ = jacobianExample
	_ = newtonSystemExample1
	_ = newtonSystemExample2
	_ = solveRungeKuttaExample
	_ = solvePredictorCorrectorExample
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
	// invExample()
	// solveMMulExample()
	// jacobianExample()
	// newtonSystemExample1()
	// newtonSystemExample2()
	// solveRungeKuttaExample()
	solvePredictorCorrectorExample()
}

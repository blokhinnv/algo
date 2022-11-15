package differential

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// Рисует несколько табличных функций
func DrawTabularFunctions(
	tfs []TabularFunction,
	titles []string,
	title string,
	dpi int,
	fname string,
) {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	n := len(tfs)
	vs := make([]any, 0)
	for i := 0; i < n; i++ {
		vs = append(vs, titles[i])
		vs = append(vs, tfs[i].toXYs())
	}
	err := plotutil.AddLinePoints(p, vs...)
	p.Legend.Top = true
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	x_dim := font.Length(dpi) * vg.Inch
	y_dim := font.Length(dpi) * vg.Inch
	if err := p.Save(x_dim, y_dim, fname); err != nil {
		panic(err)
	}
}

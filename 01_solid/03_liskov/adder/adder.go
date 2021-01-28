package adder

type Int struct{}

func (Int) Add(a, b int) int {
	return a + b
}

type Double struct{}

func (Double) Add(a, b float64) float64 {
	return a + b
}

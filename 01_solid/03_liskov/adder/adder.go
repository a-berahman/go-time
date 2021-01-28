package adder

// Int adds two integer values.
type Int struct{}

// Add returns the sum a+b.
func (Int) Add(a, b int) int { return a + b }

// Double adds two double values.
type Double struct{}

// Add returns the sum a+b.
func (Double) Add(a, b float64) float64 { return a + b }

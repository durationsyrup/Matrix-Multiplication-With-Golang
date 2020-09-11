package main

import (
	"fmt"
	"math"
	"os"
)

type mtx map[int][]float64

// No reason I use maps as the matrices
// -> [][]float64 makes a better structure for matrices
// I used a map instead just because (epic prank lol XD)
// You can rebuild this to work without a map as practice! :)

func newMatrix(i ...[]float64) mtx {
	matrix := make(mtx)
	for index, e := range i {
		matrix[index] = e
		index++
	}
	return matrix
}

func matmul(a, b mtx) mtx {
	// matrix Multiplication
	var err error

	product := make(mtx)

	if len(b) != len(a[0]) && len(b[0]) != len(a) {
		err = fmt.Errorf("cannot matrix multiply these matrices")
	}

	for i := 0; i < len(a); i++ {
		for k := 0; k < len(b[0]); k++ {
			val := 0.0
			for j := 0; j < len(b); j++ {
				val += b[j][k] * a[i][j]
			}
			product[i] = append(product[i], val)
		}
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return product
}

// 3d point
func newPoint(x, y, z float64) mtx {
	return newMatrix([]float64{x}, []float64{y}, []float64{z})
}

// rotate 3d point
func rotatePoint(rotate float64, pt mtx, dir rune) mtx {
	rotate = rotate * (math.Pi / 180)
	// degrees to radians
	var rm mtx
	if dir == 'x' {
		// rotate point around the x axis
		rm = newMatrix(
			[]float64{1, 0, 0},
			[]float64{0, math.Round(math.Cos(rotate)), math.Round(-math.Sin(rotate))},
			[]float64{0, math.Round(math.Sin(rotate)), math.Round(math.Cos(rotate))},
		)
	} else if dir == 'y' {
		// rotate point around the y axis
		rm = newMatrix(
			[]float64{math.Round(math.Cos(rotate)), 0, math.Round(math.Sin(rotate))},
			[]float64{0, 1, 0},
			[]float64{math.Round(-math.Sin(rotate)), 0, math.Round(math.Cos(rotate))},
		)
	} else if dir == 'z' {
		// rotate point around the z axis
		rm = newMatrix(
			[]float64{math.Round(math.Cos(rotate)), math.Round(-math.Sin(rotate)), 0},
			[]float64{math.Round(math.Sin(rotate)), math.Round(math.Cos(rotate)), 0},
			[]float64{0, 0, 1},
		)
	}

	return matmul(rm, pt)
}
func mR(i ...float64) []float64 {
	// replaces []float64{n1, n2, ...} with mR(n1,n2,...)
	return i
}

func main() {
	pt := newPoint(1, 0, 0)

	pt = rotatePoint(90, pt, 'z')
	// xyz(1,0,0) -> xyz(0,1,0)
	pt = rotatePoint(90, pt, 'z')
	// xyz(0,1,0) -> xyz(-1,0,0)
	fmt.Println(pt)

	mOne := newMatrix(
		mR(1, 2, 3),
		mR(1, 2, 3),
		mR(1, 2, 3),
	)
	// 3x3
	mTwo := newMatrix(
		mR(1),
		mR(1),
		mR(1),
	)
	// 1x3

	result := matmul(mOne, mTwo)
	fmt.Printf("(mOne * mTwo) = %v\n", result)
}

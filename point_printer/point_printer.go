package point_printer

import (
	"fmt"
)

type PointPrinter interface {
	Print(x, y uint)
}

func CleanScreen() {
	fmt.Printf("\033[H\033[J")
}

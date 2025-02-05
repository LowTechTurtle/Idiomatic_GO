package main

import (
	"fmt"

	format "github.com/LowTechTurtle/Idiomatic_GO/module_pkg_import/package_example/do-format"
	"github.com/LowTechTurtle/Idiomatic_GO/module_pkg_import/package_example/math"
)

func main() {
	num := math.Double(1)
	fmt.Println(format.Number(num))
}

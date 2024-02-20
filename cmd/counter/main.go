package main

import (
	"fmt"

	"github.com/abkhan/gogen/pkg/gen"
	"github.com/abkhan/gogen/pkg/sample_gen/counter"
)

func main() {

	g := gen.New("test", map[string]int{"start": 0, "end": 20, "step": 2}, counter.Counter)

	av, err := g.Exec()
	if err != nil {
		fmt.Printf("Error executing exec: %v\n", err)
	}
	fmt.Printf("Ret: %T, %v\n", av, av)

	av, err = g.Exec()
	if err != nil {
		fmt.Printf("Error executing exec: %v\n", err)
	}
	fmt.Printf("Ret: %T, %v\n", av, av)

	av, err = g.Exec()
	if err != nil {
		fmt.Printf("Error executing exec: %v\n", err)
	}
	fmt.Printf("Ret: %T, %v\n", av, av)

	av, err = g.Exec()
	if err != nil {
		fmt.Printf("Error executing exec: %v\n", err)
	}
	fmt.Printf("Ret: %T, %v\n", av, av)

	// run to read from channel
	vch, err := g.Run()
	if err != nil {
		fmt.Printf("Error executing Run: %v\n", err)
	}
	//fmt.Printf("g.Run ret: %T, %v\n", vch, vch)

	for v := range vch {
		fmt.Printf("Run-Ret: %T, %v\n", v, v)
	}
}

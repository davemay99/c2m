package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	iterations := 500000
	if len(os.Args) > 1 {
		iterations, _ = strconv.Atoi(os.Args[1])
	} else {
		fmt.Println("Missing iterations argument, using default", iterations)
	}

	fmt.Printf("Iterations = %d\n\n", iterations)
	fmt.Println("Starting multicore...")
	fmt.Println("Starting single core...")
	montepi(iterations)
}

func montepi(niter int) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	start := time.Now()
	var count uint64
	for i := 0; i < niter; i++ {
		x := rnd.Float64()
		y := rnd.Float64()
		z := x*x + y*y
		if z <= 1 {
			count++
		}
	}

	pi := float64(count) / float64(niter) * float64(4)
	total := time.Now().Sub(start)

	tps := float64(niter) / total.Seconds()
	fmt.Printf("%s # of trials= %d (%f/s), estimate of pi is %g \n", total, niter, tps, pi)
}

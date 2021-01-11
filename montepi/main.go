package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"sync"
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
	montepi_multicore(iterations)

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
	total := time.Since(start)

	tps := float64(niter) / total.Seconds()
	fmt.Printf("%s # of trials= %d (%f/s), estimate of pi is %g \n", total, niter, tps, pi)
}

// goroutine version from http://tstra.us/code/gopi/
func monte_carlo_pi(radius float64, reps int, result *int, wait *sync.WaitGroup) {
	var x, y float64
	count := 0
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	for i := 0; i < reps; i++ {
		x = random.Float64() * radius
		y = random.Float64() * radius

		if num := math.Sqrt(x*x + y*y); num < radius {
			count++
		}
	}

	*result = count
	wait.Done()
}

func montepi_multicore(samples int) {
	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)

	var wait sync.WaitGroup

	counts := make([]int, cores)

	start := time.Now()
	wait.Add(cores)

	for i := 0; i < cores; i++ {
		go monte_carlo_pi(100.0, samples/cores, &counts[i], &wait)
	}

	wait.Wait()

	total := 0
	for i := 0; i < cores; i++ {
		total += counts[i]
	}

	pi := (float64(total) / float64(samples)) * 4

	fmt.Println("Cores: ", cores)
	fmt.Println("Time: ", time.Since(start))
	fmt.Println("pi: ", pi)
	fmt.Println("")
}

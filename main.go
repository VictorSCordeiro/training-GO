package main

import (
	"fmt"
	"math"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Extract path from the request URL
	path := r.URL.Path[1:] // Removing the leading '/'

	// Write the requested path to the response
	fmt.Fprint(w, path)
}

func handlePrimeRequest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/prime/"):]
	x, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid input for x", http.StatusBadRequest)
		return
	}

	prime := getNthPrime(x)
	fmt.Fprintf(w, "%d", prime)
}

func handleTime10000thPrimeRequest(w http.ResponseWriter, r *http.Request) {
	//path := r.URL.Path[len("/MeasureTime10000thPrime/"):]

	//with the basic version of algorithm where divides by smaller numbers to check if the number is prime
	start := time.Now()
	result := getNthPrime(10000)
	elapsedTime := time.Since(start)

	responseString := fmt.Sprintf("10000th prime: %d\nTime taken to find the 10000th prime(Basic Version): %s", result, elapsedTime)

	// Write the formatted string to the response
	fmt.Fprint(w, responseString)
}

func main() {
	//var name string
	fmt.Println("Hello World")
	//fmt.Scan(&name)

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Attach handleRequest function to the root endpoint
	mux.HandleFunc("/", handleRequest)

	// Attach handlePrimeRequest function to the /prime/ endpoint
	mux.HandleFunc("/prime/", handlePrimeRequest)

	// Attach handlePrimeRequest function to the /prime/ endpoint
	mux.HandleFunc("/Time10000thPrime/", handleTime10000thPrimeRequest)

	// Start the HTTP server on port 3333
	err := http.ListenAndServe(":3333", mux)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

/* 1st version
func getNthPrime(n int) int {
	if n <= 0 {
		return 0
	}

	count := 0
	num := 2
	for {
		if isPrime(num) {
			count++
			if count == n {
				return num
			}
		}
		num++
	}
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	maxDiv := int(math.Sqrt(float64(num)))
	for i := 2; i <= maxDiv; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}*/

func getNthPrime(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 2
	}

	// Using Sieve of Atkin to generate primes with a larger multiplier
	limit := n * int(math.Log(float64(n))) * 10
	primes := sieveOfAtkinParallel(limit)

	return primes[n-1]
}

func sieveOfAtkinParallel(limit int) []int {
	isPrime := make([]bool, limit+1)
	sqrtLimit := int(math.Sqrt(float64(limit)))

	// Calculate the number of CPU cores
	numCPU := runtime.NumCPU()

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Use a Mutex to synchronize access to the isPrime array
	var mutex sync.Mutex

	// Calculate the range each goroutine should handle
	rangeSize := sqrtLimit / numCPU

	for i := 0; i < numCPU; i++ {
		start := i * rangeSize
		end := (i + 1) * rangeSize
		if i == numCPU-1 {
			end = sqrtLimit
		}

		// Launch a goroutine for each segment
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			markPrimesInRange(start, end, sqrtLimit, limit, &isPrime, &mutex)
		}(start, end)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Collect primes
	primes := make([]int, 0)
	if limit >= 2 {
		primes = append(primes, 2)
	}
	if limit >= 3 {
		primes = append(primes, 3)
	}

	for i := 5; i <= limit; i++ {
		mutex.Lock()
		if isPrime[i] {
			primes = append(primes, i)
		}
		mutex.Unlock()
	}

	return primes
}

func markPrimesInRange(start, end, sqrtLimit, limit int, isPrime *[]bool, mutex *sync.Mutex) {
	for x := 1; x <= sqrtLimit; x++ {
		for y := start; y <= end; y++ {
			n := (4 * x * x) + (y * y)
			if n <= limit && (n%12 == 1 || n%12 == 5) {
				mutex.Lock()
				(*isPrime)[n] = !(*isPrime)[n]
				mutex.Unlock()
			}

			n = (3 * x * x) + (y * y)
			if n <= limit && n%12 == 7 {
				mutex.Lock()
				(*isPrime)[n] = !(*isPrime)[n]
				mutex.Unlock()
			}

			n = (3 * x * x) - (y * y)
			if x > y && n <= limit && n%12 == 11 {
				mutex.Lock()
				(*isPrime)[n] = !(*isPrime)[n]
				mutex.Unlock()
			}
		}
	}
}

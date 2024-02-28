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

	prime := getNthPrime1(x)
	fmt.Fprintf(w, "%d", prime)
}

func handleTime10000thPrimeRequest(w http.ResponseWriter, r *http.Request) {
	//path := r.URL.Path[len("/MeasureTime10000thPrime/"):]

	//with the basic version of algorithm where divides by smaller numbers to check if the number is prime
	start1 := time.Now()
	result1 := getNthPrime1(10000)
	elapsedTime1 := time.Since(start1)

	responseString1 := fmt.Sprintf("10000th prime: %d\nTime taken to find the 10000th prime(Basic Version): %s", result1, elapsedTime1)

	//first improvement
	start2 := time.Now()
	result2 := getNthPrime2(10000)
	elapsedTime2 := time.Since(start2)

	responseString2 := fmt.Sprintf("10000th prime: %d\nTime taken to find the 10000th prime(Basic Version): %s", result2, elapsedTime2)

	//second improvement
	start3 := time.Now()
	result3 := getNthPrime3(10000)
	elapsedTime3 := time.Since(start3)

	responseString3 := fmt.Sprintf("10000th prime: %d\nTime taken to find the 10000th prime(Basic Version): %s", result3, elapsedTime3)

	//multicore improvement
	start4 := time.Now()
	result4 := getNthPrime4(10000)
	elapsedTime4 := time.Since(start4)

	responseString4 := fmt.Sprintf("10000th prime: %d\nTime taken to find the 10000th prime(Basic Version): %s", result4, elapsedTime4)

	// Write the formatted string to the response
	fmt.Fprint(w, responseString1+" \n"+responseString2+" \n"+responseString3+" \n"+responseString4)
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

func getNthPrime1(n int) int {
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
}

func getNthPrime2(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 2
	}

	// Using Sieve of Atkin to generate primes with a larger multiplier
	limit := n * int(math.Log(float64(n))) * 10
	primes := sieveOfAtkin(limit)

	return primes[n-1]
}

func sieveOfAtkin(limit int) []int {
	isPrime := make([]bool, limit+1)
	sqrtLimit := int(math.Sqrt(float64(limit)))

	// Mark sieve[n] is True if one of the following is True:
	// a) n = (4*x*x)+(y*y) has odd number of solutions, i.e., there exist odd number of distinct pairs (x, y) that satisfy the equation
	// b) n = (3*x*x)+(y*y) has odd number of solutions
	// c) n = (3*x*x)-(y*y) has odd number of solutions and x > y
	for x := 1; x <= sqrtLimit; x++ {
		for y := 1; y <= sqrtLimit; y++ {
			n := (4 * x * x) + (y * y)
			if n <= limit && (n%12 == 1 || n%12 == 5) {
				isPrime[n] = !isPrime[n]
			}

			n = (3 * x * x) + (y * y)
			if n <= limit && n%12 == 7 {
				isPrime[n] = !isPrime[n]
			}

			n = (3 * x * x) - (y * y)
			if x > y && n <= limit && n%12 == 11 {
				isPrime[n] = !isPrime[n]
			}
		}
	}

	// Mark all multiples of squares as non-prime
	for r := 5; r <= sqrtLimit; r++ {
		if isPrime[r] {
			for i := r * r; i <= limit; i += r * r {
				isPrime[i] = false
			}
		}
	}

	// Collect primes
	primes := make([]int, 0)
	if limit >= 2 {
		primes = append(primes, 2)
	}
	if limit >= 3 {
		primes = append(primes, 3)
	}

	for i := 5; i <= limit; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}

func getNthPrime3(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 2
	}

	// Using Sieve of Sundaram to generate primes
	limit := n * int(math.Log(float64(n))) * 10
	primes := sieveOfSundaram(limit)

	return primes[n-1]
}

func sieveOfSundaram(limit int) []int {
	// Adjust the limit for Sundaram
	n := (limit - 1) / 2
	sieve := make([]bool, n+1)

	for i := 1; i <= n; i++ {
		j := i
		for ; i+j+2*i*j <= n; j++ {
			sieve[i+j+2*i*j] = true
		}
	}

	// Collect primes
	primes := make([]int, 0)
	if limit >= 2 {
		primes = append(primes, 2)
	}
	for i := 1; i <= n; i++ {
		if !sieve[i] {
			primes = append(primes, 2*i+1)
		}
	}

	return primes
}

func getNthPrime4(n int) int {
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

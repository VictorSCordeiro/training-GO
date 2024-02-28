package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
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
}

/*
// Improved functionality to find Nth prime number
func isPrimeSieve(num int, primes []int) bool {
	maxDiv := int(math.Sqrt(float64(num)))
	for _, prime := range primes {
		if prime > maxDiv {
			break
		}
		if num%prime == 0 {
			return false
		}
	}
	return true
}

func sieveOfEratosthenes(limit int) []int {
	isPrime := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}

	for i := 2; i <= int(math.Sqrt(float64(limit))); i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}

	primes := make([]int, 0)
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
*/

package pe

import (
	"log"
	"sort"
)

const (
	UINT16ZERO = uint16(0)
)

// Odd prime sieve e.g. where index 0 = 1, index 1 = 3, index 2 = 5 ....
type Primes struct {
	sieve []uint16
	max   int
}

// Checks request is in the sieve's range. Bails if not
func (p Primes) InRange(n int) {
	if n < 2 || n > p.max {
		log.Fatalf("%d is not valid in sieve range of %d to %d\n", n, 2, p.max)
	}
}

// IsPrime returns true if prime, otherwise false
func (p Primes) IsPrime(n int) bool {
	p.InRange(n)

	if n%2 == 0 {
		return false
	}
	index := n / 2
	return p.sieve[index] == uint16(0)
}

// Factors returns a list of prime factors with duplicates e.g. 24 returns []int{2,2,2,3}
func (p Primes) Factors(n int) []int {
	p.InRange(n)

	factors := []int{}
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}

	index := n / 2
	for index > 0 {
		if p.sieve[index] == UINT16ZERO {
			factors = append(factors, 2*index+1)
			break
		}
		factor := int(p.sieve[index])
		factors = append(factors, factor)
		n /= factor
		index = n / 2
	}

	sort.Ints(factors)
	return factors
}

// PrimesTo returns a chan of int which delivers prime numbers up to and including n
func (p Primes) PrimesTo(n int) chan int {
	p.InRange(n)

	intChan := make(chan int)
	go p.PrimeGenerator(n, intChan)
	return intChan
}

// PrimesBelow returns the number of primes below (not including) a number n
func (p Primes) PrimesBelow(n int) int {
	if n <= 2 {
		return 0
	}
	p.InRange(n)

	maxIndex := n / 2
	count := 1
	for i := 1; i < maxIndex; i++ {
		if p.sieve[i] == UINT16ZERO {
			count++
		}
	}
	return count
}

func (p Primes) PrimeGenerator(n int, intChan chan int) {
	intChan <- 2
	for i := 3; i <= n; i += 2 {
		if p.IsPrime(i) {
			intChan <- i
		}
	}
	close(intChan)
}

// NewPrimes returns a Primes struct initialized for max size n
func NewPrimes(n int) Primes {
	max := n
	if n%2 == 1 {
		max = n + 1
	}
	sieveLen := max/2 + 1
	primes := Primes{sieve: make([]uint16, sieveLen), max: max}
	for i := 3; i <= IntSqrt(max); i += 2 {
		index := i / 2
		p := uint16(i)
		if primes.sieve[index] == UINT16ZERO {
			for j := index + i; j < sieveLen; j += i {
				primes.sieve[j] = p
			}
		}
	}
	return primes
}

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
func (p Primes) inRange(n int) {
	if n < 2 || n > p.max {
		log.Fatalf("%d is not valid in sieve range of %d to %d\n", n, 2, p.max)
	}
}

// IsPrime returns true if prime, otherwise false
func (p Primes) IsPrime(n int) bool {
	p.inRange(n)

	if n%2 == 0 {
		return false
	}
	index := n / 2
	return p.sieve[index] == uint16(0)
}

// Factors returns a list of prime factors with duplicates e.g. 24 returns []int{2,2,2,3}
func (p Primes) Factors(n int) []int {
	p.inRange(n)

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
	p.inRange(n)

	intChan := make(chan int)
	go p.PrimeGenerator(n, intChan)
	return intChan
}

// PrimesBelow returns the number of primes below (not including) a number n
func (p Primes) PrimesBelow(n int) int {
	if n <= 2 {
		return 0
	}
	p.inRange(n)

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

// FactorTuples returns a list of prime factors and their counts
// e.g. 24 -> [[2 3] [3 1]]  2**3 * 3**1
func (p Primes) FactorTuples(n int) [][]int {
	factors := p.Factors(n)
	tuples := [][]int{}
	tuples = append(tuples, []int{factors[0], 1})
	for i := 1; i < len(factors); i++ {
		f := factors[i]
		if tuples[len(tuples)-1][0] == f {
			tuples[len(tuples)-1][1]++
		} else {
			tuples = append(tuples, []int{f, 1})
		}
	}
	return tuples
}

func (p Primes) helperDivisors(tuples [][]int, n int, ch chan int) {
	if len(tuples) == 0 {
		ch <- n
		return
	}
	prime, exps := tuples[0][0], tuples[0][1]
	for exp := 0; exp <= exps; exp++ {
		p.helperDivisors(tuples[1:], n*Pow(prime, exp), ch)
	}
}

// Divisors returns sorted divisors for n
// e.g. 12 -> [1 2 3 4 6 12]
func (p Primes) Divisors(n int) []int {
	tuples := p.FactorTuples(n)
	numDivisors := 1
	for _, t := range tuples {
		numDivisors *= t[1] + 1
	}

	divisors := []int{}

	ch := make(chan int)
	go p.helperDivisors(tuples, 1, ch)
	for d := 0; d < numDivisors; d++ {
		divisors = append(divisors, <-ch)
	}
	close(ch)

	sort.Ints(divisors)
	return divisors
}

// Totient returns the euler totient of n
func (p Primes) Totient(n int) int {
	facs := p.Factors(n)
	for idx, fac := range facs {
		if idx == 0 || fac != facs[idx-1] {
			n *= fac - 1
			n /= fac
			continue
		}
	}
	return n
}

// Legendere returns factor tups for factorial
// legendere(10) ->[[2 8] [3 4] [5 2] [7 1]]
func (p Primes) Legendere(n int) [][]int {
	factorTups := [][]int{}
	for p := range p.PrimesTo(n) {
		count := 0
		N := n
		for N > 0 {
			count += N / p
			N /= p
		}
		factorTups = append(factorTups, []int{p, count})
	}
	return factorTups
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

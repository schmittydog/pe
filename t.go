package main

import (
	"fmt"
)

func IsPrime(n int) bool {
  if n == 2 {
    return true
  }
  for i := 3; i*i <= n; i += 2 {
    if n%i == 0 {
      return false
    }
  }
  return true
}

func main() {
	fmt.Println(IsPrime(107))
	fmt.Println('a')
	fmt.Println(rune(98))
}

package pe

// Pow returns n^exp (valid to 2^63-1)
func Pow(n, exp int) int {
	return PowMod(n, exp, 9223372036854775807)
}

// PowMod returns n^exp % mod
func PowMod(n, exp, mod int) int {
	ret := 1
	i := 1
	for i <= exp {
		if i&exp > 0 {
			ret *= n
			ret %= mod
		}
		i *= 2
		n *= n
		n %= mod
	}
	return ret
}

// SMod returns S(x/y) % mod
func SMod(x, y, mod int) int {
	if y*2 > x {
		y = x - y
	}
	num := 1
	den := 1
	for i := 0; i < y; i++ {
		num *= x - i
		num %= mod
		den *= i + 1
		den %= mod
	}
	ans := num * PowMod(den, mod-2, mod)
	return ans % mod
}

//máximo comun divisor //GCD greatest commom divisor
package numerico

func mcd(a, b int) int {
	if b > 0 {
		r := a % b
		a = b
		return mcd(a, r)
	}
	return a
}

package iteration

// Repeat takes one letter and iterate to return it N times
func Repeat(letter string, times int) (r string) {
	for i := 0; i < times; i++ {
		r += letter
	}
	return
}

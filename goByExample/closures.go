package goByExample

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

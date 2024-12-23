package pop


var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	totalPc := 0
	bit := x
	for bit != 0 {
		bit = bit&(bit-1)
		totalPc++
		
	}
	return totalPc
}



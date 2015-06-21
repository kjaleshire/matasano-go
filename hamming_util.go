package matasano

func HammingBitDistance(bytes1, bytes2 []byte) (distance int) {
	l := len(bytes1)
	if l != len(bytes2) {
		panic("Hamming distance needs equal-length strings")
	}

	for i := 0; i < l; i++ {
		resultByte := bytes1[i] ^ bytes2[i]
		var bitIndex uint
		for bitIndex = 0; bitIndex < 7; bitIndex++ {
			if (1<<bitIndex)&resultByte > 0 {
				distance++
			}
		}
	}
	return
}

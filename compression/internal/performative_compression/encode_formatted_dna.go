package performativecompression

var dnaMap = map[rune]uint8{
	'A': 0b00,
	'C': 0b01,
	'G': 0b10,
	'T': 0b11,
}

var refMap = map[rune]uint8{
	'I': 0b000,
	'J': 0b001,
	'O': 0b010,
	'P': 0b011,
}

func encodeFormattedDNA(formatted string) ([]byte, error) {
	var out []byte
	var buffer uint8
	var bitCount uint8

	flush := func() {
		if bitCount > 0 {
			buffer <<= (8 - bitCount)
			out = append(out, buffer)
			buffer = 0
			bitCount = 0
		}
	}

	writeBits := func(value uint8, bits uint8) {
		buffer = (buffer << bits) | value
		bitCount += bits
		if bitCount >= 8 {
			shift := bitCount - 8
			out = append(out, buffer>>shift)
			buffer &= (1<<shift - 1)
			bitCount = shift
		}
	}

	for _, r := range formatted {
		// DNA puro
		if v, ok := dnaMap[r]; ok {
			writeBits(0b00, 2) // modo DNA
			writeBits(v, 2)
			continue
		}

		// Referência
		if v, ok := refMap[r]; ok {
			writeBits(0b01, 2) // modo misto
			writeBits(1, 1)    // classe referência
			writeBits(v, 3)
			continue
		}

		continue
	}

	flush()
	return out, nil
}

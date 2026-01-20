package performativecompression

import (
	"os"

	"github.com/Dom-Garotom/estruturaDeDados/compression/pkg"
)

func Decompress(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	result, err := Decode(data)
	if err != nil {
		return err
	}

	file, _, err := pkg.CreateFile(
		"../data/decompressed/dna-performatic-decompressed",
		pkg.TXT,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(result)
	return err
}


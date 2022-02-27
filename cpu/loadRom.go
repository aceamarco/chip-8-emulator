package cpu

import (
	"os"
)

func (*Chip8) LoadRom(filename string) {

	// Open the file as a stream of binary data
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

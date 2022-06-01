package bufferAlign

import (
	"log"
	"unsafe"
)

const AlignSize = 512

func alignment(block []byte, AlignSize int) int {
	return int(uintptr(unsafe.Pointer(&block[0])) & uintptr(AlignSize-1))
}

func GetAlignedBlock(BlockSize int) []byte {
	block := make([]byte, BlockSize+AlignSize)

	a := alignment(block, AlignSize)
	offset := 0
	if a != 0 {
		offset = AlignSize - a
	}

	block = block[offset : offset+BlockSize]
	if BlockSize != 0 {
		a = alignment(block, AlignSize)
		if a != 0 {
			log.Fatal("Failed to align block")
		}
	}

	return block
}

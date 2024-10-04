package main

import (
	"fmt"
	"strings"
)

type BitSet struct {
	bits []uint64
}

func NewBitSet(size uint64) *BitSet {
	return &BitSet{
		bits: make([]uint64, (size/64)+1),
	}
}

func (bs *BitSet) Set(index uint64) {
	arrayIndex := index / 64
	bitIndex := index % 64
	bs.bits[arrayIndex] |= 1 << bitIndex
}

func (bs *BitSet) Count() int {
	count := 0
	for _, bit := range bs.bits {
		count += bitsOn(bit)
	}
	return count
}

func bitsOn(x uint64) int {
	return int(popcount(x))
}

func popcount(x uint64) uint64 {
	return uint64(strings.Count(fmt.Sprintf("%b", x), "1"))
}

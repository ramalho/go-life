package main

type BitField struct {
	width  int
	height int
	data   []uint64
}

func NewBitField(width, height int) *BitField {
	size := width * height
	dataLen := (size + 63) / 64
	return &BitField{
		width:  width,
		height: height,
		data:   make([]uint64, dataLen),
	}
}

func (bf *BitField) Read(x, y int) bool {
	index := y*bf.width + x
	arrayIndex := index / 64
	bitIndex := uint(index % 64)
	return (bf.data[arrayIndex] & (1 << bitIndex)) != 0
}

func (bf *BitField) Set(x, y int, b bool) {
	index := y*bf.width + x
	arrayIndex := index / 64
	bitIndex := uint(index % 64)

	if b {
		bf.data[arrayIndex] |= (1 << bitIndex)
	} else {
		bf.data[arrayIndex] &= ^(1 << bitIndex)
	}
}

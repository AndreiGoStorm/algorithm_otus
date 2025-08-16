package hw5

const (
	NotA  uint64 = 0xFEFEFEFEFEFEFEFE
	NotH  uint64 = 0x7F7F7F7F7F7F7F7F
	NotAB uint64 = 0xFCFCFCFCFCFCFCFC
	NotGH uint64 = 0x3F3F3F3F3F3F3F3F
)

type Bitboard struct {
	bitboard uint64
	position int
}

func NewBitboard() *Bitboard {
	return &Bitboard{}
}

func (b *Bitboard) SetPoint(point uint) {
	b.bitboard = uint64(1) << point
	for i := 0; i < 64; i++ {
		if (b.bitboard>>i)&0x01 > 0 {
			b.position = i
			return
		}
	}
}

func (b *Bitboard) CountBits() uint {
	count := uint(0)
	num := b.bitboard
	for num > 0 {
		if num&0x01 > 0 {
			count++
		}
		num = num >> 1
	}
	return count
}

func (b *Bitboard) KingMoves() {
	mask := (NotA & b.bitboard) >> 1
	mask |= (NotH & b.bitboard) << 1

	mask |= b.bitboard << 8
	mask |= b.bitboard >> 8

	mask |= (NotA & b.bitboard) >> 9
	mask |= (NotA & b.bitboard) << 7
	mask |= (NotH & b.bitboard) << 9
	mask |= (NotH & b.bitboard) >> 7
	b.bitboard = mask
}

func (b *Bitboard) KnightMoves() {
	mask := NotA & (b.bitboard<<17 | b.bitboard>>15)
	mask |= NotH & (b.bitboard<<15 | b.bitboard>>17)
	mask |= NotAB & (b.bitboard<<10 | b.bitboard>>6)
	mask |= NotGH & (b.bitboard<<6 | b.bitboard>>10)
	b.bitboard = mask
}

func (b *Bitboard) RookMoves() {
	// горизонталь (ряд)
	rank := b.position / 8
	// вертикаль (столбец)
	file := b.position % 8

	mask := uint64(0)
	// идем вверх
	for r := rank + 1; r <= 7; r++ {
		mask |= uint64(1) << (r*8 + file)
	}
	// идем вниз
	for r := rank - 1; r >= 0; r-- {
		mask |= uint64(1) << (r*8 + file)
	}
	// идем вправо
	for f := file + 1; f <= 7; f++ {
		mask |= uint64(1) << (rank*8 + f)
	}
	// идем влево
	for f := file - 1; f >= 0; f-- {
		mask |= uint64(1) << (rank*8 + f)
	}
	b.bitboard = mask
}

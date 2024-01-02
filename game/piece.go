package game

import (
	"fmt"
	"strings"
)

var (
	One Piece = [][]bool{
		{true},
	}

	Two Piece = [][]bool{
		{true, true},
	}
	I3 Piece = [][]bool{
		{true, true, true},
	}
	V3 Piece = [][]bool{
		{true, false},
		{true, true},
	}

	T4 Piece = [][]bool{
		{false, true, false},
		{true, true, true},
	}
	O Piece = [][]bool{
		{true, true},
		{true, true},
	}
	L4 Piece = [][]bool{
		{true, false, false},
		{true, true, true},
	}
	I4 Piece = [][]bool{
		{true, true, true, true},
	}
	Z4 Piece = [][]bool{
		{false, true, true},
		{true, true, false},
	}
	F Piece = [][]bool{
		{false, true, false},
		{true, true, true},
		{true, false, false},
	}
	X Piece = [][]bool{
		{false, true, false},
		{true, true, true},
		{false, true, false},
	}
	P Piece = [][]bool{
		{true, true, false},
		{true, true, true},
	}
	W Piece = [][]bool{
		{false, true, true},
		{true, true, false},
		{true, false, false},
	}
	Z5 Piece = [][]bool{
		{false, true, true},
		{false, true, false},
		{true, true, false},
	}
	Y Piece = [][]bool{
		{false, false, true, false},
		{true, true, true, true},
	}
	L5 Piece = [][]bool{
		{true, false, false, false},
		{true, true, true, true},
	}
	U Piece = [][]bool{
		{true, false, true},
		{true, true, true},
	}
	T5 Piece = [][]bool{
		{false, true, false},
		{false, true, false},
		{true, true, true},
	}
	V5 Piece = [][]bool{
		{true, false, false},
		{true, false, false},
		{true, true, true},
	}
	N Piece = [][]bool{
		{false, true, true, true},
		{true, true, false, false},
	}
	I5 Piece = [][]bool{
		{true, true, true, true, true},
	}
)

var StartingPieces = []Piece{
	One, Two, V3, I3, T4, O, L4, I4, Z4, F, X, P, W, Z5, Y, L5, U, T5, V5, N, I5,
}

type Piece [][]bool

func NewPieceFromSlice(buf [][]bool) (Piece, error) {
	// TODO: check for empty rows/cols
	p := Piece(buf)
	if p.Size() == 0 {
		return nil, fmt.Errorf("cannot have piece of size 0")
	}
	return p, nil
}

func NewPieceFromString(s string) (Piece, error) {
	var buf [][]bool
	rows := strings.Split(s, "|")
	for _, row := range rows {
		bufRow := make([]bool, 0, len(row))
		for _, c := range row {
			var v bool
			if c == 'X' {
				v = true
			} else if c == 'O' {
				v = false
			} else {
				return nil, fmt.Errorf("found unexpected char %c, expected X or O", c)
			}
			bufRow = append(bufRow, v)
		}
		buf = append(buf, bufRow)
	}
	return NewPieceFromSlice(buf)
}

func (p Piece) Rotate90DegreesRight() Piece {
	n := len(p)
	m := len(p[0])
	rotatedBuf := make([][]bool, m)
	for j := 0; j < m; j++ {
		rotatedBuf[j] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			rotatedBuf[j][n-i-1] = p[i][j]
		}
	}
	return rotatedBuf
}

func (p Piece) String() string {
	var b strings.Builder
	for i, row := range p {
		for _, val := range row {
			var c byte
			if val {
				c = 'X'
			} else {
				c = 'O'
			}
			b.WriteByte(c)
		}
		if i != len(p)-1 {
			b.WriteByte('|')
		}
	}
	return b.String()
}

func (p Piece) Size() int {
	pieceSize := 0
	for _, row := range p {
		for _, v := range row {
			if v {
				pieceSize++
			}
		}
	}
	return pieceSize
}

func (p Piece) Equals(o Piece) bool {
	for i := 0; i < 4; i++ {
		if p.String() == o.String() {
			return true
		}
		o = o.Rotate90DegreesRight()
	}
	return false
}

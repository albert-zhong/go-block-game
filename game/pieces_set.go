package game

import "fmt"

type PiecesSet struct {
	pieces map[string]struct{}
}

func NewStartingPiecesSet() *PiecesSet {
	startingPiecesSet := NewPiecesSet()
	for _, startingPiece := range StartingPieces {
		startingPiecesSet.Add(startingPiece)
	}
	return startingPiecesSet
}

func NewPiecesSet() *PiecesSet {
	return &PiecesSet{
		pieces: make(map[string]struct{}),
	}
}

func (s *PiecesSet) Add(p Piece) {
	if !s.Contains(p) {
		s.pieces[p.String()] = struct{}{}
	}
}

func (s *PiecesSet) Contains(p Piece) bool {
	for i := 0; i < 4; i++ {
		_, ok := s.pieces[p.String()]
		if ok {
			return true
		}
		p = p.Rotate90DegreesRight()
	}
	return false
}

func (s *PiecesSet) Slice() ([]Piece, error) {
	pieces := make([]Piece, 0, len(s.pieces))
	for pieceString := range s.pieces {
		pieceSlice, err := NewPieceFromString(pieceString)
		if err != nil {
			return nil, fmt.Errorf("failed to convert piece from string to slice format: %w", err)
		}
		pieces = append(pieces, pieceSlice)
	}
	return pieces, nil
}

func (s *PiecesSet) Remove(p Piece) error {
	for i := 0; i < 4; i++ {
		_, ok := s.pieces[p.String()]
		if ok {
			delete(s.pieces, p.String())
			return nil
		}
		p = p.Rotate90DegreesRight()
	}
	return fmt.Errorf("PiecesSet does not contain piece %s", p.String())
}

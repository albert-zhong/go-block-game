package game

import (
	"fmt"
	"math"
)

const (
	// grid constants
	GridWidth  = 20
	GridHeight = GridWidth

	// score constants
	OneSquareIsLastPiecePlacedBonus = 5
	AllPiecesPlacedBonus            = 15
)

type Move struct {
	player Color
	piece  Piece
	x      int
	y      int
}

type Board struct {
	grid           [][]Color
	playersByTurn  []Color
	history        []Move
	piecesByPlayer map[Color]*PiecesSet
}

func NewBoard() *Board {
	grid := make([][]Color, GridHeight)
	for i := 0; i < GridHeight; i++ {
		grid[i] = make([]Color, GridWidth)
	}
	piecesByPlayer := make(map[Color]*PiecesSet)
	for _, player := range NonEmptyColors {
		piecesByPlayer[player] = NewStartingPiecesSet()
	}
	return &Board{
		grid:           grid,
		piecesByPlayer: piecesByPlayer,
	}
}

func (b *Board) ProcessMove(move Move) error {
	piecesLeft, ok := b.piecesByPlayer[move.player]
	if !ok {
		return fmt.Errorf("player %s does not exist", move.player)
	}
	turn := len(b.history) % len(b.playersByTurn)
	if b.playersByTurn[turn] != move.player {
		return fmt.Errorf("player %s cannot move since it is not their turn, expected %s", move.player, b.playersByTurn[turn])
	}
	if !piecesLeft.Contains(move.piece) {
		return fmt.Errorf("invalid move: player %d does not have piece %s", move.player, move.piece.String())
	}

	if !playerCanPlacePiece(move.player, b.grid, move.piece, move.x, move.y) {
		return fmt.Errorf("invalid move: piece %s rooted at position (%d, %d)", move.piece, move.x, move.y)
	}

	for i := 0; i < len(move.piece); i++ {
		for j := 0; j < len(move.piece[i]); j++ {
			b.grid[move.x+i][move.y+j] = move.player
		}
	}
	piecesLeft.Remove(move.piece)
	b.history = append(b.history, move)
	return nil
}

func playerCanPlacePiece(player Color, grid [][]Color, piece Piece, x int, y int) bool {
	for i := 0; i < len(piece); i++ {
		for j := 0; j < len(piece[i]); j++ {
			if grid[x+i][y+j] != Empty {
				return false
			}
			if hasSameColorAdjacent(player, grid, x+i, y+j) {
				return false
			}
		}
	}
	return true
}

// hasSameColorAdjacent returns true if there exists an adjacent square with the given color at grid[x][y]
func hasSameColorAdjacent(color Color, grid [][]Color, x int, y int) bool {
	if x > 0 && grid[x-1][y] == color {
		return true
	}
	if x < len(grid)-1 && grid[x+1][y] == color {
		return true
	}
	if y > 0 && grid[x][y-1] == color {
		return true
	}
	if y < len(grid[x])-1 && grid[x][y+1] == color {
		return true
	}
	return false
}

func (b *Board) PlayerCanPlacePiece(player Color) (bool, error) {
	piecesSetLeft, ok := b.piecesByPlayer[player]
	if !ok {
		return false, fmt.Errorf("player %s does not exist", player)
	}
	piecesLeft, err := piecesSetLeft.Slice()
	if err != nil {
		return false, fmt.Errorf("failed to get a slice of pieces left: %w", err)
	}
	for _, piece := range piecesLeft {
		for x := 0; x < len(b.grid); x++ {
			for y := 0; y < len(b.grid[x]); y++ {
				if playerCanPlacePiece(player, b.grid, piece, x, y) {
					return true, nil
				}
			}
		}
	}
	return false, nil
}

func (b *Board) BasicScore(player Color) int {
	return 0
}

func (b *Board) AdvancedScore(player Color) (int, error) {
	piecesSetLeft, ok := b.piecesByPlayer[player]
	if !ok {
		return 0, fmt.Errorf("player %s does not exist", player)
	}
	score := 0
	piecesLeft, err := piecesSetLeft.Slice()
	if err != nil {
		return 0, fmt.Errorf("failed to get a slice of pieces left: %w", err)
	}
	for _, piece := range piecesLeft {
		score -= piece.Size()
	}
	if len(piecesLeft) == 0 {
		score += AllPiecesPlacedBonus
	}
	lastPiece, ok, err := b.lastPiecePlaced(player)
	if err != nil {
		return 0, fmt.Errorf("failed to get last piece played: %w", err)
	}
	if ok && lastPiece.Equals(One) {
		score += OneSquareIsLastPiecePlacedBonus
	}
	return score, nil
}

func (b *Board) Winner() (Color, error) {
	winner := Empty
	winnerScore := math.MinInt
	for _, player := range b.playersByTurn {
		playerScore, err := b.AdvancedScore(player)
		if err != nil {
			return Empty, fmt.Errorf("failed to calculate score: %w", err)
		}
		if playerScore > winnerScore {
			winner = player
			winnerScore = playerScore
		}
	}
	return winner, nil
}

// lastPiecePlaced returns the last piece placed by the given player. Returns true if the player has made at least 1
// move, and false otherwise. Returns an error if the player does not exist.
func (b *Board) lastPiecePlaced(player Color) (Piece, bool, error) {
	_, ok := b.piecesByPlayer[player]
	if !ok {
		return nil, false, fmt.Errorf("player %s does not exist", player)
	}
	var move Move
	for i := len(b.history) - 1; i >= 0; i-- {
		move = b.history[i]
		if move.player == player {
			return move.piece, true, nil
		}
	}
	return nil, false, nil
}

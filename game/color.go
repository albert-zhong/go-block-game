package game

type Color int

var ColorToString = map[Color]string{
	Empty:  "Empty",
	Red:    "Red",
	Blue:   "Blue",
	Green:  "Green",
	Yellow: "Yellow",
}

var NonEmptyColors = []Color{
	Red, Blue, Green, Yellow,
}

const (
	Empty Color = iota
	Red
	Blue
	Green
	Yellow
)

func (c Color) String() string {
	if s, ok := ColorToString[c]; ok {
		return s
	}
	return "unknown"
}

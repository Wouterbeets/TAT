package bestelling

import (
	"time"
)

type Type int

const (
	A Type = iota
	APlus
	B
	BPlus
)

type Bestelling struct {
	MaaltijdType Type
	Datum        time.Time
	Comentaar    string
}

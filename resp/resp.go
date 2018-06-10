package resp

import (
	"io"

	"github.com/rakin-ishmam/nfplayer/model"
)

// TeamAdder wraps team add func
type TeamAdder interface {
	TeamAdd(model.Team)
}

// Printer wraps print func
type Printer interface {
	Print(io.Writer)
}

//RepsGenerator handles player to generate player list
type RepsGenerator interface {
	TeamAdder
	Printer
}

package detectors

import (
	trn "github.com/NoelM/therailnetwork"
)

type Detector interface {
	Over(trn.Segment) bool
}

type AxleCounter struct {
	position trn.Position
}

func NewAxleCounter(pos trn.Position) AxleCounter {
	return AxleCounter{
		position: pos,
	}
}

func (a AxleCounter) Over(seg trn.Segment) bool {
	return seg.In(a.position)
}

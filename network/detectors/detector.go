// The Rail Network
// Copyright (C) 2025-2026 NoelM
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
package detectors

import (
	spatial "github.com/NoelM/therailnetwork/spatial"
)

type Detector interface {
	ID() int64

	Under(spatial.Segment) bool

	Signals() []int64
	AddSignal(id int64)
	RemoveSignal(id int64)
}

type AxleCounter struct {
	id       int64
	position spatial.Position
	signals  map[int64]bool
}

func NewAxleCounter(id int64, pos spatial.Position) *AxleCounter {
	return &AxleCounter{
		id:       id,
		position: pos,
		signals:  make(map[int64]bool),
	}
}

func (a AxleCounter) ID() int64 {
	return a.id
}

func (a AxleCounter) Under(seg spatial.Segment) bool {
	return seg.In(a.position)
}

func (a AxleCounter) Signals() []int64 {
	keys := make([]int64, 0, len(a.signals))
	for k := range a.signals {
		keys = append(keys, k)
	}

	return keys
}

func (a *AxleCounter) AddSignal(id int64) {
	a.signals[id] = true
}

func (a *AxleCounter) RemoveSignal(id int64) {
	delete(a.signals, id)
}

type TrackCircuit struct {
	id      int64
	segment spatial.Segment
	signals map[int64]bool
}

func NewTrackCircuit(id int64, seg spatial.Segment) *TrackCircuit {
	return &TrackCircuit{
		id:      id,
		segment: seg,
		signals: make(map[int64]bool),
	}
}

func (t TrackCircuit) ID() int64 {
	return t.id
}

func (t TrackCircuit) Under(seg spatial.Segment) bool {
	return seg.Overlap(t.segment)
}

func (t TrackCircuit) Signals() []int64 {
	keys := make([]int64, 0, len(t.signals))
	for k := range t.signals {
		keys = append(keys, k)
	}

	return keys
}

func (t *TrackCircuit) AddSignal(id int64) {
	t.signals[id] = true
}

func (t *TrackCircuit) RemoveSignal(id int64) {
	delete(t.signals, id)
}

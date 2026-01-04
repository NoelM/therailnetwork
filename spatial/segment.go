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
package spatial

import (
	"fmt"
)

type Segment struct {
	startPos Position
	endPos   Position
	dir      Direction
}

func NewSegment(startPos, endPos Position) (Segment, error) {
	if startPos.SectionID() != endPos.SectionID() {
		return Segment{}, fmt.Errorf("cannot create segment: SectionIDs mismatch between start & end positions")
	}

	if startPos.pk == endPos.pk {
		return Segment{}, fmt.Errorf("cannot create segment: lenght between two position must be non-zero")
	}

	var dir Direction
	if endPos.pk > startPos.pk {
		dir = Ascending
	} else {
		dir = Descending
	}

	return Segment{
		startPos: startPos,
		endPos:   endPos,
		dir:      dir,
	}, nil
}

func (s Segment) ascendingPKs() (Distance, Distance) {
	var start, end = s.startPos.pk, s.endPos.pk
	if s.dir == Descending {
		start, end = end, start
	}
	return start, end
}

func (s Segment) Start() Position {
	return s.startPos
}

func (s Segment) End() Position {
	return s.endPos
}

func (s Segment) Len() Distance {
	return abs(s.startPos.pk - s.endPos.pk)
}

func (s Segment) Direction() Direction {
	return s.dir
}

func (s *Segment) Reverse() *Segment {
	if s.dir == Ascending {
		s.dir = Descending
	} else {
		s.dir = Ascending
	}

	s.startPos, s.endPos = s.endPos, s.startPos
	return s
}

func (s Segment) SameSection(seg Segment) bool {
	return s.startPos.SameSection(seg.startPos)
}

func (s Segment) In(pos Position) bool {
	if !s.startPos.SameSection(pos) {
		return false
	}

	start, end := s.ascendingPKs()
	return start <= pos.pk && pos.pk <= end
}

func (s Segment) Overlap(seg Segment) bool {
	if !s.SameSection(seg) {
		return false
	}

	start, end := s.ascendingPKs()
	compareStart, compareEnd := seg.ascendingPKs()

	// Overlap conditions:
	//
	//    current:         |------|
	//    compared: |--------|
	//    start <= comparedEnd && end >= comparedEnd
	//
	// === OR ===
	//
	//    current:  |-----|
	//    compared:     |------|
	//    start <= comparedStart && end >= comparedStart
	//
	// === OR ===
	//
	//    current:   |-----|
	//    compared:    |-|
	//    start <= comparedStart && end >= comparedEnd
	//
	// === OR ===
	//
	//    current:     |-|
	//    compared:  |-----|
	//    start >= comparedStart && end <= comparedEnd
	//

	return ((start <= compareEnd && end >= compareEnd) ||
		(start <= compareStart && end >= compareStart) ||
		(start <= compareStart && end >= compareEnd) ||
		(start >= compareStart && end <= compareEnd))
}

func (s Segment) Intersect(seg Segment) (Segment, error) {
	if !s.Overlap(seg) {
		return Segment{}, fmt.Errorf("cannot compute intersect: segments do not overlap")
	}

	start, end := s.ascendingPKs()
	compareStart, compareEnd := s.ascendingPKs()

	intersectStart := max(start, compareStart)
	intersectEnd := min(end, compareEnd)

	sectionID := s.startPos.sectionID
	return NewSegment(NewPosition(sectionID, intersectStart), NewPosition(sectionID, intersectEnd))
}

func (s Segment) Union(seg Segment) (Segment, error) {
	if !s.Overlap(seg) {
		return Segment{}, fmt.Errorf("cannot compute union: segments do not overlap")
	}

	start, end := s.ascendingPKs()
	compareStart, compareEnd := s.ascendingPKs()

	unionStart := min(start, compareStart)
	unionEnd := max(end, compareEnd)

	sectionID := s.startPos.sectionID
	return NewSegment(NewPosition(sectionID, unionStart), NewPosition(sectionID, unionEnd))
}

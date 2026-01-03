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
}

func NewSegment(startPos, endPos Position) (Segment, error) {
	if startPos.SectionID() != endPos.SectionID() {
		return Segment{}, fmt.Errorf("cannot create segment, SectionIDs mismatch between start & end positions")
	}

	return Segment{
		startPos: startPos,
		endPos:   endPos,
	}, nil
}

func (s Segment) Start() Position {
	return s.startPos
}

func (s Segment) End() Position {
	return s.endPos
}

func (s Segment) Len() int {
	if s.startPos.PK() < s.endPos.PK() {
		return s.endPos.PK() - s.startPos.PK()
	} else {
		return s.startPos.PK() - s.endPos.PK()
	}
}

func (s Segment) Direction() Direction {
	if s.startPos.PK() < s.endPos.PK() {
		return Ascending
	} else {
		return Descending
	}
}

func (s Segment) In(pos Position) bool {
	if s.startPos.EqualSection(pos) && s.endPos.EqualSection(pos) {
		return s.startPos.PK() <= pos.PK() && pos.PK() <= s.endPos.PK()
	}
	return false
}

func (s Segment) Overlaps(seg Segment) bool {
	return seg.In(s.startPos) || seg.In(s.endPos)
}

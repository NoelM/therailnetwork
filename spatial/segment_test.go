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
	"testing"
)

func createSegment(t *testing.T, start, end Position) Segment {
	s, err := NewSegment(start, end)
	if err != nil {
		t.Errorf("unable to create segment: %s", err.Error())
	}
	return s
}

func TestSameSection(t *testing.T) {
	start := NewPosition(0, 0)
	endSuccess := NewPosition(0, 1)
	if _, err := NewSegment(start, endSuccess); err != nil {
		t.Errorf("extremities are on the same SectionID, expected to success")
	}

	endFail := NewPosition(1, 0)
	if _, err := NewSegment(start, endFail); err == nil {
		t.Errorf("extremitites are not on the same SectionID, expected to fail")
	}
}

func TestLen(t *testing.T) {
	start := NewPosition(0, 0)
	end := NewPosition(0, 1)

	ascSegment := createSegment(t, start, end)
	if ascSegment.Len() != 1 {
		t.Errorf("invalid len=%d expected 1", ascSegment.Len().Millimeter())
	}

	dscSegment := createSegment(t, end, start)
	if dscSegment.Len() != 1 {
		t.Errorf("invalid len=%d expected 1", dscSegment.Len().Millimeter())
	}
}

func TestDirection(t *testing.T) {
	start := NewPosition(0, 0)
	end := NewPosition(0, 1)

	ascSegment := createSegment(t, start, end)
	if ascSegment.Direction() != Ascending {
		t.Errorf("invalid dir=%d expected Ascending(2)", ascSegment.Direction())
	}

	dscSegment := createSegment(t, end, start)
	if dscSegment.Direction() != Descending {
		t.Errorf("invalid dir=%d expected Descending(4)", dscSegment.Direction())
	}
}

func TestIn(t *testing.T) {
	start := NewPosition(0, 0)
	end := NewPosition(0, 10)

	ascSegment := createSegment(t, start, end)
	dscSegment := createSegment(t, end, start)

	if !ascSegment.In(start) || !dscSegment.In(start) {
		t.Errorf("position IN segment, found OUT")
	}

	if !ascSegment.In(end) || !dscSegment.In(end) {
		t.Errorf("position IN segment, found OUT")
	}

	inPosition := NewPosition(0, 5)
	if !ascSegment.In(inPosition) || !dscSegment.In(inPosition) {
		t.Errorf("position IN segment, found OUT")
	}

	outPosition := NewPosition(0, 11)
	if ascSegment.In(outPosition) || dscSegment.In(outPosition) {
		t.Errorf("position OUT segment, found IN")
	}

	otherSectionPosition := NewPosition(1, 5)
	if ascSegment.In(otherSectionPosition) || dscSegment.In(otherSectionPosition) {
		t.Errorf("position OUT segment, found IN")
	}
}

func TestOverlap(t *testing.T) {
	segment := createSegment(t, NewPosition(0, 5), NewPosition(0, 10))

	segmentOverlapRight := createSegment(t, NewPosition(0, 10), NewPosition(0, 20))
	if !segment.Overlap(segmentOverlapRight) {
		t.Errorf("segments overlap, found they do not")
	}

	segmentOverlapLeft := createSegment(t, NewPosition(0, 0), NewPosition(0, 5))
	if !segment.Overlap(segmentOverlapLeft) {
		t.Errorf("segments overlap, found they do not")
	}

	segmentOverlapCenter := createSegment(t, NewPosition(0, 7), NewPosition(0, 8))
	if !segment.Overlap(segmentOverlapCenter) {
		t.Errorf("segments overlap, found they do not")
	}

	segmentOverlapLarger := createSegment(t, NewPosition(0, 0), NewPosition(0, 20))
	if !segment.Overlap(segmentOverlapLarger) {
		t.Errorf("segments overlap, found they do not")
	}

	segmentOutOfRange := createSegment(t, NewPosition(0, 100), NewPosition(0, 200))
	if segment.Overlap(segmentOutOfRange) {
		t.Errorf("segments do not overlap, found they do")
	}
}

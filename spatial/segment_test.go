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

	ascSegment, err := NewSegment(start, end)
	if err != nil {
		t.Errorf("must not fail to create segment: %s", err.Error())
	}

	if ascSegment.Len() != 1 {
		t.Errorf("invalid len=%d expected 1", ascSegment.Len())
	}

	dscSegment, err := NewSegment(end, start)
	if err != nil {
		t.Errorf("must not fail to create segment: %s", err.Error())
	}

	if dscSegment.Len() != 1 {
		t.Errorf("invalid len=%d expected 1", dscSegment.Len())
	}
}

func TestDirection(t *testing.T) {
	start := NewPosition(0, 0)
	end := NewPosition(0, 1)

	ascSegment, err := NewSegment(start, end)
	if err != nil {
		t.Errorf("must not fail to create segment: %s", err.Error())
	}

	if ascSegment.Direction() != Ascending {
		t.Errorf("invalid dir=%d expected Ascending(2)", ascSegment.Direction())
	}

	dscSegment, err := NewSegment(end, start)
	if err != nil {
		t.Errorf("must not fail to create segment: %s", err.Error())
	}

	if dscSegment.Direction() != Descending {
		t.Errorf("invalid dir=%d expected Descending(4)", dscSegment.Direction())
	}
}

func TestIn(t *testing.T) {
	start := NewPosition(0, 0)
	end := NewPosition(0, 10)

	segment, err := NewSegment(start, end)
	if err != nil {
		t.Errorf("must not fail to create segment: %s", err.Error())
	}

	if !segment.In(start) {
		t.Errorf("position IN segment, found OUT")
	}

	if !segment.In(end) {
		t.Errorf("position IN segment, found OUT")
	}

	inPosition := NewPosition(0, 5)
	if !segment.In(inPosition) {
		t.Errorf("position IN segment, found OUT")
	}

	outPosition := NewPosition(0, 11)
	if segment.In(outPosition) {
		t.Errorf("position OUT segment, found IN")
	}

	otherSectionPosition := NewPosition(1, 5)
	if segment.In(otherSectionPosition) {
		t.Errorf("position OUT segment, found IN")
	}
}

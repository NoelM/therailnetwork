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

func TestPositionSameSection(t *testing.T) {
	reference := NewPosition(0, 1*Meter)

	sameSection := NewPosition(0, 2*Meter)
	if !reference.SameSection(sameSection) {
		t.Errorf("reference & sameSection positions are on the same SectionID")
	}

	diffSection := NewPosition(1, 1*Meter)
	if reference.SameSection(diffSection) {
		t.Errorf("reference & diffSection positions are not on the same SectionID")
	}
}

func TestPositionDistance(t *testing.T) {
	reference := NewPosition(0, 1*Meter)

	ascPos := NewPosition(0, 2*Meter)
	if dist, err := reference.Distance(ascPos); dist != 1*Meter || err != nil {
		t.Errorf("distance must be 1, got %d, error must be nil, got %s", dist.Millimeter(), err)
	}

	dscPos := NewPosition(0, 0)
	if dist, err := reference.Distance(dscPos); dist != 1*Meter || err != nil {
		t.Errorf("distance must be 1, got %d, error must be nil, got %s", dist.Millimeter(), err)
	}
}

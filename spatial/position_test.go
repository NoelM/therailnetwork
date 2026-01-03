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

func TestEqualSection(t *testing.T) {
	reference := NewPosition(0, 1)
	sameSection := NewPosition(0, 2)
	if !reference.EqualSection(sameSection) {
		t.Errorf("reference & sameSection positions are on the same SectionID")
	}

	diffSection := NewPosition(1, 1)
	if reference.EqualSection(diffSection) {
		t.Errorf("reference & diffSection positions are not on the same SectionID")
	}
}

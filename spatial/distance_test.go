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

func TestUnitConversions(t *testing.T) {
	ref := 1 * Kilometer

	if ref.Millimeter() != 1_000_000 {
		t.Errorf("conversion error to Milimeter, got=%d mm instead of 1_000_000 mm", ref.Millimeter())
	}

	if ref.Meter() != 1_000 {
		t.Errorf("conversion error to Meter, got=%d m instead of 1_000 m", ref.Meter())
	}

	if ref.Kilometer() != 1 {
		t.Errorf("conversion error to Kilometer, got=%d km instead of 1 km", ref.Kilometer())
	}
}

func TestArithmetic(t *testing.T) {
	sum := 1*Millimeter + 1*Meter
	if sum.Millimeter() != 1_001 {
		t.Errorf("summation error, got=%d mm instead of 1_001 mm", sum.Millimeter())
	}

	diff := 1*Meter - 1*Millimeter
	if diff.Millimeter() != 999 {
		t.Errorf("difference error, got=%d mm instead of 999 mm", diff.Millimeter())
	}
}

func TestFunctions(t *testing.T) {
	absDist := abs(1*Millimeter - 1*Meter)
	if absDist.Millimeter() != 999 {
		t.Errorf("absolute value error, got=%d m instead of 999 mm", absDist.Millimeter())
	}

	minDist := min(Millimeter, Meter)
	if minDist != Millimeter {
		t.Errorf("mininum error, got=%d mm instead of 1 mm", minDist.Millimeter())
	}

	maxDist := max(Millimeter, Meter)
	if maxDist != Meter {
		t.Errorf("maxinum error, got=%d mm instead of 1000 mm", minDist.Millimeter())
	}

}

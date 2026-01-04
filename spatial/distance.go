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

type Distance int64

const (
	Millimeter Distance = 1
	Meter      Distance = 1000 * Millimeter
	Kilometer  Distance = 1000 * Meter
)

func (d Distance) Millimeter() int64 {
	return int64(d)
}

func (d Distance) Meter() int64 {
	return int64(d) / 1e3
}

func (d Distance) Kilometer() int64 {
	return int64(d) / 1e6
}

func abs(d Distance) Distance {
	if int64(d) < 0 {
		return -d
	}
	return d
}

func min(d1, d2 Distance) Distance {
	if d1 < d2 {
		return d1
	}
	return d2
}

func max(d1, d2 Distance) Distance {
	if d1 > d2 {
		return d1
	}
	return d2
}

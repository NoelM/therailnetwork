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

// Segment stores multiple positions,
// each position represent a limit
type Interval struct {
	segments []Segment
}

func NewInterval(segments ...Segment) Interval {
	return Interval{
		segments: segments,
	}
}

func (i *Interval) Append(segment Segment) *Interval {
	i.segments = append(i.segments, segment)
	return i
}

func (i Interval) Segments() []Segment {
	return i.segments
}

func (i Interval) Steps() int {
	return len(i.segments)
}

func (i Interval) Len() (length int) {
	for _, s := range i.segments {
		length += s.Len()
	}
	return
}

func (i Interval) In(pos Position) bool {
	for _, s := range i.segments {
		if s.In(pos) {
			return true
		}
	}
	return false
}

func (s *Interval) Overlaps(seg Segment) bool {
	for _, s := range s.segments {
		if s.Overlaps(seg) {
			return true
		}
	}
	return false
}

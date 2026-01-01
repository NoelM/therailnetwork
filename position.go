package therailnetwork

// Position designates a point which belongs
// to a track section at a given kilometer point (PK)
type Position struct {
	sectionID int
	pk        int
}

func NewPosition(secID, pk int) Position {
	return Position{
		sectionID: secID,
		pk:        pk,
	}
}

func (p Position) SectionID() int {
	return p.sectionID
}

func (p Position) PK() int {
	return p.pk
}

func (p Position) EqualSection(pos Position) bool {
	return p.sectionID == pos.SectionID()
}

// Segment stores multiple positions,
// each position represent a limit
type Segment struct {
	positions []Position
}

func NewSegment(positions []Position) Segment {
	return Segment{
		positions: positions,
	}
}

func (s Segment) Positions() []Position {
	return s.positions
}

func (s Segment) Steps() int {
	return len(s.positions)
}

func (s Segment) Lenght() int {
	return 0
}

func (s Segment) In(pos Position) bool {
	var start, stop Position

	for i := 0; i < len(s.positions)-2; i += 1 {
		start = s.positions[i]
		stop = s.positions[i+1]

		if start.EqualSection(pos) && stop.EqualSection(pos) {
			return start.PK() > pos.PK() && pos.PK() < stop.PK()
		}
	}

	return false
}

// We can have 4 cases here:
// s:     |-----|
// seg:      |----|
//
// s:     |-----|
// seg: |----|
//
// s:     |-----|
// seg:     |-|
//
// s:       |-|
// seg:   |-----|
func (s *Segment) Overlaps(seg Segment) bool {

}

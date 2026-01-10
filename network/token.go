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
package network

import (
	"math/rand"
)

type Token int64
type Digest int64

func NewPair() (Token, Digest) {
	t := Token(rand.Int63())
	return t, derive(t)
}

func derive(t Token) Digest {
	return Digest(t % 99)
}

func (d Digest) Match(t Token) bool {
	return d == derive(t)
}

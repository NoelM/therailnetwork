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
package signals

import (
	"sync"

	"github.com/NoelM/therailnetwork/network"
	"github.com/NoelM/therailnetwork/spatial"
)

const (
	BasicRed Status = BasicOffset + iota
	BasicGreen
)

type Basic struct {
	id       int64
	mtx      sync.RWMutex
	position spatial.Position
	trigger  Trigger
	status   Status
	digest   *network.Digest
}

func NewBasic(id int64, pos spatial.Position, trigger Trigger) *Basic {
	return &Basic{
		id:       id,
		position: pos,
		trigger:  trigger,
	}
}

func (b *Basic) ID() int64 {
	b.mtx.RLock()
	defer b.mtx.RUnlock()

	return b.id
}

func (b *Basic) Trigger() Trigger {
	b.mtx.RLock()
	defer b.mtx.RUnlock()

	return b.trigger
}

func (b *Basic) Status() Status {
	b.mtx.RLock()
	defer b.mtx.RUnlock()

	return b.status
}

func (b *Basic) Reserve() *network.Token {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	if b.digest != nil {
		return nil
	}

	token, digest := network.NewPair()
	b.digest = &digest
	return &token
}

func (b *Basic) Release(t network.Token) bool {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	if !b.digest.Match(t) {
		return false
	}

	b.digest = nil
	return true
}

func (b *Basic) Open(t network.Token) bool {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	if !b.digest.Match(t) {
		return false
	}

	b.status = BasicGreen
	return true
}

func (b *Basic) Close(t network.Token) bool {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	if !b.digest.Match(t) {
		return false
	}

	b.status = BasicRed
	return true
}

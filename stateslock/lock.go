package stateslock

import (
	"sync"
)

type Lock struct {
	lk *sync.Mutex

	locked bool
}

func New() *Lock {
	return &Lock{
		lk:     &sync.Mutex{},
		locked: false,
	}
}

func (lk *Lock) Lock() {
	lk.lk.Lock()
	lk.locked = true
}

func (lk *Lock) Unlock() {
	lk.lk.Unlock()
	lk.locked = false
}

func (lk *Lock) Locked() bool {
	return lk.locked
}

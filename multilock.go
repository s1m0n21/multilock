package multilock

import (
	"sync"

	"github.com/s1m0n21/multilock/stateslock"
)

type MultiLock struct {
	lk sync.RWMutex

	locks map[string]*stateslock.Lock
}

func New(cap uint) *MultiLock {
	return &MultiLock{
		locks: make(map[string]*stateslock.Lock, cap),
	}
}

func (ml *MultiLock) Add(key string) error {
	ml.lk.Lock()
	defer ml.lk.Unlock()

	if _, exist := ml.locks[key]; exist {
		return ErrLockAlreadyExist
	}

	lk := stateslock.New()
	ml.locks[key] = lk

	return nil
}

func (ml *MultiLock) Lock(key string) error {
	ml.lk.RLock()
	defer ml.lk.RUnlock()

	if _, exist := ml.locks[key]; !exist {
		ml.lk.RUnlock()
		_ = ml.Add(key)
		ml.lk.RLock()
	}

	lk := ml.locks[key]
	lk.Lock()

	return nil
}

func (ml *MultiLock) Unlock(key string) error {
	ml.lk.RLock()
	defer ml.lk.RUnlock()

	if _, exist := ml.locks[key]; !exist {
		return ErrLockNotExist
	}

	ml.locks[key].Unlock()

	return nil
}

func (ml *MultiLock) Remove(key string) error {
	ml.lk.Lock()
	defer ml.lk.Unlock()

	if _, exist := ml.locks[key]; !exist {
		return ErrLockNotExist
	}

	delete(ml.locks, key)

	return nil
}

func (ml *MultiLock) Count() int {
	ml.lk.RLock()
	defer ml.lk.RUnlock()

	return len(ml.locks)
}

func (ml *MultiLock) Locked(key string) bool {
	ml.lk.RLock()
	defer ml.lk.RUnlock()

	return ml.locks[key].Locked()
}

func (ml *MultiLock) List() map[string]bool {
	ml.lk.RLock()
	defer ml.lk.RUnlock()

	var out = make(map[string]bool, len(ml.locks))
	for k, lk := range ml.locks {
		out[k] = lk.Locked()
	}

	return out
}

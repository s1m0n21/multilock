package multilock

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var key = "test_lock"
var keys = []string{
	"test_a",
	"test_b",
	"test_c",
	"test_d",
}

func TestMultiLock(t *testing.T) {
	ml := New(0)
	t.Log("new multilock")

	err := ml.Lock(key)
	require.NoError(t, err)
	t.Logf("%s locked", key)

	go func() {
		time.Sleep(2 * time.Second)
		err := ml.Unlock(key)
		require.NoError(t, err)
	}()

	err = ml.Lock(key)
	require.NoError(t, err)
	t.Logf("%s locked again", key)

	err = ml.Unlock(key)
	require.NoError(t, err)
	t.Logf("%s unlocked", key)

	for _, k := range keys {
		err = ml.Lock(k)
		require.NoError(t, err)

		t.Logf("key: %s, locked: %v", k, ml.Locked(k))

		err = ml.Unlock(k)
		require.NoError(t, err)

		t.Logf("key: %s, locked: %v", k, ml.Locked(k))
	}

	t.Logf("ml.Count(): %d", ml.Count())
	t.Logf("ml.List(): %+v", ml.List())

	err = ml.Remove(key)
	require.NoError(t, err)

	t.Logf("ml.Count(): %d", ml.Count())
	t.Logf("ml.List(): %+v", ml.List())
}

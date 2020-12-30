package multilock

import (
	"golang.org/x/xerrors"
)

var (
	ErrLockAlreadyExist = xerrors.Errorf("lock already exist")

	ErrLockNotExist = xerrors.Errorf("lock not exist")
)


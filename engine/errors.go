package engine

import "errors"

var (
	errVersionNotExists error = errors.New("the version doesn't exist, use the command: velero-mc engine versions [ENGINE]")
)

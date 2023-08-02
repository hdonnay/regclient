package regclient

import "github.com/regclient/regclient/internal/rwfs"

// NewDirFS returns a new disk-backed filesystem implementation.
func NewDirFS(p string) rwfs.RWFS {
	return rwfs.OSNew(p)
}

// NewMemFS returns a new in-memory filesystem implementation.
func NewMemFS() rwfs.RWFS {
	return rwfs.MemNew()
}
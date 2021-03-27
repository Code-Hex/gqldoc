package pool

import (
	"strings"
	"sync"
)

var global = &sync.Pool{
	New: func() interface{} {
		return &strings.Builder{}
	},
}

func PutBuilder(b *strings.Builder) {
	// Proper usage of a sync.Pool requires each entry to have approximately
	// the same memory cost. To obtain this property when the stored type
	// contains a variably-sized buffer, we add a hard limit on the maximum buffer
	// to place back in the pool.
	//
	// See https://golang.org/issue/23199
	const maxSize = 1 << 16 // 64KiB
	if b.Cap() > maxSize {
		return
	}
	b.Reset()
	global.Put(b)
}

// GetBuilder allocates a new strings.Builder or grabs a cached one.
func GetBuilder() *strings.Builder {
	return global.Get().(*strings.Builder)
}

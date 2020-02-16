// package badidea contains no good ideas.
// Importing package badidea is a bad idea.
package badidea

import "unsafe"

// GoroutineID returns the internal id of the current goroutine.
func GoroutineID() int64 {
	m := (*m)(unsafe.Pointer(runtime_getm()))
	g := m.curg
	return g.goid
}

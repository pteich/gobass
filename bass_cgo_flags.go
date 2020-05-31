package bass
/*
#cgo LDFLAGS: -L${SRCDIR}/lib
#cgo linux,386 LDFLAGS: -lbass_linux_386
#cgo linux,amd64 LDFLAGS: -lbass_linux_amd64
#cgo windows,386 LDFLAGS: -lbass_windows_386
#cgo windows,amd64 LDFLAGS: -lbass_windows_amd64
#cgo darwin,386 LDFLAGS: -lbass_darwin_386
#cgo darwin,amd64 LDFLAGS: -lbass_darwin_amd64
#include "bass.h"
#include "stdlib.h"
#cgo CFLAGS: -I${SRCDIR}/include
*/
import "C"
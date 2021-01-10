// +build !windows

package bass
import(
	"unsafe"
)
/*
#include "bass.h"
*/
import "C"
func Init(device int, freq int, flags int, hwnd, clsid uintptr) error {
	window:=unsafe.Pointer(hwnd)
	gid:=unsafe.Pointer(clsid)
	return boolToError(C.BASS_Init(C.int(device), C.DWORD(freq), C.DWORD(flags), (window), (gid)))
}
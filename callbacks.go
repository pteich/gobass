package bass
import (
	"runtime/cgo"
	"unsafe"
)

/*
#include "bass.h"
#cgo CFLAGS: -Wno-pointer-to-int-cast
extern RECORDPROC _GoBassRecordCallbackStreamPutData; // if our callback doesn't match BASS definition of RECORDPROC, this line will tell us by complaining. Loudly.
extern SYNCPROC _GoSyncprocCallbackWrapper;
*/
import "C"
type GoSyncproc = func(sync Sync, channel Channel, data int)


//export _GoSyncprocCallback
func _GoSyncprocCallback(sync C.HSYNC, channel C.HCHANNEL, data C.DWORD, userdata unsafe.Pointer) {
fn := cgo.Handle(uintptr(userdata)).Value().(GoSyncproc)
	fn(Sync(sync), Channel(channel), int(data))
}
var (
	STREAMPROC_DEVICE = C.STREAMPROC_DEVICE
	STREAMPROC_PUSH = C.STREAMPROC_PUSH
	STREAMPROC_DUMMY = C.STREAMPROC_DUMMY
	RecordCallbackStreamPutData = C._GoBassRecordCallbackStreamPutData
	GoSyncprocCallback = &C._GoSyncprocCallbackWrapper
)
package bass

/*
#include "bass.h"
#cgo CFLAGS: -Wno-pointer-to-int-cast
STREAMPROC* _GO_STREAMPROC_DEVICE = STREAMPROC_DEVICE;
STREAMPROC* _GO_STREAMPROC_PUSH = STREAMPROC_PUSH;
STREAMPROC* _GO_STREAMPROC_DUMMY = STREAMPROC_DUMMY;
CALLBACK BOOL _GoBassRecordCallbackStreamPutData(DWORD recorder, const void* buffer, DWORD length, void* userdata) {
DWORD stream = (DWORD)(userdata);
	BASS_StreamPutData(stream, buffer, length);
	return 1;
}
RECORDPROC _GoBassRecordCallbackStreamPutData; // if our callback doesn't match BASS definition of RECORDPROC, this line will tell us by complaining. Loudly.
*/
import "C"

//BASS defines STREAMPROC_PUSH and friends weirdly using #define and Go can't find them
var (
	STREAMPROC_DEVICE = C._GO_STREAMPROC_DEVICE
	STREAMPROC_PUSH = C._GO_STREAMPROC_PUSH
	STREAMPROC_DUMMY = C._GO_STREAMPROC_DUMMY
	RecordCallbackStreamPutData = (*C.RECORDPROC)(C._GoBassRecordCallbackStreamPutData)
)
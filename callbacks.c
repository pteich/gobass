#include "bass.h"
#include "_cgo_export.h"
#include <stdlib.h>

extern CALLBACK BOOL _GoBassRecordCallbackStreamPutData(DWORD recorder, const void* buffer, DWORD length, void* userdata) {
DWORD stream = (DWORD)(userdata);BASS_StreamPutData(stream, buffer, length);
	return 1;
}

extern CALLBACK void _GoSyncprocCallbackWrapper(HSYNC sync, HCHANNEL channel, DWORD data, void* userdata) {
	_GoSyncprocCallback(sync, channel, data, userdata);
}

// if we just try to get pointers to the above functions from Go, Go will get t the wrong names for these functions as they're __stdcall on 32-bit systems, which means a different naming convension, and Go won't pick up on that. Hello, linker errors. So we need to make wrapper functions which return the right function pointers and do it that way.
RECORDPROC* _get_GoBassRecordCallbackStreamPutData() {
	return _GoBassRecordCallbackStreamPutData;
}
SYNCPROC* _get_GoSyncprocCallbackWrapper() {
	return _GoSyncprocCallbackWrapper;
}

extern CALLBACK void _SyncprocFree(HSYNC sync, HCHANNEL channel, DWORD data, void* userdata) {
	free(userdata);
}

SYNCPROC* _get_SyncprocFree() {
	return _SyncprocFree;
}
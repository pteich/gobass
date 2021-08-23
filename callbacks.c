#include "bass.h"
#include "_cgo_export.h"
extern CALLBACK BOOL _GoBassRecordCallbackStreamPutData(DWORD recorder, const void* buffer, DWORD length, void* userdata) {
DWORD stream = (DWORD)(userdata);BASS_StreamPutData(stream, buffer, length);
	return 1;
}
extern CALLBACK void _GoSyncprocCallbackWrapper(HSYNC sync, HCHANNEL channel, DWORD data, void* userdata) {
	_GoSyncprocCallback(sync, channel, data, userdata);
}
// bass project bass.go
// http://www.un4seen.com/doc/#bass/

package bass

/*
#cgo linux CFLAGS: -I/usr/include -I.
#cgo linux LDFLAGS: -L${SRCDIR} -L/usr/lib -Wl,-rpath=\$ORIGIN -lbass
#cgo windows CFLAGS: -I.
#cgo windows LDFLAGS: -lbass
#include "bass.h"
*/
import "C"

import (
	"fmt"
	"strconv"
	"unsafe"
	"errors"
)
type Channel uint64

var(
	codes=map[int]string {c.BASS_OK: "all is OK", C.BASS_ERROR_MEM:  "memory error", C.BASS_ERROR_FILEOPEN	: "can't open the file", C.BASS_ERROR_DRIVER: "can't find a free/valid driver", c.BASS_ERROR_BUFLOST: "the sample buffer was lost", C.BASS_ERROR_HANDLE: "invalid handle", BASS_ERROR_FORMAT: "unsupported sample format", C.BASS_ERROR_POSITION: "invalid position", C.BASS_ERROR_INIT: "BASS_Init has not been successfully called", C.BASS_ERROR_START: "BASS_Start has not been successfully called", C.BASS_ERROR_SSL: "SSL/HTTPS support isn't available", c.BASS_ERROR_ALREADY: "already initialized/paused/whatever", C.BASS_ERROR_NOTAUDIO: "NOTAUDIO", C.BASS_ERROR_NOCHAN: "can't get a free channel", C.BASS_ERROR_ILLTYPE: "an illegal type was specified", C.BASS_ERROR_ILLPARAM]: "an illegal parameter was specified", C.BASS_ERROR_NO3D: "no 3D support", C.BASS_ERROR_NOEAX: "no EAX support", C.BASS_ERROR_DEVICE: "illegal device number", C.BASS_ERROR_NOPLAY: "not playing", C.BASS_ERROR_FREQ: "illegal sample rate", C.BASS_ERROR_NOTFILE: "the stream is not a file stream", C.BASS_NOHW: "no hardware voices available", C.BASS_ERROR_EMPTY: "the MOD music has no sequence data", C.BASS_ERROR_NONET: "no internet connection could be opened", C.BASS_ERROR_CREATE: "couldn't create the file", C.BASS_ERROR_NOFX: "effects are not available", C.BASS_ERROR_NOTAVAIL: "requested data is not available", C.BASS_ERROR_DECODE: "the channel is/isn't a 'decoding channel", C.BASS_ERROR_DX: "a sufficient DirectX version is not installed", C.BASS_ERROR_TIMEOUT: "connection timedout", C.BASS_ERROR_FILEFORM: "unsupported file format", C.BASS_ERROR_SPEAKER: "unavailable speaker", C.BASS_ERROR_VERSION: "invalid BASS version (used by add-ons)", C.BASS_ERROR_CODEC: "codec is not available/supported", C.BASS_ERROR_ENDED: "the channel/file has ended", C.BASS_ERROR_BUSY: "the device is busy", C.BASS_ERROR_UNSTREAMABLE: "BASS_ERROR_UNSTREAMABLE", C.BASS_ERROR_UNKNOWN: "some other mystery problem"}

)
/*
Init
BOOL BASSDEF(BASS_Init)(int device, DWORD freq, DWORD flags, void *win, void *dsguid);
*/
func Init(device int, freq int, flags int) (bool, error) {
	if C.BASS_Init(C.int(device), C.DWORD(freq), C.DWORD(flags), nil, nil) != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

/*
Free
BOOL BASS_Free();
 */
func Free() (bool, error) {
	if C.BASS_Free() != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

/*
GetConfig
DWORD BASSDEF(BASS_GetConfig)(DWORD option);
*/
func GetConfig(option int) (int, error) {
	v := (int)(C.BASS_GetConfig(C.DWORD(option)))
	if v != -1 {
		return v, nil
	} else {
		return -1, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

/*
SetConfig
BOOL BASSDEF(BASS_SetConfig)(DWORD option, DWORD value);
*/
func SetConfig(option, value int) (bool, error) {
	if C.BASS_SetConfig(C.DWORD(option), C.DWORD(value)) != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// GetVol
// float BASSDEF(BASS_GetVolume)();
func GetVol() (float32, error) {
	if v := C.BASS_GetVolume(); v != -1 {
		return float32(v) * float32(100), nil
	} else {
		return -1, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// SetVol
func SetVol(v float32) (bool, error) {
	if C.BASS_SetVolume(C.float(v/100)) != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// StreamCreateURL
// HSTREAM BASSDEF(BASS_StreamCreateURL)(const char *url, DWORD offset, DWORD flags, DOWNLOADPROC *proc, void *user);
func StreamCreateURL(url string) (Channel, error) {
	ch := C.BASS_StreamCreateURL(C.CString(url), 0, C.BASS_STREAM_BLOCK|C.BASS_STREAM_STATUS|C.BASS_STREAM_AUTOFREE, nil, nil)
	if ch != 0 {
		return int(ch), nil
	} else {
		return 0, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// BASS_StreamCreateFile
// HSTREAM BASSDEF(BASS_StreamCreateFile)(BOOL mem, const void *file, QWORD offset, QWORD length, DWORD flags);
func StreamCreateFile(file string) (Channel, error) {
	ch := C.BASS_StreamCreateFile(0, unsafe.Pointer(C.CString(file)), 0, 0, C.BASS_ASYNCFILE|C.BASS_STREAM_AUTOFREE)
	if ch != 0 {
		return int(ch), nil
	} else {
		return 0, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// ChannelPlay
// BOOL BASSDEF(BASS_ChannelPlay)(DWORD handle, BOOL restart);
func (self Channel) ChannelPlay() (bool, error) {
	if C.BASS_ChannelPlay(self, 1) != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// ChannelPause
// BOOL BASSDEF(BASS_ChannelPause)(DWORD handle);
func (self Channel) ChannelPause() (bool, error) {
	if C.BASS_ChannelPause(self) != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// ChannelStop
// BOOL BASSDEF(BASS_ChannelStop)(DWORD handle);
func (self Channel) ChannelStop() (bool, error) {
	if C.BASS_ChannelStop(self) != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// ChannelStatus
// DWORD BASSDEF(BASS_ChannelIsActive)(DWORD handle);
func (self Channel) ChannelStatus() (c uint, s string) {
	c = uint(C.BASS_ChannelIsActive(self))
	switch c {
	case 0:
		s = "BASS_ACTIVE_STOPPED"
		break
	case 1:
		s = "BASS_ACTIVE_PLAYING"
		break
	case 2:
		s = "BASS_ACTIVE_STALLED"
		break
	case 3:
		s = "BASS_ACTIVE_PAUSED"
		break
	}
	return c, s
}

// ChannelGetAttribute
// BOOL BASSDEF(BASS_ChannelGetAttribute)(DWORD handle, DWORD attrib, float *value);
func (self Channel) ChannelGetAttribute(attrib int) (float32, error) {
	var cvalue C.float
	if C.BASS_ChannelGetAttribute(self, C.DWORD(attrib), &cvalue) != 0 {
		if v, err := strconv.ParseFloat(fmt.Sprintf("%v", cvalue), 32); err != nil {
			return -1, nil
		} else {
			return float32(v), nil
		}
	} else {
		return -1, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// ChannelSetAttribute
// BOOL BASSDEF(BASS_ChannelSetAttribute)(DWORD handle, DWORD attrib, float value);
func (self Channel) ChannelSetAttribute(attrib int, value float32) (bool, error) {
	if C.BASS_ChannelSetAttribute(self, C.DWORD(attrib), C.float(value)) != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

//ChannelGetLevel
//DWORD BASSDEF(BASS_ChannelGetLevel)(DWORD handle);
func (self Channel) ChannelGetLevel(ch int) (c int, e error) {
	c = int(C.BASS_ChannelGetLevel(self)
	if c == -1 {
		return 0, errMsg(int(C.BASS_ErrorGetCode()))
	}
	return c, nil
}

// ChannelGetVolume
// value: 0-100
func ChannelGetVolume(ch int) (float32, error) {
	v, err := ChannelGetAttribute(ch, BASS_ATTRIB_VOL)
	if v > 0 {
		v = v * 100
	}
	return v, err
}

// ChannelSetVolume
// value: 0-100
func ChannelSetVolume(ch int, value float32) (bool, error) {
	return ChannelSetAttribute(ch, BASS_ATTRIB_VOL, value/100)
}

// ChannelGetTags
// const char *BASSDEF(BASS_ChannelGetTags)(DWORD handle, DWORD tags);
func ChannelGetTags(ch int, tag int) string {
	return C.GoString(C.BASS_ChannelGetTags(C.DWORD(ch), C.DWORD(tag)))
}

// PluginLoad
func PluginLoad(file string) (handle int, err error) {
	if h := C.BASS_PluginLoad(C.CString(file), 0); h != 0 {
		return int(h), nil
	} else {
		return 0, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

// PluginFree
func PluginFree(handle int) (bool, error) {
	if C.BASS_PluginFree(C.HPLUGIN(handle)) != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

/* RECORD */
/*
RecordInit
BOOL BASSDEF(BASS_RecordInit)(int device);
*/
func RecordInit(device int) (bool, error) {
	if C.BASS_RecordInit(C.int(device)) != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

/*
RecordFree
BOOL BASSDEF(BASS_RecordFree)();
 */
func RecordFree() (bool, error) {
	if C.BASS_RecordFree() != 0 {
		return true, nil
	} else {
		return false, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

/*
RecordStart - Not working !!!
HRECORD BASSDEF(BASS_RecordStart)(DWORD freq, DWORD chans, DWORD flags, RECORDPROC *proc, void *user);
 */
func RecordStart(freq int, chans int, flags int, proc RecordCallback) (int, error) {
	//h := C.BASS_RecordStart(C.DWORD(freq), C.DWORD(chans), C.DWORD(flags), (*C.RECORDPROC)(unsafe.Pointer(&proc)), nil)
	h := C.BASS_RecordStart(C.DWORD(freq), C.DWORD(chans), C.DWORD(flags), nil, nil)
	if h != 0 {
		return int(h), nil
	} else {
		return 0, errMsg(int(C.BASS_ErrorGetCode()))
	}
}

//typedef BOOL (CALLBACK RECORDPROC)(HRECORD handle, const void *buffer, DWORD length, void *user);
type RecordCallback = func(handle C.HRECORD, buffer *C.char, length C.DWORD, user *C.char) bool

func errMsg(c int) error {

	return errors.New(codes[c])
}

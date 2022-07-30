package enc

import (
	"unsafe"

	"github.com/pteich/gobass"
)

/*
#include "bassenc.h"
#include "stdlib.h"
*/
import "C"

type Encoder bass.Channel

func (e Encoder) SetChannel(channel bass.Channel) error {
	res := C.BASS_Encode_SetChannel(C.DWORD(e), C.DWORD(channel))
	if res == 0 {
		return bass.GetLastError()
	} else {
		return nil
	}
}

func (e Encoder) GetChannel() (bass.Channel, error) {
	ch := C.BASS_Encode_GetChannel(C.DWORD(e))
	var err error
	if ch == 0 {
		err = bass.ErrMsg()
	}

	return bass.Channel(ch), err
}

func (e Encoder) Free() error {
	res := C.BASS_Encode_Stop(C.DWORD(e))
	if res == 0 {
		return bass.GetLastError()
	} else {
		return nil
	}
}

func (e Encoder) ToError() (Encoder, error) {
	if e == 0 {
		return 0, bass.GetLastError()
	} else {
		return e, nil
	}
}

// NewEncoder Sets up an encoder on a channel.
func NewEncoder(channel bass.Channel, cmdline string, flags EncodeStartFlags, callback *C.ENCODEPROC, userdata unsafe.Pointer) (Encoder, error) {
	ccmdline := C.CString(cmdline)
	defer C.free(unsafe.Pointer(ccmdline))

	return Encoder(C.BASS_Encode_Start(C.DWORD(channel), ccmdline, C.DWORD(flags), callback, userdata)).ToError()
}

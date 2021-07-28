package enc
import (
	"github.com/keithcat1/gobass"
)
/*
#include "bassenc.h"
*/
import "C"
type Encoder bass.Channel
// method used mostly by BASS extension bindings. It checks if there was an error by checking if self == 0, and if so returns the BASS error. Otherwise, it returns nil
func (self Encoder) ToError() (Encoder, error) {
	if self == 0 {
		return 0, bass.GetLastError()
	} else {
		return self, nil
	}
}
func (self Encoder) SetChannel(channel bass.Channel) error {
	res := C.BASS_Encode_SetChannel(C.DWORD(self), C.DWORD(channel))
	if res==0 {
		return bass.GetLastError()
	} else {
		return nil
	}
}
func (self Encoder) Free() error {
	res := C.BASS_Encode_Stop(C.DWORD(self))
	if res==0 {
		return bass.GetLastError()
	} else {
		return nil
	}
}
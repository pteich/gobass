package mix
import (
	"github.com/keithcat1/gobass"
	"unsafe"
)

/*
#include "bassmix.h"
*/
import "C"
type Mixer bass.Channel
func MixerStreamCreate(freq, chans, flags int) (Mixer, error) {
	return Mixer(C.BASS_Mixer_StreamCreate(C.DWORD(freq), C.DWORD(chans), C.DWORD(flags))).toError()
}
func (self Mixer) AddChannel(channel bass.Channel, flags int) error {
	return boolToError(C.BASS_Mixer_StreamAddChannel(C.DWORD(self), C.DWORD(channel), C.DWORD(flags)))
}
func (self Mixer) AddChannelEx(channel bass.Channel, flags int, start, length int64) error {
	return boolToError(C.BASS_Mixer_StreamAddChannelEx(C.DWORD(self), C.DWORD(channel), C.DWORD(flags), C.QWORD(start), C.QWORD(length)))
}
func (self Mixer) GetChannels(channels []bass.Channel) (int, error) {
	var ptr *C.DWORD
	if len(channels)>0 {
		ptr = (*C.DWORD)(&channels[0])
	}
	count := int(C.BASS_Mixer_StreamGetChannels(C.DWORD(self), ptr, C.DWORD(len(channels))))
	if count +1 == 0 {
		return 0, bass.GetLastError()
	} else {
		return count, nil
	}
}
func ChannelGetData(channel bass.Channel, buf []byte) (int, error) {
	var ptr unsafe.Pointer
	if len(buf)>0 {
		ptr = unsafe.Pointer(&buf[0])
	}
	count := int(C.BASS_Mixer_ChannelGetData(C.DWORD(channel), ptr, C.DWORD(len(buf))))
	if count +1 == 0 {
		return 0, bass.GetLastError()
	} else {
		return count, nil
	}
}
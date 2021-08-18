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
func (self Mixer) ToChannel() bass.Channel {
	return bass.Channel(self)
}
func StreamCreate(freq, chans int, flags bass.Flags) (Mixer, error) {
	return Mixer(C.BASS_Mixer_StreamCreate(C.DWORD(freq), C.DWORD(chans), C.DWORD(flags))).toError()
}
func (self Mixer) AddChannel(channel bass.Channel, flags bass.Flags) error {
	return boolToError(C.BASS_Mixer_StreamAddChannel(C.DWORD(self), C.DWORD(channel), C.DWORD(flags)))
}
func (self Mixer) AddChannelEx(channel bass.Channel, flags bass.Flags, start, length int64) error {
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
func ChannelRemove(channel bass.Channel) error {
	return boolToError(C.BASS_Mixer_ChannelRemove(C.DWORD(channel)))
}
func ChannelFlags(channel bass.Channel, flags, mask bass.Flags) (bass.Flags, error) {
	flags = bass.Flags(C.BASS_Mixer_ChannelFlags(C.DWORD(channel), C.DWORD(flags), C.DWORD(mask)))
	if flags +1 == 0 {
		return 0, bass.GetLastError()
	} else {
		return flags, nil
	}
}
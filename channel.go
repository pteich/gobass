package bass

/*
#include "bass.h"
#include "stdlib.h"
*/
import "C"
import (
	"unsafe"
)

type Channel uint32

func (ch Channel) cint() cuint {
	return cuint(ch)
}

func (ch Channel) SetSync(syncType int, flags Flags, param int, callback *C.SYNCPROC, userdata unsafe.Pointer) (Sync, error) {
	sync := Sync(C.BASS_ChannelSetSync(ch.cint(), cuint(syncType), culong(param|int(flags)), callback, userdata))
	if sync == 0 {
		return 0, ErrMsg()
	} else {
		return sync, nil
	}
}

func (ch Channel) StreamFree() error {
	return BoolToError(C.BASS_StreamFree(ch.cint()))
}

func (ch Channel) SlideAttribute(attrib uint64, value float64, time uint64) error {
	return BoolToError(C.BASS_ChannelSlideAttribute(ch.cint(), cuint(attrib), C.float(value), cuint(time)))
}

func (ch Channel) SetPosition(pos, mode int) error {
	return BoolToError(C.BASS_ChannelSetPosition(ch.cint(), culong(pos), cuint(mode)))
}

func (ch Channel) GetPosition(mode uint64) (int, error) {
	value := C.BASS_ChannelGetPosition(ch.cint(), C.DWORD(mode))
	if value+1 == 0 {
		return 0, ErrMsg()
	} else {
		return int(value), nil
	}
}

// Play start the channel
func (ch Channel) Play(restart bool) error {
	return BoolToError(C.BASS_ChannelPlay(ch.cint(), boolToInt(restart)))
}

// Pause pauses the channel
func (ch Channel) Pause() error {
	return BoolToError(C.BASS_ChannelPause(ch.cint()))
}

// Stop stops the channel
func (ch Channel) Stop() error {
	return BoolToError(C.BASS_ChannelStop(ch.cint()))
}

// IsActive check if the channel is active
func (ch Channel) IsActive() (int, error) {
	active := int(C.BASS_ChannelIsActive(ch.cint()))
	if active == ACTIVE_STOPPED {
		return active, ErrMsg()
	} else {
		return active, nil
	}
}

// GetAttribute	get channel attribute
func (ch Channel) GetAttribute(attrib int) (float64, error) {
	var cvalue C.float
	result := C.BASS_ChannelGetAttribute(ch.cint(), C.DWORD(attrib), &cvalue)
	return float64(cvalue), BoolToError(result)
}

// SetAttribute set channel attribute
func (ch Channel) SetAttribute(attrib int, value float64) error {
	return BoolToError(C.BASS_ChannelSetAttribute(ch.cint(), C.DWORD(attrib), C.float(value)))
}

// GetLevel get the level of the channel
func (ch Channel) GetLevel() (c int, e error) {
	c = int(C.BASS_ChannelGetLevel(ch.cint()))
	if c == -1 {
		return 0, ErrMsg()
	}
	return c, nil
}

// GetTags
// const char *BASSDEF(BASS_ChannelGetTags)(DWORD handle, DWORD tags);
func (ch Channel) GetTags(tag int) string {
	return C.GoString(C.BASS_ChannelGetTags(ch.cint(), C.DWORD(tag)))
}

// ChannelGetInfo get channel info from the header
func (ch Channel) ChannelGetInfo() (C.BASS_CHANNELINFO, error) {
	var chInfo C.BASS_CHANNELINFO
	if C.BASS_ChannelGetInfo(C.DWORD(ch), &chInfo) != 0 {
		return chInfo, nil
	}
	return chInfo, ErrMsg()
}

func (ch Channel) Bytes2Seconds(bytes int) (float64, error) {
	value := float64(C.BASS_ChannelBytes2Seconds(ch.cint(), C.QWORD(bytes)))
	if value < 0 {
		return value, ErrMsg()
	} else {
		return value, nil
	}
}

func (ch Channel) GetLength(mode uint64) (int, error) {
	result := C.BASS_ChannelGetLength(ch.cint(), C.DWORD(mode))
	if result+1 == 0 {
		return 0, ErrMsg()
	} else {
		return int(result), nil
	}
}

func (ch Channel) IsSliding(attrib uint32) bool {
	return intToBool(C.BASS_ChannelIsSliding(ch.cint(), cuint(attrib)))
}

func (ch Channel) Seconds2Bytes(pos float64) (int, error) {
	val := int(C.BASS_ChannelSeconds2Bytes(ch.cint(), C.double(pos)))
	if val < 0 {
		return 0, ErrMsg()
	} else {
		return val, nil
	}
}

func (ch Channel) Flags(flags, mask Flags) (Flags, error) {
	return Flags(C.BASS_ChannelFlags(ch.cint(), cuint(flags), cuint(mask))), ErrMsg()
}

func (ch Channel) StreamPutData(data []byte, flags Flags) (int, error) {
	var ptr unsafe.Pointer
	if len(data) > 0 {
		ptr = unsafe.Pointer(&data[0])
	}
	val := C.BASS_StreamPutData(ch.cint(), ptr, C.DWORD(len(data)|int(flags)))
	if val+1 == 0 { // -1 indicates an error, but this is an unsigned integer
		return 0, ErrMsg()
	} else {
		return int(val), nil
	}
}

func (ch Channel) GetData(data []byte, flags Flags) (int, error) {
	var ptr unsafe.Pointer
	if len(data) > 0 {
		ptr = unsafe.Pointer(&data[0])
	}
	val := C.BASS_ChannelGetData(ch.cint(), ptr, C.DWORD(len(data)|int(flags)))
	if val+1 == 0 { // -1 indicates an error, but this is an unsigned integer
		return 0, ErrMsg()
	} else {
		return int(val), nil
	}
}

func (ch Channel) Free() error {
	return BoolToError(C.BASS_ChannelFree(C.DWORD(ch)))
}

// ToError method used mostly by BASS extension bindings. It checks if there was an error by checking if self == 0, and if so returns the BASS error. Otherwise, it returns nil
func (ch Channel) ToError() (Channel, error) {
	if ch == 0 {
		return 0, GetLastError()
	} else {
		return ch, nil
	}
}

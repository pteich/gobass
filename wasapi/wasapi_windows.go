package wasapi

import (
	"github.com/pteich/gobass"
	"unsafe"
)

/*
#include "basswasapi.h"
*/
import "C"

func boolToInt(b bool) int {
	switch b {
	case false:
		return 0
	case true:
		return 1
	}
	return 1
}
func Init(device, freq, chans int, flags bass.Flags, buffer, period float64, proc *C.WASAPIPROC, userdata unsafe.Pointer) error {
	return boolToError(C.BASS_WASAPI_Init(C.int(device), C.DWORD(freq), C.DWORD(chans), C.DWORD(flags), C.float(buffer), C.float(period), proc, userdata))
}
func Free() error {
	return boolToError(C.BASS_WASAPI_Free())
}
func Start() error {
	return boolToError(C.BASS_WASAPI_Start())
}
func Stop(reset bool) error {
	return boolToError(C.BASS_WASAPI_Stop(C.BOOL(boolToInt(reset))))
}

type DeviceInfo struct {
	Name, ID                string
	Kind, MixFreq, MixChans int
	Flags                   bass.Flags
	MinPeriod, DefPeriod    float64
}

func GetDeviceInfo(device int) (DeviceInfo, error) {
	var info C.BASS_WASAPI_DEVICEINFO
	err := boolToError(C.BASS_WASAPI_GetDeviceInfo(C.DWORD(device), &info))
	if err != nil {
		return DeviceInfo{}, err
	} else {
		return DeviceInfo{Name: C.GoString(info.name), ID: C.GoString(info.id), Kind: int(info._type), Flags: bass.Flags(info.flags), MinPeriod: float64(info.minperiod), DefPeriod: float64(info.defperiod), MixFreq: int(info.mixfreq), MixChans: int(info.mixchans)}, nil
	}
}
func GetDeviceInfoFlags(device int) (bass.Flags, error) {
	var info C.BASS_WASAPI_DEVICEINFO
	err := boolToError(C.BASS_WASAPI_GetDeviceInfo(C.DWORD(device), &info))
	if err != nil {
		return 0, err
	} else {
		return bass.Flags(info.flags), err
	}
}
func PutData(buf []byte) (int, error) {
	var ptr unsafe.Pointer
	if len(buf) > 0 {
		ptr = unsafe.Pointer(&buf[0])
	}
	count := int(C.BASS_WASAPI_PutData(ptr, C.DWORD(len(buf))))
	if count+1 == 0 {
		return 0, bass.GetLastError()
	} else {
		return count, nil
	}
}

type Info struct {
	InitFlags                                            bass.Flags
	Freq, Chans, Format, BufLen, VolMin, VolMax, VolStep int
}

func GetInfo() (Info, error) {
	var info C.BASS_WASAPI_INFO
	err := boolToError(C.BASS_WASAPI_GetInfo(&info))
	if err != nil {
		return Info{}, err
	} else {
		return Info{InitFlags: bass.Flags(info.initflags), Freq: int(info.freq), Chans: int(info.chans), Format: int(info.format), BufLen: int(info.buflen), VolMin: int(info.volmin), VolMax: int(info.volmax), VolStep: int(info.volstep)}, err
	}
}

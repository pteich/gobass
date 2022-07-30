// bass project bass.go
// http://www.un4seen.com/doc/#bass/

package bass

/*
#include "bass.h"
#include "stdlib.h"
*/
import "C"

import (
	"runtime/cgo"
	"unsafe"
)

// CBytes stores a C byte array, with a pointer to the data and its length.
type CBytes struct {
	Data   unsafe.Pointer
	Length int
}
type cuint = C.DWORD
type culong = C.QWORD

type SyncProc C.SYNCPROC

// Error is the error code returned by any function.
type Error int

func (e Error) Error() string {
	return "BASS error: " + codes[e]
}

type Sample uint32

func (s Sample) cint() cuint {
	return cuint(s)
}

type Handle interface {
	cint() C.DWORD
}

var (
	codes = map[Error]string{OK: "all is OK", ERROR_MEM: "memory error", ERROR_FILEOPEN: "can't open the file", ERROR_DRIVER: "can't find a free/valid driver", ERROR_BUFLOST: "the sample buffer was lost", ERROR_HANDLE: "invalid handle", ERROR_FORMAT: "unsupported sample format", ERROR_POSITION: "invalid position", ERROR_INIT: "BASS_Init has not been successfully called", ERROR_START: "BASS_Start has not been successfully called", ERROR_SSL: "SSL/HTTPS support isn't available", ERROR_ALREADY: "already initialized/paused/whatever", ERROR_NOTAUDIO: "NOTAUDIO", ERROR_NOCHAN: "can't get a free channel", ERROR_ILLTYPE: "an illegal type was specified", ERROR_ILLPARAM: "an illegal parameter was specified", ERROR_NO3D: "no 3D support", ERROR_NOEAX: "no EAX support", ERROR_DEVICE: "illegal device number", ERROR_NOPLAY: "not playing", ERROR_FREQ: "illegal sample rate", ERROR_NOTFILE: "the stream is not a file stream", ERROR_NOHW: "no hardware voices available", ERROR_EMPTY: "the MOD music has no sequence data", ERROR_NONET: "no internet connection could be opened", ERROR_CREATE: "couldn't create the file", ERROR_NOFX: "effects are not available", ERROR_NOTAVAIL: "requested data is not available", ERROR_DECODE: "the channel is/isn't a 'decoding channel", ERROR_DX: "a sufficient DirectX version is not installed", ERROR_TIMEOUT: "connection timedout", ERROR_FILEFORM: "unsupported file format", ERROR_SPEAKER: "unavailable speaker", ERROR_VERSION: "invalid BASS version (used by add-ons)", ERROR_CODEC: "codec is not available/supported", ERROR_ENDED: "the channel/file has ended", ERROR_BUSY: "the device is busy", ERROR_UNSTREAMABLE: "Error_UNSTREAMABLE", ERROR_UNKNOWN: "some other mystery problem"}
)

func Init(device int, freq int, flags Flags, hwnd, clsid uintptr) error {
	window := (bassInitPointer)(unsafe.Pointer(hwnd))
	gid := unsafe.Pointer(clsid)
	return BoolToError(C.BASS_Init(C.int(device), C.DWORD(freq), C.DWORD(flags), window, gid))
}

/*
Free
BOOL BASS_Free();
*/
func Free() error {
	return BoolToError(C.BASS_Free())
}

/*
GetConfig
DWORD BASSDEF(BASS_GetConfig)(DWORD option);
*/
func GetConfig(option int) (int64, error) {
	return longPairToError(C.BASS_GetConfig(C.DWORD(option)))
}

/*
SetConfig
BOOL BASSDEF(BASS_SetConfig)(DWORD option, DWORD value);
*/
func SetConfig(option, value int) error {
	return BoolToError(C.BASS_SetConfig(C.DWORD(option), C.DWORD(value)))
}

// GetVolume get main volume level
// float BASSDEF(BASS_GetVolume)();
func GetVolume() (float64, error) {
	return floatPairToError(C.BASS_GetVolume())
}

// SetVolume sets main volume level
func SetVolume(v float64) error {
	return BoolToError(C.BASS_SetVolume(C.float(v)))
}

// StreamCreateURL Creates a sample stream from an MP3, MP2, MP1, OGG, WAV, AIFF or plugin supported file on the internet, optionally receiving the downloaded data in a callback function.
func StreamCreateURL(url string, flags Flags) (Channel, error) {
	curl := C.CString(url)
	defer C.free(unsafe.Pointer(curl))
	ch := C.BASS_StreamCreateURL(curl, 0, cuint(flags), nil, nil)
	return channelToError(ch)
}

// StreamCreateFile Creates a sample stream from an MP3, MP2, MP1, OGG, WAV, AIFF or plugin supported file.
func StreamCreateFile(data interface{}, offset int, flags Flags) (Channel, error) {
	var ch C.DWORD

	switch data := data.(type) {
	case CBytes:
		ch = C.BASS_StreamCreateFile(1, data.Data, culong(offset), culong(data.Length), cuint(flags))
	case string:
		cstring := unsafe.Pointer(C.CString(data))
		defer C.free(cstring)
		ch = C.BASS_StreamCreateFile(0, cstring, culong(offset), 0, cuint(flags))
	case []byte:
		cbytes := C.CBytes(data)
		ch = C.BASS_StreamCreateFile(1, cbytes, culong(offset), culong(len(data)), cuint(flags))
		// unlike BASS_SampleLoad, BASS won't make a copy of the sample data internally, which means we can't just pass a pointer to the Go bytes. Instead we need to set a sync to free the bytes when the stream it's associated with is freed
		if ch != 0 {
			channel := Channel(ch)
			_, err := channel.SetSync(SYNC_FREE, SYNC_ONETIME, 0, SyncprocFree, cbytes)
			if err != nil {
				return 0, err
			}
		}
	}

	return channelToError(ch)
}

func SampleLoad(data interface{}, offset, max, flags Flags) (Sample, error) {
	var ch C.DWORD
	switch data := data.(type) {
	case CBytes:
		ch = C.BASS_SampleLoad(1, data.Data, culong(offset), cuint(data.Length), cuint(max), cuint(flags))
	case string:
		cstring := unsafe.Pointer(C.CString(data))
		defer C.free(cstring)
		ch = C.BASS_SampleLoad(0, cstring, culong(offset), 0, cuint(max), cuint(flags))
	case []byte:
		// According to BASS documentation, BASS_SampleLoad makes an internal copy of the passed-in memory, so we don't need to worry about CGO restrictions and can just pass a pointer to Go memory
		ch = C.BASS_SampleLoad(1, unsafe.Pointer(&data[0]), culong(offset), cuint(len(data)), cuint(max), cuint(flags))
	}
	return sampleToError(ch)
}

// PluginLoad
func PluginLoad(file string, flags Flags) (handle uint32, err error) {
	cfile := C.CString(file)
	plugin := C.BASS_PluginLoad(cfile, cuint(flags))
	C.free(unsafe.Pointer(cfile))
	return uint32(plugin), ErrMsg()
}

// PluginFree
func PluginFree(handle uint32) error {
	return BoolToError(C.BASS_PluginFree(cuint(handle)))
}

/* RECORD */
/*
RecordInit
BOOL BASSDEF(BASS_RecordInit)(int device);
*/
func RecordInit(device int) error {
	return BoolToError(C.BASS_RecordInit(C.int(device)))
}

/*
RecordFree
BOOL BASSDEF(BASS_RecordFree)();
*/
func RecordFree() error {
	return BoolToError(C.BASS_RecordFree())
}

func ErrMsg() error {
	c := Error(C.BASS_ErrorGetCode())
	if c == 0 {
		return nil
	} else {
		return c
	}
}

// GetLastError returns the last error, if any, that was caused by a call to a BASS function. You should not normally call this function, all BASS functions in this package handle and return errors.
func GetLastError() error {
	return ErrMsg()
}

func (s Sample) Free() error {
	return BoolToError(C.BASS_SampleFree(s.cint()))
}

func (s Sample) GetChannel(flags Flags) (Channel, error) {
	return channelToError(C.BASS_SampleGetChannel(s.cint(), C.DWORD(flags)))
}

func boolToInt(val bool) C.int {
	switch val {
	case true:
		return 1
	case false:
		return 0
	}
	return 0 // It shouldn't get this far.
}

func (s Sample) Stop() error {
	return BoolToError(C.BASS_ChannelStop(s.cint()))
}
func IsStarted() bool {
	return C.BASS_IsStarted() != 0
}

func intToBool(val C.int) bool {
	return val != 0
}

// CopyBytes allocates C memory and coppies data to that C memory.
func CopyBytes(data []byte) CBytes {
	return CBytes{Data: C.CBytes(data), Length: len(data)}
}
func GetDevice() (int64, error) {
	return longPairToError(C.BASS_GetDevice())
}
func SetDevice(device int) error {
	return BoolToError(C.BASS_SetDevice(C.DWORD(device)))
}

func BoolToError(value C.BOOL) error {
	if value == 0 {
		return ErrMsg()
	} else {
		return nil
	}
}

func pairToError(value C.int) (int, error) {
	return int(value), BoolToError(value)
}
func floatPairToError(value C.float) (float64, error) {
	return float64(value), BoolToError(C.int(value))
}
func longPairToError(value C.DWORD) (int64, error) {
	return int64(value), BoolToError(C.int(value))
}
func channelToError(ch C.DWORD) (Channel, error) {
	return Channel(ch), BoolToError(C.int(ch))
}
func sampleToError(ch C.DWORD) (Sample, error) {
	return Sample(ch), BoolToError(C.int(ch))
}

func longlongPairToError(value C.QWORD) (int64, error) {
	return int64(value), BoolToError(C.int(value))
}

func intToError(value cuint) error {
	if value == 0 {
		return ErrMsg()
	} else {
		return nil
	}
}

func intPairToError(value C.DWORD) (int, error) {
	return int(value), intToError(value)
}

func longToError(value culong) error {
	if value != 0 {
		return nil
	} else {
		return ErrMsg()
	}
}

// StreamCreate Creates a user sample stream.
func StreamCreate(freq int, chans int, flags Flags, streamproc *C.STREAMPROC, userdata unsafe.Pointer) (Channel, error) {
	channel := C.BASS_StreamCreate(cuint(freq), cuint(chans), cuint(flags), streamproc, userdata)
	return channelToError(channel)
}

func RecordStart(freq, chans int, flags Flags, streamproc *C.RECORDPROC, userdata unsafe.Pointer) (Channel, error) {
	channel := C.BASS_RecordStart(cuint(freq), cuint(chans), cuint(flags), streamproc, userdata)
	return channelToError(channel)
}

type DeviceInfo struct {
	Name, Driver string
	Flags        Flags
	Kind         uint8
}

func RecordGetDeviceInfo(device int) (DeviceInfo, error) {
	var info C.BASS_DEVICEINFO
	err := BoolToError(C.BASS_RecordGetDeviceInfo(C.DWORD(device), &info))
	if err != nil {
		return DeviceInfo{}, err
	} else {
		return DeviceInfo{Name: C.GoString(info.name), Driver: C.GoString(info.driver), Flags: Flags(info.flags), Kind: getHighWord(info.flags)}, nil
	}
}

func GetDeviceInfo(device int) (DeviceInfo, error) {
	var info C.BASS_DEVICEINFO
	err := BoolToError(C.BASS_GetDeviceInfo(C.DWORD(device), &info))
	if err != nil {
		return DeviceInfo{}, err
	} else {
		return DeviceInfo{Name: C.GoString(info.name), Driver: C.GoString(info.driver), Flags: Flags(info.flags), Kind: getHighWord(info.flags)}, nil
	}
}

func RecordGetDeviceInfoFlags(device int) (Flags, error) {
	var info C.BASS_DEVICEINFO
	err := BoolToError(C.BASS_RecordGetDeviceInfo(C.DWORD(device), &info))
	if err != nil {
		return 0, err
	} else {
		return Flags(info.flags), nil
	}
}

func GetDeviceInfoFlags(device int) (Flags, error) {
	var info C.BASS_DEVICEINFO
	err := BoolToError(C.BASS_GetDeviceInfo(C.DWORD(device), &info))
	if err != nil {
		return 0, err
	} else {
		return Flags(info.flags), nil
	}
}

type RecordInfo struct {
	Flags, Formats, Inputs, Freq, Channels int
	SingleIn                               bool
}

func RecordGetInfo() (RecordInfo, error) {
	var info C.BASS_RECORDINFO
	err := BoolToError(C.BASS_RecordGetInfo(&info))
	if err != nil {
		return RecordInfo{}, err
	} else {
		// For some reason the channels are stored in the last byte of formats
		formats := info.formats >> 8
		ptr := (*uint8)(unsafe.Pointer(&formats))
		return RecordInfo{Flags: int(info.flags), Formats: int(info.formats), Inputs: int(info.inputs), SingleIn: info.singlein != 0, Freq: int(info.freq), Channels: int(*ptr)}, nil
	}
}

func RecordSetDevice(device int) error {
	return BoolToError(C.BASS_RecordSetDevice(C.DWORD(device)))
}

func RecordGetDevice() (int, error) {
	return intPairToError(C.BASS_RecordGetDevice())
}

// Flags is a helper type to allow setting / clearing BASS flags easily
type Flags int64

func (f Flags) Add(flag Flags) Flags {
	return Flags(int(f) | int(flag))
}

func (f Flags) Has(flag Flags) bool {
	return int(f)&int(flag) == int(flag)
}

// AddIf if condition is true, adds the specified flags and returns the result, else just returns self without modifying it
func (f Flags) AddIf(flag Flags, condition bool) Flags {
	if condition {
		return f | flag
	} else {
		return f
	}
}

// getHighWord some BASS functions like to put multiple values into a single integer
func getHighWord(v C.DWORD) uint8 {
	v = v >> 8
	ptr := (*uint8)(unsafe.Pointer(&v))
	return *ptr
}

type Sync Channel

func (ch Channel) RemoveSync(sync Sync) error {
	return BoolToError(C.BASS_ChannelRemoveSync(ch.cint(), C.DWORD(sync)))
}

// NewCGOHandle creates a new runtime/cgo.Handle and converts it to unsafe.Pointer, ready to be pased into C land
func NewCGOHandle(value interface{}) unsafe.Pointer {
	return unsafe.Pointer(cgo.NewHandle(value))
}

// DestroyCGOHandle frees an unneeded handle created by NewCGOHandle
func DestroyCGOHandle(handle unsafe.Pointer) {
	cgo.Handle(handle).Delete()
}

func GetCPU() float64 {
	return float64(C.BASS_GetCPU())
}

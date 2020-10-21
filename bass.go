// bass project bass.go
// http://www.un4seen.com/doc/#bass/

package bass

/*
#include "bass.h"
#include "stdlib.h"
*/
import "C"

import (
	"strconv"
	"unsafe"
)

//Stores a C byte array, with a pointer to the data and its length.
type CBytes struct {
	Data   unsafe.Pointer
	Length int
}
type cuint = C.DWORD
type culong = C.QWORD

type SyncProc C.SYNCPROC

type Channel uint32

// An error returned by any function. Call getErrorCode to get the BASS error code contained inside.
type BASS_Error struct {
	Kind    int
	Message string
}

func (self BASS_Error) Error() string { return self.Message }

// Gets the BASS error code contained in the value. It'll panic if this error value isn't of type BASS_Error.
func ErrorKind(err error) int {
	if err == nil {
		return 0
	}
	return err.(BASS_Error).Kind
}

// Returns true if the passed in error matches kind, which must be a BASS error code
func ErrorIsKind(err error, kind int) bool {
	return ErrorKind(err) == kind
}
func (self Channel) cint() cuint {
	return cuint(self)
}

type Sample uint32

func (self Sample) cint() cuint {
	return cuint(self)
}

type Handle interface {
	cint() C.ulong
}

var (
	codes        = map[int]string{OK: "all is OK", ERROR_MEM: "memory error", ERROR_FILEOPEN: "can't open the file", ERROR_DRIVER: "can't find a free/valid driver", ERROR_BUFLOST: "the sample buffer was lost", ERROR_HANDLE: "invalid handle", ERROR_FORMAT: "unsupported sample format", ERROR_POSITION: "invalid position", ERROR_INIT: "BASS_Init has not been successfully called", ERROR_START: "BASS_Start has not been successfully called", ERROR_SSL: "SSL/HTTPS support isn't available", ERROR_ALREADY: "already initialized/paused/whatever", ERROR_NOTAUDIO: "NOTAUDIO", ERROR_NOCHAN: "can't get a free channel", ERROR_ILLTYPE: "an illegal type was specified", ERROR_ILLPARAM: "an illegal parameter was specified", ERROR_NO3D: "no 3D support", ERROR_NOEAX: "no EAX support", ERROR_DEVICE: "illegal device number", ERROR_NOPLAY: "not playing", ERROR_FREQ: "illegal sample rate", ERROR_NOTFILE: "the stream is not a file stream", ERROR_NOHW: "no hardware voices available", ERROR_EMPTY: "the MOD music has no sequence data", ERROR_NONET: "no internet connection could be opened", ERROR_CREATE: "couldn't create the file", ERROR_NOFX: "effects are not available", ERROR_NOTAVAIL: "requested data is not available", ERROR_DECODE: "the channel is/isn't a 'decoding channel", ERROR_DX: "a sufficient DirectX version is not installed", ERROR_TIMEOUT: "connection timedout", ERROR_FILEFORM: "unsupported file format", ERROR_SPEAKER: "unavailable speaker", ERROR_VERSION: "invalid BASS version (used by add-ons)", ERROR_CODEC: "codec is not available/supported", ERROR_ENDED: "the channel/file has ended", ERROR_BUSY: "the device is busy", ERROR_UNSTREAMABLE: "BASS_ERROR_UNSTREAMABLE", ERROR_UNKNOWN: "some other mystery problem"}
	streamMemory = map[Channel]unsafe.Pointer{} // Here we store the pointers to allocated memory used to store data for a stream. This is only used if you loada  *stream* *from* memory.
)

/*
Free
BOOL BASS_Free();
*/
func Free() error {
	return boolToError(C.BASS_Free())
}

/*
GetConfig
DWORD BASSDEF(BASS_GetConfig)(DWORD option);
*/
func GetConfig(option int) (int64, error) {
	return longPairToError(C.BASS_GetConfig(C.ulong(option)))
}

/*
SetConfig
BOOL BASSDEF(BASS_SetConfig)(DWORD option, DWORD value);
*/
func SetConfig(option, value int) error {
	return boolToError(C.BASS_SetConfig(C.DWORD(option), C.DWORD(value)))
}

// GetVol
// float BASSDEF(BASS_GetVolume)();
func GetVolume() (float32, error) {
	return floatPairToError(C.BASS_GetVolume())
}

// SetVol
func SetVolume(v float32) error {
	return boolToError(C.BASS_SetVolume(C.float(v)))
}

// StreamCreateURL
// HSTREAM BASSDEF(BASS_StreamCreateURL)(const char *url, DWORD offset, DWORD flags, DOWNLOADPROC *proc, void *user);
func StreamCreateURL(url string, flags uint) (Channel, error) {
	curl := C.CString(url)
	defer C.free(unsafe.Pointer(curl))
	ch := C.BASS_StreamCreateURL(curl, 0, cuint(flags), nil, nil)
	return channelToError(ch)
}

func (self *Channel) SetSync(synctype, param uint64, callback *C.SYNCPROC, userdata unsafe.Pointer) (int64, error) {
	return longPairToError(C.BASS_ChannelSetSync(self.cint(), cuint(synctype), culong(param), callback, userdata))
}

// BASS_StreamCreateFile
// HSTREAM BASSDEF(BASS_StreamCreateFile)(BOOL mem, const void *file, QWORD offset, QWORD length, DWORD flags);
func StreamCreateFile(data interface{}, offset, flags int) (Channel, error) {
	var ch C.ulong
	switch data.(type) {
	case CBytes:
		cdata := data.(CBytes)
		ch = C.BASS_StreamCreateFile(1, cdata.Data, culong(offset), culong(cdata.Length), cuint(flags))
	case string:
		datastring := C.CString(data.(string))
		ch = C.BASS_StreamCreateFile(0, unsafe.Pointer(datastring), culong(offset), 0, cuint(flags))
		C.free(unsafe.Pointer(datastring))
	case []byte:
		databytes := C.CBytes(data.([]byte))
		ch = C.BASS_StreamCreateFile(1, databytes, culong(offset), culong(len(data.([]byte))), cuint(flags))
		streamMemory[Channel(ch)] = databytes
	}
	return channelToError(ch)
}

func SampleLoad(data interface{}, offset, max, flags int) (Sample, error) {
	var ch C.ulong
	switch data.(type) {
	case CBytes:
		cdata := data.(CBytes)
		ch = C.BASS_SampleLoad(1, cdata.Data, culong(offset), cuint(cdata.Length), cuint(max), cuint(flags))
	case string:
		datastring := C.CString(data.(string))
		ch = C.BASS_SampleLoad(0, unsafe.Pointer(datastring), culong(offset), 0, cuint(max), cuint(flags))
		C.free(unsafe.Pointer(datastring))
	case []byte:
		databytes := data.([]byte)
		ch = C.BASS_SampleLoad(1, unsafe.Pointer(&databytes[0]), culong(offset), cuint(len(databytes)), cuint(max), cuint(flags))
	}
	return sampleToError(ch)
}

// ChannelPlay
// BOOL BASSDEF(BASS_ChannelPlay)(DWORD handle, BOOL restart);
func (self Channel) Play(restart bool) error {
	return boolToError(C.BASS_ChannelPlay(self.cint(), boolToInt(restart)))
}

// ChannelPause
// BOOL BASSDEF(BASS_ChannelPause)(DWORD handle);
func (self Channel) Pause() error {
	return boolToError(C.BASS_ChannelPause(self.cint()))
}

// ChannelStop
// BOOL BASSDEF(BASS_ChannelStop)(DWORD handle);
func (self Channel) Stop() error {
	return boolToError(C.BASS_ChannelStop(self.cint()))
}

func (self Channel) IsActive() int64 {
	return int64(C.BASS_ChannelIsActive(self.cint()))
}

// ChannelGetAttribute
// BOOL BASSDEF(BASS_ChannelGetAttribute)(DWORD handle, DWORD attrib, float *value);
func (self Channel) GetAttribute(attrib int) (float32, error) {
	var cvalue C.float
	result:=C.BASS_ChannelGetAttribute(self.cint(), C.ulong(attrib), &cvalue)
	return float32(cvalue), boolToError(result)
}
// ChannelSetAttribute
// BOOL BASSDEF(BASS_ChannelSetAttribute)(DWORD handle, DWORD attrib, float value);
func (self Channel) SetAttribute(attrib int, value float32) error {
	return boolToError(C.BASS_ChannelSetAttribute(self.cint(), C.DWORD(attrib), C.float(value)))
}

//ChannelGetLevel
//DWORD BASSDEF(BASS_ChannelGetLevel)(DWORD handle);
func (self Channel) GetLevel() (c int, e error) {
	c = int(C.BASS_ChannelGetLevel(self.cint()))
	if c == -1 {
		return 0, errMsg()
	}
	return c, nil
}

// ChannelGetTags
// const char *BASSDEF(BASS_ChannelGetTags)(DWORD handle, DWORD tags);
func (self Channel) GetTags(tag int) string {
	return C.GoString(C.BASS_ChannelGetTags(self.cint(), C.DWORD(tag)))
}

// PluginLoad
func PluginLoad(file string, flags int) (handle uint32, err error) {
	cfile := C.CString(file)
	plugin := C.BASS_PluginLoad(cfile, cuint(flags))
	C.free(unsafe.Pointer(cfile))
	return uint32(plugin), errMsg()
}

// PluginFree
func PluginFree(handle uint32) error {
	return boolToError(C.BASS_PluginFree(cuint(handle)))
}

/* RECORD */
/*
RecordInit
BOOL BASSDEF(BASS_RecordInit)(int device);
*/
func RecordInit(device int) error {
	return boolToError(C.BASS_RecordInit(C.int(device)))
}

/*
RecordFree
BOOL BASSDEF(BASS_RecordFree)();
*/
func RecordFree() error {
	return boolToError(C.BASS_RecordFree())
}

/*
RecordStart - Not working !!!
HRECORD BASSDEF(BASS_RecordStart)(DWORD freq, DWORD chans, DWORD flags, RECORDPROC *proc, void *user);
*/
func errMsg() error {
	c := int(C.BASS_ErrorGetCode())
	if c == 0 {
		return nil
	}
	return error(BASS_Error{Message: "BASS error: " + strconv.Itoa(c) + ", " + codes[c], Kind: c})
}

func (self Channel) StreamFree() error {
	ptr, exists := streamMemory[self]
	if exists == true {
		C.free(ptr)
	}
	return boolToError(C.BASS_StreamFree(self.cint()))
}
func (self Channel) SlideAttribute(attrib uint64, value float32, time uint64) error {
	return boolToError(C.BASS_ChannelSlideAttribute(self.cint(), cuint(attrib), C.float(value), cuint(time)))
}

func (self Channel) SetPosition(pos, mode int64) error {
	return boolToError(C.BASS_ChannelSetPosition(self.cint(), culong(pos), cuint(mode)))
}

func (self Channel) GetPosition(mode uint64) (int64, error) {
	return longlongPairToError(C.BASS_ChannelGetPosition(self.cint(), C.ulong(mode)))
}

func (self Sample) Free() error {
	return boolToError(C.BASS_SampleFree(self.cint()))
}
func (self Sample) GetChannel(onlynew bool) (Channel, error) {
	return channelToError(C.BASS_SampleGetChannel(self.cint(), boolToInt(onlynew)))
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

func (self Sample) Stop() error {
	return boolToError(C.BASS_ChannelStop(self.cint()))
}
func IsStarted() bool {
	return intToBool(C.BASS_IsStarted())
}

func (self Channel) Bytes2Seconds(bytes int64) (float64, error) {
	value:=float64(C.BASS_ChannelBytes2Seconds(self.cint(), C.QWORD(bytes)))
	if value<0 {
		return value, errMsg()
	} else {
		return value, nil
	}
}
func (self Channel) GetLength(mode uint64) (int64, error) {
	return longlongPairToError(C.BASS_ChannelGetLength(self.cint(), C.DWORD(mode)))
}
func intToBool(val C.int) bool {
	return val != 0
}
func (self Channel) IsSliding(attrib uint32) bool {
	return intToBool(C.BASS_ChannelIsSliding(self.cint(), cuint(attrib)))
}
func (self Channel) Seconds2Bytes(pos float64) (int64, error) {
	val := int64(C.BASS_ChannelSeconds2Bytes(self.cint(), C.double(pos)))
	if val<0 {
		return val, errMsg()
	} else {
		return val, nil
	}
}
func (self Channel) Flags(a, b uint32) (uint32, error) {
	return uint32(C.BASS_ChannelFlags(self.cint(), cuint(a), cuint(b))), errMsg()
}

//Allocates C memory and coppies data to that C memory.
func CopyBytes(data []byte) CBytes {
	return CBytes{Data: C.CBytes(data), Length: len(data)}
}
func GetDevice() (int64, error) {
	return longPairToError(C.BASS_GetDevice())
}
func SetDevice(device int) error {
	return boolToError(C.BASS_SetDevice(C.ulong(device)))
}
func boolToError(value C.int) error {
	if value != 0 {
		return errMsg()
	}
	return nil
}
func pairToError(value C.int) (int, error) {
	return int(value), boolToError(value)
}
func floatPairToError(value C.float) (float32, error) {
	return float32(value), boolToError(C.int(value))
}
func longPairToError(value C.ulong) (int64, error) {
	return int64(value), boolToError(C.int(value))
}
func channelToError(ch C.ulong) (Channel, error) {
	return Channel(ch), boolToError(C.int(ch))
}
func sampleToError(ch C.ulong) (Sample, error) {
	return Sample(ch), boolToError(C.int(ch))
}
func longlongPairToError(value C.ulonglong) (int64, error) {
	return int64(value), boolToError(C.int(value))
}
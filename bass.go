// bass project bass.go
// http://www.un4seen.com/doc/#bass/

package bass


/*
#include "bass.h"
#include "stdlib.h"
*/
import "C"

import (
	"fmt"
	"strconv"
	"unsafe"
)

type cuint=C.DWORD
type culong=C.QWORD


type SyncProc C.SYNCPROC

type Channel uint32
// An error returned by any function. Call getErrorCode to get the BASS error code contained inside.
type BASS_Error struct {
	Kind int
	Message string

}
func (self BASS_Error) Error() string { return self.Message }
// Gets the BASS error code contained in the value. It'll panic if this error value isn't of type BASS_Error.
func ErrorKind(err error) int {
	if err==nil { return 0 }
	return err.(BASS_Error).Kind
}
// Returns true if the passed in error matches kind, which must be a BASS error code
func ErrorIsKind(err error, kind int) bool {
	return ErrorKind(err)==kind
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





var(
	codes=map[int]string {OK: "all is OK", ERROR_MEM:  "memory error", ERROR_FILEOPEN	: "can't open the file", ERROR_DRIVER: "can't find a free/valid driver", ERROR_BUFLOST: "the sample buffer was lost", ERROR_HANDLE: "invalid handle", ERROR_FORMAT: "unsupported sample format", ERROR_POSITION: "invalid position", ERROR_INIT: "BASS_Init has not been successfully called", ERROR_START: "BASS_Start has not been successfully called", ERROR_SSL: "SSL/HTTPS support isn't available", ERROR_ALREADY: "already initialized/paused/whatever", ERROR_NOTAUDIO: "NOTAUDIO", ERROR_NOCHAN: "can't get a free channel", ERROR_ILLTYPE: "an illegal type was specified", ERROR_ILLPARAM: "an illegal parameter was specified", ERROR_NO3D: "no 3D support", ERROR_NOEAX: "no EAX support", ERROR_DEVICE: "illegal device number", ERROR_NOPLAY: "not playing", ERROR_FREQ: "illegal sample rate", ERROR_NOTFILE: "the stream is not a file stream", ERROR_NOHW: "no hardware voices available", ERROR_EMPTY: "the MOD music has no sequence data", ERROR_NONET: "no internet connection could be opened", ERROR_CREATE: "couldn't create the file", ERROR_NOFX: "effects are not available", ERROR_NOTAVAIL: "requested data is not available", ERROR_DECODE: "the channel is/isn't a 'decoding channel", ERROR_DX: "a sufficient DirectX version is not installed", ERROR_TIMEOUT: "connection timedout", ERROR_FILEFORM: "unsupported file format", ERROR_SPEAKER: "unavailable speaker", ERROR_VERSION: "invalid BASS version (used by add-ons)", ERROR_CODEC: "codec is not available/supported", ERROR_ENDED: "the channel/file has ended", ERROR_BUSY: "the device is busy", ERROR_UNSTREAMABLE: "BASS_ERROR_UNSTREAMABLE", ERROR_UNKNOWN: "some other mystery problem"}
	streamMemory=map[Channel]unsafe.Pointer{} // Here we store the pointers to allocated memory used to store data for a stream. This is only used if you loada  *stream* *from* memory.
)


/*
Free
BOOL BASS_Free();
 */
func Free() (bool, error) {
	if C.BASS_Free() != 0 {
		return true, nil
	} else {
		return false, errMsg()
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
		return -1, errMsg()
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
		return false, errMsg()
	}
}

// GetVol
// float BASSDEF(BASS_GetVolume)();
func GetVolume() (float32, error) {
	if v := C.BASS_GetVolume(); v != -1 {
		return float32(v), nil
	} else {
		return -1, errMsg()
	}
}

// SetVol
func SetVolume(v float32) (bool, error) {
	if C.BASS_SetVolume(C.float(v)) != 0 {
		return true, nil
	} else {
		return false, errMsg()
	}
}

// StreamCreateURL
// HSTREAM BASSDEF(BASS_StreamCreateURL)(const char *url, DWORD offset, DWORD flags, DOWNLOADPROC *proc, void *user);
func StreamCreateURL(url string, flags uint) (Channel, error) {
	curl:=C.CString(url)
	ch:=Channel(C.BASS_StreamCreateURL(curl, 0, cuint(flags), nil, nil))
	C.free((unsafe.Pointer)(curl))
	if ch != 0 {
		return ch, nil
	} else {
		return 0, errMsg()
	}
}

func (self *Channel) SetSync(synctype, param uint64, callback *C.SYNCPROC, userdata unsafe.Pointer) (uint64, error) {
	c:=uint64(C.BASS_ChannelSetSync(self.cint(), cuint(synctype), culong(param), callback, userdata))
	if c!=0 { return c, nil}
	return 0, errMsg()
}
// BASS_StreamCreateFile
// HSTREAM BASSDEF(BASS_StreamCreateFile)(BOOL mem, const void *file, QWORD offset, QWORD length, DWORD flags);
func StreamCreateFile(data interface{}, offset, flags int) (Channel, error) {
	var ch Channel
	switch data.(type) {
		case unsafe.Pointer: ch=Channel(C.BASS_StreamCreateFile(1, data.(unsafe.Pointer), culong(offset), 0, cuint(flags)))
		case string:
			datastring:=C.CString(data.(string))
		ch=Channel(C.BASS_StreamCreateFile(0, unsafe.Pointer(datastring), culong(offset), 0, cuint(flags)))
		C.free(unsafe.Pointer(datastring))
		case []byte:
		databytes:=C.CBytes(data.([]byte))
		ch=Channel(C.BASS_StreamCreateFile(1, databytes, culong(offset), culong(len(data.([]byte))), cuint(flags)))
		streamMemory[ch]=databytes
}
	return ch, errMsg()
}


func SampleLoad(data interface{}, offset, max, flags int) (Sample, error) {
	var ch Sample
	switch data.(type) {
		case unsafe.Pointer: ch=Sample(C.BASS_SampleLoad(1, data.(unsafe.Pointer), culong(offset), 0, cuint(max), cuint(flags)))
		case string:
			datastring:=C.CString(data.(string))
		ch=Sample(C.BASS_SampleLoad(0, unsafe.Pointer(datastring), culong(offset), 0, cuint(max), cuint(flags)))
		C.free(unsafe.Pointer(datastring))
		case []byte:
		databytes:=data.([]byte)
		ch=Sample(C.BASS_SampleLoad(1, unsafe.Pointer(&databytes[0]), culong(offset), cuint(len(databytes)), cuint(max), cuint(flags)))
}
	return ch, errMsg()
}




// ChannelPlay
// BOOL BASSDEF(BASS_ChannelPlay)(DWORD handle, BOOL restart);
func (self Channel) Play(restart bool) (bool, error) {
	var should C.int=0
	if restart==true { should=1 }
	if C.BASS_ChannelPlay(self.cint(), should) != 0 {
		return true, nil
	} else {
		return false, errMsg()
	}
}

// ChannelPause
// BOOL BASSDEF(BASS_ChannelPause)(DWORD handle);
func (self Channel) Pause() (bool, error) {
	if C.BASS_ChannelPause(C.DWORD(self)) != 0 {
		return true, nil
	} else {
		return false, errMsg()
	}
}

// ChannelStop
// BOOL BASSDEF(BASS_ChannelStop)(DWORD handle);
func (self Channel) Stop() (bool, error) {
	if C.BASS_ChannelStop(self.cint()) != 0 {
		return true, nil
	} else {
		return false, errMsg()
	}
}


func (self Channel) IsActive() (uint, error) {
	status:=uint(C.BASS_ChannelIsActive(self.cint()))
	if status==ACTIVE_STOPPED {
		err:=errMsg()
		if err!=nil {
			return 0, err
		}
	}
	return status, nil
}


// ChannelGetAttribute
// BOOL BASSDEF(BASS_ChannelGetAttribute)(DWORD handle, DWORD attrib, float *value);
func (self Channel) GetAttribute(attrib int) (float32, error) {
	var cvalue C.float
	if C.BASS_ChannelGetAttribute(self.cint(), C.DWORD(attrib), &cvalue) != 0 {
		if v, err := strconv.ParseFloat(fmt.Sprintf("%v", cvalue), 32); err != nil {
			return -1, nil
		} else {
			return float32(v), nil
		}
	} else {
		return -1, errMsg()
	}
}

// ChannelSetAttribute
// BOOL BASSDEF(BASS_ChannelSetAttribute)(DWORD handle, DWORD attrib, float value);
func (self Channel) SetAttribute(attrib int, value float32) (error) {
	if C.BASS_ChannelSetAttribute(self.cint(), C.DWORD(attrib), C.float(value)) != 0 {
		return nil
	} else {
		return errMsg()
	}
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
	cfile:=C.CString(file)
	plugin:=C.BASS_PluginLoad(cfile, cuint(flags))
	C.free(unsafe.Pointer(cfile))
	return uint32(plugin), errMsg()
}

// PluginFree
func PluginFree(handle uint32) (bool, error) {
	return intToBool(C.BASS_PluginFree(cuint(handle))), errMsg()
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
		return false, errMsg()
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
		return false, errMsg()
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
		return 0, errMsg()
	}
}

//typedef BOOL (CALLBACK RECORDPROC)(HRECORD handle, const void *buffer, DWORD length, void *user);
type RecordCallback func(handle C.HRECORD, buffer *C.char, length C.DWORD, user *C.char) bool

func errMsg() error {
	c:=int(C.BASS_ErrorGetCode())
	if c==0 { return nil }
	return error(BASS_Error{Message: "BASS error: "+strconv.Itoa(c)+", "+codes[c], Kind: c})
}

func (self Channel) StreamFree() error {
	ptr, exists:=streamMemory[self]
	if exists==true { C.free(ptr) }
	if C.BASS_StreamFree(self.cint())!=1 {
		return errMsg()
	}
	return nil
}
func (self Channel) SlideAttribute(attrib uint64, value float32, time uint64) error {
	if C.BASS_ChannelSlideAttribute(self.cint(), cuint(attrib), C.float(value), cuint(time))!=1 {
		return errMsg()
}
	return nil
}

func (self Channel) SetPosition(pos, mode uint64) error {
	if C.BASS_ChannelSetPosition(self.cint(), culong(pos), cuint(mode))!=1 {
		return errMsg()
	}
	return nil
}


func (self Channel) GetPosition(mode uint64) (uint64, error) {
	val:=C.BASS_ChannelGetPosition(self.cint(), C.DWORD(mode))
	if val==0 { return 0, errMsg() }
	return uint64(val), nil
}



func (self Sample) Free() error {
	if C.BASS_SampleFree(self.cint())==0 {
		return errMsg()
	}
	return nil
}
func (self Sample) GetChannel(onlynew bool) (Channel, error) {
	channel:=Channel(C.BASS_SampleGetChannel(self.cint(), boolToInt(onlynew)))
	if channel==0 {
		return channel, errMsg()
	}
	return channel, nil
}
func boolToInt(val bool) C.int {
	switch val {
	case true: return 1
	case false: return 0
	}
	return 0 // It shouldn't get this far.
}


func (self Sample) Stop() error {
	if C.BASS_SampleStop(self.cint())==0 {
		return errMsg()
	}
	return nil
}
func IsStarted() bool {
	if C.BASS_IsStarted()==1 {
		return true
	}
	return false
}

func (self Channel) Bytes2Seconds(bytes uint64) (float64, error) {
	val:=C.BASS_ChannelBytes2Seconds(self.cint(), C.QWORD(bytes))
	if val==0 { return 0, errMsg() }
	return float64(val), nil
}
func (self Channel) GetLength(mode uint64) (uint64, error) {
	val:=C.BASS_ChannelGetLength(self.cint(), C.DWORD(mode))
	if val==0 { return 0, errMsg() }
	return uint64(val), nil
}
func intToBool(val C.int) bool {
	return val!=0
}
func (self Channel) IsSliding(attrib uint32) (bool, error) {
	return intToBool(C.BASS_ChannelIsSliding(self.cint(), cuint(attrib))), errMsg()
}
func (self Channel) Seconds2Bytes(pos float64) (uint64, error) {
	val:=uint64(C.BASS_ChannelSeconds2Bytes(self.cint(), C.double(pos)))
	return val, errMsg()
}
func (self Channel) Flags(a, b uint32) (uint32, error) {
	return uint32(C.BASS_ChannelFlags(self.cint(), cuint(a), cuint(b))), errMsg()
}

//Allocates C memory and coppies data to that C memory.
func CopyBytes(data []byte) unsafe.Pointer {
	return C.CBytes(data)
}
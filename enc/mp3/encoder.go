package mp3
import (
	"unsafe"
	"github.com/keithcat1/gobass"
"github.com/keithcat1/gobass/enc"
)

/*
#include "bassenc_MP3.h"
#include "stdlib.h"
*/
import "C"
func NewMP3EncoderFile(channel bass.Channel, options string, flags int, file string) (enc.Encoder, error) {
	coptions := C.CString(options)
	defer C.free(unsafe.Pointer(coptions))
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))
	return enc.Encoder(C.BASS_Encode_MP3_StartFile(C.DWORD(channel), coptions, C.DWORD(flags), cfile)).ToError()
}
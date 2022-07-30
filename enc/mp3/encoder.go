package mp3

import (
	"unsafe"

	"github.com/pteich/gobass"
	"github.com/pteich/gobass/enc"
)

/*
#include "bassenc_mp3.h"
#include "stdlib.h"
*/
import "C"

func NewEncoder(channel bass.Channel, options string, flags enc.EncodeStartFlags, callback *C.ENCODEPROCEX, userdata unsafe.Pointer) (enc.Encoder, error) {
	coptions := C.CString(options)
	defer C.free(unsafe.Pointer(coptions))

	return enc.Encoder(C.BASS_Encode_MP3_Start(C.DWORD(channel), coptions, C.DWORD(flags), callback, userdata)).ToError()
}

func NewEncoderFile(channel bass.Channel, options string, flags enc.EncodeStartFlags, file string) (enc.Encoder, error) {
	coptions := C.CString(options)
	defer C.free(unsafe.Pointer(coptions))
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))
	return enc.Encoder(C.BASS_Encode_MP3_StartFile(C.DWORD(channel), coptions, C.DWORD(flags), cfile)).ToError()
}

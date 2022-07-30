package ogg

import (
	"unsafe"

	"github.com/pteich/gobass"
	"github.com/pteich/gobass/enc"
)

/*
#include "bassenc_ogg.h"
#include "stdlib.h"
*/
import "C"

func NewEncoder(channel bass.Channel, options string, flags bass.Flags, callback *C.ENCODEPROC, userdata unsafe.Pointer) (enc.Encoder, error) {
	coptions := C.CString(options)
	defer C.free(unsafe.Pointer(coptions))

	return enc.Encoder(C.BASS_Encode_OGG_Start(C.DWORD(channel), coptions, C.DWORD(flags), callback, userdata)).ToError()
}

func NewEncoderFile(channel bass.Channel, options string, flags bass.Flags, file string) (enc.Encoder, error) {
	coptions := C.CString(options)
	defer C.free(unsafe.Pointer(coptions))
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))
	return enc.Encoder(C.BASS_Encode_OGG_StartFile(C.DWORD(channel), coptions, C.DWORD(flags), cfile)).ToError()
}

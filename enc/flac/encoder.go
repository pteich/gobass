package flac

import (
	"unsafe"

	"github.com/pteich/gobass"
	"github.com/pteich/gobass/enc"
)

/*
#include "bassenc_flac.h"
#include "stdlib.h"
*/
import "C"

func NewFLACEncoderFile(channel bass.Channel, options string, flags bass.Flags, file string) (enc.Encoder, error) {
	coptions := C.CString(options)
	defer C.free(unsafe.Pointer(coptions))
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))
	return enc.Encoder(C.BASS_Encode_FLAC_StartFile(C.DWORD(channel), coptions, C.DWORD(flags), cfile)).ToError()
}

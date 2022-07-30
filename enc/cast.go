package enc

/*
#include "bassenc.h"
#include "stdlib.h"
*/
import "C"
import (
	"unsafe"

	bass "github.com/pteich/gobass"
)

type EncodeType string

const (
	EncodeTypeMP3 EncodeType = "audio/mpeg"
	EncodeTypeOGG EncodeType = "audio/ogg"
	EncodeTypeAAC EncodeType = "audio/aacp"
)

type CastConfig struct {
	Server      string
	Password    string
	ContentType EncodeType
	Name        string
	Url         string
	Genre       string
	Description string
	Headers     string
	Bitrate     uint32
}

// CastInit Initializes sending an encoder's output to a Shoutcast or Icecast server.
func CastInit(encoder Encoder, cfg CastConfig, flags EncodeCastInitFlags) error {
	serverC := C.CString(cfg.Server)
	defer C.free(unsafe.Pointer(serverC))
	passwordC := C.CString(cfg.Password)
	defer C.free(unsafe.Pointer(passwordC))
	mimeC := C.CString(string(cfg.ContentType))
	defer C.free(unsafe.Pointer(mimeC))
	nameC := C.CString(cfg.Name)
	defer C.free(unsafe.Pointer(nameC))
	urlC := C.CString(cfg.Url)
	defer C.free(unsafe.Pointer(urlC))
	genreC := C.CString(cfg.Genre)
	defer C.free(unsafe.Pointer(genreC))
	descriptionC := C.CString(cfg.Description)
	defer C.free(unsafe.Pointer(descriptionC))
	headersC := C.CString(cfg.Headers)
	defer C.free(unsafe.Pointer(headersC))

	ret := C.BASS_Encode_CastInit(C.DWORD(encoder),
		serverC,
		passwordC,
		mimeC,
		nameC,
		urlC,
		genreC,
		descriptionC,
		headersC,
		C.DWORD(cfg.Bitrate),
		C.DWORD(flags),
	)

	if ret <= 0 {
		return bass.GetLastError()
	}

	return nil
}

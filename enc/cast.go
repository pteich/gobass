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

type CastConfig struct {
	Server      string
	Password    string
	ContentType string
	Name        string
	Url         string
	Genre       string
	Description string
	Headers     string
	Bitrate     int
	Public      bool
}

func CastInit(encoder Encoder, cfg CastConfig) error {
	serverC := C.CString(cfg.Server)
	defer C.free(unsafe.Pointer(serverC))
	passwordC := C.CString(cfg.Password)
	defer C.free(unsafe.Pointer(passwordC))
	mimeC := C.CString(cfg.ContentType)
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
	public := 0
	if cfg.Public {
		public = 1
	}

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
		C.int(public),
	)

	if ret <= 0 {
		return bass.GetLastError()
	}

	return nil
}

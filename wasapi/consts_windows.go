package wasapi

import (
	"github.com/pteich/gobass"
)

/*
#include "basswasapi.h"
*/
import "C"

// errors
const (
	ErrWasapi   = bass.Error(C.BASS_ERROR_WASAPI)
	ErrBuffer   = bass.Error(C.BASS_ERROR_WASAPI_BUFFER)
	ErrCategory = bass.Error(C.BASS_ERROR_WASAPI_CATEGORY)
	ErrDenied   = bass.Error(C.BASS_ERROR_WASAPI_DENIED)
)
const (
	Async      = C.BASS_WASAPI_ASYNC
	Autoformat = C.BASS_WASAPI_AUTOFORMAT
	Buffer     = C.BASS_WASAPI_BUFFER
	Dither     = C.BASS_WASAPI_DITHER
	Event      = C.BASS_WASAPI_EVENT
	Exclusive  = C.BASS_WASAPI_EXCLUSIVE
	Raw        = C.BASS_WASAPI_RAW
	Samples    = C.BASS_WASAPI_SAMPLES

	CategoryAlerts         = C.BASS_WASAPI_CATEGORY_ALERTS
	CategoryCommunications = C.BASS_WASAPI_CATEGORY_COMMUNICATIONS
	CategoryGamechat       = C.BASS_WASAPI_CATEGORY_GAMECHAT
	CategoryGameEffects    = C.BASS_WASAPI_CATEGORY_GAMEEFFECTS
	CategoryGameMedia      = C.BASS_WASAPI_CATEGORY_GAMEMEDIA
	CategoryMedia          = C.BASS_WASAPI_CATEGORY_MEDIA
	CategoryMovie          = C.BASS_WASAPI_CATEGORY_MOVIE
	CategoryOther          = C.BASS_WASAPI_CATEGORY_OTHER
	CategorySoundEffects   = C.BASS_WASAPI_CATEGORY_SOUNDEFFECTS
	CategorySpeech         = C.BASS_WASAPI_CATEGORY_SPEECH
)

var (
	ProcBass = C.WASAPIPROC_BASS
	ProcPush = C.WASAPIPROC_PUSH
)

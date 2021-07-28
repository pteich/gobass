package enc
import (

)
/*
#include "bassenc.h"
*/
import "C"
const (
	EncodePCM = C.BASS_ENCODE_PCM
	EncodeAutofree = C.BASS_ENCODE_AUTOFREE
)
package enc

import (
	bass "github.com/pteich/gobass"
)

/*
#include "bassenc.h"
*/
import "C"

// EncodeStartFlags BASS_Encode_Start flags
type EncodeStartFlags bass.Flags

func (f EncodeStartFlags) Add(flag EncodeStartFlags) EncodeStartFlags {
	return EncodeStartFlags(int(f) | int(flag))
}

func (f EncodeStartFlags) Has(flag EncodeStartFlags) bool {
	return int(f)&int(flag) == int(flag)
}

const (
	EncodeNoHead      EncodeStartFlags = C.BASS_ENCODE_NOHEAD       // don't send a WAV header to the encoder
	EncodeFP8Bit      EncodeStartFlags = C.BASS_ENCODE_FP_8BIT      // convert floating-point sample data to 8-bit integer
	EncodeFP16Bit     EncodeStartFlags = C.BASS_ENCODE_FP_16BIT     // convert floating-point sample data to 16-bit integer
	EncodeFP24Bit     EncodeStartFlags = C.BASS_ENCODE_FP_24BIT     // convert floating-point sample data to 24-bit integer
	EncodeFP32Bit     EncodeStartFlags = C.BASS_ENCODE_FP_32BIT     // convert floating-point sample data to 32-bit integer
	EncodeFPAuto      EncodeStartFlags = C.BASS_ENCODE_FP_AUTO      // convert floating-point sample data back to channel's format
	EncodeBigEnd      EncodeStartFlags = C.BASS_ENCODE_BIGEND       // big-endian sample data
	EncodePause       EncodeStartFlags = C.BASS_ENCODE_PAUSE        // start encording paused
	EncodePCM         EncodeStartFlags = C.BASS_ENCODE_PCM          // write PCM sample data (no encoder)
	EncodeRF64        EncodeStartFlags = C.BASS_ENCODE_RF64         // send an RF64 header
	EncodeMono        EncodeStartFlags = C.BASS_ENCODE_MONO         // convert to mono (if not already)
	EncodeQueue       EncodeStartFlags = C.BASS_ENCODE_QUEUE        // queue data to feed encoder asynchronously
	EncodeWFEXT       EncodeStartFlags = C.BASS_ENCODE_WFEXT        // WAVEFORMATEXTENSIBLE "fmt" chunk
	EncodeCastNoLimit EncodeStartFlags = C.BASS_ENCODE_CAST_NOLIMIT // don't limit casting data rate
	EncodeLimit       EncodeStartFlags = C.BASS_ENCODE_LIMIT        // limit data rate to real-time
	EncodeAIFF        EncodeStartFlags = C.BASS_ENCODE_AIFF         // send an AIFF header rather than WAV
	EncodeDither      EncodeStartFlags = C.BASS_ENCODE_DITHER       // apply dither when converting floating-point sample data to integer
	EncodeAutofree    EncodeStartFlags = C.BASS_ENCODE_AUTOFREE     // free the encoder when the channel is freed
	EncodeUnicode     EncodeStartFlags = C.BASS_UNICODE             // cmdline is in UTF-16 form. Otherwise, it is ANSI on Windows or Windows CE, and UTF-8 on other platforms
)

// EncodeCastInitFlags BASS_Encode_CastInit flags
type EncodeCastInitFlags bass.Flags

const (
	EncodeCastPublic EncodeCastInitFlags = C.BASS_ENCODE_CAST_PUBLIC
	EncodeCastPut    EncodeCastInitFlags = C.BASS_ENCODE_CAST_PUT
	EncodeCastSSL    EncodeCastInitFlags = C.BASS_ENCODE_CAST_SSL
)

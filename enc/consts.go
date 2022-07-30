package enc

import (
	bass "github.com/pteich/gobass"
)

/*
#include "bassenc.h"
*/
import "C"

// BASS_Encode_Start flags
const (
	EncodeNoHead      bass.Flags = C.BASS_ENCODE_NOHEAD       // don't send a WAV header to the encoder
	EncodeFP8Bit      bass.Flags = C.BASS_ENCODE_FP_8BIT      // convert floating-point sample data to 8-bit integer
	EncodeFP16Bit     bass.Flags = C.BASS_ENCODE_FP_16BIT     // convert floating-point sample data to 16-bit integer
	EncodeFP24Bit     bass.Flags = C.BASS_ENCODE_FP_24BIT     // convert floating-point sample data to 24-bit integer
	EncodeFP32Bit     bass.Flags = C.BASS_ENCODE_FP_32BIT     // convert floating-point sample data to 32-bit integer
	EncodeFPAuto      bass.Flags = C.BASS_ENCODE_FP_AUTO      // convert floating-point sample data back to channel's format
	EncodeBigEnd      bass.Flags = C.BASS_ENCODE_BIGEND       // big-endian sample data
	EncodePause       bass.Flags = C.BASS_ENCODE_PAUSE        // start encording paused
	EncodePCM         bass.Flags = C.BASS_ENCODE_PCM          // write PCM sample data (no encoder)
	EncodeRF64        bass.Flags = C.BASS_ENCODE_RF64         // send an RF64 header
	EncodeMono        bass.Flags = C.BASS_ENCODE_MONO         // convert to mono (if not already)
	EncodeQueue       bass.Flags = C.BASS_ENCODE_QUEUE        // queue data to feed encoder asynchronously
	EncodeWFEXT       bass.Flags = C.BASS_ENCODE_WFEXT        // WAVEFORMATEXTENSIBLE "fmt" chunk
	EncodeCastNoLimit bass.Flags = C.BASS_ENCODE_CAST_NOLIMIT // don't limit casting data rate
	EncodeLimit       bass.Flags = C.BASS_ENCODE_LIMIT        // limit data rate to real-time
	EncodeAIFF        bass.Flags = C.BASS_ENCODE_AIFF         // send an AIFF header rather than WAV
	EncodeDither      bass.Flags = C.BASS_ENCODE_DITHER       // apply dither when converting floating-point sample data to integer
	EncodeAutofree    bass.Flags = C.BASS_ENCODE_AUTOFREE     // free the encoder when the channel is freed
)

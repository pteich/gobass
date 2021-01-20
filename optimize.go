//Contains optimized functions which can do multiple things in one CGO call.

package bass
/*
#include "bass.h"
static int _BassOptimizeSetPanVolumePitch(HCHANNEL chan, float pan, float volume, float pitch) {
if(BASS_ChannelSetAttribute(chan, BASS_ATTRIB_PAN, pan)==0) { return BASS_ErrorGetCode(); }
if(BASS_ChannelSetAttribute(chan, BASS_ATTRIB_VOL, volume)==0) { return BASS_ErrorGetCode(); }
if(BASS_ChannelSetAttribute(chan, BASS_ATTRIB_FREQ, pitch)==0) { return BASS_ErrorGetCode(); }
return 0;
}
*/
import "C"

func OptimizeSetPanVolumePitch(self Channel, pan, volume, pitch float64) error {
	if C._BassOptimizeSetPanVolumePitch(self.cint(), C.float(pan), C.float(volume), C.float(pitch))==0 {
		return nil
	} else {
		return errMsg()
	}
}
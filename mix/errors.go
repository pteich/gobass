package mix
import (
	"github.com/keithcat1/gobass"
)
/*
#include "bassmix.h"
*/
import "C"

func (self Mixer) toError() (Mixer, error) {
	if self == 0 {
		return 0, bass.GetLastError()
	} else {
		return self, nil
	}
}
func boolToError(value C.BOOL) error {
	if value == 0 {
		return bass.GetLastError()
	} else {
		return nil
	}
}
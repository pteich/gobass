package wasapi
import (
	"github.com/keithcat1/gobass"
)

/*
#include "basswasapi.h"
*/
import "C"
func boolToError(value C.BOOL) error {
	if value==0 {
		return bass.GetLastError()
	} else {
		return nil
	}
}



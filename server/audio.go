package main

// #cgo pkg-config: gstreamer-1.0 glib-2.0
// #include "gstreamer.c"
import "C"
import "errors"

//  API  /////////////////////////////////////////////////////////////////////

func audioPlay(index int) (err error) {
	currentTrack = index
	currentState = StatePlaying

	cerr, err := C.play_file(C.CString(items[currentTrack].Path))
	if err != nil {
		return
	}
	if cerr != 0 {
		err = errors.New("c function didn't return 0")
		return
	}
	return
}

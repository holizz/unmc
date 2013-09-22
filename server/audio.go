package main

// #cgo pkg-config: gstreamer-1.0 glib-2.0
// #include "gstreamer.c"
import "C"
import "errors"

//  API  /////////////////////////////////////////////////////////////////////

// func audioInit() (err error) {
// 	return
// }

func audioPlay(id int) (err error) {
	i, err := itemIndex(id)
	if err != nil {
		return
	}

	currentTrack = i
	currentState = StatePlaying

	cerr, err := C.play_file(C.CString(items[i].Path))
	if err != nil {
		return
	}
	if cerr != 0 {
		err = errors.New("c function didn't return 0")
		return
	}
	return
}

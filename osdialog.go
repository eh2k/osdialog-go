package osdialog

// #cgo linux CFLAGS: -w
// #cgo linux pkg-config: gtk+-3.0
// #cgo windows LDFLAGS: -lcomdlg32
// #cgo windows CFLAGS: -w
// #cgo darwin LDFLAGS: -framework AppKit
// #cgo darwin CFLAGS: -w -x objective-c -mmacosx-version-min=10.7
// #include <osdialog/osdialog.h>
// #include <osdialog/osdialog.c>
// #ifdef linux
// #include <osdialog/osdialog_gtk3.c>
// #endif
// #ifdef WIN32
// #include <osdialog/osdialog_win.c>
// #endif
// #ifdef __APPLE__
// #include <osdialog/osdialog_mac.m>
// #endif
import "C"
import "unsafe"
import "errors"

func ShowOpenFileDialog(path string, filename string, filters string) (string, error) {

	f := C.CString(filters)
	defer C.free(unsafe.Pointer(f))

	p := C.CString(path)
	defer C.free(unsafe.Pointer(p))

	fn := C.CString(filename)
	defer C.free(unsafe.Pointer(fn))

	ff := C.osdialog_filters_parse(f)
	defer C.osdialog_filters_free(ff)

	selected := C.osdialog_file(C.OSDIALOG_OPEN, p, fn, ff)

	if unsafe.Pointer(selected) != C.NULL {
		defer C.free(unsafe.Pointer(selected))
		return C.GoString(selected), nil
	} else {
		return "", errors.New("ShowOpenFileDialog canceled")
	}
}

func ShowSaveFileDialog(path string, filename string, filters string) (string, error) {

	f := C.CString(filters)
	defer C.free(unsafe.Pointer(f))

	p := C.CString(path)
	defer C.free(unsafe.Pointer(p))

	fn := C.CString(filename)
	defer C.free(unsafe.Pointer(fn))

	ff := C.osdialog_filters_parse(f)
	defer C.osdialog_filters_free(ff)

	selected := C.osdialog_file(C.OSDIALOG_SAVE, p, fn, ff)

	if unsafe.Pointer(selected) != C.NULL {
		defer C.free(unsafe.Pointer(selected))
		return C.GoString(selected), nil
	} else {
		return "", errors.New("ShowSaveFileDialog canceled")
	}
}

func ShowOpenDirectoryDialog(path string) (string, error) {

	p := C.CString(path)
	defer C.free(unsafe.Pointer(p))

	selected := C.osdialog_file(C.OSDIALOG_OPEN_DIR, p, nil, nil)

	if unsafe.Pointer(selected) != C.NULL {
		defer C.free(unsafe.Pointer(selected))
		return C.GoString(selected), nil
	} else {
		return "", errors.New("ShowOpenDirectoryDialog canceled")
	}
}

type Level int32

const (
	Info    Level = 0
	Warning Level = 1
	Error   Level = 2
)

type Buttons int32

const (
	Ok       Buttons = 0
	OkCancel Buttons = 1
	YesNo    Buttons = 2
)

func ShowMessageBox(level Level, buttons Buttons, message string) bool {

	m := C.CString(message)
	defer C.free(unsafe.Pointer(m))

	ok := C.osdialog_message(C.osdialog_message_level(level), C.osdialog_message_buttons(buttons), m)

	return int32(ok) == 1
}

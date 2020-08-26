![build](https://github.com/eh2k/osdialog-go/workflows/build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/eh2k/osdialog-go)](https://goreportcard.com/report/github.com/eh2k/osdialog-go)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/eh2k/osdialog-go)
# osdialog-go

A Golang cross platform API for OS dialogs like file save, open, message boxes.
This library is a Go wrapper for https://github.com/AndrewBelt/osdialog

# UseCases:

MessageBox:

    ok := osdialog.ShowMessageBox(osdialog.Error, osdialog.YesNo, "Error Message!!!")

Open File:

    filename, err := osdialog.ShowOpenFileDialog(".", "こんにちは", "Source:c,cpp,m;Header:h,hpp")

Save File:

    filename, err := osdialog.ShowSaveFileDialog(".", "こんにちは", "Source:c,cpp,m;Header:h,hpp")

Open Directory:

    directory, err := osdialog.ShowOpenDirectoryDialog(".")

# platform details
* Linux: uses Gtk3 `sudo apt-get install -y libgtk-3-dev libcairo2-dev libpango1.0-dev`
* Win32: uses Win32 API
* OSX: not testet

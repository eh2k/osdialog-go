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

    directory, err := osdialog.ShowDirectoryFileDialog(".")

# platform details
* Linux: uses Gtk3 `sudo apt-get install -y libgtk-3-dev libcairo2-dev libpango1.0-dev`
* Win32: uses Win32 API
* OSX: not testet

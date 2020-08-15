package main

import "github.com/eh2k/osdialog-go"

func main() {

	yes := osdialog.ShowMessageBox(osdialog.Error, osdialog.YesNo, "Error Message!!!")
	if yes {
		osdialog.ShowMessageBox(osdialog.Info, osdialog.Ok, "Info!!!")
	}

	osdialog.ShowSaveFileDialog(".", "こんにちは", "Source:c,cpp,m;Header:h,hpp")

	osdialog.ShowOpenFileDialog(".", "こんにちは", "Source:c,cpp,m;Header:h,hpp")

	osdialog.ShowOpenDirectoryDialog("/")

}

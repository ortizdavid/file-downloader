package main

import (
	"fmt"
	"os"
	"github.com/ortizdavid/file-downloader/core"
	"github.com/ortizdavid/file-downloader/helpers"
)

func main() {

	var cliArgs = os.Args
	var lenArgs = len(cliArgs)
	var fileManager *core.FileManager
	var directory = "downloads/"

	if lenArgs < 2 {
		helpers.Usage()

	} else if lenArgs == 2{

		switch cliArgs[1] {
		case "-help":
			helpers.PrintHelp()
		case "-examples":
			helpers.Examples()
		case "-list":
			fileManager.ListFiles(directory)
		default:
			helpers.BasicInfo()
		}

	} else if lenArgs == 3{
		var downloader *core.Downloader
		var fileUrl = cliArgs[2]
		var file = helpers.GetNameFromUrl(fileUrl)

		fileManager.CreateFolder(directory)
		err := downloader.DownloadFile(directory+file, fileUrl)
		if err != nil {
			panic(err)
		}
		fmt.Println("Downloaded From: " + fileUrl)
		fileManager.GetFileInfo(directory+file)
	}
}


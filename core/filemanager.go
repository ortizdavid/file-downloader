package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"github.com/ortizdavid/file-downloader/helpers"
)

type FileManager struct {}

func (fm *FileManager) GetFileExt(path string) string {
	return filepath.Ext(path)
}

func (fm *FileManager) GetFileType(ext string) string {
	var imageExts = []string{".png", ".gif", ".jpg", ".jiff"}
	var documentExts = []string{".txt", ".pdf", ".docx", ".ppt", ".pptx"}
	var audioExts = []string{".mp3", ".aac", ".wav", ".flac"}
	var videoExts = []string{".mp4", ".mkv", ".avi", ".flv"}
	var fileType string = ""

	if helpers.Contains(imageExts, ext) {
		fileType = "Image"
	} else if helpers.Contains(documentExts, ext) {
		fileType = "Document"
	} else if helpers.Contains(audioExts, ext) {
		fileType = "Audio"
	} else if helpers.Contains(documentExts, ext) {
		fileType = "Document"
	} else if helpers.Contains(videoExts, ext) {
		fileType = "Video"
	} else {
		fileType = "Other"
	}
	return fileType
}

func (fm *FileManager) CreateFolder(folderName string) error  {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		return os.Mkdir(folderName, os.ModeDir|0755)
	}
	return nil
}

func (fm *FileManager) GetFileInfo(fileName string) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nFile Name:", fileInfo.Name())  
	fmt.Println("Extension:", fm.GetFileExt(fileName))          
	fmt.Println("Size:", fileInfo.Size(), " bytes")  
	fmt.Println("Type:", fm.GetFileType(fm.GetFileExt(fileName)))              
	fmt.Println("Last Modified:", fileInfo.ModTime()) 
	fmt.Println("Permissions:", fileInfo.Mode())     
}

func (fm *FileManager) ListFiles(dir string)  {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nAll Files in '%s':\n", dir)
	helpers.PrintChar("-", 125)
	fmt.Println("NAME\t\t\t\tTYPE\t\t\tSIZE(Bytes)\t\t\tLAST MODIFIED")
	helpers.PrintChar("-", 125)
	for _, file := range files {
		var ext = fm.GetFileExt((file.Name()))
		fmt.Printf("%s\t\t\t%s\t\t\t%d\t\t\t%s\n", file.Name(), fm.GetFileType(ext), file.Size(), file.ModTime())
	}
}
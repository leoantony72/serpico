package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"syscall"
	"strconv"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)


// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// const organizeOptions = ["Year", "Month", "File Type"];
func(a *App) OrganizeFolder(path,Organizeby string){
	switch Organizeby {
	case "File Type": OrganizeByFile(path)
	case "Year": OrganizebyYear(path)
	case "Month": OrganizebyMonth(path)
	}
}

func OrganizebyMonth(path string){
	orgf := make(map[string]int)
	files, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer files.Close()

	fileinfo, err := files.Readdir(-1)
	if err != nil {
		fmt.Println("error reading directory:", err) 
		//if directory is not read properly print error message
		return
	}

	for _, f := range fileinfo {
		d := f.Sys().(*syscall.Win32FileAttributeData)
		cTime := time.Unix(0, d.CreationTime.Nanoseconds())
		// t := cTime.Month()
		des := cTime.Month().String()
		fmt.Printf("File: %s, Year Created: %d\n", f.Name(), des)
		
		destDir := filepath.Join(path, des)
		if _,ok := orgf[des]; !ok{
			orgf[des] = 0
			os.Mkdir(destDir, 0755)
		}

		oldPath := filepath.Join(path, f.Name())
		newPath := filepath.Join(destDir, f.Name())

		// Move file to new directory
		err = os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Println("Error moving file:", err)
			continue
		}

		fmt.Printf("Moved %s to %s\n", oldPath, newPath)
	} 
}

func OrganizebyYear(path string){
	orgf := make(map[int]int)
	files, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer files.Close()

	fileinfo, err := files.Readdir(-1)
	if err != nil {
		fmt.Println("error reading directory:", err) //if directory is not read properly print error message
		return
	}

	for _, f := range fileinfo {
		d := f.Sys().(*syscall.Win32FileAttributeData)
		cTime := time.Unix(0, d.CreationTime.Nanoseconds())
		// t := cTime.Year()
		fmt.Printf("File: %s, Year Created: %d\n", f.Name(), cTime.Year())
		
		des := strconv.Itoa(cTime.Year())
		destDir := filepath.Join(path, des)
		if _,ok := orgf[cTime.Year()]; !ok{
			orgf[cTime.Year()] = 0
			os.Mkdir(destDir, 0755)
		}

		oldPath := filepath.Join(path, f.Name())
		newPath := filepath.Join(destDir, f.Name())

		// Move file to new directory
		err = os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Println("Error moving file:", err)
			continue
		}

		fmt.Printf("Moved %s to %s\n", oldPath, newPath)
	} 
}

func OrganizeByFile(path string){
	file_ext := map[string]string{
		".txt": "text file",
		".pdf": "pdf",
		".jpg": "image",
		".png": "image",
		".jpeg":"image",
		".mp3": "audio",
		".ppt": "powerpoint",
		".mkv": "video",
		".mp4": "video",
		".zip": "zip files",
		".csv": "csv files",
		".xlsx":"spreadsheets",
		".msi" :"software",
		".apk" :"software",
		".exe" :"software",
	}

	files, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer files.Close()

	fileinfo, err := files.Readdir(-1)
	if err != nil {
		fmt.Println("error reading directory:", err) //if directory is not read properly print error message
		return
	}
	for _, f := range fileinfo {

		if !f.IsDir() {
			ext := filepath.Ext(f.Name())
			fmt.Println(f.Name())
			if des, ok := file_ext[ext]; ok {
				// p := path+
				destDir := filepath.Join(path, des)
				os.Mkdir(destDir, 0755)

				oldPath := filepath.Join(path, f.Name())
				newPath := filepath.Join(destDir, f.Name())

				// Move file to new directory
				err = os.Rename(oldPath, newPath)
				if err != nil {
					fmt.Println("Error moving file:", err)
					continue
				}

				fmt.Printf("Moved %s to %s\n", oldPath, newPath)
			}

		}
	}

}

func (a *App) SelectDirectory() (string, error) {
	path, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
	  Title: "Select Folder",
	})
	if err != nil {
	  return "", err
	}
	return path, nil
}
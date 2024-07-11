package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"net/http"
	"strings"
	"syscall"
	"strconv"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"Digital_deculutterer/duplicate"
	// "sort"
)


// App struct
type App struct {
	ctx context.Context
	selectedFolder string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	http.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
		if a.selectedFolder != "" {
			filePath := strings.TrimPrefix(r.URL.Path, "/assets/")
			fullPath := filepath.Join(a.selectedFolder, filePath)
			http.ServeFile(w, r, fullPath)
		} else {
			http.Error(w, "No folder selected", http.StatusNotFound)
		}
	})

	// Start the file server in a new goroutine
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("Error starting file server:", err)
		}
	}()
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
		// fmt.Printf("File: %s, Year Created: %d\n", f.Name(), des)
		
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
func (a *App) SelectDirectoryDuplicate() (string, error) {
	path, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
	  Title: "Select Folder",
	})
	if err != nil {
	  return "", err
	}
	a.selectedFolder = path
	return path, nil
}

func (a *App) FindDuplicates(path string) map[string][]string {
	fileInfos, err := duplicate.LoadFilesFromFolder(path)
    if err != nil {
        fmt.Println("Error loading files:", err)
        return nil
    }

    fileInfos, err = duplicate.GenerateFileHashes(fileInfos)
    if err != nil {
        fmt.Println("Error generating file hashes:", err)
        return nil
    }

    duplicates := duplicate.FindDuplicates(fileInfos)

    // Log the duplicates for debugging
    fmt.Println("Found duplicates:", duplicates)

    return duplicates
}
func (a *App) DeleteFiles(filePaths []string) error {
    for _, filePath := range filePaths {
        err := os.Remove(filePath)
        if err != nil {
            fmt.Println("Error deleting file:", filePath, err)
            return err
        }
        fmt.Println("Deleted file:", filePath)
    }
    return nil
}
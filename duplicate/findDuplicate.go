package duplicate

import (
    // "fmt"
    // "github.com/glaslos/ssdeep"
    // "io"
    // "bufio"
    "encoding/hex"
    "crypto/sha256"
    "os"
    "io"
    "sort"
    // "strings"
    "time"
    // "io"
    "path/filepath"
)


type FileInfo struct {
    Path       string
    Hash       string
    Content   []byte
    CreateTime time.Time
}


// // LoadFilesFromFolder loads files from the specified folder and returns their file details
func LoadFilesFromFolder(folder string) ([]FileInfo, error) {
    var fileInfos []FileInfo

    err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            file, err := os.Open(path)
            if err != nil {
                return err
            }
            defer file.Close()

            content, err := io.ReadAll(file)
            if err != nil {
                return err
            }

            // Get the file creation time
            stat, err := os.Stat(path)
            if err != nil {
                return err
            }
            createTime := stat.ModTime()

            fileInfos = append(fileInfos, FileInfo{
                Path:       path,
                Content: content,
                CreateTime: createTime,
            })
        }
        return nil
    })

    return fileInfos, err
}

// GenerateFileHashes generates SHA-256 hashes for the files
func GenerateFileHashes(fileInfos []FileInfo) ([]FileInfo, error) {
    for i, fileInfo := range fileInfos {
        hash := sha256.Sum256(fileInfo.Content)
        fileInfos[i].Hash = hex.EncodeToString(hash[:])
    }
    return fileInfos, nil
}

func FindDuplicates(fileInfos []FileInfo) map[string][]string {
    duplicates := make(map[string][]string)
    hashMap := make(map[string][]FileInfo)

    // Group files by their hash
    for _, fileInfo := range fileInfos {
        hashMap[fileInfo.Hash] = append(hashMap[fileInfo.Hash], fileInfo)
    }

    // Iterate over groups of files with the same hash
    for _, files := range hashMap {
        if len(files) > 1 {
            // Sort files by creation date, oldest first
            sort.Slice(files, func(i, j int) bool {
                return files[i].CreateTime.Before(files[j].CreateTime)
            })

            // Use the first file (oldest) as the key
            original := files[0].Path

            // Collect all similar files except the original
            var similarFiles []string
            for _, file := range files {
                if file.Path != original {
                    similarFiles = append(similarFiles, file.Path)
                }
            }

            // Add to duplicates map only if there are duplicates
            if len(similarFiles) > 0 {
                duplicates[original] = similarFiles
            }
        }
    }

    return duplicates
}
// isCopyFile function to check if the filename contains variations of "copy"
// func isCopyFile(fileName string) bool {
// 	lowerCaseFileName := strings.ToLower(fileName)
// 	return strings.Contains(lowerCaseFileName, " copy") || strings.Contains(lowerCaseFileName, "(copy") || strings.Contains(lowerCaseFileName, "copy")
// }
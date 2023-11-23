package upload_utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFieldHandler(ctx *gin.Context) {

	forms, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id := ctx.PostForm("id")
	folder := ctx.PostForm("folder")

	if folder == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "folder name is required"})
		return
	}

	files := forms.File["files"]
	if len(files) == 0 && id != "null" {
		ctx.JSON(http.StatusOK, gin.H{"status": 200, "message": "Success"})
		return
	} else if len(files) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "No Files Found"})
		return
	}

	uploadFolder := fmt.Sprintf("./public/images/%s", folder)
	nameTime := fmt.Sprintf("_%d", time.Now().UTC().Unix())

	// Ensure the upload folder exists
	if err := ensureUploadFolder(uploadFolder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": err.Error()})
		return
	}

	var fileNameArray []string

	for _, file := range files {

		filename := filepath.Base(file.Filename)
		extension := filepath.Ext(filename)
		noSpacesFileName := strings.ReplaceAll(filename, " ", "_")

		newFilename := strings.TrimSuffix(noSpacesFileName, extension) + nameTime + extension
		fullPath := filepath.Join(uploadFolder, newFilename)

		// Save the uploaded file
		if err := ctx.SaveUploadedFile(file, fullPath); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": err.Error()})
			return
		}

		fileNameArray = append(fileNameArray, newFilename)
	}

	jsonStringfy, err := json.Marshal(fileNameArray)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": err.Error()})
		return
	}

	// Respond with a success message
	ctx.JSON(http.StatusOK, gin.H{"status": 200, "message": "Success", "filename": string(jsonStringfy)})
}

func ensureUploadFolder(uploadFolder string) error {
	// Ensure the directory for file uploads exists
	if _, err := os.Stat(uploadFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadFolder, 0755); err != nil {
			return err
		}
	}
	return nil
}

package helpers

import (
	"context"
	"fmt"
	"io"

	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/bivek/fmt_backend/constants"
)

func FileUpload(ctx context.Context, fileHeader *multipart.FileHeader, folderName constants.Folder) (string, error) {
	fileExtension := filepath.Ext(fileHeader.Filename)
	file, err := fileHeader.Open()

	defer file.Close()

	err = os.MkdirAll("./clients-image", os.ModePerm)

	dst, err := os.Create(fmt.Sprintf("./clients-image/%d%s", time.Now().UnixNano(), fileExtension))

	//defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		fmt.Printf("Failed to copy")
	}

	fmt.Printf("filename %v", fileHeader.Filename)

	return fileHeader.Filename, err
}

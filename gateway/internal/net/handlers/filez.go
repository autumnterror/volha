package handlers

import (
	"errors"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

var allowedMIMEs = map[string]string{
	"image/jpeg": ".jpg",
	"image/png":  ".png",
	"image/webp": ".webp",
	"image/gif":  ".gif",
}

func newName() string {
	return uuid.NewString()
}

const (
	formFieldName  = "img"
	ImagesDir      = "./images"
	MaxUploadBytes = 10 << 20 // 10 MB
)

// UploadFile godoc
// @Summary Загрузить изображение
// @Description Сохраняет изображение в папке сервера images/. В поле "img" передается файл с расширениями jpg, png, webp, gif
// @Tags files
// @Accept json
// @Produce json
// @Success 200 {object} views.SWGFileUploadResponse "Цвет успешно создан"
// @Failure 400 {object} views.SWGErrorResponse "Неверный формат данных"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Router /api/files/upload [post]
func (a *Apis) UploadFile(c echo.Context) error {
	fileHeader, err := c.FormFile(formFieldName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "file field 'img' is required"})
	}

	src, err := fileHeader.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "cannot open uploaded file"})
	}
	defer src.Close()

	buff := make([]byte, 512)
	n, _ := io.ReadFull(src, buff)
	if _, err := src.(io.Seeker).Seek(0, io.SeekStart); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "cannot seek file"})
	}

	detected := http.DetectContentType(buff[:n])
	ext, ok := allowedMIMEs[detected]
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "unsupported file type"})
	}

	filename := newName() + ext
	dstPath := filepath.Join(ImagesDir, filename)

	if !strings.HasPrefix(filepath.Clean(dstPath), filepath.Clean(ImagesDir)) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid path"})
	}

	if err := saveMultipartFile(src, dstPath, fileHeader); err != nil {
		var msg string
		if errors.Is(err, errTooLarge) {
			msg = "file is too large"
		} else {
			msg = "failed to save file"
		}
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": msg})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"name": filename,
		"mime": detected,
	})
}

var errTooLarge = errors.New("too large")

func saveMultipartFile(src multipart.File, dstPath string, fh *multipart.FileHeader) error {
	if fh.Size > MaxUploadBytes {
		return errTooLarge
	}

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()
	if _, err := io.Copy(dst, src); err != nil {
		_ = os.Remove(dstPath)
		return err
	}

	if err := os.Chmod(dstPath, 0644); err != nil {
		return err
	}

	return nil
}

// DeleteFile godoc
// @Summary Удалить файл
// @Description Удаляет файл. Требуется передовать только имя. Пример: example.gif
// @Tags files
// @Produce json
// @Param title query string true "название файла"
// @Success 200 {object} views.SWGSuccessResponse "файл успешно удалён"
// @Failure 400 {object} views.SWGErrorResponse "Неверное название"
// @Failure 500 {object} views.SWGErrorResponse "Ошибка на сервере"
// @Router /api/files/delete [delete]
func (a *Apis) DeleteFile(c echo.Context) error {
	filename := c.QueryParam("title")

	if strings.Contains(filename, "/") || strings.Contains(filename, "\\") {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid filename"})
	}

	path := filepath.Join("images", filename)

	if err := deleteImage(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "file not found"})
		}
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"deleted": path})
}

func deleteImage(relPath string) error {
	cleanPath := strings.TrimPrefix(relPath, "./")
	if !strings.HasPrefix(cleanPath, "images/") && !strings.HasPrefix(cleanPath, "images\\") {
		return errors.New("invalid image path " + cleanPath)
	}

	fullPath := filepath.Join(".", cleanPath)
	fullPath = filepath.Clean(fullPath)

	if !strings.HasPrefix(fullPath, filepath.Clean(ImagesDir)) {
		return errors.New("unsafe path")
	}

	if err := os.Remove(fullPath); err != nil {
		return err
	}
	return nil
}

package utils

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"go-management-auth-school/helper/str"
	"io"
	"io/ioutil"
	"math/rand"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"github.com/rs/xid"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	defaultLimit  = 10
	maxLimit      = 50
	defaultSort   = "asc"
	maxUploadSize = 5000000
	staticFile    = "./static"
)

var (
	sortWhitelist     = []string{"asc", "desc"}
	fileTypeWhitelist = []string{"image/jpeg", "image/jpg", "image/png"}
)

func GenerateProcessID() string {
	b := make([]byte, 20)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func SetPaginationParameter(page, limit int, orderBy, sort string, orderByWhiteLists, orderByStringWhiteLists []string) (int, int, int, string, string) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > maxLimit {
		limit = defaultLimit
	}

	orderBy = checkWhiteList(orderBy, orderByWhiteLists)
	if str.Contains(orderByStringWhiteLists, orderBy) {
		orderBy = `LOWER(` + orderBy + `)`
	}

	if !str.Contains(sortWhitelist, sort) {
		sort = defaultSort
	}
	offset := (page - 1) * limit

	return offset, limit, page, orderBy, sort
}

func checkWhiteList(orderBy string, whiteLists []string) string {
	for _, whiteList := range whiteLists {
		if orderBy == whiteList {
			return orderBy
		}
	}

	return "def.updated_at"
}

func UploadImage(w http.ResponseWriter, r *http.Request, imageName string) (res Image, err error) {
	file, _, err := r.FormFile(imageName)
	if err != nil {
		return res, errors.New("invalid_file")
	}
	defer file.Close()

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxUploadSize))
	err = r.ParseMultipartForm(int64(maxUploadSize))
	if err != nil {
		return res, errors.New("file_too_big")
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return res, errors.New("invalid_file")
	}

	filetype := http.DetectContentType(fileBytes)
	if !str.Contains(fileTypeWhitelist, filetype) {
		return res, errors.New("invalid_file_type")
	}

	fileType := imageName
	fileName := fileType + "-" + xid.New().String()
	fileEndings, err := mime.ExtensionsByType(filetype)
	if err != nil {
		return res, errors.New("invalid_file_type")
	}

	fullFileName := "/" + fileName + fileEndings[0]

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, staticFile)
	// Create static folder for file uploading
	if _, err := os.Stat(filesDir); os.IsNotExist(err) {
		if err = os.MkdirAll(filesDir, os.ModePerm); err != nil {
			return res, err
		} else {
			fmt.Println("Creating temporary folder ", filesDir)
		}
	}

	uploadURL := filesDir + fullFileName
	err = ioutil.WriteFile(uploadURL, fileBytes, 0644)
	if err != nil {
		return
	}

	path := "/" + fileType + fullFileName

	res = Image{
		ImagePath: path,
		UploadURL: uploadURL,
		ImageType: fileType,
		ImageName: fullFileName,
	}

	return
}

// ExtractCSVContent does file type checking, then extracts the file contents
func UploadCSV(fileIO io.Reader) (res [][]string, err error) {
	var bB bytes.Buffer
	file := io.TeeReader(fileIO, &bB)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return res, errors.New("invalid_file")
	}

	fileType := http.DetectContentType(fileBytes)
	if fileType != "text/plain; charset=utf-8" {
		return res, errors.New("invalid_file_type")
	}

	csvReader := csv.NewReader(&bB)
	csvReader.FieldsPerRecord = -1
	res, err = csvReader.ReadAll()
	if err != nil {
		return res, errors.New("unable to parse file as csv")
	}

	return
}

package response

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
)

const (
	CONTENT_TYPE     = "Content-Type"
	APPLICATION_JSON = "application/json"
)


type ApiResponseModel struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data"`
	Meta      interface{} `json:"meta,omitempty"`
	ErrorCode int         `json:"error_code,omitempty"`
	StatMsg   string      `json:"stat_msg,omitempty"`
}

func RespondSuccess(w http.ResponseWriter, statusCode int, data, meta interface{}) {
	bt, err := json.Marshal(ApiResponseModel{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
	if err != nil {
		bt, _ = json.Marshal(ApiResponseModel{
			Success: false,
			StatMsg: "",
		})
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(bt)
		return
	}
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(statusCode)
	w.Write(bt)
	return
}

func RespondError(w http.ResponseWriter, statusCode int, err error) {
	bt, err := json.Marshal(ApiResponseModel{
		Success: false,
		Data:    nil,
		StatMsg: err.Error(),
	})
	if err != nil {
		bt, _ = json.Marshal(ApiResponseModel{
			Success: false,
			StatMsg: "Internal Server Error"},
		)
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(bt)
		return
	}
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(statusCode)
	w.Write(bt)
	return
}

func RespondErrorWithMessage(w http.ResponseWriter, statusCode int, err error) {
	statMsg, exist := errMessageMaps[err]
	if !exist {
		statMsg = err.Error()
	}

	bt, err := json.Marshal(ApiResponseModel{
		Success: false,
		Data:    nil,
		StatMsg: statMsg,
	})
	if err != nil {
		bt, _ = json.Marshal(ApiResponseModel{
			Success: false,
			StatMsg: "Internal Server Error"},
		)
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(bt)
		return
	}
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(statusCode)
	w.Write(bt)
	return
}

func RespondStatMsg(w http.ResponseWriter, statusCode int, errMessages string) {
	bt, err := json.Marshal(ApiResponseModel{
		Success: false,
		Data:    nil,
		StatMsg: errMessages,
	})
	if err != nil {
		bt, _ = json.Marshal(ApiResponseModel{
			Success: false,
			StatMsg: "Internal Server Error",
		})
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(bt)
		return
	}
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	w.WriteHeader(statusCode)
	w.Write(bt)
}

// SendFileSuccess send success file into response with 200 http code.
func SendFileSuccess(w http.ResponseWriter, fileLocation, fileContentType string, unitTest bool) {
	RespondWithFile(w, 200, fileLocation, fileContentType, unitTest)
}

// RespondWithFile write file response format
func RespondWithFile(w http.ResponseWriter, httpCode int, fileLocation, fileContentType string, unitTest bool) {
	// Open result file
	fileRes, err := os.OpenFile(fileLocation, os.O_RDWR, 0644)
	if err != nil {
		RespondError(w, http.StatusNotFound, err)
		return
	}
	fi, err := fileRes.Stat()
	if err != nil {
		// Delete temporary pdf file
		fileRes.Close()
		if !unitTest {
			os.Remove(fileLocation)
		}

		RespondError(w, http.StatusNotFound, err)
		return
	}

	// copy the relevant headers. If you want to preserve the downloaded file name, extract it with go's url parser.
	w.Header().Set("Content-Disposition", "attachment; filename="+fi.Name())
	w.Header().Set(CONTENT_TYPE, fileContentType)
	w.Header().Set("Content-Length", strconv.Itoa(int(fi.Size())))

	// Stream the body to the client without fully loading it into memory
	w.WriteHeader(httpCode)
	io.Copy(w, fileRes)

	// Delete temporary pdf file
	fileRes.Close()
	if !unitTest {
		os.Remove(fileLocation)
	}
}

package response

import (
	"encoding/json"
	// "io"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
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



func RespondSuccess(c echo.Context, statusCode int, data, meta interface{}) (err error)  {
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
		c.Response().Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Response().Write(bt)
		return 
	}
	c.Response().Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	c.Response().WriteHeader(statusCode)
	c.Response().Write(bt)
	return err
}

func RespondError(c echo.Context, statusCode int, err error) (error) {
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
		c.Response().Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Response().Write(bt)
		return err
	}
	c.Response().Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	c.Response().WriteHeader(statusCode)
	c.Response().Write(bt)
	return err
}

func RespondErrorWithMessage(c echo.Context, statusCode int, err error) error {
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
		c.Response().Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Response().Write(bt)
		return err
	}
	c.Response().Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	c.Response().WriteHeader(statusCode)
	return err
}

func RespondStatMsg(c echo.Context, statusCode int, errMessages string) error {
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
		c.Response().Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		c.Response().WriteHeader(http.StatusInternalServerError)
		c.Response().Write(bt)
		return err
	}
	c.Response().Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	c.Response().WriteHeader(statusCode)
	return err
}

// SendFileSuccess send success file into response with 200 http code.
func SendFileSuccess(c echo.Context, fileLocation, fileContentType string, unitTest bool) {
	RespondWithFile(c, 200, fileLocation, fileContentType, unitTest)
}

// RespondWithFile write file response format
func RespondWithFile(c echo.Context, httpCode int, fileLocation, fileContentType string, unitTest bool) error {
	// Open result file
	fileRes, err := os.OpenFile(fileLocation, os.O_RDWR, 0644)
	if err != nil {
		return RespondError(c, http.StatusNotFound, err)
	}
	fi, err := fileRes.Stat()
	if err != nil {
		// Delete temporary pdf file
		fileRes.Close()
		if !unitTest {
			os.Remove(fileLocation)
		}

		return RespondError(c, http.StatusNotFound, err)
	}

	// copy the relevant headers. If you want to preserve the downloaded file name, extract it with go's url parser.
	c.Response().Header().Set("Content-Disposition", "attachment; filename="+fi.Name())
	c.Response().Header().Set(CONTENT_TYPE, fileContentType)
	c.Response().Header().Set("Content-Length", strconv.Itoa(int(fi.Size())))

	// Stream the body to the client without fully loading it into memory
	c.Response().WriteHeader(httpCode)
	// io.Copy(c.Response().Writer, fileRes)

	// Delete temporary pdf file
	fileRes.Close()
	if !unitTest {
		os.Remove(fileLocation)
	}
	return c.Stream(httpCode, fileContentType, fileRes)
}

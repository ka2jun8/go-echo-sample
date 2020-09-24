package test

import (
	"bytes"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ka2jun8/go-echo-sample/server"
)

// Input is type of test input information
type FilesTestInput struct {
	method    string
	path      string
	filenames []string
}

// Expect is type of test expectation
type FilesTestExpect struct {
	status  int
	message string
}

func getSampleFileBody(filenames []string) (body *bytes.Buffer, contentType string, err error) {
	fieldname := "file"
	body = &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if len(filenames) == 0 {
		// テスト用にmultipartではなくjsonを送る
		jsonStr := []byte(`{"hello":"world"}`)
		body = bytes.NewBuffer(jsonStr)
		contentType = "application/json"
		return
	}

	for _, filename := range filenames {
		file, err := os.Open("fixture/" + filename)
		if err != nil {
			fmt.Println("Error: failed to open file: " + err.Error())
			break
		}
		defer file.Close()

		part, err := writer.CreateFormFile(fieldname, filename)
		if err != nil {
			fmt.Println("Error: failed to create part: " + err.Error())
			break
		}

		if _, err = io.Copy(part, file); err != nil {
			fmt.Println("Error: io.Copy file to form writer: " + err.Error())
			break
		}

		if err != nil {
			fmt.Printf("Error: something happens: %s\n", err.Error())
			break
		}
	}
	if err != nil {
		return
	}

	contentType = writer.FormDataContentType()
	fmt.Printf("request content type is: %s\n", contentType)

	err = writer.Close()
	if err != nil {
		fmt.Printf("Error: multipart writer.Close: %s\n", err.Error())
		return
	}

	return
}

func TestFilesHandler(t *testing.T) {
	var tests = []struct {
		description string
		input       FilesTestInput
		expect      FilesTestExpect
	}{
		{
			"単体ファイル",
			FilesTestInput{
				method:    "POST",
				path:      "/files",
				filenames: []string{"sample.jpg"},
			},
			FilesTestExpect{
				status:  http.StatusOK,
				message: "[{\"index\":0,\"file_name\":\"sample.jpg\"}]\n",
			},
		},
		{
			"複数ファイル",
			FilesTestInput{
				method:    "POST",
				path:      "/files",
				filenames: []string{"sample.jpg", "sample2.png"},
			},
			FilesTestExpect{
				status:  http.StatusOK,
				message: "[{\"index\":0,\"file_name\":\"sample.jpg\"},{\"index\":1,\"file_name\":\"sample2.png\"}]\n",
			},
		},
		{
			"画像以外のファイルではBadRequestになる",
			FilesTestInput{
				method:    "POST",
				path:      "/files",
				filenames: []string{"sample.txt", "sample2.png"},
			},
			FilesTestExpect{
				status:  http.StatusBadRequest,
				message: "{\"message\":\"Bad Request\"}\n",
			},
		},
		{
			"multipart以外を送信するとBadRequestになる",
			FilesTestInput{
				method:    "POST",
				path:      "/files",
				filenames: []string{},
			},
			FilesTestExpect{
				status:  http.StatusBadRequest,
				message: "{\"message\":\"Bad Request\"}\n",
			},
		},
	}

	for _, test := range tests {
		fmt.Println(test.description)

		body, contentType, err := getSampleFileBody(test.input.filenames)
		if err != nil {
			t.Error(err)
		}

		req := httptest.NewRequest(test.input.method, test.input.path, body)
		req.Header.Set("Content-Type", contentType)
		rec := httptest.NewRecorder()
		router := server.Router()
		router.ServeHTTP(rec, req)

		status := test.expect.status
		message := test.expect.message
		actualStatus := rec.Code
		actualMessage := rec.Body.String()

		if status != actualStatus {
			t.Errorf("expected: %v, result: %q", status, actualStatus)
		}
		if message != actualMessage {
			t.Errorf("expected: %v, result: %q", message, actualMessage)
		}
	}
}

package handlers

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

// for debug
const formHTML = `<!DOCTYPE html>
<html>
  <head>
    <title>Storage</title>
    <meta charset="utf-8">
  </head>
  <body>
    <form method="POST" action="/files" enctype="multipart/form-data">
      <input type="file" name="file">
      <input type="submit">
    </form>
  </body>
</html>`

// InputFormHandler response input form HTML
func InputFormHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, formHTML)
}

type data struct {
	Index    int    `json:"index"`
	FileName string `json:"file_name"`
}

type response []data

// UploadFilesHandler is for uploading files Handler
func UploadFilesHandler(c echo.Context) error {
	r := c.Request()
	reader, err := r.MultipartReader()

	if err != nil {
		c.Echo().Logger.Errorf("Error: failed to create multipart reader: %v", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	if reader == nil {
		c.Echo().Logger.Error("Error: not a multipart request")
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	var index int
	var response response

	for {
		part, err := reader.NextPart()

		if err == io.EOF {
			c.Echo().Logger.Info("succeed to read multipart files: EOF")
			break
		}
		if err != nil {
			c.Echo().Logger.Errorf("Error: failed to read next part: %v", err.Error())
			return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
		}

		fileName := part.FileName()

		_, format, err := image.DecodeConfig(part)
		if err != nil {
			c.Echo().Logger.Errorf("Error: failed to decode image format: %v", err.Error())
			return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
		}
		c.Echo().Logger.Info("image format is " + format)

		// 画像を使ったなんらかの処理

		r := data{
			Index:    index,
			FileName: fileName,
		}
		response = append(response, r)
		index++
	}
	return c.JSON(http.StatusOK, response)
}

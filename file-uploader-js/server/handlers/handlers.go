package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"main/models"

	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("Error", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}
	fileMap := form.File
	if len(fileMap) == 0 {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	file := fileMap["file"][0]
	fileName := file.Filename

	reader, err := file.Open()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	buf := make([]byte, file.Size)
	_, err = reader.Read(buf)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	newName := "new_" + fileName
	fo, err := os.Create(newName)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	if _, err := fo.Write(buf); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(models.UploadQueryParams{Name: newName})
}

func Download(c *fiber.Ctx) error {
	params := models.DownloadArgs{}
	if err := c.QueryParser(&params); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if params.Path == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	fmt.Printf("Download params: %v\n", params)

	buf, err := os.ReadFile(params.Path)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusNotFound)
	}

	contentTypeFull := http.DetectContentType(buf)
	fmt.Println(contentTypeFull)

	contentType := strings.Split(contentTypeFull, " ")[0]
	if contentType[len(contentType)-1] == ';' {
		contentType = contentType[:len(contentType)-1]
	}

	//base64Buf := make([]byte, base64.StdEncoding.EncodedLen(len(buf)))
	//
	//base64.StdEncoding.Encode(base64Buf, buf)

	return c.Status(fiber.StatusOK).JSON(models.DownloadResponse{
		ContentType: contentType,
		File:        buf,
	})
}

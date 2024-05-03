package handlers

import (
	"fmt"
	"os"

	"main/models"

	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Println("Error", err)
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
		panic(err)
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
	return c.Status(fiber.StatusOK).SendString(params.Path+"__downloaded")
}

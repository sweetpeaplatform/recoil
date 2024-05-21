package main

import (
	//"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strings"
)

func CreateColumn(tableName, columnName, columnType string) map[string]string {
	tblJSON := map[string]string{
		"Table Name":  tableName,
		"Column Name": columnName,
		"Column Type": columnType,
	}
	return tblJSON
}

func main() {
	app := fiber.New()
	app.Get("*", func(c *fiber.Ctx) error {
		b := strings.Split(c.Path(), "/")
		log.Println(b)
		return c.JSON(CreateColumn(b[1], b[2], b[3]))
	})
	log.Fatal(app.Listen(":3000"))
}

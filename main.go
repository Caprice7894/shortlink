package main
import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"math/rand"
)
type LinkBody struct {
    Link string `json:"link" form:"link" xml:"link"`
		Base string `json:"base" form:"base" xml:"base"`
}
func main() {
	app := fiber.New()
	app.Use(logger.New())
	// genera un map de strings
	MapaEnlaces := make(map[string]string)
	app.Get("/:shortlink", func(c *fiber.Ctx) error {
		shortlink := c.Params("shortlink")

		link, found := MapaEnlaces[shortlink]

		if !found {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "shortlink not found", "data": nil})
		}
		return c.Redirect(link, 301)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		link := new(LinkBody)
		c.BodyParser(link)
		//shortlink is a random string of 5 characters
		shortlink := RandStringRunes(16)
		MapaEnlaces[shortlink] = link.Link
		base := link.Base
		if base == "" {
			base = c.Protocol() + "://" + c.Hostname() + "/"
		}
		return c.JSON(fiber.Map{
			"shortlink": base + shortlink,
		})
	})
	app.Listen(":3003")
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")


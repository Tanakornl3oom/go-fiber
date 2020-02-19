package main

import ("github.com/gofiber/fiber"
        "go-fiber/src/db"
        )


func main() {
  // Create new Fiber instance:
  app := fiber.New()

  db.ConnectDB()
  // Create route on root path, "/":
  app.Get("/", func(c *fiber.Ctx) {
    c.Send("Hello, World!")
    // => "Hello, World!"
  })

  app.Get("/value/:value?", func(c *fiber.Ctx) {
    if c.Params("value") != "" {
      c.Send("Get request with value: " + c.Params("value"))
      // => Get request with value: hello world
      return
    }
  
    c.Send("Get request without value")
  
    })
  
  app.Get("/:value", func(c *fiber.Ctx) {
	c.Send("Get request with value: " + c.Params("value"))
	// => Get request with value: hello world
  })

  // Start server on "localhost" with port "8080":
  app.Listen(8080)
}
package main

import ("github.com/gofiber/fiber"
        "go-fiber/src/db"
        "strconv"
        "encoding/json"
        "go-fiber/src/middleware"
        )


func main() {
  // Create new Fiber instance:
  app := fiber.New()
  app.Use(middleware.Cors())

  dbConn := db.ConnectDB()
  // Create route on root path, "/":
  app.Get("/", func(c *fiber.Ctx) {

    c.Send("Hello, World!")
    // => "Hello, World!"
  })

  app.Get("/users",func(c *fiber.Ctx) {

    var userList = db.GetUsers(dbConn)
    c.Send(userList)
  })

  app.Get("/user/:id",func(c *fiber.Ctx) {

    if c.Params("id") == ""  {
      c.Send("required id")
    }else{
      id, e := strconv.Atoi(c.Params("id"))
      if e != nil {
        c.Send("id should be number")
      }else{
        var user = db.GetUserById(dbConn,id)
        if user.ID == 0 {
          c.Send("user id : "+strconv.Itoa(id)+" not found ")
        }else{
          c.Send(db.PrettyPrint(user))
        }
      }
    }
  })

  app.Post("/user",func(c *fiber.Ctx) {

    var body db.User
    json.Unmarshal([]byte(c.Body()), &body)
    if body.Firstname == "" || body.Lastname == "" || body.Age == "" {
      c.Send("required firstName , lastName and age")
    } else{ 
      age, e := strconv.Atoi(body.Age)
      if e != nil {
        c.Send("age should be number")
      }else{

        var user = db.UserDB{
          Firstname : body.Firstname,
          Lastname :body.Lastname,
          Age : age}
    
        var status = db.AddUser(dbConn,user)
        if status {
          c.Send("Success add user : "+db.PrettyPrint(user))
        }else {
          c.Send("Fail add user")
        }
      }
    }
  
  })


  app.Put("/user/:id",func(c *fiber.Ctx) {

    var body db.User
    json.Unmarshal([]byte(c.Body()), &body)

    if c.Params("id") == ""  {  // require id 
      c.Send("required id")
    }else if body.Firstname == "" || body.Lastname == "" || body.Age == "" { // require use body
      c.Send("required firstName , lastName and age")
    }else{
      // check age integer
      age, e := strconv.Atoi(body.Age)
      if e != nil {
        c.Send("age should be number")
      }else{
        //check id integer
        id, e := strconv.Atoi(c.Params("id"))
        if e != nil {
          c.Send("id should be number")
        }else{
          //check user exist
          var user = db.GetUserById(dbConn,id)
          if user.ID == 0 {
            c.Send("user id : "+strconv.Itoa(id)+" not found ")
          }else{
              //edit user
              user.Firstname =  body.Firstname
              user.Lastname =  body.Lastname
              user.Age = age

              //func edit user
              var editUserResponse db.Response 
              editUserResponse = db.UpdateUser(dbConn,user)
              c.Send(editUserResponse.Response)
            }
          }
      }
    }
  })



  app.Delete("/user/:id",func(c *fiber.Ctx) {
    if c.Params("id") == ""  {  // require id 
      c.Send("required id")
    }else{
      if c.Params("id") == ""  {
        c.Send("required id")
      }else{
        id, e := strconv.Atoi(c.Params("id"))
        if e != nil {
          c.Send("id should be number")
        }else{
            var user = db.GetUserById(dbConn,id)
            if user.ID == 0 {
              c.Send("user id : "+strconv.Itoa(id)+" not found ")
            }else{
              var editUserResponse db.Response 
              editUserResponse = db.DeleteUser(dbConn,id)
              c.Send(editUserResponse.Response)
            }
        }
    }
  }
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
  app.Listen(8084)
}
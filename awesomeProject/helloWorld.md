## Hello World in GO 
![img.png](img.png)
*I have started learning about teh possibility of using GO for **Magic Eagle** simulator project but not just that, i will be adding some GO code in order to learn new things and see if its possible to integrated with the actual stack *

> This Hello World builts the most basic http response to 3000 that yo can built with GO

> #### Testing FIBER part 1
>
> - Easy (Inspired in Express)
> - Really fast (because its GO :) not like express lol

Find the following code on the main package:
        
    package main
    
    import (
    "github.com/gofiber/fiber/v2"
    _ "github.com/gofiber/fiber/v2"
    "strconv"
    _ "strconv"
    )
    
    type coche struct {
    Matricula int
    Modelo    string
    Km        int
    }
    
    func aumentarKM(c *coche) {
    for i := 1; i < 10; i++ {
    c.Km = i
    }
    }
    
    func main() {
    audi := coche{5678, "RS6", 8}
    aumentarKM(&audi)
    
    app := fiber.New()
    
    app.Get("/", func(f *fiber.Ctx) error {
    if audi.Modelo != "RS7" {
                return fiber.ErrForbidden
    } else {
                return f.SendString("Hello, World! " + audi.Modelo + " KM: " + strconv.Itoa(audi.Km))
    }
    })
    
    app.Listen(":3000")
    }


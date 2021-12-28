package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
  	e := echo.New()

  	e.POST("/users", func(c echo.Context) error {


		name := c.QueryParam("name")
		email := c.QueryParam("email")
		fmt.Printf("%s ||| %s \n",name,email)

		u := new(User)
		if err := c.Bind(u); err != nil {
			fmt.Println("Error: ", err)
			return err
		}

		//No output here
		fmt.Printf("%s ---%s \n",u.Name,u.Email)
		return c.JSON(http.StatusCreated, u)
	
	})
	e.Logger.Fatal(e.Start(":1324"))
}

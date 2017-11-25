package main

import (
	echo2 "github.com/labstack/echo"
	"net/http"
	"os"
	"io"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo2.New()
	e.GET("/", func(context echo2.Context) error {
		return context.String(http.StatusOK, "hello echo")
	})
	//路径中使用？带参数的访问(还有一种是路径作为参数的)
	e.GET("/show", show)
	/*
		接受POST的数据，同样也包括接受form中提交的数据。name为字符串，avatar为上传的图片
		curl -F "name=Joe Smith" -F "avatar=@/E:\\/avatar.png" http://localhost:1323/save
	*/
	e.POST("/save", save)
	/*
		curl -H "Content-Type: application/json" -d '{"name":"sonx","email":"winner47@163.com"}' http://localhost:1323/users
		使用绑定函数，使得请求中的json、xml、form表单数据、query中的参数，自动绑定到已定义好的数据结构中
	*/
	e.POST("/users", users)


	e.Logger.Fatal(e.Start(":1323"))
}

func show(c echo2.Context) error {
	name := c.QueryParam("name")
	return c.String(http.StatusOK, name)
}

func users(c echo2.Context) error  {
	u := new(User)
	if err := c.Bind(u); err != nil{
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func save(c echo2.Context) error {
	// Get name
	name := c.FormValue("name")
	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! " + name + "</b>")
}

package main

import (
	echo2 "github.com/labstack/echo"
	"net/http"
	"os"
	"io"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo2.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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

	// Login route
	e.POST("/login", login)

	// Unauthenticated route
	e.GET("/", accessible)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)


	e.Logger.Fatal(e.Start(":1323"))
}

func login(c echo2.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "jon" && password == "shhh!" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jon Snow"
		claims["admin"] = true
		claims["auth"] = "sonx"
		claims["exp"] = time.Now().Add(time.Second * 30).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo2.ErrUnauthorized
}

func accessible(c echo2.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo2.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"! your atuh is " + claims["auth"].(string))
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
	//return c.XML(http.StatusOK, u)
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

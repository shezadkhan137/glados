package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fedesog/webdriver"
	"github.com/gin-gonic/gin"
)

func getCookies() (cookies []webdriver.Cookie) {
	chromeDriver := webdriver.NewChromeDriver("/usr/local/bin/chromedriver")
	err := chromeDriver.Start()
	if err != nil {
		log.Println(err)
	}

	defer chromeDriver.Stop()

	desired := webdriver.Capabilities{"Platform": "OSX"}
	required := webdriver.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)
	if err != nil {
		log.Println(err)
	}

	defer session.Delete()

	err = session.Url("https://lmet.aiwip.com/accounts/login/")

	if err != nil {
		log.Println(err)
	}

	passwordEl, err := session.FindElement(webdriver.CSS_Selector, "#id_password")
	if err != nil {
		panic(err)
	}

	passwordEl.SendKeys("")

	usernameEl, err := session.FindElement(webdriver.CSS_Selector, "#id_username")
	if err != nil {
		panic(err)
	}

	usernameEl.SendKeys("shezad@letolab.com")
	if err != nil {
		panic(err)
	}

	err = usernameEl.Submit()
	if err != nil {
		panic(err)
	}

	cookies, err = session.GetCookies()
	if err != nil {
		panic(err)
	}

	time.Sleep(5 * time.Second)
	fmt.Printf("Cookies: %+v", cookies)
	return cookies
}

func main() {
	r := gin.Default()
	r.GET("/cookies", func(c *gin.Context) {
		cookies := getCookies()
		c.JSON(200, cookies)
	})
	r.Run(":8080")
}

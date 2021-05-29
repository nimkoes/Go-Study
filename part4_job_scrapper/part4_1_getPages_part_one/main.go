package main

import (
	"learngo/part4_job_scrapper/part4_1_getPages_part_one/myScrap"
	"os"
	"strings"

	"github.com/labstack/echo"
)

const FILE_NAME string = "jobs.csv"

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/myScrap", handleScrap)
	e.Logger.Fatal(e.Start(":1323"))
}

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrap(c echo.Context) error {

	// 실행 이후 서버에서 파일을 삭제한다.
	defer os.Remove(FILE_NAME)

	term := strings.ToLower(myScrap.CleanString(c.FormValue("term")))
	myScrap.MyScrap(term)
	return c.Attachment(FILE_NAME, FILE_NAME)
}

package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	getPages()
}

func getPages() int {
	res, err := http.Get(baseURL)

	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	/*
	 * res.body 는 기본적으로 byte 데이터 IO 를 하기 때문에
	 * 사용한 다음에 자원 해제 (close) 해줘야 한다.
	  * 메모리 누수를 막기 위해 defer 구문을 사용해서 닫아준다.
	*/
	doc, err := goquery.NewDocumentFromReader(res.Body)

	checkErr(err)

	doc.Find(".pagination")

	return 0
}

// err 가 있으면 프로그램 종료
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// response status 가 200 (정상) 이 아니면 프로그램 종료
func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.Status)
	}
}

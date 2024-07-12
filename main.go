package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func multiply(a int, b int) int {
	return a * b
}

func variableArgs(str ...string) {
	defer fmt.Println("is done!")
	fmt.Println(str)
}

type result struct {
	url    string
	status string
}

func main() {
	// variableArgs("gwak", "hee", "won")
	// variableArgs("gwakheewon")

	// a := 2
	// b := &a
	// a = 10
	// fmt.Println(a, *b)

	// name := [5]string{"gwak", "heewon", "won"}
	// name[3] = "천"
	// name[4] = "재"
	// fmt.Println(name)

	// slice := []string{"gwak", "heewon", "won"}
	// slice = append(slice, "천")
	// slice = append(slice, "재")
	// slice = append(slice, "다")
	// fmt.Println(slice)

	// map_test := map[string]string{"name": "heewon", "age": "23"}
	// for key, value := range map_test {
	// 	fmt.Println(key, value)
	// }

	// type person struct {
	// 	name    string
	// 	age     int
	// 	favFood []string
	// }

	// favFood := []string{"bread", "chicken"}
	// heewon := person{
	// 	name:    "heewon",
	// 	age:     23,
	// 	favFood: favFood,
	// }
	// fmt.Println(heewon)

	// account := banking.NewAccount("heewon")
	// account.Deposit(10)
	// err := account.Withdraw(11)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println(account.Balance())

	// dict := mydict.Dictionary{}

	// err := dict.Add("key", "value")
	// if err == nil {
	// 	fmt.Println(dict)
	// }

	// err2 := dict.Delete("key")
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	// v, err3 := dict.Search("key")

	// if err3 == nil {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println(err3)
	// }

	// results := make(map[string]string)
	// urls := []string{
	// 	"https://www.airbnb.com/",
	// 	"https://www.google.com/",
	// 	"https://www.amazon.com/",
	// 	"https://www.reddit.com/",
	// 	"https://www.goo2gle.com/",
	// 	"https://soundcloud.com/",
	// 	"https://www.fac2ebook.com/",
	// 	"https://www.instagram.com/",
	// 	"https://academy.nomadcoders.com/",
	// }

	// c := make(chan result)

	// for _, url := range urls {
	// 	go hitUrl(url, c)
	// }

	// for result := range c {
	// 	results[result.url] = result.status
	// }

	// for url, result := range results {
	// 	fmt.Println(url, result)
	// }

	// 치이카와 재입고 페이지에서 재입고된 상품명 가격 품절여부 가져오기
	var baseUrl string = "https://chiikawamarket.jp/collections/restock"
	ps, err := getPageListInfo(baseUrl)
	if err == nil {
		fmt.Println(ps)
	}
}

// func getTotalPageCount(url string) int {
// 	res, err := http.Get(url)
// 	checkErr(err)
// 	checkStatusCode(res)

// 	defer res.Body.Close()

// 	doc, err := goquery.NewDocumentFromReader(res.Body)
// 	checkErr(err)

// 	doc.Find(".pagination--number").Each(func(i int, s *goquery.Selection) {
// 		fmt.Println(s.Html())
// 	})

// 	return 0
// }

type ProductInfo struct {
	name     string
	price    string
	hasStock bool
}

func getPageListInfo(url string) ([]ProductInfo, error) {
	ps := []ProductInfo{}
	c := make(chan []ProductInfo)

	for i := 1; i <= 3; i++ {
		go getPageInfo(url+"?page="+strconv.Itoa(i), c)
	}

	for i := 0; i < 3; i++ {
		p := <-c
		fmt.Println(p)
		ps = append(ps, p...)
	}

	return ps, nil
}

func getPageInfo(url string, c chan<- []ProductInfo) error {
	temp := []ProductInfo{}
	res, err := http.Get(url)
	checkErr(err)
	checkStatusCode(res)
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return err
	}

	selection := doc.Find(".product--root")
	if selection.Length() == 0 {
		return errors.New("no page")
	} else {
		selection.Each(func(i int, s *goquery.Selection) {
			temp = append(temp, getProductInfo(s))
		})
	}
	c <- temp
	defer res.Body.Close()

	return nil
}

func getProductInfo(s *goquery.Selection) ProductInfo {
	name := s.Find(".product_name").Text()
	price := s.Find(".product_price").Text()
	label := s.Find(".product--label").Text()
	hasStock := true
	if label == "売り切れ" {
		hasStock = false
	}
	return ProductInfo{
		name:     strings.TrimSpace(name),
		price:    strings.TrimSpace(price),
		hasStock: hasStock,
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStatusCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("failed")
	}
}

func hitUrl(url string, c chan<- result) {
	fmt.Println("Checking.. ", url)
	status := "OK"
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		status = "Failed"
	}
	c <- result{
		url:    url,
		status: status,
	}
}

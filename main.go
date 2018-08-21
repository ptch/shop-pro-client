package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Sales 売上情報
type Sales struct {
	Sales []Sale `json:"sales"`
}

// Sale 売上情報個別
type Sale struct {
	ProductTotalPrice int `json:"product_total_price"`
}

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.shop-pro.jp/v1/sales.json?paid=true", nil)
	req.Header.Add("Authorization", "")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("%#v\n", resp)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(string(body))

	sales := Sales{}

	err = json.Unmarshal(body, &sales)
	if err != nil {
		fmt.Println(err)
	}

	sum := 0
	for _, f := range sales.Sales {
		sum = sum + f.ProductTotalPrice
	}

	fmt.Println(sum)

}

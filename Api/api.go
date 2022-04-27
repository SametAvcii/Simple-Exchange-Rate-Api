package Api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Result               string
	Time_last_update_utc string
	Base_code            string
	Conversion_rates     Conversion_rates
}

type Conversion_rates struct {
	Usd float32
	Try float32
	Eur float32
}

func GetRequest(url string) []byte {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response Response
	json.Unmarshal(body, &response)

	fmt.Println("USD/TRY", response.Conversion_rates.Try)
	fmt.Println("USD/EUR", response.Conversion_rates.Eur)
	fmt.Println("TRY/EUR", (response.Conversion_rates.Try)/(response.Conversion_rates.Eur))

	a := response.Conversion_rates.Try
	b := response.Conversion_rates.Eur
	c := ((response.Conversion_rates.Try) / (response.Conversion_rates.Eur))
	StrArr := []byte{byte(a), byte(b), byte(c)}
	return StrArr

}

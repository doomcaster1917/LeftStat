package FedStatQueryHandler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var years = []string{"2012", "2013", "2014", "2015", "2016", "2017", "2018", "2019", "2020", "2021", "2022", "2023", "2024", "2025"}

var URLRequest = "https://www.fedstat.ru/indicator/dataGrid.do?id=" + StatisticIds["minimalProductsPriceBasket"] +
	"&lineObjectIds=" + StatisticIds["name"] +
	"&lineObjectIds=" + Currencies["ruble"] +
	"&lineObjectIds=" + Regions["Russian Federation"] +
	"&columnObjectIds=3" +
	"&selectedFilterIds=3_2012" +
	"&selectedFilterIds=3_2013" +
	"&selectedFilterIds=3_2014" +
	"&selectedFilterIds=3_2015" +
	"&selectedFilterIds=3_2016" +
	"&selectedFilterIds=3_2017" +
	"&selectedFilterIds=3_2018" +
	"&selectedFilterIds=3_2019" +
	"&selectedFilterIds=3_2020" +
	"&selectedFilterIds=3_2021" +
	"&selectedFilterIds=3_2022" +
	"&selectedFilterIds=3_2023"
var URLS = map[string]string{
	"requestURL": "https://www.fedstat.ru/indicator/dataGrid.do?id=",
	"URL":        "https://www.fedstat.ru/indicator/dataGrid.do", // Пары ключ-значения являются композитными литералами для карт
}

var Months = map[string]string{
	"jan": "1540283", "feb": "1540282", "mar": "1540236",
	"apr": "1540229", "may": "1540235", "jun": "1540234",
	"jul": "1540233", "aug": "1540228", "sen": "1540276",
	"oct": "1540273", "nov": "1540272", "dec": "1540230",
}

var Regions = map[string]string{
	"Russian Federation": "57831",
}

var StatisticIds = map[string]string{
	"name":                       "0",
	"minimalProductsPriceBasket": "31481",
}

var Currencies = map[string]string{
	"ruble": "30611",
}

const FullDates = "33560"

type MinimalConsumerBasketResponse struct {
	Results []map[string]string
	Count   string `json:"__count"`
}

var testquery = "https://www.fedstat.ru/indicator/dataGrid.do?id=31481&lineObjectIds=0&lineObjectIds=30611&lineObjectIds=57831&columnObjectIds=3&selectedFilterIds=3_2012&selectedFilterIds=3_2013&selectedFilterIds=3_2014&selectedFilterIds=3_2015&selectedFilterIds=3_2016&selectedFilterIds=3_2017&selectedFilterIds=3_2018&selectedFilterIds=3_2019&selectedFilterIds=3_2020&selectedFilterIds=3_2021&selectedFilterIds=3_2022&selectedFilterIds=3_2023&selectedFilterIds=3_2024\n"

func GetStatData(query string) string {
	j, _ := json.Marshal(parseResults(query))
	fmt.Println(string(j))
	return string(j)
}

func parseResults(query string) map[string]int {
	ResultString := getResults(query)
	values := make(map[string]int)
	for i := range ResultString[0] {
		for year := range years {
			if strings.Contains(i, "dim"+years[year]) {
				values[years[year]] = convertResponse(ResultString[0][i])
			}
		}
	}
	return values
}

func getResults(query string) []map[string]string {
	var r1 MinimalConsumerBasketResponse

	println(query)
	resp, err := http.Get(URLRequest)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	err = json.Unmarshal(body, &r1)
	if err != nil {
		log.Fatal(err)
	}

	return r1.Results
}

func convertResponse(str string) int {
	filtredString := strings.ReplaceAll(str, ",", ".")
	converted, err := strconv.ParseFloat(filtredString, 16)
	if err != nil {
		log.Fatal(err)
	}
	return int(converted)
}

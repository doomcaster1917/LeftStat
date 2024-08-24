//package CommunistsStatistic

//type URLS struct {
//	RequestURL string
//}

//type Months struct {
//	Jan uint32
//	Feb uint32
//	Mar uint32
//	Apr uint32
//	May uint32
//	Jun uint32
//	Jul uint32
//	Aug uint32
//	Sen uint32
//	Oct uint32
//	Nov uint32
//	Dec uint32
//}

//var MonthsCodes = Months{Jan: 1540283, Feb: 1540282, Mar: 1540236,
//	Apr: 1540229, May: 1540235, Jun: 1540234,
//	Jul: 1540233, Aug: 1540228, Sen: 1540276,
//	Oct: 1540273, Nov: 1540272, Dec: 1540230}

//type Regions struct {
//	RussianFederation uint16
//}

//type StatisticIds struct {
//	MinimalProductsPriceBasket string
//}

//var Ids = StatisticIds{MinimalProductsPriceBasket: "dim0"}

//type Currencies struct {
//	Ruble uint32
//}
// https://www.fedstat.ru/indicator/dataGrid.do?id=31481
//var URL = URLS{RequestURL: "https://www.fedstat.ru/indicator/dataGrid.do?id="}

//var Cur = Currencies{Ruble: 30611}

//lineObjectIds: 0
//lineObjectIds: 30611
//lineObjectIds: 57831
//columnObjectIds: 3
//columnObjectIds: 33560
//selectedFilterIds: 0_31481
//selectedFilterIds: 3_2012
//selectedFilterIds: 3_2013
//selectedFilterIds: 3_2014
//selectedFilterIds: 3_2015
//selectedFilterIds: 3_2016
//selectedFilterIds: 3_2017
//selectedFilterIds: 3_2018
//selectedFilterIds: 3_2019
//selectedFilterIds: 3_2020
//selectedFilterIds: 3_2021
//selectedFilterIds: 3_2022
//selectedFilterIds: 3_2023
//selectedFilterIds: 30611_950351
//selectedFilterIds: 33560_1540228
//selectedFilterIds: 33560_1540229
//selectedFilterIds: 33560_1540230
//selectedFilterIds: 33560_1540233
//selectedFilterIds: 33560_1540234
//selectedFilterIds: 33560_1540235
//selectedFilterIds: 33560_1540236
//selectedFilterIds: 33560_1540272
//selectedFilterIds: 33560_1540273
//selectedFilterIds: 33560_1540276
//selectedFilterIds: 33560_1540282
//selectedFilterIds: 33560_1540283
//selectedFilterIds: 57831_1688487

// "https://showdata.gks.ru/x/report/274004/view/compound/?filter_1_0=2014-01-01+00%3A00%3A00%7C-56&rp_submit=t&_=1707688460439"

//var years = []int{2012, 2013, 2014, 2015, 2016, 2017, 2018, 2019, 2020, 2021, 2022, 2023, 2024}
//
//var URLRequest = "https://www.fedstat.ru/indicator/dataGrid.do?id=" + StatisticIds["minimalProductsPriceBasket"] +
//	"&lineObjectIds=" + StatisticIds["name"] +
//	"&lineObjectIds=" + Currencies["ruble"] +
//	"&lineObjectIds=" + Regions["Russian Federation"] +
//	"&columnObjectIds=3" +
//	"&selectedFilterIds=3_2012" +
//	"&selectedFilterIds=3_2013" +
//	"&selectedFilterIds=3_2014" +
//	"&selectedFilterIds=3_2015" +
//	"&selectedFilterIds=3_2016" +
//	"&selectedFilterIds=3_2017" +
//	"&selectedFilterIds=3_2018" +
//	"&selectedFilterIds=3_2019" +
//	"&selectedFilterIds=3_2020" +
//	"&selectedFilterIds=3_2021" +
//	"&selectedFilterIds=3_2022" +
//	"&selectedFilterIds=3_2023"
//var URLS = map[string]string{
//	"requestURL": "https://www.fedstat.ru/indicator/dataGrid.do?id=",
//	"URL":        "https://www.fedstat.ru/indicator/dataGrid.do", // Пары ключ-значения являются композитными литералами для карт
//}
//
//var Months = map[string]string{
//	"jan": "1540283", "feb": "1540282", "mar": "1540236",
//	"apr": "1540229", "may": "1540235", "jun": "1540234",
//	"jul": "1540233", "aug": "1540228", "sen": "1540276",
//	"oct": "1540273", "nov": "1540272", "dec": "1540230",
//}
//
//var Regions = map[string]string{
//	"Russian Federation": "57831",
//}
//
//var StatisticIds = map[string]string{
//	"name":                       "0",
//	"minimalProductsPriceBasket": "31481",
//}
//
//var Currencies = map[string]string{
//	"ruble": "30611",
//}
//
//const FullDates = "33560"

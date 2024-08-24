package main

//
//import (
//	chartsImages "CommunistsStatistic/charts_img_maker"
//	encharts "CommunistsStatistic/encharts_maker"
//	"html/template"
//	"net/http"
//	"os"
//)
//
//type charts struct {
//	enchart    template.HTML
//	chartImage *os.File
//}
//
//// Глобальный секретный ключ
//
//const hmacSampleSecret = "hmacSampleSecret"
//
//func (app *application) home(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path != "/" {
//		app.notFound(w) // Использование помощника notFound()
//		return
//	}
//
//	files := []string{
//		"./ui/html/home.page.tmpl",
//		"./ui/html/base.layout.tmpl",
//		"./ui/html/footer.partial.tmpl",
//		"./ui/html/header.partial.tmpl",
//	}
//
//	ts, err := template.ParseFiles(files...)
//	if err != nil {
//		app.serverError(w, err) // Использование помощника serverError()
//		return
//	}
//
//	dataForImageMaker := chartsImages.Chart{
//		XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
//		YValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
//	}
//
//	dataCharts := charts{
//		encharts.encharts_maker(),
//		dataForImageMaker.MakeChartImg(),
//	}
//	w.Header().Set("Content-Type", "image/png")
//	err = ts.Execute(w, dataCharts)
//}
//
//func (app *application) handler(w http.ResponseWriter, r *http.Request) {
//
//}

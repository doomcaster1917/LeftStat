package main

import (
	"CommunistsStatistic/pkg/middleware"
	"fmt"
	"net/http"
	"strings"
)

func (app *application) getViews(res http.ResponseWriter, req *http.Request) {
	response := setHeaders(res)
	fmt.Fprintf(response, "[%s]", strings.Join(middleware.GetAllViews(), `,`))
}

func (app *application) getView(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	data := middleware.GetView(id)
	response := setHeaders(res)
	response.Write(data)
}

func setHeaders(res http.ResponseWriter) http.ResponseWriter {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", fmt.Sprintf("%v", frontendAddr))
	return res
}

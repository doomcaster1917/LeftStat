package main

import (
	"CommunistsStatistic/middleware"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Login struct {
	token string
}

type WithAuth struct {
	auth string
}

type HeaderConfig interface {
	SetHeaders()
}

func (app *application) adminPanel(res http.ResponseWriter, req *http.Request) {
	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err := middleware.Auth(token)

	if err != nil {
		fmt.Println(err)
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		HeaderConfig.setHeaders(res)
	}
}

func (app *application) loginInAdmin(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	token, err := middleware.Credentials{req.Form.Get("name"), req.Form.Get("password")}.Login()
	if err != nil {
		fmt.Println(err)
		app.notFound(res)
	}
	HeaderConfig := Login{token: fmt.Sprintf("Bearer %v", token)}
	HeaderConfig.setHeaders(res)
}

func (app *application) GetCharts(res http.ResponseWriter, req *http.Request) {
	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err := middleware.Auth(token)
	if err != nil {
		fmt.Println(err)
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		fmt.Fprintf(response, "[%s]", strings.Join(middleware.GetCharts(), `,`)) //)
	}
}

func (app *application) GetChart(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err = middleware.Auth(token)
	if err != nil {
		fmt.Println(err)
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		if id == 0 {
		} else {
			fmt.Fprintf(response, "%s", middleware.GetChart(id)) //)
		}
	}
}

func (app *application) GetDataset(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err = middleware.Auth(token)
	if err != nil {
		fmt.Println(err)
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		fmt.Fprintf(response, "[%s]", strings.Join(middleware.GetDataset(id), `,`)) //)
	}
}

func (app *application) UpdateDataset(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		println(err)
	}

	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err = middleware.Auth(token)
	if err != nil {
		fmt.Println(err)
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		id, _ := strconv.Atoi(req.Form.Get("id"))
		err := middleware.UpdateDataset(req.Form.Get("sendMode"), req.Form.Get("name"),
			req.Form.Get("raw"), id)
		if err != nil {
			http.Error(res, "Ошибка", 400)
		} else {
			fmt.Fprintf(response, "success")
		}
	}
}

func (app *application) SelectDatasets(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err = middleware.Auth(token)

	fmt.Println(req.Form.Get("selected_datasets"))
	if err != nil {
		fmt.Println(err)
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		err := middleware.SelectDatasets(req.Form.Get("selected_datasets"), req.Form.Get("id"))
		if err != nil {
			http.Error(res, "Ошибка", 400)
		} else {
			fmt.Fprintf(response, "success")
		}
	}
}

func (app *application) CreateChart(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err = middleware.Auth(token)

	if err != nil {
		fmt.Println(err)
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		err := middleware.CreateChart(req.Form.Get("name"), req.Form.Get("title"))
		if err != nil {
			http.Error(res, "Ошибка", 400)
		} else {
			fmt.Fprintf(response, "success")
		}
	}
}

func (app *application) SetAxis(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err = middleware.Auth(token)

	if err != nil {
		fmt.Println(err)
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		err := middleware.SetAxis(req.Form.Get("dataset_id"), req.Form.Get("chart_id"))
		if err != nil {
			http.Error(res, "Ошибка", 400)
		} else {
			fmt.Fprintf(response, "success")
		}
	}
}

func (app *application) UpdateChart(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err = middleware.Auth(token)

	if err != nil {
		fmt.Println(err)
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		err := middleware.UpdateChart(req.Form.Get("name"), req.Form.Get("title"), req.Form.Get("chart_id"))
		if err != nil {
			http.Error(res, "Ошибка", 400)
		} else {
			fmt.Fprintf(response, "success")
		}
	}
}

func (l Login) setHeaders(res http.ResponseWriter) {
	res.Header().Set("Access-Control-Allow-Credentials", "true")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	res.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE, OPTIONS")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Content-Type-Options, access-control-allow-origin, Authorization")
	res.Header().Set("Access-Control-Expose-Headers", "Authorization")
	res.Header().Set("Authorization", l.token)
}

func (auth WithAuth) setHeaders(res http.ResponseWriter) http.ResponseWriter {
	res.Header().Set("Access-Control-Allow-Credentials", "true")
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	res.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE, OPTIONS")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Content-Type-Options, access-control-allow-origin, Authorization")
	res.Header().Set("Access-Control-Expose-Headers", "Authorization")
	res.Header().Add("Authorization", auth.auth)
	return res
}

package main

import (
	"CommunistsStatistic/pkg/middleware"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var frontendAddr = "https://leftstat.ru"

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
		app.notFound(res)
	}
	HeaderConfig := Login{token: fmt.Sprintf("Bearer %v", token)}
	HeaderConfig.setHeaders(res)
}

func (app *application) GetCharts(res http.ResponseWriter, req *http.Request) {
	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err := middleware.Auth(token)
	if err != nil {
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
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		if id == 0 {
		} else {
			fmt.Fprintf(response, "%s", middleware.GetChartAdmin(id)) //)
		}
	}
}

func (app *application) GetDataset(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err = middleware.Auth(token)
	if err != nil {
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
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		id, _ := strconv.Atoi(req.Form.Get("id"))
		err := middleware.UpdateDataset(req.Form.Get("sendMode"), req.Form.Get("name"),
			req.Form.Get("raw"), id)
		if err != nil {
			http.Error(res, fmt.Sprintf("%v", err), 400)
		} else {
			fmt.Fprintf(response, "success")
		}
	}
}

func (app *application) SelectDatasets(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
	}

	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err = middleware.Auth(token)

	if err != nil {
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		err := middleware.SelectDatasets(req.Form.Get("selected_datasets"), req.Form.Get("id"))
		if err != nil {
			http.Error(res, fmt.Sprintf("%v", err), 400)
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
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		err := middleware.CreateChart(req.Form.Get("name"), req.Form.Get("title"))
		if err != nil {
			http.Error(res, fmt.Sprintf("%v", err), 400)
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
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		err := middleware.SetAxis(req.Form.Get("dataset_id"), req.Form.Get("chart_id"))
		if err != nil {
			http.Error(res, fmt.Sprintf("%v", err), 400)
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
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		err := middleware.UpdateChart(req.Form.Get("name"), req.Form.Get("title"),
			req.Form.Get("description"), req.Form.Get("chart_id"))
		if err != nil {
			http.Error(res, fmt.Sprintf("%v", err), 400)
		} else {
			fmt.Fprintf(response, "success")
		}
	}
}

func (app *application) GetViews(res http.ResponseWriter, req *http.Request) {
	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err := middleware.Auth(token)
	if err != nil {
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		fmt.Fprintf(response, "[%s]", strings.Join(middleware.GetViews(), `,`)) //)
	}
}

func (app *application) GetView(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err = middleware.Auth(token)
	if err != nil {
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		fmt.Fprintf(response, "%s", middleware.GetViewAdmin(req.Form.Get("id")))
	}
}

func (app *application) UpdateView(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err = middleware.Auth(token)

	if err != nil {
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		err := middleware.UpdateView(req.Form.Get("name"), req.Form.Get("title"),
			req.Form.Get("seoDescription"), req.Form.Get("seoKeywords"), req.Form.Get("description"), req.Form.Get("id"))
		if err != nil {
			http.Error(res, fmt.Sprintf("%v", err), 400)
		} else {
			fmt.Fprintf(response, "success")
		}
	}
}

func (app *application) SetCharts(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
	}

	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err = middleware.Auth(token)

	if err != nil {
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		err := middleware.SetCharts(req.Form.Get("selected_charts"), req.Form.Get("id"))
		if err != nil {
			http.Error(res, fmt.Sprintf("%v", err), 400)
		} else {
			fmt.Fprintf(response, "success")
		}
	}
}

func (app *application) SetMainChart(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err = middleware.Auth(token)

	if err != nil {
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		err := middleware.SetMainChart(req.Form.Get("main_chart_id"), req.Form.Get("view_id"))
		if err != nil {
			http.Error(res, fmt.Sprintf("%v", err), 400)
		} else {
			fmt.Fprintf(response, "success")
		}
	}
}

func (app *application) createView(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err = middleware.Auth(token)

	if err != nil {
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		err := middleware.CreateView(req.Form.Get("name"), req.Form.Get("title"))
		if err != nil {
			http.Error(res, fmt.Sprintf("%v", err), 400)
		} else {
			fmt.Fprintf(response, "success")
		}
	}
}

func (app *application) GetDatasets(res http.ResponseWriter, req *http.Request) {
	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err := middleware.Auth(token)
	if err != nil {
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		fmt.Fprintf(response, "[%s]", strings.Join(middleware.GetDatasets(), `,`)) //)
	}
}

func (app *application) CreateDataset(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err = middleware.Auth(token)

	if err != nil {
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)
		err := middleware.CreateDataset(req.Form.Get("name"))
		if err != nil {
			http.Error(res, fmt.Sprintf("%v", err), 400)
		} else {
			fmt.Fprintf(response, "success")
		}
	}
}

func (app *application) deleteDataset(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err = middleware.Auth(token)
	if err != nil {
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)

		err := middleware.DeleteDataset(req.Form.Get("id"))
		if err != nil {
			http.Error(res, fmt.Sprintf("%v", err), 400)
		} else {
			fmt.Fprintf(response, "success")
		}
	}
}

func (app *application) deleteChart(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err = middleware.Auth(token)
	if err != nil {
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)

		err := middleware.DeleteChart(req.Form.Get("id"))
		if err != nil {
			http.Error(res, fmt.Sprintf("%v", err), 400)
		} else {
			fmt.Fprintf(response, "success")
		}
	}
}

func (app *application) deleteView(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err = middleware.Auth(token)
	if err != nil {
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)

		err := middleware.DeleteView(req.Form.Get("id"))
		if err != nil {
			http.Error(res, fmt.Sprintf("%v", err), 400)
		} else {
			fmt.Fprintf(response, "success")
		}
	}
}

func (app *application) createImg(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	token := strings.ReplaceAll(req.Header.Get("Authorization"), "Bearer ", "")
	err = middleware.Auth(token)
	if err != nil {
		HeaderConfig := WithAuth{auth: "false"}
		HeaderConfig.setHeaders(res)
	} else {
		HeaderConfig := WithAuth{auth: "true"}
		response := HeaderConfig.setHeaders(res)

		err := middleware.CreateImg(req.Form.Get("id"))
		if err != nil {
			http.Error(res, fmt.Sprintf("%v", err), 400)
		} else {
			fmt.Fprintf(response, "success")
		}
	}
}

func (l Login) setHeaders(res http.ResponseWriter) {
	res.Header().Set("Access-Control-Allow-Credentials", "true")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.Header().Set("Access-Control-Allow-Origin", fmt.Sprintf("%v", frontendAddr))
	res.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE, OPTIONS")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Content-Type-Options, access-control-allow-origin, Authorization")
	res.Header().Set("Access-Control-Expose-Headers", "Authorization")
	res.Header().Set("Authorization", l.token)
}

func (auth WithAuth) setHeaders(res http.ResponseWriter) http.ResponseWriter {
	res.Header().Set("Access-Control-Allow-Credentials", "true")
	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", fmt.Sprintf("%v", frontendAddr))
	res.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE, OPTIONS")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Content-Type-Options, access-control-allow-origin, Authorization")
	res.Header().Set("Access-Control-Expose-Headers", "Authorization")
	res.Header().Add("Authorization", auth.auth)
	return res
}

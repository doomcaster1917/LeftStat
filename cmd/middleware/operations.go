package middleware

import (
	"CommunistsStatistic/cmd/DataBase"
	"fmt"
	"strconv"
)

type Views struct {
	Id      int
	Name    string
	Title   string
	ImgAddr string
}

func GetAllViews() []string {
	return database.GetViews()
}

func GetView(idStr string) string {
	id, _ := strconv.Atoi(idStr)
	view := database.GetView(id)
	fmt.Println(view)
	return view
}

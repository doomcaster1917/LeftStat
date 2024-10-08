package middleware

import (
	"CommunistsStatistic/pkg/DataBase"
	"strconv"
)

func GetAllViews() []string {
	return database.GetViews()
}

func GetView(idStr string) []byte {
	id, _ := strconv.Atoi(idStr)
	return database.GetView(id)
}

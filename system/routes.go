package system

import (
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

var depot = []ToDo{}
var countID = 0

func CreateToDo(ctx *gin.Context) {
	var newToDo ToDo

	if erro := ctx.BindJSON(&newToDo); erro != nil {
		return
	}

	countID++
	newToDo.Id = countID
	depot = append(depot, newToDo)
	ctx.IndentedJSON(http.StatusCreated, newToDo)
}

func ReadToDo(ctx *gin.Context) {
	ctx.IndentedJSON(200, depot)
}

func UpdateToDo(ctx *gin.Context) {
	var putToDo ToDo

	if erro := ctx.BindJSON(&putToDo); erro != nil {
		return
	}

	var target_id, _ = strconv.Atoi(ctx.Param("id"))
	for i, item := range depot {
		if item.Id == target_id {
			log.Println("Check True")
			depot[i].Title = putToDo.Title
			depot[i].Content = putToDo.Content
			depot[i].Status = putToDo.Status
			ctx.IndentedJSON(200, "ok")
			return
		}
	}
	ctx.IndentedJSON(404, "Não realizado")
}

func removeAtIndex(slice []ToDo, index int) []ToDo {
	if index < 0 || index >= len(slice) {
		return slice // Retorna o slice original se o índice for inválido
	}
	return append(slice[:index], slice[index+1:]...)
}

// Concertar o método de removeIndex
func DeleteToDo(ctx *gin.Context) {
	var target_id, _ = strconv.Atoi(ctx.Param("id"))
	for i, item := range depot {
		if item.Id == target_id {
			removeAtIndex(depot, i)
			ctx.IndentedJSON(200, "ok")
			return
		}
	}
	ctx.IndentedJSON(404, "Não efetuado")
}

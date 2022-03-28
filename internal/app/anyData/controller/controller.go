package controller

import (
	"anyData/internal/app/anyData/middleware"
	"anyData/internal/app/anyData/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

const dtFormat = "2006-01-02 15:04:05"

func GetContent(gContext *gin.Context) {
	content, err := middleware.GetContent()
	if err != nil {
		gContext.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if content == nil {
		gContext.IndentedJSON(http.StatusNoContent, content)
		return
	}

	gContext.IndentedJSON(http.StatusOK, content)
}

func GetData(gContext *gin.Context) {
	rowId, ok := gContext.GetQuery("id")
	if !ok {
		gContext.IndentedJSON(http.StatusBadRequest, gin.H{"message": "[GIN-debug] Missing id query param."})
		return
	}
	id, _ := strconv.ParseUint(rowId, 10, 64)

	data, err := middleware.GetDataById(id)
	if err != nil {
		gContext.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if data == nil {
		gContext.IndentedJSON(http.StatusNoContent, data)
		return
	}

	gContext.IndentedJSON(http.StatusOK, data)
}

func AddData(gContext *gin.Context) {
	var newData model.AnyDataStruct
	dateTime := time.Now().Format(dtFormat)
	newData.DateUpdate, newData.DateCreate = dateTime, dateTime
	if err := gContext.BindJSON(&newData); err != nil {
		return
	}
	newId := middleware.AddData(newData)
	if newId > 0 {
		gContext.IndentedJSON(http.StatusCreated, newId)
		return
	}

	gContext.IndentedJSON(http.StatusNotAcceptable, nil)
}

func UpdateData(gContext *gin.Context) {
	var newData model.AnyDataStruct
	newData.DateUpdate = time.Now().Format(dtFormat)
	if err := gContext.BindJSON(&newData); err != nil {
		return
	}

	ok := middleware.UpdateData(newData)
	if ok {
		gContext.IndentedJSON(http.StatusOK, ok)
		return
	}

	gContext.IndentedJSON(http.StatusNotFound, nil)
}

func DeleteData(gContext *gin.Context) {
	strId, _ := gContext.GetQuery("id")
	id, _ := strconv.ParseUint(strId, 10, 64)

	ok := middleware.DeleteData(id)
	if ok {
		gContext.IndentedJSON(http.StatusOK, ok)
		return
	}

	gContext.IndentedJSON(http.StatusNotFound, nil)
}

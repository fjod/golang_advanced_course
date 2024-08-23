package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/fjod/golang_advanced_course/internal"
	data "github.com/fjod/golang_advanced_course/internal/Data"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func HTML(c *gin.Context, storage internal.StorageOperations) {
	t, err := template.New("map").Parse(`
<html>
<body>
<h1>Map Data</h1>
<ul>
{{range $key, $value := .}}
    <li>{{$key}}: {{$value}}</li>
{{end}}
</ul>
</body>
</html>
`)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(c.Writer, storage.Print())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
}

func Get(c *gin.Context, storage internal.StorageOperations) {
	metricType := c.Param("type")
	name := c.Param("name")
	g, err := storage.GetValue(name, metricType)
	if err == nil {
		c.String(http.StatusOK, g)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"err": err})
}

func GetJSON(c *gin.Context, storage internal.StorageOperations) {
	var d data.Metrics
	err := c.ShouldBindJSON(&d)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	g, err := storage.GetJsonValue(d.ID, d.MType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}

	gj, _ := json.Marshal(g)
	c.Data(http.StatusOK, "application/json", gj)
	return
}

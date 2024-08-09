package handlers

import (
	"fmt"
	"github.com/fjod/golang_advanced_course/internal"
	data "github.com/fjod/golang_advanced_course/internal/Data"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
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

func Update(c *gin.Context, storage internal.StorageOperations) {
	metricType := c.Param("type")
	if metricType == "gauge" {
		n, err := strconv.ParseFloat(c.Param("value"), 64)
		if err == nil {
			g := data.Gauge{
				Name: c.Param("name"),
				Val:  n,
			}
			err = internal.SaveMetric(g, storage)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"err": err})
				return
			}
			fmt.Println("приняли gauge ", g.GetName(), " ", g.GetValue())
			c.Status(http.StatusOK)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}

	if metricType == "counter" {
		n, err := strconv.ParseInt(c.Param("value"), 10, 64)
		if err == nil {
			g := data.Counter{
				Name: c.Param("name"),
				Val:  n,
			}
			err = internal.SaveMetric(g, storage)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"err": err})
				return
			}
			fmt.Println("приняли counter ", g.GetName(), " ", g.GetValue())
			c.Status(http.StatusOK)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"err": "unknown type"})
}

package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/fjod/golang_advanced_course/internal"
	data "github.com/fjod/golang_advanced_course/internal/Data"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Update(c *gin.Context, storage internal.StorageOperations) {
	var content = c.ContentType()
	if content == "application/json" {
		jsonUpdate(c, storage)
	} else {
		plainTextUpdate(c, storage)
	}
}

func jsonUpdate(c *gin.Context, storage internal.StorageOperations) {
	var d data.Metrics
	err := c.ShouldBindJSON(&d)
	if err == nil {
		metricType := d.MType
		if metricType == "gauge" {
			g := data.Gauge{
				Name: d.ID,
				Val:  *d.Value,
			}
			err = internal.SaveMetric(g, storage)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"err": err})
				return
			}
			fmt.Println("приняли gauge json", g.GetName(), " ", g.GetValue())
			gj, _ := json.Marshal(g)
			c.Data(http.StatusOK, "application/json", gj)
			return
		}

		if metricType == "counter" {
			g := data.Counter{
				Name: d.ID,
				Val:  *d.Delta,
			}
			err = internal.SaveMetric(g, storage)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"err": err})
				return
			}
			fmt.Println("приняли counter json", g.GetName(), " ", g.GetValue())
			metricValue, err2 := storage.GetValue(d.ID, metricType)
			if err2 != nil {
				c.JSON(http.StatusBadRequest, gin.H{"err": err2})
				return
			}
			gj, _ := json.Marshal(metricValue)
			c.Data(http.StatusOK, "application/json", gj)
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"err": "unknown type"})
	}
	c.JSON(http.StatusBadRequest, gin.H{"err": err})
}

func plainTextUpdate(c *gin.Context, storage internal.StorageOperations) {
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

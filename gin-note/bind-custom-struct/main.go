package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonymousStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func main() {
	router := gin.Default()
	router.GET("/getB", GetDataB)
	router.GET("/getC", GetDataC)
	router.GET("/getD", GetDataD)
	router.Run(":9090")
}

func GetDataB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(http.StatusOK, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func GetDataC(c *gin.Context) {
	var cs StructC
	c.Bind(&cs)
	c.JSON(http.StatusOK, gin.H{
		"a": cs.NestedStructPointer,
		"c": cs.FieldC,
	})
}

func GetDataD(c *gin.Context) {
	var d StructD
	c.Bind(&d)
	c.JSON(http.StatusOK, gin.H{
		"x": d.NestedAnonymousStruct,
		"d": d.FieldD,
	})
}

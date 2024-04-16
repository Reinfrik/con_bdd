package main

import (
	"net/http"
	"oracle23c/conexion"
	"oracle23c/oracle"
	"oracle23c/structs"

	"github.com/gin-gonic/gin"
)

var usu = []structs.Usuario{}

func main() {
	conexion.Conexiones()
	router := gin.Default()

	router.GET("/usuarios")
	router.GET("/usuarios/:usu_nrut")
	router.POST("/usuarios")
	router.DELETE("/usuarios/:usu_nrut")
	router.Run("localhost:8080")

}

func PostUsuario(c *gin.Context) {
	var newUsuario structs.Usuario

	if err := c.BindJSON(&newUsuario); err != nil {
		return

	}
	usu = append(usu, newUsuario)
	c.IndentedJSON(http.StatusCreated, newUsuario)
	oracle.Inserta_Usuario(&usu)
}

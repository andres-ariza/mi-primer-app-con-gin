package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// // This handler will match /user/john but will not match /user/ or /user
	// // obtiene un string y devuelve un json response con formato {msg: "Holis [name]"}
	// router.GET("/saludo/:nombre", func(c *gin.Context) {
	// 	name := c.Param("nombre")
	// 	c.JSON(200, gin.H{
	// 		"msg": "Holis " + name,
	// 	})
	// })
	// almacena el json en la variable datosJson
	datosJson, err := ioutil.ReadFile("transaccion.json")
	if err != nil {
		log.Fatal(err)
	}
	// instancia una estructura tipo Transacciones y la inicializa con formato json
	// mediante el unmarshall
	transacciones := Transacciones{}
	err = json.Unmarshal(datosJson, &transacciones)
	if err != nil {
		log.Fatal(err)
	}
	// utiliza un metodo anonimo
	router.GET("/transacciones", func(c *gin.Context) {
		c.JSON(200, transacciones)
	})

	router.Run(":8080")
}

type Transacciones struct {
	Id                int     `json:"id"`
	Codigo            string  `json:"codigo"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	Fecha_Transaccion string  `json:"fechaTransaccion"`
}

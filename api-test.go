package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default() //deafultnya
	//r.GET("/home", func(c *gin.Context) { //func(c *gin.Context) -> handler func
	//	c.JSON(200, gin.H{ //json itu bahasa universal untuk semua bahasa pemrograman
	//		"nama":    "Hallo Fio",
	//		"message": "welcome",
	//	})
	//})

	r.GET("/home", homehandler)
	r.GET("/books/:id", bookHandler)
	r.GET("/flash_sale", flashHandler)

	r.GET("/profile", func(c *gin.Context) { //anonymous func
		c.JSON(200, gin.H{
			"test": "test",
		})
	})

	r.Run()
}

func flashHandler(c *gin.Context) {
	flash := c.Param("flash")

	c.JSON(200, gin.H{
		"item":    flash,
		"message": "ok",
	})
}

func bookHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"Id":      id,
		"Message": "Ok",
	})
}
func homehandler(c *gin.Context) {
	c.JSON(200, gin.H{ //json itu bahasa universal untuk semua bahasa pemrograman
		"nama":    "Hallo Fio",
		"message": "welcome",
	})
}

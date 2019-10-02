package https

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/TonyChouteau/elricconnect4/ai"
)

func moveAI(c *gin.Context) {
	board := c.Param("board")
	result := ai.GetBestMove(board)
	c.JSON(200, result%7)
}

/*
Serve function
*/
func Serve() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/ai/:board", moveAI)

	err := http.ListenAndServe(":8083", r)
	//err := http.ListenAndServeTLS(":8083", "/etc/letsencrypt/live/www.domain.com/fullchain.pem", "/etc/letsencrypt/live/www.domain.com/privkey.pem", r)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

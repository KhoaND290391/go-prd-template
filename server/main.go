package main

import "github.com/gin-gonic/gin"

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/get-prime", func(c *gin.Context) {
		var req RequestPrime
		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(400, &ResponsePrime{
				Result: 0,
				Error:  err.Error(),
			})
		} else {
			result, err := GetLowerPrime(req.Number)
			res := &ResponsePrime{
				Result: result,
				Error:  "",
			}
			if err != nil {
				res.Error = err.Error()
			}
			c.JSON(200, res)
		}

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type RequestPrime struct {
	Number int64 `json:"number"`
}

type ResponsePrime struct {
	Result int64  `json:"result"`
	Error  string `json:"error"`
}

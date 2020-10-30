pakage main

 import (
        "fmt"
        "github.com.com/gin-gonic/gan"
 )

func HomePage(c *gin.Context) {
  c.JSON(200, gin.H{
  "message": "Hello wrold",
})
}
  func PostHomePage(c *gin.Context) { 
}
}

  fmt.Println("Hello world")
 r :=ginDefault()
r.GET("/", Home Page)
r.POST("/", PostHomePage)
r.Run()
}

    
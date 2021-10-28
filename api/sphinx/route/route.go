package route

import (
  "fmt"
  "sphinx/docs"
  )

func main() {
  gin.SetMode(gin.DebugMode)
  g := gin.Default()
  _ = sign.Ping()
  g.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
  agentGroup := g.Group("/api")
  _ = agent.Register(agentGroup)
  verify.Register(g.Group("/review"))
  g.Run("0.0.0.0:8080")
  fmt.Println("server up")
  }




package main

import (
	"frame/gee"
)

func main() {
	engine := gee.New()

	engine.GET("/", func(c *gee.Context) {
		c.String(200, "Hello, World")
	})

	engine.GET("/home", func(c *gee.Context) {
		c.HTML(200, `

			<body style='text-align: center'>

				<h1>  home page </h1>

				<p> content </p>
				
			</body>
		
		`)
	})

	engine.GET("/user", func(c *gee.Context) {
		c.String(200, "userId: %s", c.Query("id"))
	})

	engine.Run(":9001")
}

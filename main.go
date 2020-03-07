package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/static"
	"golang.org/x/sync/errgroup"
	"github.com/gin-gonic/gin"

	// "github.com/gin-gonic/autotls"
	// "golang.org/x/crypto/acme/autocert"
)

//for multiple servers
var (
	g errgroup.Group
)

//runs all the things

func main() {

	// serves frontend on port 80
	// should change to 443 if not using certbot

	server01 := &http.Server{
		Addr:         ":80",
		Handler:      frontend(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	g.Go(func() error {
		return server01.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

} //end main func

//actual frontend function

func frontend() http.Handler {

	app := gin.New()

	app.Use(gin.Recovery())
	//the important thing to use to serve your static build
	app.Use(static.Serve("/", static.LocalFile("./build", false)))

	//used for setting up the https
	// m := autocert.Manager{
	//     Prompt:    autocert.AcceptTOS,
	//     HostPolicy: autocert.HostWhitelist("your-url.com"
	//     Cache:     autocert.DirCache("/var/www/.cache"),
	// }
	// //actually sets up the https
	// autotls.RunWithManager(app, &m)

	return app
} //end actual frontend function

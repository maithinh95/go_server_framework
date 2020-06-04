package server

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type server struct {
	ginEngine *gin.Engine
}

/*
Server is repository
*/
type Server interface {
	Engine() *gin.Engine
	ListenAndServe(c chan<- os.Signal, certFile, keyFile string, srv *http.Server)
}

/*
NewServer create new server reporitory
*/
func NewServer(mode string) Server {
	gin.SetMode(mode)
	engine := gin.Default()
	// default allow all origins
	handlerFunc := func() gin.HandlerFunc {
		config := cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodOptions},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "*"},
			AllowCredentials: false,
			MaxAge:           12 * time.Hour,
		}
		return cors.New(config)
	}()
	engine.Use(handlerFunc)
	// insert StaticFile
	func(relativePath, filepath string) {
		handler := func(c *gin.Context) {
			defer func(begin time.Time) {
				logrus.Warningf("| %v | %15v | %15v | %10v | %v", c.Writer.Status(), time.Since(begin), c.ClientIP(), c.Request.Method, c.Request.RequestURI)
			}(time.Now())
			c.File(filepath)
		}
		engine.Handle(http.MethodGet, relativePath, handler)
		engine.Handle(http.MethodHead, relativePath, handler)
	}("/favicon.ico", "./static/favicon.png")
	engine.Handle(http.MethodGet, "/contactus.html", func(c *gin.Context) {
		defer func(begin time.Time) {
			logrus.Warningf("| %v | %15v | %15v | %10v | %v", c.Writer.Status(), time.Since(begin), c.ClientIP(), c.Request.Method, c.Request.RequestURI)
		}(time.Now())
		c.File("./static/contactus.html")
	})
	return &server{
		ginEngine: engine,
	}
}

func (ins *server) Engine() *gin.Engine {
	return ins.ginEngine
}

func (ins *server) ListenAndServe(c chan<- os.Signal, certFile, keyFile string, srv *http.Server) {
	if certFile == "" && keyFile == "" {
		logrus.Infof("[ROUTE] Listening and serving HTTP on %v", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			logrus.Error(err)
		}
	} else {
		logrus.Infof("[ROUTE] Listening and serving HTTPS on %v", srv.Addr)
		if err := srv.ListenAndServeTLS(certFile, keyFile); err != nil {
			logrus.Error(err)
		}
	}
	c <- os.Interrupt
}

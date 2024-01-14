package main

import (

	// "log"
	"bytes"
	"context"
	"io"

	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"time"

	// "strings"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"

	Conf "orders/config"

	Corders "orders/handler/httphandler"

	// for swagger

	healthcheck "github.com/RaMin0/gin-health-check"
	"github.com/alexcesaro/log/stdlog"
	logger "github.com/sirupsen/logrus"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {

	logkoe := stdlog.GetFromFlags()
	addr, err := Conf.DetermineListenAddress()
	if err != nil {
		logger.Fatal(err)
		logger.Println(err)
		logkoe.Info(err)
	}
	host, errhost := Conf.Hostname()
	if errhost != nil {
		logger.Fatal(errhost)
		logger.Println(errhost)
		logkoe.Info(err)
	}
	env, errenv := Conf.Environment()
	if errenv != nil {
		logger.Println(errenv)
		logkoe.Info(errenv)
	} else {
		if env == "production" {
			gin.SetMode(gin.ReleaseMode)
			// router := gin.New()
		} else if env == "development" {
			gin.SetMode(gin.DebugMode)
		}
	}

	router := gin.New()
	LoggingActivity()
	router.Use(RequestLoggerActivity())

	v1 := router.Group("/orders/v1")
	{
		//Modul Wallet
		walletGroup := v1.Group("/wallets")
		{
			walletGroup.GET("/", Corders.GetWallet)
			walletGroup.GET("/userid", Corders.GetWalletByUserId)
		}

		//Modul Transaction
		transactionGroup := v1.Group("/transactions")
		{
			transactionGroup.GET("/", Corders.GetTransaction)
			transactionGroup.PATCH("/", Corders.TransferTransaction)
		}
	}

	c := cors.AllowAll()
	handler := c.Handler(router)
	subhost := host[7:len(host)]
	router.Use(healthcheck.Default())
	// end health cek
	//tes start service
	server := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		logger.Println("Service orders is shutting down...")
		// atomic.StoreInt32(&healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown the service orders: %v\n", err)
		}
		close(done)
	}()

	logger.Println("Service order is ready to handle requests at", subhost, addr)
	// atomic.StoreInt32(&healthy, 1)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", addr, err)
	}

	<-done
	logger.Println("Service orders stopped")
	//end tes start service
	// logger.Info("Server is running in ", subhost, addr)
	// logger.Fatal(http.ListenAndServe(addr, handler))
}
func RequestLoggerActivity() gin.HandlerFunc {
	return func(c *gin.Context) {
		logkoe := stdlog.GetFromFlags()
		// end added
		// LoggingActivity()
		if string(c.Request.Method) == "POST" || string(c.Request.Method) == "PUT" {
			buf, _ := ioutil.ReadAll(c.Request.Body)
			rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
			rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.
			// fmt.Println(readBody(rdr1)) // Print request body
			re := regexp.MustCompile(`\r?\n`)
			var request = re.ReplaceAllString(readBody(rdr1), "")
			logger.WithFields(logger.Fields{
				"method":  c.Request.Method,
				"url":     c.Request.URL,
				"request": re.ReplaceAllString(readBody(rdr1), ""),
			}).Info("HTTP Request Method")
			logkoe.Info("HTTP Request Method", "method=", c.Request.Method, "url=", c.Request.URL, "param=", request)
			c.Request.Body = rdr2
			c.Next()
		} else {
			logger.WithFields(logger.Fields{
				"method": c.Request.Method,
				"url":    c.Request.URL,
			}).Info("HTTP Request Method")
			logkoe.Info("HTTP Request Method", "method=", c.Request.Method, "url=", c.Request.URL)

		}
	}
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	s := buf.String()
	return s
}
func LoggingActivity() {

	dt := time.Now()
	//fmt.Println(dt)
	date := dt.Format("20060102")
	var filename string = "log/log" + date + ".log"

	// Create the log file if doesn't exist. And append to it if it already exists.
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		logger.Fatal(err)
	}
	Formatter := new(logger.TextFormatter)
	// You can change the Timestamp format. But you have to use the same date and time.

	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true

	logger.SetFormatter(Formatter)
	logger.SetOutput(f)

}

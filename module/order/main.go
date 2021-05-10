package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/KarasWinds/The-trading-platform-of-Pokemon-Card-Game/global"
	_ "github.com/KarasWinds/The-trading-platform-of-Pokemon-Card-Game/internal/init"
	"github.com/KarasWinds/The-trading-platform-of-Pokemon-Card-Game/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err: %v", err)
		}
	}()

	log.Printf("ListenAndServe : %v\n", s.Addr)
	// 等待中斷訊號
	quit := make(chan os.Signal)
	// 接收 syscall.SIGINT 和 syscall.SIGTERM 訊號
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shuting down server...")

	// 最大時間控制，用於通知該服務端它有5秒時間來處理原有的請求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/order", Create)
		apiv1.GET("/orders", Query)
	}
	return r
}

func Create(c *gin.Context) {
	param := service.CreateOrderRequest{}
	// valid, errs := app.BindAndValid(c, &param)
	if err := c.ShouldBind(&param); err != nil {
		log.Printf("app.BindAndValid errs: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Invalid Params",
		})
		return
	}
	switch param.OrderType {
	case "buy":
		svc := service.New(c.Request.Context())
		sellOrder, err := svc.QueryMinSellOrder(&param)
		if err != nil {
			log.Printf("svc.QueryMinSellOrder: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "db query error",
			})
			return
		}
		if sellOrder.Price <= param.Price && sellOrder.Price != 0 {
			if err := svc.CompleteOrderAndCreateTrade(sellOrder, &param); err != nil {
				log.Printf("svc.CompleteOrderAndCreateTrade: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"err": "db query error",
				})
				return
			}
		} else {
			if err := svc.CreateOrder(&param); err != nil {
				log.Printf("svc.CreateOrder: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"err": "db query error",
				})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"create": "success",
		})
		return
	case "sell":
		svc := service.New(c.Request.Context())
		buyOrder, err := svc.QueryMaxBuyOrder(&param)
		if err != nil {
			log.Printf("svc.QueryMaxBuyOrder: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "db query error",
			})
			return
		}
		if buyOrder.Price >= param.Price && buyOrder.Price != 0 {
			if err := svc.CompleteOrderAndCreateTrade(buyOrder, &param); err != nil {
				log.Printf("svc.CompleteOrderAndCreateTrade: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"err": "db query error",
				})
				return
			}
		} else {
			if err := svc.CreateOrder(&param); err != nil {
				log.Printf("svc.CreateOrder: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"err": "db query error",
				})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"create": "success",
		})
		return
	}
}

func Query(c *gin.Context) {
	traderID, err := strconv.Atoi(c.Query("trader_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Invalid Params",
		})
		return
	}
	param := service.QueryOrderRequest{
		TraderID: traderID,
	}
	svc := service.New(c.Request.Context())
	orderList, err := svc.QueryTraderOrder(&param)
	c.JSON(http.StatusOK, gin.H{
		"trader_id": traderID,
		"order":     orderList,
	})
	return
}

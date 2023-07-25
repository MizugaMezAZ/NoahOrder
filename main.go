package main

import (
	"context"
	"fmt"
	"gorder/api/controller"
	"gorder/api/repository"
	"gorder/api/router"
	"gorder/api/service"
	"gorder/database"
	"gorder/logger"
	"gorder/logger/zap"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	file := logger.GetRotationWriter("./log/" + "")
	l := zap.New(file, zap.DebugLevel, true)
	logger.SetDefaultLogger(l)

}

func main() {
	router := fiveLionFit()

	engine := gin.Default()

	pprof.Register(engine)
	router.SetupRoute(engine)

	addr := fmt.Sprintf("%s:%s", "", "7123")

	srv := &http.Server{
		Handler:      engine,
		Addr:         addr,
		WriteTimeout: 60 * 10 * time.Second,
		ReadTimeout:  60 * 10 * time.Second,
	}

	//啟動server
	go func() {
		fmt.Println("訂單系統初始化完畢 service listening:", addr, "立馬啟動 : HTTP模式")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("starting http service has error: %v", err)
		}
	}()

	// listening 中斷信號優雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("訂單系統 shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("訂單系統 shutdown err :", err)
	}

	fmt.Println("訂單系統已停止")
}

func fiveLionFit() *router.HttpRouter {
	authRepo := repository.NewAuthRepository(database.DB)
	billRepo := repository.NewBillRepository(database.DB, database.RDB)

	authService := service.NewAuthService(authRepo)
	billService := service.NewBillService(billRepo)

	authController := controller.NewAuthController(authService)
	billController := controller.NewBillController(billService)

	router := router.NewRouter(router.RouterParam{
		AuthController: authController,
		BillController: billController,
	})

	return router
}

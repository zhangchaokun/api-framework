package app

import (
	"api-framework/actions"
	"api-framework/core/config"
	"api-framework/core/logging"
	"api-framework/models"
	"api-framework/store/gredis"
	"fmt"
	"net/http"
)

func Run() {
	loader()

	router := actions.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.ServerConfig.HttpPort),
		Handler:        router,
		ReadTimeout:    config.ServerConfig.ReadTimeout,
		WriteTimeout:   config.ServerConfig.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()

}

func loader() {
	config.InitConfig()
	logging.Setup()
	if config.AppConfig.UseDatabase {
		models.InitModel()
	}
	if config.AppConfig.UseRedis {
		_ = gredis.InitRedis()
	}
}

/**
endless优雅重启
 */
/*
func main() {
	endless.DefaultReadTimeOut    = setting.ReadTimeout
	endless.DefaultWriteTimeOut   = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20

	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
*/

/**
http.Server - Shutdown() 优雅重启
*/
/*
func main() {

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
*/


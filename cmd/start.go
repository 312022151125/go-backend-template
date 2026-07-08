package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/312022151125/go-backend-template/internal/conf"
	_ "github.com/312022151125/go-backend-template/internal/server/handlers"
	"github.com/312022151125/go-backend-template/internal/server/middleware"
	"github.com/312022151125/go-backend-template/internal/server/router"
	"github.com/312022151125/go-backend-template/internal/store"

	"github.com/charmbracelet/log"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var startConfig string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start " + conf.APP_NAME,
	PreRun: func(cmd *cobra.Command, args []string) {
		conf.PrintBanner()
		log.SetReportCaller(true)
		if err := conf.Load(startConfig); err != nil {
			log.Fatalf("load config failed: %v", err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if conf.IsDebug() {
			gin.SetMode(gin.DebugMode)
		} else {
			gin.SetMode(gin.ReleaseMode)
		}

		r := gin.New()

		r.Use(middleware.Cors())
		r.Use(middleware.Logger())
		r.Use(middleware.StaticLocal("/", "static"))
		cookieStore := cookie.NewStore([]byte(conf.APP_NAME))
		cookieStore.Options(sessions.Options{
			Path:     "/",
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})
		r.Use(sessions.Sessions(conf.APP_NAME, cookieStore))

		router.RegisterAll(r)

		if err := store.InitDB(); err != nil {
			log.Errorf("database init error: %v", err)
			return
		}
		defer func() {
			if err := store.Close(); err != nil {
				log.Errorf("database close error: %v", err)
			}
		}()

		if err := store.UserInit(); err != nil {
			log.Errorf("user init error: %v", err)
			return
		}

		addr := fmt.Sprintf("%s:%d", conf.AppConfig.Server.Host, conf.AppConfig.Server.Port)
		log.Infof("http server listening on http://%s", addr)
		httpSrv := &http.Server{Addr: addr, Handler: r}

		go func() {
			if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Errorf("http server listen and serve error: %v", err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		// Block on Ctrl+C so the command exits cleanly instead of Windows exit code 0xc000013a.
		<-quit
	},
}

func init() {
	startCmd.Flags().StringVar(&startConfig, "config", "", "config file (default is ./data/config.json)")
	rootCmd.AddCommand(startCmd)
}

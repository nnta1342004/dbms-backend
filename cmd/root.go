package cmd

import (
	"github.com/gin-gonic/gin"
	goservice "github.com/leductoan3082004/go-sdk"
	"github.com/leductoan3082004/go-sdk/plugin/aws"
	mailer "github.com/leductoan3082004/go-sdk/plugin/mailer/mail"
	"github.com/leductoan3082004/go-sdk/plugin/storage/sdkgorm"
	"github.com/leductoan3082004/go-sdk/plugin/storage/sdkredis"
	jwtProvider "github.com/leductoan3082004/go-sdk/plugin/tokenprovider/jwt"
	"github.com/spf13/cobra"
	"hareta/appCommon"
	"hareta/cmd/handlers"
	"hareta/plugin/rabbitmq"
	"hareta/subscriber"
	"os"
)

func newService() goservice.Service {
	service := goservice.New(
		goservice.WithName("hareta"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(jwtProvider.NewJwtProvider("jwt", appCommon.PluginJwt)),
		goservice.WithInitRunnable(mailer.NewMailEngine("mail", appCommon.PluginMailer)),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("gorm", appCommon.PluginGorm)),
		goservice.WithInitRunnable(aws.New(appCommon.PluginAws)),
		goservice.WithInitRunnable(sdkredis.NewRedisDB(appCommon.PluginRedis, appCommon.PluginRedis)),
		goservice.WithInitRunnable(
			rabbitmq.NewRabbitMQ(
				appCommon.PluginRabbitMQ,
				appCommon.TopicSendMailRecoveryPassword,
			),
		),
	)
	if err := service.Init(); err != nil {
		panic(err)
	}
	return service
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Hareta backend",
	Run: func(cmd *cobra.Command, args []string) {
		service := newService()
		serviceLogger := service.Logger("service")

		service.HTTPServer().AddHandler(
			func(engine *gin.Engine) {
				handlers.MainRoute(engine, service)
			},
		)

		if err := subscriber.NewEngine(service).Start(); err != nil {
			serviceLogger.Fatalln(err)
		}
		if err := service.Start(); err != nil {
			serviceLogger.Fatalln(err)
		}

	},
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

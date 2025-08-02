package initialize

import (
	"encoding/json"
	"log"
	"schedule_gateway/internal/middlewares"
	"schedule_gateway/internal/routers"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func InitRouter(r *gin.Engine) {
	r.Use(middlewares.TrackLogMiddleware())
	r.Use(middlewares.CORSMiddleware())

	permRouter := routers.RouterGroupApp.AuthorizationRouterEnter.PermissionRouter
	roleRouter := routers.RouterGroupApp.AuthorizationRouterEnter.RoleRouter
	tokenRouter := routers.RouterGroupApp.AuthenticationRouterEnter.TokenRouter
	authRouter := routers.RouterGroupApp.AuthenticationRouterEnter.AuthRouter
	userRouter := routers.RouterGroupApp.UserRouterEnter.UserRouter

	MainGroup := r.Group("api/v1/")
	{
		MainGroup.GET("/health", func(c *gin.Context) {

			conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
			if err != nil {
				log.Fatal("Failed to connect:", err)
			}
			defer conn.Close()

			// Create a channel
			ch, err := conn.Channel()
			if err != nil {
				log.Fatal("Failed to open channel:", err)
			}
			defer ch.Close()

			// queueName := "send_otp_mail_queue"

			// Publish a message
			type Message struct {
				Email string `json:"email"`
				OTP   string `json:"otp"`
			}
			msg := Message{
				Email: "sinh@gmai.com",
				OTP:   "123456",
			}
			message, err := json.Marshal(msg)
			if err != nil {
				log.Fatal("Failed to marshal message:", err)
			}
			err = ch.Publish(
				"send_otp_mail_exchange",    // exchange
				"send_otp_mail_routing_key", // routing key
				false,                       // mandatory
				false,                       // immediate
				amqp.Publishing{
					ContentType: "application/json",
					Body:        message,
				},
			)
			if err != nil {
				log.Fatal("Failed to publish message:", err)
			}

			log.Println("Message sent")
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Gateway is running",
			})
		})
	}
	{
		permRouter.InitPermissionRouter(MainGroup)
		roleRouter.InitRoleRouter(MainGroup)
		tokenRouter.InitTokenRouter(MainGroup)
		authRouter.InitAuthRouter(MainGroup)
		userRouter.InitUserRouter(MainGroup)
	}
}

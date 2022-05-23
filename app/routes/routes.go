package routes
import (
	"{{ .AppName }}/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func {{ .DomainName }}Routes(a *fiber.App){
    route:=a.Group("/")
    route.Post("/{{ .DomainName | ToLower }}",controllers.Create{{ .DomainName }})
    route.Get("/{{ .DomainName | ToLower }}/:id",controllers.Get{{ .DomainName }})
    route.Get("/{{ .DomainName | ToLower }}",controllers.GetAll{{ .DomainName }})
    route.Put("/{{ .DomainName | ToLower }}/:id",controllers.Update{{ .DomainName }})
    route.Delete("/{{ .DomainName | ToLower }}",controllers.Delete{{ .DomainName }})
}

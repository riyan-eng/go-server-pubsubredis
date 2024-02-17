package route

import (
	"server/internal/app"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

func ExampleRoute(a *fiber.App, handler *app.ServiceServer, enforcer *casbin.Enforcer) {
	a.Post("/example-import", handler.ImportExample)
	a.Get("/example-template", handler.TemplateExample)
	a.Get("/example-download", handler.DownloadExample)

	// route := a.Group("/example", middleware.AuthorizeJwt(), middleware.PermitCasbin(enforcer))
	route := a.Group("/example")
	route.Get("/", handler.ListExample)
	route.Post("/", handler.CreateExample)
	route.Get("/:id", handler.DetailExample)
	route.Put("/:id", handler.PutExample)
	route.Patch("/:id", handler.PatchExample)
	route.Delete("/:id", handler.DeleteExample)
	route.Delete("/:id", handler.DeleteExample)

}

// func ExampleRoute2(mux *http.ServeMux, handler *app.ServiceServer, enforcer *casbin.Enforcer) {
// 	apiMux := http.NewServeMux()
// 	apiMux.HandleFunc("/", middlewarerrrr(handler.CreateExample2))
// 	apiMux.HandleFunc("/welcome2", middlewarerrrr(ServeHTTP2))
// 	mux.Handle("/", apiMux)
// }

// func middlewarerrrr(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// Middleware logic before calling next handler
// 		// For example, checking authentication, logging, etc.
// 		// Example: if !handler.IsAuthenticated(r) { http.Error(w, "Unauthorized", http.StatusUnauthorized); return }
// 		fmt.Println("this is my middleware")
// 		next(w, r) // calling the next handler

// 		// Middleware logic after calling next handler
// 	}
// }

// func ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the API home page!")
// }

// func ServeHTTP2(w http.ResponseWriter, req *http.Request) {
// 	fmt.Println("this is my function")
// 	fmt.Fprintf(w, "Welcome to the API home page 2!")
// }

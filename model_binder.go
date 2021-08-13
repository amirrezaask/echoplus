package echoplus

import "github.com/labstack/echo/v4"

// func main() {
// 	e := echo.New()
// 	type RootModel struct {
// 		ID string `json:"id"`
// 	}

// 	e.Use(MakeMiddlewareFrom(func(ctx echo.Context) (interface{}, error) {
// 		id := ctx.Param("id")
// 		return &RootModel{ID: id}, nil
// 	}))
// 	e.GET("/:id", func(ctx echo.Context) error {
// 		req := ctx.Get("Model").(*RootModel)
// 		return ctx.JSON(200, map[string]interface{}{"id": req.ID})
// 	})
// 	if err := e.Start(":8080"); err != nil {
// 		panic(err)
// 	}
// }

func MakeModelBinder(binder func(ctx echo.Context) (interface{}, error)) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			model, err := binder(ctx)
			if err != nil {
				return err
			}
			ctx.Set("Model", model)
			return next(ctx)
		}
	}
}

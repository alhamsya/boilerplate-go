package cronMiddleware

import "context"

type FuncOrigin func(context.Context) (interface{}, error)

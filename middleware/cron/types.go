package cronMiddleware

import "context"

type FuncScheduler func()
type FuncOrigin func(context.Context) (interface{}, error)

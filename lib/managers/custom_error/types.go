package customError

import "google.golang.org/grpc/codes"

type Error struct {
	HttpCode int
	GrpcCode codes.Code
	Err      error
}

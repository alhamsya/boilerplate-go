package constCommon

import "time"

const (
	GRPCStatusOk                 = "OK"
	GRPCStatusCancelled          = "CANCELLED"
	GRPCStatusUnknown            = "UNKNOWN"
	GRPCStatusInvalidArgument    = "INVALID_ARGUMENT"
	GRPCStatusDeadlineExceeded   = "DEADLINE_EXCEEDED"
	GRPCStatusNotFound           = "NOT_FOUND"
	GRPCStatusAlreadyExists      = "ALREADY_EXISTS"
	GRPCStatusPermissionDenied   = "PERMISSION_DENIED"
	GRPCStatusResourceExhausted  = "RESOURCE_EXHAUSTED"
	GRPCStatusFailedPrecondition = "FAILED_PRECONDITION"
	GRPCStatusAborted            = "ABORTED"
	GRPCStatusOutOfRange         = "OUT_OF_RANGE"
	GRPCStatusUnimplemented      = "UNIMPLEMENTED"
	GRPCStatusInternal           = "INTERNAL"
	GRPCStatusUnavailable        = "UNAVAILABLE"
	GRPCStatusDataLoss           = "DATA_LOSS"
	GRPCStatusUnauthenticated    = "UNAUTHENTICATED"
)

const (
	FuncHealthCheck = "/grpc.health.v1.Health/Check"

	FuncGRPCCreateUserPayment  = "/protos.DebtUser/CreateUserPayment"
	FuncGRPCCreateOrUpdateDebt = "/protos.DebtUser/CreateOrUpdateDebt"
)

const (
	TypeREST = "REST"
	TypeGRPC = "GRPC"
)

const (
	DefaultGRCPMaxConnectionIdle  = 3 * time.Second
	DefaultGRPCContextWithTimeout = 3 * time.Second
	DefaultServerGRPCGraceTimeout = 10 * time.Second
)

const (
	DefaultServerHTTPGraceTimeout = 10 * time.Second
)

const (
	DefaultServerCRONGraceTimeout = 10 * time.Second
)

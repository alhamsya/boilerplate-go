package circuitBreaker

const (
	ApiTypeHTTP = "http"
	ApiTypeGRPC = "grpc"
)

const (
	// CallTypeInternal : internal/local services dependencies
	CallTypeInternal = "internal"

	// CallTypeExternal : 3rd party dependencies
	CallTypeExternal = "external"

	// CallTypeDatabase : databases dependencies
	CallTypeDatabase = "database"

	// CallTypeCache : cache dependencies
	CallTypeCache = "cache"
)

var EnumAPIType = map[string]bool{
	ApiTypeHTTP: true,
	ApiTypeGRPC: true,
}

var EnumCallType = map[string]bool{
	CallTypeInternal: true,
	CallTypeExternal: true,
	CallTypeDatabase: true,
	CallTypeCache:    true,
}

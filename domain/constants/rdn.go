package constCommon

const (
	StrRdnBCA     = "bca"
	StrRdnPermata = "permata"
	StrRdnJago    = "jago"

	NumRdnBCA     = 1
	NumRdnPermata = 2
	NumRdnJago    = 3
)

var (
	EnumStrRdn = map[string]int{
		StrRdnBCA:     NumRdnBCA,
		StrRdnPermata: NumRdnPermata,
		StrRdnJago:    NumRdnJago,
	}

	EnumNumRdn = map[int]string{
		NumRdnBCA:     StrRdnBCA,
		NumRdnPermata: StrRdnPermata,
		NumRdnJago:    StrRdnJago,
	}
)

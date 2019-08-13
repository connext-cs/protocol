package api

const (
	ServiceName = "gateway"

	NameSpace        = "go.micro.api"
	RegisterTtl      = 60
	RegisterInterval = 60
	APIPath          = "/"
	ReadTimeout      = 10
	WriteTimeout     = 10
	EtcdRootPrefix   = "/gateway/"
	BackendPrefix    = "/backend"
)

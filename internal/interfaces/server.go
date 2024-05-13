package interfaces

type IServer interface {
	REST() IRESTServer
	AMPQ() IAMPQServer
}

type IRESTServer interface {
	Run()
}

type IAMPQServer interface {
	Handle() (err error)
}

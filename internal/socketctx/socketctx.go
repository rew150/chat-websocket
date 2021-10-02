package socketctx

type SocketCtx string

const BlankCtx SocketCtx = ""

func (c SocketCtx) String() string {
	return string(c)
}

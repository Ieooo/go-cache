package peer

type PeerCache interface {
	Get(key string) string
	Set(key string, val string)
	Del(key string)
}

package peer

type peerClient struct {
	baseUrl string
}

func NewPeerClient(endpoint string) *peerClient {
	return &peerClient{
		baseUrl: endpoint,
	}
}

func (p peerClient) Get(key string) string {
	return ""
}
func (p peerClient) Set(key string, val string) {
}
func (p peerClient) Del(key string) {
}

var _ PeerCache = (*peerClient)(nil)

package gzkwrapper

type NodeInfo struct {
	Key  string
	Data *NodeData
}

type PulseHandlerFunc func(key string, nodedata *NodeData, err error)
type NodeHandlerFunc func(online []*NodeInfo, offline []*NodeInfo)
type WatchHandlerFunc func(path string, data []byte, err error)

type INodeNotifyHandler interface {
	OnZkWrapperPulseHandlerFunc(key string, nodedata *NodeData, err error)
	OnZkWrapperNodeHandlerFunc(online []*NodeInfo, offline []*NodeInfo)
}

func (fn PulseHandlerFunc) OnZkWrapperPulseHandlerFunc(key string, nodedata *NodeData, err error) {
	fn(key, nodedata, err)
}

func (fn NodeHandlerFunc) OnZkWrapperNodeHandlerFunc(online []*NodeInfo, offline []*NodeInfo) {
	fn(online, offline)
}

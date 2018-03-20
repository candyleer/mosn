package proxy

import (
	"gitlab.alipay-inc.com/afe/mosn/pkg/api/v2"
	"gitlab.alipay-inc.com/afe/mosn/pkg/types"
	"gitlab.alipay-inc.com/afe/mosn/pkg/server"
	"gitlab.alipay-inc.com/afe/mosn/pkg/proxy"
)

type GenericProxyFilterConfigFactory struct {
	Proxy *v2.Proxy
}

func (gfcf *GenericProxyFilterConfigFactory) CreateFilterFactory(clusterManager types.ClusterManager) server.NetworkFilterFactoryCb {
	return func(manager types.FilterManager) {
		manager.AddReadFilter(proxy.NewProxy(gfcf.Proxy, clusterManager))
	}
}
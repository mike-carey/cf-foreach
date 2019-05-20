package cfforeach

import (
	cfclient "github.com/cloudfoundry-community/go-cfclient"
	"github.com/mike-carey/cf-foreach/cf"
	"github.com/mike-carey/cf-foreach/foreach"
)

//go:generate ifacemaker -f client.go -s Client -i CFForEach -p cfforeach -o interface.go

type Client struct {
	CFClient cf.CFClient
}

func (c *Client) ForEachServiceBindingToApp(serviceBindings []cfclient.ServiceBinding) ([]cfclient.App, []error) {
	return foreach.ServiceBindingToApp(serviceBindings, func(serviceBinding cfclient.ServiceBinding) (cfclient.App, error) {
		return c.CFClient.AppByGuid(serviceBinding.AppGuid)
	})
}

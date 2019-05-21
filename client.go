package cfforeach

import (
	cfclient "github.com/cloudfoundry-community/go-cfclient"
	"github.com/mike-carey/cf-foreach/cf"
	"github.com/mike-carey/cf-foreach/foreach"
)

type Client struct {
	CFClient cf.CFClient
}

func (c *Client) ServiceBindingToApp(serviceBindings []cfclient.ServiceBinding) ([]cfclient.App, []error) {
	return foreach.ServiceBindingToApp(serviceBindings, func(serviceBinding cfclient.ServiceBinding) (cfclient.App, error) {
		return c.CFClient.AppByGuid(serviceBinding.AppGuid)
	})
}

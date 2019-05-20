//

package cfforeach

import (
	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

// CFForEach ...
type CFForEach interface {
	ForEachServiceBindingToApp(serviceBindings []cfclient.ServiceBinding) ([]cfclient.App, []error)
}

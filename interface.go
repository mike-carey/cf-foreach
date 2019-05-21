//

package cfforeach

import (
	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

// CFForEach ...
type CFForEach interface {
	ServiceBindingToApp(serviceBindings []cfclient.ServiceBinding) ([]cfclient.App, []error)
}

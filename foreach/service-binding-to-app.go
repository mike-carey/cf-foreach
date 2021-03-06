// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package foreach

import (
	"sync"

	"github.com/cloudfoundry-community/go-cfclient"
)

func ServiceBindingToApp(these []cfclient.ServiceBinding, do func(cfclient.ServiceBinding) (cfclient.App, error)) ([]cfclient.App, []error) {
	pool := make([]cfclient.App, 0)
	errs := make([]error, 0)

	if len(these) > 0 {
		var wg sync.WaitGroup

		wg.Add(len(these))

		poolCh := make(chan cfclient.App, len(these))
		errsCh := make(chan error, len(these))
		for _, this := range these {
			go func(this cfclient.ServiceBinding) {
				defer wg.Done()

				t, e := do(this)
				if e != nil {
					errsCh <- e
				} else {
					poolCh <- t
				}
			}(this)
		}

		wg.Wait()

		for _ = range these {
			select {
			case this := <-poolCh:
				pool = append(pool, this)
			case err := <-errsCh:
				errs = append(errs, err)
			}
		}
	}

	return pool, errs
}

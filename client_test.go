package cfforeach_test

import (
	cfclient "github.com/cloudfoundry-community/go-cfclient"
	. "github.com/mike-carey/cf-foreach"
	fakes "github.com/mike-carey/cf-foreach/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CFForeachClient", func() {
	var (
		fakeClient *fakes.FakeCFClient
		client     Client
	)

	BeforeEach(func() {
		fakeClient = new(fakes.FakeCFClient)
		client = Client{
			CFClient: fakeClient,
		}
	})

	It("Should find all Apps from ServiceBindings", func() {
		app := cfclient.App{
			Guid: "app1",
			Name: "App 1",
		}

		sb := cfclient.ServiceBinding{
			AppGuid: "app1",
		}

		fakeClient.AppByGuidReturns(app, nil)

		apps, errs := client.ServiceBindingToApp([]cfclient.ServiceBinding{sb})

		Expect(errs).To(BeEmpty())
		Expect(apps).Should(ConsistOf(app))
	})

})

package cf

//go:generate ifacemaker -f ../vendor/github.com/cloudfoundry-community/go-cfclient/app_update.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/app_usage_events.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/appevents.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/apps.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/buildpacks.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/cf_error.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/client.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/domains.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/environmentvariablegroups.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/error.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/events.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/gen_error.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/info.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/isolationsegments.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/org_quotas.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/orgs.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/processes.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/route_mappings.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/routes.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/secgroups.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/service_bindings.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/service_brokers.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/service_instances.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/service_keys.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/service_plan_visibilities.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/service_plans.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/service_usage_events.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/services.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/space_quotas.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/spaces.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/stacks.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/tasks.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/types.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/user_provided_service_instances.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/users.go -f ../vendor/github.com/cloudfoundry-community/go-cfclient/v3types.go -s Client -i CFClient -p cf -o client.go
//go:generate ./patch.sh

//go:generate counterfeiter -o ../fakes/fake_cf_client.go client.go CFClient

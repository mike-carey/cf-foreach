package cfforeach

//go:generate ifacemaker -f client.go -s Client -i CFForEach -p cfforeach -o interface.go

//go:generate counterfeiter -o fakes/fake_cfforeach_client.go interface.go CFForEach

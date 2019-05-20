package foreach

//go:generate echo "Generating ServiceBindingToApp"
//go:generate genny -in=../vendor/github.com/mike-carey/async/foreach.go -out=service-binding-to-app.go -pkg foreach gen "Input=cfclient.ServiceBinding Output=cfclient.App"

//go:generate ./patch.sh

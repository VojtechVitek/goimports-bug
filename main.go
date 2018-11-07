package main

// This used to work until recently:
import "github.com/VojtechVitek/goimports-bug/pkg/ibm"

// Only explicit pkg rename works now:
// import redhat "github.com/VojtechVitek/goimports-bug/pkg/ibm"

func main() {
	redhat.Foo()
}

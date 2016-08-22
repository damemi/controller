package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/openshift/origin/pkg/api/install"
	osclient "github.com/openshift/origin/pkg/client"
	"github.com/openshift/origin/pkg/cmd/util/clientcmd"

	"github.com/spf13/pflag"
	kclient "k8s.io/kubernetes/pkg/client/unversioned"
)

func main() {
	var openshiftClient osclient.Interface
	config, err := clientcmd.DefaultClientConfig(pflag.NewFlagSet("empty", pflag.ContinueOnError)).ClientConfig()
	oc, err := osclient.New(config)
	kubeClient, err := kclient.New(config)
	if err != nil {
		log.Printf("Error creating cluster config: %s", err)
		os.Exit(1)
	}
	openshiftClient, err = osclient.New(config)
	if err != nil {
		log.Printf("Error creating OpenShift client: %s", err)
		os.Exit(2)
	}

	fmt.Printf("Connected to OpenShift!\n")
}

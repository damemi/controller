package controller

import (
	"fmt"

	osclient "github.com/openshift/origin/pkg/client"

	kclient "k8s.io/kubernetes/pkg/client/unversioned"
)

type Controller struct {
	openshiftClient *osclient.Client
	kubeClient      *kclient.Client
}

func NewController(os *osclient.Client, kc *kclient.Client) *Controller {
	return &Controller{
		openshiftClient: os,
		kubeClient:      kc,
	}
}

func (c *Controller) Run() {
	fmt.Printf("Success!\n")
}

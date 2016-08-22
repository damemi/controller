package controller

import (
	"fmt"

	osclient "github.com/openshift/origin/pkg/client"

	kapi "k8s.io/kubernetes/pkg/api"
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
	projects, err := c.openshiftClient.Projects()
	for _, project := range projects.List(kapi.ListOptions{}) {
		fmt.Printf("%s\n", project.ObjectMeta.Name)
	}
}

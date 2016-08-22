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
	projects, err := c.openshiftClient.Projects().List(kapi.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}
	for _, project := range projects.Items {
		fmt.Printf("%s\n", project.ObjectMeta.Name)
	}
}

package controller

import (
	"fmt"

	osclient "github.com/openshift/origin/pkg/client"
	"github.com/openshift/origin/pkg/cmd/util/clientcmd"
	projectapi "github.com/openshift/origin/pkg/project/api"

	"github.com/spf13/pflag"
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/meta"
	kclient "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/kubectl/resource"
	"k8s.io/kubernetes/pkg/runtime"
)

type Controller struct {
	openshiftClient *osclient.Client
	kubeClient      *kclient.Client
	mapper          meta.RESTMapper
	typer           runtime.ObjectTyper
	f               *clientcmd.Factory
}

func NewController(os *osclient.Client, kc *kclient.Client) *Controller {

	f := clientcmd.New(pflag.NewFlagSet("empty", pflag.ContinueOnError))
	mapper, typer := f.Object()

	return &Controller{
		openshiftClient: os,
		kubeClient:      kc,
		mapper:          mapper,
		typer:           typer,
		f:               f,
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

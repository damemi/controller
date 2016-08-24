package controller

import (
	"fmt"
	"time"

	osclient "github.com/openshift/origin/pkg/client"
	"github.com/openshift/origin/pkg/cmd/util/clientcmd"

	"github.com/spf13/pflag"
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/meta"
	kclient "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/wait"
	"k8s.io/kubernetes/pkg/watch"
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

func (c *Controller) Run(stopChan <-chan struct{}) {
	go wait.Until(func() {
		w, err := c.kubeClient.Pods(kapi.NamespaceAll).Watch(kapi.ListOptions{})
		if err != nil {
			fmt.Println(err)
		}
		if w == nil {
			return
		}

		for {
			select {
			case event, ok := <-w.ResultChan():
				c.ProcessEvent(event, ok)
			}
		}
	}, 1*time.Millisecond, stopChan)
}

func (c *Controller) ProcessEvent(event watch.Event, ok bool) {
	if !ok {
		fmt.Println("Error received from watch channel")
	}
	if event.Type == watch.Error {
		fmt.Println("Watch channel error")
	}

	switch t := event.Object.(type) {
	case *kapi.Pod:
		fmt.Printf("%s pod %s in namespace %s\n", event.Type, t.ObjectMeta.Name, t.ObjectMeta.Namespace)
	default:
		fmt.Printf("Unknown type\n")
	}
}

package main

import (
	"flag"
	"time"

	"github.com/golang/glog"
	clientset "github.com/krallistic/redis-operator-demo/pkg/client/clientset/versioned"
	informers "github.com/krallistic/redis-operator-demo/pkg/client/informers/externalversions"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/sample-controller/pkg/signals"

	c "github.com/krallistic/redis-operator-demo/controller"
)

var (
	kubeconfig = flag.String("kubeconfig", "/Users/jakobkaralus/.kube/config", "Path to a kubeconfig. Only required if out-of-cluster.")
	masterURL  = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
)

func main() {

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(*masterURL, *kubeconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	redisClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %s", err.Error())
	}

	redisInformerFactory := informers.NewSharedInformerFactory(redisClient, time.Second*30)

	controller := c.NewController(kubeClient, redisClient, redisInformerFactory)

	go redisInformerFactory.Start(stopCh)

	if err = controller.Run(stopCh); err != nil {
		glog.Fatalf("Error running controller: %s", err.Error())
	}

}

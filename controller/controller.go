/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"fmt"

	"github.com/golang/glog"
	appsv1beta2 "k8s.io/api/apps/v1beta2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"

	redisv1 "github.com/krallistic/redis-operator-demo/pkg/apis/krallistic.com/v1"
	clientset "github.com/krallistic/redis-operator-demo/pkg/client/clientset/versioned"
	redisscheme "github.com/krallistic/redis-operator-demo/pkg/client/clientset/versioned/scheme"

	informers "github.com/krallistic/redis-operator-demo/pkg/client/informers/externalversions"
	listers "github.com/krallistic/redis-operator-demo/pkg/client/listers/krallistic/v1"
)

const controllerAgentName = "sample-controller"

const (
	// SuccessSynced is used as part of the Event 'reason' when a Foo is synced
	SuccessSynced = "Synced"
	// ErrResourceExists is used as part of the Event 'reason' when a Foo fails
	// to sync due to a Deployment of the same name already existing.
	ErrResourceExists = "ErrResourceExists"

	// MessageResourceExists is the message used for Events when a resource
	// fails to sync due to a Deployment already existing
	MessageResourceExists = "Resource %q already exists and is not managed by Foo"
	// MessageResourceSynced is the message used for an Event fired when a Foo
	// is synced successfully
	MessageResourceSynced = "Foo synced successfully"
)

// Controller is the controller implementation for Foo resources
type Controller struct {
	// kubeclientset is a standard kubernetes clientset
	kubeclientset kubernetes.Interface
	// redisclientset is a clientset for our own API group
	redisclientset clientset.Interface

	redisLister listers.RedisdbLister
	redisSynced cache.InformerSynced

	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	recorder record.EventRecorder
}

// NewController returns a new sample controller
func NewController(
	kubeclientset kubernetes.Interface,
	redisclientset clientset.Interface,
	redisInformerFactory informers.SharedInformerFactory) *Controller {

	// obtain references to shared index informers for the Deployment and Foo
	// types.
	redisInformer := redisInformerFactory.Krallistic().V1().Redisdbs()

	// Create event broadcaster
	// Add sample-controller types to the default Kubernetes Scheme so Events can be
	// logged for sample-controller types.
	redisscheme.AddToScheme(scheme.Scheme)
	glog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	controller := &Controller{
		kubeclientset:  kubeclientset,
		redisclientset: redisclientset,
		redisLister:    redisInformer.Lister(),
		redisSynced:    redisInformer.Informer().HasSynced,
		recorder:       recorder,
	}

	glog.Info("Setting up event handlers")
	// Set up an event handler for when Foo resources change
	redisInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    controller.added,
		UpdateFunc: controller.updated,
		DeleteFunc: controller.deleted,
	})

	return controller
}

func (c *Controller) Run(stopCh <-chan struct{}) error {
	defer runtime.HandleCrash()

	// Start the informer factories to begin populating the informer caches
	glog.Info("Starting Redis controller")

	// Wait for the caches to be synced before starting workers
	glog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, c.redisSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	glog.Info("Started workers")
	<-stopCh
	glog.Info("Shutting down workers")

	return nil
}

func (c *Controller) added(obj interface{}) {
	redis := obj.(*redisv1.Redisdb)
	deployment := newDeployment(redis)

	glog.Info("Added Redis Object")
	c.kubeclientset.Apps().Deployments("default").Create(deployment)

}

func (c *Controller) deleted(obj interface{}) {
	redis := obj.(*redisv1.Redisdb)
	deployment := newDeployment(redis)

	glog.Info("Deleted Redis Object")
	c.kubeclientset.Apps().Deployments("default").Delete(deployment.Name, &metav1.DeleteOptions{})
}

//THIS is currently not idempotent!!
func (c *Controller) updated(new, old interface{}) {
	redis_new := new.(*redisv1.Redisdb)
	redis_old := old.(*redisv1.Redisdb)
	glog.Info("Updated Redis Object")
	//Detect Changes: Up/DownScales etc...
	if *redis_new.Spec.Replicas > *redis_old.Spec.Replicas {
		//Upscale
		glog.Info("Upscale, join new cluster")
		joinNewNode(redis_new)
	} else if *redis_new.Spec.Replicas < *redis_old.Spec.Replicas {
		glog.Info("Downsclae, migrate Data")
		moveDataFromNode(redis_new, redis_old)
	}
	deployment := newDeployment(redis_new)

	c.kubeclientset.Apps().Deployments("default").Update(deployment)
}

// newDeployment creates a new Deployment for a Foo resource. It also sets
// the appropriate OwnerReferences on the resource so handleObject can discover
// the Foo resource that 'owns' it.
func newDeployment(redis *redisv1.Redisdb) *appsv1beta2.Deployment {
	labels := map[string]string{
		"app":        "redis",
		"controller": redis.Name,
	}
	return &appsv1beta2.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      redis.Name,
			Namespace: redis.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(redis, schema.GroupVersionKind{
					Group:   redisv1.SchemeGroupVersion.Group,
					Version: redisv1.SchemeGroupVersion.Version,
					Kind:    "Redis",
				}),
			},
		},
		Spec: appsv1beta2.DeploymentSpec{
			Replicas: redis.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "redis",
							Image: "redis:3.2",
						},
					},
				},
			},
		},
	}
}

func joinNewNode(redis *redisv1.Redisdb) {
	//Demo Stubs
}

func moveDataFromNode(new, old *redisv1.Redisdb) {
	//Demo Stubs
}

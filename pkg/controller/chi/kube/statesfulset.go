// Copyright 2019 Altinity Ltd and/or its affiliates. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kube

import (
	"context"
	"fmt"
	apps "k8s.io/api/apps/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	kube "k8s.io/client-go/kubernetes"

	api "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
	"github.com/altinity/clickhouse-operator/pkg/controller"
	"github.com/altinity/clickhouse-operator/pkg/interfaces"
)

type STS struct {
	kubeClient kube.Interface
	namer      interfaces.INameManager
}

func NewSTS(kubeClient kube.Interface, namer interfaces.INameManager) *STS {
	return &STS{
		kubeClient: kubeClient,
		namer:      namer,
	}
}

// Get gets StatefulSet. Accepted types:
//  1. *meta.ObjectMeta
//  2. *chop.Host
func (c *STS) Get(ctx context.Context, obj any) (*apps.StatefulSet, error) {
	switch obj := obj.(type) {
	case meta.Object:
		return c.kubeClient.AppsV1().StatefulSets(obj.GetNamespace()).Get(controller.NewContext(), obj.GetName(), controller.NewGetOptions())
	case *api.Host:
		// Namespaced name
		name := c.namer.Name(interfaces.NameStatefulSet, obj)
		namespace := obj.Runtime.Address.Namespace

		return c.kubeClient.AppsV1().StatefulSets(namespace).Get(ctx, name, controller.NewGetOptions())
	}
	return nil, fmt.Errorf("unknown type")
}

func (c *STS) Create(ctx context.Context, statefulSet *apps.StatefulSet) (*apps.StatefulSet, error) {
	return c.kubeClient.AppsV1().StatefulSets(statefulSet.Namespace).Create(ctx, statefulSet, controller.NewCreateOptions())
}

// updateStatefulSet is an internal function, used in reconcileStatefulSet only
func (c *STS) Update(ctx context.Context, sts *apps.StatefulSet) (*apps.StatefulSet, error) {
	return c.kubeClient.AppsV1().StatefulSets(sts.Namespace).Update(ctx, sts, controller.NewUpdateOptions())
}

// deleteStatefulSet gracefully deletes StatefulSet through zeroing Pod's count
func (c *STS) Delete(ctx context.Context, namespace, name string) error {
	return c.kubeClient.AppsV1().StatefulSets(namespace).Delete(ctx, name, controller.NewDeleteOptions())
}

func (c *STS) List(ctx context.Context, namespace string, opts meta.ListOptions) ([]apps.StatefulSet, error) {
	list, err := c.kubeClient.AppsV1().StatefulSets(namespace).List(ctx, opts)
	if err != nil {
		return nil, err
	}
	if list == nil {
		return nil, err
	}
	return list.Items, nil
}
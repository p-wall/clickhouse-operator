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

	policy "k8s.io/api/policy/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/altinity/clickhouse-operator/pkg/controller"
)

type PDB struct {
	kubeClient client.Client
}

func NewPDB(kubeClient client.Client) *PDB {
	return &PDB{
		kubeClient: kubeClient,
	}
}

func (c *PDB) Create(ctx context.Context, pdb *policy.PodDisruptionBudget) (*policy.PodDisruptionBudget, error) {
	err := c.kubeClient.Create(ctx, pdb)
	return pdb, err
}

func (c *PDB) Get(ctx context.Context, namespace, name string) (*policy.PodDisruptionBudget, error) {
	pdb := &policy.PodDisruptionBudget{}
	err := c.kubeClient.Get(controller.NewContext(), types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}, pdb)
	return pdb, err
}

func (c *PDB) Update(ctx context.Context, pdb *policy.PodDisruptionBudget) (*policy.PodDisruptionBudget, error) {
	err := c.kubeClient.Update(controller.NewContext(), pdb)
	return pdb, err
}

func (c *PDB) Delete(ctx context.Context, namespace, name string) error {
	pvc := &policy.PodDisruptionBudget{
		ObjectMeta: meta.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
	}
	return c.kubeClient.Delete(ctx, pvc)
}

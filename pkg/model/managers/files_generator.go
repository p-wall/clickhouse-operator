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

package managers

import (
	api "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
	"github.com/altinity/clickhouse-operator/pkg/interfaces"
	chiConfig "github.com/altinity/clickhouse-operator/pkg/model/chi/config"
	chkConfig "github.com/altinity/clickhouse-operator/pkg/model/chk/config"
)

type FilesGeneratorType string

const (
	FilesGeneratorTypeClickHouse FilesGeneratorType = "clickhouse"
	FilesGeneratorTypeKeeper     FilesGeneratorType = "keeper"
)

func NewConfigFilesGenerator(what FilesGeneratorType, cr api.ICustomResource, opts any) interfaces.IConfigFilesGenerator {
	switch what {
	case FilesGeneratorTypeClickHouse:
		return chiConfig.NewConfigFilesGeneratorClickHouse(cr, NewNameManager(NameManagerTypeClickHouse), opts.(*chiConfig.GeneratorOptions))
	case FilesGeneratorTypeKeeper:
		return chkConfig.NewConfigFilesGeneratorKeeper(cr, opts.(*chkConfig.GeneratorOptions))
	}
	panic("unknown config files generator type")
}
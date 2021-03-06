// Copyright (C) 2019-2020 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License.

#include <condition_variable>
#include <deque>
#include <list>
#include <memory>
#include <mutex>
#include <queue>
#include <string>
#include <thread>
#include <unordered_map>
#include <vector>

#include "db/engine/EngineFactory.h"
#include "db/engine/ExecutionEngine.h"
#include "db/meta/MetaTypes.h"

namespace milvus {
namespace scheduler {

using TableFileSchemaPtr = engine::meta::TableFileSchemaPtr;
using TableFileSchema = engine::meta::TableFileSchema;

using ExecutionEnginePtr = engine::ExecutionEnginePtr;
using EngineFactory = engine::EngineFactory;
using EngineType = engine::EngineType;
using MetricType = engine::MetricType;

}  // namespace scheduler
}  // namespace milvus

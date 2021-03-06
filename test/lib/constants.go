/*
Copyright 2019 Google LLC

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

package lib

const (
	ProwProjectKey = "E2E_PROJECT_ID"

	EventCountMetricType     = "custom.googleapis.com/cloud.google.com/source/event_count"
	GlobalMetricResourceType = "global"
	StorageResourceGroup     = "storages.events.cloud.google.com"
	PubsubResourceGroup      = "pubsubs.events.cloud.google.com"

	BrokerEventCountMetricType = "knative.dev/eventing/broker/event_count"
	BrokerMetricResourceType   = "knative_broker"

	TriggerEventCountMetricType       = "knative.dev/eventing/trigger/event_count"
	TriggerEventDispatchLatencyType   = "knative.dev/eventing/trigger/event_dispatch_latencies"
	TriggerEventProcessingLatencyType = "knative.dev/eventing/trigger/event_processing_latencies"
	TriggerMonitoredResourceType      = "knative_trigger"

	EventType          = "type"
	EventSource        = "source"
	EventDataSchema    = "dataschema"
	EventSubject       = "subject"
	EventSubjectPrefix = "subject-prefix"
	EventID            = "id"
	EventData          = "data"

	E2ERespEventIDPrefix    = "e2e-testing-resp-event-id"
	E2EPubSubRespEventID    = E2ERespEventIDPrefix + "-pubsub"
	E2EBuildRespEventID     = E2ERespEventIDPrefix + "-build"
	E2EStorageRespEventID   = E2ERespEventIDPrefix + "-storage"
	E2EAuditLogsRespEventID = E2ERespEventIDPrefix + "-auditlogs"
	E2ESchedulerRespEventID = E2ERespEventIDPrefix + "-scheduler"
	E2ESampleRespEventID    = E2ERespEventIDPrefix + "-sample"

	E2ERespEventTypePrefix  = "e2e-testing-resp-event-type"
	E2EPubSubRespEventType  = E2ERespEventTypePrefix + "-pubsub"
	E2EBuildRespEventType   = E2ERespEventTypePrefix + "-build"
	E2EStorageRespEventType = E2ERespEventTypePrefix + "-storage"
	E2EAuditLogsRespType    = E2ERespEventTypePrefix + "-auditlogs"
	E2ESchedulerRespType    = E2ERespEventTypePrefix + "-scheduler"
	E2ESampleRespEventType  = E2ERespEventTypePrefix + "-sample"

	// Used in ../test_images/sender, ../test_images/receiver and ../test_images/receiver
	// E2ESampleEventID is the id of the event sent by image `sender`
	E2ESampleEventID = "e2e-sample-event-id"
	// E2ESampleEventType is the type of the event sent by image `sender`
	E2ESampleEventType = "e2e-sample-event-type"
	// E2ESampleEventSource is the source of the event sent by image `sender`
	E2ESampleEventSource = "e2e-sample-event-source"
	// E2ESampleRespEventSource is the source of the resp event sent by image `receiver`
	E2ESampleRespEventSource = "e2e-sample-resp-event-source"
)

// +build go1.9

// Copyright 2017 Microsoft Corporation
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

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: dab57ee609fffdc578f48519d5cdc980efd8cc00

package media

import original "github.com/Azure/azure-sdk-for-go/services/mediaservices/mgmt/2015-10-01/media"

type OperationsClient = original.OperationsClient
type ServiceClient = original.ServiceClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type EntityNameUnavailabilityReason = original.EntityNameUnavailabilityReason

const (
	AlreadyExists	EntityNameUnavailabilityReason	= original.AlreadyExists
	Invalid		EntityNameUnavailabilityReason	= original.Invalid
	None		EntityNameUnavailabilityReason	= original.None
)

type KeyType = original.KeyType

const (
	Primary		KeyType	= original.Primary
	Secondary	KeyType	= original.Secondary
)

type ResourceType = original.ResourceType

const (
	Mediaservices ResourceType = original.Mediaservices
)

type APIEndpoint = original.APIEndpoint
type APIError = original.APIError
type CheckNameAvailabilityInput = original.CheckNameAvailabilityInput
type CheckNameAvailabilityOutput = original.CheckNameAvailabilityOutput
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type RegenerateKeyInput = original.RegenerateKeyInput
type RegenerateKeyOutput = original.RegenerateKeyOutput
type Resource = original.Resource
type Service = original.Service
type ServiceCollection = original.ServiceCollection
type ServiceKeys = original.ServiceKeys
type ServiceProperties = original.ServiceProperties
type StorageAccount = original.StorageAccount
type SyncStorageKeysInput = original.SyncStorageKeysInput

func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewServiceClient(subscriptionID string) ServiceClient {
	return original.NewServiceClient(subscriptionID)
}
func NewServiceClientWithBaseURI(baseURI string, subscriptionID string) ServiceClient {
	return original.NewServiceClientWithBaseURI(baseURI, subscriptionID)
}

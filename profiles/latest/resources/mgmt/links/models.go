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

package links

import original "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-09-01/links"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient
type Filter = original.Filter

const (
	AtScope Filter = original.AtScope
)

type ResourceLink = original.ResourceLink
type ResourceLinkFilter = original.ResourceLinkFilter
type ResourceLinkProperties = original.ResourceLinkProperties
type ResourceLinkResult = original.ResourceLinkResult
type ResourceLinksClient = original.ResourceLinksClient

func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewResourceLinksClient(subscriptionID string) ResourceLinksClient {
	return original.NewResourceLinksClient(subscriptionID)
}
func NewResourceLinksClientWithBaseURI(baseURI string, subscriptionID string) ResourceLinksClient {
	return original.NewResourceLinksClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}

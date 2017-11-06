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

package servicefabric

import original "github.com/Azure/azure-sdk-for-go/services/servicefabric/mgmt/2016-09-01/servicefabric"

type ClustersClient = original.ClustersClient
type ClusterVersionsClient = original.ClusterVersionsClient
type ClusterState = original.ClusterState

const (
	AutoScale			ClusterState	= original.AutoScale
	BaselineUpgrade			ClusterState	= original.BaselineUpgrade
	Deploying			ClusterState	= original.Deploying
	EnforcingClusterVersion		ClusterState	= original.EnforcingClusterVersion
	Ready				ClusterState	= original.Ready
	UpdatingInfrastructure		ClusterState	= original.UpdatingInfrastructure
	UpdatingUserCertificate		ClusterState	= original.UpdatingUserCertificate
	UpdatingUserConfiguration	ClusterState	= original.UpdatingUserConfiguration
	UpgradeServiceUnreachable	ClusterState	= original.UpgradeServiceUnreachable
	WaitingForNodes			ClusterState	= original.WaitingForNodes
)

type DurabilityLevel = original.DurabilityLevel

const (
	Bronze	DurabilityLevel	= original.Bronze
	Gold	DurabilityLevel	= original.Gold
	Silver	DurabilityLevel	= original.Silver
)

type Environment = original.Environment

const (
	Linux	Environment	= original.Linux
	Windows	Environment	= original.Windows
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled	ProvisioningState	= original.Canceled
	Failed		ProvisioningState	= original.Failed
	Succeeded	ProvisioningState	= original.Succeeded
	Updating	ProvisioningState	= original.Updating
)

type ReliabilityLevel = original.ReliabilityLevel

const (
	ReliabilityLevelBronze	ReliabilityLevel	= original.ReliabilityLevelBronze
	ReliabilityLevelGold	ReliabilityLevel	= original.ReliabilityLevelGold
	ReliabilityLevelSilver	ReliabilityLevel	= original.ReliabilityLevelSilver
)

type ReliabilityLevel1 = original.ReliabilityLevel1

const (
	ReliabilityLevel1Bronze		ReliabilityLevel1	= original.ReliabilityLevel1Bronze
	ReliabilityLevel1Gold		ReliabilityLevel1	= original.ReliabilityLevel1Gold
	ReliabilityLevel1Platinum	ReliabilityLevel1	= original.ReliabilityLevel1Platinum
	ReliabilityLevel1Silver		ReliabilityLevel1	= original.ReliabilityLevel1Silver
)

type UpgradeMode = original.UpgradeMode

const (
	Automatic	UpgradeMode	= original.Automatic
	Manual		UpgradeMode	= original.Manual
)

type UpgradeMode1 = original.UpgradeMode1

const (
	UpgradeMode1Automatic	UpgradeMode1	= original.UpgradeMode1Automatic
	UpgradeMode1Manual	UpgradeMode1	= original.UpgradeMode1Manual
)

type X509StoreName = original.X509StoreName

const (
	AddressBook		X509StoreName	= original.AddressBook
	AuthRoot		X509StoreName	= original.AuthRoot
	CertificateAuthority	X509StoreName	= original.CertificateAuthority
	Disallowed		X509StoreName	= original.Disallowed
	My			X509StoreName	= original.My
	Root			X509StoreName	= original.Root
	TrustedPeople		X509StoreName	= original.TrustedPeople
	TrustedPublisher	X509StoreName	= original.TrustedPublisher
)

type AvailableOperationDisplay = original.AvailableOperationDisplay
type AzureActiveDirectory = original.AzureActiveDirectory
type CertificateDescription = original.CertificateDescription
type ClientCertificateCommonName = original.ClientCertificateCommonName
type ClientCertificateThumbprint = original.ClientCertificateThumbprint
type Cluster = original.Cluster
type ClusterCodeVersionsListResult = original.ClusterCodeVersionsListResult
type ClusterCodeVersionsResult = original.ClusterCodeVersionsResult
type ClusterHealthPolicy = original.ClusterHealthPolicy
type ClusterListResult = original.ClusterListResult
type ClusterProperties = original.ClusterProperties
type ClusterPropertiesUpdateParameters = original.ClusterPropertiesUpdateParameters
type ClusterUpdateParameters = original.ClusterUpdateParameters
type ClusterUpgradeDeltaHealthPolicy = original.ClusterUpgradeDeltaHealthPolicy
type ClusterUpgradePolicy = original.ClusterUpgradePolicy
type ClusterVersionDetails = original.ClusterVersionDetails
type DiagnosticsStorageAccountConfig = original.DiagnosticsStorageAccountConfig
type EndpointRangeDescription = original.EndpointRangeDescription
type ErrorModel = original.ErrorModel
type ErrorModelError = original.ErrorModelError
type NodeTypeDescription = original.NodeTypeDescription
type OperationListResult = original.OperationListResult
type OperationResult = original.OperationResult
type Resource = original.Resource
type SettingsParameterDescription = original.SettingsParameterDescription
type SettingsSectionDescription = original.SettingsSectionDescription
type OperationsClient = original.OperationsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient

func NewClusterVersionsClient(subscriptionID string) ClusterVersionsClient {
	return original.NewClusterVersionsClient(subscriptionID)
}
func NewClusterVersionsClientWithBaseURI(baseURI string, subscriptionID string) ClusterVersionsClient {
	return original.NewClusterVersionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
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
func NewClustersClient(subscriptionID string) ClustersClient {
	return original.NewClustersClient(subscriptionID)
}
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return original.NewClustersClientWithBaseURI(baseURI, subscriptionID)
}

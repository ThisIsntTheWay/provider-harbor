/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// (lornest) embedding schema and metadata files
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	configauthCluster "github.com/thisisnttheway/provider-harbor/config/cluster/configauth"
	configsecurityCluster "github.com/thisisnttheway/provider-harbor/config/cluster/configsecurity"
	configsystemCluster "github.com/thisisnttheway/provider-harbor/config/cluster/configsystem"
	garbagecollectionCluster "github.com/thisisnttheway/provider-harbor/config/cluster/garbagecollection"
	groupCluster "github.com/thisisnttheway/provider-harbor/config/cluster/group"
	immutabletagruleCluster "github.com/thisisnttheway/provider-harbor/config/cluster/immutabletagrule"
	interrogationservicesCluster "github.com/thisisnttheway/provider-harbor/config/cluster/interrogationservices"
	labelCluster "github.com/thisisnttheway/provider-harbor/config/cluster/label"
	membergroupCluster "github.com/thisisnttheway/provider-harbor/config/cluster/membergroup"
	memberuserCluster "github.com/thisisnttheway/provider-harbor/config/cluster/memberuser"
	preheatinstanceCluster "github.com/thisisnttheway/provider-harbor/config/cluster/preheatinstance"
	projectCluster "github.com/thisisnttheway/provider-harbor/config/cluster/project"
	purgeauditlogCluster "github.com/thisisnttheway/provider-harbor/config/cluster/purgeauditlog"
	registryCluster "github.com/thisisnttheway/provider-harbor/config/cluster/registry"
	replicationCluster "github.com/thisisnttheway/provider-harbor/config/cluster/replication"
	retentionpolicyCluster "github.com/thisisnttheway/provider-harbor/config/cluster/retentionpolicy"
	robotaccountCluster "github.com/thisisnttheway/provider-harbor/config/cluster/robotaccount"
	tasksCluster "github.com/thisisnttheway/provider-harbor/config/cluster/tasks"
	userCluster "github.com/thisisnttheway/provider-harbor/config/cluster/user"
	webhookCluster "github.com/thisisnttheway/provider-harbor/config/cluster/webhook"

	configauthNamespaced "github.com/thisisnttheway/provider-harbor/config/namespaced/configauth"
	configsecurityNamespaced "github.com/thisisnttheway/provider-harbor/config/namespaced/configsecurity"
	configsystemNamespaced "github.com/thisisnttheway/provider-harbor/config/namespaced/configsystem"
	garbagecollectionNamespaced "github.com/thisisnttheway/provider-harbor/config/namespaced/garbagecollection"
	groupNamespaced "github.com/thisisnttheway/provider-harbor/config/namespaced/group"
	immutabletagruleNamespaced "github.com/thisisnttheway/provider-harbor/config/namespaced/immutabletagrule"
	interrogationservicesNamespaced "github.com/thisisnttheway/provider-harbor/config/namespaced/interrogationservices"
	labelNamespaced "github.com/thisisnttheway/provider-harbor/config/namespaced/label"
	membergroupNamespaced "github.com/thisisnttheway/provider-harbor/config/namespaced/membergroup"
	memberuserNamespaced "github.com/thisisnttheway/provider-harbor/config/namespaced/memberuser"
	preheatinstanceNamespaced "github.com/thisisnttheway/provider-harbor/config/namespaced/preheatinstance"
	projectNamespaced "github.com/thisisnttheway/provider-harbor/config/namespaced/project"
	purgeauditlogNamespaced "github.com/thisisnttheway/provider-harbor/config/namespaced/purgeauditlog"
	registryNamespaced "github.com/thisisnttheway/provider-harbor/config/namespaced/registry"
	replicationNamespaced "github.com/thisisnttheway/provider-harbor/config/namespaced/replication"
	retentionpolicyNamespaced "github.com/thisisnttheway/provider-harbor/config/namespaced/retentionpolicy"
	robotaccountNamespaced "github.com/thisisnttheway/provider-harbor/config/namespaced/robotaccount"
	tasksNamespaced "github.com/thisisnttheway/provider-harbor/config/namespaced/tasks"
	userNamespaced "github.com/thisisnttheway/provider-harbor/config/namespaced/user"
	webhookNamespaced "github.com/thisisnttheway/provider-harbor/config/namespaced/webhook"
)

const (
	resourcePrefix = "harbor"
	modulePath     = "github.com/thisisnttheway/provider-harbor"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider(
		[]byte(providerSchema),
		resourcePrefix,
		modulePath,
		[]byte(providerMetadata),
		ujconfig.WithRootGroup("harbor.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		configauthCluster.Configure,
		configsecurityCluster.Configure,
		configsystemCluster.Configure,
		garbagecollectionCluster.Configure,
		groupCluster.Configure,
		immutabletagruleCluster.Configure,
		interrogationservicesCluster.Configure,
		labelCluster.Configure,
		preheatinstanceCluster.Configure,
		projectCluster.Configure,
		membergroupCluster.Configure,
		memberuserCluster.Configure,
		webhookCluster.Configure,
		purgeauditlogCluster.Configure,
		registryCluster.Configure,
		replicationCluster.Configure,
		retentionpolicyCluster.Configure,
		robotaccountCluster.Configure,
		tasksCluster.Configure,
		userCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProvider returns provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider(
		[]byte(providerSchema),
		resourcePrefix,
		modulePath,
		[]byte(providerMetadata),
		ujconfig.WithRootGroup("harbor.m.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		configauthNamespaced.Configure,
		configsecurityNamespaced.Configure,
		configsystemNamespaced.Configure,
		garbagecollectionNamespaced.Configure,
		groupNamespaced.Configure,
		immutabletagruleNamespaced.Configure,
		interrogationservicesNamespaced.Configure,
		labelNamespaced.Configure,
		preheatinstanceNamespaced.Configure,
		projectNamespaced.Configure,
		membergroupNamespaced.Configure,
		memberuserNamespaced.Configure,
		webhookNamespaced.Configure,
		purgeauditlogNamespaced.Configure,
		registryNamespaced.Configure,
		replicationNamespaced.Configure,
		retentionpolicyNamespaced.Configure,
		robotaccountNamespaced.Configure,
		tasksNamespaced.Configure,
		userNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

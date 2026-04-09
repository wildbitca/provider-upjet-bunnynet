// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	accesslist "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/cdn/accesslist"
	edgerule "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/cdn/edgerule"
	hostname "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/cdn/hostname"
	optimizerclass "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/cdn/optimizerclass"
	pullzone "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/cdn/pullzone"
	ratelimitrule "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/cdn/ratelimitrule"
	shield "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/cdn/shield"
	wafrule "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/cdn/wafrule"
	containerapp "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/compute/containerapp"
	containerimageregistry "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/compute/containerimageregistry"
	database "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/compute/database"
	script "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/compute/script"
	scriptsecret "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/compute/scriptsecret"
	scriptvariable "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/compute/scriptvariable"
	record "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/dns/record"
	scriptdns "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/dns/script"
	scriptvariabledns "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/dns/scriptvariable"
	zone "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/dns/zone"
	providerconfig "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/providerconfig"
	file "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/storage/file"
	zonestorage "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/storage/zone"
	collection "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/stream/collection"
	library "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/stream/library"
	video "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/namespaced/stream/video"
)

// Setup_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesslist.Setup,
		edgerule.Setup,
		hostname.Setup,
		optimizerclass.Setup,
		pullzone.Setup,
		ratelimitrule.Setup,
		shield.Setup,
		wafrule.Setup,
		containerapp.Setup,
		containerimageregistry.Setup,
		database.Setup,
		script.Setup,
		scriptsecret.Setup,
		scriptvariable.Setup,
		record.Setup,
		scriptdns.Setup,
		scriptvariabledns.Setup,
		zone.Setup,
		providerconfig.Setup,
		file.Setup,
		zonestorage.Setup,
		collection.Setup,
		library.Setup,
		video.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesslist.SetupGated,
		edgerule.SetupGated,
		hostname.SetupGated,
		optimizerclass.SetupGated,
		pullzone.SetupGated,
		ratelimitrule.SetupGated,
		shield.SetupGated,
		wafrule.SetupGated,
		containerapp.SetupGated,
		containerimageregistry.SetupGated,
		database.SetupGated,
		script.SetupGated,
		scriptsecret.SetupGated,
		scriptvariable.SetupGated,
		record.SetupGated,
		scriptdns.SetupGated,
		scriptvariabledns.SetupGated,
		zone.SetupGated,
		providerconfig.SetupGated,
		file.SetupGated,
		zonestorage.SetupGated,
		collection.SetupGated,
		library.SetupGated,
		video.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

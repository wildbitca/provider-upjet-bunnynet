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
)

// Setup_cdn creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_cdn(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesslist.Setup,
		edgerule.Setup,
		hostname.Setup,
		optimizerclass.Setup,
		pullzone.Setup,
		ratelimitrule.Setup,
		shield.Setup,
		wafrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_cdn creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_cdn(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		accesslist.SetupGated,
		edgerule.SetupGated,
		hostname.SetupGated,
		optimizerclass.SetupGated,
		pullzone.SetupGated,
		ratelimitrule.SetupGated,
		shield.SetupGated,
		wafrule.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

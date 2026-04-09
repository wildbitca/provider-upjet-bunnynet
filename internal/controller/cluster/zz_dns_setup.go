// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	record "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/cluster/dns/record"
	script "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/cluster/dns/script"
	scriptvariable "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/cluster/dns/scriptvariable"
	zone "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/cluster/dns/zone"
)

// Setup_dns creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dns(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		record.Setup,
		script.Setup,
		scriptvariable.Setup,
		zone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_dns creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_dns(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		record.SetupGated,
		script.SetupGated,
		scriptvariable.SetupGated,
		zone.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

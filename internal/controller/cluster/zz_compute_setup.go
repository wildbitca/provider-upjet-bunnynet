// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	containerapp "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/cluster/compute/containerapp"
	containerimageregistry "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/cluster/compute/containerimageregistry"
	database "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/cluster/compute/database"
	script "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/cluster/compute/script"
	scriptsecret "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/cluster/compute/scriptsecret"
	scriptvariable "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/cluster/compute/scriptvariable"
)

// Setup_compute creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_compute(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		containerapp.Setup,
		containerimageregistry.Setup,
		database.Setup,
		script.Setup,
		scriptsecret.Setup,
		scriptvariable.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_compute creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_compute(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		containerapp.SetupGated,
		containerimageregistry.SetupGated,
		database.SetupGated,
		script.SetupGated,
		scriptsecret.SetupGated,
		scriptvariable.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

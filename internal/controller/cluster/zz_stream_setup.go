// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	collection "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/cluster/stream/collection"
	library "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/cluster/stream/library"
	video "github.com/wildbitca/provider-upjet-bunnynet/internal/controller/cluster/stream/video"
)

// Setup_stream creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_stream(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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

// SetupGated_stream creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_stream(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
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

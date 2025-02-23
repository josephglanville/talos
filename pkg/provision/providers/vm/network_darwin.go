// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package vm

import (
	"context"
	"net"

	"github.com/siderolabs/gen/xslices"
	"github.com/siderolabs/talos/pkg/provision"
)

// CreateNetwork does nothing on darwin.
func (p *Provisioner) CreateNetwork(ctx context.Context, state *State, network provision.NetworkRequest, options provision.Options) error {
	ifaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	ifNames := xslices.Map(ifaces, func(iface net.Interface) string { return iface.Name })
	state.BridgeName = getVmnetInterfaceName(ifNames)

	return nil
}

// DestroyNetwork does nothing on darwin.
func (p *Provisioner) DestroyNetwork(state *State) error {
	return nil
}

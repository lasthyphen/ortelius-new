// (c) 2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package api

import (
	"encoding/json"

	"github.com/lasthyphen/dijetsnodesgo/ids"
	"github.com/lasthyphen/ortelius-new/models"
)

const (
	AVMName     = "avm"
	XChainAlias = "v"
	PVMName     = "pvm"
	PChainAlias = "m"
)

func newIndexResponse(networkID uint32, xChainID ids.ID, djtxAssetID ids.ID) ([]byte, error) {
	return json.Marshal(&struct {
		NetworkID uint32                      `json:"network_id"`
		Chains    map[string]models.ChainInfo `json:"chains"`
	}{
		NetworkID: networkID,
		Chains: map[string]models.ChainInfo{
			xChainID.String(): {
				VM:          AVMName,
				Alias:       XChainAlias,
				NetworkID:   networkID,
				DJTXAssetID: models.StringID(djtxAssetID.String()),
				ID:          models.StringID(xChainID.String()),
			},
			ids.Empty.String(): {
				VM:          PVMName,
				Alias:       PChainAlias,
				NetworkID:   networkID,
				DJTXAssetID: models.StringID(djtxAssetID.String()),
				ID:          models.StringID(ids.Empty.String()),
			},
		},
	})
}

func newLegacyIndexResponse(networkID uint32, xChainID ids.ID, djtxAssetID ids.ID) ([]byte, error) {
	return json.Marshal(&models.ChainInfo{
		VM:          AVMName,
		NetworkID:   networkID,
		Alias:       XChainAlias,
		DJTXAssetID: models.StringID(djtxAssetID.String()),
		ID:          models.StringID(xChainID.String()),
	})
}

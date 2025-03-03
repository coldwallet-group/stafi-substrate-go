package expand

import (
	"github.com/coldwallet-group/stafi-substrate-go/expand/acala"
	"github.com/coldwallet-group/stafi-substrate-go/expand/base"
	"github.com/coldwallet-group/stafi-substrate-go/expand/chainX"
	"github.com/coldwallet-group/stafi-substrate-go/expand/crust"
	"github.com/coldwallet-group/stafi-substrate-go/expand/darwinia"
	"github.com/coldwallet-group/stafi-substrate-go/expand/ori"
	"github.com/coldwallet-group/stafi-substrate-go/expand/polkadot"
	"github.com/coldwallet-group/stafi-substrate-go/expand/stafi"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
	"strings"
)

type IEventRecords interface {
	GetBalancesTransfer() []types.EventBalancesTransfer
	GetSystemExtrinsicSuccess() []types.EventSystemExtrinsicSuccess
	GetSystemExtrinsicFailed() []types.EventSystemExtrinsicFailed
}

/*
扩展： 解析event
*/
func DecodeEventRecords(meta *types.Metadata, rawData string, chainName string) (IEventRecords, error) {
	e := types.EventRecordsRaw(types.MustHexDecodeString(rawData))
	var ier IEventRecords
	switch strings.ToLower(chainName) {
	case "chainx":
		var events chainX.ChainXEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "crab", "darwinia":
		var events darwinia.DarwiniaEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "crust":
		var events crust.CRustEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "mandala": // acala mandala 网络
		var events acala.AcalaEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "node": //stafi
		var events stafi.StafiEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "orion":
		var events ori.OrionEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	case "polkadot":
		var events polkadot.PolkadotEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	default:
		var events base.BaseEventRecords
		err := e.DecodeEventRecords(meta, &events)
		if err != nil {
			return nil, err
		}
		ier = &events
	}
	return ier, nil
}

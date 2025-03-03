package test

import (
	"encoding/json"
	"fmt"
	"github.com/coldwallet-group/stafi-substrate-go/client"
	"github.com/stafiprotocol/go-substrate-rpc-client/types"
	"testing"
)

func Test_GetStorage(t *testing.T) {
	blockHash := "0x02c5ac1e520188c91fee16a27a1627dfbe8e5f6d349cae87e2889beece19ec9f"
	c, err := client.New("wss://mainnet-rpc.stafi.io")
	if err != nil {
		t.Fatal(err)
	}
	meta, err := c.C.RPC.State.GetMetadataLatest()
	if err != nil {
		t.Fatal(err)
	}
	storage, err := types.CreateStorageKey(meta, "System", "Events", nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	key := storage.Hex()
	var r interface{}
	err = c.C.Client.Call(&r, "state_getStorage", key, blockHash)
	if err != nil {
		t.Fatal(err)
	}

	e := types.EventRecordsRaw(types.MustHexDecodeString(r.(string)))
	events := types.EventRecords{}
	err = e.DecodeEventRecords(meta, &events)
	if err != nil {
		t.Fatal(err)
	}
	dd, _ := json.Marshal(events)
	fmt.Println(string(dd))
}

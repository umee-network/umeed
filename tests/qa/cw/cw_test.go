//go:build test_qa
// +build test_qa

package cw

import (
	"encoding/json"
	"math/rand"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"gotest.tools/v3/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/umee-network/umee/v4/app"
	"github.com/umee-network/umee/v4/client"
	cwutil "github.com/umee-network/umee/v4/tests/util"
)

const (
	cwGroupPath = "../../artifacts/cw4_group-aarch64.wasm"
)

var (
	SucceessRespCode   = uint32(0)
	TotalAccs          = 1000
	TotalTxsExec       = 100
	cwGroupMsgExecFunc func(name string, msg GroupExecMsg, wg *sync.WaitGroup, accSeq uint64)
)

func TestCWPlusGroup(t *testing.T) {
	cw := &cwutil.Cosmwasm{}
	cw.SetTestingF(t)

	accAddrs := make([]sdk.AccAddress, 0)
	for i := 0; i < TotalAccs; i++ {
		privateKey := secp256k1.GenPrivKey()
		pubKey := privateKey.PubKey()
		accAddrs = append(accAddrs, sdk.AccAddress(pubKey.Address()))
	}

	// remove if old keyring exists for testing
	os.RemoveAll("./keyring-test")
	encConfig := app.MakeEncodingConfig()
	cc, err := ReadConfig("./config_example.yaml")
	assert.NilError(t, err)
	// umee client
	client, err := client.NewClient(cc.ChainID, cc.RPC, cc.GRPC, cc.Mnemonics, 1.5, encConfig)
	assert.NilError(t, err)
	cw.SetUmeeClient(client)
	cw.DeployWasmContract(cwGroupPath)

	admin := client.Tx.SenderAddr()
	// instantiate Contract
	initMsg := GroupInitMsg{
		Admin:   admin.String(),
		Members: make([]Member, 0),
	}

	initMsg.Members = append(initMsg.Members, Member{Addr: admin.String(), Weight: 1})
	for i := 0; i < TotalAccs; i++ {
		initMsg.Members = append(initMsg.Members, Member{
			Addr:   accAddrs[i].String(),
			Weight: 1,
		})
	}

	msg, err := initMsg.Marshal()
	assert.NilError(t, err)

	cw.InstantiateContract(msg)

	// query the contract
	cwGroupQuery := GroupQuery{
		ListMembers: &ListMembers{Limit: 30},
	}
	queryMsg, err := json.Marshal(cwGroupQuery)
	assert.NilError(t, err)

	queryResp := cw.CWQuery(queryMsg)
	var listResp ListMembersResponse
	err = json.Unmarshal([]byte(queryResp.Data), &listResp)
	assert.NilError(t, err)
	assert.Equal(t, 30, len(listResp.Members))

	// doing random txs to flood the cosmwasm network
	wg := &sync.WaitGroup{}
	accSeq, err := client.QueryAuthSeq(admin.String())
	assert.NilError(t, err)
	total := 0

	cwGroupMsgExecFunc = func(name string, msg GroupExecMsg, wg *sync.WaitGroup, accSeq uint64) {
		execMsg, err := msg.Marshal()
		assert.NilError(t, err)
		txResp, err := cw.CWExecuteWithSeqAndAsyncResp(execMsg, accSeq)
		if err != nil && strings.Contains(err.Error(), "account sequence") {
			time.Sleep(time.Second * 1)
			cwGroupMsgExecFunc(name, msg, wg, accSeq)
		}
		if txResp == nil || (txResp != nil && txResp.Code != SucceessRespCode) {
			time.Sleep(time.Second * 1)
			cwGroupMsgExecFunc(name, msg, wg, accSeq)
		}
		if txResp != nil && txResp.Code == SucceessRespCode {
			total = total + 1
			t.Log(name, " total txs successfully executed =", total)
			// TODO: needs to fix panic: sync: negative WaitGroup counter
			wg.Done()
		}
	}

	for i := 0; i < TotalTxsExec; i++ {
		index := rand.Intn(1000)
		updateMembers := GroupExecMsg{
			UpdateMembers: &UpdateMembers{
				Remove: []string{},
				Add: []Member{
					{
						Addr:   accAddrs[index].String(),
						Weight: 1,
					},
				},
			},
		}
		accSeq = accSeq + 1
		wg.Add(1)
		go func(wg *sync.WaitGroup, accSeq uint64) {
			cwGroupMsgExecFunc("update_members", updateMembers, wg, accSeq)
		}(wg, accSeq)
	}

	// updating the admin...
	updateAdmin := GroupExecMsg{
		UpdateAdmin: &UpdateAdmin{
			Admin: accAddrs[1].String(),
		},
	}
	wg.Add(1)
	go func(wg *sync.WaitGroup, accSeq uint64) {
		cwGroupMsgExecFunc("update_admin", updateAdmin, wg, accSeq)
	}(wg, accSeq+1)

	// waiting for go routine to finish
	wg.Wait()

	time.Sleep(time.Second * 6)
	// query the update admin info
	cwGroupQuery = GroupQuery{
		Admin: &Admin{},
	}
	queryMsg, err = json.Marshal(cwGroupQuery)
	assert.NilError(t, err)
	queryResp = cw.CWQuery(queryMsg)
	var adminQuery AdminResp
	err = json.Unmarshal([]byte(queryResp.Data), &adminQuery)
	assert.NilError(t, err)
	assert.Equal(t, updateAdmin.UpdateAdmin.Admin, adminQuery.Admin)
}

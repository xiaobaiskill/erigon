package handlers

import (
	"github.com/ledgerwatch/erigon/cmd/lightclient/sentinel/communication"
	"github.com/ledgerwatch/erigon/cmd/lightclient/sentinel/communication/p2p"
	"github.com/ledgerwatch/erigon/cmd/lightclient/utils"
	"github.com/ledgerwatch/log/v3"
)

// type safe handlers which all have access to the original stream & decompressed data
// ping handler
func pingHandler(ctx *communication.StreamContext, dat *p2p.Ping) error {
	// since packets are just structs, they can be resent with no issue
	_, err := ctx.Codec.WritePacket(dat, SuccessfullResponsePrefix)
	if err != nil {
		return err
	}
	return nil
}

// does nothing
func nilHandler(ctx *communication.StreamContext, dat *communication.EmptyPacket) error {
	return nil
}

// TODO: Actually respond with proper status
func statusHandler(ctx *communication.StreamContext, dat *p2p.Status) error {
	log.Debug("[ReqResp] Status",
		"epoch", dat.FinalizedEpoch,
		"final root", utils.BytesToHex(dat.FinalizedRoot),
		"head root", utils.BytesToHex(dat.HeadRoot),
		"head slot", dat.HeadSlot,
		"fork digest", utils.BytesToHex(dat.ForkDigest),
	)
	_, err := ctx.Codec.WritePacket(dat, SuccessfullResponsePrefix)
	if err != nil {
		return err
	}
	return nil
}
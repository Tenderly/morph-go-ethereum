// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
)

var _ = (*rollupBatchMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (r RollupBatch) MarshalJSON() ([]byte, error) {
	type RollupBatch struct {
		Index                    hexutil.Uint64
		Hash                     common.Hash
		Version                  hexutil.Uint
		ParentBatchHeader        hexutil.Bytes
		Chunks                   []hexutil.Bytes
		SkippedL1MessageBitmap   hexutil.Bytes
		CurrentSequencerSetBytes hexutil.Bytes
		PrevStateRoot            common.Hash
		PostStateRoot            common.Hash
		WithdrawRoot             common.Hash
		Sidecar                  *BlobTxSidecar `rlp:"-"`
	}
	var enc RollupBatch
	enc.Index = hexutil.Uint64(r.Index)
	enc.Hash = r.Hash
	enc.Version = hexutil.Uint(r.Version)
	enc.ParentBatchHeader = r.ParentBatchHeader
	if r.Chunks != nil {
		enc.Chunks = make([]hexutil.Bytes, len(r.Chunks))
		for k, v := range r.Chunks {
			enc.Chunks[k] = v
		}
	}
	enc.SkippedL1MessageBitmap = r.SkippedL1MessageBitmap
	enc.CurrentSequencerSetBytes = r.CurrentSequencerSetBytes
	enc.PrevStateRoot = r.PrevStateRoot
	enc.PostStateRoot = r.PostStateRoot
	enc.WithdrawRoot = r.WithdrawRoot
	enc.Sidecar = r.Sidecar
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (r *RollupBatch) UnmarshalJSON(input []byte) error {
	type RollupBatch struct {
		Index                    *hexutil.Uint64
		Hash                     *common.Hash
		Version                  *hexutil.Uint
		ParentBatchHeader        *hexutil.Bytes
		Chunks                   []hexutil.Bytes
		SkippedL1MessageBitmap   *hexutil.Bytes
		CurrentSequencerSetBytes *hexutil.Bytes
		PrevStateRoot            *common.Hash
		PostStateRoot            *common.Hash
		WithdrawRoot             *common.Hash
		Sidecar                  *BlobTxSidecar `rlp:"-"`
	}
	var dec RollupBatch
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Index != nil {
		r.Index = uint64(*dec.Index)
	}
	if dec.Hash != nil {
		r.Hash = *dec.Hash
	}
	if dec.Version != nil {
		r.Version = uint(*dec.Version)
	}
	if dec.ParentBatchHeader != nil {
		r.ParentBatchHeader = *dec.ParentBatchHeader
	}
	if dec.Chunks != nil {
		r.Chunks = make([][]byte, len(dec.Chunks))
		for k, v := range dec.Chunks {
			r.Chunks[k] = v
		}
	}
	if dec.SkippedL1MessageBitmap != nil {
		r.SkippedL1MessageBitmap = *dec.SkippedL1MessageBitmap
	}
	if dec.CurrentSequencerSetBytes != nil {
		r.CurrentSequencerSetBytes = *dec.CurrentSequencerSetBytes
	}
	if dec.PrevStateRoot != nil {
		r.PrevStateRoot = *dec.PrevStateRoot
	}
	if dec.PostStateRoot != nil {
		r.PostStateRoot = *dec.PostStateRoot
	}
	if dec.WithdrawRoot != nil {
		r.WithdrawRoot = *dec.WithdrawRoot
	}
	if dec.Sidecar != nil {
		r.Sidecar = dec.Sidecar
	}
	return nil
}

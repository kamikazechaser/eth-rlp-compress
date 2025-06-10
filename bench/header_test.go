package bench

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/klauspost/compress/zstd"
)

var (
	encoder, _ = zstd.NewWriter(nil)
	decoder, _ = zstd.NewReader(nil)
)

func TestHeaderRLPHexDecode(t *testing.T) {
	const rawRLPHexFile = "37644747_block_rlp_hex.txt"

	rawHex, err := os.ReadFile(rawRLPHexFile)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", rawRLPHexFile, err)
	}
	log.Printf("Hex content size: %d bytes", len(rawHex))

	hexStr := strings.TrimSpace(string(rawHex))
	if strings.HasPrefix(hexStr, "0x") {
		hexStr = hexStr[2:]
	}
	comressedHex := encoder.EncodeAll([]byte(hexStr), make([]byte, 0, len(hexStr)))
	log.Printf("Hex compressed size: %d bytes", len(comressedHex))

	rawRLP, err := hex.DecodeString(hexStr)
	if err != nil {
		t.Fatalf("Failed to decode hex: %v", err)
	}
	log.Printf("RLP bytes size: %d bytes", len(rawRLP))
	compressedRLP := encoder.EncodeAll(rawRLP, make([]byte, 0, len(rawRLP)))
	log.Printf("RLP compressed size: %d bytes", len(compressedRLP))

	var block types.Block
	if err := rlp.DecodeBytes(rawRLP, &block); err != nil {
		t.Fatalf("Failed to decode RLP: %v", err)
	}

	log.Printf("Decoded block: %d", block.NumberU64())
}

func TestHeaderJSONDecode(t *testing.T) {
	const rawJSONFile = "37644747_block.json"

	rawJSON, err := os.ReadFile(rawJSONFile)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", rawJSONFile, err)
	}
	log.Printf("JSON content size: %d bytes", len(rawJSON))
	compressedJSON := encoder.EncodeAll(rawJSON, make([]byte, 0, len(rawJSON)))
	log.Printf("JSON compressed size: %d bytes", len(compressedJSON))

	var header types.Header
	if err := json.Unmarshal(rawJSON, &header); err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	var blockExtraData struct {
		Transactions []*types.Transaction `json:"transactions"`
		Withdrawals  []*types.Withdrawal  `json:"withdrawals"`
	}
	if err := json.Unmarshal(rawJSON, &blockExtraData); err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	block := types.NewBlockWithHeader(&header).WithBody(types.Body{
		Transactions: blockExtraData.Transactions,
		Withdrawals:  blockExtraData.Withdrawals,
	})

	log.Printf("Decoded block: %v", block.NumberU64())
}

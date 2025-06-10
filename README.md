### Run

* go test -v ./bench/...

### Results

```
=== RUN   TestHeaderRLPHexDecode
2025/06/10 13:29:30 Hex content size: 30366 bytes
2025/06/10 13:29:30 Hex compressed size: 5496 bytes
2025/06/10 13:29:30 RLP bytes size: 15183 bytes
2025/06/10 13:29:30 RLP compressed size: 5444 bytes
2025/06/10 13:29:30 Decoded block: 37644747
--- PASS: TestHeaderRLPHexDecode (0.01s)
=== RUN   TestHeaderJSONDecode
2025/06/10 13:29:30 JSON content size: 42493 bytes
2025/06/10 13:29:30 JSON compressed size: 7705 bytes
2025/06/10 13:29:30 Decoded block: 37644747
--- PASS: TestHeaderJSONDecode (0.00s)
PASS
ok      github.com/kamikazechaser/eth-cache/bench       0.026s
```

### Inference

| Format | Original Size | Compressed Size | Compression Ratio | Space Savings |
|--------|---------------|-----------------|-------------------|---------------|
| **RLP Binary** | 15,183 bytes | **5,444 bytes** | **64.1%** | **35.9%** |
| **Hex String** | 30,366 bytes | 5,496 bytes | 81.9% | 18.1% |
| **JSON** | 42,493 bytes | 7,705 bytes | 81.9% | 18.1% |
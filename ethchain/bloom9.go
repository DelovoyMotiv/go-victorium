package ethchain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/vm"
)

func CreateBloom(txs Transactions) []byte {
	bin := new(big.Int)
	for _, tx := range txs {
		bin.Or(bin, LogsBloom(tx.logs))
	}

	return bin.Bytes()
}

func LogsBloom(logs []vm.Log) *big.Int {
	bin := new(big.Int)
	for _, log := range logs {
		data := [][]byte{log.Address}
		for _, topic := range log.Topics {
			data = append(data, topic)
		}
		data = append(data, log.Data)

		for _, b := range data {
			bin.Or(bin, bloom9(b))
		}
	}

	return bin
}

func bloom9(b []byte) *big.Int {
	r := new(big.Int)
	for _, i := range []int{0, 2, 4} {
		t := big.NewInt(1)

		//r |= 1 << (uint64(b[i+1]) + 256*(uint64(b[i])&1))
		r.Or(r, t.Rsh(t, uint(b[i+1])+256*(uint(b[i])&1)))
	}

	return r
}
package types

import "encoding/binary"

const (
    // ModuleName defines the module name
    ModuleName = "callback"

    // StoreKey defines the primary module store key
    StoreKey = ModuleName

    // RouterKey is the message route for slashing
    RouterKey = ModuleName

    // QuerierRoute defines the module's querier route
    QuerierRoute = ModuleName
)

var (
    KeyPrefixCallbackByHeight = []byte{0x01}
)

func GetCallbackByHeightKeyPrefix(height int64) []byte {
    return append(KeyPrefixCallbackByHeight, Int64ToBytes(height)...)
}

func Int64ToBytes(i int64) []byte {
    buf := make([]byte, 8)
    binary.BigEndian.PutUint64(buf, uint64(i))
    return buf
}
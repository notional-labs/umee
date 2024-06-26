package store

import (
	"math"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"

	"github.com/umee-network/umee/v6/tests/tsdk"
	"github.com/umee-network/umee/v6/util"
)

var (
	keyPrefixUint32 = []byte{0x01}
	keyPrefixOther  = []byte{0x02}
)

func TestGetAndSetDec(t *testing.T) {
	t.Parallel()
	store := tsdk.KVStore(t)
	key := []byte("decKey")
	v1 := sdk.MustNewDecFromStr("1234.5679")
	v2, ok := GetDec(store, key, "no error")
	assert.Equal(t, false, ok)
	assert.DeepEqual(t, sdk.ZeroDec(), v2)

	err := SetDec(store, key, v1, "no error")
	assert.NilError(t, err)

	v2, ok = GetDec(store, key, "no error")
	assert.Equal(t, true, ok)
	assert.DeepEqual(t, v2, v1)
}

func TestGetAndSetInt(t *testing.T) {
	t.Parallel()
	store := tsdk.KVStore(t)
	key := []byte("intKey")
	v2, ok := GetInt(store, key, "no error")
	assert.Equal(t, false, ok)
	assert.DeepEqual(t, sdk.ZeroInt(), v2)

	v1, ok := sdk.NewIntFromString("1234")
	assert.Equal(t, true, ok)
	err := SetInt(store, key, v1, "no error")
	assert.NilError(t, err)

	v2, ok = GetInt(store, key, "no error")
	assert.Equal(t, true, ok)
	assert.DeepEqual(t, v2, v1)
}

func checkStoreNumber[T Integer](name string, val T, store sdk.KVStore, key []byte, t *testing.T) {
	SetInteger(store, key, val)
	vOut, ok := GetInteger[T](store, key)
	require.True(t, ok)
	require.Equal(t, val, vOut, name)
}

func TestStoreNumber(t *testing.T) {
	t.Parallel()
	store := tsdk.KVStore(t)
	key := []byte("integer")

	checkStoreNumber("int32-0", int32(0), store, key, t)
	checkStoreNumber("int32-min", int32(math.MinInt32), store, key, t)
	checkStoreNumber("int32-max", int32(math.MaxInt32), store, key, t)
	checkStoreNumber("uint32-0", uint32(0), store, key, t)
	checkStoreNumber("uint32-max", uint32(math.MaxUint32), store, key, t)
	checkStoreNumber("int64-0", int64(0), store, key, t)
	checkStoreNumber("int64-min", int64(math.MinInt64), store, key, t)
	checkStoreNumber("int64-max", int64(math.MaxInt64), store, key, t)
	checkStoreNumber("uint64-0", uint64(0), store, key, t)
	checkStoreNumber("uint64-max", uint64(math.MaxUint64), store, key, t)
}

func TestSetAndGetAddress(t *testing.T) {
	store := tsdk.KVStore(t)
	key := []byte("uint32")
	val := sdk.AccAddress("1234")

	SetAddress(store, key, val)
	assert.DeepEqual(t, val, GetAddress(store, key))
}

func TestGetAndSetTime(t *testing.T) {
	t.Parallel()
	store := tsdk.KVStore(t)
	key := []byte("tKey")

	_, ok := GetTimeMs(store, key)
	assert.Equal(t, false, ok)

	val := time.Now()
	SetTimeMs(store, key, val)

	val2, ok := GetTimeMs(store, key)
	assert.Equal(t, true, ok)
	val = val.Truncate(time.Millisecond)
	assert.Equal(t, val, val2)
}

func TestLoadAll(t *testing.T) {
	t.Parallel()
	store := tsdk.KVStore(t)

	o1 := newCoin("umee", 1312)
	o2 := newCoin("atom", 2)
	o3 := newCoin("atom", 10)
	o4 := newCoin("gg1", 6)
	o150 := newCoin("aa2", 42)
	assert.NilError(t, SetValue(store, uint32Key(150), &o150, ""))
	assert.NilError(t, SetValue(store, uint32Key(1), &o1, ""))
	assert.NilError(t, SetValue(store, uint32Key(2), &o2, ""))
	assert.NilError(t, SetValue(store, uint32Key(3), &o3, ""))
	assert.NilError(t, SetValue(store, uint32Key(4), &o4, ""))
	// set other data with different prefix
	store.Set(keyPrefixOther, []byte{1})
	store.Set(append(keyPrefixOther, 1), []byte{1})

	elems, err := LoadAllKV[*Uint32, Uint32, *sdk.Coin](store, keyPrefixUint32)
	assert.NilError(t, err)

	newKV := func(k uint32, v sdk.Coin) KV[Uint32, sdk.Coin] {
		return KV[Uint32, sdk.Coin]{Uint32(k), v}
	}

	assert.DeepEqual(t, elems, []KV[Uint32, sdk.Coin]{
		newKV(1, o1),
		newKV(2, o2),
		newKV(3, o3),
		newKV(4, o4),
		newKV(150, o150),
	})

	vals, err := LoadAll[*sdk.Coin](store, keyPrefixUint32)
	assert.NilError(t, err)
	assert.DeepEqual(t, vals, []sdk.Coin{o1, o2, o3, o4, o150})
}

func uint32Key(id uint32) []byte {
	return util.KeyWithUint32(keyPrefixUint32, id)
}

func newCoin(denom string, amount int64) sdk.Coin {
	return sdk.NewCoin(denom, sdk.NewInt(amount))
}

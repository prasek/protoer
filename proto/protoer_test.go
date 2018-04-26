package proto

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelpers(t *testing.T) {

	var b bool = false
	var i32 int32 = 32
	var i int = 32
	var i64 int64 = 64
	var f32 float32 = .32
	var f64 float64 = .64
	var u32 uint32 = 32
	var u64 uint64 = 64
	var s string = "foo"

	require.Equal(t, &b != Bool(b), true)
	require.Equal(t, &i32 != Int32(i32), true)
	require.Equal(t, &i64 != Int64(i64), true)
	require.Equal(t, &f32 != Float32(f32), true)
	require.Equal(t, &f64 != Float64(f64), true)
	require.Equal(t, &u32 != Uint32(u32), true)
	require.Equal(t, &u64 != Uint64(u64), true)
	require.Equal(t, &s != String(s), true)

	require.Equal(t, b, *Bool(b))
	require.Equal(t, i32, *Int32(i32))
	require.Equal(t, int32(i), *Int(i))
	require.Equal(t, i64, *Int64(i64))
	require.Equal(t, f32, *Float32(f32))
	require.Equal(t, f64, *Float64(f64))
	require.Equal(t, u32, *Uint32(u32))
	require.Equal(t, u64, *Uint64(u64))
	require.Equal(t, s, *String(s))
}

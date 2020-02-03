package hmoj

import "testing"

import "reflect"

// regression tests. these will change if you change the emoji map in emoji.go
func TestStd(t *testing.T) {
	testCases := []struct {
		b    []byte
		want string
	}{
		{[]byte("hello"), "🙅‍♀️ 68656c6c6f"},
		{[]byte("long long long this string is long 1234567890"), "🦐 6c6f6e67206c6f"},
		{[]byte{0x01, 0x02, 0x03, 0x04}, "🙃 01020304"},
		{[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}, "🏊‍♀ 01020304050607"},
	}

	for _, tc := range testCases {
		got := Encode(tc.b)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("want Encode(%x)=%s, got %s", tc.b, tc.want, got)
		}
	}
}

package homedir

import (
	"os/user"
	"path/filepath"
	"testing"
)

func BenchmarkDir(b *testing.B) {
	// We do this for any "warmups"
	for i := 0; i < 10; i++ {
		Dir()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Dir()
	}
}

func TestDir(t *testing.T) {
	u, err := user.Current()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	dir, err := Dir()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if u.HomeDir != dir {
		t.Fatalf("%#v != %#v", u.HomeDir, dir)
	}
}

func TestExpand(t *testing.T) {
	u, err := user.Current()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	cases := []struct {
		Input  string
		Output string
		Err    bool
	}{
		{
			"/foo",
			"/foo",
			false,
		},

		{
			"~/foo",
			filepath.Join(u.HomeDir, "foo"),
			false,
		},

		{
			"",
			"",
			false,
		},

		{
			"~",
			u.HomeDir,
			false,
		},

		{
			"~foo/foo",
			"",
			true,
		},
	}

	for _, tc := range cases {
		actual, err := Expand(tc.Input)
		if (err != nil) != tc.Err {
			t.Fatalf("Input: %#v\n\nErr: %s", tc.Input, err)
		}

		if actual != tc.Output {
			t.Fatalf("Input: %#v\n\nOutput: %#v", tc.Input, actual)
		}
	}
}

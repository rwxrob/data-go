package data

import "testing"

func init() { ForceDir = "testdata/datadir" }

func TestDir(t *testing.T) {
	t.Log(Dir())
}

func TestPath(t *testing.T) {
	t.Log(Path("foo", "data.yml"))
}

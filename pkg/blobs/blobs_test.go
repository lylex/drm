package blobs

import (
	"os"
	"testing"
	"time"

	"github.com/lylex/drm/pkg/utils"
)

func TestCreate(t *testing.T) {
	path := "test"
	res := Create(path)
	if res.FileName != path {
		t.Errorf("Generate blob failed")
	}
}

func TestName(t *testing.T) {
	ct, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	randStr := utils.GenerateRandString(6)
	blob := &Blob{
		FileName:  "testname",
		Dir:       "testdir",
		CreatedAt: ct,
		ID:        randStr,
	}
	ret := blob.Name()
	expectName := "1136214245000000000_ti2SMt_testname"
	if ret != expectName {
		t.Errorf("Failed to get name, expect %s, got %s", expectName, ret)
	}
}
func TestSave(t *testing.T) {
	os.Setenv("FOO", "1")
}

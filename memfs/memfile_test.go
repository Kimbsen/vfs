package memfs

import (
	"github.com/Kimbsen/vfs"
	"testing"
)

func TestFileInterface(t *testing.T) {
	_ = vfs.File(NewMemFile("", nil, nil))
}

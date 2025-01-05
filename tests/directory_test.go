package tests

import (
	"os"
	"path"
	"testing"

	"github.com/pazifical/onyx/internal/filesystem"
)

func TestDirectoryTreeCreation(t *testing.T) {
	rootDir := "testdata/directory_tree"

	// Preparation

	err := os.MkdirAll(path.Join(rootDir, "a1", "a1b1"), 0755)
	if err != nil {
		t.Error(err)
	}
	f, err := os.Create(path.Join(rootDir, "a1", "file1.md"))
	f.Close()
	f, err = os.Create(path.Join(rootDir, "a1", "file2.md"))
	f.Close()

	err = os.MkdirAll(path.Join(rootDir, "a1", "a1b2"), 0755)
	if err != nil {
		t.Error(err)
	}
	err = os.MkdirAll(path.Join(rootDir, "a2", "a2b1"), 0755)
	if err != nil {
		t.Error(err)
	}
	err = os.MkdirAll(path.Join(rootDir, "a2", "a2b2"), 0755)
	if err != nil {
		t.Error(err)
	}
	err = os.MkdirAll(path.Join(rootDir, "a2", "a2b3"), 0755)
	if err != nil {
		t.Error(err)
	}

	// Exectution

	dirTree, err := filesystem.CreateDirectoryTree(rootDir)
	if err != nil {
		t.Error(err)
	}

	// Assertion

	if dirTree.Path != rootDir {
		t.Errorf("Expected path %s, got %s", rootDir, dirTree.Path)
	}

	if len(dirTree.Directories) != 2 {
		t.Errorf("Expected directories count to be 2, got %d", len(dirTree.Directories))
	}

	if len(dirTree.Filenames) != 0 {
		t.Errorf("Expected files count to be 0, got %d", len(dirTree.Filenames))
	}

	for _, subDir := range dirTree.Directories {
		if subDir.Path == path.Join(rootDir, "a1") {
			if len(subDir.Directories) != 2 {
				t.Errorf("Expected directories count to be 2 in %s, got %d", subDir.Path, len(subDir.Directories))
			}
			if len(subDir.Filenames) != 2 {
				t.Errorf("Expected files count to be 2 in %s, got %d", subDir.Path, len(subDir.Filenames))
			}
		} else if subDir.Path == path.Join(rootDir, "a2") {
			if len(subDir.Directories) != 3 {
				t.Errorf("Expected directories count to be 3 in %s, got %d", subDir.Path, len(subDir.Directories))
			}
			if len(subDir.Filenames) != 0 {
				t.Errorf("Expected files count to be 0 in %s, got %d", subDir.Path, len(subDir.Filenames))
			}
		} else {
			t.Errorf("Unexpected directory: %s", subDir.Path)
		}
	}

	// Cleanup

	err = os.RemoveAll(rootDir)
	if err != nil {
		t.Error(err)
	}
}

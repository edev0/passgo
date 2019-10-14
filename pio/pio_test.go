package pio

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestDirExists(t *testing.T) {
	t.Run("File with same name exists", func(t *testing.T) {
		home, err := getHomeDir()
		if err != nil {
			t.Fatal(err)
		}
		f, err := ioutil.TempFile(home, ".passgo*")
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			if err := os.Remove(f.Name()); err != nil {
				t.Log("Failed to remove temp file " + f.Name() + ". Please delete it manually.")
			}
		}()
		exists, err := dirExists(filepath.Base(f.Name()))
		if err == nil {
			t.Fatal("PassDirExists should have returned an error")
		}
		if !exists {
			t.Fatalf("Unexpected PassDirExists result. Want:true Have:%v", exists)
		}
	})
	t.Run("Dir already exists", func(t *testing.T) {
		home, err := getHomeDir()
		if err != nil {
			t.Fatal(err)
		}
		d, err := ioutil.TempDir(home, ".passgo")
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			if err := os.Remove(d); err != nil {
				t.Log("Failed to remove temp dir " + d + ". Please delete it manually.")
			}
		}()
		exists, err := dirExists(filepath.Base(d))
		if err != nil {
			t.Fatalf("PassDirExists should have returned an error: %v", err)
		}
		if !exists {
			t.Fatalf("Unexpected PassDirExists result. Want:true Have:%v", exists)
		}
	})
}

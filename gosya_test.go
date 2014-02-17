package gosya

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMerge(t *testing.T) {
	// Make temporary folder and files for configs
	// /path/to/conf/
	// /path/to/conf/settings.yml
	// /path/to/conf/settings.local.yml
	// /path/to/conf/settings/development.yml
	dir_conf, err := ioutil.TempDir("", "conf")
	if err != nil {
		t.Errorf(`ioutil.TempDir("", "conf"): %v`, err)
	}
	dir_settings := dir_conf + "/settings"
	err = os.Mkdir(dir_settings, 0777)
	if err != nil {
		t.Fatalf("os.Mkdir(dir_settings, 0777): %q %v", dir_settings, err)
	}
	// t.Logf("list dirs\ndir_conf: %q\ndir_settings: %q", dir_conf, dir_settings)
	files_settings := []struct {
		name string
		body string
	}{
		{"settings.yml", "foo: global value"},
		{"settings/development.yml", "foo: development value"},
		{"settings.local.yml", "foo: local value"},
	}
	for _, file := range files_settings {
		fname := dir_conf + "/" + file.name
		if err := ioutil.WriteFile(fname, []byte(file.body), 0644); err != nil {
			t.Fatalf("WriteFile %s: %v", fname, err)
		}
	}

	// Init type of struct for settings
	type Settings struct {
		Foo string `yaml:"foo"`
	}
	// Var as target for settings
	s := Settings{}

	err = Merge(&s, dir_conf, "development")
	if err != nil {
		t.Errorf(`Merge(s, dir_conf, "development"): %v`, err)
	}
	if s.Foo != "local value" {
		t.Errorf("Error merging: expect %q, got: %q", "local value", s.Foo)
	}
	// t.Logf("%#v", s)

	err = os.RemoveAll(dir_conf)
	if err != nil {
		t.Fatalf("os.RemoveAll(dir_conf): %q %v", dir_conf, err)
	}
}

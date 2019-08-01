package cmd

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timfpark/yaml"
)

func TestUnset(t *testing.T) {
	// This test changes the cwd. Must change back so any tests following don't break
	cwd, err := os.Getwd()
	assert.Nil(t, err)
	defer func() {
		_ = os.Chdir(cwd)
	}()

	err = os.Chdir("../test/fixtures/unset")
	_ = os.RemoveAll("config")
	assert.Nil(t, err)

	// Set config for two nested subcomponents
	err = Set("common", "a", []string{"x.y.z=abc", "foo=foo"}, false, "")
	err = Set("common", "b", []string{"xyz=abc", "foo.bar=baz"}, false, "")
	assert.Nil(t, err)

	// Remove the config for x.y.z; x.y should be an empty map afterwards
	err = unset([]string{"x.y.z"}, "common", "a")
	assert.Nil(t, err)

	// Removing the a non existant key should return an error
	err = unset([]string{"x.y.i.do.not.exist"}, "common", "a")
	assert.NotNil(t, err)

	// Unsetting a key path which contains a key to a non-map should return an error
	err = unset([]string{"foo.bar.baz"}, "common", "b")
	assert.NotNil(t, err)

	// Read the config yaml
	commonConfig := map[string]interface{}{}
	commonBytes, err := ioutil.ReadFile("config/common.yaml")
	assert.Nil(t, err)
	yaml.Unmarshal(commonBytes, &commonConfig)

	assert.EqualValues(
		t,
		map[string]interface{}{
			"subcomponents": map[string]interface{}{
				"a": map[string]interface{}{
					"config": map[string]interface{}{
						"x": map[string]interface{}{
							"y": map[string]interface{}{},
						},
						"foo": "foo",
					},
				},
				"b": map[string]interface{}{
					"config": map[string]interface{}{
						"xyz": "abc",
						"foo": map[string]interface{}{
							"bar": "baz",
						},
					},
				},
			},
		},
		commonConfig)
}

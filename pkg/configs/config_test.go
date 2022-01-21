package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	const configFile = "testdata/test.ini"
	got, err := LoadConfig(configFile)
	assert.Nil(t, err)

	t.Run("with int value", func(t *testing.T) {
		actual, err := got.GetInt("App.i", 0)
		assert.Nil(t, err)
		assert.Equal(t, 1, actual)
	})

	t.Run("with string value", func(t *testing.T) {
		actual := got.GetString("App.s", "")
		assert.Equal(t, "a", actual)

		actual = got.GetString("App.ss", "d")
		assert.Equal(t, "d", actual)
	})

	t.Run("with bool value", func(t *testing.T) {
		actual, err := got.GetBool("App.b", false)
		assert.Nil(t, err)
		assert.Equal(t, true, actual)
	})

	t.Run("with float value", func(t *testing.T) {
		actual, err := got.GetFloat("App.f", 1)
		assert.Nil(t, err)
		assert.Equal(t, 1.200, actual)
	})
}

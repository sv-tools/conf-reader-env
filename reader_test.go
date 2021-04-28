package confenv_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sv-tools/conf"

	confenv "github.com/sv-tools/conf-reader-env"
)

const envName = "TEST_FOO"

func TestNew(t *testing.T) {
	t.Cleanup(func() {
		require.NoError(t, os.Unsetenv(envName))
	})
	require.NoError(t, os.Setenv(envName, "42"))

	c := conf.New().WithReaders(confenv.New(map[string]string{envName: "foo"}, ""))
	require.NoError(t, c.Load(context.Background()))

	require.Equal(t, 42, c.GetInt("foo"))
}

func ExampleNew() {
	if err := os.Setenv(envName, "42"); err != nil {
		panic(err)
	}

	c := conf.New().WithReaders(confenv.New(map[string]string{envName: "foo"}, ""))
	if err := c.Load(context.Background()); err != nil {
		panic(err)
	}

	fmt.Println(c.GetInt("foo"))
	// Output: 42

	if err := os.Unsetenv(envName); err != nil {
		panic(err)
	}
}

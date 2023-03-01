package confenv

import (
	"context"
	"os"

	"github.com/sv-tools/conf"
)

type envReader struct {
	mapEnvKey map[string]string
	prefix    string
}

func (r *envReader) Prefix() string {
	return r.prefix
}

func (r *envReader) Read(ctx context.Context) (interface{}, error) {
	res := map[string]string{}
	for env, key := range r.mapEnvKey {
		if value, ok := os.LookupEnv(env); ok {
			res[key] = value
		}
	}

	return res, ctx.Err()
}

// New creates the Env reader
//
//	`mapEnvKey` is a map of the names of the environment variables and the configuration keys
//	`prefix` is a default prefix that will be added to all configuration keys
func New(mapEnvKey map[string]string, prefix string) conf.Reader {
	return &envReader{
		mapEnvKey: mapEnvKey,
		prefix:    prefix,
	}
}

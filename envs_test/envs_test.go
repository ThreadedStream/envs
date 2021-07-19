package envs_test

import (
	"github.com/ThreadedStream/envs"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	var (
		exported = struct {
			Host   string `env:"HOST" fallback:"127.0.0.1"`
			Port   int    `env:"PORT" fallback:"3500"`
			UseSSL bool   `env:"USE_SSL" fallback:"true"`
		}{}

		unexported = struct {
			host   string `env:"HOST" fallback:"127.0.0.1"`
			port   int    `env:"PORT" fallback:"3500"`
			UseSSL bool   `env:"USE_SSL" fallback:"false"`
		}{}

		network = struct {
			Host     string `env:"HOST" fallback:"http://localhost:3000"`
			UseProxy bool   `env:"USE_PROXY" fallback:"false"`
		}{}

		err error
	)

	assert.Equal(t, err, nil)

	err = os.Setenv("HOST", "0.0.0.0")
	assert.Equal(t, err, nil)

	err = os.Setenv("PORT", "7000")
	assert.Equal(t, err, nil)

	err = envs.Parse(&exported)
	assert.Equal(t, err, nil)

	assert.Equal(t, exported.Host, "0.0.0.0")
	assert.Equal(t, exported.Port, 7000)
	assert.Equal(t, exported.UseSSL, true)

	err = envs.Parse(&unexported)
	assert.Equal(t, err, nil)

	assert.Equal(t, unexported.host, "")
	assert.Equal(t, unexported.port, 0)
	assert.Equal(t, unexported.UseSSL, false)

	err = os.Setenv("HOST", "http://localhost:3500")
	assert.Equal(t, err, nil)

	err = envs.Parse(&network)
	assert.Equal(t, err, nil)

	assert.Equal(t, network.Host, "http://localhost:3500")
	assert.Equal(t, network.UseProxy, false)
}

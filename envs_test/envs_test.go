package envs_test

import (
	"github.com/ThreadedStream/envs"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	var (
		c = struct {
			Host   string `env:"HOST" fallback:"127.0.0.1"`
			Port   int    `env:"PORT" fallback:"3500"`
			UseSSL bool   `env:"USE_SSL" fallback:"true"`
		}{}
		err error
	)

	assert.Equal(t, err, nil)

	err = os.Setenv("HOST", "0.0.0.0")
	assert.Equal(t, err, nil)

	err = os.Setenv("PORT", "7000")
	assert.Equal(t, err, nil)

	err = envs.Parse(&c)
	assert.Equal(t, err, nil)

	assert.Equal(t, c.Host, "0.0.0.0")
	assert.Equal(t, c.Port, 7000)
	assert.Equal(t, c.UseSSL, true)
}

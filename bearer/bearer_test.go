package bearer_test

import (
	"testing"

	"github.com/nspcc-dev/neo-go/pkg/crypto/keys"
	"github.com/nspcc-dev/neofs-sdk-go/bearer"
	tokentest "github.com/nspcc-dev/neofs-sdk-go/bearer/test"
	"github.com/nspcc-dev/neofs-sdk-go/eacl"
	"github.com/nspcc-dev/neofs-sdk-go/user"
	"github.com/stretchr/testify/require"
)

func TestBearerToken_Issuer(t *testing.T) {
	var bearerToken bearer.Token

	t.Run("non signed token", func(t *testing.T) {
		_, ok := bearerToken.Issuer()
		require.False(t, ok)
	})

	t.Run("signed token", func(t *testing.T) {
		p, err := keys.NewPrivateKey()
		require.NoError(t, err)

		var id user.ID
		user.IDFromKey(&id, p.PrivateKey.PublicKey)

		bearerToken.SetEACLTable(*eacl.NewTable())
		require.NoError(t, bearerToken.Sign(p.PrivateKey))
		issuer, ok := bearerToken.Issuer()
		require.True(t, ok)
		require.True(t, id.Equals(issuer))
	})
}

func TestFilterEncoding(t *testing.T) {
	f := tokentest.Token()

	t.Run("binary", func(t *testing.T) {
		data := f.Marshal()

		var f2 bearer.Token
		require.NoError(t, f2.Unmarshal(data))

		require.Equal(t, f, f2)
	})

	t.Run("json", func(t *testing.T) {
		data, err := f.MarshalJSON()
		require.NoError(t, err)

		var d2 bearer.Token
		require.NoError(t, d2.UnmarshalJSON(data))

		require.Equal(t, f, d2)
	})
}

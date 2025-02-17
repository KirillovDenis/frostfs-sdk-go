package reputation_test

import (
	"testing"

	reputationV2 "github.com/nspcc-dev/neofs-api-go/v2/reputation"
	reputationtestV2 "github.com/nspcc-dev/neofs-api-go/v2/reputation/test"
	"github.com/nspcc-dev/neofs-sdk-go/reputation"
	reputationtest "github.com/nspcc-dev/neofs-sdk-go/reputation/test"
	"github.com/nspcc-dev/neofs-sdk-go/version"
	"github.com/stretchr/testify/require"
)

func TestTrust(t *testing.T) {
	trust := reputation.NewTrust()

	id := reputationtest.PeerID()
	trust.SetPeer(id)
	require.Equal(t, id, trust.Peer())

	val := 1.5
	trust.SetValue(val)
	require.Equal(t, val, trust.Value())

	t.Run("binary encoding", func(t *testing.T) {
		trust := reputationtest.Trust()
		data, err := trust.Marshal()
		require.NoError(t, err)

		trust2 := reputation.NewTrust()
		require.NoError(t, trust2.Unmarshal(data))
		require.Equal(t, trust, trust2)
	})

	t.Run("JSON encoding", func(t *testing.T) {
		trust := reputationtest.Trust()
		data, err := trust.MarshalJSON()
		require.NoError(t, err)

		trust2 := reputation.NewTrust()
		require.NoError(t, trust2.UnmarshalJSON(data))
		require.Equal(t, trust, trust2)
	})
}

func TestPeerToPeerTrust(t *testing.T) {
	t.Run("v2", func(t *testing.T) {
		p2ptV2 := reputationtestV2.GeneratePeerToPeerTrust(false)

		p2pt := reputation.PeerToPeerTrustFromV2(p2ptV2)

		require.Equal(t, p2ptV2, p2pt.ToV2())
	})

	t.Run("getters+setters", func(t *testing.T) {
		p2pt := reputation.NewPeerToPeerTrust()

		require.Nil(t, p2pt.TrustingPeer())
		require.Nil(t, p2pt.Trust())

		trusting := reputationtest.PeerID()
		p2pt.SetTrustingPeer(trusting)
		require.Equal(t, trusting, p2pt.TrustingPeer())

		trust := reputationtest.Trust()
		p2pt.SetTrust(trust)
		require.Equal(t, trust, p2pt.Trust())
	})

	t.Run("encoding", func(t *testing.T) {
		p2pt := reputationtest.PeerToPeerTrust()

		t.Run("binary", func(t *testing.T) {
			data, err := p2pt.Marshal()
			require.NoError(t, err)

			p2pt2 := reputation.NewPeerToPeerTrust()
			require.NoError(t, p2pt2.Unmarshal(data))
			require.Equal(t, p2pt, p2pt2)
		})

		t.Run("JSON", func(t *testing.T) {
			data, err := p2pt.MarshalJSON()
			require.NoError(t, err)

			p2pt2 := reputation.NewPeerToPeerTrust()
			require.NoError(t, p2pt2.UnmarshalJSON(data))
			require.Equal(t, p2pt, p2pt2)
		})
	})
}

func TestGlobalTrust(t *testing.T) {
	t.Run("v2", func(t *testing.T) {
		gtV2 := reputationtestV2.GenerateGlobalTrust(false)

		gt := reputation.GlobalTrustFromV2(gtV2)

		require.Equal(t, gtV2, gt.ToV2())
	})

	t.Run("getters+setters", func(t *testing.T) {
		gt := reputation.NewGlobalTrust()

		require.Equal(t, version.Current(), *gt.Version())
		require.Nil(t, gt.Manager())
		require.Nil(t, gt.Trust())

		var ver version.Version
		ver.SetMajor(13)
		ver.SetMinor(31)
		gt.SetVersion(&ver)
		require.Equal(t, ver, *gt.Version())

		mngr := reputationtest.PeerID()
		gt.SetManager(mngr)
		require.Equal(t, mngr, gt.Manager())

		trust := reputationtest.Trust()
		gt.SetTrust(trust)
		require.Equal(t, trust, gt.Trust())
	})

	t.Run("sign+verify", func(t *testing.T) {
		gt := reputationtest.SignedGlobalTrust(t)

		err := gt.VerifySignature()
		require.NoError(t, err)
	})

	t.Run("encoding", func(t *testing.T) {
		t.Run("binary", func(t *testing.T) {
			gt := reputationtest.SignedGlobalTrust(t)

			data, err := gt.Marshal()
			require.NoError(t, err)

			gt2 := reputation.NewGlobalTrust()
			require.NoError(t, gt2.Unmarshal(data))
			require.Equal(t, gt, gt2)
		})

		t.Run("JSON", func(t *testing.T) {
			gt := reputationtest.SignedGlobalTrust(t)
			data, err := gt.MarshalJSON()
			require.NoError(t, err)

			gt2 := reputation.NewGlobalTrust()
			require.NoError(t, gt2.UnmarshalJSON(data))
			require.Equal(t, gt, gt2)
		})
	})
}

func TestTrustFromV2(t *testing.T) {
	t.Run("from nil", func(t *testing.T) {
		var x *reputationV2.Trust

		require.Nil(t, reputation.TrustFromV2(x))
	})
}

func TestPeerToPeerTrustFromV2(t *testing.T) {
	t.Run("from nil", func(t *testing.T) {
		var x *reputationV2.PeerToPeerTrust

		require.Nil(t, reputation.PeerToPeerTrustFromV2(x))
	})
}

func TestGlobalTrustFromV2(t *testing.T) {
	t.Run("from nil", func(t *testing.T) {
		var x *reputationV2.GlobalTrust

		require.Nil(t, reputation.GlobalTrustFromV2(x))
	})
}

func TestTrust_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *reputation.Trust

		require.Nil(t, x.ToV2())
	})
}

func TestPeerToPeerTrust_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *reputation.PeerToPeerTrust

		require.Nil(t, x.ToV2())
	})
}

func TestGlobalTrust_ToV2(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var x *reputation.GlobalTrust

		require.Nil(t, x.ToV2())
	})
}

func TestNewTrust(t *testing.T) {
	t.Run("default values", func(t *testing.T) {
		trust := reputation.NewTrust()

		// check initial values
		require.Zero(t, trust.Value())
		require.Nil(t, trust.Peer())

		// convert to v2 message
		trustV2 := trust.ToV2()

		require.Zero(t, trustV2.GetValue())
		require.Nil(t, trustV2.GetPeer())
	})
}

func TestNewPeerToPeerTrust(t *testing.T) {
	t.Run("default values", func(t *testing.T) {
		trust := reputation.NewPeerToPeerTrust()

		// check initial values
		require.Nil(t, trust.Trust())
		require.Nil(t, trust.TrustingPeer())

		// convert to v2 message
		trustV2 := trust.ToV2()

		require.Nil(t, trustV2.GetTrust())
		require.Nil(t, trustV2.GetTrustingPeer())
	})
}

func TestNewGlobalTrust(t *testing.T) {
	t.Run("default values", func(t *testing.T) {
		trust := reputation.NewGlobalTrust()

		// check initial values
		require.Nil(t, trust.Manager())
		require.Nil(t, trust.Trust())

		require.Equal(t, version.Current(), *trust.Version())

		// convert to v2 message
		trustV2 := trust.ToV2()

		require.Nil(t, trustV2.GetBody())
	})
}

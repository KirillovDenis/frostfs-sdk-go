package reputation

import (
	"crypto/ecdsa"
	"errors"
	"fmt"

	"github.com/nspcc-dev/neofs-api-go/v2/refs"
	"github.com/nspcc-dev/neofs-api-go/v2/reputation"
	neofscrypto "github.com/nspcc-dev/neofs-sdk-go/crypto"
	neofsecdsa "github.com/nspcc-dev/neofs-sdk-go/crypto/ecdsa"
	"github.com/nspcc-dev/neofs-sdk-go/version"
)

// Trust represents peer's trust compatible with NeoFS API v2.
type Trust reputation.Trust

// NewTrust creates and returns blank Trust.
//
// Defaults:
//  - value: 0;
//  - PeerID: nil.
func NewTrust() *Trust {
	return TrustFromV2(new(reputation.Trust))
}

// TrustFromV2 converts NeoFS API v2
// reputation.Trust message structure to Trust.
//
// Nil reputation.Trust converts to nil.
func TrustFromV2(t *reputation.Trust) *Trust {
	return (*Trust)(t)
}

// ToV2 converts Trust to NeoFS API v2
// reputation.Trust message structure.
//
// Nil Trust converts to nil.
func (x *Trust) ToV2() *reputation.Trust {
	return (*reputation.Trust)(x)
}

// TrustsToV2 converts slice of Trust's to slice of
// NeoFS API v2 reputation.Trust message structures.
func TrustsToV2(xs []Trust) (res []reputation.Trust) {
	if xs != nil {
		res = make([]reputation.Trust, len(xs))

		for i := range xs {
			res[i] = *xs[i].ToV2()
		}
	}

	return
}

// SetPeer sets trusted peer ID.
func (x *Trust) SetPeer(id *PeerID) {
	(*reputation.Trust)(x).SetPeer(id.ToV2())
}

// Peer returns trusted peer ID.
func (x *Trust) Peer() *PeerID {
	return PeerIDFromV2(
		(*reputation.Trust)(x).GetPeer())
}

// SetValue sets trust value.
func (x *Trust) SetValue(val float64) {
	(*reputation.Trust)(x).SetValue(val)
}

// Value returns trust value.
func (x *Trust) Value() float64 {
	return (*reputation.Trust)(x).GetValue()
}

// Marshal marshals Trust into a protobuf binary form.
func (x *Trust) Marshal() ([]byte, error) {
	return (*reputation.Trust)(x).StableMarshal(nil)
}

// Unmarshal unmarshals protobuf binary representation of Trust.
func (x *Trust) Unmarshal(data []byte) error {
	return (*reputation.Trust)(x).Unmarshal(data)
}

// MarshalJSON encodes Trust to protobuf JSON format.
func (x *Trust) MarshalJSON() ([]byte, error) {
	return (*reputation.Trust)(x).MarshalJSON()
}

// UnmarshalJSON decodes Trust from protobuf JSON format.
func (x *Trust) UnmarshalJSON(data []byte) error {
	return (*reputation.Trust)(x).UnmarshalJSON(data)
}

// PeerToPeerTrust represents directed peer-to-peer trust
// compatible with NeoFS API v2.
type PeerToPeerTrust reputation.PeerToPeerTrust

// NewPeerToPeerTrust creates and returns blank PeerToPeerTrust.
//
// Defaults:
//  - trusting: nil;
//  - trust: nil.
func NewPeerToPeerTrust() *PeerToPeerTrust {
	return PeerToPeerTrustFromV2(new(reputation.PeerToPeerTrust))
}

// PeerToPeerTrustFromV2 converts NeoFS API v2
// reputation.PeerToPeerTrust message structure to PeerToPeerTrust.
//
// Nil reputation.PeerToPeerTrust converts to nil.
func PeerToPeerTrustFromV2(t *reputation.PeerToPeerTrust) *PeerToPeerTrust {
	return (*PeerToPeerTrust)(t)
}

// ToV2 converts PeerToPeerTrust to NeoFS API v2
// reputation.PeerToPeerTrust message structure.
//
// Nil PeerToPeerTrust converts to nil.
func (x *PeerToPeerTrust) ToV2() *reputation.PeerToPeerTrust {
	return (*reputation.PeerToPeerTrust)(x)
}

// SetTrustingPeer sets trusting peer ID.
func (x *PeerToPeerTrust) SetTrustingPeer(id *PeerID) {
	(*reputation.PeerToPeerTrust)(x).SetTrustingPeer(id.ToV2())
}

// TrustingPeer returns trusting peer ID.
func (x *PeerToPeerTrust) TrustingPeer() *PeerID {
	return PeerIDFromV2(
		(*reputation.PeerToPeerTrust)(x).
			GetTrustingPeer(),
	)
}

// SetTrust sets trust value of the trusting peer to the trusted one.
func (x *PeerToPeerTrust) SetTrust(t *Trust) {
	(*reputation.PeerToPeerTrust)(x).SetTrust(t.ToV2())
}

// Trust returns trust value of the trusting peer to the trusted one.
func (x *PeerToPeerTrust) Trust() *Trust {
	return TrustFromV2(
		(*reputation.PeerToPeerTrust)(x).GetTrust())
}

// Marshal marshals PeerToPeerTrust into a protobuf binary form.
func (x *PeerToPeerTrust) Marshal() ([]byte, error) {
	return (*reputation.PeerToPeerTrust)(x).StableMarshal(nil)
}

// Unmarshal unmarshals protobuf binary representation of PeerToPeerTrust.
func (x *PeerToPeerTrust) Unmarshal(data []byte) error {
	return (*reputation.PeerToPeerTrust)(x).Unmarshal(data)
}

// MarshalJSON encodes PeerToPeerTrust to protobuf JSON format.
func (x *PeerToPeerTrust) MarshalJSON() ([]byte, error) {
	return (*reputation.PeerToPeerTrust)(x).MarshalJSON()
}

// UnmarshalJSON decodes PeerToPeerTrust from protobuf JSON format.
func (x *PeerToPeerTrust) UnmarshalJSON(data []byte) error {
	return (*reputation.PeerToPeerTrust)(x).UnmarshalJSON(data)
}

// GlobalTrust represents peer's global trust compatible with NeoFS API v2.
type GlobalTrust reputation.GlobalTrust

// NewGlobalTrust creates and returns blank GlobalTrust.
//
// Defaults:
// 	- version: version.Current();
//  - manager: nil;
//  - trust: nil.
func NewGlobalTrust() *GlobalTrust {
	gt := GlobalTrustFromV2(new(reputation.GlobalTrust))
	ver := version.Current()
	gt.SetVersion(&ver)

	return gt
}

// GlobalTrustFromV2 converts NeoFS API v2
// reputation.GlobalTrust message structure to GlobalTrust.
//
// Nil reputation.GlobalTrust converts to nil.
func GlobalTrustFromV2(t *reputation.GlobalTrust) *GlobalTrust {
	return (*GlobalTrust)(t)
}

// ToV2 converts GlobalTrust to NeoFS API v2
// reputation.GlobalTrust message structure.
//
// Nil GlobalTrust converts to nil.
func (x *GlobalTrust) ToV2() *reputation.GlobalTrust {
	return (*reputation.GlobalTrust)(x)
}

// SetVersion sets GlobalTrust's protocol version.
func (x *GlobalTrust) SetVersion(version *version.Version) {
	var verV2 refs.Version
	version.WriteToV2(&verV2)
	(*reputation.GlobalTrust)(x).SetVersion(&verV2)
}

// Version returns GlobalTrust's protocol version.
func (x *GlobalTrust) Version() *version.Version {
	var ver version.Version
	if verV2 := (*reputation.GlobalTrust)(x).GetVersion(); verV2 != nil {
		ver.ReadFromV2(*verV2)
	}
	return &ver
}

func (x *GlobalTrust) setBodyField(setter func(*reputation.GlobalTrustBody)) {
	if x != nil {
		v2 := (*reputation.GlobalTrust)(x)

		body := v2.GetBody()
		if body == nil {
			body = new(reputation.GlobalTrustBody)
			v2.SetBody(body)
		}

		setter(body)
	}
}

// SetManager sets node manager ID.
func (x *GlobalTrust) SetManager(id *PeerID) {
	x.setBodyField(func(body *reputation.GlobalTrustBody) {
		body.SetManager(id.ToV2())
	})
}

// Manager returns node manager ID.
func (x *GlobalTrust) Manager() *PeerID {
	return PeerIDFromV2(
		(*reputation.GlobalTrust)(x).
			GetBody().
			GetManager(),
	)
}

// SetTrust sets global trust value.
func (x *GlobalTrust) SetTrust(trust *Trust) {
	x.setBodyField(func(body *reputation.GlobalTrustBody) {
		body.SetTrust(trust.ToV2())
	})
}

// Trust returns global trust value.
func (x *GlobalTrust) Trust() *Trust {
	return TrustFromV2(
		(*reputation.GlobalTrust)(x).
			GetBody().
			GetTrust(),
	)
}

// Sign signs global trust value with key.
func (x *GlobalTrust) Sign(key *ecdsa.PrivateKey) error {
	if key == nil {
		return errors.New("nil private key")
	}

	m := (*reputation.GlobalTrust)(x)

	data, err := m.GetBody().StableMarshal(nil)
	if err != nil {
		return fmt.Errorf("marshal body: %w", err)
	}

	var sig neofscrypto.Signature

	err = sig.Calculate(neofsecdsa.Signer(*key), data)
	if err != nil {
		return fmt.Errorf("calculate signature: %w", err)
	}

	var sigv2 refs.Signature

	sig.WriteToV2(&sigv2)

	m.SetSignature(&sigv2)

	return nil
}

// VerifySignature verifies global trust signature.
func (x *GlobalTrust) VerifySignature() error {
	m := (*reputation.GlobalTrust)(x)

	sigV2 := m.GetSignature()
	if sigV2 == nil {
		return errors.New("missing signature")
	}

	data, err := m.GetBody().StableMarshal(nil)
	if err != nil {
		return fmt.Errorf("marshal body: %w", err)
	}

	var sig neofscrypto.Signature
	sig.ReadFromV2(*sigV2)

	if !sig.Verify(data) {
		return errors.New("wrong signature")
	}

	return nil
}

// Marshal marshals GlobalTrust into a protobuf binary form.
func (x *GlobalTrust) Marshal() ([]byte, error) {
	return (*reputation.GlobalTrust)(x).StableMarshal(nil)
}

// Unmarshal unmarshals protobuf binary representation of GlobalTrust.
func (x *GlobalTrust) Unmarshal(data []byte) error {
	return (*reputation.GlobalTrust)(x).Unmarshal(data)
}

// MarshalJSON encodes GlobalTrust to protobuf JSON format.
func (x *GlobalTrust) MarshalJSON() ([]byte, error) {
	return (*reputation.GlobalTrust)(x).MarshalJSON()
}

// UnmarshalJSON decodes GlobalTrust from protobuf JSON format.
func (x *GlobalTrust) UnmarshalJSON(data []byte) error {
	return (*reputation.GlobalTrust)(x).UnmarshalJSON(data)
}

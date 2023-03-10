// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package crypto

import (
	"crypto"
	"fmt"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that crypto.Signer is mocked completely
var _ crypto.Signer = (*MoqSigner_mock)(nil)

// MoqSigner holds the state of a moq of the Signer type
type MoqSigner struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSigner_mock

	ResultsByParams_Public []MoqSigner_Public_resultsByParams
	ResultsByParams_Sign   []MoqSigner_Sign_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Public struct{}
			Sign   struct {
				Rand   moq.ParamIndexing
				Digest moq.ParamIndexing
				Opts   moq.ParamIndexing
			}
		}
	}
	// MoqSigner_mock isolates the mock interface of the Signer type
}

type MoqSigner_mock struct {
	Moq *MoqSigner
}

// MoqSigner_recorder isolates the recorder interface of the Signer type
type MoqSigner_recorder struct {
	Moq *MoqSigner
}

// MoqSigner_Public_params holds the params of the Signer type
type MoqSigner_Public_params struct{}

// MoqSigner_Public_paramsKey holds the map key params of the Signer type
type MoqSigner_Public_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqSigner_Public_resultsByParams contains the results for a given set of
// parameters for the Signer type
type MoqSigner_Public_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSigner_Public_paramsKey]*MoqSigner_Public_results
}

// MoqSigner_Public_doFn defines the type of function needed when calling AndDo
// for the Signer type
type MoqSigner_Public_doFn func()

// MoqSigner_Public_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Signer type
type MoqSigner_Public_doReturnFn func() crypto.PublicKey

// MoqSigner_Public_results holds the results of the Signer type
type MoqSigner_Public_results struct {
	Params  MoqSigner_Public_params
	Results []struct {
		Values *struct {
			Result1 crypto.PublicKey
		}
		Sequence   uint32
		DoFn       MoqSigner_Public_doFn
		DoReturnFn MoqSigner_Public_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSigner_Public_fnRecorder routes recorded function calls to the MoqSigner
// moq
type MoqSigner_Public_fnRecorder struct {
	Params    MoqSigner_Public_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSigner_Public_results
	Moq       *MoqSigner
}

// MoqSigner_Public_anyParams isolates the any params functions of the Signer
// type
type MoqSigner_Public_anyParams struct {
	Recorder *MoqSigner_Public_fnRecorder
}

// MoqSigner_Sign_params holds the params of the Signer type
type MoqSigner_Sign_params struct {
	Rand   io.Reader
	Digest []byte
	Opts   crypto.SignerOpts
}

// MoqSigner_Sign_paramsKey holds the map key params of the Signer type
type MoqSigner_Sign_paramsKey struct {
	Params struct {
		Rand io.Reader
		Opts crypto.SignerOpts
	}
	Hashes struct {
		Rand   hash.Hash
		Digest hash.Hash
		Opts   hash.Hash
	}
}

// MoqSigner_Sign_resultsByParams contains the results for a given set of
// parameters for the Signer type
type MoqSigner_Sign_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSigner_Sign_paramsKey]*MoqSigner_Sign_results
}

// MoqSigner_Sign_doFn defines the type of function needed when calling AndDo
// for the Signer type
type MoqSigner_Sign_doFn func(rand io.Reader, digest []byte, opts crypto.SignerOpts)

// MoqSigner_Sign_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Signer type
type MoqSigner_Sign_doReturnFn func(rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error)

// MoqSigner_Sign_results holds the results of the Signer type
type MoqSigner_Sign_results struct {
	Params  MoqSigner_Sign_params
	Results []struct {
		Values *struct {
			Signature []byte
			Err       error
		}
		Sequence   uint32
		DoFn       MoqSigner_Sign_doFn
		DoReturnFn MoqSigner_Sign_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSigner_Sign_fnRecorder routes recorded function calls to the MoqSigner
// moq
type MoqSigner_Sign_fnRecorder struct {
	Params    MoqSigner_Sign_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSigner_Sign_results
	Moq       *MoqSigner
}

// MoqSigner_Sign_anyParams isolates the any params functions of the Signer
// type
type MoqSigner_Sign_anyParams struct {
	Recorder *MoqSigner_Sign_fnRecorder
}

// NewMoqSigner creates a new moq of the Signer type
func NewMoqSigner(scene *moq.Scene, config *moq.Config) *MoqSigner {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSigner{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSigner_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Public struct{}
				Sign   struct {
					Rand   moq.ParamIndexing
					Digest moq.ParamIndexing
					Opts   moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Public struct{}
			Sign   struct {
				Rand   moq.ParamIndexing
				Digest moq.ParamIndexing
				Opts   moq.ParamIndexing
			}
		}{
			Public: struct{}{},
			Sign: struct {
				Rand   moq.ParamIndexing
				Digest moq.ParamIndexing
				Opts   moq.ParamIndexing
			}{
				Rand:   moq.ParamIndexByHash,
				Digest: moq.ParamIndexByHash,
				Opts:   moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Signer type
func (m *MoqSigner) Mock() *MoqSigner_mock { return m.Moq }

func (m *MoqSigner_mock) Public() (result1 crypto.PublicKey) {
	m.Moq.Scene.T.Helper()
	params := MoqSigner_Public_params{}
	var results *MoqSigner_Public_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Public {
		paramsKey := m.Moq.ParamsKey_Public(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Public(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Public(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Public(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn()
	}
	return
}

func (m *MoqSigner_mock) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqSigner_Sign_params{
		Rand:   rand,
		Digest: digest,
		Opts:   opts,
	}
	var results *MoqSigner_Sign_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Sign {
		paramsKey := m.Moq.ParamsKey_Sign(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Sign(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Sign(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Sign(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(rand, digest, opts)
	}

	if result.Values != nil {
		signature = result.Values.Signature
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		signature, err = result.DoReturnFn(rand, digest, opts)
	}
	return
}

// OnCall returns the recorder implementation of the Signer type
func (m *MoqSigner) OnCall() *MoqSigner_recorder {
	return &MoqSigner_recorder{
		Moq: m,
	}
}

func (m *MoqSigner_recorder) Public() *MoqSigner_Public_fnRecorder {
	return &MoqSigner_Public_fnRecorder{
		Params:   MoqSigner_Public_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqSigner_Public_fnRecorder) Any() *MoqSigner_Public_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Public(r.Params))
		return nil
	}
	return &MoqSigner_Public_anyParams{Recorder: r}
}

func (r *MoqSigner_Public_fnRecorder) Seq() *MoqSigner_Public_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Public(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSigner_Public_fnRecorder) NoSeq() *MoqSigner_Public_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Public(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSigner_Public_fnRecorder) ReturnResults(result1 crypto.PublicKey) *MoqSigner_Public_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 crypto.PublicKey
		}
		Sequence   uint32
		DoFn       MoqSigner_Public_doFn
		DoReturnFn MoqSigner_Public_doReturnFn
	}{
		Values: &struct {
			Result1 crypto.PublicKey
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSigner_Public_fnRecorder) AndDo(fn MoqSigner_Public_doFn) *MoqSigner_Public_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSigner_Public_fnRecorder) DoReturnResults(fn MoqSigner_Public_doReturnFn) *MoqSigner_Public_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 crypto.PublicKey
		}
		Sequence   uint32
		DoFn       MoqSigner_Public_doFn
		DoReturnFn MoqSigner_Public_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSigner_Public_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSigner_Public_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Public {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqSigner_Public_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSigner_Public_paramsKey]*MoqSigner_Public_results{},
		}
		r.Moq.ResultsByParams_Public = append(r.Moq.ResultsByParams_Public, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Public) {
			copy(r.Moq.ResultsByParams_Public[insertAt+1:], r.Moq.ResultsByParams_Public[insertAt:0])
			r.Moq.ResultsByParams_Public[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Public(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqSigner_Public_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSigner_Public_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSigner_Public_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults or DoReturnResults must be called before calling Repeat")
		return nil
	}
	r.Results.Repeat.Repeat(r.Moq.Scene.T, repeaters)
	last := r.Results.Results[len(r.Results.Results)-1]
	for n := 0; n < r.Results.Repeat.ResultCount-1; n++ {
		if r.Sequence {
			last = struct {
				Values *struct {
					Result1 crypto.PublicKey
				}
				Sequence   uint32
				DoFn       MoqSigner_Public_doFn
				DoReturnFn MoqSigner_Public_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSigner) PrettyParams_Public(params MoqSigner_Public_params) string {
	return fmt.Sprintf("Public()")
}

func (m *MoqSigner) ParamsKey_Public(params MoqSigner_Public_params, anyParams uint64) MoqSigner_Public_paramsKey {
	m.Scene.T.Helper()
	return MoqSigner_Public_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqSigner_recorder) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) *MoqSigner_Sign_fnRecorder {
	return &MoqSigner_Sign_fnRecorder{
		Params: MoqSigner_Sign_params{
			Rand:   rand,
			Digest: digest,
			Opts:   opts,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqSigner_Sign_fnRecorder) Any() *MoqSigner_Sign_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Sign(r.Params))
		return nil
	}
	return &MoqSigner_Sign_anyParams{Recorder: r}
}

func (a *MoqSigner_Sign_anyParams) Rand() *MoqSigner_Sign_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqSigner_Sign_anyParams) Digest() *MoqSigner_Sign_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqSigner_Sign_anyParams) Opts() *MoqSigner_Sign_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqSigner_Sign_fnRecorder) Seq() *MoqSigner_Sign_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Sign(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSigner_Sign_fnRecorder) NoSeq() *MoqSigner_Sign_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Sign(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSigner_Sign_fnRecorder) ReturnResults(signature []byte, err error) *MoqSigner_Sign_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Signature []byte
			Err       error
		}
		Sequence   uint32
		DoFn       MoqSigner_Sign_doFn
		DoReturnFn MoqSigner_Sign_doReturnFn
	}{
		Values: &struct {
			Signature []byte
			Err       error
		}{
			Signature: signature,
			Err:       err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSigner_Sign_fnRecorder) AndDo(fn MoqSigner_Sign_doFn) *MoqSigner_Sign_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSigner_Sign_fnRecorder) DoReturnResults(fn MoqSigner_Sign_doReturnFn) *MoqSigner_Sign_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Signature []byte
			Err       error
		}
		Sequence   uint32
		DoFn       MoqSigner_Sign_doFn
		DoReturnFn MoqSigner_Sign_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSigner_Sign_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSigner_Sign_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Sign {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqSigner_Sign_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSigner_Sign_paramsKey]*MoqSigner_Sign_results{},
		}
		r.Moq.ResultsByParams_Sign = append(r.Moq.ResultsByParams_Sign, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Sign) {
			copy(r.Moq.ResultsByParams_Sign[insertAt+1:], r.Moq.ResultsByParams_Sign[insertAt:0])
			r.Moq.ResultsByParams_Sign[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Sign(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqSigner_Sign_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSigner_Sign_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSigner_Sign_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults or DoReturnResults must be called before calling Repeat")
		return nil
	}
	r.Results.Repeat.Repeat(r.Moq.Scene.T, repeaters)
	last := r.Results.Results[len(r.Results.Results)-1]
	for n := 0; n < r.Results.Repeat.ResultCount-1; n++ {
		if r.Sequence {
			last = struct {
				Values *struct {
					Signature []byte
					Err       error
				}
				Sequence   uint32
				DoFn       MoqSigner_Sign_doFn
				DoReturnFn MoqSigner_Sign_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSigner) PrettyParams_Sign(params MoqSigner_Sign_params) string {
	return fmt.Sprintf("Sign(%#v, %#v, %#v)", params.Rand, params.Digest, params.Opts)
}

func (m *MoqSigner) ParamsKey_Sign(params MoqSigner_Sign_params, anyParams uint64) MoqSigner_Sign_paramsKey {
	m.Scene.T.Helper()
	var randUsed io.Reader
	var randUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Sign.Rand == moq.ParamIndexByValue {
			randUsed = params.Rand
		} else {
			randUsedHash = hash.DeepHash(params.Rand)
		}
	}
	var digestUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Sign.Digest == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The digest parameter of the Sign function can't be indexed by value")
		}
		digestUsedHash = hash.DeepHash(params.Digest)
	}
	var optsUsed crypto.SignerOpts
	var optsUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Sign.Opts == moq.ParamIndexByValue {
			optsUsed = params.Opts
		} else {
			optsUsedHash = hash.DeepHash(params.Opts)
		}
	}
	return MoqSigner_Sign_paramsKey{
		Params: struct {
			Rand io.Reader
			Opts crypto.SignerOpts
		}{
			Rand: randUsed,
			Opts: optsUsed,
		},
		Hashes: struct {
			Rand   hash.Hash
			Digest hash.Hash
			Opts   hash.Hash
		}{
			Rand:   randUsedHash,
			Digest: digestUsedHash,
			Opts:   optsUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSigner) Reset() { m.ResultsByParams_Public = nil; m.ResultsByParams_Sign = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSigner) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Public {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Public(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Sign {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Sign(results.Params))
			}
		}
	}
}

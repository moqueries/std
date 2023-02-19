// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package rsa

import (
	"crypto"
	"crypto/rsa"
	"fmt"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// SignPSS_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type SignPSS_genType func(rand io.Reader, priv *rsa.PrivateKey, hash crypto.Hash, hashed []byte, opts *rsa.PSSOptions) ([]byte, error)

// MoqSignPSS_genType holds the state of a moq of the SignPSS_genType type
type MoqSignPSS_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSignPSS_genType_mock

	ResultsByParams []MoqSignPSS_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Rand   moq.ParamIndexing
			Priv   moq.ParamIndexing
			Hash   moq.ParamIndexing
			Hashed moq.ParamIndexing
			Opts   moq.ParamIndexing
		}
	}
}

// MoqSignPSS_genType_mock isolates the mock interface of the SignPSS_genType
// type
type MoqSignPSS_genType_mock struct {
	Moq *MoqSignPSS_genType
}

// MoqSignPSS_genType_params holds the params of the SignPSS_genType type
type MoqSignPSS_genType_params struct {
	Rand   io.Reader
	Priv   *rsa.PrivateKey
	Hash   crypto.Hash
	Hashed []byte
	Opts   *rsa.PSSOptions
}

// MoqSignPSS_genType_paramsKey holds the map key params of the SignPSS_genType
// type
type MoqSignPSS_genType_paramsKey struct {
	Params struct {
		Rand io.Reader
		Priv *rsa.PrivateKey
		Hash crypto.Hash
		Opts *rsa.PSSOptions
	}
	Hashes struct {
		Rand   hash.Hash
		Priv   hash.Hash
		Hash   hash.Hash
		Hashed hash.Hash
		Opts   hash.Hash
	}
}

// MoqSignPSS_genType_resultsByParams contains the results for a given set of
// parameters for the SignPSS_genType type
type MoqSignPSS_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSignPSS_genType_paramsKey]*MoqSignPSS_genType_results
}

// MoqSignPSS_genType_doFn defines the type of function needed when calling
// AndDo for the SignPSS_genType type
type MoqSignPSS_genType_doFn func(rand io.Reader, priv *rsa.PrivateKey, hash crypto.Hash, hashed []byte, opts *rsa.PSSOptions)

// MoqSignPSS_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the SignPSS_genType type
type MoqSignPSS_genType_doReturnFn func(rand io.Reader, priv *rsa.PrivateKey, hash crypto.Hash, hashed []byte, opts *rsa.PSSOptions) ([]byte, error)

// MoqSignPSS_genType_results holds the results of the SignPSS_genType type
type MoqSignPSS_genType_results struct {
	Params  MoqSignPSS_genType_params
	Results []struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqSignPSS_genType_doFn
		DoReturnFn MoqSignPSS_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSignPSS_genType_fnRecorder routes recorded function calls to the
// MoqSignPSS_genType moq
type MoqSignPSS_genType_fnRecorder struct {
	Params    MoqSignPSS_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSignPSS_genType_results
	Moq       *MoqSignPSS_genType
}

// MoqSignPSS_genType_anyParams isolates the any params functions of the
// SignPSS_genType type
type MoqSignPSS_genType_anyParams struct {
	Recorder *MoqSignPSS_genType_fnRecorder
}

// NewMoqSignPSS_genType creates a new moq of the SignPSS_genType type
func NewMoqSignPSS_genType(scene *moq.Scene, config *moq.Config) *MoqSignPSS_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSignPSS_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSignPSS_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Rand   moq.ParamIndexing
				Priv   moq.ParamIndexing
				Hash   moq.ParamIndexing
				Hashed moq.ParamIndexing
				Opts   moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Rand   moq.ParamIndexing
			Priv   moq.ParamIndexing
			Hash   moq.ParamIndexing
			Hashed moq.ParamIndexing
			Opts   moq.ParamIndexing
		}{
			Rand:   moq.ParamIndexByHash,
			Priv:   moq.ParamIndexByHash,
			Hash:   moq.ParamIndexByValue,
			Hashed: moq.ParamIndexByHash,
			Opts:   moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the SignPSS_genType type
func (m *MoqSignPSS_genType) Mock() SignPSS_genType {
	return func(rand io.Reader, priv *rsa.PrivateKey, hash crypto.Hash, hashed []byte, opts *rsa.PSSOptions) ([]byte, error) {
		m.Scene.T.Helper()
		moq := &MoqSignPSS_genType_mock{Moq: m}
		return moq.Fn(rand, priv, hash, hashed, opts)
	}
}

func (m *MoqSignPSS_genType_mock) Fn(rand io.Reader, priv *rsa.PrivateKey, hash crypto.Hash, hashed []byte, opts *rsa.PSSOptions) (result1 []byte, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqSignPSS_genType_params{
		Rand:   rand,
		Priv:   priv,
		Hash:   hash,
		Hashed: hashed,
		Opts:   opts,
	}
	var results *MoqSignPSS_genType_results
	for _, resultsByParams := range m.Moq.ResultsByParams {
		paramsKey := m.Moq.ParamsKey(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(rand, priv, hash, hashed, opts)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(rand, priv, hash, hashed, opts)
	}
	return
}

func (m *MoqSignPSS_genType) OnCall(rand io.Reader, priv *rsa.PrivateKey, hash crypto.Hash, hashed []byte, opts *rsa.PSSOptions) *MoqSignPSS_genType_fnRecorder {
	return &MoqSignPSS_genType_fnRecorder{
		Params: MoqSignPSS_genType_params{
			Rand:   rand,
			Priv:   priv,
			Hash:   hash,
			Hashed: hashed,
			Opts:   opts,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSignPSS_genType_fnRecorder) Any() *MoqSignPSS_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSignPSS_genType_anyParams{Recorder: r}
}

func (a *MoqSignPSS_genType_anyParams) Rand() *MoqSignPSS_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqSignPSS_genType_anyParams) Priv() *MoqSignPSS_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqSignPSS_genType_anyParams) Hash() *MoqSignPSS_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (a *MoqSignPSS_genType_anyParams) Hashed() *MoqSignPSS_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 3
	return a.Recorder
}

func (a *MoqSignPSS_genType_anyParams) Opts() *MoqSignPSS_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 4
	return a.Recorder
}

func (r *MoqSignPSS_genType_fnRecorder) Seq() *MoqSignPSS_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSignPSS_genType_fnRecorder) NoSeq() *MoqSignPSS_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSignPSS_genType_fnRecorder) ReturnResults(result1 []byte, result2 error) *MoqSignPSS_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqSignPSS_genType_doFn
		DoReturnFn MoqSignPSS_genType_doReturnFn
	}{
		Values: &struct {
			Result1 []byte
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSignPSS_genType_fnRecorder) AndDo(fn MoqSignPSS_genType_doFn) *MoqSignPSS_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSignPSS_genType_fnRecorder) DoReturnResults(fn MoqSignPSS_genType_doReturnFn) *MoqSignPSS_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqSignPSS_genType_doFn
		DoReturnFn MoqSignPSS_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSignPSS_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSignPSS_genType_resultsByParams
	for n, res := range r.Moq.ResultsByParams {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqSignPSS_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSignPSS_genType_paramsKey]*MoqSignPSS_genType_results{},
		}
		r.Moq.ResultsByParams = append(r.Moq.ResultsByParams, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams) {
			copy(r.Moq.ResultsByParams[insertAt+1:], r.Moq.ResultsByParams[insertAt:0])
			r.Moq.ResultsByParams[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqSignPSS_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSignPSS_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSignPSS_genType_fnRecorder {
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
					Result1 []byte
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqSignPSS_genType_doFn
				DoReturnFn MoqSignPSS_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSignPSS_genType) PrettyParams(params MoqSignPSS_genType_params) string {
	return fmt.Sprintf("SignPSS_genType(%#v, %#v, %#v, %#v, %#v)", params.Rand, params.Priv, params.Hash, params.Hashed, params.Opts)
}

func (m *MoqSignPSS_genType) ParamsKey(params MoqSignPSS_genType_params, anyParams uint64) MoqSignPSS_genType_paramsKey {
	m.Scene.T.Helper()
	var randUsed io.Reader
	var randUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Rand == moq.ParamIndexByValue {
			randUsed = params.Rand
		} else {
			randUsedHash = hash.DeepHash(params.Rand)
		}
	}
	var privUsed *rsa.PrivateKey
	var privUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Priv == moq.ParamIndexByValue {
			privUsed = params.Priv
		} else {
			privUsedHash = hash.DeepHash(params.Priv)
		}
	}
	var hashUsed crypto.Hash
	var hashUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Hash == moq.ParamIndexByValue {
			hashUsed = params.Hash
		} else {
			hashUsedHash = hash.DeepHash(params.Hash)
		}
	}
	var hashedUsedHash hash.Hash
	if anyParams&(1<<3) == 0 {
		if m.Runtime.ParameterIndexing.Hashed == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The hashed parameter can't be indexed by value")
		}
		hashedUsedHash = hash.DeepHash(params.Hashed)
	}
	var optsUsed *rsa.PSSOptions
	var optsUsedHash hash.Hash
	if anyParams&(1<<4) == 0 {
		if m.Runtime.ParameterIndexing.Opts == moq.ParamIndexByValue {
			optsUsed = params.Opts
		} else {
			optsUsedHash = hash.DeepHash(params.Opts)
		}
	}
	return MoqSignPSS_genType_paramsKey{
		Params: struct {
			Rand io.Reader
			Priv *rsa.PrivateKey
			Hash crypto.Hash
			Opts *rsa.PSSOptions
		}{
			Rand: randUsed,
			Priv: privUsed,
			Hash: hashUsed,
			Opts: optsUsed,
		},
		Hashes: struct {
			Rand   hash.Hash
			Priv   hash.Hash
			Hash   hash.Hash
			Hashed hash.Hash
			Opts   hash.Hash
		}{
			Rand:   randUsedHash,
			Priv:   privUsedHash,
			Hash:   hashUsedHash,
			Hashed: hashedUsedHash,
			Opts:   optsUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSignPSS_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSignPSS_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams(results.Params))
			}
		}
	}
}

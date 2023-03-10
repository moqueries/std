// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package ed25519

import (
	"crypto/ed25519"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// NewKeyFromSeed_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type NewKeyFromSeed_genType func(seed []byte) ed25519.PrivateKey

// MoqNewKeyFromSeed_genType holds the state of a moq of the
// NewKeyFromSeed_genType type
type MoqNewKeyFromSeed_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNewKeyFromSeed_genType_mock

	ResultsByParams []MoqNewKeyFromSeed_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Seed moq.ParamIndexing
		}
	}
}

// MoqNewKeyFromSeed_genType_mock isolates the mock interface of the
// NewKeyFromSeed_genType type
type MoqNewKeyFromSeed_genType_mock struct {
	Moq *MoqNewKeyFromSeed_genType
}

// MoqNewKeyFromSeed_genType_params holds the params of the
// NewKeyFromSeed_genType type
type MoqNewKeyFromSeed_genType_params struct{ Seed []byte }

// MoqNewKeyFromSeed_genType_paramsKey holds the map key params of the
// NewKeyFromSeed_genType type
type MoqNewKeyFromSeed_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ Seed hash.Hash }
}

// MoqNewKeyFromSeed_genType_resultsByParams contains the results for a given
// set of parameters for the NewKeyFromSeed_genType type
type MoqNewKeyFromSeed_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNewKeyFromSeed_genType_paramsKey]*MoqNewKeyFromSeed_genType_results
}

// MoqNewKeyFromSeed_genType_doFn defines the type of function needed when
// calling AndDo for the NewKeyFromSeed_genType type
type MoqNewKeyFromSeed_genType_doFn func(seed []byte)

// MoqNewKeyFromSeed_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the NewKeyFromSeed_genType type
type MoqNewKeyFromSeed_genType_doReturnFn func(seed []byte) ed25519.PrivateKey

// MoqNewKeyFromSeed_genType_results holds the results of the
// NewKeyFromSeed_genType type
type MoqNewKeyFromSeed_genType_results struct {
	Params  MoqNewKeyFromSeed_genType_params
	Results []struct {
		Values *struct {
			Result1 ed25519.PrivateKey
		}
		Sequence   uint32
		DoFn       MoqNewKeyFromSeed_genType_doFn
		DoReturnFn MoqNewKeyFromSeed_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNewKeyFromSeed_genType_fnRecorder routes recorded function calls to the
// MoqNewKeyFromSeed_genType moq
type MoqNewKeyFromSeed_genType_fnRecorder struct {
	Params    MoqNewKeyFromSeed_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNewKeyFromSeed_genType_results
	Moq       *MoqNewKeyFromSeed_genType
}

// MoqNewKeyFromSeed_genType_anyParams isolates the any params functions of the
// NewKeyFromSeed_genType type
type MoqNewKeyFromSeed_genType_anyParams struct {
	Recorder *MoqNewKeyFromSeed_genType_fnRecorder
}

// NewMoqNewKeyFromSeed_genType creates a new moq of the NewKeyFromSeed_genType
// type
func NewMoqNewKeyFromSeed_genType(scene *moq.Scene, config *moq.Config) *MoqNewKeyFromSeed_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNewKeyFromSeed_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNewKeyFromSeed_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Seed moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Seed moq.ParamIndexing
		}{
			Seed: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the NewKeyFromSeed_genType type
func (m *MoqNewKeyFromSeed_genType) Mock() NewKeyFromSeed_genType {
	return func(seed []byte) ed25519.PrivateKey {
		m.Scene.T.Helper()
		moq := &MoqNewKeyFromSeed_genType_mock{Moq: m}
		return moq.Fn(seed)
	}
}

func (m *MoqNewKeyFromSeed_genType_mock) Fn(seed []byte) (result1 ed25519.PrivateKey) {
	m.Moq.Scene.T.Helper()
	params := MoqNewKeyFromSeed_genType_params{
		Seed: seed,
	}
	var results *MoqNewKeyFromSeed_genType_results
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
		result.DoFn(seed)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(seed)
	}
	return
}

func (m *MoqNewKeyFromSeed_genType) OnCall(seed []byte) *MoqNewKeyFromSeed_genType_fnRecorder {
	return &MoqNewKeyFromSeed_genType_fnRecorder{
		Params: MoqNewKeyFromSeed_genType_params{
			Seed: seed,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNewKeyFromSeed_genType_fnRecorder) Any() *MoqNewKeyFromSeed_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNewKeyFromSeed_genType_anyParams{Recorder: r}
}

func (a *MoqNewKeyFromSeed_genType_anyParams) Seed() *MoqNewKeyFromSeed_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqNewKeyFromSeed_genType_fnRecorder) Seq() *MoqNewKeyFromSeed_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNewKeyFromSeed_genType_fnRecorder) NoSeq() *MoqNewKeyFromSeed_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNewKeyFromSeed_genType_fnRecorder) ReturnResults(result1 ed25519.PrivateKey) *MoqNewKeyFromSeed_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 ed25519.PrivateKey
		}
		Sequence   uint32
		DoFn       MoqNewKeyFromSeed_genType_doFn
		DoReturnFn MoqNewKeyFromSeed_genType_doReturnFn
	}{
		Values: &struct {
			Result1 ed25519.PrivateKey
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNewKeyFromSeed_genType_fnRecorder) AndDo(fn MoqNewKeyFromSeed_genType_doFn) *MoqNewKeyFromSeed_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNewKeyFromSeed_genType_fnRecorder) DoReturnResults(fn MoqNewKeyFromSeed_genType_doReturnFn) *MoqNewKeyFromSeed_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 ed25519.PrivateKey
		}
		Sequence   uint32
		DoFn       MoqNewKeyFromSeed_genType_doFn
		DoReturnFn MoqNewKeyFromSeed_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNewKeyFromSeed_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNewKeyFromSeed_genType_resultsByParams
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
		results = &MoqNewKeyFromSeed_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNewKeyFromSeed_genType_paramsKey]*MoqNewKeyFromSeed_genType_results{},
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
		r.Results = &MoqNewKeyFromSeed_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNewKeyFromSeed_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNewKeyFromSeed_genType_fnRecorder {
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
					Result1 ed25519.PrivateKey
				}
				Sequence   uint32
				DoFn       MoqNewKeyFromSeed_genType_doFn
				DoReturnFn MoqNewKeyFromSeed_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNewKeyFromSeed_genType) PrettyParams(params MoqNewKeyFromSeed_genType_params) string {
	return fmt.Sprintf("NewKeyFromSeed_genType(%#v)", params.Seed)
}

func (m *MoqNewKeyFromSeed_genType) ParamsKey(params MoqNewKeyFromSeed_genType_params, anyParams uint64) MoqNewKeyFromSeed_genType_paramsKey {
	m.Scene.T.Helper()
	var seedUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Seed == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The seed parameter can't be indexed by value")
		}
		seedUsedHash = hash.DeepHash(params.Seed)
	}
	return MoqNewKeyFromSeed_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ Seed hash.Hash }{
			Seed: seedUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNewKeyFromSeed_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNewKeyFromSeed_genType) AssertExpectationsMet() {
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

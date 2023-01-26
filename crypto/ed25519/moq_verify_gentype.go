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

// Verify_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Verify_genType func(publicKey ed25519.PublicKey, message, sig []byte) bool

// MoqVerify_genType holds the state of a moq of the Verify_genType type
type MoqVerify_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqVerify_genType_mock

	ResultsByParams []MoqVerify_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			PublicKey moq.ParamIndexing
			Message   moq.ParamIndexing
			Sig       moq.ParamIndexing
		}
	}
}

// MoqVerify_genType_mock isolates the mock interface of the Verify_genType
// type
type MoqVerify_genType_mock struct {
	Moq *MoqVerify_genType
}

// MoqVerify_genType_params holds the params of the Verify_genType type
type MoqVerify_genType_params struct {
	PublicKey    ed25519.PublicKey
	Message, Sig []byte
}

// MoqVerify_genType_paramsKey holds the map key params of the Verify_genType
// type
type MoqVerify_genType_paramsKey struct {
	Params struct{}
	Hashes struct {
		PublicKey    hash.Hash
		Message, Sig hash.Hash
	}
}

// MoqVerify_genType_resultsByParams contains the results for a given set of
// parameters for the Verify_genType type
type MoqVerify_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqVerify_genType_paramsKey]*MoqVerify_genType_results
}

// MoqVerify_genType_doFn defines the type of function needed when calling
// AndDo for the Verify_genType type
type MoqVerify_genType_doFn func(publicKey ed25519.PublicKey, message, sig []byte)

// MoqVerify_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Verify_genType type
type MoqVerify_genType_doReturnFn func(publicKey ed25519.PublicKey, message, sig []byte) bool

// MoqVerify_genType_results holds the results of the Verify_genType type
type MoqVerify_genType_results struct {
	Params  MoqVerify_genType_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqVerify_genType_doFn
		DoReturnFn MoqVerify_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqVerify_genType_fnRecorder routes recorded function calls to the
// MoqVerify_genType moq
type MoqVerify_genType_fnRecorder struct {
	Params    MoqVerify_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqVerify_genType_results
	Moq       *MoqVerify_genType
}

// MoqVerify_genType_anyParams isolates the any params functions of the
// Verify_genType type
type MoqVerify_genType_anyParams struct {
	Recorder *MoqVerify_genType_fnRecorder
}

// NewMoqVerify_genType creates a new moq of the Verify_genType type
func NewMoqVerify_genType(scene *moq.Scene, config *moq.Config) *MoqVerify_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqVerify_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqVerify_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				PublicKey moq.ParamIndexing
				Message   moq.ParamIndexing
				Sig       moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			PublicKey moq.ParamIndexing
			Message   moq.ParamIndexing
			Sig       moq.ParamIndexing
		}{
			PublicKey: moq.ParamIndexByHash,
			Message:   moq.ParamIndexByHash,
			Sig:       moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Verify_genType type
func (m *MoqVerify_genType) Mock() Verify_genType {
	return func(publicKey ed25519.PublicKey, message, sig []byte) bool {
		m.Scene.T.Helper()
		moq := &MoqVerify_genType_mock{Moq: m}
		return moq.Fn(publicKey, message, sig)
	}
}

func (m *MoqVerify_genType_mock) Fn(publicKey ed25519.PublicKey, message, sig []byte) (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqVerify_genType_params{
		PublicKey: publicKey,
		Message:   message,
		Sig:       sig,
	}
	var results *MoqVerify_genType_results
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
		result.DoFn(publicKey, message, sig)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(publicKey, message, sig)
	}
	return
}

func (m *MoqVerify_genType) OnCall(publicKey ed25519.PublicKey, message, sig []byte) *MoqVerify_genType_fnRecorder {
	return &MoqVerify_genType_fnRecorder{
		Params: MoqVerify_genType_params{
			PublicKey: publicKey,
			Message:   message,
			Sig:       sig,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqVerify_genType_fnRecorder) Any() *MoqVerify_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqVerify_genType_anyParams{Recorder: r}
}

func (a *MoqVerify_genType_anyParams) PublicKey() *MoqVerify_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqVerify_genType_anyParams) Message() *MoqVerify_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqVerify_genType_anyParams) Sig() *MoqVerify_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqVerify_genType_fnRecorder) Seq() *MoqVerify_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqVerify_genType_fnRecorder) NoSeq() *MoqVerify_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqVerify_genType_fnRecorder) ReturnResults(result1 bool) *MoqVerify_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqVerify_genType_doFn
		DoReturnFn MoqVerify_genType_doReturnFn
	}{
		Values: &struct {
			Result1 bool
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqVerify_genType_fnRecorder) AndDo(fn MoqVerify_genType_doFn) *MoqVerify_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqVerify_genType_fnRecorder) DoReturnResults(fn MoqVerify_genType_doReturnFn) *MoqVerify_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqVerify_genType_doFn
		DoReturnFn MoqVerify_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqVerify_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqVerify_genType_resultsByParams
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
		results = &MoqVerify_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqVerify_genType_paramsKey]*MoqVerify_genType_results{},
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
		r.Results = &MoqVerify_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqVerify_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqVerify_genType_fnRecorder {
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
					Result1 bool
				}
				Sequence   uint32
				DoFn       MoqVerify_genType_doFn
				DoReturnFn MoqVerify_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqVerify_genType) PrettyParams(params MoqVerify_genType_params) string {
	return fmt.Sprintf("Verify_genType(%#v, %#v, %#v)", params.PublicKey, params.Message, params.Sig)
}

func (m *MoqVerify_genType) ParamsKey(params MoqVerify_genType_params, anyParams uint64) MoqVerify_genType_paramsKey {
	m.Scene.T.Helper()
	var publicKeyUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.PublicKey == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The publicKey parameter can't be indexed by value")
		}
		publicKeyUsedHash = hash.DeepHash(params.PublicKey)
	}
	var messageUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Message == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The message parameter can't be indexed by value")
		}
		messageUsedHash = hash.DeepHash(params.Message)
	}
	var sigUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Sig == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The sig parameter can't be indexed by value")
		}
		sigUsedHash = hash.DeepHash(params.Sig)
	}
	return MoqVerify_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct {
			PublicKey    hash.Hash
			Message, Sig hash.Hash
		}{
			PublicKey: publicKeyUsedHash,
			Message:   messageUsedHash,
			Sig:       sigUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqVerify_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqVerify_genType) AssertExpectationsMet() {
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

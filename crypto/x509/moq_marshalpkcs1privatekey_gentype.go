// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package x509

import (
	"crypto/rsa"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// MarshalPKCS1PrivateKey_genType is the fabricated implementation type of this
// mock (emitted when mocking functions directly and not from a function type)
type MarshalPKCS1PrivateKey_genType func(key *rsa.PrivateKey) []byte

// MoqMarshalPKCS1PrivateKey_genType holds the state of a moq of the
// MarshalPKCS1PrivateKey_genType type
type MoqMarshalPKCS1PrivateKey_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqMarshalPKCS1PrivateKey_genType_mock

	ResultsByParams []MoqMarshalPKCS1PrivateKey_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Key moq.ParamIndexing
		}
	}
}

// MoqMarshalPKCS1PrivateKey_genType_mock isolates the mock interface of the
// MarshalPKCS1PrivateKey_genType type
type MoqMarshalPKCS1PrivateKey_genType_mock struct {
	Moq *MoqMarshalPKCS1PrivateKey_genType
}

// MoqMarshalPKCS1PrivateKey_genType_params holds the params of the
// MarshalPKCS1PrivateKey_genType type
type MoqMarshalPKCS1PrivateKey_genType_params struct{ Key *rsa.PrivateKey }

// MoqMarshalPKCS1PrivateKey_genType_paramsKey holds the map key params of the
// MarshalPKCS1PrivateKey_genType type
type MoqMarshalPKCS1PrivateKey_genType_paramsKey struct {
	Params struct{ Key *rsa.PrivateKey }
	Hashes struct{ Key hash.Hash }
}

// MoqMarshalPKCS1PrivateKey_genType_resultsByParams contains the results for a
// given set of parameters for the MarshalPKCS1PrivateKey_genType type
type MoqMarshalPKCS1PrivateKey_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqMarshalPKCS1PrivateKey_genType_paramsKey]*MoqMarshalPKCS1PrivateKey_genType_results
}

// MoqMarshalPKCS1PrivateKey_genType_doFn defines the type of function needed
// when calling AndDo for the MarshalPKCS1PrivateKey_genType type
type MoqMarshalPKCS1PrivateKey_genType_doFn func(key *rsa.PrivateKey)

// MoqMarshalPKCS1PrivateKey_genType_doReturnFn defines the type of function
// needed when calling DoReturnResults for the MarshalPKCS1PrivateKey_genType
// type
type MoqMarshalPKCS1PrivateKey_genType_doReturnFn func(key *rsa.PrivateKey) []byte

// MoqMarshalPKCS1PrivateKey_genType_results holds the results of the
// MarshalPKCS1PrivateKey_genType type
type MoqMarshalPKCS1PrivateKey_genType_results struct {
	Params  MoqMarshalPKCS1PrivateKey_genType_params
	Results []struct {
		Values *struct {
			Result1 []byte
		}
		Sequence   uint32
		DoFn       MoqMarshalPKCS1PrivateKey_genType_doFn
		DoReturnFn MoqMarshalPKCS1PrivateKey_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqMarshalPKCS1PrivateKey_genType_fnRecorder routes recorded function calls
// to the MoqMarshalPKCS1PrivateKey_genType moq
type MoqMarshalPKCS1PrivateKey_genType_fnRecorder struct {
	Params    MoqMarshalPKCS1PrivateKey_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqMarshalPKCS1PrivateKey_genType_results
	Moq       *MoqMarshalPKCS1PrivateKey_genType
}

// MoqMarshalPKCS1PrivateKey_genType_anyParams isolates the any params
// functions of the MarshalPKCS1PrivateKey_genType type
type MoqMarshalPKCS1PrivateKey_genType_anyParams struct {
	Recorder *MoqMarshalPKCS1PrivateKey_genType_fnRecorder
}

// NewMoqMarshalPKCS1PrivateKey_genType creates a new moq of the
// MarshalPKCS1PrivateKey_genType type
func NewMoqMarshalPKCS1PrivateKey_genType(scene *moq.Scene, config *moq.Config) *MoqMarshalPKCS1PrivateKey_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqMarshalPKCS1PrivateKey_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqMarshalPKCS1PrivateKey_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Key moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Key moq.ParamIndexing
		}{
			Key: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the MarshalPKCS1PrivateKey_genType
// type
func (m *MoqMarshalPKCS1PrivateKey_genType) Mock() MarshalPKCS1PrivateKey_genType {
	return func(key *rsa.PrivateKey) []byte {
		m.Scene.T.Helper()
		moq := &MoqMarshalPKCS1PrivateKey_genType_mock{Moq: m}
		return moq.Fn(key)
	}
}

func (m *MoqMarshalPKCS1PrivateKey_genType_mock) Fn(key *rsa.PrivateKey) (result1 []byte) {
	m.Moq.Scene.T.Helper()
	params := MoqMarshalPKCS1PrivateKey_genType_params{
		Key: key,
	}
	var results *MoqMarshalPKCS1PrivateKey_genType_results
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
		result.DoFn(key)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(key)
	}
	return
}

func (m *MoqMarshalPKCS1PrivateKey_genType) OnCall(key *rsa.PrivateKey) *MoqMarshalPKCS1PrivateKey_genType_fnRecorder {
	return &MoqMarshalPKCS1PrivateKey_genType_fnRecorder{
		Params: MoqMarshalPKCS1PrivateKey_genType_params{
			Key: key,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqMarshalPKCS1PrivateKey_genType_fnRecorder) Any() *MoqMarshalPKCS1PrivateKey_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqMarshalPKCS1PrivateKey_genType_anyParams{Recorder: r}
}

func (a *MoqMarshalPKCS1PrivateKey_genType_anyParams) Key() *MoqMarshalPKCS1PrivateKey_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqMarshalPKCS1PrivateKey_genType_fnRecorder) Seq() *MoqMarshalPKCS1PrivateKey_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqMarshalPKCS1PrivateKey_genType_fnRecorder) NoSeq() *MoqMarshalPKCS1PrivateKey_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqMarshalPKCS1PrivateKey_genType_fnRecorder) ReturnResults(result1 []byte) *MoqMarshalPKCS1PrivateKey_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []byte
		}
		Sequence   uint32
		DoFn       MoqMarshalPKCS1PrivateKey_genType_doFn
		DoReturnFn MoqMarshalPKCS1PrivateKey_genType_doReturnFn
	}{
		Values: &struct {
			Result1 []byte
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqMarshalPKCS1PrivateKey_genType_fnRecorder) AndDo(fn MoqMarshalPKCS1PrivateKey_genType_doFn) *MoqMarshalPKCS1PrivateKey_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqMarshalPKCS1PrivateKey_genType_fnRecorder) DoReturnResults(fn MoqMarshalPKCS1PrivateKey_genType_doReturnFn) *MoqMarshalPKCS1PrivateKey_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []byte
		}
		Sequence   uint32
		DoFn       MoqMarshalPKCS1PrivateKey_genType_doFn
		DoReturnFn MoqMarshalPKCS1PrivateKey_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqMarshalPKCS1PrivateKey_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqMarshalPKCS1PrivateKey_genType_resultsByParams
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
		results = &MoqMarshalPKCS1PrivateKey_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqMarshalPKCS1PrivateKey_genType_paramsKey]*MoqMarshalPKCS1PrivateKey_genType_results{},
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
		r.Results = &MoqMarshalPKCS1PrivateKey_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqMarshalPKCS1PrivateKey_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqMarshalPKCS1PrivateKey_genType_fnRecorder {
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
				}
				Sequence   uint32
				DoFn       MoqMarshalPKCS1PrivateKey_genType_doFn
				DoReturnFn MoqMarshalPKCS1PrivateKey_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqMarshalPKCS1PrivateKey_genType) PrettyParams(params MoqMarshalPKCS1PrivateKey_genType_params) string {
	return fmt.Sprintf("MarshalPKCS1PrivateKey_genType(%#v)", params.Key)
}

func (m *MoqMarshalPKCS1PrivateKey_genType) ParamsKey(params MoqMarshalPKCS1PrivateKey_genType_params, anyParams uint64) MoqMarshalPKCS1PrivateKey_genType_paramsKey {
	m.Scene.T.Helper()
	var keyUsed *rsa.PrivateKey
	var keyUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Key == moq.ParamIndexByValue {
			keyUsed = params.Key
		} else {
			keyUsedHash = hash.DeepHash(params.Key)
		}
	}
	return MoqMarshalPKCS1PrivateKey_genType_paramsKey{
		Params: struct{ Key *rsa.PrivateKey }{
			Key: keyUsed,
		},
		Hashes: struct{ Key hash.Hash }{
			Key: keyUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqMarshalPKCS1PrivateKey_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqMarshalPKCS1PrivateKey_genType) AssertExpectationsMet() {
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

// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package bytes

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// HasSuffix_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type HasSuffix_genType func(s, suffix []byte) bool

// MoqHasSuffix_genType holds the state of a moq of the HasSuffix_genType type
type MoqHasSuffix_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqHasSuffix_genType_mock

	ResultsByParams []MoqHasSuffix_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			S      moq.ParamIndexing
			Suffix moq.ParamIndexing
		}
	}
}

// MoqHasSuffix_genType_mock isolates the mock interface of the
// HasSuffix_genType type
type MoqHasSuffix_genType_mock struct {
	Moq *MoqHasSuffix_genType
}

// MoqHasSuffix_genType_params holds the params of the HasSuffix_genType type
type MoqHasSuffix_genType_params struct{ S, Suffix []byte }

// MoqHasSuffix_genType_paramsKey holds the map key params of the
// HasSuffix_genType type
type MoqHasSuffix_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ S, Suffix hash.Hash }
}

// MoqHasSuffix_genType_resultsByParams contains the results for a given set of
// parameters for the HasSuffix_genType type
type MoqHasSuffix_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqHasSuffix_genType_paramsKey]*MoqHasSuffix_genType_results
}

// MoqHasSuffix_genType_doFn defines the type of function needed when calling
// AndDo for the HasSuffix_genType type
type MoqHasSuffix_genType_doFn func(s, suffix []byte)

// MoqHasSuffix_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the HasSuffix_genType type
type MoqHasSuffix_genType_doReturnFn func(s, suffix []byte) bool

// MoqHasSuffix_genType_results holds the results of the HasSuffix_genType type
type MoqHasSuffix_genType_results struct {
	Params  MoqHasSuffix_genType_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqHasSuffix_genType_doFn
		DoReturnFn MoqHasSuffix_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqHasSuffix_genType_fnRecorder routes recorded function calls to the
// MoqHasSuffix_genType moq
type MoqHasSuffix_genType_fnRecorder struct {
	Params    MoqHasSuffix_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqHasSuffix_genType_results
	Moq       *MoqHasSuffix_genType
}

// MoqHasSuffix_genType_anyParams isolates the any params functions of the
// HasSuffix_genType type
type MoqHasSuffix_genType_anyParams struct {
	Recorder *MoqHasSuffix_genType_fnRecorder
}

// NewMoqHasSuffix_genType creates a new moq of the HasSuffix_genType type
func NewMoqHasSuffix_genType(scene *moq.Scene, config *moq.Config) *MoqHasSuffix_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqHasSuffix_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqHasSuffix_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				S      moq.ParamIndexing
				Suffix moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			S      moq.ParamIndexing
			Suffix moq.ParamIndexing
		}{
			S:      moq.ParamIndexByHash,
			Suffix: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the HasSuffix_genType type
func (m *MoqHasSuffix_genType) Mock() HasSuffix_genType {
	return func(s, suffix []byte) bool {
		m.Scene.T.Helper()
		moq := &MoqHasSuffix_genType_mock{Moq: m}
		return moq.Fn(s, suffix)
	}
}

func (m *MoqHasSuffix_genType_mock) Fn(s, suffix []byte) (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqHasSuffix_genType_params{
		S:      s,
		Suffix: suffix,
	}
	var results *MoqHasSuffix_genType_results
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
		result.DoFn(s, suffix)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(s, suffix)
	}
	return
}

func (m *MoqHasSuffix_genType) OnCall(s, suffix []byte) *MoqHasSuffix_genType_fnRecorder {
	return &MoqHasSuffix_genType_fnRecorder{
		Params: MoqHasSuffix_genType_params{
			S:      s,
			Suffix: suffix,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqHasSuffix_genType_fnRecorder) Any() *MoqHasSuffix_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqHasSuffix_genType_anyParams{Recorder: r}
}

func (a *MoqHasSuffix_genType_anyParams) S() *MoqHasSuffix_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqHasSuffix_genType_anyParams) Suffix() *MoqHasSuffix_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqHasSuffix_genType_fnRecorder) Seq() *MoqHasSuffix_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqHasSuffix_genType_fnRecorder) NoSeq() *MoqHasSuffix_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqHasSuffix_genType_fnRecorder) ReturnResults(result1 bool) *MoqHasSuffix_genType_fnRecorder {
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
		DoFn       MoqHasSuffix_genType_doFn
		DoReturnFn MoqHasSuffix_genType_doReturnFn
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

func (r *MoqHasSuffix_genType_fnRecorder) AndDo(fn MoqHasSuffix_genType_doFn) *MoqHasSuffix_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqHasSuffix_genType_fnRecorder) DoReturnResults(fn MoqHasSuffix_genType_doReturnFn) *MoqHasSuffix_genType_fnRecorder {
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
		DoFn       MoqHasSuffix_genType_doFn
		DoReturnFn MoqHasSuffix_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqHasSuffix_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqHasSuffix_genType_resultsByParams
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
		results = &MoqHasSuffix_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqHasSuffix_genType_paramsKey]*MoqHasSuffix_genType_results{},
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
		r.Results = &MoqHasSuffix_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqHasSuffix_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqHasSuffix_genType_fnRecorder {
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
				DoFn       MoqHasSuffix_genType_doFn
				DoReturnFn MoqHasSuffix_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqHasSuffix_genType) PrettyParams(params MoqHasSuffix_genType_params) string {
	return fmt.Sprintf("HasSuffix_genType(%#v, %#v)", params.S, params.Suffix)
}

func (m *MoqHasSuffix_genType) ParamsKey(params MoqHasSuffix_genType_params, anyParams uint64) MoqHasSuffix_genType_paramsKey {
	m.Scene.T.Helper()
	var sUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.S == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The s parameter can't be indexed by value")
		}
		sUsedHash = hash.DeepHash(params.S)
	}
	var suffixUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Suffix == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The suffix parameter can't be indexed by value")
		}
		suffixUsedHash = hash.DeepHash(params.Suffix)
	}
	return MoqHasSuffix_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ S, Suffix hash.Hash }{
			S:      sUsedHash,
			Suffix: suffixUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqHasSuffix_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqHasSuffix_genType) AssertExpectationsMet() {
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

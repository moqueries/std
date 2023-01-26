// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package bytes

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// SplitN_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type SplitN_genType func(s, sep []byte, n int) [][]byte

// MoqSplitN_genType holds the state of a moq of the SplitN_genType type
type MoqSplitN_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSplitN_genType_mock

	ResultsByParams []MoqSplitN_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			S   moq.ParamIndexing
			Sep moq.ParamIndexing
			N   moq.ParamIndexing
		}
	}
}

// MoqSplitN_genType_mock isolates the mock interface of the SplitN_genType
// type
type MoqSplitN_genType_mock struct {
	Moq *MoqSplitN_genType
}

// MoqSplitN_genType_params holds the params of the SplitN_genType type
type MoqSplitN_genType_params struct {
	S, Sep []byte
	N      int
}

// MoqSplitN_genType_paramsKey holds the map key params of the SplitN_genType
// type
type MoqSplitN_genType_paramsKey struct {
	Params struct{ N int }
	Hashes struct {
		S, Sep hash.Hash
		N      hash.Hash
	}
}

// MoqSplitN_genType_resultsByParams contains the results for a given set of
// parameters for the SplitN_genType type
type MoqSplitN_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSplitN_genType_paramsKey]*MoqSplitN_genType_results
}

// MoqSplitN_genType_doFn defines the type of function needed when calling
// AndDo for the SplitN_genType type
type MoqSplitN_genType_doFn func(s, sep []byte, n int)

// MoqSplitN_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the SplitN_genType type
type MoqSplitN_genType_doReturnFn func(s, sep []byte, n int) [][]byte

// MoqSplitN_genType_results holds the results of the SplitN_genType type
type MoqSplitN_genType_results struct {
	Params  MoqSplitN_genType_params
	Results []struct {
		Values *struct {
			Result1 [][]byte
		}
		Sequence   uint32
		DoFn       MoqSplitN_genType_doFn
		DoReturnFn MoqSplitN_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSplitN_genType_fnRecorder routes recorded function calls to the
// MoqSplitN_genType moq
type MoqSplitN_genType_fnRecorder struct {
	Params    MoqSplitN_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSplitN_genType_results
	Moq       *MoqSplitN_genType
}

// MoqSplitN_genType_anyParams isolates the any params functions of the
// SplitN_genType type
type MoqSplitN_genType_anyParams struct {
	Recorder *MoqSplitN_genType_fnRecorder
}

// NewMoqSplitN_genType creates a new moq of the SplitN_genType type
func NewMoqSplitN_genType(scene *moq.Scene, config *moq.Config) *MoqSplitN_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSplitN_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSplitN_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				S   moq.ParamIndexing
				Sep moq.ParamIndexing
				N   moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			S   moq.ParamIndexing
			Sep moq.ParamIndexing
			N   moq.ParamIndexing
		}{
			S:   moq.ParamIndexByHash,
			Sep: moq.ParamIndexByHash,
			N:   moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the SplitN_genType type
func (m *MoqSplitN_genType) Mock() SplitN_genType {
	return func(s, sep []byte, n int) [][]byte {
		m.Scene.T.Helper()
		moq := &MoqSplitN_genType_mock{Moq: m}
		return moq.Fn(s, sep, n)
	}
}

func (m *MoqSplitN_genType_mock) Fn(s, sep []byte, n int) (result1 [][]byte) {
	m.Moq.Scene.T.Helper()
	params := MoqSplitN_genType_params{
		S:   s,
		Sep: sep,
		N:   n,
	}
	var results *MoqSplitN_genType_results
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
		result.DoFn(s, sep, n)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(s, sep, n)
	}
	return
}

func (m *MoqSplitN_genType) OnCall(s, sep []byte, n int) *MoqSplitN_genType_fnRecorder {
	return &MoqSplitN_genType_fnRecorder{
		Params: MoqSplitN_genType_params{
			S:   s,
			Sep: sep,
			N:   n,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSplitN_genType_fnRecorder) Any() *MoqSplitN_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSplitN_genType_anyParams{Recorder: r}
}

func (a *MoqSplitN_genType_anyParams) S() *MoqSplitN_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqSplitN_genType_anyParams) Sep() *MoqSplitN_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqSplitN_genType_anyParams) N() *MoqSplitN_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqSplitN_genType_fnRecorder) Seq() *MoqSplitN_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSplitN_genType_fnRecorder) NoSeq() *MoqSplitN_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSplitN_genType_fnRecorder) ReturnResults(result1 [][]byte) *MoqSplitN_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 [][]byte
		}
		Sequence   uint32
		DoFn       MoqSplitN_genType_doFn
		DoReturnFn MoqSplitN_genType_doReturnFn
	}{
		Values: &struct {
			Result1 [][]byte
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSplitN_genType_fnRecorder) AndDo(fn MoqSplitN_genType_doFn) *MoqSplitN_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSplitN_genType_fnRecorder) DoReturnResults(fn MoqSplitN_genType_doReturnFn) *MoqSplitN_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 [][]byte
		}
		Sequence   uint32
		DoFn       MoqSplitN_genType_doFn
		DoReturnFn MoqSplitN_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSplitN_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSplitN_genType_resultsByParams
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
		results = &MoqSplitN_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSplitN_genType_paramsKey]*MoqSplitN_genType_results{},
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
		r.Results = &MoqSplitN_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSplitN_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSplitN_genType_fnRecorder {
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
					Result1 [][]byte
				}
				Sequence   uint32
				DoFn       MoqSplitN_genType_doFn
				DoReturnFn MoqSplitN_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSplitN_genType) PrettyParams(params MoqSplitN_genType_params) string {
	return fmt.Sprintf("SplitN_genType(%#v, %#v, %#v)", params.S, params.Sep, params.N)
}

func (m *MoqSplitN_genType) ParamsKey(params MoqSplitN_genType_params, anyParams uint64) MoqSplitN_genType_paramsKey {
	m.Scene.T.Helper()
	var sUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.S == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The s parameter can't be indexed by value")
		}
		sUsedHash = hash.DeepHash(params.S)
	}
	var sepUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Sep == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The sep parameter can't be indexed by value")
		}
		sepUsedHash = hash.DeepHash(params.Sep)
	}
	var nUsed int
	var nUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.N == moq.ParamIndexByValue {
			nUsed = params.N
		} else {
			nUsedHash = hash.DeepHash(params.N)
		}
	}
	return MoqSplitN_genType_paramsKey{
		Params: struct{ N int }{
			N: nUsed,
		},
		Hashes: struct {
			S, Sep hash.Hash
			N      hash.Hash
		}{
			S:   sUsedHash,
			Sep: sepUsedHash,
			N:   nUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSplitN_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSplitN_genType) AssertExpectationsMet() {
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

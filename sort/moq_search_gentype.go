// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package sort

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Search_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Search_genType func(n int, f func(int) bool) int

// MoqSearch_genType holds the state of a moq of the Search_genType type
type MoqSearch_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSearch_genType_mock

	ResultsByParams []MoqSearch_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			N moq.ParamIndexing
			F moq.ParamIndexing
		}
	}
}

// MoqSearch_genType_mock isolates the mock interface of the Search_genType
// type
type MoqSearch_genType_mock struct {
	Moq *MoqSearch_genType
}

// MoqSearch_genType_params holds the params of the Search_genType type
type MoqSearch_genType_params struct {
	N int
	F func(int) bool
}

// MoqSearch_genType_paramsKey holds the map key params of the Search_genType
// type
type MoqSearch_genType_paramsKey struct {
	Params struct{ N int }
	Hashes struct {
		N hash.Hash
		F hash.Hash
	}
}

// MoqSearch_genType_resultsByParams contains the results for a given set of
// parameters for the Search_genType type
type MoqSearch_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSearch_genType_paramsKey]*MoqSearch_genType_results
}

// MoqSearch_genType_doFn defines the type of function needed when calling
// AndDo for the Search_genType type
type MoqSearch_genType_doFn func(n int, f func(int) bool)

// MoqSearch_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Search_genType type
type MoqSearch_genType_doReturnFn func(n int, f func(int) bool) int

// MoqSearch_genType_results holds the results of the Search_genType type
type MoqSearch_genType_results struct {
	Params  MoqSearch_genType_params
	Results []struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqSearch_genType_doFn
		DoReturnFn MoqSearch_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSearch_genType_fnRecorder routes recorded function calls to the
// MoqSearch_genType moq
type MoqSearch_genType_fnRecorder struct {
	Params    MoqSearch_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSearch_genType_results
	Moq       *MoqSearch_genType
}

// MoqSearch_genType_anyParams isolates the any params functions of the
// Search_genType type
type MoqSearch_genType_anyParams struct {
	Recorder *MoqSearch_genType_fnRecorder
}

// NewMoqSearch_genType creates a new moq of the Search_genType type
func NewMoqSearch_genType(scene *moq.Scene, config *moq.Config) *MoqSearch_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSearch_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSearch_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				N moq.ParamIndexing
				F moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			N moq.ParamIndexing
			F moq.ParamIndexing
		}{
			N: moq.ParamIndexByValue,
			F: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Search_genType type
func (m *MoqSearch_genType) Mock() Search_genType {
	return func(n int, f func(int) bool) int {
		m.Scene.T.Helper()
		moq := &MoqSearch_genType_mock{Moq: m}
		return moq.Fn(n, f)
	}
}

func (m *MoqSearch_genType_mock) Fn(n int, f func(int) bool) (result1 int) {
	m.Moq.Scene.T.Helper()
	params := MoqSearch_genType_params{
		N: n,
		F: f,
	}
	var results *MoqSearch_genType_results
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
		result.DoFn(n, f)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(n, f)
	}
	return
}

func (m *MoqSearch_genType) OnCall(n int, f func(int) bool) *MoqSearch_genType_fnRecorder {
	return &MoqSearch_genType_fnRecorder{
		Params: MoqSearch_genType_params{
			N: n,
			F: f,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSearch_genType_fnRecorder) Any() *MoqSearch_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSearch_genType_anyParams{Recorder: r}
}

func (a *MoqSearch_genType_anyParams) N() *MoqSearch_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqSearch_genType_anyParams) F() *MoqSearch_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqSearch_genType_fnRecorder) Seq() *MoqSearch_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSearch_genType_fnRecorder) NoSeq() *MoqSearch_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSearch_genType_fnRecorder) ReturnResults(result1 int) *MoqSearch_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqSearch_genType_doFn
		DoReturnFn MoqSearch_genType_doReturnFn
	}{
		Values: &struct {
			Result1 int
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSearch_genType_fnRecorder) AndDo(fn MoqSearch_genType_doFn) *MoqSearch_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSearch_genType_fnRecorder) DoReturnResults(fn MoqSearch_genType_doReturnFn) *MoqSearch_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqSearch_genType_doFn
		DoReturnFn MoqSearch_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSearch_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSearch_genType_resultsByParams
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
		results = &MoqSearch_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSearch_genType_paramsKey]*MoqSearch_genType_results{},
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
		r.Results = &MoqSearch_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSearch_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSearch_genType_fnRecorder {
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
					Result1 int
				}
				Sequence   uint32
				DoFn       MoqSearch_genType_doFn
				DoReturnFn MoqSearch_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSearch_genType) PrettyParams(params MoqSearch_genType_params) string {
	return fmt.Sprintf("Search_genType(%#v, %#v)", params.N, moq.FnString(params.F))
}

func (m *MoqSearch_genType) ParamsKey(params MoqSearch_genType_params, anyParams uint64) MoqSearch_genType_paramsKey {
	m.Scene.T.Helper()
	var nUsed int
	var nUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.N == moq.ParamIndexByValue {
			nUsed = params.N
		} else {
			nUsedHash = hash.DeepHash(params.N)
		}
	}
	var fUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.F == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The f parameter can't be indexed by value")
		}
		fUsedHash = hash.DeepHash(params.F)
	}
	return MoqSearch_genType_paramsKey{
		Params: struct{ N int }{
			N: nUsed,
		},
		Hashes: struct {
			N hash.Hash
			F hash.Hash
		}{
			N: nUsedHash,
			F: fUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSearch_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSearch_genType) AssertExpectationsMet() {
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

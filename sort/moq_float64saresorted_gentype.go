// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package sort

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Float64sAreSorted_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type Float64sAreSorted_genType func(a []float64) bool

// MoqFloat64sAreSorted_genType holds the state of a moq of the
// Float64sAreSorted_genType type
type MoqFloat64sAreSorted_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqFloat64sAreSorted_genType_mock

	ResultsByParams []MoqFloat64sAreSorted_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			A moq.ParamIndexing
		}
	}
}

// MoqFloat64sAreSorted_genType_mock isolates the mock interface of the
// Float64sAreSorted_genType type
type MoqFloat64sAreSorted_genType_mock struct {
	Moq *MoqFloat64sAreSorted_genType
}

// MoqFloat64sAreSorted_genType_params holds the params of the
// Float64sAreSorted_genType type
type MoqFloat64sAreSorted_genType_params struct{ A []float64 }

// MoqFloat64sAreSorted_genType_paramsKey holds the map key params of the
// Float64sAreSorted_genType type
type MoqFloat64sAreSorted_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ A hash.Hash }
}

// MoqFloat64sAreSorted_genType_resultsByParams contains the results for a
// given set of parameters for the Float64sAreSorted_genType type
type MoqFloat64sAreSorted_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFloat64sAreSorted_genType_paramsKey]*MoqFloat64sAreSorted_genType_results
}

// MoqFloat64sAreSorted_genType_doFn defines the type of function needed when
// calling AndDo for the Float64sAreSorted_genType type
type MoqFloat64sAreSorted_genType_doFn func(a []float64)

// MoqFloat64sAreSorted_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Float64sAreSorted_genType type
type MoqFloat64sAreSorted_genType_doReturnFn func(a []float64) bool

// MoqFloat64sAreSorted_genType_results holds the results of the
// Float64sAreSorted_genType type
type MoqFloat64sAreSorted_genType_results struct {
	Params  MoqFloat64sAreSorted_genType_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqFloat64sAreSorted_genType_doFn
		DoReturnFn MoqFloat64sAreSorted_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFloat64sAreSorted_genType_fnRecorder routes recorded function calls to
// the MoqFloat64sAreSorted_genType moq
type MoqFloat64sAreSorted_genType_fnRecorder struct {
	Params    MoqFloat64sAreSorted_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFloat64sAreSorted_genType_results
	Moq       *MoqFloat64sAreSorted_genType
}

// MoqFloat64sAreSorted_genType_anyParams isolates the any params functions of
// the Float64sAreSorted_genType type
type MoqFloat64sAreSorted_genType_anyParams struct {
	Recorder *MoqFloat64sAreSorted_genType_fnRecorder
}

// NewMoqFloat64sAreSorted_genType creates a new moq of the
// Float64sAreSorted_genType type
func NewMoqFloat64sAreSorted_genType(scene *moq.Scene, config *moq.Config) *MoqFloat64sAreSorted_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqFloat64sAreSorted_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqFloat64sAreSorted_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				A moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			A moq.ParamIndexing
		}{
			A: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Float64sAreSorted_genType type
func (m *MoqFloat64sAreSorted_genType) Mock() Float64sAreSorted_genType {
	return func(a []float64) bool {
		m.Scene.T.Helper()
		moq := &MoqFloat64sAreSorted_genType_mock{Moq: m}
		return moq.Fn(a)
	}
}

func (m *MoqFloat64sAreSorted_genType_mock) Fn(a []float64) (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqFloat64sAreSorted_genType_params{
		A: a,
	}
	var results *MoqFloat64sAreSorted_genType_results
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
		result.DoFn(a)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(a)
	}
	return
}

func (m *MoqFloat64sAreSorted_genType) OnCall(a []float64) *MoqFloat64sAreSorted_genType_fnRecorder {
	return &MoqFloat64sAreSorted_genType_fnRecorder{
		Params: MoqFloat64sAreSorted_genType_params{
			A: a,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqFloat64sAreSorted_genType_fnRecorder) Any() *MoqFloat64sAreSorted_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqFloat64sAreSorted_genType_anyParams{Recorder: r}
}

func (a *MoqFloat64sAreSorted_genType_anyParams) A() *MoqFloat64sAreSorted_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqFloat64sAreSorted_genType_fnRecorder) Seq() *MoqFloat64sAreSorted_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFloat64sAreSorted_genType_fnRecorder) NoSeq() *MoqFloat64sAreSorted_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFloat64sAreSorted_genType_fnRecorder) ReturnResults(result1 bool) *MoqFloat64sAreSorted_genType_fnRecorder {
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
		DoFn       MoqFloat64sAreSorted_genType_doFn
		DoReturnFn MoqFloat64sAreSorted_genType_doReturnFn
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

func (r *MoqFloat64sAreSorted_genType_fnRecorder) AndDo(fn MoqFloat64sAreSorted_genType_doFn) *MoqFloat64sAreSorted_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFloat64sAreSorted_genType_fnRecorder) DoReturnResults(fn MoqFloat64sAreSorted_genType_doReturnFn) *MoqFloat64sAreSorted_genType_fnRecorder {
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
		DoFn       MoqFloat64sAreSorted_genType_doFn
		DoReturnFn MoqFloat64sAreSorted_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFloat64sAreSorted_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFloat64sAreSorted_genType_resultsByParams
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
		results = &MoqFloat64sAreSorted_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFloat64sAreSorted_genType_paramsKey]*MoqFloat64sAreSorted_genType_results{},
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
		r.Results = &MoqFloat64sAreSorted_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFloat64sAreSorted_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFloat64sAreSorted_genType_fnRecorder {
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
				DoFn       MoqFloat64sAreSorted_genType_doFn
				DoReturnFn MoqFloat64sAreSorted_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFloat64sAreSorted_genType) PrettyParams(params MoqFloat64sAreSorted_genType_params) string {
	return fmt.Sprintf("Float64sAreSorted_genType(%#v)", params.A)
}

func (m *MoqFloat64sAreSorted_genType) ParamsKey(params MoqFloat64sAreSorted_genType_params, anyParams uint64) MoqFloat64sAreSorted_genType_paramsKey {
	m.Scene.T.Helper()
	var aUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.A == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The a parameter can't be indexed by value")
		}
		aUsedHash = hash.DeepHash(params.A)
	}
	return MoqFloat64sAreSorted_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ A hash.Hash }{
			A: aUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqFloat64sAreSorted_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqFloat64sAreSorted_genType) AssertExpectationsMet() {
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

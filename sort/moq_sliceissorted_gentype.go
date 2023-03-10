// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package sort

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// SliceIsSorted_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type SliceIsSorted_genType func(x any, less func(i, j int) bool) bool

// MoqSliceIsSorted_genType holds the state of a moq of the
// SliceIsSorted_genType type
type MoqSliceIsSorted_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSliceIsSorted_genType_mock

	ResultsByParams []MoqSliceIsSorted_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			X    moq.ParamIndexing
			Less moq.ParamIndexing
		}
	}
}

// MoqSliceIsSorted_genType_mock isolates the mock interface of the
// SliceIsSorted_genType type
type MoqSliceIsSorted_genType_mock struct {
	Moq *MoqSliceIsSorted_genType
}

// MoqSliceIsSorted_genType_params holds the params of the
// SliceIsSorted_genType type
type MoqSliceIsSorted_genType_params struct {
	X    any
	Less func(i, j int) bool
}

// MoqSliceIsSorted_genType_paramsKey holds the map key params of the
// SliceIsSorted_genType type
type MoqSliceIsSorted_genType_paramsKey struct {
	Params struct{ X any }
	Hashes struct {
		X    hash.Hash
		Less hash.Hash
	}
}

// MoqSliceIsSorted_genType_resultsByParams contains the results for a given
// set of parameters for the SliceIsSorted_genType type
type MoqSliceIsSorted_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSliceIsSorted_genType_paramsKey]*MoqSliceIsSorted_genType_results
}

// MoqSliceIsSorted_genType_doFn defines the type of function needed when
// calling AndDo for the SliceIsSorted_genType type
type MoqSliceIsSorted_genType_doFn func(x any, less func(i, j int) bool)

// MoqSliceIsSorted_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the SliceIsSorted_genType type
type MoqSliceIsSorted_genType_doReturnFn func(x any, less func(i, j int) bool) bool

// MoqSliceIsSorted_genType_results holds the results of the
// SliceIsSorted_genType type
type MoqSliceIsSorted_genType_results struct {
	Params  MoqSliceIsSorted_genType_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqSliceIsSorted_genType_doFn
		DoReturnFn MoqSliceIsSorted_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSliceIsSorted_genType_fnRecorder routes recorded function calls to the
// MoqSliceIsSorted_genType moq
type MoqSliceIsSorted_genType_fnRecorder struct {
	Params    MoqSliceIsSorted_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSliceIsSorted_genType_results
	Moq       *MoqSliceIsSorted_genType
}

// MoqSliceIsSorted_genType_anyParams isolates the any params functions of the
// SliceIsSorted_genType type
type MoqSliceIsSorted_genType_anyParams struct {
	Recorder *MoqSliceIsSorted_genType_fnRecorder
}

// NewMoqSliceIsSorted_genType creates a new moq of the SliceIsSorted_genType
// type
func NewMoqSliceIsSorted_genType(scene *moq.Scene, config *moq.Config) *MoqSliceIsSorted_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSliceIsSorted_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSliceIsSorted_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				X    moq.ParamIndexing
				Less moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			X    moq.ParamIndexing
			Less moq.ParamIndexing
		}{
			X:    moq.ParamIndexByValue,
			Less: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the SliceIsSorted_genType type
func (m *MoqSliceIsSorted_genType) Mock() SliceIsSorted_genType {
	return func(x any, less func(i, j int) bool) bool {
		m.Scene.T.Helper()
		moq := &MoqSliceIsSorted_genType_mock{Moq: m}
		return moq.Fn(x, less)
	}
}

func (m *MoqSliceIsSorted_genType_mock) Fn(x any, less func(i, j int) bool) (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqSliceIsSorted_genType_params{
		X:    x,
		Less: less,
	}
	var results *MoqSliceIsSorted_genType_results
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
		result.DoFn(x, less)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(x, less)
	}
	return
}

func (m *MoqSliceIsSorted_genType) OnCall(x any, less func(i, j int) bool) *MoqSliceIsSorted_genType_fnRecorder {
	return &MoqSliceIsSorted_genType_fnRecorder{
		Params: MoqSliceIsSorted_genType_params{
			X:    x,
			Less: less,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSliceIsSorted_genType_fnRecorder) Any() *MoqSliceIsSorted_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSliceIsSorted_genType_anyParams{Recorder: r}
}

func (a *MoqSliceIsSorted_genType_anyParams) X() *MoqSliceIsSorted_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqSliceIsSorted_genType_anyParams) Less() *MoqSliceIsSorted_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqSliceIsSorted_genType_fnRecorder) Seq() *MoqSliceIsSorted_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSliceIsSorted_genType_fnRecorder) NoSeq() *MoqSliceIsSorted_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSliceIsSorted_genType_fnRecorder) ReturnResults(result1 bool) *MoqSliceIsSorted_genType_fnRecorder {
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
		DoFn       MoqSliceIsSorted_genType_doFn
		DoReturnFn MoqSliceIsSorted_genType_doReturnFn
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

func (r *MoqSliceIsSorted_genType_fnRecorder) AndDo(fn MoqSliceIsSorted_genType_doFn) *MoqSliceIsSorted_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSliceIsSorted_genType_fnRecorder) DoReturnResults(fn MoqSliceIsSorted_genType_doReturnFn) *MoqSliceIsSorted_genType_fnRecorder {
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
		DoFn       MoqSliceIsSorted_genType_doFn
		DoReturnFn MoqSliceIsSorted_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSliceIsSorted_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSliceIsSorted_genType_resultsByParams
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
		results = &MoqSliceIsSorted_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSliceIsSorted_genType_paramsKey]*MoqSliceIsSorted_genType_results{},
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
		r.Results = &MoqSliceIsSorted_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSliceIsSorted_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSliceIsSorted_genType_fnRecorder {
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
				DoFn       MoqSliceIsSorted_genType_doFn
				DoReturnFn MoqSliceIsSorted_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSliceIsSorted_genType) PrettyParams(params MoqSliceIsSorted_genType_params) string {
	return fmt.Sprintf("SliceIsSorted_genType(%#v, %#v)", params.X, moq.FnString(params.Less))
}

func (m *MoqSliceIsSorted_genType) ParamsKey(params MoqSliceIsSorted_genType_params, anyParams uint64) MoqSliceIsSorted_genType_paramsKey {
	m.Scene.T.Helper()
	var xUsed any
	var xUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.X == moq.ParamIndexByValue {
			xUsed = params.X
		} else {
			xUsedHash = hash.DeepHash(params.X)
		}
	}
	var lessUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Less == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The less parameter can't be indexed by value")
		}
		lessUsedHash = hash.DeepHash(params.Less)
	}
	return MoqSliceIsSorted_genType_paramsKey{
		Params: struct{ X any }{
			X: xUsed,
		},
		Hashes: struct {
			X    hash.Hash
			Less hash.Hash
		}{
			X:    xUsedHash,
			Less: lessUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSliceIsSorted_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSliceIsSorted_genType) AssertExpectationsMet() {
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

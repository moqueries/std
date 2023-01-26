// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package sort

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// SliceStable_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type SliceStable_genType func(slice interface{}, less func(i, j int) bool)

// MoqSliceStable_genType holds the state of a moq of the SliceStable_genType
// type
type MoqSliceStable_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSliceStable_genType_mock

	ResultsByParams []MoqSliceStable_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Slice moq.ParamIndexing
			Less  moq.ParamIndexing
		}
	}
}

// MoqSliceStable_genType_mock isolates the mock interface of the
// SliceStable_genType type
type MoqSliceStable_genType_mock struct {
	Moq *MoqSliceStable_genType
}

// MoqSliceStable_genType_params holds the params of the SliceStable_genType
// type
type MoqSliceStable_genType_params struct {
	Slice interface{}
	Less  func(i, j int) bool
}

// MoqSliceStable_genType_paramsKey holds the map key params of the
// SliceStable_genType type
type MoqSliceStable_genType_paramsKey struct {
	Params struct{ Slice interface{} }
	Hashes struct {
		Slice hash.Hash
		Less  hash.Hash
	}
}

// MoqSliceStable_genType_resultsByParams contains the results for a given set
// of parameters for the SliceStable_genType type
type MoqSliceStable_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSliceStable_genType_paramsKey]*MoqSliceStable_genType_results
}

// MoqSliceStable_genType_doFn defines the type of function needed when calling
// AndDo for the SliceStable_genType type
type MoqSliceStable_genType_doFn func(slice interface{}, less func(i, j int) bool)

// MoqSliceStable_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the SliceStable_genType type
type MoqSliceStable_genType_doReturnFn func(slice interface{}, less func(i, j int) bool)

// MoqSliceStable_genType_results holds the results of the SliceStable_genType
// type
type MoqSliceStable_genType_results struct {
	Params  MoqSliceStable_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSliceStable_genType_doFn
		DoReturnFn MoqSliceStable_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSliceStable_genType_fnRecorder routes recorded function calls to the
// MoqSliceStable_genType moq
type MoqSliceStable_genType_fnRecorder struct {
	Params    MoqSliceStable_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSliceStable_genType_results
	Moq       *MoqSliceStable_genType
}

// MoqSliceStable_genType_anyParams isolates the any params functions of the
// SliceStable_genType type
type MoqSliceStable_genType_anyParams struct {
	Recorder *MoqSliceStable_genType_fnRecorder
}

// NewMoqSliceStable_genType creates a new moq of the SliceStable_genType type
func NewMoqSliceStable_genType(scene *moq.Scene, config *moq.Config) *MoqSliceStable_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSliceStable_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSliceStable_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Slice moq.ParamIndexing
				Less  moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Slice moq.ParamIndexing
			Less  moq.ParamIndexing
		}{
			Slice: moq.ParamIndexByHash,
			Less:  moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the SliceStable_genType type
func (m *MoqSliceStable_genType) Mock() SliceStable_genType {
	return func(slice interface{}, less func(i, j int) bool) {
		m.Scene.T.Helper()
		moq := &MoqSliceStable_genType_mock{Moq: m}
		moq.Fn(slice, less)
	}
}

func (m *MoqSliceStable_genType_mock) Fn(slice interface{}, less func(i, j int) bool) {
	m.Moq.Scene.T.Helper()
	params := MoqSliceStable_genType_params{
		Slice: slice,
		Less:  less,
	}
	var results *MoqSliceStable_genType_results
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
		result.DoFn(slice, less)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(slice, less)
	}
	return
}

func (m *MoqSliceStable_genType) OnCall(slice interface{}, less func(i, j int) bool) *MoqSliceStable_genType_fnRecorder {
	return &MoqSliceStable_genType_fnRecorder{
		Params: MoqSliceStable_genType_params{
			Slice: slice,
			Less:  less,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSliceStable_genType_fnRecorder) Any() *MoqSliceStable_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSliceStable_genType_anyParams{Recorder: r}
}

func (a *MoqSliceStable_genType_anyParams) Slice() *MoqSliceStable_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqSliceStable_genType_anyParams) Less() *MoqSliceStable_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqSliceStable_genType_fnRecorder) Seq() *MoqSliceStable_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSliceStable_genType_fnRecorder) NoSeq() *MoqSliceStable_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSliceStable_genType_fnRecorder) ReturnResults() *MoqSliceStable_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSliceStable_genType_doFn
		DoReturnFn MoqSliceStable_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSliceStable_genType_fnRecorder) AndDo(fn MoqSliceStable_genType_doFn) *MoqSliceStable_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSliceStable_genType_fnRecorder) DoReturnResults(fn MoqSliceStable_genType_doReturnFn) *MoqSliceStable_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSliceStable_genType_doFn
		DoReturnFn MoqSliceStable_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSliceStable_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSliceStable_genType_resultsByParams
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
		results = &MoqSliceStable_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSliceStable_genType_paramsKey]*MoqSliceStable_genType_results{},
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
		r.Results = &MoqSliceStable_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSliceStable_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSliceStable_genType_fnRecorder {
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
				Values     *struct{}
				Sequence   uint32
				DoFn       MoqSliceStable_genType_doFn
				DoReturnFn MoqSliceStable_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSliceStable_genType) PrettyParams(params MoqSliceStable_genType_params) string {
	return fmt.Sprintf("SliceStable_genType(%#v, %#v)", params.Slice, params.Less)
}

func (m *MoqSliceStable_genType) ParamsKey(params MoqSliceStable_genType_params, anyParams uint64) MoqSliceStable_genType_paramsKey {
	m.Scene.T.Helper()
	var sliceUsed interface{}
	var sliceUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Slice == moq.ParamIndexByValue {
			sliceUsed = params.Slice
		} else {
			sliceUsedHash = hash.DeepHash(params.Slice)
		}
	}
	var lessUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Less == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The less parameter can't be indexed by value")
		}
		lessUsedHash = hash.DeepHash(params.Less)
	}
	return MoqSliceStable_genType_paramsKey{
		Params: struct{ Slice interface{} }{
			Slice: sliceUsed,
		},
		Hashes: struct {
			Slice hash.Hash
			Less  hash.Hash
		}{
			Slice: sliceUsedHash,
			Less:  lessUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSliceStable_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSliceStable_genType) AssertExpectationsMet() {
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

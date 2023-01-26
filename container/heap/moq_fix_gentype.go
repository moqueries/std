// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package heap

import (
	"container/heap"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Fix_genType is the fabricated implementation type of this mock (emitted when
// mocking functions directly and not from a function type)
type Fix_genType func(h heap.Interface, i int)

// MoqFix_genType holds the state of a moq of the Fix_genType type
type MoqFix_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqFix_genType_mock

	ResultsByParams []MoqFix_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			H      moq.ParamIndexing
			Param2 moq.ParamIndexing
		}
	}
}

// MoqFix_genType_mock isolates the mock interface of the Fix_genType type
type MoqFix_genType_mock struct {
	Moq *MoqFix_genType
}

// MoqFix_genType_params holds the params of the Fix_genType type
type MoqFix_genType_params struct {
	H      heap.Interface
	Param2 int
}

// MoqFix_genType_paramsKey holds the map key params of the Fix_genType type
type MoqFix_genType_paramsKey struct {
	Params struct {
		H      heap.Interface
		Param2 int
	}
	Hashes struct {
		H      hash.Hash
		Param2 hash.Hash
	}
}

// MoqFix_genType_resultsByParams contains the results for a given set of
// parameters for the Fix_genType type
type MoqFix_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFix_genType_paramsKey]*MoqFix_genType_results
}

// MoqFix_genType_doFn defines the type of function needed when calling AndDo
// for the Fix_genType type
type MoqFix_genType_doFn func(h heap.Interface, i int)

// MoqFix_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Fix_genType type
type MoqFix_genType_doReturnFn func(h heap.Interface, i int)

// MoqFix_genType_results holds the results of the Fix_genType type
type MoqFix_genType_results struct {
	Params  MoqFix_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqFix_genType_doFn
		DoReturnFn MoqFix_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFix_genType_fnRecorder routes recorded function calls to the
// MoqFix_genType moq
type MoqFix_genType_fnRecorder struct {
	Params    MoqFix_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFix_genType_results
	Moq       *MoqFix_genType
}

// MoqFix_genType_anyParams isolates the any params functions of the
// Fix_genType type
type MoqFix_genType_anyParams struct {
	Recorder *MoqFix_genType_fnRecorder
}

// NewMoqFix_genType creates a new moq of the Fix_genType type
func NewMoqFix_genType(scene *moq.Scene, config *moq.Config) *MoqFix_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqFix_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqFix_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				H      moq.ParamIndexing
				Param2 moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			H      moq.ParamIndexing
			Param2 moq.ParamIndexing
		}{
			H:      moq.ParamIndexByHash,
			Param2: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Fix_genType type
func (m *MoqFix_genType) Mock() Fix_genType {
	return func(h heap.Interface, param2 int) {
		m.Scene.T.Helper()
		moq := &MoqFix_genType_mock{Moq: m}
		moq.Fn(h, param2)
	}
}

func (m *MoqFix_genType_mock) Fn(h heap.Interface, param2 int) {
	m.Moq.Scene.T.Helper()
	params := MoqFix_genType_params{
		H:      h,
		Param2: param2,
	}
	var results *MoqFix_genType_results
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
		result.DoFn(h, param2)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(h, param2)
	}
	return
}

func (m *MoqFix_genType) OnCall(h heap.Interface, param2 int) *MoqFix_genType_fnRecorder {
	return &MoqFix_genType_fnRecorder{
		Params: MoqFix_genType_params{
			H:      h,
			Param2: param2,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqFix_genType_fnRecorder) Any() *MoqFix_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqFix_genType_anyParams{Recorder: r}
}

func (a *MoqFix_genType_anyParams) H() *MoqFix_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqFix_genType_anyParams) Param2() *MoqFix_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqFix_genType_fnRecorder) Seq() *MoqFix_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFix_genType_fnRecorder) NoSeq() *MoqFix_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFix_genType_fnRecorder) ReturnResults() *MoqFix_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqFix_genType_doFn
		DoReturnFn MoqFix_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqFix_genType_fnRecorder) AndDo(fn MoqFix_genType_doFn) *MoqFix_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFix_genType_fnRecorder) DoReturnResults(fn MoqFix_genType_doReturnFn) *MoqFix_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqFix_genType_doFn
		DoReturnFn MoqFix_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFix_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFix_genType_resultsByParams
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
		results = &MoqFix_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFix_genType_paramsKey]*MoqFix_genType_results{},
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
		r.Results = &MoqFix_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFix_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFix_genType_fnRecorder {
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
				DoFn       MoqFix_genType_doFn
				DoReturnFn MoqFix_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFix_genType) PrettyParams(params MoqFix_genType_params) string {
	return fmt.Sprintf("Fix_genType(%#v, %#v)", params.H, params.Param2)
}

func (m *MoqFix_genType) ParamsKey(params MoqFix_genType_params, anyParams uint64) MoqFix_genType_paramsKey {
	m.Scene.T.Helper()
	var hUsed heap.Interface
	var hUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.H == moq.ParamIndexByValue {
			hUsed = params.H
		} else {
			hUsedHash = hash.DeepHash(params.H)
		}
	}
	var param2Used int
	var param2UsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Param2 == moq.ParamIndexByValue {
			param2Used = params.Param2
		} else {
			param2UsedHash = hash.DeepHash(params.Param2)
		}
	}
	return MoqFix_genType_paramsKey{
		Params: struct {
			H      heap.Interface
			Param2 int
		}{
			H:      hUsed,
			Param2: param2Used,
		},
		Hashes: struct {
			H      hash.Hash
			Param2 hash.Hash
		}{
			H:      hUsedHash,
			Param2: param2UsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqFix_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqFix_genType) AssertExpectationsMet() {
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

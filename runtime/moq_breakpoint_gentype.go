// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package runtime

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// Breakpoint_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type Breakpoint_genType func()

// MoqBreakpoint_genType holds the state of a moq of the Breakpoint_genType
// type
type MoqBreakpoint_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqBreakpoint_genType_mock

	ResultsByParams []MoqBreakpoint_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct{}
	}
}

// MoqBreakpoint_genType_mock isolates the mock interface of the
// Breakpoint_genType type
type MoqBreakpoint_genType_mock struct {
	Moq *MoqBreakpoint_genType
}

// MoqBreakpoint_genType_params holds the params of the Breakpoint_genType type
type MoqBreakpoint_genType_params struct{}

// MoqBreakpoint_genType_paramsKey holds the map key params of the
// Breakpoint_genType type
type MoqBreakpoint_genType_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqBreakpoint_genType_resultsByParams contains the results for a given set
// of parameters for the Breakpoint_genType type
type MoqBreakpoint_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqBreakpoint_genType_paramsKey]*MoqBreakpoint_genType_results
}

// MoqBreakpoint_genType_doFn defines the type of function needed when calling
// AndDo for the Breakpoint_genType type
type MoqBreakpoint_genType_doFn func()

// MoqBreakpoint_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Breakpoint_genType type
type MoqBreakpoint_genType_doReturnFn func()

// MoqBreakpoint_genType_results holds the results of the Breakpoint_genType
// type
type MoqBreakpoint_genType_results struct {
	Params  MoqBreakpoint_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqBreakpoint_genType_doFn
		DoReturnFn MoqBreakpoint_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqBreakpoint_genType_fnRecorder routes recorded function calls to the
// MoqBreakpoint_genType moq
type MoqBreakpoint_genType_fnRecorder struct {
	Params    MoqBreakpoint_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqBreakpoint_genType_results
	Moq       *MoqBreakpoint_genType
}

// MoqBreakpoint_genType_anyParams isolates the any params functions of the
// Breakpoint_genType type
type MoqBreakpoint_genType_anyParams struct {
	Recorder *MoqBreakpoint_genType_fnRecorder
}

// NewMoqBreakpoint_genType creates a new moq of the Breakpoint_genType type
func NewMoqBreakpoint_genType(scene *moq.Scene, config *moq.Config) *MoqBreakpoint_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqBreakpoint_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqBreakpoint_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct{}
		}{ParameterIndexing: struct{}{}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Breakpoint_genType type
func (m *MoqBreakpoint_genType) Mock() Breakpoint_genType {
	return func() { m.Scene.T.Helper(); moq := &MoqBreakpoint_genType_mock{Moq: m}; moq.Fn() }
}

func (m *MoqBreakpoint_genType_mock) Fn() {
	m.Moq.Scene.T.Helper()
	params := MoqBreakpoint_genType_params{}
	var results *MoqBreakpoint_genType_results
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
		result.DoFn()
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn()
	}
	return
}

func (m *MoqBreakpoint_genType) OnCall() *MoqBreakpoint_genType_fnRecorder {
	return &MoqBreakpoint_genType_fnRecorder{
		Params:   MoqBreakpoint_genType_params{},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqBreakpoint_genType_fnRecorder) Any() *MoqBreakpoint_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqBreakpoint_genType_anyParams{Recorder: r}
}

func (r *MoqBreakpoint_genType_fnRecorder) Seq() *MoqBreakpoint_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqBreakpoint_genType_fnRecorder) NoSeq() *MoqBreakpoint_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqBreakpoint_genType_fnRecorder) ReturnResults() *MoqBreakpoint_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqBreakpoint_genType_doFn
		DoReturnFn MoqBreakpoint_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqBreakpoint_genType_fnRecorder) AndDo(fn MoqBreakpoint_genType_doFn) *MoqBreakpoint_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqBreakpoint_genType_fnRecorder) DoReturnResults(fn MoqBreakpoint_genType_doReturnFn) *MoqBreakpoint_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqBreakpoint_genType_doFn
		DoReturnFn MoqBreakpoint_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqBreakpoint_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqBreakpoint_genType_resultsByParams
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
		results = &MoqBreakpoint_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqBreakpoint_genType_paramsKey]*MoqBreakpoint_genType_results{},
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
		r.Results = &MoqBreakpoint_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqBreakpoint_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqBreakpoint_genType_fnRecorder {
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
				DoFn       MoqBreakpoint_genType_doFn
				DoReturnFn MoqBreakpoint_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqBreakpoint_genType) PrettyParams(params MoqBreakpoint_genType_params) string {
	return fmt.Sprintf("Breakpoint_genType()")
}

func (m *MoqBreakpoint_genType) ParamsKey(params MoqBreakpoint_genType_params, anyParams uint64) MoqBreakpoint_genType_paramsKey {
	m.Scene.T.Helper()
	return MoqBreakpoint_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqBreakpoint_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqBreakpoint_genType) AssertExpectationsMet() {
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
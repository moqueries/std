// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package pprof

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// StopCPUProfile_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type StopCPUProfile_genType func()

// MoqStopCPUProfile_genType holds the state of a moq of the
// StopCPUProfile_genType type
type MoqStopCPUProfile_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqStopCPUProfile_genType_mock

	ResultsByParams []MoqStopCPUProfile_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct{}
	}
}

// MoqStopCPUProfile_genType_mock isolates the mock interface of the
// StopCPUProfile_genType type
type MoqStopCPUProfile_genType_mock struct {
	Moq *MoqStopCPUProfile_genType
}

// MoqStopCPUProfile_genType_params holds the params of the
// StopCPUProfile_genType type
type MoqStopCPUProfile_genType_params struct{}

// MoqStopCPUProfile_genType_paramsKey holds the map key params of the
// StopCPUProfile_genType type
type MoqStopCPUProfile_genType_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqStopCPUProfile_genType_resultsByParams contains the results for a given
// set of parameters for the StopCPUProfile_genType type
type MoqStopCPUProfile_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqStopCPUProfile_genType_paramsKey]*MoqStopCPUProfile_genType_results
}

// MoqStopCPUProfile_genType_doFn defines the type of function needed when
// calling AndDo for the StopCPUProfile_genType type
type MoqStopCPUProfile_genType_doFn func()

// MoqStopCPUProfile_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the StopCPUProfile_genType type
type MoqStopCPUProfile_genType_doReturnFn func()

// MoqStopCPUProfile_genType_results holds the results of the
// StopCPUProfile_genType type
type MoqStopCPUProfile_genType_results struct {
	Params  MoqStopCPUProfile_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqStopCPUProfile_genType_doFn
		DoReturnFn MoqStopCPUProfile_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqStopCPUProfile_genType_fnRecorder routes recorded function calls to the
// MoqStopCPUProfile_genType moq
type MoqStopCPUProfile_genType_fnRecorder struct {
	Params    MoqStopCPUProfile_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqStopCPUProfile_genType_results
	Moq       *MoqStopCPUProfile_genType
}

// MoqStopCPUProfile_genType_anyParams isolates the any params functions of the
// StopCPUProfile_genType type
type MoqStopCPUProfile_genType_anyParams struct {
	Recorder *MoqStopCPUProfile_genType_fnRecorder
}

// NewMoqStopCPUProfile_genType creates a new moq of the StopCPUProfile_genType
// type
func NewMoqStopCPUProfile_genType(scene *moq.Scene, config *moq.Config) *MoqStopCPUProfile_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqStopCPUProfile_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqStopCPUProfile_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct{}
		}{ParameterIndexing: struct{}{}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the StopCPUProfile_genType type
func (m *MoqStopCPUProfile_genType) Mock() StopCPUProfile_genType {
	return func() { m.Scene.T.Helper(); moq := &MoqStopCPUProfile_genType_mock{Moq: m}; moq.Fn() }
}

func (m *MoqStopCPUProfile_genType_mock) Fn() {
	m.Moq.Scene.T.Helper()
	params := MoqStopCPUProfile_genType_params{}
	var results *MoqStopCPUProfile_genType_results
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

func (m *MoqStopCPUProfile_genType) OnCall() *MoqStopCPUProfile_genType_fnRecorder {
	return &MoqStopCPUProfile_genType_fnRecorder{
		Params:   MoqStopCPUProfile_genType_params{},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqStopCPUProfile_genType_fnRecorder) Any() *MoqStopCPUProfile_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqStopCPUProfile_genType_anyParams{Recorder: r}
}

func (r *MoqStopCPUProfile_genType_fnRecorder) Seq() *MoqStopCPUProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqStopCPUProfile_genType_fnRecorder) NoSeq() *MoqStopCPUProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqStopCPUProfile_genType_fnRecorder) ReturnResults() *MoqStopCPUProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqStopCPUProfile_genType_doFn
		DoReturnFn MoqStopCPUProfile_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqStopCPUProfile_genType_fnRecorder) AndDo(fn MoqStopCPUProfile_genType_doFn) *MoqStopCPUProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqStopCPUProfile_genType_fnRecorder) DoReturnResults(fn MoqStopCPUProfile_genType_doReturnFn) *MoqStopCPUProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqStopCPUProfile_genType_doFn
		DoReturnFn MoqStopCPUProfile_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqStopCPUProfile_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqStopCPUProfile_genType_resultsByParams
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
		results = &MoqStopCPUProfile_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqStopCPUProfile_genType_paramsKey]*MoqStopCPUProfile_genType_results{},
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
		r.Results = &MoqStopCPUProfile_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqStopCPUProfile_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqStopCPUProfile_genType_fnRecorder {
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
				DoFn       MoqStopCPUProfile_genType_doFn
				DoReturnFn MoqStopCPUProfile_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqStopCPUProfile_genType) PrettyParams(params MoqStopCPUProfile_genType_params) string {
	return fmt.Sprintf("StopCPUProfile_genType()")
}

func (m *MoqStopCPUProfile_genType) ParamsKey(params MoqStopCPUProfile_genType_params, anyParams uint64) MoqStopCPUProfile_genType_paramsKey {
	m.Scene.T.Helper()
	return MoqStopCPUProfile_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqStopCPUProfile_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqStopCPUProfile_genType) AssertExpectationsMet() {
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

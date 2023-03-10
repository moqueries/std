// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package debug

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// PrintStack_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type PrintStack_genType func()

// MoqPrintStack_genType holds the state of a moq of the PrintStack_genType
// type
type MoqPrintStack_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqPrintStack_genType_mock

	ResultsByParams []MoqPrintStack_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct{}
	}
}

// MoqPrintStack_genType_mock isolates the mock interface of the
// PrintStack_genType type
type MoqPrintStack_genType_mock struct {
	Moq *MoqPrintStack_genType
}

// MoqPrintStack_genType_params holds the params of the PrintStack_genType type
type MoqPrintStack_genType_params struct{}

// MoqPrintStack_genType_paramsKey holds the map key params of the
// PrintStack_genType type
type MoqPrintStack_genType_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqPrintStack_genType_resultsByParams contains the results for a given set
// of parameters for the PrintStack_genType type
type MoqPrintStack_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqPrintStack_genType_paramsKey]*MoqPrintStack_genType_results
}

// MoqPrintStack_genType_doFn defines the type of function needed when calling
// AndDo for the PrintStack_genType type
type MoqPrintStack_genType_doFn func()

// MoqPrintStack_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the PrintStack_genType type
type MoqPrintStack_genType_doReturnFn func()

// MoqPrintStack_genType_results holds the results of the PrintStack_genType
// type
type MoqPrintStack_genType_results struct {
	Params  MoqPrintStack_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqPrintStack_genType_doFn
		DoReturnFn MoqPrintStack_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqPrintStack_genType_fnRecorder routes recorded function calls to the
// MoqPrintStack_genType moq
type MoqPrintStack_genType_fnRecorder struct {
	Params    MoqPrintStack_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqPrintStack_genType_results
	Moq       *MoqPrintStack_genType
}

// MoqPrintStack_genType_anyParams isolates the any params functions of the
// PrintStack_genType type
type MoqPrintStack_genType_anyParams struct {
	Recorder *MoqPrintStack_genType_fnRecorder
}

// NewMoqPrintStack_genType creates a new moq of the PrintStack_genType type
func NewMoqPrintStack_genType(scene *moq.Scene, config *moq.Config) *MoqPrintStack_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqPrintStack_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqPrintStack_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct{}
		}{ParameterIndexing: struct{}{}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the PrintStack_genType type
func (m *MoqPrintStack_genType) Mock() PrintStack_genType {
	return func() { m.Scene.T.Helper(); moq := &MoqPrintStack_genType_mock{Moq: m}; moq.Fn() }
}

func (m *MoqPrintStack_genType_mock) Fn() {
	m.Moq.Scene.T.Helper()
	params := MoqPrintStack_genType_params{}
	var results *MoqPrintStack_genType_results
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

func (m *MoqPrintStack_genType) OnCall() *MoqPrintStack_genType_fnRecorder {
	return &MoqPrintStack_genType_fnRecorder{
		Params:   MoqPrintStack_genType_params{},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqPrintStack_genType_fnRecorder) Any() *MoqPrintStack_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqPrintStack_genType_anyParams{Recorder: r}
}

func (r *MoqPrintStack_genType_fnRecorder) Seq() *MoqPrintStack_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqPrintStack_genType_fnRecorder) NoSeq() *MoqPrintStack_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqPrintStack_genType_fnRecorder) ReturnResults() *MoqPrintStack_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqPrintStack_genType_doFn
		DoReturnFn MoqPrintStack_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqPrintStack_genType_fnRecorder) AndDo(fn MoqPrintStack_genType_doFn) *MoqPrintStack_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqPrintStack_genType_fnRecorder) DoReturnResults(fn MoqPrintStack_genType_doReturnFn) *MoqPrintStack_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqPrintStack_genType_doFn
		DoReturnFn MoqPrintStack_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqPrintStack_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqPrintStack_genType_resultsByParams
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
		results = &MoqPrintStack_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqPrintStack_genType_paramsKey]*MoqPrintStack_genType_results{},
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
		r.Results = &MoqPrintStack_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqPrintStack_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqPrintStack_genType_fnRecorder {
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
				DoFn       MoqPrintStack_genType_doFn
				DoReturnFn MoqPrintStack_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqPrintStack_genType) PrettyParams(params MoqPrintStack_genType_params) string {
	return fmt.Sprintf("PrintStack_genType()")
}

func (m *MoqPrintStack_genType) ParamsKey(params MoqPrintStack_genType_params, anyParams uint64) MoqPrintStack_genType_paramsKey {
	m.Scene.T.Helper()
	return MoqPrintStack_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqPrintStack_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqPrintStack_genType) AssertExpectationsMet() {
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

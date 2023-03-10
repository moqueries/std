// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package testing

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that testing.M_starGenType is mocked
// completely
var _ M_starGenType = (*MoqM_starGenType_mock)(nil)

// M_starGenType is the fabricated implementation type of this mock (emitted
// when mocking a collections of methods directly and not from an interface
// type)
type M_starGenType interface {
	Run() (code int)
}

// MoqM_starGenType holds the state of a moq of the M_starGenType type
type MoqM_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqM_starGenType_mock

	ResultsByParams_Run []MoqM_starGenType_Run_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Run struct{}
		}
	}
}

// MoqM_starGenType_mock isolates the mock interface of the M_starGenType type
type MoqM_starGenType_mock struct {
	Moq *MoqM_starGenType
}

// MoqM_starGenType_recorder isolates the recorder interface of the
// M_starGenType type
type MoqM_starGenType_recorder struct {
	Moq *MoqM_starGenType
}

// MoqM_starGenType_Run_params holds the params of the M_starGenType type
type MoqM_starGenType_Run_params struct{}

// MoqM_starGenType_Run_paramsKey holds the map key params of the M_starGenType
// type
type MoqM_starGenType_Run_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqM_starGenType_Run_resultsByParams contains the results for a given set of
// parameters for the M_starGenType type
type MoqM_starGenType_Run_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqM_starGenType_Run_paramsKey]*MoqM_starGenType_Run_results
}

// MoqM_starGenType_Run_doFn defines the type of function needed when calling
// AndDo for the M_starGenType type
type MoqM_starGenType_Run_doFn func()

// MoqM_starGenType_Run_doReturnFn defines the type of function needed when
// calling DoReturnResults for the M_starGenType type
type MoqM_starGenType_Run_doReturnFn func() (code int)

// MoqM_starGenType_Run_results holds the results of the M_starGenType type
type MoqM_starGenType_Run_results struct {
	Params  MoqM_starGenType_Run_params
	Results []struct {
		Values     *struct{ Code int }
		Sequence   uint32
		DoFn       MoqM_starGenType_Run_doFn
		DoReturnFn MoqM_starGenType_Run_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqM_starGenType_Run_fnRecorder routes recorded function calls to the
// MoqM_starGenType moq
type MoqM_starGenType_Run_fnRecorder struct {
	Params    MoqM_starGenType_Run_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqM_starGenType_Run_results
	Moq       *MoqM_starGenType
}

// MoqM_starGenType_Run_anyParams isolates the any params functions of the
// M_starGenType type
type MoqM_starGenType_Run_anyParams struct {
	Recorder *MoqM_starGenType_Run_fnRecorder
}

// NewMoqM_starGenType creates a new moq of the M_starGenType type
func NewMoqM_starGenType(scene *moq.Scene, config *moq.Config) *MoqM_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqM_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqM_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Run struct{}
			}
		}{ParameterIndexing: struct {
			Run struct{}
		}{
			Run: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the M_starGenType type
func (m *MoqM_starGenType) Mock() *MoqM_starGenType_mock { return m.Moq }

func (m *MoqM_starGenType_mock) Run() (code int) {
	m.Moq.Scene.T.Helper()
	params := MoqM_starGenType_Run_params{}
	var results *MoqM_starGenType_Run_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Run {
		paramsKey := m.Moq.ParamsKey_Run(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Run(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Run(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Run(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		code = result.Values.Code
	}
	if result.DoReturnFn != nil {
		code = result.DoReturnFn()
	}
	return
}

// OnCall returns the recorder implementation of the M_starGenType type
func (m *MoqM_starGenType) OnCall() *MoqM_starGenType_recorder {
	return &MoqM_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqM_starGenType_recorder) Run() *MoqM_starGenType_Run_fnRecorder {
	return &MoqM_starGenType_Run_fnRecorder{
		Params:   MoqM_starGenType_Run_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqM_starGenType_Run_fnRecorder) Any() *MoqM_starGenType_Run_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Run(r.Params))
		return nil
	}
	return &MoqM_starGenType_Run_anyParams{Recorder: r}
}

func (r *MoqM_starGenType_Run_fnRecorder) Seq() *MoqM_starGenType_Run_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Run(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqM_starGenType_Run_fnRecorder) NoSeq() *MoqM_starGenType_Run_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Run(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqM_starGenType_Run_fnRecorder) ReturnResults(code int) *MoqM_starGenType_Run_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Code int }
		Sequence   uint32
		DoFn       MoqM_starGenType_Run_doFn
		DoReturnFn MoqM_starGenType_Run_doReturnFn
	}{
		Values: &struct{ Code int }{
			Code: code,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqM_starGenType_Run_fnRecorder) AndDo(fn MoqM_starGenType_Run_doFn) *MoqM_starGenType_Run_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqM_starGenType_Run_fnRecorder) DoReturnResults(fn MoqM_starGenType_Run_doReturnFn) *MoqM_starGenType_Run_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Code int }
		Sequence   uint32
		DoFn       MoqM_starGenType_Run_doFn
		DoReturnFn MoqM_starGenType_Run_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqM_starGenType_Run_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqM_starGenType_Run_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Run {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqM_starGenType_Run_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqM_starGenType_Run_paramsKey]*MoqM_starGenType_Run_results{},
		}
		r.Moq.ResultsByParams_Run = append(r.Moq.ResultsByParams_Run, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Run) {
			copy(r.Moq.ResultsByParams_Run[insertAt+1:], r.Moq.ResultsByParams_Run[insertAt:0])
			r.Moq.ResultsByParams_Run[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Run(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqM_starGenType_Run_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqM_starGenType_Run_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqM_starGenType_Run_fnRecorder {
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
				Values     *struct{ Code int }
				Sequence   uint32
				DoFn       MoqM_starGenType_Run_doFn
				DoReturnFn MoqM_starGenType_Run_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqM_starGenType) PrettyParams_Run(params MoqM_starGenType_Run_params) string {
	return fmt.Sprintf("Run()")
}

func (m *MoqM_starGenType) ParamsKey_Run(params MoqM_starGenType_Run_params, anyParams uint64) MoqM_starGenType_Run_paramsKey {
	m.Scene.T.Helper()
	return MoqM_starGenType_Run_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqM_starGenType) Reset() { m.ResultsByParams_Run = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqM_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Run {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Run(results.Params))
			}
		}
	}
}

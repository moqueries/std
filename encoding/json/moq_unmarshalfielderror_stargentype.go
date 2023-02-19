// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package json

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that
// json.UnmarshalFieldError_starGenType is mocked completely
var _ UnmarshalFieldError_starGenType = (*MoqUnmarshalFieldError_starGenType_mock)(nil)

// UnmarshalFieldError_starGenType is the fabricated implementation type of
// this mock (emitted when mocking a collections of methods directly and not
// from an interface type)
type UnmarshalFieldError_starGenType interface {
	Error() string
}

// MoqUnmarshalFieldError_starGenType holds the state of a moq of the
// UnmarshalFieldError_starGenType type
type MoqUnmarshalFieldError_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqUnmarshalFieldError_starGenType_mock

	ResultsByParams_Error []MoqUnmarshalFieldError_starGenType_Error_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Error struct{}
		}
	}
}

// MoqUnmarshalFieldError_starGenType_mock isolates the mock interface of the
// UnmarshalFieldError_starGenType type
type MoqUnmarshalFieldError_starGenType_mock struct {
	Moq *MoqUnmarshalFieldError_starGenType
}

// MoqUnmarshalFieldError_starGenType_recorder isolates the recorder interface
// of the UnmarshalFieldError_starGenType type
type MoqUnmarshalFieldError_starGenType_recorder struct {
	Moq *MoqUnmarshalFieldError_starGenType
}

// MoqUnmarshalFieldError_starGenType_Error_params holds the params of the
// UnmarshalFieldError_starGenType type
type MoqUnmarshalFieldError_starGenType_Error_params struct{}

// MoqUnmarshalFieldError_starGenType_Error_paramsKey holds the map key params
// of the UnmarshalFieldError_starGenType type
type MoqUnmarshalFieldError_starGenType_Error_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqUnmarshalFieldError_starGenType_Error_resultsByParams contains the
// results for a given set of parameters for the
// UnmarshalFieldError_starGenType type
type MoqUnmarshalFieldError_starGenType_Error_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqUnmarshalFieldError_starGenType_Error_paramsKey]*MoqUnmarshalFieldError_starGenType_Error_results
}

// MoqUnmarshalFieldError_starGenType_Error_doFn defines the type of function
// needed when calling AndDo for the UnmarshalFieldError_starGenType type
type MoqUnmarshalFieldError_starGenType_Error_doFn func()

// MoqUnmarshalFieldError_starGenType_Error_doReturnFn defines the type of
// function needed when calling DoReturnResults for the
// UnmarshalFieldError_starGenType type
type MoqUnmarshalFieldError_starGenType_Error_doReturnFn func() string

// MoqUnmarshalFieldError_starGenType_Error_results holds the results of the
// UnmarshalFieldError_starGenType type
type MoqUnmarshalFieldError_starGenType_Error_results struct {
	Params  MoqUnmarshalFieldError_starGenType_Error_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqUnmarshalFieldError_starGenType_Error_doFn
		DoReturnFn MoqUnmarshalFieldError_starGenType_Error_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqUnmarshalFieldError_starGenType_Error_fnRecorder routes recorded function
// calls to the MoqUnmarshalFieldError_starGenType moq
type MoqUnmarshalFieldError_starGenType_Error_fnRecorder struct {
	Params    MoqUnmarshalFieldError_starGenType_Error_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqUnmarshalFieldError_starGenType_Error_results
	Moq       *MoqUnmarshalFieldError_starGenType
}

// MoqUnmarshalFieldError_starGenType_Error_anyParams isolates the any params
// functions of the UnmarshalFieldError_starGenType type
type MoqUnmarshalFieldError_starGenType_Error_anyParams struct {
	Recorder *MoqUnmarshalFieldError_starGenType_Error_fnRecorder
}

// NewMoqUnmarshalFieldError_starGenType creates a new moq of the
// UnmarshalFieldError_starGenType type
func NewMoqUnmarshalFieldError_starGenType(scene *moq.Scene, config *moq.Config) *MoqUnmarshalFieldError_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqUnmarshalFieldError_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqUnmarshalFieldError_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Error struct{}
			}
		}{ParameterIndexing: struct {
			Error struct{}
		}{
			Error: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the UnmarshalFieldError_starGenType
// type
func (m *MoqUnmarshalFieldError_starGenType) Mock() *MoqUnmarshalFieldError_starGenType_mock {
	return m.Moq
}

func (m *MoqUnmarshalFieldError_starGenType_mock) Error() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqUnmarshalFieldError_starGenType_Error_params{}
	var results *MoqUnmarshalFieldError_starGenType_Error_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Error {
		paramsKey := m.Moq.ParamsKey_Error(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Error(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Error(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Error(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn()
	}
	return
}

// OnCall returns the recorder implementation of the
// UnmarshalFieldError_starGenType type
func (m *MoqUnmarshalFieldError_starGenType) OnCall() *MoqUnmarshalFieldError_starGenType_recorder {
	return &MoqUnmarshalFieldError_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqUnmarshalFieldError_starGenType_recorder) Error() *MoqUnmarshalFieldError_starGenType_Error_fnRecorder {
	return &MoqUnmarshalFieldError_starGenType_Error_fnRecorder{
		Params:   MoqUnmarshalFieldError_starGenType_Error_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqUnmarshalFieldError_starGenType_Error_fnRecorder) Any() *MoqUnmarshalFieldError_starGenType_Error_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	return &MoqUnmarshalFieldError_starGenType_Error_anyParams{Recorder: r}
}

func (r *MoqUnmarshalFieldError_starGenType_Error_fnRecorder) Seq() *MoqUnmarshalFieldError_starGenType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqUnmarshalFieldError_starGenType_Error_fnRecorder) NoSeq() *MoqUnmarshalFieldError_starGenType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqUnmarshalFieldError_starGenType_Error_fnRecorder) ReturnResults(result1 string) *MoqUnmarshalFieldError_starGenType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqUnmarshalFieldError_starGenType_Error_doFn
		DoReturnFn MoqUnmarshalFieldError_starGenType_Error_doReturnFn
	}{
		Values: &struct {
			Result1 string
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqUnmarshalFieldError_starGenType_Error_fnRecorder) AndDo(fn MoqUnmarshalFieldError_starGenType_Error_doFn) *MoqUnmarshalFieldError_starGenType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqUnmarshalFieldError_starGenType_Error_fnRecorder) DoReturnResults(fn MoqUnmarshalFieldError_starGenType_Error_doReturnFn) *MoqUnmarshalFieldError_starGenType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqUnmarshalFieldError_starGenType_Error_doFn
		DoReturnFn MoqUnmarshalFieldError_starGenType_Error_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqUnmarshalFieldError_starGenType_Error_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqUnmarshalFieldError_starGenType_Error_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Error {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqUnmarshalFieldError_starGenType_Error_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqUnmarshalFieldError_starGenType_Error_paramsKey]*MoqUnmarshalFieldError_starGenType_Error_results{},
		}
		r.Moq.ResultsByParams_Error = append(r.Moq.ResultsByParams_Error, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Error) {
			copy(r.Moq.ResultsByParams_Error[insertAt+1:], r.Moq.ResultsByParams_Error[insertAt:0])
			r.Moq.ResultsByParams_Error[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Error(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqUnmarshalFieldError_starGenType_Error_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqUnmarshalFieldError_starGenType_Error_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqUnmarshalFieldError_starGenType_Error_fnRecorder {
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
					Result1 string
				}
				Sequence   uint32
				DoFn       MoqUnmarshalFieldError_starGenType_Error_doFn
				DoReturnFn MoqUnmarshalFieldError_starGenType_Error_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqUnmarshalFieldError_starGenType) PrettyParams_Error(params MoqUnmarshalFieldError_starGenType_Error_params) string {
	return fmt.Sprintf("Error()")
}

func (m *MoqUnmarshalFieldError_starGenType) ParamsKey_Error(params MoqUnmarshalFieldError_starGenType_Error_params, anyParams uint64) MoqUnmarshalFieldError_starGenType_Error_paramsKey {
	m.Scene.T.Helper()
	return MoqUnmarshalFieldError_starGenType_Error_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqUnmarshalFieldError_starGenType) Reset() { m.ResultsByParams_Error = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqUnmarshalFieldError_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Error {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Error(results.Params))
			}
		}
	}
}
// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package json

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that
// json.UnsupportedTypeError_starGenType is mocked completely
var _ UnsupportedTypeError_starGenType = (*MoqUnsupportedTypeError_starGenType_mock)(nil)

// UnsupportedTypeError_starGenType is the fabricated implementation type of
// this mock (emitted when mocking a collections of methods directly and not
// from an interface type)
type UnsupportedTypeError_starGenType interface {
	Error() string
}

// MoqUnsupportedTypeError_starGenType holds the state of a moq of the
// UnsupportedTypeError_starGenType type
type MoqUnsupportedTypeError_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqUnsupportedTypeError_starGenType_mock

	ResultsByParams_Error []MoqUnsupportedTypeError_starGenType_Error_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Error struct{}
		}
	}
}

// MoqUnsupportedTypeError_starGenType_mock isolates the mock interface of the
// UnsupportedTypeError_starGenType type
type MoqUnsupportedTypeError_starGenType_mock struct {
	Moq *MoqUnsupportedTypeError_starGenType
}

// MoqUnsupportedTypeError_starGenType_recorder isolates the recorder interface
// of the UnsupportedTypeError_starGenType type
type MoqUnsupportedTypeError_starGenType_recorder struct {
	Moq *MoqUnsupportedTypeError_starGenType
}

// MoqUnsupportedTypeError_starGenType_Error_params holds the params of the
// UnsupportedTypeError_starGenType type
type MoqUnsupportedTypeError_starGenType_Error_params struct{}

// MoqUnsupportedTypeError_starGenType_Error_paramsKey holds the map key params
// of the UnsupportedTypeError_starGenType type
type MoqUnsupportedTypeError_starGenType_Error_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqUnsupportedTypeError_starGenType_Error_resultsByParams contains the
// results for a given set of parameters for the
// UnsupportedTypeError_starGenType type
type MoqUnsupportedTypeError_starGenType_Error_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqUnsupportedTypeError_starGenType_Error_paramsKey]*MoqUnsupportedTypeError_starGenType_Error_results
}

// MoqUnsupportedTypeError_starGenType_Error_doFn defines the type of function
// needed when calling AndDo for the UnsupportedTypeError_starGenType type
type MoqUnsupportedTypeError_starGenType_Error_doFn func()

// MoqUnsupportedTypeError_starGenType_Error_doReturnFn defines the type of
// function needed when calling DoReturnResults for the
// UnsupportedTypeError_starGenType type
type MoqUnsupportedTypeError_starGenType_Error_doReturnFn func() string

// MoqUnsupportedTypeError_starGenType_Error_results holds the results of the
// UnsupportedTypeError_starGenType type
type MoqUnsupportedTypeError_starGenType_Error_results struct {
	Params  MoqUnsupportedTypeError_starGenType_Error_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqUnsupportedTypeError_starGenType_Error_doFn
		DoReturnFn MoqUnsupportedTypeError_starGenType_Error_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqUnsupportedTypeError_starGenType_Error_fnRecorder routes recorded
// function calls to the MoqUnsupportedTypeError_starGenType moq
type MoqUnsupportedTypeError_starGenType_Error_fnRecorder struct {
	Params    MoqUnsupportedTypeError_starGenType_Error_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqUnsupportedTypeError_starGenType_Error_results
	Moq       *MoqUnsupportedTypeError_starGenType
}

// MoqUnsupportedTypeError_starGenType_Error_anyParams isolates the any params
// functions of the UnsupportedTypeError_starGenType type
type MoqUnsupportedTypeError_starGenType_Error_anyParams struct {
	Recorder *MoqUnsupportedTypeError_starGenType_Error_fnRecorder
}

// NewMoqUnsupportedTypeError_starGenType creates a new moq of the
// UnsupportedTypeError_starGenType type
func NewMoqUnsupportedTypeError_starGenType(scene *moq.Scene, config *moq.Config) *MoqUnsupportedTypeError_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqUnsupportedTypeError_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqUnsupportedTypeError_starGenType_mock{},

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

// Mock returns the mock implementation of the UnsupportedTypeError_starGenType
// type
func (m *MoqUnsupportedTypeError_starGenType) Mock() *MoqUnsupportedTypeError_starGenType_mock {
	return m.Moq
}

func (m *MoqUnsupportedTypeError_starGenType_mock) Error() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqUnsupportedTypeError_starGenType_Error_params{}
	var results *MoqUnsupportedTypeError_starGenType_Error_results
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
// UnsupportedTypeError_starGenType type
func (m *MoqUnsupportedTypeError_starGenType) OnCall() *MoqUnsupportedTypeError_starGenType_recorder {
	return &MoqUnsupportedTypeError_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqUnsupportedTypeError_starGenType_recorder) Error() *MoqUnsupportedTypeError_starGenType_Error_fnRecorder {
	return &MoqUnsupportedTypeError_starGenType_Error_fnRecorder{
		Params:   MoqUnsupportedTypeError_starGenType_Error_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqUnsupportedTypeError_starGenType_Error_fnRecorder) Any() *MoqUnsupportedTypeError_starGenType_Error_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	return &MoqUnsupportedTypeError_starGenType_Error_anyParams{Recorder: r}
}

func (r *MoqUnsupportedTypeError_starGenType_Error_fnRecorder) Seq() *MoqUnsupportedTypeError_starGenType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqUnsupportedTypeError_starGenType_Error_fnRecorder) NoSeq() *MoqUnsupportedTypeError_starGenType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqUnsupportedTypeError_starGenType_Error_fnRecorder) ReturnResults(result1 string) *MoqUnsupportedTypeError_starGenType_Error_fnRecorder {
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
		DoFn       MoqUnsupportedTypeError_starGenType_Error_doFn
		DoReturnFn MoqUnsupportedTypeError_starGenType_Error_doReturnFn
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

func (r *MoqUnsupportedTypeError_starGenType_Error_fnRecorder) AndDo(fn MoqUnsupportedTypeError_starGenType_Error_doFn) *MoqUnsupportedTypeError_starGenType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqUnsupportedTypeError_starGenType_Error_fnRecorder) DoReturnResults(fn MoqUnsupportedTypeError_starGenType_Error_doReturnFn) *MoqUnsupportedTypeError_starGenType_Error_fnRecorder {
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
		DoFn       MoqUnsupportedTypeError_starGenType_Error_doFn
		DoReturnFn MoqUnsupportedTypeError_starGenType_Error_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqUnsupportedTypeError_starGenType_Error_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqUnsupportedTypeError_starGenType_Error_resultsByParams
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
		results = &MoqUnsupportedTypeError_starGenType_Error_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqUnsupportedTypeError_starGenType_Error_paramsKey]*MoqUnsupportedTypeError_starGenType_Error_results{},
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
		r.Results = &MoqUnsupportedTypeError_starGenType_Error_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqUnsupportedTypeError_starGenType_Error_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqUnsupportedTypeError_starGenType_Error_fnRecorder {
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
				DoFn       MoqUnsupportedTypeError_starGenType_Error_doFn
				DoReturnFn MoqUnsupportedTypeError_starGenType_Error_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqUnsupportedTypeError_starGenType) PrettyParams_Error(params MoqUnsupportedTypeError_starGenType_Error_params) string {
	return fmt.Sprintf("Error()")
}

func (m *MoqUnsupportedTypeError_starGenType) ParamsKey_Error(params MoqUnsupportedTypeError_starGenType_Error_params, anyParams uint64) MoqUnsupportedTypeError_starGenType_Error_paramsKey {
	m.Scene.T.Helper()
	return MoqUnsupportedTypeError_starGenType_Error_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqUnsupportedTypeError_starGenType) Reset() { m.ResultsByParams_Error = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqUnsupportedTypeError_starGenType) AssertExpectationsMet() {
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

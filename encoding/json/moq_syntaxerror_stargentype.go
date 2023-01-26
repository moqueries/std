// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package json

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that json.SyntaxError_starGenType is
// mocked completely
var _ SyntaxError_starGenType = (*MoqSyntaxError_starGenType_mock)(nil)

// SyntaxError_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type SyntaxError_starGenType interface {
	Error() string
}

// MoqSyntaxError_starGenType holds the state of a moq of the
// SyntaxError_starGenType type
type MoqSyntaxError_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSyntaxError_starGenType_mock

	ResultsByParams_Error []MoqSyntaxError_starGenType_Error_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Error struct{}
		}
	}
}

// MoqSyntaxError_starGenType_mock isolates the mock interface of the
// SyntaxError_starGenType type
type MoqSyntaxError_starGenType_mock struct {
	Moq *MoqSyntaxError_starGenType
}

// MoqSyntaxError_starGenType_recorder isolates the recorder interface of the
// SyntaxError_starGenType type
type MoqSyntaxError_starGenType_recorder struct {
	Moq *MoqSyntaxError_starGenType
}

// MoqSyntaxError_starGenType_Error_params holds the params of the
// SyntaxError_starGenType type
type MoqSyntaxError_starGenType_Error_params struct{}

// MoqSyntaxError_starGenType_Error_paramsKey holds the map key params of the
// SyntaxError_starGenType type
type MoqSyntaxError_starGenType_Error_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqSyntaxError_starGenType_Error_resultsByParams contains the results for a
// given set of parameters for the SyntaxError_starGenType type
type MoqSyntaxError_starGenType_Error_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSyntaxError_starGenType_Error_paramsKey]*MoqSyntaxError_starGenType_Error_results
}

// MoqSyntaxError_starGenType_Error_doFn defines the type of function needed
// when calling AndDo for the SyntaxError_starGenType type
type MoqSyntaxError_starGenType_Error_doFn func()

// MoqSyntaxError_starGenType_Error_doReturnFn defines the type of function
// needed when calling DoReturnResults for the SyntaxError_starGenType type
type MoqSyntaxError_starGenType_Error_doReturnFn func() string

// MoqSyntaxError_starGenType_Error_results holds the results of the
// SyntaxError_starGenType type
type MoqSyntaxError_starGenType_Error_results struct {
	Params  MoqSyntaxError_starGenType_Error_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqSyntaxError_starGenType_Error_doFn
		DoReturnFn MoqSyntaxError_starGenType_Error_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSyntaxError_starGenType_Error_fnRecorder routes recorded function calls
// to the MoqSyntaxError_starGenType moq
type MoqSyntaxError_starGenType_Error_fnRecorder struct {
	Params    MoqSyntaxError_starGenType_Error_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSyntaxError_starGenType_Error_results
	Moq       *MoqSyntaxError_starGenType
}

// MoqSyntaxError_starGenType_Error_anyParams isolates the any params functions
// of the SyntaxError_starGenType type
type MoqSyntaxError_starGenType_Error_anyParams struct {
	Recorder *MoqSyntaxError_starGenType_Error_fnRecorder
}

// NewMoqSyntaxError_starGenType creates a new moq of the
// SyntaxError_starGenType type
func NewMoqSyntaxError_starGenType(scene *moq.Scene, config *moq.Config) *MoqSyntaxError_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSyntaxError_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSyntaxError_starGenType_mock{},

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

// Mock returns the mock implementation of the SyntaxError_starGenType type
func (m *MoqSyntaxError_starGenType) Mock() *MoqSyntaxError_starGenType_mock { return m.Moq }

func (m *MoqSyntaxError_starGenType_mock) Error() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqSyntaxError_starGenType_Error_params{}
	var results *MoqSyntaxError_starGenType_Error_results
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

// OnCall returns the recorder implementation of the SyntaxError_starGenType
// type
func (m *MoqSyntaxError_starGenType) OnCall() *MoqSyntaxError_starGenType_recorder {
	return &MoqSyntaxError_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqSyntaxError_starGenType_recorder) Error() *MoqSyntaxError_starGenType_Error_fnRecorder {
	return &MoqSyntaxError_starGenType_Error_fnRecorder{
		Params:   MoqSyntaxError_starGenType_Error_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqSyntaxError_starGenType_Error_fnRecorder) Any() *MoqSyntaxError_starGenType_Error_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	return &MoqSyntaxError_starGenType_Error_anyParams{Recorder: r}
}

func (r *MoqSyntaxError_starGenType_Error_fnRecorder) Seq() *MoqSyntaxError_starGenType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSyntaxError_starGenType_Error_fnRecorder) NoSeq() *MoqSyntaxError_starGenType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSyntaxError_starGenType_Error_fnRecorder) ReturnResults(result1 string) *MoqSyntaxError_starGenType_Error_fnRecorder {
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
		DoFn       MoqSyntaxError_starGenType_Error_doFn
		DoReturnFn MoqSyntaxError_starGenType_Error_doReturnFn
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

func (r *MoqSyntaxError_starGenType_Error_fnRecorder) AndDo(fn MoqSyntaxError_starGenType_Error_doFn) *MoqSyntaxError_starGenType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSyntaxError_starGenType_Error_fnRecorder) DoReturnResults(fn MoqSyntaxError_starGenType_Error_doReturnFn) *MoqSyntaxError_starGenType_Error_fnRecorder {
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
		DoFn       MoqSyntaxError_starGenType_Error_doFn
		DoReturnFn MoqSyntaxError_starGenType_Error_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSyntaxError_starGenType_Error_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSyntaxError_starGenType_Error_resultsByParams
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
		results = &MoqSyntaxError_starGenType_Error_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSyntaxError_starGenType_Error_paramsKey]*MoqSyntaxError_starGenType_Error_results{},
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
		r.Results = &MoqSyntaxError_starGenType_Error_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSyntaxError_starGenType_Error_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSyntaxError_starGenType_Error_fnRecorder {
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
				DoFn       MoqSyntaxError_starGenType_Error_doFn
				DoReturnFn MoqSyntaxError_starGenType_Error_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSyntaxError_starGenType) PrettyParams_Error(params MoqSyntaxError_starGenType_Error_params) string {
	return fmt.Sprintf("Error()")
}

func (m *MoqSyntaxError_starGenType) ParamsKey_Error(params MoqSyntaxError_starGenType_Error_params, anyParams uint64) MoqSyntaxError_starGenType_Error_paramsKey {
	m.Scene.T.Helper()
	return MoqSyntaxError_starGenType_Error_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqSyntaxError_starGenType) Reset() { m.ResultsByParams_Error = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSyntaxError_starGenType) AssertExpectationsMet() {
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
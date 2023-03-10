// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package base32

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that base32.CorruptInputError_genType
// is mocked completely
var _ CorruptInputError_genType = (*MoqCorruptInputError_genType_mock)(nil)

// CorruptInputError_genType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type CorruptInputError_genType interface {
	Error() string
}

// MoqCorruptInputError_genType holds the state of a moq of the
// CorruptInputError_genType type
type MoqCorruptInputError_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqCorruptInputError_genType_mock

	ResultsByParams_Error []MoqCorruptInputError_genType_Error_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Error struct{}
		}
	}
}

// MoqCorruptInputError_genType_mock isolates the mock interface of the
// CorruptInputError_genType type
type MoqCorruptInputError_genType_mock struct {
	Moq *MoqCorruptInputError_genType
}

// MoqCorruptInputError_genType_recorder isolates the recorder interface of the
// CorruptInputError_genType type
type MoqCorruptInputError_genType_recorder struct {
	Moq *MoqCorruptInputError_genType
}

// MoqCorruptInputError_genType_Error_params holds the params of the
// CorruptInputError_genType type
type MoqCorruptInputError_genType_Error_params struct{}

// MoqCorruptInputError_genType_Error_paramsKey holds the map key params of the
// CorruptInputError_genType type
type MoqCorruptInputError_genType_Error_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqCorruptInputError_genType_Error_resultsByParams contains the results for
// a given set of parameters for the CorruptInputError_genType type
type MoqCorruptInputError_genType_Error_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqCorruptInputError_genType_Error_paramsKey]*MoqCorruptInputError_genType_Error_results
}

// MoqCorruptInputError_genType_Error_doFn defines the type of function needed
// when calling AndDo for the CorruptInputError_genType type
type MoqCorruptInputError_genType_Error_doFn func()

// MoqCorruptInputError_genType_Error_doReturnFn defines the type of function
// needed when calling DoReturnResults for the CorruptInputError_genType type
type MoqCorruptInputError_genType_Error_doReturnFn func() string

// MoqCorruptInputError_genType_Error_results holds the results of the
// CorruptInputError_genType type
type MoqCorruptInputError_genType_Error_results struct {
	Params  MoqCorruptInputError_genType_Error_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqCorruptInputError_genType_Error_doFn
		DoReturnFn MoqCorruptInputError_genType_Error_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqCorruptInputError_genType_Error_fnRecorder routes recorded function calls
// to the MoqCorruptInputError_genType moq
type MoqCorruptInputError_genType_Error_fnRecorder struct {
	Params    MoqCorruptInputError_genType_Error_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqCorruptInputError_genType_Error_results
	Moq       *MoqCorruptInputError_genType
}

// MoqCorruptInputError_genType_Error_anyParams isolates the any params
// functions of the CorruptInputError_genType type
type MoqCorruptInputError_genType_Error_anyParams struct {
	Recorder *MoqCorruptInputError_genType_Error_fnRecorder
}

// NewMoqCorruptInputError_genType creates a new moq of the
// CorruptInputError_genType type
func NewMoqCorruptInputError_genType(scene *moq.Scene, config *moq.Config) *MoqCorruptInputError_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqCorruptInputError_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqCorruptInputError_genType_mock{},

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

// Mock returns the mock implementation of the CorruptInputError_genType type
func (m *MoqCorruptInputError_genType) Mock() *MoqCorruptInputError_genType_mock { return m.Moq }

func (m *MoqCorruptInputError_genType_mock) Error() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqCorruptInputError_genType_Error_params{}
	var results *MoqCorruptInputError_genType_Error_results
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

// OnCall returns the recorder implementation of the CorruptInputError_genType
// type
func (m *MoqCorruptInputError_genType) OnCall() *MoqCorruptInputError_genType_recorder {
	return &MoqCorruptInputError_genType_recorder{
		Moq: m,
	}
}

func (m *MoqCorruptInputError_genType_recorder) Error() *MoqCorruptInputError_genType_Error_fnRecorder {
	return &MoqCorruptInputError_genType_Error_fnRecorder{
		Params:   MoqCorruptInputError_genType_Error_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqCorruptInputError_genType_Error_fnRecorder) Any() *MoqCorruptInputError_genType_Error_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	return &MoqCorruptInputError_genType_Error_anyParams{Recorder: r}
}

func (r *MoqCorruptInputError_genType_Error_fnRecorder) Seq() *MoqCorruptInputError_genType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqCorruptInputError_genType_Error_fnRecorder) NoSeq() *MoqCorruptInputError_genType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqCorruptInputError_genType_Error_fnRecorder) ReturnResults(result1 string) *MoqCorruptInputError_genType_Error_fnRecorder {
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
		DoFn       MoqCorruptInputError_genType_Error_doFn
		DoReturnFn MoqCorruptInputError_genType_Error_doReturnFn
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

func (r *MoqCorruptInputError_genType_Error_fnRecorder) AndDo(fn MoqCorruptInputError_genType_Error_doFn) *MoqCorruptInputError_genType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqCorruptInputError_genType_Error_fnRecorder) DoReturnResults(fn MoqCorruptInputError_genType_Error_doReturnFn) *MoqCorruptInputError_genType_Error_fnRecorder {
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
		DoFn       MoqCorruptInputError_genType_Error_doFn
		DoReturnFn MoqCorruptInputError_genType_Error_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqCorruptInputError_genType_Error_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqCorruptInputError_genType_Error_resultsByParams
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
		results = &MoqCorruptInputError_genType_Error_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqCorruptInputError_genType_Error_paramsKey]*MoqCorruptInputError_genType_Error_results{},
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
		r.Results = &MoqCorruptInputError_genType_Error_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqCorruptInputError_genType_Error_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqCorruptInputError_genType_Error_fnRecorder {
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
				DoFn       MoqCorruptInputError_genType_Error_doFn
				DoReturnFn MoqCorruptInputError_genType_Error_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqCorruptInputError_genType) PrettyParams_Error(params MoqCorruptInputError_genType_Error_params) string {
	return fmt.Sprintf("Error()")
}

func (m *MoqCorruptInputError_genType) ParamsKey_Error(params MoqCorruptInputError_genType_Error_params, anyParams uint64) MoqCorruptInputError_genType_Error_paramsKey {
	m.Scene.T.Helper()
	return MoqCorruptInputError_genType_Error_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqCorruptInputError_genType) Reset() { m.ResultsByParams_Error = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqCorruptInputError_genType) AssertExpectationsMet() {
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

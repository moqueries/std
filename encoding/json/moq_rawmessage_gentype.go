// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package json

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that json.RawMessage_genType is mocked
// completely
var _ RawMessage_genType = (*MoqRawMessage_genType_mock)(nil)

// RawMessage_genType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type RawMessage_genType interface {
	MarshalJSON() ([]byte, error)
}

// MoqRawMessage_genType holds the state of a moq of the RawMessage_genType
// type
type MoqRawMessage_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqRawMessage_genType_mock

	ResultsByParams_MarshalJSON []MoqRawMessage_genType_MarshalJSON_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			MarshalJSON struct{}
		}
	}
}

// MoqRawMessage_genType_mock isolates the mock interface of the
// RawMessage_genType type
type MoqRawMessage_genType_mock struct {
	Moq *MoqRawMessage_genType
}

// MoqRawMessage_genType_recorder isolates the recorder interface of the
// RawMessage_genType type
type MoqRawMessage_genType_recorder struct {
	Moq *MoqRawMessage_genType
}

// MoqRawMessage_genType_MarshalJSON_params holds the params of the
// RawMessage_genType type
type MoqRawMessage_genType_MarshalJSON_params struct{}

// MoqRawMessage_genType_MarshalJSON_paramsKey holds the map key params of the
// RawMessage_genType type
type MoqRawMessage_genType_MarshalJSON_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqRawMessage_genType_MarshalJSON_resultsByParams contains the results for a
// given set of parameters for the RawMessage_genType type
type MoqRawMessage_genType_MarshalJSON_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqRawMessage_genType_MarshalJSON_paramsKey]*MoqRawMessage_genType_MarshalJSON_results
}

// MoqRawMessage_genType_MarshalJSON_doFn defines the type of function needed
// when calling AndDo for the RawMessage_genType type
type MoqRawMessage_genType_MarshalJSON_doFn func()

// MoqRawMessage_genType_MarshalJSON_doReturnFn defines the type of function
// needed when calling DoReturnResults for the RawMessage_genType type
type MoqRawMessage_genType_MarshalJSON_doReturnFn func() ([]byte, error)

// MoqRawMessage_genType_MarshalJSON_results holds the results of the
// RawMessage_genType type
type MoqRawMessage_genType_MarshalJSON_results struct {
	Params  MoqRawMessage_genType_MarshalJSON_params
	Results []struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqRawMessage_genType_MarshalJSON_doFn
		DoReturnFn MoqRawMessage_genType_MarshalJSON_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqRawMessage_genType_MarshalJSON_fnRecorder routes recorded function calls
// to the MoqRawMessage_genType moq
type MoqRawMessage_genType_MarshalJSON_fnRecorder struct {
	Params    MoqRawMessage_genType_MarshalJSON_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqRawMessage_genType_MarshalJSON_results
	Moq       *MoqRawMessage_genType
}

// MoqRawMessage_genType_MarshalJSON_anyParams isolates the any params
// functions of the RawMessage_genType type
type MoqRawMessage_genType_MarshalJSON_anyParams struct {
	Recorder *MoqRawMessage_genType_MarshalJSON_fnRecorder
}

// NewMoqRawMessage_genType creates a new moq of the RawMessage_genType type
func NewMoqRawMessage_genType(scene *moq.Scene, config *moq.Config) *MoqRawMessage_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqRawMessage_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqRawMessage_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				MarshalJSON struct{}
			}
		}{ParameterIndexing: struct {
			MarshalJSON struct{}
		}{
			MarshalJSON: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the RawMessage_genType type
func (m *MoqRawMessage_genType) Mock() *MoqRawMessage_genType_mock { return m.Moq }

func (m *MoqRawMessage_genType_mock) MarshalJSON() (result1 []byte, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqRawMessage_genType_MarshalJSON_params{}
	var results *MoqRawMessage_genType_MarshalJSON_results
	for _, resultsByParams := range m.Moq.ResultsByParams_MarshalJSON {
		paramsKey := m.Moq.ParamsKey_MarshalJSON(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_MarshalJSON(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_MarshalJSON(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_MarshalJSON(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn()
	}
	return
}

// OnCall returns the recorder implementation of the RawMessage_genType type
func (m *MoqRawMessage_genType) OnCall() *MoqRawMessage_genType_recorder {
	return &MoqRawMessage_genType_recorder{
		Moq: m,
	}
}

func (m *MoqRawMessage_genType_recorder) MarshalJSON() *MoqRawMessage_genType_MarshalJSON_fnRecorder {
	return &MoqRawMessage_genType_MarshalJSON_fnRecorder{
		Params:   MoqRawMessage_genType_MarshalJSON_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqRawMessage_genType_MarshalJSON_fnRecorder) Any() *MoqRawMessage_genType_MarshalJSON_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_MarshalJSON(r.Params))
		return nil
	}
	return &MoqRawMessage_genType_MarshalJSON_anyParams{Recorder: r}
}

func (r *MoqRawMessage_genType_MarshalJSON_fnRecorder) Seq() *MoqRawMessage_genType_MarshalJSON_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_MarshalJSON(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqRawMessage_genType_MarshalJSON_fnRecorder) NoSeq() *MoqRawMessage_genType_MarshalJSON_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_MarshalJSON(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqRawMessage_genType_MarshalJSON_fnRecorder) ReturnResults(result1 []byte, result2 error) *MoqRawMessage_genType_MarshalJSON_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqRawMessage_genType_MarshalJSON_doFn
		DoReturnFn MoqRawMessage_genType_MarshalJSON_doReturnFn
	}{
		Values: &struct {
			Result1 []byte
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqRawMessage_genType_MarshalJSON_fnRecorder) AndDo(fn MoqRawMessage_genType_MarshalJSON_doFn) *MoqRawMessage_genType_MarshalJSON_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqRawMessage_genType_MarshalJSON_fnRecorder) DoReturnResults(fn MoqRawMessage_genType_MarshalJSON_doReturnFn) *MoqRawMessage_genType_MarshalJSON_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqRawMessage_genType_MarshalJSON_doFn
		DoReturnFn MoqRawMessage_genType_MarshalJSON_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqRawMessage_genType_MarshalJSON_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqRawMessage_genType_MarshalJSON_resultsByParams
	for n, res := range r.Moq.ResultsByParams_MarshalJSON {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqRawMessage_genType_MarshalJSON_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqRawMessage_genType_MarshalJSON_paramsKey]*MoqRawMessage_genType_MarshalJSON_results{},
		}
		r.Moq.ResultsByParams_MarshalJSON = append(r.Moq.ResultsByParams_MarshalJSON, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_MarshalJSON) {
			copy(r.Moq.ResultsByParams_MarshalJSON[insertAt+1:], r.Moq.ResultsByParams_MarshalJSON[insertAt:0])
			r.Moq.ResultsByParams_MarshalJSON[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_MarshalJSON(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqRawMessage_genType_MarshalJSON_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqRawMessage_genType_MarshalJSON_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqRawMessage_genType_MarshalJSON_fnRecorder {
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
					Result1 []byte
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqRawMessage_genType_MarshalJSON_doFn
				DoReturnFn MoqRawMessage_genType_MarshalJSON_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqRawMessage_genType) PrettyParams_MarshalJSON(params MoqRawMessage_genType_MarshalJSON_params) string {
	return fmt.Sprintf("MarshalJSON()")
}

func (m *MoqRawMessage_genType) ParamsKey_MarshalJSON(params MoqRawMessage_genType_MarshalJSON_params, anyParams uint64) MoqRawMessage_genType_MarshalJSON_paramsKey {
	m.Scene.T.Helper()
	return MoqRawMessage_genType_MarshalJSON_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqRawMessage_genType) Reset() { m.ResultsByParams_MarshalJSON = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqRawMessage_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_MarshalJSON {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_MarshalJSON(results.Params))
			}
		}
	}
}

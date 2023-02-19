// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package sql

import (
	"database/sql/driver"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that sql.NullInt64_genType is mocked
// completely
var _ NullInt64_genType = (*MoqNullInt64_genType_mock)(nil)

// NullInt64_genType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type NullInt64_genType interface {
	Value() (driver.Value, error)
}

// MoqNullInt64_genType holds the state of a moq of the NullInt64_genType type
type MoqNullInt64_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNullInt64_genType_mock

	ResultsByParams_Value []MoqNullInt64_genType_Value_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Value struct{}
		}
	}
}

// MoqNullInt64_genType_mock isolates the mock interface of the
// NullInt64_genType type
type MoqNullInt64_genType_mock struct {
	Moq *MoqNullInt64_genType
}

// MoqNullInt64_genType_recorder isolates the recorder interface of the
// NullInt64_genType type
type MoqNullInt64_genType_recorder struct {
	Moq *MoqNullInt64_genType
}

// MoqNullInt64_genType_Value_params holds the params of the NullInt64_genType
// type
type MoqNullInt64_genType_Value_params struct{}

// MoqNullInt64_genType_Value_paramsKey holds the map key params of the
// NullInt64_genType type
type MoqNullInt64_genType_Value_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqNullInt64_genType_Value_resultsByParams contains the results for a given
// set of parameters for the NullInt64_genType type
type MoqNullInt64_genType_Value_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNullInt64_genType_Value_paramsKey]*MoqNullInt64_genType_Value_results
}

// MoqNullInt64_genType_Value_doFn defines the type of function needed when
// calling AndDo for the NullInt64_genType type
type MoqNullInt64_genType_Value_doFn func()

// MoqNullInt64_genType_Value_doReturnFn defines the type of function needed
// when calling DoReturnResults for the NullInt64_genType type
type MoqNullInt64_genType_Value_doReturnFn func() (driver.Value, error)

// MoqNullInt64_genType_Value_results holds the results of the
// NullInt64_genType type
type MoqNullInt64_genType_Value_results struct {
	Params  MoqNullInt64_genType_Value_params
	Results []struct {
		Values *struct {
			Result1 driver.Value
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqNullInt64_genType_Value_doFn
		DoReturnFn MoqNullInt64_genType_Value_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNullInt64_genType_Value_fnRecorder routes recorded function calls to the
// MoqNullInt64_genType moq
type MoqNullInt64_genType_Value_fnRecorder struct {
	Params    MoqNullInt64_genType_Value_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNullInt64_genType_Value_results
	Moq       *MoqNullInt64_genType
}

// MoqNullInt64_genType_Value_anyParams isolates the any params functions of
// the NullInt64_genType type
type MoqNullInt64_genType_Value_anyParams struct {
	Recorder *MoqNullInt64_genType_Value_fnRecorder
}

// NewMoqNullInt64_genType creates a new moq of the NullInt64_genType type
func NewMoqNullInt64_genType(scene *moq.Scene, config *moq.Config) *MoqNullInt64_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNullInt64_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNullInt64_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Value struct{}
			}
		}{ParameterIndexing: struct {
			Value struct{}
		}{
			Value: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the NullInt64_genType type
func (m *MoqNullInt64_genType) Mock() *MoqNullInt64_genType_mock { return m.Moq }

func (m *MoqNullInt64_genType_mock) Value() (result1 driver.Value, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqNullInt64_genType_Value_params{}
	var results *MoqNullInt64_genType_Value_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Value {
		paramsKey := m.Moq.ParamsKey_Value(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Value(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Value(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Value(params))
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

// OnCall returns the recorder implementation of the NullInt64_genType type
func (m *MoqNullInt64_genType) OnCall() *MoqNullInt64_genType_recorder {
	return &MoqNullInt64_genType_recorder{
		Moq: m,
	}
}

func (m *MoqNullInt64_genType_recorder) Value() *MoqNullInt64_genType_Value_fnRecorder {
	return &MoqNullInt64_genType_Value_fnRecorder{
		Params:   MoqNullInt64_genType_Value_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqNullInt64_genType_Value_fnRecorder) Any() *MoqNullInt64_genType_Value_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Value(r.Params))
		return nil
	}
	return &MoqNullInt64_genType_Value_anyParams{Recorder: r}
}

func (r *MoqNullInt64_genType_Value_fnRecorder) Seq() *MoqNullInt64_genType_Value_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Value(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNullInt64_genType_Value_fnRecorder) NoSeq() *MoqNullInt64_genType_Value_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Value(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNullInt64_genType_Value_fnRecorder) ReturnResults(result1 driver.Value, result2 error) *MoqNullInt64_genType_Value_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 driver.Value
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqNullInt64_genType_Value_doFn
		DoReturnFn MoqNullInt64_genType_Value_doReturnFn
	}{
		Values: &struct {
			Result1 driver.Value
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNullInt64_genType_Value_fnRecorder) AndDo(fn MoqNullInt64_genType_Value_doFn) *MoqNullInt64_genType_Value_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNullInt64_genType_Value_fnRecorder) DoReturnResults(fn MoqNullInt64_genType_Value_doReturnFn) *MoqNullInt64_genType_Value_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 driver.Value
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqNullInt64_genType_Value_doFn
		DoReturnFn MoqNullInt64_genType_Value_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNullInt64_genType_Value_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNullInt64_genType_Value_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Value {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqNullInt64_genType_Value_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNullInt64_genType_Value_paramsKey]*MoqNullInt64_genType_Value_results{},
		}
		r.Moq.ResultsByParams_Value = append(r.Moq.ResultsByParams_Value, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Value) {
			copy(r.Moq.ResultsByParams_Value[insertAt+1:], r.Moq.ResultsByParams_Value[insertAt:0])
			r.Moq.ResultsByParams_Value[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Value(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqNullInt64_genType_Value_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNullInt64_genType_Value_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNullInt64_genType_Value_fnRecorder {
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
					Result1 driver.Value
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqNullInt64_genType_Value_doFn
				DoReturnFn MoqNullInt64_genType_Value_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNullInt64_genType) PrettyParams_Value(params MoqNullInt64_genType_Value_params) string {
	return fmt.Sprintf("Value()")
}

func (m *MoqNullInt64_genType) ParamsKey_Value(params MoqNullInt64_genType_Value_params, anyParams uint64) MoqNullInt64_genType_Value_paramsKey {
	m.Scene.T.Helper()
	return MoqNullInt64_genType_Value_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqNullInt64_genType) Reset() { m.ResultsByParams_Value = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNullInt64_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Value {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Value(results.Params))
			}
		}
	}
}
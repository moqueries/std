// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package sql

import (
	"database/sql/driver"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that sql.NullString_genType is mocked
// completely
var _ NullString_genType = (*MoqNullString_genType_mock)(nil)

// NullString_genType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type NullString_genType interface {
	Value() (driver.Value, error)
}

// MoqNullString_genType holds the state of a moq of the NullString_genType
// type
type MoqNullString_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNullString_genType_mock

	ResultsByParams_Value []MoqNullString_genType_Value_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Value struct{}
		}
	}
}

// MoqNullString_genType_mock isolates the mock interface of the
// NullString_genType type
type MoqNullString_genType_mock struct {
	Moq *MoqNullString_genType
}

// MoqNullString_genType_recorder isolates the recorder interface of the
// NullString_genType type
type MoqNullString_genType_recorder struct {
	Moq *MoqNullString_genType
}

// MoqNullString_genType_Value_params holds the params of the
// NullString_genType type
type MoqNullString_genType_Value_params struct{}

// MoqNullString_genType_Value_paramsKey holds the map key params of the
// NullString_genType type
type MoqNullString_genType_Value_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqNullString_genType_Value_resultsByParams contains the results for a given
// set of parameters for the NullString_genType type
type MoqNullString_genType_Value_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNullString_genType_Value_paramsKey]*MoqNullString_genType_Value_results
}

// MoqNullString_genType_Value_doFn defines the type of function needed when
// calling AndDo for the NullString_genType type
type MoqNullString_genType_Value_doFn func()

// MoqNullString_genType_Value_doReturnFn defines the type of function needed
// when calling DoReturnResults for the NullString_genType type
type MoqNullString_genType_Value_doReturnFn func() (driver.Value, error)

// MoqNullString_genType_Value_results holds the results of the
// NullString_genType type
type MoqNullString_genType_Value_results struct {
	Params  MoqNullString_genType_Value_params
	Results []struct {
		Values *struct {
			Result1 driver.Value
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqNullString_genType_Value_doFn
		DoReturnFn MoqNullString_genType_Value_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNullString_genType_Value_fnRecorder routes recorded function calls to the
// MoqNullString_genType moq
type MoqNullString_genType_Value_fnRecorder struct {
	Params    MoqNullString_genType_Value_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNullString_genType_Value_results
	Moq       *MoqNullString_genType
}

// MoqNullString_genType_Value_anyParams isolates the any params functions of
// the NullString_genType type
type MoqNullString_genType_Value_anyParams struct {
	Recorder *MoqNullString_genType_Value_fnRecorder
}

// NewMoqNullString_genType creates a new moq of the NullString_genType type
func NewMoqNullString_genType(scene *moq.Scene, config *moq.Config) *MoqNullString_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNullString_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNullString_genType_mock{},

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

// Mock returns the mock implementation of the NullString_genType type
func (m *MoqNullString_genType) Mock() *MoqNullString_genType_mock { return m.Moq }

func (m *MoqNullString_genType_mock) Value() (result1 driver.Value, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqNullString_genType_Value_params{}
	var results *MoqNullString_genType_Value_results
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

// OnCall returns the recorder implementation of the NullString_genType type
func (m *MoqNullString_genType) OnCall() *MoqNullString_genType_recorder {
	return &MoqNullString_genType_recorder{
		Moq: m,
	}
}

func (m *MoqNullString_genType_recorder) Value() *MoqNullString_genType_Value_fnRecorder {
	return &MoqNullString_genType_Value_fnRecorder{
		Params:   MoqNullString_genType_Value_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqNullString_genType_Value_fnRecorder) Any() *MoqNullString_genType_Value_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Value(r.Params))
		return nil
	}
	return &MoqNullString_genType_Value_anyParams{Recorder: r}
}

func (r *MoqNullString_genType_Value_fnRecorder) Seq() *MoqNullString_genType_Value_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Value(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNullString_genType_Value_fnRecorder) NoSeq() *MoqNullString_genType_Value_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Value(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNullString_genType_Value_fnRecorder) ReturnResults(result1 driver.Value, result2 error) *MoqNullString_genType_Value_fnRecorder {
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
		DoFn       MoqNullString_genType_Value_doFn
		DoReturnFn MoqNullString_genType_Value_doReturnFn
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

func (r *MoqNullString_genType_Value_fnRecorder) AndDo(fn MoqNullString_genType_Value_doFn) *MoqNullString_genType_Value_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNullString_genType_Value_fnRecorder) DoReturnResults(fn MoqNullString_genType_Value_doReturnFn) *MoqNullString_genType_Value_fnRecorder {
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
		DoFn       MoqNullString_genType_Value_doFn
		DoReturnFn MoqNullString_genType_Value_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNullString_genType_Value_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNullString_genType_Value_resultsByParams
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
		results = &MoqNullString_genType_Value_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNullString_genType_Value_paramsKey]*MoqNullString_genType_Value_results{},
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
		r.Results = &MoqNullString_genType_Value_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNullString_genType_Value_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNullString_genType_Value_fnRecorder {
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
				DoFn       MoqNullString_genType_Value_doFn
				DoReturnFn MoqNullString_genType_Value_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNullString_genType) PrettyParams_Value(params MoqNullString_genType_Value_params) string {
	return fmt.Sprintf("Value()")
}

func (m *MoqNullString_genType) ParamsKey_Value(params MoqNullString_genType_Value_params, anyParams uint64) MoqNullString_genType_Value_paramsKey {
	m.Scene.T.Helper()
	return MoqNullString_genType_Value_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqNullString_genType) Reset() { m.ResultsByParams_Value = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNullString_genType) AssertExpectationsMet() {
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

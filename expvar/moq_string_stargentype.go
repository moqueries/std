// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package expvar

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that expvar.String_starGenType is
// mocked completely
var _ String_starGenType = (*MoqString_starGenType_mock)(nil)

// String_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type String_starGenType interface {
	Value() string
	String() string
	Set(value string)
}

// MoqString_starGenType holds the state of a moq of the String_starGenType
// type
type MoqString_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqString_starGenType_mock

	ResultsByParams_Value  []MoqString_starGenType_Value_resultsByParams
	ResultsByParams_String []MoqString_starGenType_String_resultsByParams
	ResultsByParams_Set    []MoqString_starGenType_Set_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Value  struct{}
			String struct{}
			Set    struct {
				Value moq.ParamIndexing
			}
		}
	}
	// MoqString_starGenType_mock isolates the mock interface of the
}

// String_starGenType type
type MoqString_starGenType_mock struct {
	Moq *MoqString_starGenType
}

// MoqString_starGenType_recorder isolates the recorder interface of the
// String_starGenType type
type MoqString_starGenType_recorder struct {
	Moq *MoqString_starGenType
}

// MoqString_starGenType_Value_params holds the params of the
// String_starGenType type
type MoqString_starGenType_Value_params struct{}

// MoqString_starGenType_Value_paramsKey holds the map key params of the
// String_starGenType type
type MoqString_starGenType_Value_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqString_starGenType_Value_resultsByParams contains the results for a given
// set of parameters for the String_starGenType type
type MoqString_starGenType_Value_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqString_starGenType_Value_paramsKey]*MoqString_starGenType_Value_results
}

// MoqString_starGenType_Value_doFn defines the type of function needed when
// calling AndDo for the String_starGenType type
type MoqString_starGenType_Value_doFn func()

// MoqString_starGenType_Value_doReturnFn defines the type of function needed
// when calling DoReturnResults for the String_starGenType type
type MoqString_starGenType_Value_doReturnFn func() string

// MoqString_starGenType_Value_results holds the results of the
// String_starGenType type
type MoqString_starGenType_Value_results struct {
	Params  MoqString_starGenType_Value_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqString_starGenType_Value_doFn
		DoReturnFn MoqString_starGenType_Value_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqString_starGenType_Value_fnRecorder routes recorded function calls to the
// MoqString_starGenType moq
type MoqString_starGenType_Value_fnRecorder struct {
	Params    MoqString_starGenType_Value_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqString_starGenType_Value_results
	Moq       *MoqString_starGenType
}

// MoqString_starGenType_Value_anyParams isolates the any params functions of
// the String_starGenType type
type MoqString_starGenType_Value_anyParams struct {
	Recorder *MoqString_starGenType_Value_fnRecorder
}

// MoqString_starGenType_String_params holds the params of the
// String_starGenType type
type MoqString_starGenType_String_params struct{}

// MoqString_starGenType_String_paramsKey holds the map key params of the
// String_starGenType type
type MoqString_starGenType_String_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqString_starGenType_String_resultsByParams contains the results for a
// given set of parameters for the String_starGenType type
type MoqString_starGenType_String_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqString_starGenType_String_paramsKey]*MoqString_starGenType_String_results
}

// MoqString_starGenType_String_doFn defines the type of function needed when
// calling AndDo for the String_starGenType type
type MoqString_starGenType_String_doFn func()

// MoqString_starGenType_String_doReturnFn defines the type of function needed
// when calling DoReturnResults for the String_starGenType type
type MoqString_starGenType_String_doReturnFn func() string

// MoqString_starGenType_String_results holds the results of the
// String_starGenType type
type MoqString_starGenType_String_results struct {
	Params  MoqString_starGenType_String_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqString_starGenType_String_doFn
		DoReturnFn MoqString_starGenType_String_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqString_starGenType_String_fnRecorder routes recorded function calls to
// the MoqString_starGenType moq
type MoqString_starGenType_String_fnRecorder struct {
	Params    MoqString_starGenType_String_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqString_starGenType_String_results
	Moq       *MoqString_starGenType
}

// MoqString_starGenType_String_anyParams isolates the any params functions of
// the String_starGenType type
type MoqString_starGenType_String_anyParams struct {
	Recorder *MoqString_starGenType_String_fnRecorder
}

// MoqString_starGenType_Set_params holds the params of the String_starGenType
// type
type MoqString_starGenType_Set_params struct{ Value string }

// MoqString_starGenType_Set_paramsKey holds the map key params of the
// String_starGenType type
type MoqString_starGenType_Set_paramsKey struct {
	Params struct{ Value string }
	Hashes struct{ Value hash.Hash }
}

// MoqString_starGenType_Set_resultsByParams contains the results for a given
// set of parameters for the String_starGenType type
type MoqString_starGenType_Set_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqString_starGenType_Set_paramsKey]*MoqString_starGenType_Set_results
}

// MoqString_starGenType_Set_doFn defines the type of function needed when
// calling AndDo for the String_starGenType type
type MoqString_starGenType_Set_doFn func(value string)

// MoqString_starGenType_Set_doReturnFn defines the type of function needed
// when calling DoReturnResults for the String_starGenType type
type MoqString_starGenType_Set_doReturnFn func(value string)

// MoqString_starGenType_Set_results holds the results of the
// String_starGenType type
type MoqString_starGenType_Set_results struct {
	Params  MoqString_starGenType_Set_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqString_starGenType_Set_doFn
		DoReturnFn MoqString_starGenType_Set_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqString_starGenType_Set_fnRecorder routes recorded function calls to the
// MoqString_starGenType moq
type MoqString_starGenType_Set_fnRecorder struct {
	Params    MoqString_starGenType_Set_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqString_starGenType_Set_results
	Moq       *MoqString_starGenType
}

// MoqString_starGenType_Set_anyParams isolates the any params functions of the
// String_starGenType type
type MoqString_starGenType_Set_anyParams struct {
	Recorder *MoqString_starGenType_Set_fnRecorder
}

// NewMoqString_starGenType creates a new moq of the String_starGenType type
func NewMoqString_starGenType(scene *moq.Scene, config *moq.Config) *MoqString_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqString_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqString_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Value  struct{}
				String struct{}
				Set    struct {
					Value moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Value  struct{}
			String struct{}
			Set    struct {
				Value moq.ParamIndexing
			}
		}{
			Value:  struct{}{},
			String: struct{}{},
			Set: struct {
				Value moq.ParamIndexing
			}{
				Value: moq.ParamIndexByValue,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the String_starGenType type
func (m *MoqString_starGenType) Mock() *MoqString_starGenType_mock { return m.Moq }

func (m *MoqString_starGenType_mock) Value() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqString_starGenType_Value_params{}
	var results *MoqString_starGenType_Value_results
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
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn()
	}
	return
}

func (m *MoqString_starGenType_mock) String() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqString_starGenType_String_params{}
	var results *MoqString_starGenType_String_results
	for _, resultsByParams := range m.Moq.ResultsByParams_String {
		paramsKey := m.Moq.ParamsKey_String(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_String(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_String(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_String(params))
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

func (m *MoqString_starGenType_mock) Set(value string) {
	m.Moq.Scene.T.Helper()
	params := MoqString_starGenType_Set_params{
		Value: value,
	}
	var results *MoqString_starGenType_Set_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Set {
		paramsKey := m.Moq.ParamsKey_Set(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Set(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Set(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Set(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(value)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(value)
	}
	return
}

// OnCall returns the recorder implementation of the String_starGenType type
func (m *MoqString_starGenType) OnCall() *MoqString_starGenType_recorder {
	return &MoqString_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqString_starGenType_recorder) Value() *MoqString_starGenType_Value_fnRecorder {
	return &MoqString_starGenType_Value_fnRecorder{
		Params:   MoqString_starGenType_Value_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqString_starGenType_Value_fnRecorder) Any() *MoqString_starGenType_Value_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Value(r.Params))
		return nil
	}
	return &MoqString_starGenType_Value_anyParams{Recorder: r}
}

func (r *MoqString_starGenType_Value_fnRecorder) Seq() *MoqString_starGenType_Value_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Value(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqString_starGenType_Value_fnRecorder) NoSeq() *MoqString_starGenType_Value_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Value(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqString_starGenType_Value_fnRecorder) ReturnResults(result1 string) *MoqString_starGenType_Value_fnRecorder {
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
		DoFn       MoqString_starGenType_Value_doFn
		DoReturnFn MoqString_starGenType_Value_doReturnFn
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

func (r *MoqString_starGenType_Value_fnRecorder) AndDo(fn MoqString_starGenType_Value_doFn) *MoqString_starGenType_Value_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqString_starGenType_Value_fnRecorder) DoReturnResults(fn MoqString_starGenType_Value_doReturnFn) *MoqString_starGenType_Value_fnRecorder {
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
		DoFn       MoqString_starGenType_Value_doFn
		DoReturnFn MoqString_starGenType_Value_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqString_starGenType_Value_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqString_starGenType_Value_resultsByParams
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
		results = &MoqString_starGenType_Value_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqString_starGenType_Value_paramsKey]*MoqString_starGenType_Value_results{},
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
		r.Results = &MoqString_starGenType_Value_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqString_starGenType_Value_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqString_starGenType_Value_fnRecorder {
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
				DoFn       MoqString_starGenType_Value_doFn
				DoReturnFn MoqString_starGenType_Value_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqString_starGenType) PrettyParams_Value(params MoqString_starGenType_Value_params) string {
	return fmt.Sprintf("Value()")
}

func (m *MoqString_starGenType) ParamsKey_Value(params MoqString_starGenType_Value_params, anyParams uint64) MoqString_starGenType_Value_paramsKey {
	m.Scene.T.Helper()
	return MoqString_starGenType_Value_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqString_starGenType_recorder) String() *MoqString_starGenType_String_fnRecorder {
	return &MoqString_starGenType_String_fnRecorder{
		Params:   MoqString_starGenType_String_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqString_starGenType_String_fnRecorder) Any() *MoqString_starGenType_String_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	return &MoqString_starGenType_String_anyParams{Recorder: r}
}

func (r *MoqString_starGenType_String_fnRecorder) Seq() *MoqString_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqString_starGenType_String_fnRecorder) NoSeq() *MoqString_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqString_starGenType_String_fnRecorder) ReturnResults(result1 string) *MoqString_starGenType_String_fnRecorder {
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
		DoFn       MoqString_starGenType_String_doFn
		DoReturnFn MoqString_starGenType_String_doReturnFn
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

func (r *MoqString_starGenType_String_fnRecorder) AndDo(fn MoqString_starGenType_String_doFn) *MoqString_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqString_starGenType_String_fnRecorder) DoReturnResults(fn MoqString_starGenType_String_doReturnFn) *MoqString_starGenType_String_fnRecorder {
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
		DoFn       MoqString_starGenType_String_doFn
		DoReturnFn MoqString_starGenType_String_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqString_starGenType_String_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqString_starGenType_String_resultsByParams
	for n, res := range r.Moq.ResultsByParams_String {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqString_starGenType_String_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqString_starGenType_String_paramsKey]*MoqString_starGenType_String_results{},
		}
		r.Moq.ResultsByParams_String = append(r.Moq.ResultsByParams_String, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_String) {
			copy(r.Moq.ResultsByParams_String[insertAt+1:], r.Moq.ResultsByParams_String[insertAt:0])
			r.Moq.ResultsByParams_String[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_String(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqString_starGenType_String_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqString_starGenType_String_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqString_starGenType_String_fnRecorder {
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
				DoFn       MoqString_starGenType_String_doFn
				DoReturnFn MoqString_starGenType_String_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqString_starGenType) PrettyParams_String(params MoqString_starGenType_String_params) string {
	return fmt.Sprintf("String()")
}

func (m *MoqString_starGenType) ParamsKey_String(params MoqString_starGenType_String_params, anyParams uint64) MoqString_starGenType_String_paramsKey {
	m.Scene.T.Helper()
	return MoqString_starGenType_String_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqString_starGenType_recorder) Set(value string) *MoqString_starGenType_Set_fnRecorder {
	return &MoqString_starGenType_Set_fnRecorder{
		Params: MoqString_starGenType_Set_params{
			Value: value,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqString_starGenType_Set_fnRecorder) Any() *MoqString_starGenType_Set_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Set(r.Params))
		return nil
	}
	return &MoqString_starGenType_Set_anyParams{Recorder: r}
}

func (a *MoqString_starGenType_Set_anyParams) Value() *MoqString_starGenType_Set_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqString_starGenType_Set_fnRecorder) Seq() *MoqString_starGenType_Set_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Set(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqString_starGenType_Set_fnRecorder) NoSeq() *MoqString_starGenType_Set_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Set(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqString_starGenType_Set_fnRecorder) ReturnResults() *MoqString_starGenType_Set_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqString_starGenType_Set_doFn
		DoReturnFn MoqString_starGenType_Set_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqString_starGenType_Set_fnRecorder) AndDo(fn MoqString_starGenType_Set_doFn) *MoqString_starGenType_Set_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqString_starGenType_Set_fnRecorder) DoReturnResults(fn MoqString_starGenType_Set_doReturnFn) *MoqString_starGenType_Set_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqString_starGenType_Set_doFn
		DoReturnFn MoqString_starGenType_Set_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqString_starGenType_Set_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqString_starGenType_Set_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Set {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqString_starGenType_Set_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqString_starGenType_Set_paramsKey]*MoqString_starGenType_Set_results{},
		}
		r.Moq.ResultsByParams_Set = append(r.Moq.ResultsByParams_Set, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Set) {
			copy(r.Moq.ResultsByParams_Set[insertAt+1:], r.Moq.ResultsByParams_Set[insertAt:0])
			r.Moq.ResultsByParams_Set[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Set(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqString_starGenType_Set_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqString_starGenType_Set_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqString_starGenType_Set_fnRecorder {
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
				DoFn       MoqString_starGenType_Set_doFn
				DoReturnFn MoqString_starGenType_Set_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqString_starGenType) PrettyParams_Set(params MoqString_starGenType_Set_params) string {
	return fmt.Sprintf("Set(%#v)", params.Value)
}

func (m *MoqString_starGenType) ParamsKey_Set(params MoqString_starGenType_Set_params, anyParams uint64) MoqString_starGenType_Set_paramsKey {
	m.Scene.T.Helper()
	var valueUsed string
	var valueUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Set.Value == moq.ParamIndexByValue {
			valueUsed = params.Value
		} else {
			valueUsedHash = hash.DeepHash(params.Value)
		}
	}
	return MoqString_starGenType_Set_paramsKey{
		Params: struct{ Value string }{
			Value: valueUsed,
		},
		Hashes: struct{ Value hash.Hash }{
			Value: valueUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqString_starGenType) Reset() {
	m.ResultsByParams_Value = nil
	m.ResultsByParams_String = nil
	m.ResultsByParams_Set = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqString_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Value {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Value(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_String {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_String(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Set {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Set(results.Params))
			}
		}
	}
}

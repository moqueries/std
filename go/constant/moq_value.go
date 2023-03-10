// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package constant

import (
	"fmt"
	"go/constant"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that Value_reduced is mocked completely
var _ Value_reduced = (*MoqValue_mock)(nil)

// Value_reduced is the fabricated implementation type of this mock (emitted
// when the original interface contains non-exported methods)
type Value_reduced interface {
	// Kind returns the value kind.
	Kind() constant.Kind

	// String returns a short, quoted (human-readable) form of the value.
	// For numeric values, the result may be an approximation;
	// for String values the result may be a shortened string.
	// Use ExactString for a string representing a value exactly.
	String() string

	// ExactString returns an exact, quoted (human-readable) form of the value.
	// If the Value is of Kind String, use StringVal to obtain the unquoted string.
	ExactString() string
}

// MoqValue holds the state of a moq of the Value_reduced type
type MoqValue struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqValue_mock

	ResultsByParams_Kind        []MoqValue_Kind_resultsByParams
	ResultsByParams_String      []MoqValue_String_resultsByParams
	ResultsByParams_ExactString []MoqValue_ExactString_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Kind        struct{}
			String      struct{}
			ExactString struct{}
		}
	}
}

// MoqValue_mock isolates the mock interface of the Value type
type MoqValue_mock struct {
	Moq *MoqValue
}

// MoqValue_recorder isolates the recorder interface of the Value type
type MoqValue_recorder struct {
	Moq *MoqValue
}

// MoqValue_Kind_params holds the params of the Value type
type MoqValue_Kind_params struct{}

// MoqValue_Kind_paramsKey holds the map key params of the Value type
type MoqValue_Kind_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqValue_Kind_resultsByParams contains the results for a given set of
// parameters for the Value type
type MoqValue_Kind_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqValue_Kind_paramsKey]*MoqValue_Kind_results
}

// MoqValue_Kind_doFn defines the type of function needed when calling AndDo
// for the Value type
type MoqValue_Kind_doFn func()

// MoqValue_Kind_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Value type
type MoqValue_Kind_doReturnFn func() constant.Kind

// MoqValue_Kind_results holds the results of the Value type
type MoqValue_Kind_results struct {
	Params  MoqValue_Kind_params
	Results []struct {
		Values *struct {
			Result1 constant.Kind
		}
		Sequence   uint32
		DoFn       MoqValue_Kind_doFn
		DoReturnFn MoqValue_Kind_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqValue_Kind_fnRecorder routes recorded function calls to the MoqValue moq
type MoqValue_Kind_fnRecorder struct {
	Params    MoqValue_Kind_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqValue_Kind_results
	Moq       *MoqValue
}

// MoqValue_Kind_anyParams isolates the any params functions of the Value type
type MoqValue_Kind_anyParams struct {
	Recorder *MoqValue_Kind_fnRecorder
}

// MoqValue_String_params holds the params of the Value type
type MoqValue_String_params struct{}

// MoqValue_String_paramsKey holds the map key params of the Value type
type MoqValue_String_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqValue_String_resultsByParams contains the results for a given set of
// parameters for the Value type
type MoqValue_String_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqValue_String_paramsKey]*MoqValue_String_results
}

// MoqValue_String_doFn defines the type of function needed when calling AndDo
// for the Value type
type MoqValue_String_doFn func()

// MoqValue_String_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Value type
type MoqValue_String_doReturnFn func() string

// MoqValue_String_results holds the results of the Value type
type MoqValue_String_results struct {
	Params  MoqValue_String_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqValue_String_doFn
		DoReturnFn MoqValue_String_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqValue_String_fnRecorder routes recorded function calls to the MoqValue
// moq
type MoqValue_String_fnRecorder struct {
	Params    MoqValue_String_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqValue_String_results
	Moq       *MoqValue
}

// MoqValue_String_anyParams isolates the any params functions of the Value
// type
type MoqValue_String_anyParams struct {
	Recorder *MoqValue_String_fnRecorder
}

// MoqValue_ExactString_params holds the params of the Value type
type MoqValue_ExactString_params struct{}

// MoqValue_ExactString_paramsKey holds the map key params of the Value type
type MoqValue_ExactString_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqValue_ExactString_resultsByParams contains the results for a given set of
// parameters for the Value type
type MoqValue_ExactString_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqValue_ExactString_paramsKey]*MoqValue_ExactString_results
}

// MoqValue_ExactString_doFn defines the type of function needed when calling
// AndDo for the Value type
type MoqValue_ExactString_doFn func()

// MoqValue_ExactString_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Value type
type MoqValue_ExactString_doReturnFn func() string

// MoqValue_ExactString_results holds the results of the Value type
type MoqValue_ExactString_results struct {
	Params  MoqValue_ExactString_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqValue_ExactString_doFn
		DoReturnFn MoqValue_ExactString_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqValue_ExactString_fnRecorder routes recorded function calls to the
// MoqValue moq
type MoqValue_ExactString_fnRecorder struct {
	Params    MoqValue_ExactString_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqValue_ExactString_results
	Moq       *MoqValue
}

// MoqValue_ExactString_anyParams isolates the any params functions of the
// Value type
type MoqValue_ExactString_anyParams struct {
	Recorder *MoqValue_ExactString_fnRecorder
}

// NewMoqValue creates a new moq of the Value type
func NewMoqValue(scene *moq.Scene, config *moq.Config) *MoqValue {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqValue{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqValue_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Kind        struct{}
				String      struct{}
				ExactString struct{}
			}
		}{ParameterIndexing: struct {
			Kind        struct{}
			String      struct{}
			ExactString struct{}
		}{
			Kind:        struct{}{},
			String:      struct{}{},
			ExactString: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Value type
func (m *MoqValue) Mock() *MoqValue_mock { return m.Moq }

func (m *MoqValue_mock) Kind() (result1 constant.Kind) {
	m.Moq.Scene.T.Helper()
	params := MoqValue_Kind_params{}
	var results *MoqValue_Kind_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Kind {
		paramsKey := m.Moq.ParamsKey_Kind(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Kind(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Kind(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Kind(params))
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

func (m *MoqValue_mock) String() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqValue_String_params{}
	var results *MoqValue_String_results
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

func (m *MoqValue_mock) ExactString() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqValue_ExactString_params{}
	var results *MoqValue_ExactString_results
	for _, resultsByParams := range m.Moq.ResultsByParams_ExactString {
		paramsKey := m.Moq.ParamsKey_ExactString(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_ExactString(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_ExactString(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_ExactString(params))
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

// OnCall returns the recorder implementation of the Value type
func (m *MoqValue) OnCall() *MoqValue_recorder {
	return &MoqValue_recorder{
		Moq: m,
	}
}

func (m *MoqValue_recorder) Kind() *MoqValue_Kind_fnRecorder {
	return &MoqValue_Kind_fnRecorder{
		Params:   MoqValue_Kind_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqValue_Kind_fnRecorder) Any() *MoqValue_Kind_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Kind(r.Params))
		return nil
	}
	return &MoqValue_Kind_anyParams{Recorder: r}
}

func (r *MoqValue_Kind_fnRecorder) Seq() *MoqValue_Kind_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Kind(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqValue_Kind_fnRecorder) NoSeq() *MoqValue_Kind_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Kind(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqValue_Kind_fnRecorder) ReturnResults(result1 constant.Kind) *MoqValue_Kind_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 constant.Kind
		}
		Sequence   uint32
		DoFn       MoqValue_Kind_doFn
		DoReturnFn MoqValue_Kind_doReturnFn
	}{
		Values: &struct {
			Result1 constant.Kind
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqValue_Kind_fnRecorder) AndDo(fn MoqValue_Kind_doFn) *MoqValue_Kind_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqValue_Kind_fnRecorder) DoReturnResults(fn MoqValue_Kind_doReturnFn) *MoqValue_Kind_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 constant.Kind
		}
		Sequence   uint32
		DoFn       MoqValue_Kind_doFn
		DoReturnFn MoqValue_Kind_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqValue_Kind_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqValue_Kind_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Kind {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqValue_Kind_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqValue_Kind_paramsKey]*MoqValue_Kind_results{},
		}
		r.Moq.ResultsByParams_Kind = append(r.Moq.ResultsByParams_Kind, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Kind) {
			copy(r.Moq.ResultsByParams_Kind[insertAt+1:], r.Moq.ResultsByParams_Kind[insertAt:0])
			r.Moq.ResultsByParams_Kind[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Kind(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqValue_Kind_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqValue_Kind_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqValue_Kind_fnRecorder {
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
					Result1 constant.Kind
				}
				Sequence   uint32
				DoFn       MoqValue_Kind_doFn
				DoReturnFn MoqValue_Kind_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqValue) PrettyParams_Kind(params MoqValue_Kind_params) string {
	return fmt.Sprintf("Kind()")
}

func (m *MoqValue) ParamsKey_Kind(params MoqValue_Kind_params, anyParams uint64) MoqValue_Kind_paramsKey {
	m.Scene.T.Helper()
	return MoqValue_Kind_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqValue_recorder) String() *MoqValue_String_fnRecorder {
	return &MoqValue_String_fnRecorder{
		Params:   MoqValue_String_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqValue_String_fnRecorder) Any() *MoqValue_String_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	return &MoqValue_String_anyParams{Recorder: r}
}

func (r *MoqValue_String_fnRecorder) Seq() *MoqValue_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqValue_String_fnRecorder) NoSeq() *MoqValue_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqValue_String_fnRecorder) ReturnResults(result1 string) *MoqValue_String_fnRecorder {
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
		DoFn       MoqValue_String_doFn
		DoReturnFn MoqValue_String_doReturnFn
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

func (r *MoqValue_String_fnRecorder) AndDo(fn MoqValue_String_doFn) *MoqValue_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqValue_String_fnRecorder) DoReturnResults(fn MoqValue_String_doReturnFn) *MoqValue_String_fnRecorder {
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
		DoFn       MoqValue_String_doFn
		DoReturnFn MoqValue_String_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqValue_String_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqValue_String_resultsByParams
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
		results = &MoqValue_String_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqValue_String_paramsKey]*MoqValue_String_results{},
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
		r.Results = &MoqValue_String_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqValue_String_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqValue_String_fnRecorder {
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
				DoFn       MoqValue_String_doFn
				DoReturnFn MoqValue_String_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqValue) PrettyParams_String(params MoqValue_String_params) string {
	return fmt.Sprintf("String()")
}

func (m *MoqValue) ParamsKey_String(params MoqValue_String_params, anyParams uint64) MoqValue_String_paramsKey {
	m.Scene.T.Helper()
	return MoqValue_String_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqValue_recorder) ExactString() *MoqValue_ExactString_fnRecorder {
	return &MoqValue_ExactString_fnRecorder{
		Params:   MoqValue_ExactString_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqValue_ExactString_fnRecorder) Any() *MoqValue_ExactString_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ExactString(r.Params))
		return nil
	}
	return &MoqValue_ExactString_anyParams{Recorder: r}
}

func (r *MoqValue_ExactString_fnRecorder) Seq() *MoqValue_ExactString_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ExactString(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqValue_ExactString_fnRecorder) NoSeq() *MoqValue_ExactString_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ExactString(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqValue_ExactString_fnRecorder) ReturnResults(result1 string) *MoqValue_ExactString_fnRecorder {
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
		DoFn       MoqValue_ExactString_doFn
		DoReturnFn MoqValue_ExactString_doReturnFn
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

func (r *MoqValue_ExactString_fnRecorder) AndDo(fn MoqValue_ExactString_doFn) *MoqValue_ExactString_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqValue_ExactString_fnRecorder) DoReturnResults(fn MoqValue_ExactString_doReturnFn) *MoqValue_ExactString_fnRecorder {
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
		DoFn       MoqValue_ExactString_doFn
		DoReturnFn MoqValue_ExactString_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqValue_ExactString_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqValue_ExactString_resultsByParams
	for n, res := range r.Moq.ResultsByParams_ExactString {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqValue_ExactString_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqValue_ExactString_paramsKey]*MoqValue_ExactString_results{},
		}
		r.Moq.ResultsByParams_ExactString = append(r.Moq.ResultsByParams_ExactString, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_ExactString) {
			copy(r.Moq.ResultsByParams_ExactString[insertAt+1:], r.Moq.ResultsByParams_ExactString[insertAt:0])
			r.Moq.ResultsByParams_ExactString[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_ExactString(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqValue_ExactString_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqValue_ExactString_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqValue_ExactString_fnRecorder {
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
				DoFn       MoqValue_ExactString_doFn
				DoReturnFn MoqValue_ExactString_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqValue) PrettyParams_ExactString(params MoqValue_ExactString_params) string {
	return fmt.Sprintf("ExactString()")
}

func (m *MoqValue) ParamsKey_ExactString(params MoqValue_ExactString_params, anyParams uint64) MoqValue_ExactString_paramsKey {
	m.Scene.T.Helper()
	return MoqValue_ExactString_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqValue) Reset() {
	m.ResultsByParams_Kind = nil
	m.ResultsByParams_String = nil
	m.ResultsByParams_ExactString = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqValue) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Kind {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Kind(results.Params))
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
	for _, res := range m.ResultsByParams_ExactString {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_ExactString(results.Params))
			}
		}
	}
}

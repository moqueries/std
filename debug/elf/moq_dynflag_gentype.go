// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package elf

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that elf.DynFlag_genType is mocked
// completely
var _ DynFlag_genType = (*MoqDynFlag_genType_mock)(nil)

// DynFlag_genType is the fabricated implementation type of this mock (emitted
// when mocking a collections of methods directly and not from an interface
// type)
type DynFlag_genType interface {
	String() string
	GoString() string
}

// MoqDynFlag_genType holds the state of a moq of the DynFlag_genType type
type MoqDynFlag_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqDynFlag_genType_mock

	ResultsByParams_String   []MoqDynFlag_genType_String_resultsByParams
	ResultsByParams_GoString []MoqDynFlag_genType_GoString_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			String   struct{}
			GoString struct{}
		}
	}
}

// MoqDynFlag_genType_mock isolates the mock interface of the DynFlag_genType
// type
type MoqDynFlag_genType_mock struct {
	Moq *MoqDynFlag_genType
}

// MoqDynFlag_genType_recorder isolates the recorder interface of the
// DynFlag_genType type
type MoqDynFlag_genType_recorder struct {
	Moq *MoqDynFlag_genType
}

// MoqDynFlag_genType_String_params holds the params of the DynFlag_genType
// type
type MoqDynFlag_genType_String_params struct{}

// MoqDynFlag_genType_String_paramsKey holds the map key params of the
// DynFlag_genType type
type MoqDynFlag_genType_String_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqDynFlag_genType_String_resultsByParams contains the results for a given
// set of parameters for the DynFlag_genType type
type MoqDynFlag_genType_String_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqDynFlag_genType_String_paramsKey]*MoqDynFlag_genType_String_results
}

// MoqDynFlag_genType_String_doFn defines the type of function needed when
// calling AndDo for the DynFlag_genType type
type MoqDynFlag_genType_String_doFn func()

// MoqDynFlag_genType_String_doReturnFn defines the type of function needed
// when calling DoReturnResults for the DynFlag_genType type
type MoqDynFlag_genType_String_doReturnFn func() string

// MoqDynFlag_genType_String_results holds the results of the DynFlag_genType
// type
type MoqDynFlag_genType_String_results struct {
	Params  MoqDynFlag_genType_String_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqDynFlag_genType_String_doFn
		DoReturnFn MoqDynFlag_genType_String_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqDynFlag_genType_String_fnRecorder routes recorded function calls to the
// MoqDynFlag_genType moq
type MoqDynFlag_genType_String_fnRecorder struct {
	Params    MoqDynFlag_genType_String_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqDynFlag_genType_String_results
	Moq       *MoqDynFlag_genType
}

// MoqDynFlag_genType_String_anyParams isolates the any params functions of the
// DynFlag_genType type
type MoqDynFlag_genType_String_anyParams struct {
	Recorder *MoqDynFlag_genType_String_fnRecorder
}

// MoqDynFlag_genType_GoString_params holds the params of the DynFlag_genType
// type
type MoqDynFlag_genType_GoString_params struct{}

// MoqDynFlag_genType_GoString_paramsKey holds the map key params of the
// DynFlag_genType type
type MoqDynFlag_genType_GoString_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqDynFlag_genType_GoString_resultsByParams contains the results for a given
// set of parameters for the DynFlag_genType type
type MoqDynFlag_genType_GoString_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqDynFlag_genType_GoString_paramsKey]*MoqDynFlag_genType_GoString_results
}

// MoqDynFlag_genType_GoString_doFn defines the type of function needed when
// calling AndDo for the DynFlag_genType type
type MoqDynFlag_genType_GoString_doFn func()

// MoqDynFlag_genType_GoString_doReturnFn defines the type of function needed
// when calling DoReturnResults for the DynFlag_genType type
type MoqDynFlag_genType_GoString_doReturnFn func() string

// MoqDynFlag_genType_GoString_results holds the results of the DynFlag_genType
// type
type MoqDynFlag_genType_GoString_results struct {
	Params  MoqDynFlag_genType_GoString_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqDynFlag_genType_GoString_doFn
		DoReturnFn MoqDynFlag_genType_GoString_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqDynFlag_genType_GoString_fnRecorder routes recorded function calls to the
// MoqDynFlag_genType moq
type MoqDynFlag_genType_GoString_fnRecorder struct {
	Params    MoqDynFlag_genType_GoString_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqDynFlag_genType_GoString_results
	Moq       *MoqDynFlag_genType
}

// MoqDynFlag_genType_GoString_anyParams isolates the any params functions of
// the DynFlag_genType type
type MoqDynFlag_genType_GoString_anyParams struct {
	Recorder *MoqDynFlag_genType_GoString_fnRecorder
}

// NewMoqDynFlag_genType creates a new moq of the DynFlag_genType type
func NewMoqDynFlag_genType(scene *moq.Scene, config *moq.Config) *MoqDynFlag_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqDynFlag_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqDynFlag_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				String   struct{}
				GoString struct{}
			}
		}{ParameterIndexing: struct {
			String   struct{}
			GoString struct{}
		}{
			String:   struct{}{},
			GoString: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the DynFlag_genType type
func (m *MoqDynFlag_genType) Mock() *MoqDynFlag_genType_mock { return m.Moq }

func (m *MoqDynFlag_genType_mock) String() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqDynFlag_genType_String_params{}
	var results *MoqDynFlag_genType_String_results
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

func (m *MoqDynFlag_genType_mock) GoString() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqDynFlag_genType_GoString_params{}
	var results *MoqDynFlag_genType_GoString_results
	for _, resultsByParams := range m.Moq.ResultsByParams_GoString {
		paramsKey := m.Moq.ParamsKey_GoString(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_GoString(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_GoString(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_GoString(params))
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

// OnCall returns the recorder implementation of the DynFlag_genType type
func (m *MoqDynFlag_genType) OnCall() *MoqDynFlag_genType_recorder {
	return &MoqDynFlag_genType_recorder{
		Moq: m,
	}
}

func (m *MoqDynFlag_genType_recorder) String() *MoqDynFlag_genType_String_fnRecorder {
	return &MoqDynFlag_genType_String_fnRecorder{
		Params:   MoqDynFlag_genType_String_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqDynFlag_genType_String_fnRecorder) Any() *MoqDynFlag_genType_String_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	return &MoqDynFlag_genType_String_anyParams{Recorder: r}
}

func (r *MoqDynFlag_genType_String_fnRecorder) Seq() *MoqDynFlag_genType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqDynFlag_genType_String_fnRecorder) NoSeq() *MoqDynFlag_genType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqDynFlag_genType_String_fnRecorder) ReturnResults(result1 string) *MoqDynFlag_genType_String_fnRecorder {
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
		DoFn       MoqDynFlag_genType_String_doFn
		DoReturnFn MoqDynFlag_genType_String_doReturnFn
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

func (r *MoqDynFlag_genType_String_fnRecorder) AndDo(fn MoqDynFlag_genType_String_doFn) *MoqDynFlag_genType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqDynFlag_genType_String_fnRecorder) DoReturnResults(fn MoqDynFlag_genType_String_doReturnFn) *MoqDynFlag_genType_String_fnRecorder {
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
		DoFn       MoqDynFlag_genType_String_doFn
		DoReturnFn MoqDynFlag_genType_String_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqDynFlag_genType_String_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqDynFlag_genType_String_resultsByParams
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
		results = &MoqDynFlag_genType_String_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqDynFlag_genType_String_paramsKey]*MoqDynFlag_genType_String_results{},
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
		r.Results = &MoqDynFlag_genType_String_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqDynFlag_genType_String_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqDynFlag_genType_String_fnRecorder {
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
				DoFn       MoqDynFlag_genType_String_doFn
				DoReturnFn MoqDynFlag_genType_String_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqDynFlag_genType) PrettyParams_String(params MoqDynFlag_genType_String_params) string {
	return fmt.Sprintf("String()")
}

func (m *MoqDynFlag_genType) ParamsKey_String(params MoqDynFlag_genType_String_params, anyParams uint64) MoqDynFlag_genType_String_paramsKey {
	m.Scene.T.Helper()
	return MoqDynFlag_genType_String_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqDynFlag_genType_recorder) GoString() *MoqDynFlag_genType_GoString_fnRecorder {
	return &MoqDynFlag_genType_GoString_fnRecorder{
		Params:   MoqDynFlag_genType_GoString_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqDynFlag_genType_GoString_fnRecorder) Any() *MoqDynFlag_genType_GoString_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_GoString(r.Params))
		return nil
	}
	return &MoqDynFlag_genType_GoString_anyParams{Recorder: r}
}

func (r *MoqDynFlag_genType_GoString_fnRecorder) Seq() *MoqDynFlag_genType_GoString_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_GoString(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqDynFlag_genType_GoString_fnRecorder) NoSeq() *MoqDynFlag_genType_GoString_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_GoString(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqDynFlag_genType_GoString_fnRecorder) ReturnResults(result1 string) *MoqDynFlag_genType_GoString_fnRecorder {
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
		DoFn       MoqDynFlag_genType_GoString_doFn
		DoReturnFn MoqDynFlag_genType_GoString_doReturnFn
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

func (r *MoqDynFlag_genType_GoString_fnRecorder) AndDo(fn MoqDynFlag_genType_GoString_doFn) *MoqDynFlag_genType_GoString_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqDynFlag_genType_GoString_fnRecorder) DoReturnResults(fn MoqDynFlag_genType_GoString_doReturnFn) *MoqDynFlag_genType_GoString_fnRecorder {
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
		DoFn       MoqDynFlag_genType_GoString_doFn
		DoReturnFn MoqDynFlag_genType_GoString_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqDynFlag_genType_GoString_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqDynFlag_genType_GoString_resultsByParams
	for n, res := range r.Moq.ResultsByParams_GoString {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqDynFlag_genType_GoString_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqDynFlag_genType_GoString_paramsKey]*MoqDynFlag_genType_GoString_results{},
		}
		r.Moq.ResultsByParams_GoString = append(r.Moq.ResultsByParams_GoString, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_GoString) {
			copy(r.Moq.ResultsByParams_GoString[insertAt+1:], r.Moq.ResultsByParams_GoString[insertAt:0])
			r.Moq.ResultsByParams_GoString[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_GoString(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqDynFlag_genType_GoString_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqDynFlag_genType_GoString_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqDynFlag_genType_GoString_fnRecorder {
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
				DoFn       MoqDynFlag_genType_GoString_doFn
				DoReturnFn MoqDynFlag_genType_GoString_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqDynFlag_genType) PrettyParams_GoString(params MoqDynFlag_genType_GoString_params) string {
	return fmt.Sprintf("GoString()")
}

func (m *MoqDynFlag_genType) ParamsKey_GoString(params MoqDynFlag_genType_GoString_params, anyParams uint64) MoqDynFlag_genType_GoString_paramsKey {
	m.Scene.T.Helper()
	return MoqDynFlag_genType_GoString_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqDynFlag_genType) Reset() {
	m.ResultsByParams_String = nil
	m.ResultsByParams_GoString = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqDynFlag_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_String {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_String(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_GoString {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_GoString(results.Params))
			}
		}
	}
}

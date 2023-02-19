// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package dwarf

import (
	"debug/dwarf"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that dwarf.BasicType_starGenType is
// mocked completely
var _ BasicType_starGenType = (*MoqBasicType_starGenType_mock)(nil)

// BasicType_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type BasicType_starGenType interface {
	Basic() *dwarf.BasicType
	String() string
}

// MoqBasicType_starGenType holds the state of a moq of the
// BasicType_starGenType type
type MoqBasicType_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqBasicType_starGenType_mock

	ResultsByParams_Basic  []MoqBasicType_starGenType_Basic_resultsByParams
	ResultsByParams_String []MoqBasicType_starGenType_String_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Basic  struct{}
			String struct{}
		}
	}
}

// MoqBasicType_starGenType_mock isolates the mock interface of the
// BasicType_starGenType type
type MoqBasicType_starGenType_mock struct {
	Moq *MoqBasicType_starGenType
}

// MoqBasicType_starGenType_recorder isolates the recorder interface of the
// BasicType_starGenType type
type MoqBasicType_starGenType_recorder struct {
	Moq *MoqBasicType_starGenType
}

// MoqBasicType_starGenType_Basic_params holds the params of the
// BasicType_starGenType type
type MoqBasicType_starGenType_Basic_params struct{}

// MoqBasicType_starGenType_Basic_paramsKey holds the map key params of the
// BasicType_starGenType type
type MoqBasicType_starGenType_Basic_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqBasicType_starGenType_Basic_resultsByParams contains the results for a
// given set of parameters for the BasicType_starGenType type
type MoqBasicType_starGenType_Basic_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqBasicType_starGenType_Basic_paramsKey]*MoqBasicType_starGenType_Basic_results
}

// MoqBasicType_starGenType_Basic_doFn defines the type of function needed when
// calling AndDo for the BasicType_starGenType type
type MoqBasicType_starGenType_Basic_doFn func()

// MoqBasicType_starGenType_Basic_doReturnFn defines the type of function
// needed when calling DoReturnResults for the BasicType_starGenType type
type MoqBasicType_starGenType_Basic_doReturnFn func() *dwarf.BasicType

// MoqBasicType_starGenType_Basic_results holds the results of the
// BasicType_starGenType type
type MoqBasicType_starGenType_Basic_results struct {
	Params  MoqBasicType_starGenType_Basic_params
	Results []struct {
		Values *struct {
			Result1 *dwarf.BasicType
		}
		Sequence   uint32
		DoFn       MoqBasicType_starGenType_Basic_doFn
		DoReturnFn MoqBasicType_starGenType_Basic_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqBasicType_starGenType_Basic_fnRecorder routes recorded function calls to
// the MoqBasicType_starGenType moq
type MoqBasicType_starGenType_Basic_fnRecorder struct {
	Params    MoqBasicType_starGenType_Basic_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqBasicType_starGenType_Basic_results
	Moq       *MoqBasicType_starGenType
}

// MoqBasicType_starGenType_Basic_anyParams isolates the any params functions
// of the BasicType_starGenType type
type MoqBasicType_starGenType_Basic_anyParams struct {
	Recorder *MoqBasicType_starGenType_Basic_fnRecorder
}

// MoqBasicType_starGenType_String_params holds the params of the
// BasicType_starGenType type
type MoqBasicType_starGenType_String_params struct{}

// MoqBasicType_starGenType_String_paramsKey holds the map key params of the
// BasicType_starGenType type
type MoqBasicType_starGenType_String_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqBasicType_starGenType_String_resultsByParams contains the results for a
// given set of parameters for the BasicType_starGenType type
type MoqBasicType_starGenType_String_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqBasicType_starGenType_String_paramsKey]*MoqBasicType_starGenType_String_results
}

// MoqBasicType_starGenType_String_doFn defines the type of function needed
// when calling AndDo for the BasicType_starGenType type
type MoqBasicType_starGenType_String_doFn func()

// MoqBasicType_starGenType_String_doReturnFn defines the type of function
// needed when calling DoReturnResults for the BasicType_starGenType type
type MoqBasicType_starGenType_String_doReturnFn func() string

// MoqBasicType_starGenType_String_results holds the results of the
// BasicType_starGenType type
type MoqBasicType_starGenType_String_results struct {
	Params  MoqBasicType_starGenType_String_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqBasicType_starGenType_String_doFn
		DoReturnFn MoqBasicType_starGenType_String_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqBasicType_starGenType_String_fnRecorder routes recorded function calls to
// the MoqBasicType_starGenType moq
type MoqBasicType_starGenType_String_fnRecorder struct {
	Params    MoqBasicType_starGenType_String_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqBasicType_starGenType_String_results
	Moq       *MoqBasicType_starGenType
}

// MoqBasicType_starGenType_String_anyParams isolates the any params functions
// of the BasicType_starGenType type
type MoqBasicType_starGenType_String_anyParams struct {
	Recorder *MoqBasicType_starGenType_String_fnRecorder
}

// NewMoqBasicType_starGenType creates a new moq of the BasicType_starGenType
// type
func NewMoqBasicType_starGenType(scene *moq.Scene, config *moq.Config) *MoqBasicType_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqBasicType_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqBasicType_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Basic  struct{}
				String struct{}
			}
		}{ParameterIndexing: struct {
			Basic  struct{}
			String struct{}
		}{
			Basic:  struct{}{},
			String: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the BasicType_starGenType type
func (m *MoqBasicType_starGenType) Mock() *MoqBasicType_starGenType_mock { return m.Moq }

func (m *MoqBasicType_starGenType_mock) Basic() (result1 *dwarf.BasicType) {
	m.Moq.Scene.T.Helper()
	params := MoqBasicType_starGenType_Basic_params{}
	var results *MoqBasicType_starGenType_Basic_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Basic {
		paramsKey := m.Moq.ParamsKey_Basic(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Basic(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Basic(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Basic(params))
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

func (m *MoqBasicType_starGenType_mock) String() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqBasicType_starGenType_String_params{}
	var results *MoqBasicType_starGenType_String_results
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

// OnCall returns the recorder implementation of the BasicType_starGenType type
func (m *MoqBasicType_starGenType) OnCall() *MoqBasicType_starGenType_recorder {
	return &MoqBasicType_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqBasicType_starGenType_recorder) Basic() *MoqBasicType_starGenType_Basic_fnRecorder {
	return &MoqBasicType_starGenType_Basic_fnRecorder{
		Params:   MoqBasicType_starGenType_Basic_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqBasicType_starGenType_Basic_fnRecorder) Any() *MoqBasicType_starGenType_Basic_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Basic(r.Params))
		return nil
	}
	return &MoqBasicType_starGenType_Basic_anyParams{Recorder: r}
}

func (r *MoqBasicType_starGenType_Basic_fnRecorder) Seq() *MoqBasicType_starGenType_Basic_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Basic(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqBasicType_starGenType_Basic_fnRecorder) NoSeq() *MoqBasicType_starGenType_Basic_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Basic(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqBasicType_starGenType_Basic_fnRecorder) ReturnResults(result1 *dwarf.BasicType) *MoqBasicType_starGenType_Basic_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *dwarf.BasicType
		}
		Sequence   uint32
		DoFn       MoqBasicType_starGenType_Basic_doFn
		DoReturnFn MoqBasicType_starGenType_Basic_doReturnFn
	}{
		Values: &struct {
			Result1 *dwarf.BasicType
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqBasicType_starGenType_Basic_fnRecorder) AndDo(fn MoqBasicType_starGenType_Basic_doFn) *MoqBasicType_starGenType_Basic_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqBasicType_starGenType_Basic_fnRecorder) DoReturnResults(fn MoqBasicType_starGenType_Basic_doReturnFn) *MoqBasicType_starGenType_Basic_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *dwarf.BasicType
		}
		Sequence   uint32
		DoFn       MoqBasicType_starGenType_Basic_doFn
		DoReturnFn MoqBasicType_starGenType_Basic_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqBasicType_starGenType_Basic_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqBasicType_starGenType_Basic_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Basic {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqBasicType_starGenType_Basic_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqBasicType_starGenType_Basic_paramsKey]*MoqBasicType_starGenType_Basic_results{},
		}
		r.Moq.ResultsByParams_Basic = append(r.Moq.ResultsByParams_Basic, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Basic) {
			copy(r.Moq.ResultsByParams_Basic[insertAt+1:], r.Moq.ResultsByParams_Basic[insertAt:0])
			r.Moq.ResultsByParams_Basic[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Basic(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqBasicType_starGenType_Basic_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqBasicType_starGenType_Basic_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqBasicType_starGenType_Basic_fnRecorder {
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
					Result1 *dwarf.BasicType
				}
				Sequence   uint32
				DoFn       MoqBasicType_starGenType_Basic_doFn
				DoReturnFn MoqBasicType_starGenType_Basic_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqBasicType_starGenType) PrettyParams_Basic(params MoqBasicType_starGenType_Basic_params) string {
	return fmt.Sprintf("Basic()")
}

func (m *MoqBasicType_starGenType) ParamsKey_Basic(params MoqBasicType_starGenType_Basic_params, anyParams uint64) MoqBasicType_starGenType_Basic_paramsKey {
	m.Scene.T.Helper()
	return MoqBasicType_starGenType_Basic_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqBasicType_starGenType_recorder) String() *MoqBasicType_starGenType_String_fnRecorder {
	return &MoqBasicType_starGenType_String_fnRecorder{
		Params:   MoqBasicType_starGenType_String_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqBasicType_starGenType_String_fnRecorder) Any() *MoqBasicType_starGenType_String_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	return &MoqBasicType_starGenType_String_anyParams{Recorder: r}
}

func (r *MoqBasicType_starGenType_String_fnRecorder) Seq() *MoqBasicType_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqBasicType_starGenType_String_fnRecorder) NoSeq() *MoqBasicType_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqBasicType_starGenType_String_fnRecorder) ReturnResults(result1 string) *MoqBasicType_starGenType_String_fnRecorder {
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
		DoFn       MoqBasicType_starGenType_String_doFn
		DoReturnFn MoqBasicType_starGenType_String_doReturnFn
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

func (r *MoqBasicType_starGenType_String_fnRecorder) AndDo(fn MoqBasicType_starGenType_String_doFn) *MoqBasicType_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqBasicType_starGenType_String_fnRecorder) DoReturnResults(fn MoqBasicType_starGenType_String_doReturnFn) *MoqBasicType_starGenType_String_fnRecorder {
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
		DoFn       MoqBasicType_starGenType_String_doFn
		DoReturnFn MoqBasicType_starGenType_String_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqBasicType_starGenType_String_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqBasicType_starGenType_String_resultsByParams
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
		results = &MoqBasicType_starGenType_String_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqBasicType_starGenType_String_paramsKey]*MoqBasicType_starGenType_String_results{},
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
		r.Results = &MoqBasicType_starGenType_String_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqBasicType_starGenType_String_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqBasicType_starGenType_String_fnRecorder {
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
				DoFn       MoqBasicType_starGenType_String_doFn
				DoReturnFn MoqBasicType_starGenType_String_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqBasicType_starGenType) PrettyParams_String(params MoqBasicType_starGenType_String_params) string {
	return fmt.Sprintf("String()")
}

func (m *MoqBasicType_starGenType) ParamsKey_String(params MoqBasicType_starGenType_String_params, anyParams uint64) MoqBasicType_starGenType_String_paramsKey {
	m.Scene.T.Helper()
	return MoqBasicType_starGenType_String_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqBasicType_starGenType) Reset() {
	m.ResultsByParams_Basic = nil
	m.ResultsByParams_String = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqBasicType_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Basic {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Basic(results.Params))
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
}
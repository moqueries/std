// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package elf

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that elf.OSABI_genType is mocked
// completely
var _ OSABI_genType = (*MoqOSABI_genType_mock)(nil)

// OSABI_genType is the fabricated implementation type of this mock (emitted
// when mocking a collections of methods directly and not from an interface
// type)
type OSABI_genType interface {
	String() string
	GoString() string
}

// MoqOSABI_genType holds the state of a moq of the OSABI_genType type
type MoqOSABI_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqOSABI_genType_mock

	ResultsByParams_String   []MoqOSABI_genType_String_resultsByParams
	ResultsByParams_GoString []MoqOSABI_genType_GoString_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			String   struct{}
			GoString struct{}
		}
	}
}

// MoqOSABI_genType_mock isolates the mock interface of the OSABI_genType type
type MoqOSABI_genType_mock struct {
	Moq *MoqOSABI_genType
}

// MoqOSABI_genType_recorder isolates the recorder interface of the
// OSABI_genType type
type MoqOSABI_genType_recorder struct {
	Moq *MoqOSABI_genType
}

// MoqOSABI_genType_String_params holds the params of the OSABI_genType type
type MoqOSABI_genType_String_params struct{}

// MoqOSABI_genType_String_paramsKey holds the map key params of the
// OSABI_genType type
type MoqOSABI_genType_String_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqOSABI_genType_String_resultsByParams contains the results for a given set
// of parameters for the OSABI_genType type
type MoqOSABI_genType_String_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqOSABI_genType_String_paramsKey]*MoqOSABI_genType_String_results
}

// MoqOSABI_genType_String_doFn defines the type of function needed when
// calling AndDo for the OSABI_genType type
type MoqOSABI_genType_String_doFn func()

// MoqOSABI_genType_String_doReturnFn defines the type of function needed when
// calling DoReturnResults for the OSABI_genType type
type MoqOSABI_genType_String_doReturnFn func() string

// MoqOSABI_genType_String_results holds the results of the OSABI_genType type
type MoqOSABI_genType_String_results struct {
	Params  MoqOSABI_genType_String_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqOSABI_genType_String_doFn
		DoReturnFn MoqOSABI_genType_String_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqOSABI_genType_String_fnRecorder routes recorded function calls to the
// MoqOSABI_genType moq
type MoqOSABI_genType_String_fnRecorder struct {
	Params    MoqOSABI_genType_String_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqOSABI_genType_String_results
	Moq       *MoqOSABI_genType
}

// MoqOSABI_genType_String_anyParams isolates the any params functions of the
// OSABI_genType type
type MoqOSABI_genType_String_anyParams struct {
	Recorder *MoqOSABI_genType_String_fnRecorder
}

// MoqOSABI_genType_GoString_params holds the params of the OSABI_genType type
type MoqOSABI_genType_GoString_params struct{}

// MoqOSABI_genType_GoString_paramsKey holds the map key params of the
// OSABI_genType type
type MoqOSABI_genType_GoString_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqOSABI_genType_GoString_resultsByParams contains the results for a given
// set of parameters for the OSABI_genType type
type MoqOSABI_genType_GoString_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqOSABI_genType_GoString_paramsKey]*MoqOSABI_genType_GoString_results
}

// MoqOSABI_genType_GoString_doFn defines the type of function needed when
// calling AndDo for the OSABI_genType type
type MoqOSABI_genType_GoString_doFn func()

// MoqOSABI_genType_GoString_doReturnFn defines the type of function needed
// when calling DoReturnResults for the OSABI_genType type
type MoqOSABI_genType_GoString_doReturnFn func() string

// MoqOSABI_genType_GoString_results holds the results of the OSABI_genType
// type
type MoqOSABI_genType_GoString_results struct {
	Params  MoqOSABI_genType_GoString_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqOSABI_genType_GoString_doFn
		DoReturnFn MoqOSABI_genType_GoString_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqOSABI_genType_GoString_fnRecorder routes recorded function calls to the
// MoqOSABI_genType moq
type MoqOSABI_genType_GoString_fnRecorder struct {
	Params    MoqOSABI_genType_GoString_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqOSABI_genType_GoString_results
	Moq       *MoqOSABI_genType
}

// MoqOSABI_genType_GoString_anyParams isolates the any params functions of the
// OSABI_genType type
type MoqOSABI_genType_GoString_anyParams struct {
	Recorder *MoqOSABI_genType_GoString_fnRecorder
}

// NewMoqOSABI_genType creates a new moq of the OSABI_genType type
func NewMoqOSABI_genType(scene *moq.Scene, config *moq.Config) *MoqOSABI_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqOSABI_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqOSABI_genType_mock{},

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

// Mock returns the mock implementation of the OSABI_genType type
func (m *MoqOSABI_genType) Mock() *MoqOSABI_genType_mock { return m.Moq }

func (m *MoqOSABI_genType_mock) String() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqOSABI_genType_String_params{}
	var results *MoqOSABI_genType_String_results
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

func (m *MoqOSABI_genType_mock) GoString() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqOSABI_genType_GoString_params{}
	var results *MoqOSABI_genType_GoString_results
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

// OnCall returns the recorder implementation of the OSABI_genType type
func (m *MoqOSABI_genType) OnCall() *MoqOSABI_genType_recorder {
	return &MoqOSABI_genType_recorder{
		Moq: m,
	}
}

func (m *MoqOSABI_genType_recorder) String() *MoqOSABI_genType_String_fnRecorder {
	return &MoqOSABI_genType_String_fnRecorder{
		Params:   MoqOSABI_genType_String_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqOSABI_genType_String_fnRecorder) Any() *MoqOSABI_genType_String_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	return &MoqOSABI_genType_String_anyParams{Recorder: r}
}

func (r *MoqOSABI_genType_String_fnRecorder) Seq() *MoqOSABI_genType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqOSABI_genType_String_fnRecorder) NoSeq() *MoqOSABI_genType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqOSABI_genType_String_fnRecorder) ReturnResults(result1 string) *MoqOSABI_genType_String_fnRecorder {
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
		DoFn       MoqOSABI_genType_String_doFn
		DoReturnFn MoqOSABI_genType_String_doReturnFn
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

func (r *MoqOSABI_genType_String_fnRecorder) AndDo(fn MoqOSABI_genType_String_doFn) *MoqOSABI_genType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqOSABI_genType_String_fnRecorder) DoReturnResults(fn MoqOSABI_genType_String_doReturnFn) *MoqOSABI_genType_String_fnRecorder {
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
		DoFn       MoqOSABI_genType_String_doFn
		DoReturnFn MoqOSABI_genType_String_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqOSABI_genType_String_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqOSABI_genType_String_resultsByParams
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
		results = &MoqOSABI_genType_String_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqOSABI_genType_String_paramsKey]*MoqOSABI_genType_String_results{},
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
		r.Results = &MoqOSABI_genType_String_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqOSABI_genType_String_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqOSABI_genType_String_fnRecorder {
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
				DoFn       MoqOSABI_genType_String_doFn
				DoReturnFn MoqOSABI_genType_String_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqOSABI_genType) PrettyParams_String(params MoqOSABI_genType_String_params) string {
	return fmt.Sprintf("String()")
}

func (m *MoqOSABI_genType) ParamsKey_String(params MoqOSABI_genType_String_params, anyParams uint64) MoqOSABI_genType_String_paramsKey {
	m.Scene.T.Helper()
	return MoqOSABI_genType_String_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqOSABI_genType_recorder) GoString() *MoqOSABI_genType_GoString_fnRecorder {
	return &MoqOSABI_genType_GoString_fnRecorder{
		Params:   MoqOSABI_genType_GoString_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqOSABI_genType_GoString_fnRecorder) Any() *MoqOSABI_genType_GoString_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_GoString(r.Params))
		return nil
	}
	return &MoqOSABI_genType_GoString_anyParams{Recorder: r}
}

func (r *MoqOSABI_genType_GoString_fnRecorder) Seq() *MoqOSABI_genType_GoString_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_GoString(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqOSABI_genType_GoString_fnRecorder) NoSeq() *MoqOSABI_genType_GoString_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_GoString(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqOSABI_genType_GoString_fnRecorder) ReturnResults(result1 string) *MoqOSABI_genType_GoString_fnRecorder {
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
		DoFn       MoqOSABI_genType_GoString_doFn
		DoReturnFn MoqOSABI_genType_GoString_doReturnFn
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

func (r *MoqOSABI_genType_GoString_fnRecorder) AndDo(fn MoqOSABI_genType_GoString_doFn) *MoqOSABI_genType_GoString_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqOSABI_genType_GoString_fnRecorder) DoReturnResults(fn MoqOSABI_genType_GoString_doReturnFn) *MoqOSABI_genType_GoString_fnRecorder {
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
		DoFn       MoqOSABI_genType_GoString_doFn
		DoReturnFn MoqOSABI_genType_GoString_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqOSABI_genType_GoString_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqOSABI_genType_GoString_resultsByParams
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
		results = &MoqOSABI_genType_GoString_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqOSABI_genType_GoString_paramsKey]*MoqOSABI_genType_GoString_results{},
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
		r.Results = &MoqOSABI_genType_GoString_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqOSABI_genType_GoString_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqOSABI_genType_GoString_fnRecorder {
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
				DoFn       MoqOSABI_genType_GoString_doFn
				DoReturnFn MoqOSABI_genType_GoString_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqOSABI_genType) PrettyParams_GoString(params MoqOSABI_genType_GoString_params) string {
	return fmt.Sprintf("GoString()")
}

func (m *MoqOSABI_genType) ParamsKey_GoString(params MoqOSABI_genType_GoString_params, anyParams uint64) MoqOSABI_genType_GoString_paramsKey {
	m.Scene.T.Helper()
	return MoqOSABI_genType_GoString_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqOSABI_genType) Reset() { m.ResultsByParams_String = nil; m.ResultsByParams_GoString = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqOSABI_genType) AssertExpectationsMet() {
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
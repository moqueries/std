// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package net

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that net.IPMask_genType is mocked
// completely
var _ IPMask_genType = (*MoqIPMask_genType_mock)(nil)

// IPMask_genType is the fabricated implementation type of this mock (emitted
// when mocking a collections of methods directly and not from an interface
// type)
type IPMask_genType interface {
	Size() (ones, bits int)
	String() string
}

// MoqIPMask_genType holds the state of a moq of the IPMask_genType type
type MoqIPMask_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqIPMask_genType_mock

	ResultsByParams_Size   []MoqIPMask_genType_Size_resultsByParams
	ResultsByParams_String []MoqIPMask_genType_String_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Size   struct{}
			String struct{}
		}
	}
}

// MoqIPMask_genType_mock isolates the mock interface of the IPMask_genType
// type
type MoqIPMask_genType_mock struct {
	Moq *MoqIPMask_genType
}

// MoqIPMask_genType_recorder isolates the recorder interface of the
// IPMask_genType type
type MoqIPMask_genType_recorder struct {
	Moq *MoqIPMask_genType
}

// MoqIPMask_genType_Size_params holds the params of the IPMask_genType type
type MoqIPMask_genType_Size_params struct{}

// MoqIPMask_genType_Size_paramsKey holds the map key params of the
// IPMask_genType type
type MoqIPMask_genType_Size_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqIPMask_genType_Size_resultsByParams contains the results for a given set
// of parameters for the IPMask_genType type
type MoqIPMask_genType_Size_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqIPMask_genType_Size_paramsKey]*MoqIPMask_genType_Size_results
}

// MoqIPMask_genType_Size_doFn defines the type of function needed when calling
// AndDo for the IPMask_genType type
type MoqIPMask_genType_Size_doFn func()

// MoqIPMask_genType_Size_doReturnFn defines the type of function needed when
// calling DoReturnResults for the IPMask_genType type
type MoqIPMask_genType_Size_doReturnFn func() (ones, bits int)

// MoqIPMask_genType_Size_results holds the results of the IPMask_genType type
type MoqIPMask_genType_Size_results struct {
	Params  MoqIPMask_genType_Size_params
	Results []struct {
		Values     *struct{ Ones, Bits int }
		Sequence   uint32
		DoFn       MoqIPMask_genType_Size_doFn
		DoReturnFn MoqIPMask_genType_Size_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqIPMask_genType_Size_fnRecorder routes recorded function calls to the
// MoqIPMask_genType moq
type MoqIPMask_genType_Size_fnRecorder struct {
	Params    MoqIPMask_genType_Size_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqIPMask_genType_Size_results
	Moq       *MoqIPMask_genType
}

// MoqIPMask_genType_Size_anyParams isolates the any params functions of the
// IPMask_genType type
type MoqIPMask_genType_Size_anyParams struct {
	Recorder *MoqIPMask_genType_Size_fnRecorder
}

// MoqIPMask_genType_String_params holds the params of the IPMask_genType type
type MoqIPMask_genType_String_params struct{}

// MoqIPMask_genType_String_paramsKey holds the map key params of the
// IPMask_genType type
type MoqIPMask_genType_String_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqIPMask_genType_String_resultsByParams contains the results for a given
// set of parameters for the IPMask_genType type
type MoqIPMask_genType_String_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqIPMask_genType_String_paramsKey]*MoqIPMask_genType_String_results
}

// MoqIPMask_genType_String_doFn defines the type of function needed when
// calling AndDo for the IPMask_genType type
type MoqIPMask_genType_String_doFn func()

// MoqIPMask_genType_String_doReturnFn defines the type of function needed when
// calling DoReturnResults for the IPMask_genType type
type MoqIPMask_genType_String_doReturnFn func() string

// MoqIPMask_genType_String_results holds the results of the IPMask_genType
// type
type MoqIPMask_genType_String_results struct {
	Params  MoqIPMask_genType_String_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqIPMask_genType_String_doFn
		DoReturnFn MoqIPMask_genType_String_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqIPMask_genType_String_fnRecorder routes recorded function calls to the
// MoqIPMask_genType moq
type MoqIPMask_genType_String_fnRecorder struct {
	Params    MoqIPMask_genType_String_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqIPMask_genType_String_results
	Moq       *MoqIPMask_genType
}

// MoqIPMask_genType_String_anyParams isolates the any params functions of the
// IPMask_genType type
type MoqIPMask_genType_String_anyParams struct {
	Recorder *MoqIPMask_genType_String_fnRecorder
}

// NewMoqIPMask_genType creates a new moq of the IPMask_genType type
func NewMoqIPMask_genType(scene *moq.Scene, config *moq.Config) *MoqIPMask_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqIPMask_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqIPMask_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Size   struct{}
				String struct{}
			}
		}{ParameterIndexing: struct {
			Size   struct{}
			String struct{}
		}{
			Size:   struct{}{},
			String: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the IPMask_genType type
func (m *MoqIPMask_genType) Mock() *MoqIPMask_genType_mock { return m.Moq }

func (m *MoqIPMask_genType_mock) Size() (ones, bits int) {
	m.Moq.Scene.T.Helper()
	params := MoqIPMask_genType_Size_params{}
	var results *MoqIPMask_genType_Size_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Size {
		paramsKey := m.Moq.ParamsKey_Size(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Size(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Size(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Size(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		ones = result.Values.Ones
		bits = result.Values.Bits
	}
	if result.DoReturnFn != nil {
		ones, bits = result.DoReturnFn()
	}
	return
}

func (m *MoqIPMask_genType_mock) String() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqIPMask_genType_String_params{}
	var results *MoqIPMask_genType_String_results
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

// OnCall returns the recorder implementation of the IPMask_genType type
func (m *MoqIPMask_genType) OnCall() *MoqIPMask_genType_recorder {
	return &MoqIPMask_genType_recorder{
		Moq: m,
	}
}

func (m *MoqIPMask_genType_recorder) Size() *MoqIPMask_genType_Size_fnRecorder {
	return &MoqIPMask_genType_Size_fnRecorder{
		Params:   MoqIPMask_genType_Size_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqIPMask_genType_Size_fnRecorder) Any() *MoqIPMask_genType_Size_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Size(r.Params))
		return nil
	}
	return &MoqIPMask_genType_Size_anyParams{Recorder: r}
}

func (r *MoqIPMask_genType_Size_fnRecorder) Seq() *MoqIPMask_genType_Size_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Size(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqIPMask_genType_Size_fnRecorder) NoSeq() *MoqIPMask_genType_Size_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Size(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqIPMask_genType_Size_fnRecorder) ReturnResults(ones, bits int) *MoqIPMask_genType_Size_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Ones, Bits int }
		Sequence   uint32
		DoFn       MoqIPMask_genType_Size_doFn
		DoReturnFn MoqIPMask_genType_Size_doReturnFn
	}{
		Values: &struct{ Ones, Bits int }{
			Ones: ones,
			Bits: bits,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqIPMask_genType_Size_fnRecorder) AndDo(fn MoqIPMask_genType_Size_doFn) *MoqIPMask_genType_Size_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqIPMask_genType_Size_fnRecorder) DoReturnResults(fn MoqIPMask_genType_Size_doReturnFn) *MoqIPMask_genType_Size_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Ones, Bits int }
		Sequence   uint32
		DoFn       MoqIPMask_genType_Size_doFn
		DoReturnFn MoqIPMask_genType_Size_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqIPMask_genType_Size_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqIPMask_genType_Size_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Size {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqIPMask_genType_Size_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqIPMask_genType_Size_paramsKey]*MoqIPMask_genType_Size_results{},
		}
		r.Moq.ResultsByParams_Size = append(r.Moq.ResultsByParams_Size, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Size) {
			copy(r.Moq.ResultsByParams_Size[insertAt+1:], r.Moq.ResultsByParams_Size[insertAt:0])
			r.Moq.ResultsByParams_Size[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Size(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqIPMask_genType_Size_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqIPMask_genType_Size_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqIPMask_genType_Size_fnRecorder {
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
				Values     *struct{ Ones, Bits int }
				Sequence   uint32
				DoFn       MoqIPMask_genType_Size_doFn
				DoReturnFn MoqIPMask_genType_Size_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqIPMask_genType) PrettyParams_Size(params MoqIPMask_genType_Size_params) string {
	return fmt.Sprintf("Size()")
}

func (m *MoqIPMask_genType) ParamsKey_Size(params MoqIPMask_genType_Size_params, anyParams uint64) MoqIPMask_genType_Size_paramsKey {
	m.Scene.T.Helper()
	return MoqIPMask_genType_Size_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqIPMask_genType_recorder) String() *MoqIPMask_genType_String_fnRecorder {
	return &MoqIPMask_genType_String_fnRecorder{
		Params:   MoqIPMask_genType_String_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqIPMask_genType_String_fnRecorder) Any() *MoqIPMask_genType_String_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	return &MoqIPMask_genType_String_anyParams{Recorder: r}
}

func (r *MoqIPMask_genType_String_fnRecorder) Seq() *MoqIPMask_genType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqIPMask_genType_String_fnRecorder) NoSeq() *MoqIPMask_genType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqIPMask_genType_String_fnRecorder) ReturnResults(result1 string) *MoqIPMask_genType_String_fnRecorder {
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
		DoFn       MoqIPMask_genType_String_doFn
		DoReturnFn MoqIPMask_genType_String_doReturnFn
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

func (r *MoqIPMask_genType_String_fnRecorder) AndDo(fn MoqIPMask_genType_String_doFn) *MoqIPMask_genType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqIPMask_genType_String_fnRecorder) DoReturnResults(fn MoqIPMask_genType_String_doReturnFn) *MoqIPMask_genType_String_fnRecorder {
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
		DoFn       MoqIPMask_genType_String_doFn
		DoReturnFn MoqIPMask_genType_String_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqIPMask_genType_String_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqIPMask_genType_String_resultsByParams
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
		results = &MoqIPMask_genType_String_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqIPMask_genType_String_paramsKey]*MoqIPMask_genType_String_results{},
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
		r.Results = &MoqIPMask_genType_String_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqIPMask_genType_String_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqIPMask_genType_String_fnRecorder {
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
				DoFn       MoqIPMask_genType_String_doFn
				DoReturnFn MoqIPMask_genType_String_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqIPMask_genType) PrettyParams_String(params MoqIPMask_genType_String_params) string {
	return fmt.Sprintf("String()")
}

func (m *MoqIPMask_genType) ParamsKey_String(params MoqIPMask_genType_String_params, anyParams uint64) MoqIPMask_genType_String_paramsKey {
	m.Scene.T.Helper()
	return MoqIPMask_genType_String_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqIPMask_genType) Reset() { m.ResultsByParams_Size = nil; m.ResultsByParams_String = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqIPMask_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Size {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Size(results.Params))
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

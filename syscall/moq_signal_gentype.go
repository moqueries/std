// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that syscall.Signal_genType is mocked
// completely
var _ Signal_genType = (*MoqSignal_genType_mock)(nil)

// Signal_genType is the fabricated implementation type of this mock (emitted
// when mocking a collections of methods directly and not from an interface
// type)
type Signal_genType interface {
	Signal()
	String() string
}

// MoqSignal_genType holds the state of a moq of the Signal_genType type
type MoqSignal_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSignal_genType_mock

	ResultsByParams_Signal []MoqSignal_genType_Signal_resultsByParams
	ResultsByParams_String []MoqSignal_genType_String_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Signal struct{}
			String struct{}
		}
	}
}

// MoqSignal_genType_mock isolates the mock interface of the Signal_genType
// type
type MoqSignal_genType_mock struct {
	Moq *MoqSignal_genType
}

// MoqSignal_genType_recorder isolates the recorder interface of the
// Signal_genType type
type MoqSignal_genType_recorder struct {
	Moq *MoqSignal_genType
}

// MoqSignal_genType_Signal_params holds the params of the Signal_genType type
type MoqSignal_genType_Signal_params struct{}

// MoqSignal_genType_Signal_paramsKey holds the map key params of the
// Signal_genType type
type MoqSignal_genType_Signal_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqSignal_genType_Signal_resultsByParams contains the results for a given
// set of parameters for the Signal_genType type
type MoqSignal_genType_Signal_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSignal_genType_Signal_paramsKey]*MoqSignal_genType_Signal_results
}

// MoqSignal_genType_Signal_doFn defines the type of function needed when
// calling AndDo for the Signal_genType type
type MoqSignal_genType_Signal_doFn func()

// MoqSignal_genType_Signal_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Signal_genType type
type MoqSignal_genType_Signal_doReturnFn func()

// MoqSignal_genType_Signal_results holds the results of the Signal_genType
// type
type MoqSignal_genType_Signal_results struct {
	Params  MoqSignal_genType_Signal_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSignal_genType_Signal_doFn
		DoReturnFn MoqSignal_genType_Signal_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSignal_genType_Signal_fnRecorder routes recorded function calls to the
// MoqSignal_genType moq
type MoqSignal_genType_Signal_fnRecorder struct {
	Params    MoqSignal_genType_Signal_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSignal_genType_Signal_results
	Moq       *MoqSignal_genType
}

// MoqSignal_genType_Signal_anyParams isolates the any params functions of the
// Signal_genType type
type MoqSignal_genType_Signal_anyParams struct {
	Recorder *MoqSignal_genType_Signal_fnRecorder
}

// MoqSignal_genType_String_params holds the params of the Signal_genType type
type MoqSignal_genType_String_params struct{}

// MoqSignal_genType_String_paramsKey holds the map key params of the
// Signal_genType type
type MoqSignal_genType_String_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqSignal_genType_String_resultsByParams contains the results for a given
// set of parameters for the Signal_genType type
type MoqSignal_genType_String_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSignal_genType_String_paramsKey]*MoqSignal_genType_String_results
}

// MoqSignal_genType_String_doFn defines the type of function needed when
// calling AndDo for the Signal_genType type
type MoqSignal_genType_String_doFn func()

// MoqSignal_genType_String_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Signal_genType type
type MoqSignal_genType_String_doReturnFn func() string

// MoqSignal_genType_String_results holds the results of the Signal_genType
// type
type MoqSignal_genType_String_results struct {
	Params  MoqSignal_genType_String_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqSignal_genType_String_doFn
		DoReturnFn MoqSignal_genType_String_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSignal_genType_String_fnRecorder routes recorded function calls to the
// MoqSignal_genType moq
type MoqSignal_genType_String_fnRecorder struct {
	Params    MoqSignal_genType_String_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSignal_genType_String_results
	Moq       *MoqSignal_genType
}

// MoqSignal_genType_String_anyParams isolates the any params functions of the
// Signal_genType type
type MoqSignal_genType_String_anyParams struct {
	Recorder *MoqSignal_genType_String_fnRecorder
}

// NewMoqSignal_genType creates a new moq of the Signal_genType type
func NewMoqSignal_genType(scene *moq.Scene, config *moq.Config) *MoqSignal_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSignal_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSignal_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Signal struct{}
				String struct{}
			}
		}{ParameterIndexing: struct {
			Signal struct{}
			String struct{}
		}{
			Signal: struct{}{},
			String: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Signal_genType type
func (m *MoqSignal_genType) Mock() *MoqSignal_genType_mock { return m.Moq }

func (m *MoqSignal_genType_mock) Signal() {
	m.Moq.Scene.T.Helper()
	params := MoqSignal_genType_Signal_params{}
	var results *MoqSignal_genType_Signal_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Signal {
		paramsKey := m.Moq.ParamsKey_Signal(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Signal(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Signal(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Signal(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn()
	}
	return
}

func (m *MoqSignal_genType_mock) String() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqSignal_genType_String_params{}
	var results *MoqSignal_genType_String_results
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

// OnCall returns the recorder implementation of the Signal_genType type
func (m *MoqSignal_genType) OnCall() *MoqSignal_genType_recorder {
	return &MoqSignal_genType_recorder{
		Moq: m,
	}
}

func (m *MoqSignal_genType_recorder) Signal() *MoqSignal_genType_Signal_fnRecorder {
	return &MoqSignal_genType_Signal_fnRecorder{
		Params:   MoqSignal_genType_Signal_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqSignal_genType_Signal_fnRecorder) Any() *MoqSignal_genType_Signal_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Signal(r.Params))
		return nil
	}
	return &MoqSignal_genType_Signal_anyParams{Recorder: r}
}

func (r *MoqSignal_genType_Signal_fnRecorder) Seq() *MoqSignal_genType_Signal_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Signal(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSignal_genType_Signal_fnRecorder) NoSeq() *MoqSignal_genType_Signal_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Signal(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSignal_genType_Signal_fnRecorder) ReturnResults() *MoqSignal_genType_Signal_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSignal_genType_Signal_doFn
		DoReturnFn MoqSignal_genType_Signal_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSignal_genType_Signal_fnRecorder) AndDo(fn MoqSignal_genType_Signal_doFn) *MoqSignal_genType_Signal_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSignal_genType_Signal_fnRecorder) DoReturnResults(fn MoqSignal_genType_Signal_doReturnFn) *MoqSignal_genType_Signal_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSignal_genType_Signal_doFn
		DoReturnFn MoqSignal_genType_Signal_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSignal_genType_Signal_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSignal_genType_Signal_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Signal {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqSignal_genType_Signal_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSignal_genType_Signal_paramsKey]*MoqSignal_genType_Signal_results{},
		}
		r.Moq.ResultsByParams_Signal = append(r.Moq.ResultsByParams_Signal, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Signal) {
			copy(r.Moq.ResultsByParams_Signal[insertAt+1:], r.Moq.ResultsByParams_Signal[insertAt:0])
			r.Moq.ResultsByParams_Signal[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Signal(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqSignal_genType_Signal_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSignal_genType_Signal_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSignal_genType_Signal_fnRecorder {
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
				DoFn       MoqSignal_genType_Signal_doFn
				DoReturnFn MoqSignal_genType_Signal_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSignal_genType) PrettyParams_Signal(params MoqSignal_genType_Signal_params) string {
	return fmt.Sprintf("Signal()")
}

func (m *MoqSignal_genType) ParamsKey_Signal(params MoqSignal_genType_Signal_params, anyParams uint64) MoqSignal_genType_Signal_paramsKey {
	m.Scene.T.Helper()
	return MoqSignal_genType_Signal_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqSignal_genType_recorder) String() *MoqSignal_genType_String_fnRecorder {
	return &MoqSignal_genType_String_fnRecorder{
		Params:   MoqSignal_genType_String_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqSignal_genType_String_fnRecorder) Any() *MoqSignal_genType_String_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	return &MoqSignal_genType_String_anyParams{Recorder: r}
}

func (r *MoqSignal_genType_String_fnRecorder) Seq() *MoqSignal_genType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSignal_genType_String_fnRecorder) NoSeq() *MoqSignal_genType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSignal_genType_String_fnRecorder) ReturnResults(result1 string) *MoqSignal_genType_String_fnRecorder {
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
		DoFn       MoqSignal_genType_String_doFn
		DoReturnFn MoqSignal_genType_String_doReturnFn
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

func (r *MoqSignal_genType_String_fnRecorder) AndDo(fn MoqSignal_genType_String_doFn) *MoqSignal_genType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSignal_genType_String_fnRecorder) DoReturnResults(fn MoqSignal_genType_String_doReturnFn) *MoqSignal_genType_String_fnRecorder {
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
		DoFn       MoqSignal_genType_String_doFn
		DoReturnFn MoqSignal_genType_String_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSignal_genType_String_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSignal_genType_String_resultsByParams
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
		results = &MoqSignal_genType_String_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSignal_genType_String_paramsKey]*MoqSignal_genType_String_results{},
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
		r.Results = &MoqSignal_genType_String_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSignal_genType_String_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSignal_genType_String_fnRecorder {
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
				DoFn       MoqSignal_genType_String_doFn
				DoReturnFn MoqSignal_genType_String_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSignal_genType) PrettyParams_String(params MoqSignal_genType_String_params) string {
	return fmt.Sprintf("String()")
}

func (m *MoqSignal_genType) ParamsKey_String(params MoqSignal_genType_String_params, anyParams uint64) MoqSignal_genType_String_paramsKey {
	m.Scene.T.Helper()
	return MoqSignal_genType_String_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqSignal_genType) Reset() { m.ResultsByParams_Signal = nil; m.ResultsByParams_String = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSignal_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Signal {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Signal(results.Params))
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
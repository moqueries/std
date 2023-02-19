// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package net

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that net.IPAddr_starGenType is mocked
// completely
var _ IPAddr_starGenType = (*MoqIPAddr_starGenType_mock)(nil)

// IPAddr_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type IPAddr_starGenType interface {
	Network() string
	String() string
}

// MoqIPAddr_starGenType holds the state of a moq of the IPAddr_starGenType
// type
type MoqIPAddr_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqIPAddr_starGenType_mock

	ResultsByParams_Network []MoqIPAddr_starGenType_Network_resultsByParams
	ResultsByParams_String  []MoqIPAddr_starGenType_String_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Network struct{}
			String  struct{}
		}
	}
}

// MoqIPAddr_starGenType_mock isolates the mock interface of the
// IPAddr_starGenType type
type MoqIPAddr_starGenType_mock struct {
	Moq *MoqIPAddr_starGenType
}

// MoqIPAddr_starGenType_recorder isolates the recorder interface of the
// IPAddr_starGenType type
type MoqIPAddr_starGenType_recorder struct {
	Moq *MoqIPAddr_starGenType
}

// MoqIPAddr_starGenType_Network_params holds the params of the
// IPAddr_starGenType type
type MoqIPAddr_starGenType_Network_params struct{}

// MoqIPAddr_starGenType_Network_paramsKey holds the map key params of the
// IPAddr_starGenType type
type MoqIPAddr_starGenType_Network_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqIPAddr_starGenType_Network_resultsByParams contains the results for a
// given set of parameters for the IPAddr_starGenType type
type MoqIPAddr_starGenType_Network_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqIPAddr_starGenType_Network_paramsKey]*MoqIPAddr_starGenType_Network_results
}

// MoqIPAddr_starGenType_Network_doFn defines the type of function needed when
// calling AndDo for the IPAddr_starGenType type
type MoqIPAddr_starGenType_Network_doFn func()

// MoqIPAddr_starGenType_Network_doReturnFn defines the type of function needed
// when calling DoReturnResults for the IPAddr_starGenType type
type MoqIPAddr_starGenType_Network_doReturnFn func() string

// MoqIPAddr_starGenType_Network_results holds the results of the
// IPAddr_starGenType type
type MoqIPAddr_starGenType_Network_results struct {
	Params  MoqIPAddr_starGenType_Network_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqIPAddr_starGenType_Network_doFn
		DoReturnFn MoqIPAddr_starGenType_Network_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqIPAddr_starGenType_Network_fnRecorder routes recorded function calls to
// the MoqIPAddr_starGenType moq
type MoqIPAddr_starGenType_Network_fnRecorder struct {
	Params    MoqIPAddr_starGenType_Network_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqIPAddr_starGenType_Network_results
	Moq       *MoqIPAddr_starGenType
}

// MoqIPAddr_starGenType_Network_anyParams isolates the any params functions of
// the IPAddr_starGenType type
type MoqIPAddr_starGenType_Network_anyParams struct {
	Recorder *MoqIPAddr_starGenType_Network_fnRecorder
}

// MoqIPAddr_starGenType_String_params holds the params of the
// IPAddr_starGenType type
type MoqIPAddr_starGenType_String_params struct{}

// MoqIPAddr_starGenType_String_paramsKey holds the map key params of the
// IPAddr_starGenType type
type MoqIPAddr_starGenType_String_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqIPAddr_starGenType_String_resultsByParams contains the results for a
// given set of parameters for the IPAddr_starGenType type
type MoqIPAddr_starGenType_String_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqIPAddr_starGenType_String_paramsKey]*MoqIPAddr_starGenType_String_results
}

// MoqIPAddr_starGenType_String_doFn defines the type of function needed when
// calling AndDo for the IPAddr_starGenType type
type MoqIPAddr_starGenType_String_doFn func()

// MoqIPAddr_starGenType_String_doReturnFn defines the type of function needed
// when calling DoReturnResults for the IPAddr_starGenType type
type MoqIPAddr_starGenType_String_doReturnFn func() string

// MoqIPAddr_starGenType_String_results holds the results of the
// IPAddr_starGenType type
type MoqIPAddr_starGenType_String_results struct {
	Params  MoqIPAddr_starGenType_String_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqIPAddr_starGenType_String_doFn
		DoReturnFn MoqIPAddr_starGenType_String_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqIPAddr_starGenType_String_fnRecorder routes recorded function calls to
// the MoqIPAddr_starGenType moq
type MoqIPAddr_starGenType_String_fnRecorder struct {
	Params    MoqIPAddr_starGenType_String_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqIPAddr_starGenType_String_results
	Moq       *MoqIPAddr_starGenType
}

// MoqIPAddr_starGenType_String_anyParams isolates the any params functions of
// the IPAddr_starGenType type
type MoqIPAddr_starGenType_String_anyParams struct {
	Recorder *MoqIPAddr_starGenType_String_fnRecorder
}

// NewMoqIPAddr_starGenType creates a new moq of the IPAddr_starGenType type
func NewMoqIPAddr_starGenType(scene *moq.Scene, config *moq.Config) *MoqIPAddr_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqIPAddr_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqIPAddr_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Network struct{}
				String  struct{}
			}
		}{ParameterIndexing: struct {
			Network struct{}
			String  struct{}
		}{
			Network: struct{}{},
			String:  struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the IPAddr_starGenType type
func (m *MoqIPAddr_starGenType) Mock() *MoqIPAddr_starGenType_mock { return m.Moq }

func (m *MoqIPAddr_starGenType_mock) Network() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqIPAddr_starGenType_Network_params{}
	var results *MoqIPAddr_starGenType_Network_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Network {
		paramsKey := m.Moq.ParamsKey_Network(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Network(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Network(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Network(params))
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

func (m *MoqIPAddr_starGenType_mock) String() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqIPAddr_starGenType_String_params{}
	var results *MoqIPAddr_starGenType_String_results
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

// OnCall returns the recorder implementation of the IPAddr_starGenType type
func (m *MoqIPAddr_starGenType) OnCall() *MoqIPAddr_starGenType_recorder {
	return &MoqIPAddr_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqIPAddr_starGenType_recorder) Network() *MoqIPAddr_starGenType_Network_fnRecorder {
	return &MoqIPAddr_starGenType_Network_fnRecorder{
		Params:   MoqIPAddr_starGenType_Network_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqIPAddr_starGenType_Network_fnRecorder) Any() *MoqIPAddr_starGenType_Network_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Network(r.Params))
		return nil
	}
	return &MoqIPAddr_starGenType_Network_anyParams{Recorder: r}
}

func (r *MoqIPAddr_starGenType_Network_fnRecorder) Seq() *MoqIPAddr_starGenType_Network_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Network(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqIPAddr_starGenType_Network_fnRecorder) NoSeq() *MoqIPAddr_starGenType_Network_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Network(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqIPAddr_starGenType_Network_fnRecorder) ReturnResults(result1 string) *MoqIPAddr_starGenType_Network_fnRecorder {
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
		DoFn       MoqIPAddr_starGenType_Network_doFn
		DoReturnFn MoqIPAddr_starGenType_Network_doReturnFn
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

func (r *MoqIPAddr_starGenType_Network_fnRecorder) AndDo(fn MoqIPAddr_starGenType_Network_doFn) *MoqIPAddr_starGenType_Network_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqIPAddr_starGenType_Network_fnRecorder) DoReturnResults(fn MoqIPAddr_starGenType_Network_doReturnFn) *MoqIPAddr_starGenType_Network_fnRecorder {
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
		DoFn       MoqIPAddr_starGenType_Network_doFn
		DoReturnFn MoqIPAddr_starGenType_Network_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqIPAddr_starGenType_Network_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqIPAddr_starGenType_Network_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Network {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqIPAddr_starGenType_Network_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqIPAddr_starGenType_Network_paramsKey]*MoqIPAddr_starGenType_Network_results{},
		}
		r.Moq.ResultsByParams_Network = append(r.Moq.ResultsByParams_Network, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Network) {
			copy(r.Moq.ResultsByParams_Network[insertAt+1:], r.Moq.ResultsByParams_Network[insertAt:0])
			r.Moq.ResultsByParams_Network[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Network(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqIPAddr_starGenType_Network_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqIPAddr_starGenType_Network_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqIPAddr_starGenType_Network_fnRecorder {
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
				DoFn       MoqIPAddr_starGenType_Network_doFn
				DoReturnFn MoqIPAddr_starGenType_Network_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqIPAddr_starGenType) PrettyParams_Network(params MoqIPAddr_starGenType_Network_params) string {
	return fmt.Sprintf("Network()")
}

func (m *MoqIPAddr_starGenType) ParamsKey_Network(params MoqIPAddr_starGenType_Network_params, anyParams uint64) MoqIPAddr_starGenType_Network_paramsKey {
	m.Scene.T.Helper()
	return MoqIPAddr_starGenType_Network_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqIPAddr_starGenType_recorder) String() *MoqIPAddr_starGenType_String_fnRecorder {
	return &MoqIPAddr_starGenType_String_fnRecorder{
		Params:   MoqIPAddr_starGenType_String_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqIPAddr_starGenType_String_fnRecorder) Any() *MoqIPAddr_starGenType_String_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	return &MoqIPAddr_starGenType_String_anyParams{Recorder: r}
}

func (r *MoqIPAddr_starGenType_String_fnRecorder) Seq() *MoqIPAddr_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqIPAddr_starGenType_String_fnRecorder) NoSeq() *MoqIPAddr_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqIPAddr_starGenType_String_fnRecorder) ReturnResults(result1 string) *MoqIPAddr_starGenType_String_fnRecorder {
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
		DoFn       MoqIPAddr_starGenType_String_doFn
		DoReturnFn MoqIPAddr_starGenType_String_doReturnFn
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

func (r *MoqIPAddr_starGenType_String_fnRecorder) AndDo(fn MoqIPAddr_starGenType_String_doFn) *MoqIPAddr_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqIPAddr_starGenType_String_fnRecorder) DoReturnResults(fn MoqIPAddr_starGenType_String_doReturnFn) *MoqIPAddr_starGenType_String_fnRecorder {
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
		DoFn       MoqIPAddr_starGenType_String_doFn
		DoReturnFn MoqIPAddr_starGenType_String_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqIPAddr_starGenType_String_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqIPAddr_starGenType_String_resultsByParams
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
		results = &MoqIPAddr_starGenType_String_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqIPAddr_starGenType_String_paramsKey]*MoqIPAddr_starGenType_String_results{},
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
		r.Results = &MoqIPAddr_starGenType_String_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqIPAddr_starGenType_String_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqIPAddr_starGenType_String_fnRecorder {
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
				DoFn       MoqIPAddr_starGenType_String_doFn
				DoReturnFn MoqIPAddr_starGenType_String_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqIPAddr_starGenType) PrettyParams_String(params MoqIPAddr_starGenType_String_params) string {
	return fmt.Sprintf("String()")
}

func (m *MoqIPAddr_starGenType) ParamsKey_String(params MoqIPAddr_starGenType_String_params, anyParams uint64) MoqIPAddr_starGenType_String_paramsKey {
	m.Scene.T.Helper()
	return MoqIPAddr_starGenType_String_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqIPAddr_starGenType) Reset() {
	m.ResultsByParams_Network = nil
	m.ResultsByParams_String = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqIPAddr_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Network {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Network(results.Params))
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

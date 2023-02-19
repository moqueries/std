// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package net

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that net.AddrError_starGenType is
// mocked completely
var _ AddrError_starGenType = (*MoqAddrError_starGenType_mock)(nil)

// AddrError_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type AddrError_starGenType interface {
	Error() string
	Timeout() bool
	Temporary() bool
}

// MoqAddrError_starGenType holds the state of a moq of the
// AddrError_starGenType type
type MoqAddrError_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqAddrError_starGenType_mock

	ResultsByParams_Error     []MoqAddrError_starGenType_Error_resultsByParams
	ResultsByParams_Timeout   []MoqAddrError_starGenType_Timeout_resultsByParams
	ResultsByParams_Temporary []MoqAddrError_starGenType_Temporary_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Error     struct{}
			Timeout   struct{}
			Temporary struct{}
		}
	}
}

// MoqAddrError_starGenType_mock isolates the mock interface of the
// AddrError_starGenType type
type MoqAddrError_starGenType_mock struct {
	Moq *MoqAddrError_starGenType
}

// MoqAddrError_starGenType_recorder isolates the recorder interface of the
// AddrError_starGenType type
type MoqAddrError_starGenType_recorder struct {
	Moq *MoqAddrError_starGenType
}

// MoqAddrError_starGenType_Error_params holds the params of the
// AddrError_starGenType type
type MoqAddrError_starGenType_Error_params struct{}

// MoqAddrError_starGenType_Error_paramsKey holds the map key params of the
// AddrError_starGenType type
type MoqAddrError_starGenType_Error_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqAddrError_starGenType_Error_resultsByParams contains the results for a
// given set of parameters for the AddrError_starGenType type
type MoqAddrError_starGenType_Error_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqAddrError_starGenType_Error_paramsKey]*MoqAddrError_starGenType_Error_results
}

// MoqAddrError_starGenType_Error_doFn defines the type of function needed when
// calling AndDo for the AddrError_starGenType type
type MoqAddrError_starGenType_Error_doFn func()

// MoqAddrError_starGenType_Error_doReturnFn defines the type of function
// needed when calling DoReturnResults for the AddrError_starGenType type
type MoqAddrError_starGenType_Error_doReturnFn func() string

// MoqAddrError_starGenType_Error_results holds the results of the
// AddrError_starGenType type
type MoqAddrError_starGenType_Error_results struct {
	Params  MoqAddrError_starGenType_Error_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqAddrError_starGenType_Error_doFn
		DoReturnFn MoqAddrError_starGenType_Error_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqAddrError_starGenType_Error_fnRecorder routes recorded function calls to
// the MoqAddrError_starGenType moq
type MoqAddrError_starGenType_Error_fnRecorder struct {
	Params    MoqAddrError_starGenType_Error_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqAddrError_starGenType_Error_results
	Moq       *MoqAddrError_starGenType
}

// MoqAddrError_starGenType_Error_anyParams isolates the any params functions
// of the AddrError_starGenType type
type MoqAddrError_starGenType_Error_anyParams struct {
	Recorder *MoqAddrError_starGenType_Error_fnRecorder
}

// MoqAddrError_starGenType_Timeout_params holds the params of the
// AddrError_starGenType type
type MoqAddrError_starGenType_Timeout_params struct{}

// MoqAddrError_starGenType_Timeout_paramsKey holds the map key params of the
// AddrError_starGenType type
type MoqAddrError_starGenType_Timeout_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqAddrError_starGenType_Timeout_resultsByParams contains the results for a
// given set of parameters for the AddrError_starGenType type
type MoqAddrError_starGenType_Timeout_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqAddrError_starGenType_Timeout_paramsKey]*MoqAddrError_starGenType_Timeout_results
}

// MoqAddrError_starGenType_Timeout_doFn defines the type of function needed
// when calling AndDo for the AddrError_starGenType type
type MoqAddrError_starGenType_Timeout_doFn func()

// MoqAddrError_starGenType_Timeout_doReturnFn defines the type of function
// needed when calling DoReturnResults for the AddrError_starGenType type
type MoqAddrError_starGenType_Timeout_doReturnFn func() bool

// MoqAddrError_starGenType_Timeout_results holds the results of the
// AddrError_starGenType type
type MoqAddrError_starGenType_Timeout_results struct {
	Params  MoqAddrError_starGenType_Timeout_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqAddrError_starGenType_Timeout_doFn
		DoReturnFn MoqAddrError_starGenType_Timeout_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqAddrError_starGenType_Timeout_fnRecorder routes recorded function calls
// to the MoqAddrError_starGenType moq
type MoqAddrError_starGenType_Timeout_fnRecorder struct {
	Params    MoqAddrError_starGenType_Timeout_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqAddrError_starGenType_Timeout_results
	Moq       *MoqAddrError_starGenType
}

// MoqAddrError_starGenType_Timeout_anyParams isolates the any params functions
// of the AddrError_starGenType type
type MoqAddrError_starGenType_Timeout_anyParams struct {
	Recorder *MoqAddrError_starGenType_Timeout_fnRecorder
}

// MoqAddrError_starGenType_Temporary_params holds the params of the
// AddrError_starGenType type
type MoqAddrError_starGenType_Temporary_params struct{}

// MoqAddrError_starGenType_Temporary_paramsKey holds the map key params of the
// AddrError_starGenType type
type MoqAddrError_starGenType_Temporary_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqAddrError_starGenType_Temporary_resultsByParams contains the results for
// a given set of parameters for the AddrError_starGenType type
type MoqAddrError_starGenType_Temporary_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqAddrError_starGenType_Temporary_paramsKey]*MoqAddrError_starGenType_Temporary_results
}

// MoqAddrError_starGenType_Temporary_doFn defines the type of function needed
// when calling AndDo for the AddrError_starGenType type
type MoqAddrError_starGenType_Temporary_doFn func()

// MoqAddrError_starGenType_Temporary_doReturnFn defines the type of function
// needed when calling DoReturnResults for the AddrError_starGenType type
type MoqAddrError_starGenType_Temporary_doReturnFn func() bool

// MoqAddrError_starGenType_Temporary_results holds the results of the
// AddrError_starGenType type
type MoqAddrError_starGenType_Temporary_results struct {
	Params  MoqAddrError_starGenType_Temporary_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqAddrError_starGenType_Temporary_doFn
		DoReturnFn MoqAddrError_starGenType_Temporary_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqAddrError_starGenType_Temporary_fnRecorder routes recorded function calls
// to the MoqAddrError_starGenType moq
type MoqAddrError_starGenType_Temporary_fnRecorder struct {
	Params    MoqAddrError_starGenType_Temporary_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqAddrError_starGenType_Temporary_results
	Moq       *MoqAddrError_starGenType
}

// MoqAddrError_starGenType_Temporary_anyParams isolates the any params
// functions of the AddrError_starGenType type
type MoqAddrError_starGenType_Temporary_anyParams struct {
	Recorder *MoqAddrError_starGenType_Temporary_fnRecorder
}

// NewMoqAddrError_starGenType creates a new moq of the AddrError_starGenType
// type
func NewMoqAddrError_starGenType(scene *moq.Scene, config *moq.Config) *MoqAddrError_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqAddrError_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqAddrError_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Error     struct{}
				Timeout   struct{}
				Temporary struct{}
			}
		}{ParameterIndexing: struct {
			Error     struct{}
			Timeout   struct{}
			Temporary struct{}
		}{
			Error:     struct{}{},
			Timeout:   struct{}{},
			Temporary: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the AddrError_starGenType type
func (m *MoqAddrError_starGenType) Mock() *MoqAddrError_starGenType_mock { return m.Moq }

func (m *MoqAddrError_starGenType_mock) Error() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqAddrError_starGenType_Error_params{}
	var results *MoqAddrError_starGenType_Error_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Error {
		paramsKey := m.Moq.ParamsKey_Error(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Error(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Error(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Error(params))
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

func (m *MoqAddrError_starGenType_mock) Timeout() (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqAddrError_starGenType_Timeout_params{}
	var results *MoqAddrError_starGenType_Timeout_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Timeout {
		paramsKey := m.Moq.ParamsKey_Timeout(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Timeout(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Timeout(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Timeout(params))
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

func (m *MoqAddrError_starGenType_mock) Temporary() (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqAddrError_starGenType_Temporary_params{}
	var results *MoqAddrError_starGenType_Temporary_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Temporary {
		paramsKey := m.Moq.ParamsKey_Temporary(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Temporary(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Temporary(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Temporary(params))
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

// OnCall returns the recorder implementation of the AddrError_starGenType type
func (m *MoqAddrError_starGenType) OnCall() *MoqAddrError_starGenType_recorder {
	return &MoqAddrError_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqAddrError_starGenType_recorder) Error() *MoqAddrError_starGenType_Error_fnRecorder {
	return &MoqAddrError_starGenType_Error_fnRecorder{
		Params:   MoqAddrError_starGenType_Error_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqAddrError_starGenType_Error_fnRecorder) Any() *MoqAddrError_starGenType_Error_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	return &MoqAddrError_starGenType_Error_anyParams{Recorder: r}
}

func (r *MoqAddrError_starGenType_Error_fnRecorder) Seq() *MoqAddrError_starGenType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqAddrError_starGenType_Error_fnRecorder) NoSeq() *MoqAddrError_starGenType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqAddrError_starGenType_Error_fnRecorder) ReturnResults(result1 string) *MoqAddrError_starGenType_Error_fnRecorder {
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
		DoFn       MoqAddrError_starGenType_Error_doFn
		DoReturnFn MoqAddrError_starGenType_Error_doReturnFn
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

func (r *MoqAddrError_starGenType_Error_fnRecorder) AndDo(fn MoqAddrError_starGenType_Error_doFn) *MoqAddrError_starGenType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqAddrError_starGenType_Error_fnRecorder) DoReturnResults(fn MoqAddrError_starGenType_Error_doReturnFn) *MoqAddrError_starGenType_Error_fnRecorder {
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
		DoFn       MoqAddrError_starGenType_Error_doFn
		DoReturnFn MoqAddrError_starGenType_Error_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqAddrError_starGenType_Error_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqAddrError_starGenType_Error_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Error {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqAddrError_starGenType_Error_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqAddrError_starGenType_Error_paramsKey]*MoqAddrError_starGenType_Error_results{},
		}
		r.Moq.ResultsByParams_Error = append(r.Moq.ResultsByParams_Error, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Error) {
			copy(r.Moq.ResultsByParams_Error[insertAt+1:], r.Moq.ResultsByParams_Error[insertAt:0])
			r.Moq.ResultsByParams_Error[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Error(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqAddrError_starGenType_Error_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqAddrError_starGenType_Error_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqAddrError_starGenType_Error_fnRecorder {
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
				DoFn       MoqAddrError_starGenType_Error_doFn
				DoReturnFn MoqAddrError_starGenType_Error_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqAddrError_starGenType) PrettyParams_Error(params MoqAddrError_starGenType_Error_params) string {
	return fmt.Sprintf("Error()")
}

func (m *MoqAddrError_starGenType) ParamsKey_Error(params MoqAddrError_starGenType_Error_params, anyParams uint64) MoqAddrError_starGenType_Error_paramsKey {
	m.Scene.T.Helper()
	return MoqAddrError_starGenType_Error_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqAddrError_starGenType_recorder) Timeout() *MoqAddrError_starGenType_Timeout_fnRecorder {
	return &MoqAddrError_starGenType_Timeout_fnRecorder{
		Params:   MoqAddrError_starGenType_Timeout_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqAddrError_starGenType_Timeout_fnRecorder) Any() *MoqAddrError_starGenType_Timeout_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Timeout(r.Params))
		return nil
	}
	return &MoqAddrError_starGenType_Timeout_anyParams{Recorder: r}
}

func (r *MoqAddrError_starGenType_Timeout_fnRecorder) Seq() *MoqAddrError_starGenType_Timeout_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Timeout(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqAddrError_starGenType_Timeout_fnRecorder) NoSeq() *MoqAddrError_starGenType_Timeout_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Timeout(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqAddrError_starGenType_Timeout_fnRecorder) ReturnResults(result1 bool) *MoqAddrError_starGenType_Timeout_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqAddrError_starGenType_Timeout_doFn
		DoReturnFn MoqAddrError_starGenType_Timeout_doReturnFn
	}{
		Values: &struct {
			Result1 bool
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqAddrError_starGenType_Timeout_fnRecorder) AndDo(fn MoqAddrError_starGenType_Timeout_doFn) *MoqAddrError_starGenType_Timeout_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqAddrError_starGenType_Timeout_fnRecorder) DoReturnResults(fn MoqAddrError_starGenType_Timeout_doReturnFn) *MoqAddrError_starGenType_Timeout_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqAddrError_starGenType_Timeout_doFn
		DoReturnFn MoqAddrError_starGenType_Timeout_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqAddrError_starGenType_Timeout_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqAddrError_starGenType_Timeout_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Timeout {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqAddrError_starGenType_Timeout_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqAddrError_starGenType_Timeout_paramsKey]*MoqAddrError_starGenType_Timeout_results{},
		}
		r.Moq.ResultsByParams_Timeout = append(r.Moq.ResultsByParams_Timeout, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Timeout) {
			copy(r.Moq.ResultsByParams_Timeout[insertAt+1:], r.Moq.ResultsByParams_Timeout[insertAt:0])
			r.Moq.ResultsByParams_Timeout[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Timeout(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqAddrError_starGenType_Timeout_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqAddrError_starGenType_Timeout_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqAddrError_starGenType_Timeout_fnRecorder {
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
					Result1 bool
				}
				Sequence   uint32
				DoFn       MoqAddrError_starGenType_Timeout_doFn
				DoReturnFn MoqAddrError_starGenType_Timeout_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqAddrError_starGenType) PrettyParams_Timeout(params MoqAddrError_starGenType_Timeout_params) string {
	return fmt.Sprintf("Timeout()")
}

func (m *MoqAddrError_starGenType) ParamsKey_Timeout(params MoqAddrError_starGenType_Timeout_params, anyParams uint64) MoqAddrError_starGenType_Timeout_paramsKey {
	m.Scene.T.Helper()
	return MoqAddrError_starGenType_Timeout_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqAddrError_starGenType_recorder) Temporary() *MoqAddrError_starGenType_Temporary_fnRecorder {
	return &MoqAddrError_starGenType_Temporary_fnRecorder{
		Params:   MoqAddrError_starGenType_Temporary_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqAddrError_starGenType_Temporary_fnRecorder) Any() *MoqAddrError_starGenType_Temporary_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Temporary(r.Params))
		return nil
	}
	return &MoqAddrError_starGenType_Temporary_anyParams{Recorder: r}
}

func (r *MoqAddrError_starGenType_Temporary_fnRecorder) Seq() *MoqAddrError_starGenType_Temporary_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Temporary(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqAddrError_starGenType_Temporary_fnRecorder) NoSeq() *MoqAddrError_starGenType_Temporary_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Temporary(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqAddrError_starGenType_Temporary_fnRecorder) ReturnResults(result1 bool) *MoqAddrError_starGenType_Temporary_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqAddrError_starGenType_Temporary_doFn
		DoReturnFn MoqAddrError_starGenType_Temporary_doReturnFn
	}{
		Values: &struct {
			Result1 bool
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqAddrError_starGenType_Temporary_fnRecorder) AndDo(fn MoqAddrError_starGenType_Temporary_doFn) *MoqAddrError_starGenType_Temporary_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqAddrError_starGenType_Temporary_fnRecorder) DoReturnResults(fn MoqAddrError_starGenType_Temporary_doReturnFn) *MoqAddrError_starGenType_Temporary_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqAddrError_starGenType_Temporary_doFn
		DoReturnFn MoqAddrError_starGenType_Temporary_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqAddrError_starGenType_Temporary_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqAddrError_starGenType_Temporary_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Temporary {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqAddrError_starGenType_Temporary_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqAddrError_starGenType_Temporary_paramsKey]*MoqAddrError_starGenType_Temporary_results{},
		}
		r.Moq.ResultsByParams_Temporary = append(r.Moq.ResultsByParams_Temporary, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Temporary) {
			copy(r.Moq.ResultsByParams_Temporary[insertAt+1:], r.Moq.ResultsByParams_Temporary[insertAt:0])
			r.Moq.ResultsByParams_Temporary[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Temporary(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqAddrError_starGenType_Temporary_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqAddrError_starGenType_Temporary_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqAddrError_starGenType_Temporary_fnRecorder {
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
					Result1 bool
				}
				Sequence   uint32
				DoFn       MoqAddrError_starGenType_Temporary_doFn
				DoReturnFn MoqAddrError_starGenType_Temporary_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqAddrError_starGenType) PrettyParams_Temporary(params MoqAddrError_starGenType_Temporary_params) string {
	return fmt.Sprintf("Temporary()")
}

func (m *MoqAddrError_starGenType) ParamsKey_Temporary(params MoqAddrError_starGenType_Temporary_params, anyParams uint64) MoqAddrError_starGenType_Temporary_paramsKey {
	m.Scene.T.Helper()
	return MoqAddrError_starGenType_Temporary_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqAddrError_starGenType) Reset() {
	m.ResultsByParams_Error = nil
	m.ResultsByParams_Timeout = nil
	m.ResultsByParams_Temporary = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqAddrError_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Error {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Error(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Timeout {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Timeout(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Temporary {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Temporary(results.Params))
			}
		}
	}
}

// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package x509

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that x509.SystemRootsError_genType is
// mocked completely
var _ SystemRootsError_genType = (*MoqSystemRootsError_genType_mock)(nil)

// SystemRootsError_genType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type SystemRootsError_genType interface {
	Error() string
	Unwrap() error
}

// MoqSystemRootsError_genType holds the state of a moq of the
// SystemRootsError_genType type
type MoqSystemRootsError_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSystemRootsError_genType_mock

	ResultsByParams_Error  []MoqSystemRootsError_genType_Error_resultsByParams
	ResultsByParams_Unwrap []MoqSystemRootsError_genType_Unwrap_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Error  struct{}
			Unwrap struct{}
		}
	}
}

// MoqSystemRootsError_genType_mock isolates the mock interface of the
// SystemRootsError_genType type
type MoqSystemRootsError_genType_mock struct {
	Moq *MoqSystemRootsError_genType
}

// MoqSystemRootsError_genType_recorder isolates the recorder interface of the
// SystemRootsError_genType type
type MoqSystemRootsError_genType_recorder struct {
	Moq *MoqSystemRootsError_genType
}

// MoqSystemRootsError_genType_Error_params holds the params of the
// SystemRootsError_genType type
type MoqSystemRootsError_genType_Error_params struct{}

// MoqSystemRootsError_genType_Error_paramsKey holds the map key params of the
// SystemRootsError_genType type
type MoqSystemRootsError_genType_Error_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqSystemRootsError_genType_Error_resultsByParams contains the results for a
// given set of parameters for the SystemRootsError_genType type
type MoqSystemRootsError_genType_Error_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSystemRootsError_genType_Error_paramsKey]*MoqSystemRootsError_genType_Error_results
}

// MoqSystemRootsError_genType_Error_doFn defines the type of function needed
// when calling AndDo for the SystemRootsError_genType type
type MoqSystemRootsError_genType_Error_doFn func()

// MoqSystemRootsError_genType_Error_doReturnFn defines the type of function
// needed when calling DoReturnResults for the SystemRootsError_genType type
type MoqSystemRootsError_genType_Error_doReturnFn func() string

// MoqSystemRootsError_genType_Error_results holds the results of the
// SystemRootsError_genType type
type MoqSystemRootsError_genType_Error_results struct {
	Params  MoqSystemRootsError_genType_Error_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqSystemRootsError_genType_Error_doFn
		DoReturnFn MoqSystemRootsError_genType_Error_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSystemRootsError_genType_Error_fnRecorder routes recorded function calls
// to the MoqSystemRootsError_genType moq
type MoqSystemRootsError_genType_Error_fnRecorder struct {
	Params    MoqSystemRootsError_genType_Error_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSystemRootsError_genType_Error_results
	Moq       *MoqSystemRootsError_genType
}

// MoqSystemRootsError_genType_Error_anyParams isolates the any params
// functions of the SystemRootsError_genType type
type MoqSystemRootsError_genType_Error_anyParams struct {
	Recorder *MoqSystemRootsError_genType_Error_fnRecorder
}

// MoqSystemRootsError_genType_Unwrap_params holds the params of the
// SystemRootsError_genType type
type MoqSystemRootsError_genType_Unwrap_params struct{}

// MoqSystemRootsError_genType_Unwrap_paramsKey holds the map key params of the
// SystemRootsError_genType type
type MoqSystemRootsError_genType_Unwrap_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqSystemRootsError_genType_Unwrap_resultsByParams contains the results for
// a given set of parameters for the SystemRootsError_genType type
type MoqSystemRootsError_genType_Unwrap_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSystemRootsError_genType_Unwrap_paramsKey]*MoqSystemRootsError_genType_Unwrap_results
}

// MoqSystemRootsError_genType_Unwrap_doFn defines the type of function needed
// when calling AndDo for the SystemRootsError_genType type
type MoqSystemRootsError_genType_Unwrap_doFn func()

// MoqSystemRootsError_genType_Unwrap_doReturnFn defines the type of function
// needed when calling DoReturnResults for the SystemRootsError_genType type
type MoqSystemRootsError_genType_Unwrap_doReturnFn func() error

// MoqSystemRootsError_genType_Unwrap_results holds the results of the
// SystemRootsError_genType type
type MoqSystemRootsError_genType_Unwrap_results struct {
	Params  MoqSystemRootsError_genType_Unwrap_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqSystemRootsError_genType_Unwrap_doFn
		DoReturnFn MoqSystemRootsError_genType_Unwrap_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSystemRootsError_genType_Unwrap_fnRecorder routes recorded function calls
// to the MoqSystemRootsError_genType moq
type MoqSystemRootsError_genType_Unwrap_fnRecorder struct {
	Params    MoqSystemRootsError_genType_Unwrap_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSystemRootsError_genType_Unwrap_results
	Moq       *MoqSystemRootsError_genType
}

// MoqSystemRootsError_genType_Unwrap_anyParams isolates the any params
// functions of the SystemRootsError_genType type
type MoqSystemRootsError_genType_Unwrap_anyParams struct {
	Recorder *MoqSystemRootsError_genType_Unwrap_fnRecorder
}

// NewMoqSystemRootsError_genType creates a new moq of the
// SystemRootsError_genType type
func NewMoqSystemRootsError_genType(scene *moq.Scene, config *moq.Config) *MoqSystemRootsError_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSystemRootsError_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSystemRootsError_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Error  struct{}
				Unwrap struct{}
			}
		}{ParameterIndexing: struct {
			Error  struct{}
			Unwrap struct{}
		}{
			Error:  struct{}{},
			Unwrap: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the SystemRootsError_genType type
func (m *MoqSystemRootsError_genType) Mock() *MoqSystemRootsError_genType_mock { return m.Moq }

func (m *MoqSystemRootsError_genType_mock) Error() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqSystemRootsError_genType_Error_params{}
	var results *MoqSystemRootsError_genType_Error_results
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

func (m *MoqSystemRootsError_genType_mock) Unwrap() (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqSystemRootsError_genType_Unwrap_params{}
	var results *MoqSystemRootsError_genType_Unwrap_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Unwrap {
		paramsKey := m.Moq.ParamsKey_Unwrap(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Unwrap(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Unwrap(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Unwrap(params))
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

// OnCall returns the recorder implementation of the SystemRootsError_genType
// type
func (m *MoqSystemRootsError_genType) OnCall() *MoqSystemRootsError_genType_recorder {
	return &MoqSystemRootsError_genType_recorder{
		Moq: m,
	}
}

func (m *MoqSystemRootsError_genType_recorder) Error() *MoqSystemRootsError_genType_Error_fnRecorder {
	return &MoqSystemRootsError_genType_Error_fnRecorder{
		Params:   MoqSystemRootsError_genType_Error_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqSystemRootsError_genType_Error_fnRecorder) Any() *MoqSystemRootsError_genType_Error_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	return &MoqSystemRootsError_genType_Error_anyParams{Recorder: r}
}

func (r *MoqSystemRootsError_genType_Error_fnRecorder) Seq() *MoqSystemRootsError_genType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSystemRootsError_genType_Error_fnRecorder) NoSeq() *MoqSystemRootsError_genType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSystemRootsError_genType_Error_fnRecorder) ReturnResults(result1 string) *MoqSystemRootsError_genType_Error_fnRecorder {
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
		DoFn       MoqSystemRootsError_genType_Error_doFn
		DoReturnFn MoqSystemRootsError_genType_Error_doReturnFn
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

func (r *MoqSystemRootsError_genType_Error_fnRecorder) AndDo(fn MoqSystemRootsError_genType_Error_doFn) *MoqSystemRootsError_genType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSystemRootsError_genType_Error_fnRecorder) DoReturnResults(fn MoqSystemRootsError_genType_Error_doReturnFn) *MoqSystemRootsError_genType_Error_fnRecorder {
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
		DoFn       MoqSystemRootsError_genType_Error_doFn
		DoReturnFn MoqSystemRootsError_genType_Error_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSystemRootsError_genType_Error_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSystemRootsError_genType_Error_resultsByParams
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
		results = &MoqSystemRootsError_genType_Error_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSystemRootsError_genType_Error_paramsKey]*MoqSystemRootsError_genType_Error_results{},
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
		r.Results = &MoqSystemRootsError_genType_Error_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSystemRootsError_genType_Error_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSystemRootsError_genType_Error_fnRecorder {
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
				DoFn       MoqSystemRootsError_genType_Error_doFn
				DoReturnFn MoqSystemRootsError_genType_Error_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSystemRootsError_genType) PrettyParams_Error(params MoqSystemRootsError_genType_Error_params) string {
	return fmt.Sprintf("Error()")
}

func (m *MoqSystemRootsError_genType) ParamsKey_Error(params MoqSystemRootsError_genType_Error_params, anyParams uint64) MoqSystemRootsError_genType_Error_paramsKey {
	m.Scene.T.Helper()
	return MoqSystemRootsError_genType_Error_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqSystemRootsError_genType_recorder) Unwrap() *MoqSystemRootsError_genType_Unwrap_fnRecorder {
	return &MoqSystemRootsError_genType_Unwrap_fnRecorder{
		Params:   MoqSystemRootsError_genType_Unwrap_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqSystemRootsError_genType_Unwrap_fnRecorder) Any() *MoqSystemRootsError_genType_Unwrap_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Unwrap(r.Params))
		return nil
	}
	return &MoqSystemRootsError_genType_Unwrap_anyParams{Recorder: r}
}

func (r *MoqSystemRootsError_genType_Unwrap_fnRecorder) Seq() *MoqSystemRootsError_genType_Unwrap_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Unwrap(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSystemRootsError_genType_Unwrap_fnRecorder) NoSeq() *MoqSystemRootsError_genType_Unwrap_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Unwrap(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSystemRootsError_genType_Unwrap_fnRecorder) ReturnResults(result1 error) *MoqSystemRootsError_genType_Unwrap_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqSystemRootsError_genType_Unwrap_doFn
		DoReturnFn MoqSystemRootsError_genType_Unwrap_doReturnFn
	}{
		Values: &struct {
			Result1 error
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSystemRootsError_genType_Unwrap_fnRecorder) AndDo(fn MoqSystemRootsError_genType_Unwrap_doFn) *MoqSystemRootsError_genType_Unwrap_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSystemRootsError_genType_Unwrap_fnRecorder) DoReturnResults(fn MoqSystemRootsError_genType_Unwrap_doReturnFn) *MoqSystemRootsError_genType_Unwrap_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqSystemRootsError_genType_Unwrap_doFn
		DoReturnFn MoqSystemRootsError_genType_Unwrap_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSystemRootsError_genType_Unwrap_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSystemRootsError_genType_Unwrap_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Unwrap {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqSystemRootsError_genType_Unwrap_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSystemRootsError_genType_Unwrap_paramsKey]*MoqSystemRootsError_genType_Unwrap_results{},
		}
		r.Moq.ResultsByParams_Unwrap = append(r.Moq.ResultsByParams_Unwrap, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Unwrap) {
			copy(r.Moq.ResultsByParams_Unwrap[insertAt+1:], r.Moq.ResultsByParams_Unwrap[insertAt:0])
			r.Moq.ResultsByParams_Unwrap[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Unwrap(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqSystemRootsError_genType_Unwrap_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSystemRootsError_genType_Unwrap_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSystemRootsError_genType_Unwrap_fnRecorder {
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
					Result1 error
				}
				Sequence   uint32
				DoFn       MoqSystemRootsError_genType_Unwrap_doFn
				DoReturnFn MoqSystemRootsError_genType_Unwrap_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSystemRootsError_genType) PrettyParams_Unwrap(params MoqSystemRootsError_genType_Unwrap_params) string {
	return fmt.Sprintf("Unwrap()")
}

func (m *MoqSystemRootsError_genType) ParamsKey_Unwrap(params MoqSystemRootsError_genType_Unwrap_params, anyParams uint64) MoqSystemRootsError_genType_Unwrap_paramsKey {
	m.Scene.T.Helper()
	return MoqSystemRootsError_genType_Unwrap_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqSystemRootsError_genType) Reset() {
	m.ResultsByParams_Error = nil
	m.ResultsByParams_Unwrap = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSystemRootsError_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Error {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Error(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Unwrap {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Unwrap(results.Params))
			}
		}
	}
}

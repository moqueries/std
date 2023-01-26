// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that syscall.Errno_genType is mocked
// completely
var _ Errno_genType = (*MoqErrno_genType_mock)(nil)

// Errno_genType is the fabricated implementation type of this mock (emitted
// when mocking a collections of methods directly and not from an interface
// type)
type Errno_genType interface {
	Error() string
	Is(target error) bool
	Temporary() bool
	Timeout() bool
}

// MoqErrno_genType holds the state of a moq of the Errno_genType type
type MoqErrno_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqErrno_genType_mock

	ResultsByParams_Error     []MoqErrno_genType_Error_resultsByParams
	ResultsByParams_Is        []MoqErrno_genType_Is_resultsByParams
	ResultsByParams_Temporary []MoqErrno_genType_Temporary_resultsByParams
	ResultsByParams_Timeout   []MoqErrno_genType_Timeout_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Error struct{}
			Is    struct {
				Target moq.ParamIndexing
			}
			Temporary struct{}
			Timeout   struct{}
		}
	}
}

// MoqErrno_genType_mock isolates the mock interface of the Errno_genType type
type MoqErrno_genType_mock struct {
	Moq *MoqErrno_genType
}

// MoqErrno_genType_recorder isolates the recorder interface of the
// Errno_genType type
type MoqErrno_genType_recorder struct {
	Moq *MoqErrno_genType
}

// MoqErrno_genType_Error_params holds the params of the Errno_genType type
type MoqErrno_genType_Error_params struct{}

// MoqErrno_genType_Error_paramsKey holds the map key params of the
// Errno_genType type
type MoqErrno_genType_Error_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqErrno_genType_Error_resultsByParams contains the results for a given set
// of parameters for the Errno_genType type
type MoqErrno_genType_Error_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqErrno_genType_Error_paramsKey]*MoqErrno_genType_Error_results
}

// MoqErrno_genType_Error_doFn defines the type of function needed when calling
// AndDo for the Errno_genType type
type MoqErrno_genType_Error_doFn func()

// MoqErrno_genType_Error_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Errno_genType type
type MoqErrno_genType_Error_doReturnFn func() string

// MoqErrno_genType_Error_results holds the results of the Errno_genType type
type MoqErrno_genType_Error_results struct {
	Params  MoqErrno_genType_Error_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqErrno_genType_Error_doFn
		DoReturnFn MoqErrno_genType_Error_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqErrno_genType_Error_fnRecorder routes recorded function calls to the
// MoqErrno_genType moq
type MoqErrno_genType_Error_fnRecorder struct {
	Params    MoqErrno_genType_Error_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqErrno_genType_Error_results
	Moq       *MoqErrno_genType
}

// MoqErrno_genType_Error_anyParams isolates the any params functions of the
// Errno_genType type
type MoqErrno_genType_Error_anyParams struct {
	Recorder *MoqErrno_genType_Error_fnRecorder
}

// MoqErrno_genType_Is_params holds the params of the Errno_genType type
type MoqErrno_genType_Is_params struct{ Target error }

// MoqErrno_genType_Is_paramsKey holds the map key params of the Errno_genType
// type
type MoqErrno_genType_Is_paramsKey struct {
	Params struct{ Target error }
	Hashes struct{ Target hash.Hash }
}

// MoqErrno_genType_Is_resultsByParams contains the results for a given set of
// parameters for the Errno_genType type
type MoqErrno_genType_Is_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqErrno_genType_Is_paramsKey]*MoqErrno_genType_Is_results
}

// MoqErrno_genType_Is_doFn defines the type of function needed when calling
// AndDo for the Errno_genType type
type MoqErrno_genType_Is_doFn func(target error)

// MoqErrno_genType_Is_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Errno_genType type
type MoqErrno_genType_Is_doReturnFn func(target error) bool

// MoqErrno_genType_Is_results holds the results of the Errno_genType type
type MoqErrno_genType_Is_results struct {
	Params  MoqErrno_genType_Is_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqErrno_genType_Is_doFn
		DoReturnFn MoqErrno_genType_Is_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqErrno_genType_Is_fnRecorder routes recorded function calls to the
// MoqErrno_genType moq
type MoqErrno_genType_Is_fnRecorder struct {
	Params    MoqErrno_genType_Is_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqErrno_genType_Is_results
	Moq       *MoqErrno_genType
}

// MoqErrno_genType_Is_anyParams isolates the any params functions of the
// Errno_genType type
type MoqErrno_genType_Is_anyParams struct {
	Recorder *MoqErrno_genType_Is_fnRecorder
}

// MoqErrno_genType_Temporary_params holds the params of the Errno_genType type
type MoqErrno_genType_Temporary_params struct{}

// MoqErrno_genType_Temporary_paramsKey holds the map key params of the
// Errno_genType type
type MoqErrno_genType_Temporary_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqErrno_genType_Temporary_resultsByParams contains the results for a given
// set of parameters for the Errno_genType type
type MoqErrno_genType_Temporary_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqErrno_genType_Temporary_paramsKey]*MoqErrno_genType_Temporary_results
}

// MoqErrno_genType_Temporary_doFn defines the type of function needed when
// calling AndDo for the Errno_genType type
type MoqErrno_genType_Temporary_doFn func()

// MoqErrno_genType_Temporary_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Errno_genType type
type MoqErrno_genType_Temporary_doReturnFn func() bool

// MoqErrno_genType_Temporary_results holds the results of the Errno_genType
// type
type MoqErrno_genType_Temporary_results struct {
	Params  MoqErrno_genType_Temporary_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqErrno_genType_Temporary_doFn
		DoReturnFn MoqErrno_genType_Temporary_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqErrno_genType_Temporary_fnRecorder routes recorded function calls to the
// MoqErrno_genType moq
type MoqErrno_genType_Temporary_fnRecorder struct {
	Params    MoqErrno_genType_Temporary_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqErrno_genType_Temporary_results
	Moq       *MoqErrno_genType
}

// MoqErrno_genType_Temporary_anyParams isolates the any params functions of
// the Errno_genType type
type MoqErrno_genType_Temporary_anyParams struct {
	Recorder *MoqErrno_genType_Temporary_fnRecorder
}

// MoqErrno_genType_Timeout_params holds the params of the Errno_genType type
type MoqErrno_genType_Timeout_params struct{}

// MoqErrno_genType_Timeout_paramsKey holds the map key params of the
// Errno_genType type
type MoqErrno_genType_Timeout_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqErrno_genType_Timeout_resultsByParams contains the results for a given
// set of parameters for the Errno_genType type
type MoqErrno_genType_Timeout_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqErrno_genType_Timeout_paramsKey]*MoqErrno_genType_Timeout_results
}

// MoqErrno_genType_Timeout_doFn defines the type of function needed when
// calling AndDo for the Errno_genType type
type MoqErrno_genType_Timeout_doFn func()

// MoqErrno_genType_Timeout_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Errno_genType type
type MoqErrno_genType_Timeout_doReturnFn func() bool

// MoqErrno_genType_Timeout_results holds the results of the Errno_genType type
type MoqErrno_genType_Timeout_results struct {
	Params  MoqErrno_genType_Timeout_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqErrno_genType_Timeout_doFn
		DoReturnFn MoqErrno_genType_Timeout_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqErrno_genType_Timeout_fnRecorder routes recorded function calls to the
// MoqErrno_genType moq
type MoqErrno_genType_Timeout_fnRecorder struct {
	Params    MoqErrno_genType_Timeout_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqErrno_genType_Timeout_results
	Moq       *MoqErrno_genType
}

// MoqErrno_genType_Timeout_anyParams isolates the any params functions of the
// Errno_genType type
type MoqErrno_genType_Timeout_anyParams struct {
	Recorder *MoqErrno_genType_Timeout_fnRecorder
}

// NewMoqErrno_genType creates a new moq of the Errno_genType type
func NewMoqErrno_genType(scene *moq.Scene, config *moq.Config) *MoqErrno_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqErrno_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqErrno_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Error struct{}
				Is    struct {
					Target moq.ParamIndexing
				}
				Temporary struct{}
				Timeout   struct{}
			}
		}{ParameterIndexing: struct {
			Error struct{}
			Is    struct {
				Target moq.ParamIndexing
			}
			Temporary struct{}
			Timeout   struct{}
		}{
			Error: struct{}{},
			Is: struct {
				Target moq.ParamIndexing
			}{
				Target: moq.ParamIndexByHash,
			},
			Temporary: struct{}{},
			Timeout:   struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Errno_genType type
func (m *MoqErrno_genType) Mock() *MoqErrno_genType_mock { return m.Moq }

func (m *MoqErrno_genType_mock) Error() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqErrno_genType_Error_params{}
	var results *MoqErrno_genType_Error_results
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

func (m *MoqErrno_genType_mock) Is(target error) (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqErrno_genType_Is_params{
		Target: target,
	}
	var results *MoqErrno_genType_Is_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Is {
		paramsKey := m.Moq.ParamsKey_Is(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Is(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Is(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Is(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(target)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(target)
	}
	return
}

func (m *MoqErrno_genType_mock) Temporary() (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqErrno_genType_Temporary_params{}
	var results *MoqErrno_genType_Temporary_results
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

func (m *MoqErrno_genType_mock) Timeout() (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqErrno_genType_Timeout_params{}
	var results *MoqErrno_genType_Timeout_results
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

// OnCall returns the recorder implementation of the Errno_genType type
func (m *MoqErrno_genType) OnCall() *MoqErrno_genType_recorder {
	return &MoqErrno_genType_recorder{
		Moq: m,
	}
}

func (m *MoqErrno_genType_recorder) Error() *MoqErrno_genType_Error_fnRecorder {
	return &MoqErrno_genType_Error_fnRecorder{
		Params:   MoqErrno_genType_Error_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqErrno_genType_Error_fnRecorder) Any() *MoqErrno_genType_Error_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	return &MoqErrno_genType_Error_anyParams{Recorder: r}
}

func (r *MoqErrno_genType_Error_fnRecorder) Seq() *MoqErrno_genType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqErrno_genType_Error_fnRecorder) NoSeq() *MoqErrno_genType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqErrno_genType_Error_fnRecorder) ReturnResults(result1 string) *MoqErrno_genType_Error_fnRecorder {
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
		DoFn       MoqErrno_genType_Error_doFn
		DoReturnFn MoqErrno_genType_Error_doReturnFn
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

func (r *MoqErrno_genType_Error_fnRecorder) AndDo(fn MoqErrno_genType_Error_doFn) *MoqErrno_genType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqErrno_genType_Error_fnRecorder) DoReturnResults(fn MoqErrno_genType_Error_doReturnFn) *MoqErrno_genType_Error_fnRecorder {
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
		DoFn       MoqErrno_genType_Error_doFn
		DoReturnFn MoqErrno_genType_Error_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqErrno_genType_Error_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqErrno_genType_Error_resultsByParams
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
		results = &MoqErrno_genType_Error_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqErrno_genType_Error_paramsKey]*MoqErrno_genType_Error_results{},
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
		r.Results = &MoqErrno_genType_Error_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqErrno_genType_Error_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqErrno_genType_Error_fnRecorder {
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
				DoFn       MoqErrno_genType_Error_doFn
				DoReturnFn MoqErrno_genType_Error_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqErrno_genType) PrettyParams_Error(params MoqErrno_genType_Error_params) string {
	return fmt.Sprintf("Error()")
}

func (m *MoqErrno_genType) ParamsKey_Error(params MoqErrno_genType_Error_params, anyParams uint64) MoqErrno_genType_Error_paramsKey {
	m.Scene.T.Helper()
	return MoqErrno_genType_Error_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqErrno_genType_recorder) Is(target error) *MoqErrno_genType_Is_fnRecorder {
	return &MoqErrno_genType_Is_fnRecorder{
		Params: MoqErrno_genType_Is_params{
			Target: target,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqErrno_genType_Is_fnRecorder) Any() *MoqErrno_genType_Is_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Is(r.Params))
		return nil
	}
	return &MoqErrno_genType_Is_anyParams{Recorder: r}
}

func (a *MoqErrno_genType_Is_anyParams) Target() *MoqErrno_genType_Is_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqErrno_genType_Is_fnRecorder) Seq() *MoqErrno_genType_Is_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Is(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqErrno_genType_Is_fnRecorder) NoSeq() *MoqErrno_genType_Is_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Is(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqErrno_genType_Is_fnRecorder) ReturnResults(result1 bool) *MoqErrno_genType_Is_fnRecorder {
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
		DoFn       MoqErrno_genType_Is_doFn
		DoReturnFn MoqErrno_genType_Is_doReturnFn
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

func (r *MoqErrno_genType_Is_fnRecorder) AndDo(fn MoqErrno_genType_Is_doFn) *MoqErrno_genType_Is_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqErrno_genType_Is_fnRecorder) DoReturnResults(fn MoqErrno_genType_Is_doReturnFn) *MoqErrno_genType_Is_fnRecorder {
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
		DoFn       MoqErrno_genType_Is_doFn
		DoReturnFn MoqErrno_genType_Is_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqErrno_genType_Is_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqErrno_genType_Is_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Is {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqErrno_genType_Is_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqErrno_genType_Is_paramsKey]*MoqErrno_genType_Is_results{},
		}
		r.Moq.ResultsByParams_Is = append(r.Moq.ResultsByParams_Is, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Is) {
			copy(r.Moq.ResultsByParams_Is[insertAt+1:], r.Moq.ResultsByParams_Is[insertAt:0])
			r.Moq.ResultsByParams_Is[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Is(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqErrno_genType_Is_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqErrno_genType_Is_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqErrno_genType_Is_fnRecorder {
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
				DoFn       MoqErrno_genType_Is_doFn
				DoReturnFn MoqErrno_genType_Is_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqErrno_genType) PrettyParams_Is(params MoqErrno_genType_Is_params) string {
	return fmt.Sprintf("Is(%#v)", params.Target)
}

func (m *MoqErrno_genType) ParamsKey_Is(params MoqErrno_genType_Is_params, anyParams uint64) MoqErrno_genType_Is_paramsKey {
	m.Scene.T.Helper()
	var targetUsed error
	var targetUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Is.Target == moq.ParamIndexByValue {
			targetUsed = params.Target
		} else {
			targetUsedHash = hash.DeepHash(params.Target)
		}
	}
	return MoqErrno_genType_Is_paramsKey{
		Params: struct{ Target error }{
			Target: targetUsed,
		},
		Hashes: struct{ Target hash.Hash }{
			Target: targetUsedHash,
		},
	}
}

func (m *MoqErrno_genType_recorder) Temporary() *MoqErrno_genType_Temporary_fnRecorder {
	return &MoqErrno_genType_Temporary_fnRecorder{
		Params:   MoqErrno_genType_Temporary_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqErrno_genType_Temporary_fnRecorder) Any() *MoqErrno_genType_Temporary_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Temporary(r.Params))
		return nil
	}
	return &MoqErrno_genType_Temporary_anyParams{Recorder: r}
}

func (r *MoqErrno_genType_Temporary_fnRecorder) Seq() *MoqErrno_genType_Temporary_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Temporary(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqErrno_genType_Temporary_fnRecorder) NoSeq() *MoqErrno_genType_Temporary_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Temporary(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqErrno_genType_Temporary_fnRecorder) ReturnResults(result1 bool) *MoqErrno_genType_Temporary_fnRecorder {
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
		DoFn       MoqErrno_genType_Temporary_doFn
		DoReturnFn MoqErrno_genType_Temporary_doReturnFn
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

func (r *MoqErrno_genType_Temporary_fnRecorder) AndDo(fn MoqErrno_genType_Temporary_doFn) *MoqErrno_genType_Temporary_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqErrno_genType_Temporary_fnRecorder) DoReturnResults(fn MoqErrno_genType_Temporary_doReturnFn) *MoqErrno_genType_Temporary_fnRecorder {
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
		DoFn       MoqErrno_genType_Temporary_doFn
		DoReturnFn MoqErrno_genType_Temporary_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqErrno_genType_Temporary_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqErrno_genType_Temporary_resultsByParams
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
		results = &MoqErrno_genType_Temporary_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqErrno_genType_Temporary_paramsKey]*MoqErrno_genType_Temporary_results{},
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
		r.Results = &MoqErrno_genType_Temporary_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqErrno_genType_Temporary_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqErrno_genType_Temporary_fnRecorder {
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
				DoFn       MoqErrno_genType_Temporary_doFn
				DoReturnFn MoqErrno_genType_Temporary_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqErrno_genType) PrettyParams_Temporary(params MoqErrno_genType_Temporary_params) string {
	return fmt.Sprintf("Temporary()")
}

func (m *MoqErrno_genType) ParamsKey_Temporary(params MoqErrno_genType_Temporary_params, anyParams uint64) MoqErrno_genType_Temporary_paramsKey {
	m.Scene.T.Helper()
	return MoqErrno_genType_Temporary_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqErrno_genType_recorder) Timeout() *MoqErrno_genType_Timeout_fnRecorder {
	return &MoqErrno_genType_Timeout_fnRecorder{
		Params:   MoqErrno_genType_Timeout_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqErrno_genType_Timeout_fnRecorder) Any() *MoqErrno_genType_Timeout_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Timeout(r.Params))
		return nil
	}
	return &MoqErrno_genType_Timeout_anyParams{Recorder: r}
}

func (r *MoqErrno_genType_Timeout_fnRecorder) Seq() *MoqErrno_genType_Timeout_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Timeout(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqErrno_genType_Timeout_fnRecorder) NoSeq() *MoqErrno_genType_Timeout_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Timeout(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqErrno_genType_Timeout_fnRecorder) ReturnResults(result1 bool) *MoqErrno_genType_Timeout_fnRecorder {
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
		DoFn       MoqErrno_genType_Timeout_doFn
		DoReturnFn MoqErrno_genType_Timeout_doReturnFn
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

func (r *MoqErrno_genType_Timeout_fnRecorder) AndDo(fn MoqErrno_genType_Timeout_doFn) *MoqErrno_genType_Timeout_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqErrno_genType_Timeout_fnRecorder) DoReturnResults(fn MoqErrno_genType_Timeout_doReturnFn) *MoqErrno_genType_Timeout_fnRecorder {
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
		DoFn       MoqErrno_genType_Timeout_doFn
		DoReturnFn MoqErrno_genType_Timeout_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqErrno_genType_Timeout_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqErrno_genType_Timeout_resultsByParams
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
		results = &MoqErrno_genType_Timeout_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqErrno_genType_Timeout_paramsKey]*MoqErrno_genType_Timeout_results{},
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
		r.Results = &MoqErrno_genType_Timeout_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqErrno_genType_Timeout_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqErrno_genType_Timeout_fnRecorder {
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
				DoFn       MoqErrno_genType_Timeout_doFn
				DoReturnFn MoqErrno_genType_Timeout_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqErrno_genType) PrettyParams_Timeout(params MoqErrno_genType_Timeout_params) string {
	return fmt.Sprintf("Timeout()")
}

func (m *MoqErrno_genType) ParamsKey_Timeout(params MoqErrno_genType_Timeout_params, anyParams uint64) MoqErrno_genType_Timeout_paramsKey {
	m.Scene.T.Helper()
	return MoqErrno_genType_Timeout_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqErrno_genType) Reset() {
	m.ResultsByParams_Error = nil
	m.ResultsByParams_Is = nil
	m.ResultsByParams_Temporary = nil
	m.ResultsByParams_Timeout = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqErrno_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Error {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Error(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Is {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Is(results.Params))
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
	for _, res := range m.ResultsByParams_Timeout {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Timeout(results.Params))
			}
		}
	}
}
// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package context

import (
	"context"
	"fmt"
	"math/bits"
	"sync/atomic"
	"time"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that context.Context is mocked
// completely
var _ context.Context = (*MoqContext_mock)(nil)

// MoqContext holds the state of a moq of the Context type
type MoqContext struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqContext_mock

	ResultsByParams_Deadline []MoqContext_Deadline_resultsByParams
	ResultsByParams_Done     []MoqContext_Done_resultsByParams
	ResultsByParams_Err      []MoqContext_Err_resultsByParams
	ResultsByParams_Value    []MoqContext_Value_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Deadline struct{}
			Done     struct{}
			Err      struct{}
			Value    struct {
				Key moq.ParamIndexing
			}
		}
	}
	// MoqContext_mock isolates the mock interface of the Context type
}

type MoqContext_mock struct {
	Moq *MoqContext
}

// MoqContext_recorder isolates the recorder interface of the Context type
type MoqContext_recorder struct {
	Moq *MoqContext
}

// MoqContext_Deadline_params holds the params of the Context type
type MoqContext_Deadline_params struct{}

// MoqContext_Deadline_paramsKey holds the map key params of the Context type
type MoqContext_Deadline_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqContext_Deadline_resultsByParams contains the results for a given set of
// parameters for the Context type
type MoqContext_Deadline_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqContext_Deadline_paramsKey]*MoqContext_Deadline_results
}

// MoqContext_Deadline_doFn defines the type of function needed when calling
// AndDo for the Context type
type MoqContext_Deadline_doFn func()

// MoqContext_Deadline_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Context type
type MoqContext_Deadline_doReturnFn func() (deadline time.Time, ok bool)

// MoqContext_Deadline_results holds the results of the Context type
type MoqContext_Deadline_results struct {
	Params  MoqContext_Deadline_params
	Results []struct {
		Values *struct {
			Deadline time.Time
			Ok       bool
		}
		Sequence   uint32
		DoFn       MoqContext_Deadline_doFn
		DoReturnFn MoqContext_Deadline_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqContext_Deadline_fnRecorder routes recorded function calls to the
// MoqContext moq
type MoqContext_Deadline_fnRecorder struct {
	Params    MoqContext_Deadline_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqContext_Deadline_results
	Moq       *MoqContext
}

// MoqContext_Deadline_anyParams isolates the any params functions of the
// Context type
type MoqContext_Deadline_anyParams struct {
	Recorder *MoqContext_Deadline_fnRecorder
}

// MoqContext_Done_params holds the params of the Context type
type MoqContext_Done_params struct{}

// MoqContext_Done_paramsKey holds the map key params of the Context type
type MoqContext_Done_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqContext_Done_resultsByParams contains the results for a given set of
// parameters for the Context type
type MoqContext_Done_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqContext_Done_paramsKey]*MoqContext_Done_results
}

// MoqContext_Done_doFn defines the type of function needed when calling AndDo
// for the Context type
type MoqContext_Done_doFn func()

// MoqContext_Done_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Context type
type MoqContext_Done_doReturnFn func() <-chan struct{}

// MoqContext_Done_results holds the results of the Context type
type MoqContext_Done_results struct {
	Params  MoqContext_Done_params
	Results []struct {
		Values *struct {
			Result1 <-chan struct{}
		}
		Sequence   uint32
		DoFn       MoqContext_Done_doFn
		DoReturnFn MoqContext_Done_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqContext_Done_fnRecorder routes recorded function calls to the MoqContext
// moq
type MoqContext_Done_fnRecorder struct {
	Params    MoqContext_Done_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqContext_Done_results
	Moq       *MoqContext
}

// MoqContext_Done_anyParams isolates the any params functions of the Context
// type
type MoqContext_Done_anyParams struct {
	Recorder *MoqContext_Done_fnRecorder
}

// MoqContext_Err_params holds the params of the Context type
type MoqContext_Err_params struct{}

// MoqContext_Err_paramsKey holds the map key params of the Context type
type MoqContext_Err_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqContext_Err_resultsByParams contains the results for a given set of
// parameters for the Context type
type MoqContext_Err_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqContext_Err_paramsKey]*MoqContext_Err_results
}

// MoqContext_Err_doFn defines the type of function needed when calling AndDo
// for the Context type
type MoqContext_Err_doFn func()

// MoqContext_Err_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Context type
type MoqContext_Err_doReturnFn func() error

// MoqContext_Err_results holds the results of the Context type
type MoqContext_Err_results struct {
	Params  MoqContext_Err_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqContext_Err_doFn
		DoReturnFn MoqContext_Err_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqContext_Err_fnRecorder routes recorded function calls to the MoqContext
// moq
type MoqContext_Err_fnRecorder struct {
	Params    MoqContext_Err_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqContext_Err_results
	Moq       *MoqContext
}

// MoqContext_Err_anyParams isolates the any params functions of the Context
// type
type MoqContext_Err_anyParams struct {
	Recorder *MoqContext_Err_fnRecorder
}

// MoqContext_Value_params holds the params of the Context type
type MoqContext_Value_params struct{ Key any }

// MoqContext_Value_paramsKey holds the map key params of the Context type
type MoqContext_Value_paramsKey struct {
	Params struct{ Key any }
	Hashes struct{ Key hash.Hash }
}

// MoqContext_Value_resultsByParams contains the results for a given set of
// parameters for the Context type
type MoqContext_Value_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqContext_Value_paramsKey]*MoqContext_Value_results
}

// MoqContext_Value_doFn defines the type of function needed when calling AndDo
// for the Context type
type MoqContext_Value_doFn func(key any)

// MoqContext_Value_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Context type
type MoqContext_Value_doReturnFn func(key any) any

// MoqContext_Value_results holds the results of the Context type
type MoqContext_Value_results struct {
	Params  MoqContext_Value_params
	Results []struct {
		Values *struct {
			Result1 any
		}
		Sequence   uint32
		DoFn       MoqContext_Value_doFn
		DoReturnFn MoqContext_Value_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqContext_Value_fnRecorder routes recorded function calls to the MoqContext
// moq
type MoqContext_Value_fnRecorder struct {
	Params    MoqContext_Value_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqContext_Value_results
	Moq       *MoqContext
}

// MoqContext_Value_anyParams isolates the any params functions of the Context
// type
type MoqContext_Value_anyParams struct {
	Recorder *MoqContext_Value_fnRecorder
}

// NewMoqContext creates a new moq of the Context type
func NewMoqContext(scene *moq.Scene, config *moq.Config) *MoqContext {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqContext{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqContext_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Deadline struct{}
				Done     struct{}
				Err      struct{}
				Value    struct {
					Key moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Deadline struct{}
			Done     struct{}
			Err      struct{}
			Value    struct {
				Key moq.ParamIndexing
			}
		}{
			Deadline: struct{}{},
			Done:     struct{}{},
			Err:      struct{}{},
			Value: struct {
				Key moq.ParamIndexing
			}{
				Key: moq.ParamIndexByValue,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Context type
func (m *MoqContext) Mock() *MoqContext_mock { return m.Moq }

func (m *MoqContext_mock) Deadline() (deadline time.Time, ok bool) {
	m.Moq.Scene.T.Helper()
	params := MoqContext_Deadline_params{}
	var results *MoqContext_Deadline_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Deadline {
		paramsKey := m.Moq.ParamsKey_Deadline(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Deadline(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Deadline(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Deadline(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		deadline = result.Values.Deadline
		ok = result.Values.Ok
	}
	if result.DoReturnFn != nil {
		deadline, ok = result.DoReturnFn()
	}
	return
}

func (m *MoqContext_mock) Done() (result1 <-chan struct{}) {
	m.Moq.Scene.T.Helper()
	params := MoqContext_Done_params{}
	var results *MoqContext_Done_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Done {
		paramsKey := m.Moq.ParamsKey_Done(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Done(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Done(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Done(params))
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

func (m *MoqContext_mock) Err() (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqContext_Err_params{}
	var results *MoqContext_Err_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Err {
		paramsKey := m.Moq.ParamsKey_Err(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Err(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Err(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Err(params))
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

func (m *MoqContext_mock) Value(key any) (result1 any) {
	m.Moq.Scene.T.Helper()
	params := MoqContext_Value_params{
		Key: key,
	}
	var results *MoqContext_Value_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Value {
		paramsKey := m.Moq.ParamsKey_Value(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Value(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Value(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Value(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(key)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(key)
	}
	return
}

// OnCall returns the recorder implementation of the Context type
func (m *MoqContext) OnCall() *MoqContext_recorder {
	return &MoqContext_recorder{
		Moq: m,
	}
}

func (m *MoqContext_recorder) Deadline() *MoqContext_Deadline_fnRecorder {
	return &MoqContext_Deadline_fnRecorder{
		Params:   MoqContext_Deadline_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqContext_Deadline_fnRecorder) Any() *MoqContext_Deadline_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Deadline(r.Params))
		return nil
	}
	return &MoqContext_Deadline_anyParams{Recorder: r}
}

func (r *MoqContext_Deadline_fnRecorder) Seq() *MoqContext_Deadline_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Deadline(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqContext_Deadline_fnRecorder) NoSeq() *MoqContext_Deadline_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Deadline(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqContext_Deadline_fnRecorder) ReturnResults(deadline time.Time, ok bool) *MoqContext_Deadline_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Deadline time.Time
			Ok       bool
		}
		Sequence   uint32
		DoFn       MoqContext_Deadline_doFn
		DoReturnFn MoqContext_Deadline_doReturnFn
	}{
		Values: &struct {
			Deadline time.Time
			Ok       bool
		}{
			Deadline: deadline,
			Ok:       ok,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqContext_Deadline_fnRecorder) AndDo(fn MoqContext_Deadline_doFn) *MoqContext_Deadline_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqContext_Deadline_fnRecorder) DoReturnResults(fn MoqContext_Deadline_doReturnFn) *MoqContext_Deadline_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Deadline time.Time
			Ok       bool
		}
		Sequence   uint32
		DoFn       MoqContext_Deadline_doFn
		DoReturnFn MoqContext_Deadline_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqContext_Deadline_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqContext_Deadline_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Deadline {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqContext_Deadline_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqContext_Deadline_paramsKey]*MoqContext_Deadline_results{},
		}
		r.Moq.ResultsByParams_Deadline = append(r.Moq.ResultsByParams_Deadline, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Deadline) {
			copy(r.Moq.ResultsByParams_Deadline[insertAt+1:], r.Moq.ResultsByParams_Deadline[insertAt:0])
			r.Moq.ResultsByParams_Deadline[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Deadline(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqContext_Deadline_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqContext_Deadline_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqContext_Deadline_fnRecorder {
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
					Deadline time.Time
					Ok       bool
				}
				Sequence   uint32
				DoFn       MoqContext_Deadline_doFn
				DoReturnFn MoqContext_Deadline_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqContext) PrettyParams_Deadline(params MoqContext_Deadline_params) string {
	return fmt.Sprintf("Deadline()")
}

func (m *MoqContext) ParamsKey_Deadline(params MoqContext_Deadline_params, anyParams uint64) MoqContext_Deadline_paramsKey {
	m.Scene.T.Helper()
	return MoqContext_Deadline_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqContext_recorder) Done() *MoqContext_Done_fnRecorder {
	return &MoqContext_Done_fnRecorder{
		Params:   MoqContext_Done_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqContext_Done_fnRecorder) Any() *MoqContext_Done_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Done(r.Params))
		return nil
	}
	return &MoqContext_Done_anyParams{Recorder: r}
}

func (r *MoqContext_Done_fnRecorder) Seq() *MoqContext_Done_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Done(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqContext_Done_fnRecorder) NoSeq() *MoqContext_Done_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Done(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqContext_Done_fnRecorder) ReturnResults(result1 <-chan struct{}) *MoqContext_Done_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 <-chan struct{}
		}
		Sequence   uint32
		DoFn       MoqContext_Done_doFn
		DoReturnFn MoqContext_Done_doReturnFn
	}{
		Values: &struct {
			Result1 <-chan struct{}
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqContext_Done_fnRecorder) AndDo(fn MoqContext_Done_doFn) *MoqContext_Done_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqContext_Done_fnRecorder) DoReturnResults(fn MoqContext_Done_doReturnFn) *MoqContext_Done_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 <-chan struct{}
		}
		Sequence   uint32
		DoFn       MoqContext_Done_doFn
		DoReturnFn MoqContext_Done_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqContext_Done_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqContext_Done_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Done {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqContext_Done_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqContext_Done_paramsKey]*MoqContext_Done_results{},
		}
		r.Moq.ResultsByParams_Done = append(r.Moq.ResultsByParams_Done, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Done) {
			copy(r.Moq.ResultsByParams_Done[insertAt+1:], r.Moq.ResultsByParams_Done[insertAt:0])
			r.Moq.ResultsByParams_Done[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Done(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqContext_Done_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqContext_Done_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqContext_Done_fnRecorder {
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
					Result1 <-chan struct{}
				}
				Sequence   uint32
				DoFn       MoqContext_Done_doFn
				DoReturnFn MoqContext_Done_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqContext) PrettyParams_Done(params MoqContext_Done_params) string {
	return fmt.Sprintf("Done()")
}

func (m *MoqContext) ParamsKey_Done(params MoqContext_Done_params, anyParams uint64) MoqContext_Done_paramsKey {
	m.Scene.T.Helper()
	return MoqContext_Done_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqContext_recorder) Err() *MoqContext_Err_fnRecorder {
	return &MoqContext_Err_fnRecorder{
		Params:   MoqContext_Err_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqContext_Err_fnRecorder) Any() *MoqContext_Err_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Err(r.Params))
		return nil
	}
	return &MoqContext_Err_anyParams{Recorder: r}
}

func (r *MoqContext_Err_fnRecorder) Seq() *MoqContext_Err_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Err(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqContext_Err_fnRecorder) NoSeq() *MoqContext_Err_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Err(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqContext_Err_fnRecorder) ReturnResults(result1 error) *MoqContext_Err_fnRecorder {
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
		DoFn       MoqContext_Err_doFn
		DoReturnFn MoqContext_Err_doReturnFn
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

func (r *MoqContext_Err_fnRecorder) AndDo(fn MoqContext_Err_doFn) *MoqContext_Err_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqContext_Err_fnRecorder) DoReturnResults(fn MoqContext_Err_doReturnFn) *MoqContext_Err_fnRecorder {
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
		DoFn       MoqContext_Err_doFn
		DoReturnFn MoqContext_Err_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqContext_Err_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqContext_Err_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Err {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqContext_Err_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqContext_Err_paramsKey]*MoqContext_Err_results{},
		}
		r.Moq.ResultsByParams_Err = append(r.Moq.ResultsByParams_Err, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Err) {
			copy(r.Moq.ResultsByParams_Err[insertAt+1:], r.Moq.ResultsByParams_Err[insertAt:0])
			r.Moq.ResultsByParams_Err[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Err(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqContext_Err_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqContext_Err_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqContext_Err_fnRecorder {
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
				DoFn       MoqContext_Err_doFn
				DoReturnFn MoqContext_Err_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqContext) PrettyParams_Err(params MoqContext_Err_params) string {
	return fmt.Sprintf("Err()")
}

func (m *MoqContext) ParamsKey_Err(params MoqContext_Err_params, anyParams uint64) MoqContext_Err_paramsKey {
	m.Scene.T.Helper()
	return MoqContext_Err_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqContext_recorder) Value(key any) *MoqContext_Value_fnRecorder {
	return &MoqContext_Value_fnRecorder{
		Params: MoqContext_Value_params{
			Key: key,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqContext_Value_fnRecorder) Any() *MoqContext_Value_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Value(r.Params))
		return nil
	}
	return &MoqContext_Value_anyParams{Recorder: r}
}

func (a *MoqContext_Value_anyParams) Key() *MoqContext_Value_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqContext_Value_fnRecorder) Seq() *MoqContext_Value_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Value(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqContext_Value_fnRecorder) NoSeq() *MoqContext_Value_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Value(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqContext_Value_fnRecorder) ReturnResults(result1 any) *MoqContext_Value_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 any
		}
		Sequence   uint32
		DoFn       MoqContext_Value_doFn
		DoReturnFn MoqContext_Value_doReturnFn
	}{
		Values: &struct {
			Result1 any
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqContext_Value_fnRecorder) AndDo(fn MoqContext_Value_doFn) *MoqContext_Value_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqContext_Value_fnRecorder) DoReturnResults(fn MoqContext_Value_doReturnFn) *MoqContext_Value_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 any
		}
		Sequence   uint32
		DoFn       MoqContext_Value_doFn
		DoReturnFn MoqContext_Value_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqContext_Value_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqContext_Value_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Value {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqContext_Value_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqContext_Value_paramsKey]*MoqContext_Value_results{},
		}
		r.Moq.ResultsByParams_Value = append(r.Moq.ResultsByParams_Value, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Value) {
			copy(r.Moq.ResultsByParams_Value[insertAt+1:], r.Moq.ResultsByParams_Value[insertAt:0])
			r.Moq.ResultsByParams_Value[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Value(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqContext_Value_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqContext_Value_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqContext_Value_fnRecorder {
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
					Result1 any
				}
				Sequence   uint32
				DoFn       MoqContext_Value_doFn
				DoReturnFn MoqContext_Value_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqContext) PrettyParams_Value(params MoqContext_Value_params) string {
	return fmt.Sprintf("Value(%#v)", params.Key)
}

func (m *MoqContext) ParamsKey_Value(params MoqContext_Value_params, anyParams uint64) MoqContext_Value_paramsKey {
	m.Scene.T.Helper()
	var keyUsed any
	var keyUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Value.Key == moq.ParamIndexByValue {
			keyUsed = params.Key
		} else {
			keyUsedHash = hash.DeepHash(params.Key)
		}
	}
	return MoqContext_Value_paramsKey{
		Params: struct{ Key any }{
			Key: keyUsed,
		},
		Hashes: struct{ Key hash.Hash }{
			Key: keyUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqContext) Reset() {
	m.ResultsByParams_Deadline = nil
	m.ResultsByParams_Done = nil
	m.ResultsByParams_Err = nil
	m.ResultsByParams_Value = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqContext) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Deadline {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Deadline(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Done {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Done(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Err {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Err(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Value {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Value(results.Params))
			}
		}
	}
}

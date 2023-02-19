// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package sync

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that sync.Cond_starGenType is mocked
// completely
var _ Cond_starGenType = (*MoqCond_starGenType_mock)(nil)

// Cond_starGenType is the fabricated implementation type of this mock (emitted
// when mocking a collections of methods directly and not from an interface
// type)
type Cond_starGenType interface {
	Wait()
	Signal()
	Broadcast()
}

// MoqCond_starGenType holds the state of a moq of the Cond_starGenType type
type MoqCond_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqCond_starGenType_mock

	ResultsByParams_Wait      []MoqCond_starGenType_Wait_resultsByParams
	ResultsByParams_Signal    []MoqCond_starGenType_Signal_resultsByParams
	ResultsByParams_Broadcast []MoqCond_starGenType_Broadcast_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Wait      struct{}
			Signal    struct{}
			Broadcast struct{}
		}
	}
}

// MoqCond_starGenType_mock isolates the mock interface of the Cond_starGenType
// type
type MoqCond_starGenType_mock struct {
	Moq *MoqCond_starGenType
}

// MoqCond_starGenType_recorder isolates the recorder interface of the
// Cond_starGenType type
type MoqCond_starGenType_recorder struct {
	Moq *MoqCond_starGenType
}

// MoqCond_starGenType_Wait_params holds the params of the Cond_starGenType
// type
type MoqCond_starGenType_Wait_params struct{}

// MoqCond_starGenType_Wait_paramsKey holds the map key params of the
// Cond_starGenType type
type MoqCond_starGenType_Wait_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqCond_starGenType_Wait_resultsByParams contains the results for a given
// set of parameters for the Cond_starGenType type
type MoqCond_starGenType_Wait_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqCond_starGenType_Wait_paramsKey]*MoqCond_starGenType_Wait_results
}

// MoqCond_starGenType_Wait_doFn defines the type of function needed when
// calling AndDo for the Cond_starGenType type
type MoqCond_starGenType_Wait_doFn func()

// MoqCond_starGenType_Wait_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Cond_starGenType type
type MoqCond_starGenType_Wait_doReturnFn func()

// MoqCond_starGenType_Wait_results holds the results of the Cond_starGenType
// type
type MoqCond_starGenType_Wait_results struct {
	Params  MoqCond_starGenType_Wait_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqCond_starGenType_Wait_doFn
		DoReturnFn MoqCond_starGenType_Wait_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqCond_starGenType_Wait_fnRecorder routes recorded function calls to the
// MoqCond_starGenType moq
type MoqCond_starGenType_Wait_fnRecorder struct {
	Params    MoqCond_starGenType_Wait_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqCond_starGenType_Wait_results
	Moq       *MoqCond_starGenType
}

// MoqCond_starGenType_Wait_anyParams isolates the any params functions of the
// Cond_starGenType type
type MoqCond_starGenType_Wait_anyParams struct {
	Recorder *MoqCond_starGenType_Wait_fnRecorder
}

// MoqCond_starGenType_Signal_params holds the params of the Cond_starGenType
// type
type MoqCond_starGenType_Signal_params struct{}

// MoqCond_starGenType_Signal_paramsKey holds the map key params of the
// Cond_starGenType type
type MoqCond_starGenType_Signal_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqCond_starGenType_Signal_resultsByParams contains the results for a given
// set of parameters for the Cond_starGenType type
type MoqCond_starGenType_Signal_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqCond_starGenType_Signal_paramsKey]*MoqCond_starGenType_Signal_results
}

// MoqCond_starGenType_Signal_doFn defines the type of function needed when
// calling AndDo for the Cond_starGenType type
type MoqCond_starGenType_Signal_doFn func()

// MoqCond_starGenType_Signal_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Cond_starGenType type
type MoqCond_starGenType_Signal_doReturnFn func()

// MoqCond_starGenType_Signal_results holds the results of the Cond_starGenType
// type
type MoqCond_starGenType_Signal_results struct {
	Params  MoqCond_starGenType_Signal_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqCond_starGenType_Signal_doFn
		DoReturnFn MoqCond_starGenType_Signal_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqCond_starGenType_Signal_fnRecorder routes recorded function calls to the
// MoqCond_starGenType moq
type MoqCond_starGenType_Signal_fnRecorder struct {
	Params    MoqCond_starGenType_Signal_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqCond_starGenType_Signal_results
	Moq       *MoqCond_starGenType
}

// MoqCond_starGenType_Signal_anyParams isolates the any params functions of
// the Cond_starGenType type
type MoqCond_starGenType_Signal_anyParams struct {
	Recorder *MoqCond_starGenType_Signal_fnRecorder
}

// MoqCond_starGenType_Broadcast_params holds the params of the
// Cond_starGenType type
type MoqCond_starGenType_Broadcast_params struct{}

// MoqCond_starGenType_Broadcast_paramsKey holds the map key params of the
// Cond_starGenType type
type MoqCond_starGenType_Broadcast_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqCond_starGenType_Broadcast_resultsByParams contains the results for a
// given set of parameters for the Cond_starGenType type
type MoqCond_starGenType_Broadcast_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqCond_starGenType_Broadcast_paramsKey]*MoqCond_starGenType_Broadcast_results
}

// MoqCond_starGenType_Broadcast_doFn defines the type of function needed when
// calling AndDo for the Cond_starGenType type
type MoqCond_starGenType_Broadcast_doFn func()

// MoqCond_starGenType_Broadcast_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Cond_starGenType type
type MoqCond_starGenType_Broadcast_doReturnFn func()

// MoqCond_starGenType_Broadcast_results holds the results of the
// Cond_starGenType type
type MoqCond_starGenType_Broadcast_results struct {
	Params  MoqCond_starGenType_Broadcast_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqCond_starGenType_Broadcast_doFn
		DoReturnFn MoqCond_starGenType_Broadcast_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqCond_starGenType_Broadcast_fnRecorder routes recorded function calls to
// the MoqCond_starGenType moq
type MoqCond_starGenType_Broadcast_fnRecorder struct {
	Params    MoqCond_starGenType_Broadcast_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqCond_starGenType_Broadcast_results
	Moq       *MoqCond_starGenType
}

// MoqCond_starGenType_Broadcast_anyParams isolates the any params functions of
// the Cond_starGenType type
type MoqCond_starGenType_Broadcast_anyParams struct {
	Recorder *MoqCond_starGenType_Broadcast_fnRecorder
}

// NewMoqCond_starGenType creates a new moq of the Cond_starGenType type
func NewMoqCond_starGenType(scene *moq.Scene, config *moq.Config) *MoqCond_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqCond_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqCond_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Wait      struct{}
				Signal    struct{}
				Broadcast struct{}
			}
		}{ParameterIndexing: struct {
			Wait      struct{}
			Signal    struct{}
			Broadcast struct{}
		}{
			Wait:      struct{}{},
			Signal:    struct{}{},
			Broadcast: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Cond_starGenType type
func (m *MoqCond_starGenType) Mock() *MoqCond_starGenType_mock { return m.Moq }

func (m *MoqCond_starGenType_mock) Wait() {
	m.Moq.Scene.T.Helper()
	params := MoqCond_starGenType_Wait_params{}
	var results *MoqCond_starGenType_Wait_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Wait {
		paramsKey := m.Moq.ParamsKey_Wait(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Wait(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Wait(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Wait(params))
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

func (m *MoqCond_starGenType_mock) Signal() {
	m.Moq.Scene.T.Helper()
	params := MoqCond_starGenType_Signal_params{}
	var results *MoqCond_starGenType_Signal_results
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

func (m *MoqCond_starGenType_mock) Broadcast() {
	m.Moq.Scene.T.Helper()
	params := MoqCond_starGenType_Broadcast_params{}
	var results *MoqCond_starGenType_Broadcast_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Broadcast {
		paramsKey := m.Moq.ParamsKey_Broadcast(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Broadcast(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Broadcast(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Broadcast(params))
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

// OnCall returns the recorder implementation of the Cond_starGenType type
func (m *MoqCond_starGenType) OnCall() *MoqCond_starGenType_recorder {
	return &MoqCond_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqCond_starGenType_recorder) Wait() *MoqCond_starGenType_Wait_fnRecorder {
	return &MoqCond_starGenType_Wait_fnRecorder{
		Params:   MoqCond_starGenType_Wait_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqCond_starGenType_Wait_fnRecorder) Any() *MoqCond_starGenType_Wait_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Wait(r.Params))
		return nil
	}
	return &MoqCond_starGenType_Wait_anyParams{Recorder: r}
}

func (r *MoqCond_starGenType_Wait_fnRecorder) Seq() *MoqCond_starGenType_Wait_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Wait(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqCond_starGenType_Wait_fnRecorder) NoSeq() *MoqCond_starGenType_Wait_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Wait(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqCond_starGenType_Wait_fnRecorder) ReturnResults() *MoqCond_starGenType_Wait_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqCond_starGenType_Wait_doFn
		DoReturnFn MoqCond_starGenType_Wait_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqCond_starGenType_Wait_fnRecorder) AndDo(fn MoqCond_starGenType_Wait_doFn) *MoqCond_starGenType_Wait_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqCond_starGenType_Wait_fnRecorder) DoReturnResults(fn MoqCond_starGenType_Wait_doReturnFn) *MoqCond_starGenType_Wait_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqCond_starGenType_Wait_doFn
		DoReturnFn MoqCond_starGenType_Wait_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqCond_starGenType_Wait_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqCond_starGenType_Wait_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Wait {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqCond_starGenType_Wait_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqCond_starGenType_Wait_paramsKey]*MoqCond_starGenType_Wait_results{},
		}
		r.Moq.ResultsByParams_Wait = append(r.Moq.ResultsByParams_Wait, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Wait) {
			copy(r.Moq.ResultsByParams_Wait[insertAt+1:], r.Moq.ResultsByParams_Wait[insertAt:0])
			r.Moq.ResultsByParams_Wait[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Wait(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqCond_starGenType_Wait_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqCond_starGenType_Wait_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqCond_starGenType_Wait_fnRecorder {
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
				DoFn       MoqCond_starGenType_Wait_doFn
				DoReturnFn MoqCond_starGenType_Wait_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqCond_starGenType) PrettyParams_Wait(params MoqCond_starGenType_Wait_params) string {
	return fmt.Sprintf("Wait()")
}

func (m *MoqCond_starGenType) ParamsKey_Wait(params MoqCond_starGenType_Wait_params, anyParams uint64) MoqCond_starGenType_Wait_paramsKey {
	m.Scene.T.Helper()
	return MoqCond_starGenType_Wait_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqCond_starGenType_recorder) Signal() *MoqCond_starGenType_Signal_fnRecorder {
	return &MoqCond_starGenType_Signal_fnRecorder{
		Params:   MoqCond_starGenType_Signal_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqCond_starGenType_Signal_fnRecorder) Any() *MoqCond_starGenType_Signal_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Signal(r.Params))
		return nil
	}
	return &MoqCond_starGenType_Signal_anyParams{Recorder: r}
}

func (r *MoqCond_starGenType_Signal_fnRecorder) Seq() *MoqCond_starGenType_Signal_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Signal(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqCond_starGenType_Signal_fnRecorder) NoSeq() *MoqCond_starGenType_Signal_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Signal(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqCond_starGenType_Signal_fnRecorder) ReturnResults() *MoqCond_starGenType_Signal_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqCond_starGenType_Signal_doFn
		DoReturnFn MoqCond_starGenType_Signal_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqCond_starGenType_Signal_fnRecorder) AndDo(fn MoqCond_starGenType_Signal_doFn) *MoqCond_starGenType_Signal_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqCond_starGenType_Signal_fnRecorder) DoReturnResults(fn MoqCond_starGenType_Signal_doReturnFn) *MoqCond_starGenType_Signal_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqCond_starGenType_Signal_doFn
		DoReturnFn MoqCond_starGenType_Signal_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqCond_starGenType_Signal_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqCond_starGenType_Signal_resultsByParams
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
		results = &MoqCond_starGenType_Signal_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqCond_starGenType_Signal_paramsKey]*MoqCond_starGenType_Signal_results{},
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
		r.Results = &MoqCond_starGenType_Signal_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqCond_starGenType_Signal_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqCond_starGenType_Signal_fnRecorder {
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
				DoFn       MoqCond_starGenType_Signal_doFn
				DoReturnFn MoqCond_starGenType_Signal_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqCond_starGenType) PrettyParams_Signal(params MoqCond_starGenType_Signal_params) string {
	return fmt.Sprintf("Signal()")
}

func (m *MoqCond_starGenType) ParamsKey_Signal(params MoqCond_starGenType_Signal_params, anyParams uint64) MoqCond_starGenType_Signal_paramsKey {
	m.Scene.T.Helper()
	return MoqCond_starGenType_Signal_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqCond_starGenType_recorder) Broadcast() *MoqCond_starGenType_Broadcast_fnRecorder {
	return &MoqCond_starGenType_Broadcast_fnRecorder{
		Params:   MoqCond_starGenType_Broadcast_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqCond_starGenType_Broadcast_fnRecorder) Any() *MoqCond_starGenType_Broadcast_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Broadcast(r.Params))
		return nil
	}
	return &MoqCond_starGenType_Broadcast_anyParams{Recorder: r}
}

func (r *MoqCond_starGenType_Broadcast_fnRecorder) Seq() *MoqCond_starGenType_Broadcast_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Broadcast(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqCond_starGenType_Broadcast_fnRecorder) NoSeq() *MoqCond_starGenType_Broadcast_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Broadcast(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqCond_starGenType_Broadcast_fnRecorder) ReturnResults() *MoqCond_starGenType_Broadcast_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqCond_starGenType_Broadcast_doFn
		DoReturnFn MoqCond_starGenType_Broadcast_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqCond_starGenType_Broadcast_fnRecorder) AndDo(fn MoqCond_starGenType_Broadcast_doFn) *MoqCond_starGenType_Broadcast_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqCond_starGenType_Broadcast_fnRecorder) DoReturnResults(fn MoqCond_starGenType_Broadcast_doReturnFn) *MoqCond_starGenType_Broadcast_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqCond_starGenType_Broadcast_doFn
		DoReturnFn MoqCond_starGenType_Broadcast_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqCond_starGenType_Broadcast_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqCond_starGenType_Broadcast_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Broadcast {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqCond_starGenType_Broadcast_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqCond_starGenType_Broadcast_paramsKey]*MoqCond_starGenType_Broadcast_results{},
		}
		r.Moq.ResultsByParams_Broadcast = append(r.Moq.ResultsByParams_Broadcast, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Broadcast) {
			copy(r.Moq.ResultsByParams_Broadcast[insertAt+1:], r.Moq.ResultsByParams_Broadcast[insertAt:0])
			r.Moq.ResultsByParams_Broadcast[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Broadcast(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqCond_starGenType_Broadcast_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqCond_starGenType_Broadcast_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqCond_starGenType_Broadcast_fnRecorder {
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
				DoFn       MoqCond_starGenType_Broadcast_doFn
				DoReturnFn MoqCond_starGenType_Broadcast_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqCond_starGenType) PrettyParams_Broadcast(params MoqCond_starGenType_Broadcast_params) string {
	return fmt.Sprintf("Broadcast()")
}

func (m *MoqCond_starGenType) ParamsKey_Broadcast(params MoqCond_starGenType_Broadcast_params, anyParams uint64) MoqCond_starGenType_Broadcast_paramsKey {
	m.Scene.T.Helper()
	return MoqCond_starGenType_Broadcast_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqCond_starGenType) Reset() {
	m.ResultsByParams_Wait = nil
	m.ResultsByParams_Signal = nil
	m.ResultsByParams_Broadcast = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqCond_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Wait {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Wait(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Signal {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Signal(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Broadcast {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Broadcast(results.Params))
			}
		}
	}
}

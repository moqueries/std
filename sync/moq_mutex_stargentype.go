// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package sync

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that sync.Mutex_starGenType is mocked
// completely
var _ Mutex_starGenType = (*MoqMutex_starGenType_mock)(nil)

// Mutex_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type Mutex_starGenType interface {
	Lock()
	Unlock()
}

// MoqMutex_starGenType holds the state of a moq of the Mutex_starGenType type
type MoqMutex_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqMutex_starGenType_mock

	ResultsByParams_Lock   []MoqMutex_starGenType_Lock_resultsByParams
	ResultsByParams_Unlock []MoqMutex_starGenType_Unlock_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Lock   struct{}
			Unlock struct{}
		}
	}
}

// MoqMutex_starGenType_mock isolates the mock interface of the
// Mutex_starGenType type
type MoqMutex_starGenType_mock struct {
	Moq *MoqMutex_starGenType
}

// MoqMutex_starGenType_recorder isolates the recorder interface of the
// Mutex_starGenType type
type MoqMutex_starGenType_recorder struct {
	Moq *MoqMutex_starGenType
}

// MoqMutex_starGenType_Lock_params holds the params of the Mutex_starGenType
// type
type MoqMutex_starGenType_Lock_params struct{}

// MoqMutex_starGenType_Lock_paramsKey holds the map key params of the
// Mutex_starGenType type
type MoqMutex_starGenType_Lock_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqMutex_starGenType_Lock_resultsByParams contains the results for a given
// set of parameters for the Mutex_starGenType type
type MoqMutex_starGenType_Lock_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqMutex_starGenType_Lock_paramsKey]*MoqMutex_starGenType_Lock_results
}

// MoqMutex_starGenType_Lock_doFn defines the type of function needed when
// calling AndDo for the Mutex_starGenType type
type MoqMutex_starGenType_Lock_doFn func()

// MoqMutex_starGenType_Lock_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Mutex_starGenType type
type MoqMutex_starGenType_Lock_doReturnFn func()

// MoqMutex_starGenType_Lock_results holds the results of the Mutex_starGenType
// type
type MoqMutex_starGenType_Lock_results struct {
	Params  MoqMutex_starGenType_Lock_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqMutex_starGenType_Lock_doFn
		DoReturnFn MoqMutex_starGenType_Lock_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqMutex_starGenType_Lock_fnRecorder routes recorded function calls to the
// MoqMutex_starGenType moq
type MoqMutex_starGenType_Lock_fnRecorder struct {
	Params    MoqMutex_starGenType_Lock_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqMutex_starGenType_Lock_results
	Moq       *MoqMutex_starGenType
}

// MoqMutex_starGenType_Lock_anyParams isolates the any params functions of the
// Mutex_starGenType type
type MoqMutex_starGenType_Lock_anyParams struct {
	Recorder *MoqMutex_starGenType_Lock_fnRecorder
}

// MoqMutex_starGenType_Unlock_params holds the params of the Mutex_starGenType
// type
type MoqMutex_starGenType_Unlock_params struct{}

// MoqMutex_starGenType_Unlock_paramsKey holds the map key params of the
// Mutex_starGenType type
type MoqMutex_starGenType_Unlock_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqMutex_starGenType_Unlock_resultsByParams contains the results for a given
// set of parameters for the Mutex_starGenType type
type MoqMutex_starGenType_Unlock_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqMutex_starGenType_Unlock_paramsKey]*MoqMutex_starGenType_Unlock_results
}

// MoqMutex_starGenType_Unlock_doFn defines the type of function needed when
// calling AndDo for the Mutex_starGenType type
type MoqMutex_starGenType_Unlock_doFn func()

// MoqMutex_starGenType_Unlock_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Mutex_starGenType type
type MoqMutex_starGenType_Unlock_doReturnFn func()

// MoqMutex_starGenType_Unlock_results holds the results of the
// Mutex_starGenType type
type MoqMutex_starGenType_Unlock_results struct {
	Params  MoqMutex_starGenType_Unlock_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqMutex_starGenType_Unlock_doFn
		DoReturnFn MoqMutex_starGenType_Unlock_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqMutex_starGenType_Unlock_fnRecorder routes recorded function calls to the
// MoqMutex_starGenType moq
type MoqMutex_starGenType_Unlock_fnRecorder struct {
	Params    MoqMutex_starGenType_Unlock_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqMutex_starGenType_Unlock_results
	Moq       *MoqMutex_starGenType
}

// MoqMutex_starGenType_Unlock_anyParams isolates the any params functions of
// the Mutex_starGenType type
type MoqMutex_starGenType_Unlock_anyParams struct {
	Recorder *MoqMutex_starGenType_Unlock_fnRecorder
}

// NewMoqMutex_starGenType creates a new moq of the Mutex_starGenType type
func NewMoqMutex_starGenType(scene *moq.Scene, config *moq.Config) *MoqMutex_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqMutex_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqMutex_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Lock   struct{}
				Unlock struct{}
			}
		}{ParameterIndexing: struct {
			Lock   struct{}
			Unlock struct{}
		}{
			Lock:   struct{}{},
			Unlock: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Mutex_starGenType type
func (m *MoqMutex_starGenType) Mock() *MoqMutex_starGenType_mock { return m.Moq }

func (m *MoqMutex_starGenType_mock) Lock() {
	m.Moq.Scene.T.Helper()
	params := MoqMutex_starGenType_Lock_params{}
	var results *MoqMutex_starGenType_Lock_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Lock {
		paramsKey := m.Moq.ParamsKey_Lock(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Lock(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Lock(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Lock(params))
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

func (m *MoqMutex_starGenType_mock) Unlock() {
	m.Moq.Scene.T.Helper()
	params := MoqMutex_starGenType_Unlock_params{}
	var results *MoqMutex_starGenType_Unlock_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Unlock {
		paramsKey := m.Moq.ParamsKey_Unlock(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Unlock(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Unlock(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Unlock(params))
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

// OnCall returns the recorder implementation of the Mutex_starGenType type
func (m *MoqMutex_starGenType) OnCall() *MoqMutex_starGenType_recorder {
	return &MoqMutex_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqMutex_starGenType_recorder) Lock() *MoqMutex_starGenType_Lock_fnRecorder {
	return &MoqMutex_starGenType_Lock_fnRecorder{
		Params:   MoqMutex_starGenType_Lock_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqMutex_starGenType_Lock_fnRecorder) Any() *MoqMutex_starGenType_Lock_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Lock(r.Params))
		return nil
	}
	return &MoqMutex_starGenType_Lock_anyParams{Recorder: r}
}

func (r *MoqMutex_starGenType_Lock_fnRecorder) Seq() *MoqMutex_starGenType_Lock_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Lock(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqMutex_starGenType_Lock_fnRecorder) NoSeq() *MoqMutex_starGenType_Lock_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Lock(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqMutex_starGenType_Lock_fnRecorder) ReturnResults() *MoqMutex_starGenType_Lock_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqMutex_starGenType_Lock_doFn
		DoReturnFn MoqMutex_starGenType_Lock_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqMutex_starGenType_Lock_fnRecorder) AndDo(fn MoqMutex_starGenType_Lock_doFn) *MoqMutex_starGenType_Lock_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqMutex_starGenType_Lock_fnRecorder) DoReturnResults(fn MoqMutex_starGenType_Lock_doReturnFn) *MoqMutex_starGenType_Lock_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqMutex_starGenType_Lock_doFn
		DoReturnFn MoqMutex_starGenType_Lock_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqMutex_starGenType_Lock_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqMutex_starGenType_Lock_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Lock {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqMutex_starGenType_Lock_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqMutex_starGenType_Lock_paramsKey]*MoqMutex_starGenType_Lock_results{},
		}
		r.Moq.ResultsByParams_Lock = append(r.Moq.ResultsByParams_Lock, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Lock) {
			copy(r.Moq.ResultsByParams_Lock[insertAt+1:], r.Moq.ResultsByParams_Lock[insertAt:0])
			r.Moq.ResultsByParams_Lock[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Lock(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqMutex_starGenType_Lock_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqMutex_starGenType_Lock_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqMutex_starGenType_Lock_fnRecorder {
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
				DoFn       MoqMutex_starGenType_Lock_doFn
				DoReturnFn MoqMutex_starGenType_Lock_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqMutex_starGenType) PrettyParams_Lock(params MoqMutex_starGenType_Lock_params) string {
	return fmt.Sprintf("Lock()")
}

func (m *MoqMutex_starGenType) ParamsKey_Lock(params MoqMutex_starGenType_Lock_params, anyParams uint64) MoqMutex_starGenType_Lock_paramsKey {
	m.Scene.T.Helper()
	return MoqMutex_starGenType_Lock_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqMutex_starGenType_recorder) Unlock() *MoqMutex_starGenType_Unlock_fnRecorder {
	return &MoqMutex_starGenType_Unlock_fnRecorder{
		Params:   MoqMutex_starGenType_Unlock_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqMutex_starGenType_Unlock_fnRecorder) Any() *MoqMutex_starGenType_Unlock_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Unlock(r.Params))
		return nil
	}
	return &MoqMutex_starGenType_Unlock_anyParams{Recorder: r}
}

func (r *MoqMutex_starGenType_Unlock_fnRecorder) Seq() *MoqMutex_starGenType_Unlock_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Unlock(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqMutex_starGenType_Unlock_fnRecorder) NoSeq() *MoqMutex_starGenType_Unlock_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Unlock(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqMutex_starGenType_Unlock_fnRecorder) ReturnResults() *MoqMutex_starGenType_Unlock_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqMutex_starGenType_Unlock_doFn
		DoReturnFn MoqMutex_starGenType_Unlock_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqMutex_starGenType_Unlock_fnRecorder) AndDo(fn MoqMutex_starGenType_Unlock_doFn) *MoqMutex_starGenType_Unlock_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqMutex_starGenType_Unlock_fnRecorder) DoReturnResults(fn MoqMutex_starGenType_Unlock_doReturnFn) *MoqMutex_starGenType_Unlock_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqMutex_starGenType_Unlock_doFn
		DoReturnFn MoqMutex_starGenType_Unlock_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqMutex_starGenType_Unlock_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqMutex_starGenType_Unlock_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Unlock {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqMutex_starGenType_Unlock_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqMutex_starGenType_Unlock_paramsKey]*MoqMutex_starGenType_Unlock_results{},
		}
		r.Moq.ResultsByParams_Unlock = append(r.Moq.ResultsByParams_Unlock, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Unlock) {
			copy(r.Moq.ResultsByParams_Unlock[insertAt+1:], r.Moq.ResultsByParams_Unlock[insertAt:0])
			r.Moq.ResultsByParams_Unlock[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Unlock(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqMutex_starGenType_Unlock_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqMutex_starGenType_Unlock_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqMutex_starGenType_Unlock_fnRecorder {
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
				DoFn       MoqMutex_starGenType_Unlock_doFn
				DoReturnFn MoqMutex_starGenType_Unlock_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqMutex_starGenType) PrettyParams_Unlock(params MoqMutex_starGenType_Unlock_params) string {
	return fmt.Sprintf("Unlock()")
}

func (m *MoqMutex_starGenType) ParamsKey_Unlock(params MoqMutex_starGenType_Unlock_params, anyParams uint64) MoqMutex_starGenType_Unlock_paramsKey {
	m.Scene.T.Helper()
	return MoqMutex_starGenType_Unlock_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqMutex_starGenType) Reset() { m.ResultsByParams_Lock = nil; m.ResultsByParams_Unlock = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqMutex_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Lock {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Lock(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Unlock {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Unlock(results.Params))
			}
		}
	}
}

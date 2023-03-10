// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package time

import (
	"fmt"
	"math/bits"
	"sync/atomic"
	"time"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that time.Ticker_starGenType is mocked
// completely
var _ Ticker_starGenType = (*MoqTicker_starGenType_mock)(nil)

// Ticker_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type Ticker_starGenType interface {
	Stop()
	Reset(d time.Duration)
}

// MoqTicker_starGenType holds the state of a moq of the Ticker_starGenType
// type
type MoqTicker_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqTicker_starGenType_mock

	ResultsByParams_Stop  []MoqTicker_starGenType_Stop_resultsByParams
	ResultsByParams_Reset []MoqTicker_starGenType_Reset_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Stop  struct{}
			Reset struct {
				D moq.ParamIndexing
			}
		}
	}
	// MoqTicker_starGenType_mock isolates the mock interface of the
}

// Ticker_starGenType type
type MoqTicker_starGenType_mock struct {
	Moq *MoqTicker_starGenType
}

// MoqTicker_starGenType_recorder isolates the recorder interface of the
// Ticker_starGenType type
type MoqTicker_starGenType_recorder struct {
	Moq *MoqTicker_starGenType
}

// MoqTicker_starGenType_Stop_params holds the params of the Ticker_starGenType
// type
type MoqTicker_starGenType_Stop_params struct{}

// MoqTicker_starGenType_Stop_paramsKey holds the map key params of the
// Ticker_starGenType type
type MoqTicker_starGenType_Stop_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqTicker_starGenType_Stop_resultsByParams contains the results for a given
// set of parameters for the Ticker_starGenType type
type MoqTicker_starGenType_Stop_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqTicker_starGenType_Stop_paramsKey]*MoqTicker_starGenType_Stop_results
}

// MoqTicker_starGenType_Stop_doFn defines the type of function needed when
// calling AndDo for the Ticker_starGenType type
type MoqTicker_starGenType_Stop_doFn func()

// MoqTicker_starGenType_Stop_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Ticker_starGenType type
type MoqTicker_starGenType_Stop_doReturnFn func()

// MoqTicker_starGenType_Stop_results holds the results of the
// Ticker_starGenType type
type MoqTicker_starGenType_Stop_results struct {
	Params  MoqTicker_starGenType_Stop_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqTicker_starGenType_Stop_doFn
		DoReturnFn MoqTicker_starGenType_Stop_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqTicker_starGenType_Stop_fnRecorder routes recorded function calls to the
// MoqTicker_starGenType moq
type MoqTicker_starGenType_Stop_fnRecorder struct {
	Params    MoqTicker_starGenType_Stop_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqTicker_starGenType_Stop_results
	Moq       *MoqTicker_starGenType
}

// MoqTicker_starGenType_Stop_anyParams isolates the any params functions of
// the Ticker_starGenType type
type MoqTicker_starGenType_Stop_anyParams struct {
	Recorder *MoqTicker_starGenType_Stop_fnRecorder
}

// MoqTicker_starGenType_Reset_params holds the params of the
// Ticker_starGenType type
type MoqTicker_starGenType_Reset_params struct{ D time.Duration }

// MoqTicker_starGenType_Reset_paramsKey holds the map key params of the
// Ticker_starGenType type
type MoqTicker_starGenType_Reset_paramsKey struct {
	Params struct{ D time.Duration }
	Hashes struct{ D hash.Hash }
}

// MoqTicker_starGenType_Reset_resultsByParams contains the results for a given
// set of parameters for the Ticker_starGenType type
type MoqTicker_starGenType_Reset_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqTicker_starGenType_Reset_paramsKey]*MoqTicker_starGenType_Reset_results
}

// MoqTicker_starGenType_Reset_doFn defines the type of function needed when
// calling AndDo for the Ticker_starGenType type
type MoqTicker_starGenType_Reset_doFn func(d time.Duration)

// MoqTicker_starGenType_Reset_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Ticker_starGenType type
type MoqTicker_starGenType_Reset_doReturnFn func(d time.Duration)

// MoqTicker_starGenType_Reset_results holds the results of the
// Ticker_starGenType type
type MoqTicker_starGenType_Reset_results struct {
	Params  MoqTicker_starGenType_Reset_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqTicker_starGenType_Reset_doFn
		DoReturnFn MoqTicker_starGenType_Reset_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqTicker_starGenType_Reset_fnRecorder routes recorded function calls to the
// MoqTicker_starGenType moq
type MoqTicker_starGenType_Reset_fnRecorder struct {
	Params    MoqTicker_starGenType_Reset_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqTicker_starGenType_Reset_results
	Moq       *MoqTicker_starGenType
}

// MoqTicker_starGenType_Reset_anyParams isolates the any params functions of
// the Ticker_starGenType type
type MoqTicker_starGenType_Reset_anyParams struct {
	Recorder *MoqTicker_starGenType_Reset_fnRecorder
}

// NewMoqTicker_starGenType creates a new moq of the Ticker_starGenType type
func NewMoqTicker_starGenType(scene *moq.Scene, config *moq.Config) *MoqTicker_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqTicker_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqTicker_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Stop  struct{}
				Reset struct {
					D moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Stop  struct{}
			Reset struct {
				D moq.ParamIndexing
			}
		}{
			Stop: struct{}{},
			Reset: struct {
				D moq.ParamIndexing
			}{
				D: moq.ParamIndexByValue,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Ticker_starGenType type
func (m *MoqTicker_starGenType) Mock() *MoqTicker_starGenType_mock { return m.Moq }

func (m *MoqTicker_starGenType_mock) Stop() {
	m.Moq.Scene.T.Helper()
	params := MoqTicker_starGenType_Stop_params{}
	var results *MoqTicker_starGenType_Stop_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Stop {
		paramsKey := m.Moq.ParamsKey_Stop(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Stop(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Stop(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Stop(params))
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

func (m *MoqTicker_starGenType_mock) Reset(d time.Duration) {
	m.Moq.Scene.T.Helper()
	params := MoqTicker_starGenType_Reset_params{
		D: d,
	}
	var results *MoqTicker_starGenType_Reset_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Reset {
		paramsKey := m.Moq.ParamsKey_Reset(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Reset(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Reset(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Reset(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(d)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(d)
	}
	return
}

// OnCall returns the recorder implementation of the Ticker_starGenType type
func (m *MoqTicker_starGenType) OnCall() *MoqTicker_starGenType_recorder {
	return &MoqTicker_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqTicker_starGenType_recorder) Stop() *MoqTicker_starGenType_Stop_fnRecorder {
	return &MoqTicker_starGenType_Stop_fnRecorder{
		Params:   MoqTicker_starGenType_Stop_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqTicker_starGenType_Stop_fnRecorder) Any() *MoqTicker_starGenType_Stop_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Stop(r.Params))
		return nil
	}
	return &MoqTicker_starGenType_Stop_anyParams{Recorder: r}
}

func (r *MoqTicker_starGenType_Stop_fnRecorder) Seq() *MoqTicker_starGenType_Stop_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Stop(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqTicker_starGenType_Stop_fnRecorder) NoSeq() *MoqTicker_starGenType_Stop_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Stop(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqTicker_starGenType_Stop_fnRecorder) ReturnResults() *MoqTicker_starGenType_Stop_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqTicker_starGenType_Stop_doFn
		DoReturnFn MoqTicker_starGenType_Stop_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqTicker_starGenType_Stop_fnRecorder) AndDo(fn MoqTicker_starGenType_Stop_doFn) *MoqTicker_starGenType_Stop_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqTicker_starGenType_Stop_fnRecorder) DoReturnResults(fn MoqTicker_starGenType_Stop_doReturnFn) *MoqTicker_starGenType_Stop_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqTicker_starGenType_Stop_doFn
		DoReturnFn MoqTicker_starGenType_Stop_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqTicker_starGenType_Stop_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqTicker_starGenType_Stop_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Stop {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqTicker_starGenType_Stop_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqTicker_starGenType_Stop_paramsKey]*MoqTicker_starGenType_Stop_results{},
		}
		r.Moq.ResultsByParams_Stop = append(r.Moq.ResultsByParams_Stop, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Stop) {
			copy(r.Moq.ResultsByParams_Stop[insertAt+1:], r.Moq.ResultsByParams_Stop[insertAt:0])
			r.Moq.ResultsByParams_Stop[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Stop(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqTicker_starGenType_Stop_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqTicker_starGenType_Stop_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqTicker_starGenType_Stop_fnRecorder {
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
				DoFn       MoqTicker_starGenType_Stop_doFn
				DoReturnFn MoqTicker_starGenType_Stop_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqTicker_starGenType) PrettyParams_Stop(params MoqTicker_starGenType_Stop_params) string {
	return fmt.Sprintf("Stop()")
}

func (m *MoqTicker_starGenType) ParamsKey_Stop(params MoqTicker_starGenType_Stop_params, anyParams uint64) MoqTicker_starGenType_Stop_paramsKey {
	m.Scene.T.Helper()
	return MoqTicker_starGenType_Stop_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqTicker_starGenType_recorder) Reset(d time.Duration) *MoqTicker_starGenType_Reset_fnRecorder {
	return &MoqTicker_starGenType_Reset_fnRecorder{
		Params: MoqTicker_starGenType_Reset_params{
			D: d,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqTicker_starGenType_Reset_fnRecorder) Any() *MoqTicker_starGenType_Reset_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Reset(r.Params))
		return nil
	}
	return &MoqTicker_starGenType_Reset_anyParams{Recorder: r}
}

func (a *MoqTicker_starGenType_Reset_anyParams) D() *MoqTicker_starGenType_Reset_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqTicker_starGenType_Reset_fnRecorder) Seq() *MoqTicker_starGenType_Reset_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Reset(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqTicker_starGenType_Reset_fnRecorder) NoSeq() *MoqTicker_starGenType_Reset_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Reset(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqTicker_starGenType_Reset_fnRecorder) ReturnResults() *MoqTicker_starGenType_Reset_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqTicker_starGenType_Reset_doFn
		DoReturnFn MoqTicker_starGenType_Reset_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqTicker_starGenType_Reset_fnRecorder) AndDo(fn MoqTicker_starGenType_Reset_doFn) *MoqTicker_starGenType_Reset_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqTicker_starGenType_Reset_fnRecorder) DoReturnResults(fn MoqTicker_starGenType_Reset_doReturnFn) *MoqTicker_starGenType_Reset_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqTicker_starGenType_Reset_doFn
		DoReturnFn MoqTicker_starGenType_Reset_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqTicker_starGenType_Reset_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqTicker_starGenType_Reset_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Reset {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqTicker_starGenType_Reset_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqTicker_starGenType_Reset_paramsKey]*MoqTicker_starGenType_Reset_results{},
		}
		r.Moq.ResultsByParams_Reset = append(r.Moq.ResultsByParams_Reset, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Reset) {
			copy(r.Moq.ResultsByParams_Reset[insertAt+1:], r.Moq.ResultsByParams_Reset[insertAt:0])
			r.Moq.ResultsByParams_Reset[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Reset(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqTicker_starGenType_Reset_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqTicker_starGenType_Reset_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqTicker_starGenType_Reset_fnRecorder {
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
				DoFn       MoqTicker_starGenType_Reset_doFn
				DoReturnFn MoqTicker_starGenType_Reset_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqTicker_starGenType) PrettyParams_Reset(params MoqTicker_starGenType_Reset_params) string {
	return fmt.Sprintf("Reset(%#v)", params.D)
}

func (m *MoqTicker_starGenType) ParamsKey_Reset(params MoqTicker_starGenType_Reset_params, anyParams uint64) MoqTicker_starGenType_Reset_paramsKey {
	m.Scene.T.Helper()
	var dUsed time.Duration
	var dUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Reset.D == moq.ParamIndexByValue {
			dUsed = params.D
		} else {
			dUsedHash = hash.DeepHash(params.D)
		}
	}
	return MoqTicker_starGenType_Reset_paramsKey{
		Params: struct{ D time.Duration }{
			D: dUsed,
		},
		Hashes: struct{ D hash.Hash }{
			D: dUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqTicker_starGenType) Reset() { m.ResultsByParams_Stop = nil; m.ResultsByParams_Reset = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqTicker_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Stop {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Stop(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Reset {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Reset(results.Params))
			}
		}
	}
}

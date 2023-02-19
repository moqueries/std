// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package debug

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// SetPanicOnFault_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type SetPanicOnFault_genType func(enabled bool) bool

// MoqSetPanicOnFault_genType holds the state of a moq of the
// SetPanicOnFault_genType type
type MoqSetPanicOnFault_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSetPanicOnFault_genType_mock

	ResultsByParams []MoqSetPanicOnFault_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Enabled moq.ParamIndexing
		}
	}
}

// MoqSetPanicOnFault_genType_mock isolates the mock interface of the
// SetPanicOnFault_genType type
type MoqSetPanicOnFault_genType_mock struct {
	Moq *MoqSetPanicOnFault_genType
}

// MoqSetPanicOnFault_genType_params holds the params of the
// SetPanicOnFault_genType type
type MoqSetPanicOnFault_genType_params struct{ Enabled bool }

// MoqSetPanicOnFault_genType_paramsKey holds the map key params of the
// SetPanicOnFault_genType type
type MoqSetPanicOnFault_genType_paramsKey struct {
	Params struct{ Enabled bool }
	Hashes struct{ Enabled hash.Hash }
}

// MoqSetPanicOnFault_genType_resultsByParams contains the results for a given
// set of parameters for the SetPanicOnFault_genType type
type MoqSetPanicOnFault_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSetPanicOnFault_genType_paramsKey]*MoqSetPanicOnFault_genType_results
}

// MoqSetPanicOnFault_genType_doFn defines the type of function needed when
// calling AndDo for the SetPanicOnFault_genType type
type MoqSetPanicOnFault_genType_doFn func(enabled bool)

// MoqSetPanicOnFault_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the SetPanicOnFault_genType type
type MoqSetPanicOnFault_genType_doReturnFn func(enabled bool) bool

// MoqSetPanicOnFault_genType_results holds the results of the
// SetPanicOnFault_genType type
type MoqSetPanicOnFault_genType_results struct {
	Params  MoqSetPanicOnFault_genType_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqSetPanicOnFault_genType_doFn
		DoReturnFn MoqSetPanicOnFault_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSetPanicOnFault_genType_fnRecorder routes recorded function calls to the
// MoqSetPanicOnFault_genType moq
type MoqSetPanicOnFault_genType_fnRecorder struct {
	Params    MoqSetPanicOnFault_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSetPanicOnFault_genType_results
	Moq       *MoqSetPanicOnFault_genType
}

// MoqSetPanicOnFault_genType_anyParams isolates the any params functions of
// the SetPanicOnFault_genType type
type MoqSetPanicOnFault_genType_anyParams struct {
	Recorder *MoqSetPanicOnFault_genType_fnRecorder
}

// NewMoqSetPanicOnFault_genType creates a new moq of the
// SetPanicOnFault_genType type
func NewMoqSetPanicOnFault_genType(scene *moq.Scene, config *moq.Config) *MoqSetPanicOnFault_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSetPanicOnFault_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSetPanicOnFault_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Enabled moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Enabled moq.ParamIndexing
		}{
			Enabled: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the SetPanicOnFault_genType type
func (m *MoqSetPanicOnFault_genType) Mock() SetPanicOnFault_genType {
	return func(enabled bool) bool {
		m.Scene.T.Helper()
		moq := &MoqSetPanicOnFault_genType_mock{Moq: m}
		return moq.Fn(enabled)
	}
}

func (m *MoqSetPanicOnFault_genType_mock) Fn(enabled bool) (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqSetPanicOnFault_genType_params{
		Enabled: enabled,
	}
	var results *MoqSetPanicOnFault_genType_results
	for _, resultsByParams := range m.Moq.ResultsByParams {
		paramsKey := m.Moq.ParamsKey(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(enabled)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(enabled)
	}
	return
}

func (m *MoqSetPanicOnFault_genType) OnCall(enabled bool) *MoqSetPanicOnFault_genType_fnRecorder {
	return &MoqSetPanicOnFault_genType_fnRecorder{
		Params: MoqSetPanicOnFault_genType_params{
			Enabled: enabled,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSetPanicOnFault_genType_fnRecorder) Any() *MoqSetPanicOnFault_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSetPanicOnFault_genType_anyParams{Recorder: r}
}

func (a *MoqSetPanicOnFault_genType_anyParams) Enabled() *MoqSetPanicOnFault_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqSetPanicOnFault_genType_fnRecorder) Seq() *MoqSetPanicOnFault_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSetPanicOnFault_genType_fnRecorder) NoSeq() *MoqSetPanicOnFault_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSetPanicOnFault_genType_fnRecorder) ReturnResults(result1 bool) *MoqSetPanicOnFault_genType_fnRecorder {
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
		DoFn       MoqSetPanicOnFault_genType_doFn
		DoReturnFn MoqSetPanicOnFault_genType_doReturnFn
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

func (r *MoqSetPanicOnFault_genType_fnRecorder) AndDo(fn MoqSetPanicOnFault_genType_doFn) *MoqSetPanicOnFault_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSetPanicOnFault_genType_fnRecorder) DoReturnResults(fn MoqSetPanicOnFault_genType_doReturnFn) *MoqSetPanicOnFault_genType_fnRecorder {
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
		DoFn       MoqSetPanicOnFault_genType_doFn
		DoReturnFn MoqSetPanicOnFault_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSetPanicOnFault_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSetPanicOnFault_genType_resultsByParams
	for n, res := range r.Moq.ResultsByParams {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqSetPanicOnFault_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSetPanicOnFault_genType_paramsKey]*MoqSetPanicOnFault_genType_results{},
		}
		r.Moq.ResultsByParams = append(r.Moq.ResultsByParams, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams) {
			copy(r.Moq.ResultsByParams[insertAt+1:], r.Moq.ResultsByParams[insertAt:0])
			r.Moq.ResultsByParams[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqSetPanicOnFault_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSetPanicOnFault_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSetPanicOnFault_genType_fnRecorder {
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
				DoFn       MoqSetPanicOnFault_genType_doFn
				DoReturnFn MoqSetPanicOnFault_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSetPanicOnFault_genType) PrettyParams(params MoqSetPanicOnFault_genType_params) string {
	return fmt.Sprintf("SetPanicOnFault_genType(%#v)", params.Enabled)
}

func (m *MoqSetPanicOnFault_genType) ParamsKey(params MoqSetPanicOnFault_genType_params, anyParams uint64) MoqSetPanicOnFault_genType_paramsKey {
	m.Scene.T.Helper()
	var enabledUsed bool
	var enabledUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Enabled == moq.ParamIndexByValue {
			enabledUsed = params.Enabled
		} else {
			enabledUsedHash = hash.DeepHash(params.Enabled)
		}
	}
	return MoqSetPanicOnFault_genType_paramsKey{
		Params: struct{ Enabled bool }{
			Enabled: enabledUsed,
		},
		Hashes: struct{ Enabled hash.Hash }{
			Enabled: enabledUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSetPanicOnFault_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSetPanicOnFault_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams(results.Params))
			}
		}
	}
}
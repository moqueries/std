// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package runtime

import (
	"fmt"
	"math/bits"
	"runtime"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// FuncForPC_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type FuncForPC_genType func(pc uintptr) *runtime.Func

// MoqFuncForPC_genType holds the state of a moq of the FuncForPC_genType type
type MoqFuncForPC_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqFuncForPC_genType_mock

	ResultsByParams []MoqFuncForPC_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Pc moq.ParamIndexing
		}
	}
}

// MoqFuncForPC_genType_mock isolates the mock interface of the
// FuncForPC_genType type
type MoqFuncForPC_genType_mock struct {
	Moq *MoqFuncForPC_genType
}

// MoqFuncForPC_genType_params holds the params of the FuncForPC_genType type
type MoqFuncForPC_genType_params struct{ Pc uintptr }

// MoqFuncForPC_genType_paramsKey holds the map key params of the
// FuncForPC_genType type
type MoqFuncForPC_genType_paramsKey struct {
	Params struct{ Pc uintptr }
	Hashes struct{ Pc hash.Hash }
}

// MoqFuncForPC_genType_resultsByParams contains the results for a given set of
// parameters for the FuncForPC_genType type
type MoqFuncForPC_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFuncForPC_genType_paramsKey]*MoqFuncForPC_genType_results
}

// MoqFuncForPC_genType_doFn defines the type of function needed when calling
// AndDo for the FuncForPC_genType type
type MoqFuncForPC_genType_doFn func(pc uintptr)

// MoqFuncForPC_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the FuncForPC_genType type
type MoqFuncForPC_genType_doReturnFn func(pc uintptr) *runtime.Func

// MoqFuncForPC_genType_results holds the results of the FuncForPC_genType type
type MoqFuncForPC_genType_results struct {
	Params  MoqFuncForPC_genType_params
	Results []struct {
		Values *struct {
			Result1 *runtime.Func
		}
		Sequence   uint32
		DoFn       MoqFuncForPC_genType_doFn
		DoReturnFn MoqFuncForPC_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFuncForPC_genType_fnRecorder routes recorded function calls to the
// MoqFuncForPC_genType moq
type MoqFuncForPC_genType_fnRecorder struct {
	Params    MoqFuncForPC_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFuncForPC_genType_results
	Moq       *MoqFuncForPC_genType
}

// MoqFuncForPC_genType_anyParams isolates the any params functions of the
// FuncForPC_genType type
type MoqFuncForPC_genType_anyParams struct {
	Recorder *MoqFuncForPC_genType_fnRecorder
}

// NewMoqFuncForPC_genType creates a new moq of the FuncForPC_genType type
func NewMoqFuncForPC_genType(scene *moq.Scene, config *moq.Config) *MoqFuncForPC_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqFuncForPC_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqFuncForPC_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Pc moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Pc moq.ParamIndexing
		}{
			Pc: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the FuncForPC_genType type
func (m *MoqFuncForPC_genType) Mock() FuncForPC_genType {
	return func(pc uintptr) *runtime.Func {
		m.Scene.T.Helper()
		moq := &MoqFuncForPC_genType_mock{Moq: m}
		return moq.Fn(pc)
	}
}

func (m *MoqFuncForPC_genType_mock) Fn(pc uintptr) (result1 *runtime.Func) {
	m.Moq.Scene.T.Helper()
	params := MoqFuncForPC_genType_params{
		Pc: pc,
	}
	var results *MoqFuncForPC_genType_results
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
		result.DoFn(pc)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(pc)
	}
	return
}

func (m *MoqFuncForPC_genType) OnCall(pc uintptr) *MoqFuncForPC_genType_fnRecorder {
	return &MoqFuncForPC_genType_fnRecorder{
		Params: MoqFuncForPC_genType_params{
			Pc: pc,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqFuncForPC_genType_fnRecorder) Any() *MoqFuncForPC_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqFuncForPC_genType_anyParams{Recorder: r}
}

func (a *MoqFuncForPC_genType_anyParams) Pc() *MoqFuncForPC_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqFuncForPC_genType_fnRecorder) Seq() *MoqFuncForPC_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFuncForPC_genType_fnRecorder) NoSeq() *MoqFuncForPC_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFuncForPC_genType_fnRecorder) ReturnResults(result1 *runtime.Func) *MoqFuncForPC_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *runtime.Func
		}
		Sequence   uint32
		DoFn       MoqFuncForPC_genType_doFn
		DoReturnFn MoqFuncForPC_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *runtime.Func
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqFuncForPC_genType_fnRecorder) AndDo(fn MoqFuncForPC_genType_doFn) *MoqFuncForPC_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFuncForPC_genType_fnRecorder) DoReturnResults(fn MoqFuncForPC_genType_doReturnFn) *MoqFuncForPC_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *runtime.Func
		}
		Sequence   uint32
		DoFn       MoqFuncForPC_genType_doFn
		DoReturnFn MoqFuncForPC_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFuncForPC_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFuncForPC_genType_resultsByParams
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
		results = &MoqFuncForPC_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFuncForPC_genType_paramsKey]*MoqFuncForPC_genType_results{},
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
		r.Results = &MoqFuncForPC_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFuncForPC_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFuncForPC_genType_fnRecorder {
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
					Result1 *runtime.Func
				}
				Sequence   uint32
				DoFn       MoqFuncForPC_genType_doFn
				DoReturnFn MoqFuncForPC_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFuncForPC_genType) PrettyParams(params MoqFuncForPC_genType_params) string {
	return fmt.Sprintf("FuncForPC_genType(%#v)", params.Pc)
}

func (m *MoqFuncForPC_genType) ParamsKey(params MoqFuncForPC_genType_params, anyParams uint64) MoqFuncForPC_genType_paramsKey {
	m.Scene.T.Helper()
	var pcUsed uintptr
	var pcUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Pc == moq.ParamIndexByValue {
			pcUsed = params.Pc
		} else {
			pcUsedHash = hash.DeepHash(params.Pc)
		}
	}
	return MoqFuncForPC_genType_paramsKey{
		Params: struct{ Pc uintptr }{
			Pc: pcUsed,
		},
		Hashes: struct{ Pc hash.Hash }{
			Pc: pcUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqFuncForPC_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqFuncForPC_genType) AssertExpectationsMet() {
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

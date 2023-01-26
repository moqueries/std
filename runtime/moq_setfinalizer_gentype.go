// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package runtime

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// SetFinalizer_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type SetFinalizer_genType func(obj interface{}, finalizer interface{})

// MoqSetFinalizer_genType holds the state of a moq of the SetFinalizer_genType
// type
type MoqSetFinalizer_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSetFinalizer_genType_mock

	ResultsByParams []MoqSetFinalizer_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Obj       moq.ParamIndexing
			Finalizer moq.ParamIndexing
		}
	}
}

// MoqSetFinalizer_genType_mock isolates the mock interface of the
// SetFinalizer_genType type
type MoqSetFinalizer_genType_mock struct {
	Moq *MoqSetFinalizer_genType
}

// MoqSetFinalizer_genType_params holds the params of the SetFinalizer_genType
// type
type MoqSetFinalizer_genType_params struct {
	Obj       interface{}
	Finalizer interface{}
}

// MoqSetFinalizer_genType_paramsKey holds the map key params of the
// SetFinalizer_genType type
type MoqSetFinalizer_genType_paramsKey struct {
	Params struct {
		Obj       interface{}
		Finalizer interface{}
	}
	Hashes struct {
		Obj       hash.Hash
		Finalizer hash.Hash
	}
}

// MoqSetFinalizer_genType_resultsByParams contains the results for a given set
// of parameters for the SetFinalizer_genType type
type MoqSetFinalizer_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSetFinalizer_genType_paramsKey]*MoqSetFinalizer_genType_results
}

// MoqSetFinalizer_genType_doFn defines the type of function needed when
// calling AndDo for the SetFinalizer_genType type
type MoqSetFinalizer_genType_doFn func(obj interface{}, finalizer interface{})

// MoqSetFinalizer_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the SetFinalizer_genType type
type MoqSetFinalizer_genType_doReturnFn func(obj interface{}, finalizer interface{})

// MoqSetFinalizer_genType_results holds the results of the
// SetFinalizer_genType type
type MoqSetFinalizer_genType_results struct {
	Params  MoqSetFinalizer_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSetFinalizer_genType_doFn
		DoReturnFn MoqSetFinalizer_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSetFinalizer_genType_fnRecorder routes recorded function calls to the
// MoqSetFinalizer_genType moq
type MoqSetFinalizer_genType_fnRecorder struct {
	Params    MoqSetFinalizer_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSetFinalizer_genType_results
	Moq       *MoqSetFinalizer_genType
}

// MoqSetFinalizer_genType_anyParams isolates the any params functions of the
// SetFinalizer_genType type
type MoqSetFinalizer_genType_anyParams struct {
	Recorder *MoqSetFinalizer_genType_fnRecorder
}

// NewMoqSetFinalizer_genType creates a new moq of the SetFinalizer_genType
// type
func NewMoqSetFinalizer_genType(scene *moq.Scene, config *moq.Config) *MoqSetFinalizer_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSetFinalizer_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSetFinalizer_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Obj       moq.ParamIndexing
				Finalizer moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Obj       moq.ParamIndexing
			Finalizer moq.ParamIndexing
		}{
			Obj:       moq.ParamIndexByHash,
			Finalizer: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the SetFinalizer_genType type
func (m *MoqSetFinalizer_genType) Mock() SetFinalizer_genType {
	return func(obj interface{}, finalizer interface{}) {
		m.Scene.T.Helper()
		moq := &MoqSetFinalizer_genType_mock{Moq: m}
		moq.Fn(obj, finalizer)
	}
}

func (m *MoqSetFinalizer_genType_mock) Fn(obj interface{}, finalizer interface{}) {
	m.Moq.Scene.T.Helper()
	params := MoqSetFinalizer_genType_params{
		Obj:       obj,
		Finalizer: finalizer,
	}
	var results *MoqSetFinalizer_genType_results
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
		result.DoFn(obj, finalizer)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(obj, finalizer)
	}
	return
}

func (m *MoqSetFinalizer_genType) OnCall(obj interface{}, finalizer interface{}) *MoqSetFinalizer_genType_fnRecorder {
	return &MoqSetFinalizer_genType_fnRecorder{
		Params: MoqSetFinalizer_genType_params{
			Obj:       obj,
			Finalizer: finalizer,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSetFinalizer_genType_fnRecorder) Any() *MoqSetFinalizer_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSetFinalizer_genType_anyParams{Recorder: r}
}

func (a *MoqSetFinalizer_genType_anyParams) Obj() *MoqSetFinalizer_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqSetFinalizer_genType_anyParams) Finalizer() *MoqSetFinalizer_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqSetFinalizer_genType_fnRecorder) Seq() *MoqSetFinalizer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSetFinalizer_genType_fnRecorder) NoSeq() *MoqSetFinalizer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSetFinalizer_genType_fnRecorder) ReturnResults() *MoqSetFinalizer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSetFinalizer_genType_doFn
		DoReturnFn MoqSetFinalizer_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSetFinalizer_genType_fnRecorder) AndDo(fn MoqSetFinalizer_genType_doFn) *MoqSetFinalizer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSetFinalizer_genType_fnRecorder) DoReturnResults(fn MoqSetFinalizer_genType_doReturnFn) *MoqSetFinalizer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSetFinalizer_genType_doFn
		DoReturnFn MoqSetFinalizer_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSetFinalizer_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSetFinalizer_genType_resultsByParams
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
		results = &MoqSetFinalizer_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSetFinalizer_genType_paramsKey]*MoqSetFinalizer_genType_results{},
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
		r.Results = &MoqSetFinalizer_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSetFinalizer_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSetFinalizer_genType_fnRecorder {
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
				DoFn       MoqSetFinalizer_genType_doFn
				DoReturnFn MoqSetFinalizer_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSetFinalizer_genType) PrettyParams(params MoqSetFinalizer_genType_params) string {
	return fmt.Sprintf("SetFinalizer_genType(%#v, %#v)", params.Obj, params.Finalizer)
}

func (m *MoqSetFinalizer_genType) ParamsKey(params MoqSetFinalizer_genType_params, anyParams uint64) MoqSetFinalizer_genType_paramsKey {
	m.Scene.T.Helper()
	var objUsed interface{}
	var objUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Obj == moq.ParamIndexByValue {
			objUsed = params.Obj
		} else {
			objUsedHash = hash.DeepHash(params.Obj)
		}
	}
	var finalizerUsed interface{}
	var finalizerUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Finalizer == moq.ParamIndexByValue {
			finalizerUsed = params.Finalizer
		} else {
			finalizerUsedHash = hash.DeepHash(params.Finalizer)
		}
	}
	return MoqSetFinalizer_genType_paramsKey{
		Params: struct {
			Obj       interface{}
			Finalizer interface{}
		}{
			Obj:       objUsed,
			Finalizer: finalizerUsed,
		},
		Hashes: struct {
			Obj       hash.Hash
			Finalizer hash.Hash
		}{
			Obj:       objUsedHash,
			Finalizer: finalizerUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSetFinalizer_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSetFinalizer_genType) AssertExpectationsMet() {
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
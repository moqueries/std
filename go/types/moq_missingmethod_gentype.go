// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package types

import (
	"fmt"
	"go/types"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// MissingMethod_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type MissingMethod_genType func(V types.Type, T *types.Interface, static bool) (method *types.Func, wrongType bool)

// MoqMissingMethod_genType holds the state of a moq of the
// MissingMethod_genType type
type MoqMissingMethod_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqMissingMethod_genType_mock

	ResultsByParams []MoqMissingMethod_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			V      moq.ParamIndexing
			T      moq.ParamIndexing
			Static moq.ParamIndexing
		}
	}
}

// MoqMissingMethod_genType_mock isolates the mock interface of the
// MissingMethod_genType type
type MoqMissingMethod_genType_mock struct {
	Moq *MoqMissingMethod_genType
}

// MoqMissingMethod_genType_params holds the params of the
// MissingMethod_genType type
type MoqMissingMethod_genType_params struct {
	V      types.Type
	T      *types.Interface
	Static bool
}

// MoqMissingMethod_genType_paramsKey holds the map key params of the
// MissingMethod_genType type
type MoqMissingMethod_genType_paramsKey struct {
	Params struct {
		V      types.Type
		T      *types.Interface
		Static bool
	}
	Hashes struct {
		V      hash.Hash
		T      hash.Hash
		Static hash.Hash
	}
}

// MoqMissingMethod_genType_resultsByParams contains the results for a given
// set of parameters for the MissingMethod_genType type
type MoqMissingMethod_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqMissingMethod_genType_paramsKey]*MoqMissingMethod_genType_results
}

// MoqMissingMethod_genType_doFn defines the type of function needed when
// calling AndDo for the MissingMethod_genType type
type MoqMissingMethod_genType_doFn func(V types.Type, T *types.Interface, static bool)

// MoqMissingMethod_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the MissingMethod_genType type
type MoqMissingMethod_genType_doReturnFn func(V types.Type, T *types.Interface, static bool) (method *types.Func, wrongType bool)

// MoqMissingMethod_genType_results holds the results of the
// MissingMethod_genType type
type MoqMissingMethod_genType_results struct {
	Params  MoqMissingMethod_genType_params
	Results []struct {
		Values *struct {
			Method    *types.Func
			WrongType bool
		}
		Sequence   uint32
		DoFn       MoqMissingMethod_genType_doFn
		DoReturnFn MoqMissingMethod_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqMissingMethod_genType_fnRecorder routes recorded function calls to the
// MoqMissingMethod_genType moq
type MoqMissingMethod_genType_fnRecorder struct {
	Params    MoqMissingMethod_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqMissingMethod_genType_results
	Moq       *MoqMissingMethod_genType
}

// MoqMissingMethod_genType_anyParams isolates the any params functions of the
// MissingMethod_genType type
type MoqMissingMethod_genType_anyParams struct {
	Recorder *MoqMissingMethod_genType_fnRecorder
}

// NewMoqMissingMethod_genType creates a new moq of the MissingMethod_genType
// type
func NewMoqMissingMethod_genType(scene *moq.Scene, config *moq.Config) *MoqMissingMethod_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqMissingMethod_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqMissingMethod_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				V      moq.ParamIndexing
				T      moq.ParamIndexing
				Static moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			V      moq.ParamIndexing
			T      moq.ParamIndexing
			Static moq.ParamIndexing
		}{
			V:      moq.ParamIndexByHash,
			T:      moq.ParamIndexByHash,
			Static: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the MissingMethod_genType type
func (m *MoqMissingMethod_genType) Mock() MissingMethod_genType {
	return func(V types.Type, T *types.Interface, static bool) (_ *types.Func, _ bool) {
		m.Scene.T.Helper()
		moq := &MoqMissingMethod_genType_mock{Moq: m}
		return moq.Fn(V, T, static)
	}
}

func (m *MoqMissingMethod_genType_mock) Fn(V types.Type, T *types.Interface, static bool) (method *types.Func, wrongType bool) {
	m.Moq.Scene.T.Helper()
	params := MoqMissingMethod_genType_params{
		V:      V,
		T:      T,
		Static: static,
	}
	var results *MoqMissingMethod_genType_results
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
		result.DoFn(V, T, static)
	}

	if result.Values != nil {
		method = result.Values.Method
		wrongType = result.Values.WrongType
	}
	if result.DoReturnFn != nil {
		method, wrongType = result.DoReturnFn(V, T, static)
	}
	return
}

func (m *MoqMissingMethod_genType) OnCall(V types.Type, T *types.Interface, static bool) *MoqMissingMethod_genType_fnRecorder {
	return &MoqMissingMethod_genType_fnRecorder{
		Params: MoqMissingMethod_genType_params{
			V:      V,
			T:      T,
			Static: static,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqMissingMethod_genType_fnRecorder) Any() *MoqMissingMethod_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqMissingMethod_genType_anyParams{Recorder: r}
}

func (a *MoqMissingMethod_genType_anyParams) V() *MoqMissingMethod_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqMissingMethod_genType_anyParams) T() *MoqMissingMethod_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqMissingMethod_genType_anyParams) Static() *MoqMissingMethod_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqMissingMethod_genType_fnRecorder) Seq() *MoqMissingMethod_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqMissingMethod_genType_fnRecorder) NoSeq() *MoqMissingMethod_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqMissingMethod_genType_fnRecorder) ReturnResults(method *types.Func, wrongType bool) *MoqMissingMethod_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Method    *types.Func
			WrongType bool
		}
		Sequence   uint32
		DoFn       MoqMissingMethod_genType_doFn
		DoReturnFn MoqMissingMethod_genType_doReturnFn
	}{
		Values: &struct {
			Method    *types.Func
			WrongType bool
		}{
			Method:    method,
			WrongType: wrongType,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqMissingMethod_genType_fnRecorder) AndDo(fn MoqMissingMethod_genType_doFn) *MoqMissingMethod_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqMissingMethod_genType_fnRecorder) DoReturnResults(fn MoqMissingMethod_genType_doReturnFn) *MoqMissingMethod_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Method    *types.Func
			WrongType bool
		}
		Sequence   uint32
		DoFn       MoqMissingMethod_genType_doFn
		DoReturnFn MoqMissingMethod_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqMissingMethod_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqMissingMethod_genType_resultsByParams
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
		results = &MoqMissingMethod_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqMissingMethod_genType_paramsKey]*MoqMissingMethod_genType_results{},
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
		r.Results = &MoqMissingMethod_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqMissingMethod_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqMissingMethod_genType_fnRecorder {
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
					Method    *types.Func
					WrongType bool
				}
				Sequence   uint32
				DoFn       MoqMissingMethod_genType_doFn
				DoReturnFn MoqMissingMethod_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqMissingMethod_genType) PrettyParams(params MoqMissingMethod_genType_params) string {
	return fmt.Sprintf("MissingMethod_genType(%#v, %#v, %#v)", params.V, params.T, params.Static)
}

func (m *MoqMissingMethod_genType) ParamsKey(params MoqMissingMethod_genType_params, anyParams uint64) MoqMissingMethod_genType_paramsKey {
	m.Scene.T.Helper()
	var VUsed types.Type
	var VUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.V == moq.ParamIndexByValue {
			VUsed = params.V
		} else {
			VUsedHash = hash.DeepHash(params.V)
		}
	}
	var TUsed *types.Interface
	var TUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.T == moq.ParamIndexByValue {
			TUsed = params.T
		} else {
			TUsedHash = hash.DeepHash(params.T)
		}
	}
	var staticUsed bool
	var staticUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Static == moq.ParamIndexByValue {
			staticUsed = params.Static
		} else {
			staticUsedHash = hash.DeepHash(params.Static)
		}
	}
	return MoqMissingMethod_genType_paramsKey{
		Params: struct {
			V      types.Type
			T      *types.Interface
			Static bool
		}{
			V:      VUsed,
			T:      TUsed,
			Static: staticUsed,
		},
		Hashes: struct {
			V      hash.Hash
			T      hash.Hash
			Static hash.Hash
		}{
			V:      VUsedHash,
			T:      TUsedHash,
			Static: staticUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqMissingMethod_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqMissingMethod_genType) AssertExpectationsMet() {
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

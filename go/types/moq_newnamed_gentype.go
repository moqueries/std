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

// NewNamed_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type NewNamed_genType func(obj *types.TypeName, underlying types.Type, methods []*types.Func) *types.Named

// MoqNewNamed_genType holds the state of a moq of the NewNamed_genType type
type MoqNewNamed_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNewNamed_genType_mock

	ResultsByParams []MoqNewNamed_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Obj        moq.ParamIndexing
			Underlying moq.ParamIndexing
			Methods    moq.ParamIndexing
		}
	}
}

// MoqNewNamed_genType_mock isolates the mock interface of the NewNamed_genType
// type
type MoqNewNamed_genType_mock struct {
	Moq *MoqNewNamed_genType
}

// MoqNewNamed_genType_params holds the params of the NewNamed_genType type
type MoqNewNamed_genType_params struct {
	Obj        *types.TypeName
	Underlying types.Type
	Methods    []*types.Func
}

// MoqNewNamed_genType_paramsKey holds the map key params of the
// NewNamed_genType type
type MoqNewNamed_genType_paramsKey struct {
	Params struct {
		Obj        *types.TypeName
		Underlying types.Type
	}
	Hashes struct {
		Obj        hash.Hash
		Underlying hash.Hash
		Methods    hash.Hash
	}
}

// MoqNewNamed_genType_resultsByParams contains the results for a given set of
// parameters for the NewNamed_genType type
type MoqNewNamed_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNewNamed_genType_paramsKey]*MoqNewNamed_genType_results
}

// MoqNewNamed_genType_doFn defines the type of function needed when calling
// AndDo for the NewNamed_genType type
type MoqNewNamed_genType_doFn func(obj *types.TypeName, underlying types.Type, methods []*types.Func)

// MoqNewNamed_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the NewNamed_genType type
type MoqNewNamed_genType_doReturnFn func(obj *types.TypeName, underlying types.Type, methods []*types.Func) *types.Named

// MoqNewNamed_genType_results holds the results of the NewNamed_genType type
type MoqNewNamed_genType_results struct {
	Params  MoqNewNamed_genType_params
	Results []struct {
		Values *struct {
			Result1 *types.Named
		}
		Sequence   uint32
		DoFn       MoqNewNamed_genType_doFn
		DoReturnFn MoqNewNamed_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNewNamed_genType_fnRecorder routes recorded function calls to the
// MoqNewNamed_genType moq
type MoqNewNamed_genType_fnRecorder struct {
	Params    MoqNewNamed_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNewNamed_genType_results
	Moq       *MoqNewNamed_genType
}

// MoqNewNamed_genType_anyParams isolates the any params functions of the
// NewNamed_genType type
type MoqNewNamed_genType_anyParams struct {
	Recorder *MoqNewNamed_genType_fnRecorder
}

// NewMoqNewNamed_genType creates a new moq of the NewNamed_genType type
func NewMoqNewNamed_genType(scene *moq.Scene, config *moq.Config) *MoqNewNamed_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNewNamed_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNewNamed_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Obj        moq.ParamIndexing
				Underlying moq.ParamIndexing
				Methods    moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Obj        moq.ParamIndexing
			Underlying moq.ParamIndexing
			Methods    moq.ParamIndexing
		}{
			Obj:        moq.ParamIndexByHash,
			Underlying: moq.ParamIndexByHash,
			Methods:    moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the NewNamed_genType type
func (m *MoqNewNamed_genType) Mock() NewNamed_genType {
	return func(obj *types.TypeName, underlying types.Type, methods []*types.Func) *types.Named {
		m.Scene.T.Helper()
		moq := &MoqNewNamed_genType_mock{Moq: m}
		return moq.Fn(obj, underlying, methods)
	}
}

func (m *MoqNewNamed_genType_mock) Fn(obj *types.TypeName, underlying types.Type, methods []*types.Func) (result1 *types.Named) {
	m.Moq.Scene.T.Helper()
	params := MoqNewNamed_genType_params{
		Obj:        obj,
		Underlying: underlying,
		Methods:    methods,
	}
	var results *MoqNewNamed_genType_results
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
		result.DoFn(obj, underlying, methods)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(obj, underlying, methods)
	}
	return
}

func (m *MoqNewNamed_genType) OnCall(obj *types.TypeName, underlying types.Type, methods []*types.Func) *MoqNewNamed_genType_fnRecorder {
	return &MoqNewNamed_genType_fnRecorder{
		Params: MoqNewNamed_genType_params{
			Obj:        obj,
			Underlying: underlying,
			Methods:    methods,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNewNamed_genType_fnRecorder) Any() *MoqNewNamed_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNewNamed_genType_anyParams{Recorder: r}
}

func (a *MoqNewNamed_genType_anyParams) Obj() *MoqNewNamed_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqNewNamed_genType_anyParams) Underlying() *MoqNewNamed_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqNewNamed_genType_anyParams) Methods() *MoqNewNamed_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqNewNamed_genType_fnRecorder) Seq() *MoqNewNamed_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNewNamed_genType_fnRecorder) NoSeq() *MoqNewNamed_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNewNamed_genType_fnRecorder) ReturnResults(result1 *types.Named) *MoqNewNamed_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *types.Named
		}
		Sequence   uint32
		DoFn       MoqNewNamed_genType_doFn
		DoReturnFn MoqNewNamed_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *types.Named
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNewNamed_genType_fnRecorder) AndDo(fn MoqNewNamed_genType_doFn) *MoqNewNamed_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNewNamed_genType_fnRecorder) DoReturnResults(fn MoqNewNamed_genType_doReturnFn) *MoqNewNamed_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *types.Named
		}
		Sequence   uint32
		DoFn       MoqNewNamed_genType_doFn
		DoReturnFn MoqNewNamed_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNewNamed_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNewNamed_genType_resultsByParams
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
		results = &MoqNewNamed_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNewNamed_genType_paramsKey]*MoqNewNamed_genType_results{},
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
		r.Results = &MoqNewNamed_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNewNamed_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNewNamed_genType_fnRecorder {
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
					Result1 *types.Named
				}
				Sequence   uint32
				DoFn       MoqNewNamed_genType_doFn
				DoReturnFn MoqNewNamed_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNewNamed_genType) PrettyParams(params MoqNewNamed_genType_params) string {
	return fmt.Sprintf("NewNamed_genType(%#v, %#v, %#v)", params.Obj, params.Underlying, params.Methods)
}

func (m *MoqNewNamed_genType) ParamsKey(params MoqNewNamed_genType_params, anyParams uint64) MoqNewNamed_genType_paramsKey {
	m.Scene.T.Helper()
	var objUsed *types.TypeName
	var objUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Obj == moq.ParamIndexByValue {
			objUsed = params.Obj
		} else {
			objUsedHash = hash.DeepHash(params.Obj)
		}
	}
	var underlyingUsed types.Type
	var underlyingUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Underlying == moq.ParamIndexByValue {
			underlyingUsed = params.Underlying
		} else {
			underlyingUsedHash = hash.DeepHash(params.Underlying)
		}
	}
	var methodsUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Methods == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The methods parameter can't be indexed by value")
		}
		methodsUsedHash = hash.DeepHash(params.Methods)
	}
	return MoqNewNamed_genType_paramsKey{
		Params: struct {
			Obj        *types.TypeName
			Underlying types.Type
		}{
			Obj:        objUsed,
			Underlying: underlyingUsed,
		},
		Hashes: struct {
			Obj        hash.Hash
			Underlying hash.Hash
			Methods    hash.Hash
		}{
			Obj:        objUsedHash,
			Underlying: underlyingUsedHash,
			Methods:    methodsUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNewNamed_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNewNamed_genType) AssertExpectationsMet() {
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

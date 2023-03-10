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

// ConvertibleTo_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type ConvertibleTo_genType func(V, T types.Type) bool

// MoqConvertibleTo_genType holds the state of a moq of the
// ConvertibleTo_genType type
type MoqConvertibleTo_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqConvertibleTo_genType_mock

	ResultsByParams []MoqConvertibleTo_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			V moq.ParamIndexing
			T moq.ParamIndexing
		}
	}
}

// MoqConvertibleTo_genType_mock isolates the mock interface of the
// ConvertibleTo_genType type
type MoqConvertibleTo_genType_mock struct {
	Moq *MoqConvertibleTo_genType
}

// MoqConvertibleTo_genType_params holds the params of the
// ConvertibleTo_genType type
type MoqConvertibleTo_genType_params struct{ V, T types.Type }

// MoqConvertibleTo_genType_paramsKey holds the map key params of the
// ConvertibleTo_genType type
type MoqConvertibleTo_genType_paramsKey struct {
	Params struct{ V, T types.Type }
	Hashes struct{ V, T hash.Hash }
}

// MoqConvertibleTo_genType_resultsByParams contains the results for a given
// set of parameters for the ConvertibleTo_genType type
type MoqConvertibleTo_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqConvertibleTo_genType_paramsKey]*MoqConvertibleTo_genType_results
}

// MoqConvertibleTo_genType_doFn defines the type of function needed when
// calling AndDo for the ConvertibleTo_genType type
type MoqConvertibleTo_genType_doFn func(V, T types.Type)

// MoqConvertibleTo_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ConvertibleTo_genType type
type MoqConvertibleTo_genType_doReturnFn func(V, T types.Type) bool

// MoqConvertibleTo_genType_results holds the results of the
// ConvertibleTo_genType type
type MoqConvertibleTo_genType_results struct {
	Params  MoqConvertibleTo_genType_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqConvertibleTo_genType_doFn
		DoReturnFn MoqConvertibleTo_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqConvertibleTo_genType_fnRecorder routes recorded function calls to the
// MoqConvertibleTo_genType moq
type MoqConvertibleTo_genType_fnRecorder struct {
	Params    MoqConvertibleTo_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqConvertibleTo_genType_results
	Moq       *MoqConvertibleTo_genType
}

// MoqConvertibleTo_genType_anyParams isolates the any params functions of the
// ConvertibleTo_genType type
type MoqConvertibleTo_genType_anyParams struct {
	Recorder *MoqConvertibleTo_genType_fnRecorder
}

// NewMoqConvertibleTo_genType creates a new moq of the ConvertibleTo_genType
// type
func NewMoqConvertibleTo_genType(scene *moq.Scene, config *moq.Config) *MoqConvertibleTo_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqConvertibleTo_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqConvertibleTo_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				V moq.ParamIndexing
				T moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			V moq.ParamIndexing
			T moq.ParamIndexing
		}{
			V: moq.ParamIndexByHash,
			T: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ConvertibleTo_genType type
func (m *MoqConvertibleTo_genType) Mock() ConvertibleTo_genType {
	return func(V, T types.Type) bool {
		m.Scene.T.Helper()
		moq := &MoqConvertibleTo_genType_mock{Moq: m}
		return moq.Fn(V, T)
	}
}

func (m *MoqConvertibleTo_genType_mock) Fn(V, T types.Type) (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqConvertibleTo_genType_params{
		V: V,
		T: T,
	}
	var results *MoqConvertibleTo_genType_results
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
		result.DoFn(V, T)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(V, T)
	}
	return
}

func (m *MoqConvertibleTo_genType) OnCall(V, T types.Type) *MoqConvertibleTo_genType_fnRecorder {
	return &MoqConvertibleTo_genType_fnRecorder{
		Params: MoqConvertibleTo_genType_params{
			V: V,
			T: T,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqConvertibleTo_genType_fnRecorder) Any() *MoqConvertibleTo_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqConvertibleTo_genType_anyParams{Recorder: r}
}

func (a *MoqConvertibleTo_genType_anyParams) V() *MoqConvertibleTo_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqConvertibleTo_genType_anyParams) T() *MoqConvertibleTo_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqConvertibleTo_genType_fnRecorder) Seq() *MoqConvertibleTo_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqConvertibleTo_genType_fnRecorder) NoSeq() *MoqConvertibleTo_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqConvertibleTo_genType_fnRecorder) ReturnResults(result1 bool) *MoqConvertibleTo_genType_fnRecorder {
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
		DoFn       MoqConvertibleTo_genType_doFn
		DoReturnFn MoqConvertibleTo_genType_doReturnFn
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

func (r *MoqConvertibleTo_genType_fnRecorder) AndDo(fn MoqConvertibleTo_genType_doFn) *MoqConvertibleTo_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqConvertibleTo_genType_fnRecorder) DoReturnResults(fn MoqConvertibleTo_genType_doReturnFn) *MoqConvertibleTo_genType_fnRecorder {
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
		DoFn       MoqConvertibleTo_genType_doFn
		DoReturnFn MoqConvertibleTo_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqConvertibleTo_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqConvertibleTo_genType_resultsByParams
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
		results = &MoqConvertibleTo_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqConvertibleTo_genType_paramsKey]*MoqConvertibleTo_genType_results{},
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
		r.Results = &MoqConvertibleTo_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqConvertibleTo_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqConvertibleTo_genType_fnRecorder {
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
				DoFn       MoqConvertibleTo_genType_doFn
				DoReturnFn MoqConvertibleTo_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqConvertibleTo_genType) PrettyParams(params MoqConvertibleTo_genType_params) string {
	return fmt.Sprintf("ConvertibleTo_genType(%#v, %#v)", params.V, params.T)
}

func (m *MoqConvertibleTo_genType) ParamsKey(params MoqConvertibleTo_genType_params, anyParams uint64) MoqConvertibleTo_genType_paramsKey {
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
	var TUsed types.Type
	var TUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.T == moq.ParamIndexByValue {
			TUsed = params.T
		} else {
			TUsedHash = hash.DeepHash(params.T)
		}
	}
	return MoqConvertibleTo_genType_paramsKey{
		Params: struct{ V, T types.Type }{
			V: VUsed,
			T: TUsed,
		},
		Hashes: struct{ V, T hash.Hash }{
			V: VUsedHash,
			T: TUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqConvertibleTo_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqConvertibleTo_genType) AssertExpectationsMet() {
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

// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package reflect

import (
	"fmt"
	"math/bits"
	"reflect"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// PtrTo_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type PtrTo_genType func(t reflect.Type) reflect.Type

// MoqPtrTo_genType holds the state of a moq of the PtrTo_genType type
type MoqPtrTo_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqPtrTo_genType_mock

	ResultsByParams []MoqPtrTo_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			T moq.ParamIndexing
		}
	}
}

// MoqPtrTo_genType_mock isolates the mock interface of the PtrTo_genType type
type MoqPtrTo_genType_mock struct {
	Moq *MoqPtrTo_genType
}

// MoqPtrTo_genType_params holds the params of the PtrTo_genType type
type MoqPtrTo_genType_params struct{ T reflect.Type }

// MoqPtrTo_genType_paramsKey holds the map key params of the PtrTo_genType
// type
type MoqPtrTo_genType_paramsKey struct {
	Params struct{ T reflect.Type }
	Hashes struct{ T hash.Hash }
}

// MoqPtrTo_genType_resultsByParams contains the results for a given set of
// parameters for the PtrTo_genType type
type MoqPtrTo_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqPtrTo_genType_paramsKey]*MoqPtrTo_genType_results
}

// MoqPtrTo_genType_doFn defines the type of function needed when calling AndDo
// for the PtrTo_genType type
type MoqPtrTo_genType_doFn func(t reflect.Type)

// MoqPtrTo_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the PtrTo_genType type
type MoqPtrTo_genType_doReturnFn func(t reflect.Type) reflect.Type

// MoqPtrTo_genType_results holds the results of the PtrTo_genType type
type MoqPtrTo_genType_results struct {
	Params  MoqPtrTo_genType_params
	Results []struct {
		Values *struct {
			Result1 reflect.Type
		}
		Sequence   uint32
		DoFn       MoqPtrTo_genType_doFn
		DoReturnFn MoqPtrTo_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqPtrTo_genType_fnRecorder routes recorded function calls to the
// MoqPtrTo_genType moq
type MoqPtrTo_genType_fnRecorder struct {
	Params    MoqPtrTo_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqPtrTo_genType_results
	Moq       *MoqPtrTo_genType
}

// MoqPtrTo_genType_anyParams isolates the any params functions of the
// PtrTo_genType type
type MoqPtrTo_genType_anyParams struct {
	Recorder *MoqPtrTo_genType_fnRecorder
}

// NewMoqPtrTo_genType creates a new moq of the PtrTo_genType type
func NewMoqPtrTo_genType(scene *moq.Scene, config *moq.Config) *MoqPtrTo_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqPtrTo_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqPtrTo_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				T moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			T moq.ParamIndexing
		}{
			T: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the PtrTo_genType type
func (m *MoqPtrTo_genType) Mock() PtrTo_genType {
	return func(t reflect.Type) reflect.Type {
		m.Scene.T.Helper()
		moq := &MoqPtrTo_genType_mock{Moq: m}
		return moq.Fn(t)
	}
}

func (m *MoqPtrTo_genType_mock) Fn(t reflect.Type) (result1 reflect.Type) {
	m.Moq.Scene.T.Helper()
	params := MoqPtrTo_genType_params{
		T: t,
	}
	var results *MoqPtrTo_genType_results
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
		result.DoFn(t)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(t)
	}
	return
}

func (m *MoqPtrTo_genType) OnCall(t reflect.Type) *MoqPtrTo_genType_fnRecorder {
	return &MoqPtrTo_genType_fnRecorder{
		Params: MoqPtrTo_genType_params{
			T: t,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqPtrTo_genType_fnRecorder) Any() *MoqPtrTo_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqPtrTo_genType_anyParams{Recorder: r}
}

func (a *MoqPtrTo_genType_anyParams) T() *MoqPtrTo_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqPtrTo_genType_fnRecorder) Seq() *MoqPtrTo_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqPtrTo_genType_fnRecorder) NoSeq() *MoqPtrTo_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqPtrTo_genType_fnRecorder) ReturnResults(result1 reflect.Type) *MoqPtrTo_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 reflect.Type
		}
		Sequence   uint32
		DoFn       MoqPtrTo_genType_doFn
		DoReturnFn MoqPtrTo_genType_doReturnFn
	}{
		Values: &struct {
			Result1 reflect.Type
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqPtrTo_genType_fnRecorder) AndDo(fn MoqPtrTo_genType_doFn) *MoqPtrTo_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqPtrTo_genType_fnRecorder) DoReturnResults(fn MoqPtrTo_genType_doReturnFn) *MoqPtrTo_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 reflect.Type
		}
		Sequence   uint32
		DoFn       MoqPtrTo_genType_doFn
		DoReturnFn MoqPtrTo_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqPtrTo_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqPtrTo_genType_resultsByParams
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
		results = &MoqPtrTo_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqPtrTo_genType_paramsKey]*MoqPtrTo_genType_results{},
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
		r.Results = &MoqPtrTo_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqPtrTo_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqPtrTo_genType_fnRecorder {
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
					Result1 reflect.Type
				}
				Sequence   uint32
				DoFn       MoqPtrTo_genType_doFn
				DoReturnFn MoqPtrTo_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqPtrTo_genType) PrettyParams(params MoqPtrTo_genType_params) string {
	return fmt.Sprintf("PtrTo_genType(%#v)", params.T)
}

func (m *MoqPtrTo_genType) ParamsKey(params MoqPtrTo_genType_params, anyParams uint64) MoqPtrTo_genType_paramsKey {
	m.Scene.T.Helper()
	var tUsed reflect.Type
	var tUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.T == moq.ParamIndexByValue {
			tUsed = params.T
		} else {
			tUsedHash = hash.DeepHash(params.T)
		}
	}
	return MoqPtrTo_genType_paramsKey{
		Params: struct{ T reflect.Type }{
			T: tUsed,
		},
		Hashes: struct{ T hash.Hash }{
			T: tUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqPtrTo_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqPtrTo_genType) AssertExpectationsMet() {
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

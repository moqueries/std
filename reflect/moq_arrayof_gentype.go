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

// ArrayOf_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type ArrayOf_genType func(length int, elem reflect.Type) reflect.Type

// MoqArrayOf_genType holds the state of a moq of the ArrayOf_genType type
type MoqArrayOf_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqArrayOf_genType_mock

	ResultsByParams []MoqArrayOf_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Length moq.ParamIndexing
			Elem   moq.ParamIndexing
		}
	}
}

// MoqArrayOf_genType_mock isolates the mock interface of the ArrayOf_genType
// type
type MoqArrayOf_genType_mock struct {
	Moq *MoqArrayOf_genType
}

// MoqArrayOf_genType_params holds the params of the ArrayOf_genType type
type MoqArrayOf_genType_params struct {
	Length int
	Elem   reflect.Type
}

// MoqArrayOf_genType_paramsKey holds the map key params of the ArrayOf_genType
// type
type MoqArrayOf_genType_paramsKey struct {
	Params struct {
		Length int
		Elem   reflect.Type
	}
	Hashes struct {
		Length hash.Hash
		Elem   hash.Hash
	}
}

// MoqArrayOf_genType_resultsByParams contains the results for a given set of
// parameters for the ArrayOf_genType type
type MoqArrayOf_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqArrayOf_genType_paramsKey]*MoqArrayOf_genType_results
}

// MoqArrayOf_genType_doFn defines the type of function needed when calling
// AndDo for the ArrayOf_genType type
type MoqArrayOf_genType_doFn func(length int, elem reflect.Type)

// MoqArrayOf_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ArrayOf_genType type
type MoqArrayOf_genType_doReturnFn func(length int, elem reflect.Type) reflect.Type

// MoqArrayOf_genType_results holds the results of the ArrayOf_genType type
type MoqArrayOf_genType_results struct {
	Params  MoqArrayOf_genType_params
	Results []struct {
		Values *struct {
			Result1 reflect.Type
		}
		Sequence   uint32
		DoFn       MoqArrayOf_genType_doFn
		DoReturnFn MoqArrayOf_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqArrayOf_genType_fnRecorder routes recorded function calls to the
// MoqArrayOf_genType moq
type MoqArrayOf_genType_fnRecorder struct {
	Params    MoqArrayOf_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqArrayOf_genType_results
	Moq       *MoqArrayOf_genType
}

// MoqArrayOf_genType_anyParams isolates the any params functions of the
// ArrayOf_genType type
type MoqArrayOf_genType_anyParams struct {
	Recorder *MoqArrayOf_genType_fnRecorder
}

// NewMoqArrayOf_genType creates a new moq of the ArrayOf_genType type
func NewMoqArrayOf_genType(scene *moq.Scene, config *moq.Config) *MoqArrayOf_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqArrayOf_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqArrayOf_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Length moq.ParamIndexing
				Elem   moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Length moq.ParamIndexing
			Elem   moq.ParamIndexing
		}{
			Length: moq.ParamIndexByValue,
			Elem:   moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ArrayOf_genType type
func (m *MoqArrayOf_genType) Mock() ArrayOf_genType {
	return func(length int, elem reflect.Type) reflect.Type {
		m.Scene.T.Helper()
		moq := &MoqArrayOf_genType_mock{Moq: m}
		return moq.Fn(length, elem)
	}
}

func (m *MoqArrayOf_genType_mock) Fn(length int, elem reflect.Type) (result1 reflect.Type) {
	m.Moq.Scene.T.Helper()
	params := MoqArrayOf_genType_params{
		Length: length,
		Elem:   elem,
	}
	var results *MoqArrayOf_genType_results
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
		result.DoFn(length, elem)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(length, elem)
	}
	return
}

func (m *MoqArrayOf_genType) OnCall(length int, elem reflect.Type) *MoqArrayOf_genType_fnRecorder {
	return &MoqArrayOf_genType_fnRecorder{
		Params: MoqArrayOf_genType_params{
			Length: length,
			Elem:   elem,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqArrayOf_genType_fnRecorder) Any() *MoqArrayOf_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqArrayOf_genType_anyParams{Recorder: r}
}

func (a *MoqArrayOf_genType_anyParams) Length() *MoqArrayOf_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqArrayOf_genType_anyParams) Elem() *MoqArrayOf_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqArrayOf_genType_fnRecorder) Seq() *MoqArrayOf_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqArrayOf_genType_fnRecorder) NoSeq() *MoqArrayOf_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqArrayOf_genType_fnRecorder) ReturnResults(result1 reflect.Type) *MoqArrayOf_genType_fnRecorder {
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
		DoFn       MoqArrayOf_genType_doFn
		DoReturnFn MoqArrayOf_genType_doReturnFn
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

func (r *MoqArrayOf_genType_fnRecorder) AndDo(fn MoqArrayOf_genType_doFn) *MoqArrayOf_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqArrayOf_genType_fnRecorder) DoReturnResults(fn MoqArrayOf_genType_doReturnFn) *MoqArrayOf_genType_fnRecorder {
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
		DoFn       MoqArrayOf_genType_doFn
		DoReturnFn MoqArrayOf_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqArrayOf_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqArrayOf_genType_resultsByParams
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
		results = &MoqArrayOf_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqArrayOf_genType_paramsKey]*MoqArrayOf_genType_results{},
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
		r.Results = &MoqArrayOf_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqArrayOf_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqArrayOf_genType_fnRecorder {
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
				DoFn       MoqArrayOf_genType_doFn
				DoReturnFn MoqArrayOf_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqArrayOf_genType) PrettyParams(params MoqArrayOf_genType_params) string {
	return fmt.Sprintf("ArrayOf_genType(%#v, %#v)", params.Length, params.Elem)
}

func (m *MoqArrayOf_genType) ParamsKey(params MoqArrayOf_genType_params, anyParams uint64) MoqArrayOf_genType_paramsKey {
	m.Scene.T.Helper()
	var lengthUsed int
	var lengthUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Length == moq.ParamIndexByValue {
			lengthUsed = params.Length
		} else {
			lengthUsedHash = hash.DeepHash(params.Length)
		}
	}
	var elemUsed reflect.Type
	var elemUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Elem == moq.ParamIndexByValue {
			elemUsed = params.Elem
		} else {
			elemUsedHash = hash.DeepHash(params.Elem)
		}
	}
	return MoqArrayOf_genType_paramsKey{
		Params: struct {
			Length int
			Elem   reflect.Type
		}{
			Length: lengthUsed,
			Elem:   elemUsed,
		},
		Hashes: struct {
			Length hash.Hash
			Elem   hash.Hash
		}{
			Length: lengthUsedHash,
			Elem:   elemUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqArrayOf_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqArrayOf_genType) AssertExpectationsMet() {
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

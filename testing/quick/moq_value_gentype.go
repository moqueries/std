// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package quick

import (
	"fmt"
	"math/bits"
	"math/rand"
	"reflect"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Value_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Value_genType func(t reflect.Type, rand *rand.Rand) (value reflect.Value, ok bool)

// MoqValue_genType holds the state of a moq of the Value_genType type
type MoqValue_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqValue_genType_mock

	ResultsByParams []MoqValue_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			T    moq.ParamIndexing
			Rand moq.ParamIndexing
		}
	}
}

// MoqValue_genType_mock isolates the mock interface of the Value_genType type
type MoqValue_genType_mock struct {
	Moq *MoqValue_genType
}

// MoqValue_genType_params holds the params of the Value_genType type
type MoqValue_genType_params struct {
	T    reflect.Type
	Rand *rand.Rand
}

// MoqValue_genType_paramsKey holds the map key params of the Value_genType
// type
type MoqValue_genType_paramsKey struct {
	Params struct {
		T    reflect.Type
		Rand *rand.Rand
	}
	Hashes struct {
		T    hash.Hash
		Rand hash.Hash
	}
}

// MoqValue_genType_resultsByParams contains the results for a given set of
// parameters for the Value_genType type
type MoqValue_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqValue_genType_paramsKey]*MoqValue_genType_results
}

// MoqValue_genType_doFn defines the type of function needed when calling AndDo
// for the Value_genType type
type MoqValue_genType_doFn func(t reflect.Type, rand *rand.Rand)

// MoqValue_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Value_genType type
type MoqValue_genType_doReturnFn func(t reflect.Type, rand *rand.Rand) (value reflect.Value, ok bool)

// MoqValue_genType_results holds the results of the Value_genType type
type MoqValue_genType_results struct {
	Params  MoqValue_genType_params
	Results []struct {
		Values *struct {
			Value reflect.Value
			Ok    bool
		}
		Sequence   uint32
		DoFn       MoqValue_genType_doFn
		DoReturnFn MoqValue_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqValue_genType_fnRecorder routes recorded function calls to the
// MoqValue_genType moq
type MoqValue_genType_fnRecorder struct {
	Params    MoqValue_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqValue_genType_results
	Moq       *MoqValue_genType
}

// MoqValue_genType_anyParams isolates the any params functions of the
// Value_genType type
type MoqValue_genType_anyParams struct {
	Recorder *MoqValue_genType_fnRecorder
}

// NewMoqValue_genType creates a new moq of the Value_genType type
func NewMoqValue_genType(scene *moq.Scene, config *moq.Config) *MoqValue_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqValue_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqValue_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				T    moq.ParamIndexing
				Rand moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			T    moq.ParamIndexing
			Rand moq.ParamIndexing
		}{
			T:    moq.ParamIndexByHash,
			Rand: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Value_genType type
func (m *MoqValue_genType) Mock() Value_genType {
	return func(t reflect.Type, rand *rand.Rand) (_ reflect.Value, _ bool) {
		m.Scene.T.Helper()
		moq := &MoqValue_genType_mock{Moq: m}
		return moq.Fn(t, rand)
	}
}

func (m *MoqValue_genType_mock) Fn(t reflect.Type, rand *rand.Rand) (value reflect.Value, ok bool) {
	m.Moq.Scene.T.Helper()
	params := MoqValue_genType_params{
		T:    t,
		Rand: rand,
	}
	var results *MoqValue_genType_results
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
		result.DoFn(t, rand)
	}

	if result.Values != nil {
		value = result.Values.Value
		ok = result.Values.Ok
	}
	if result.DoReturnFn != nil {
		value, ok = result.DoReturnFn(t, rand)
	}
	return
}

func (m *MoqValue_genType) OnCall(t reflect.Type, rand *rand.Rand) *MoqValue_genType_fnRecorder {
	return &MoqValue_genType_fnRecorder{
		Params: MoqValue_genType_params{
			T:    t,
			Rand: rand,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqValue_genType_fnRecorder) Any() *MoqValue_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqValue_genType_anyParams{Recorder: r}
}

func (a *MoqValue_genType_anyParams) T() *MoqValue_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqValue_genType_anyParams) Rand() *MoqValue_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqValue_genType_fnRecorder) Seq() *MoqValue_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqValue_genType_fnRecorder) NoSeq() *MoqValue_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqValue_genType_fnRecorder) ReturnResults(value reflect.Value, ok bool) *MoqValue_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Value reflect.Value
			Ok    bool
		}
		Sequence   uint32
		DoFn       MoqValue_genType_doFn
		DoReturnFn MoqValue_genType_doReturnFn
	}{
		Values: &struct {
			Value reflect.Value
			Ok    bool
		}{
			Value: value,
			Ok:    ok,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqValue_genType_fnRecorder) AndDo(fn MoqValue_genType_doFn) *MoqValue_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqValue_genType_fnRecorder) DoReturnResults(fn MoqValue_genType_doReturnFn) *MoqValue_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Value reflect.Value
			Ok    bool
		}
		Sequence   uint32
		DoFn       MoqValue_genType_doFn
		DoReturnFn MoqValue_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqValue_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqValue_genType_resultsByParams
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
		results = &MoqValue_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqValue_genType_paramsKey]*MoqValue_genType_results{},
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
		r.Results = &MoqValue_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqValue_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqValue_genType_fnRecorder {
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
					Value reflect.Value
					Ok    bool
				}
				Sequence   uint32
				DoFn       MoqValue_genType_doFn
				DoReturnFn MoqValue_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqValue_genType) PrettyParams(params MoqValue_genType_params) string {
	return fmt.Sprintf("Value_genType(%#v, %#v)", params.T, params.Rand)
}

func (m *MoqValue_genType) ParamsKey(params MoqValue_genType_params, anyParams uint64) MoqValue_genType_paramsKey {
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
	var randUsed *rand.Rand
	var randUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Rand == moq.ParamIndexByValue {
			randUsed = params.Rand
		} else {
			randUsedHash = hash.DeepHash(params.Rand)
		}
	}
	return MoqValue_genType_paramsKey{
		Params: struct {
			T    reflect.Type
			Rand *rand.Rand
		}{
			T:    tUsed,
			Rand: randUsed,
		},
		Hashes: struct {
			T    hash.Hash
			Rand hash.Hash
		}{
			T:    tUsedHash,
			Rand: randUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqValue_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqValue_genType) AssertExpectationsMet() {
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

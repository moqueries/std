// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package log

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Panic_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Panic_genType func(v ...interface{})

// MoqPanic_genType holds the state of a moq of the Panic_genType type
type MoqPanic_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqPanic_genType_mock

	ResultsByParams []MoqPanic_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			V moq.ParamIndexing
		}
	}
}

// MoqPanic_genType_mock isolates the mock interface of the Panic_genType type
type MoqPanic_genType_mock struct {
	Moq *MoqPanic_genType
}

// MoqPanic_genType_params holds the params of the Panic_genType type
type MoqPanic_genType_params struct{ V []interface{} }

// MoqPanic_genType_paramsKey holds the map key params of the Panic_genType
// type
type MoqPanic_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ V hash.Hash }
}

// MoqPanic_genType_resultsByParams contains the results for a given set of
// parameters for the Panic_genType type
type MoqPanic_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqPanic_genType_paramsKey]*MoqPanic_genType_results
}

// MoqPanic_genType_doFn defines the type of function needed when calling AndDo
// for the Panic_genType type
type MoqPanic_genType_doFn func(v ...interface{})

// MoqPanic_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Panic_genType type
type MoqPanic_genType_doReturnFn func(v ...interface{})

// MoqPanic_genType_results holds the results of the Panic_genType type
type MoqPanic_genType_results struct {
	Params  MoqPanic_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqPanic_genType_doFn
		DoReturnFn MoqPanic_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqPanic_genType_fnRecorder routes recorded function calls to the
// MoqPanic_genType moq
type MoqPanic_genType_fnRecorder struct {
	Params    MoqPanic_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqPanic_genType_results
	Moq       *MoqPanic_genType
}

// MoqPanic_genType_anyParams isolates the any params functions of the
// Panic_genType type
type MoqPanic_genType_anyParams struct {
	Recorder *MoqPanic_genType_fnRecorder
}

// NewMoqPanic_genType creates a new moq of the Panic_genType type
func NewMoqPanic_genType(scene *moq.Scene, config *moq.Config) *MoqPanic_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqPanic_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqPanic_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				V moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			V moq.ParamIndexing
		}{
			V: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Panic_genType type
func (m *MoqPanic_genType) Mock() Panic_genType {
	return func(v ...interface{}) { m.Scene.T.Helper(); moq := &MoqPanic_genType_mock{Moq: m}; moq.Fn(v...) }
}

func (m *MoqPanic_genType_mock) Fn(v ...interface{}) {
	m.Moq.Scene.T.Helper()
	params := MoqPanic_genType_params{
		V: v,
	}
	var results *MoqPanic_genType_results
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
		result.DoFn(v...)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(v...)
	}
	return
}

func (m *MoqPanic_genType) OnCall(v ...interface{}) *MoqPanic_genType_fnRecorder {
	return &MoqPanic_genType_fnRecorder{
		Params: MoqPanic_genType_params{
			V: v,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqPanic_genType_fnRecorder) Any() *MoqPanic_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqPanic_genType_anyParams{Recorder: r}
}

func (a *MoqPanic_genType_anyParams) V() *MoqPanic_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqPanic_genType_fnRecorder) Seq() *MoqPanic_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqPanic_genType_fnRecorder) NoSeq() *MoqPanic_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqPanic_genType_fnRecorder) ReturnResults() *MoqPanic_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqPanic_genType_doFn
		DoReturnFn MoqPanic_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqPanic_genType_fnRecorder) AndDo(fn MoqPanic_genType_doFn) *MoqPanic_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqPanic_genType_fnRecorder) DoReturnResults(fn MoqPanic_genType_doReturnFn) *MoqPanic_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqPanic_genType_doFn
		DoReturnFn MoqPanic_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqPanic_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqPanic_genType_resultsByParams
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
		results = &MoqPanic_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqPanic_genType_paramsKey]*MoqPanic_genType_results{},
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
		r.Results = &MoqPanic_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqPanic_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqPanic_genType_fnRecorder {
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
				DoFn       MoqPanic_genType_doFn
				DoReturnFn MoqPanic_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqPanic_genType) PrettyParams(params MoqPanic_genType_params) string {
	return fmt.Sprintf("Panic_genType(%#v)", params.V)
}

func (m *MoqPanic_genType) ParamsKey(params MoqPanic_genType_params, anyParams uint64) MoqPanic_genType_paramsKey {
	m.Scene.T.Helper()
	var vUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.V == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The v parameter can't be indexed by value")
		}
		vUsedHash = hash.DeepHash(params.V)
	}
	return MoqPanic_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ V hash.Hash }{
			V: vUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqPanic_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqPanic_genType) AssertExpectationsMet() {
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

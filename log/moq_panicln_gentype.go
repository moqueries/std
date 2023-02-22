// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package log

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Panicln_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Panicln_genType func(v ...any)

// MoqPanicln_genType holds the state of a moq of the Panicln_genType type
type MoqPanicln_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqPanicln_genType_mock

	ResultsByParams []MoqPanicln_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			V moq.ParamIndexing
		}
	}
}

// MoqPanicln_genType_mock isolates the mock interface of the Panicln_genType
// type
type MoqPanicln_genType_mock struct {
	Moq *MoqPanicln_genType
}

// MoqPanicln_genType_params holds the params of the Panicln_genType type
type MoqPanicln_genType_params struct{ V []any }

// MoqPanicln_genType_paramsKey holds the map key params of the Panicln_genType
// type
type MoqPanicln_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ V hash.Hash }
}

// MoqPanicln_genType_resultsByParams contains the results for a given set of
// parameters for the Panicln_genType type
type MoqPanicln_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqPanicln_genType_paramsKey]*MoqPanicln_genType_results
}

// MoqPanicln_genType_doFn defines the type of function needed when calling
// AndDo for the Panicln_genType type
type MoqPanicln_genType_doFn func(v ...any)

// MoqPanicln_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Panicln_genType type
type MoqPanicln_genType_doReturnFn func(v ...any)

// MoqPanicln_genType_results holds the results of the Panicln_genType type
type MoqPanicln_genType_results struct {
	Params  MoqPanicln_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqPanicln_genType_doFn
		DoReturnFn MoqPanicln_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqPanicln_genType_fnRecorder routes recorded function calls to the
// MoqPanicln_genType moq
type MoqPanicln_genType_fnRecorder struct {
	Params    MoqPanicln_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqPanicln_genType_results
	Moq       *MoqPanicln_genType
}

// MoqPanicln_genType_anyParams isolates the any params functions of the
// Panicln_genType type
type MoqPanicln_genType_anyParams struct {
	Recorder *MoqPanicln_genType_fnRecorder
}

// NewMoqPanicln_genType creates a new moq of the Panicln_genType type
func NewMoqPanicln_genType(scene *moq.Scene, config *moq.Config) *MoqPanicln_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqPanicln_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqPanicln_genType_mock{},

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

// Mock returns the moq implementation of the Panicln_genType type
func (m *MoqPanicln_genType) Mock() Panicln_genType {
	return func(v ...any) { m.Scene.T.Helper(); moq := &MoqPanicln_genType_mock{Moq: m}; moq.Fn(v...) }
}

func (m *MoqPanicln_genType_mock) Fn(v ...any) {
	m.Moq.Scene.T.Helper()
	params := MoqPanicln_genType_params{
		V: v,
	}
	var results *MoqPanicln_genType_results
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

func (m *MoqPanicln_genType) OnCall(v ...any) *MoqPanicln_genType_fnRecorder {
	return &MoqPanicln_genType_fnRecorder{
		Params: MoqPanicln_genType_params{
			V: v,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqPanicln_genType_fnRecorder) Any() *MoqPanicln_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqPanicln_genType_anyParams{Recorder: r}
}

func (a *MoqPanicln_genType_anyParams) V() *MoqPanicln_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqPanicln_genType_fnRecorder) Seq() *MoqPanicln_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqPanicln_genType_fnRecorder) NoSeq() *MoqPanicln_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqPanicln_genType_fnRecorder) ReturnResults() *MoqPanicln_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqPanicln_genType_doFn
		DoReturnFn MoqPanicln_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqPanicln_genType_fnRecorder) AndDo(fn MoqPanicln_genType_doFn) *MoqPanicln_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqPanicln_genType_fnRecorder) DoReturnResults(fn MoqPanicln_genType_doReturnFn) *MoqPanicln_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqPanicln_genType_doFn
		DoReturnFn MoqPanicln_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqPanicln_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqPanicln_genType_resultsByParams
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
		results = &MoqPanicln_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqPanicln_genType_paramsKey]*MoqPanicln_genType_results{},
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
		r.Results = &MoqPanicln_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqPanicln_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqPanicln_genType_fnRecorder {
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
				DoFn       MoqPanicln_genType_doFn
				DoReturnFn MoqPanicln_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqPanicln_genType) PrettyParams(params MoqPanicln_genType_params) string {
	return fmt.Sprintf("Panicln_genType(%#v)", params.V)
}

func (m *MoqPanicln_genType) ParamsKey(params MoqPanicln_genType_params, anyParams uint64) MoqPanicln_genType_paramsKey {
	m.Scene.T.Helper()
	var vUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.V == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The v parameter can't be indexed by value")
		}
		vUsedHash = hash.DeepHash(params.V)
	}
	return MoqPanicln_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ V hash.Hash }{
			V: vUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqPanicln_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqPanicln_genType) AssertExpectationsMet() {
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

// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package math

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// IsNaN_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type IsNaN_genType func(f float64) (is bool)

// MoqIsNaN_genType holds the state of a moq of the IsNaN_genType type
type MoqIsNaN_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqIsNaN_genType_mock

	ResultsByParams []MoqIsNaN_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			F moq.ParamIndexing
		}
	}
}

// MoqIsNaN_genType_mock isolates the mock interface of the IsNaN_genType type
type MoqIsNaN_genType_mock struct {
	Moq *MoqIsNaN_genType
}

// MoqIsNaN_genType_params holds the params of the IsNaN_genType type
type MoqIsNaN_genType_params struct{ F float64 }

// MoqIsNaN_genType_paramsKey holds the map key params of the IsNaN_genType
// type
type MoqIsNaN_genType_paramsKey struct {
	Params struct{ F float64 }
	Hashes struct{ F hash.Hash }
}

// MoqIsNaN_genType_resultsByParams contains the results for a given set of
// parameters for the IsNaN_genType type
type MoqIsNaN_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqIsNaN_genType_paramsKey]*MoqIsNaN_genType_results
}

// MoqIsNaN_genType_doFn defines the type of function needed when calling AndDo
// for the IsNaN_genType type
type MoqIsNaN_genType_doFn func(f float64)

// MoqIsNaN_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the IsNaN_genType type
type MoqIsNaN_genType_doReturnFn func(f float64) (is bool)

// MoqIsNaN_genType_results holds the results of the IsNaN_genType type
type MoqIsNaN_genType_results struct {
	Params  MoqIsNaN_genType_params
	Results []struct {
		Values     *struct{ Is bool }
		Sequence   uint32
		DoFn       MoqIsNaN_genType_doFn
		DoReturnFn MoqIsNaN_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqIsNaN_genType_fnRecorder routes recorded function calls to the
// MoqIsNaN_genType moq
type MoqIsNaN_genType_fnRecorder struct {
	Params    MoqIsNaN_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqIsNaN_genType_results
	Moq       *MoqIsNaN_genType
}

// MoqIsNaN_genType_anyParams isolates the any params functions of the
// IsNaN_genType type
type MoqIsNaN_genType_anyParams struct {
	Recorder *MoqIsNaN_genType_fnRecorder
}

// NewMoqIsNaN_genType creates a new moq of the IsNaN_genType type
func NewMoqIsNaN_genType(scene *moq.Scene, config *moq.Config) *MoqIsNaN_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqIsNaN_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqIsNaN_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				F moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			F moq.ParamIndexing
		}{
			F: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the IsNaN_genType type
func (m *MoqIsNaN_genType) Mock() IsNaN_genType {
	return func(f float64) (_ bool) { m.Scene.T.Helper(); moq := &MoqIsNaN_genType_mock{Moq: m}; return moq.Fn(f) }
}

func (m *MoqIsNaN_genType_mock) Fn(f float64) (is bool) {
	m.Moq.Scene.T.Helper()
	params := MoqIsNaN_genType_params{
		F: f,
	}
	var results *MoqIsNaN_genType_results
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
		result.DoFn(f)
	}

	if result.Values != nil {
		is = result.Values.Is
	}
	if result.DoReturnFn != nil {
		is = result.DoReturnFn(f)
	}
	return
}

func (m *MoqIsNaN_genType) OnCall(f float64) *MoqIsNaN_genType_fnRecorder {
	return &MoqIsNaN_genType_fnRecorder{
		Params: MoqIsNaN_genType_params{
			F: f,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqIsNaN_genType_fnRecorder) Any() *MoqIsNaN_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqIsNaN_genType_anyParams{Recorder: r}
}

func (a *MoqIsNaN_genType_anyParams) F() *MoqIsNaN_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqIsNaN_genType_fnRecorder) Seq() *MoqIsNaN_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqIsNaN_genType_fnRecorder) NoSeq() *MoqIsNaN_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqIsNaN_genType_fnRecorder) ReturnResults(is bool) *MoqIsNaN_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Is bool }
		Sequence   uint32
		DoFn       MoqIsNaN_genType_doFn
		DoReturnFn MoqIsNaN_genType_doReturnFn
	}{
		Values: &struct{ Is bool }{
			Is: is,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqIsNaN_genType_fnRecorder) AndDo(fn MoqIsNaN_genType_doFn) *MoqIsNaN_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqIsNaN_genType_fnRecorder) DoReturnResults(fn MoqIsNaN_genType_doReturnFn) *MoqIsNaN_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Is bool }
		Sequence   uint32
		DoFn       MoqIsNaN_genType_doFn
		DoReturnFn MoqIsNaN_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqIsNaN_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqIsNaN_genType_resultsByParams
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
		results = &MoqIsNaN_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqIsNaN_genType_paramsKey]*MoqIsNaN_genType_results{},
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
		r.Results = &MoqIsNaN_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqIsNaN_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqIsNaN_genType_fnRecorder {
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
				Values     *struct{ Is bool }
				Sequence   uint32
				DoFn       MoqIsNaN_genType_doFn
				DoReturnFn MoqIsNaN_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqIsNaN_genType) PrettyParams(params MoqIsNaN_genType_params) string {
	return fmt.Sprintf("IsNaN_genType(%#v)", params.F)
}

func (m *MoqIsNaN_genType) ParamsKey(params MoqIsNaN_genType_params, anyParams uint64) MoqIsNaN_genType_paramsKey {
	m.Scene.T.Helper()
	var fUsed float64
	var fUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.F == moq.ParamIndexByValue {
			fUsed = params.F
		} else {
			fUsedHash = hash.DeepHash(params.F)
		}
	}
	return MoqIsNaN_genType_paramsKey{
		Params: struct{ F float64 }{
			F: fUsed,
		},
		Hashes: struct{ F hash.Hash }{
			F: fUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqIsNaN_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqIsNaN_genType) AssertExpectationsMet() {
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

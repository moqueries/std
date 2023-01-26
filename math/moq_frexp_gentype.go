// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package math

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Frexp_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Frexp_genType func(f float64) (frac float64, exp int)

// MoqFrexp_genType holds the state of a moq of the Frexp_genType type
type MoqFrexp_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqFrexp_genType_mock

	ResultsByParams []MoqFrexp_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			F moq.ParamIndexing
		}
	}
}

// MoqFrexp_genType_mock isolates the mock interface of the Frexp_genType type
type MoqFrexp_genType_mock struct {
	Moq *MoqFrexp_genType
}

// MoqFrexp_genType_params holds the params of the Frexp_genType type
type MoqFrexp_genType_params struct{ F float64 }

// MoqFrexp_genType_paramsKey holds the map key params of the Frexp_genType
// type
type MoqFrexp_genType_paramsKey struct {
	Params struct{ F float64 }
	Hashes struct{ F hash.Hash }
}

// MoqFrexp_genType_resultsByParams contains the results for a given set of
// parameters for the Frexp_genType type
type MoqFrexp_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFrexp_genType_paramsKey]*MoqFrexp_genType_results
}

// MoqFrexp_genType_doFn defines the type of function needed when calling AndDo
// for the Frexp_genType type
type MoqFrexp_genType_doFn func(f float64)

// MoqFrexp_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Frexp_genType type
type MoqFrexp_genType_doReturnFn func(f float64) (frac float64, exp int)

// MoqFrexp_genType_results holds the results of the Frexp_genType type
type MoqFrexp_genType_results struct {
	Params  MoqFrexp_genType_params
	Results []struct {
		Values *struct {
			Frac float64
			Exp  int
		}
		Sequence   uint32
		DoFn       MoqFrexp_genType_doFn
		DoReturnFn MoqFrexp_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFrexp_genType_fnRecorder routes recorded function calls to the
// MoqFrexp_genType moq
type MoqFrexp_genType_fnRecorder struct {
	Params    MoqFrexp_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFrexp_genType_results
	Moq       *MoqFrexp_genType
}

// MoqFrexp_genType_anyParams isolates the any params functions of the
// Frexp_genType type
type MoqFrexp_genType_anyParams struct {
	Recorder *MoqFrexp_genType_fnRecorder
}

// NewMoqFrexp_genType creates a new moq of the Frexp_genType type
func NewMoqFrexp_genType(scene *moq.Scene, config *moq.Config) *MoqFrexp_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqFrexp_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqFrexp_genType_mock{},

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

// Mock returns the moq implementation of the Frexp_genType type
func (m *MoqFrexp_genType) Mock() Frexp_genType {
	return func(f float64) (_ float64, _ int) {
		m.Scene.T.Helper()
		moq := &MoqFrexp_genType_mock{Moq: m}
		return moq.Fn(f)
	}
}

func (m *MoqFrexp_genType_mock) Fn(f float64) (frac float64, exp int) {
	m.Moq.Scene.T.Helper()
	params := MoqFrexp_genType_params{
		F: f,
	}
	var results *MoqFrexp_genType_results
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
		frac = result.Values.Frac
		exp = result.Values.Exp
	}
	if result.DoReturnFn != nil {
		frac, exp = result.DoReturnFn(f)
	}
	return
}

func (m *MoqFrexp_genType) OnCall(f float64) *MoqFrexp_genType_fnRecorder {
	return &MoqFrexp_genType_fnRecorder{
		Params: MoqFrexp_genType_params{
			F: f,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqFrexp_genType_fnRecorder) Any() *MoqFrexp_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqFrexp_genType_anyParams{Recorder: r}
}

func (a *MoqFrexp_genType_anyParams) F() *MoqFrexp_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqFrexp_genType_fnRecorder) Seq() *MoqFrexp_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFrexp_genType_fnRecorder) NoSeq() *MoqFrexp_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFrexp_genType_fnRecorder) ReturnResults(frac float64, exp int) *MoqFrexp_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Frac float64
			Exp  int
		}
		Sequence   uint32
		DoFn       MoqFrexp_genType_doFn
		DoReturnFn MoqFrexp_genType_doReturnFn
	}{
		Values: &struct {
			Frac float64
			Exp  int
		}{
			Frac: frac,
			Exp:  exp,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqFrexp_genType_fnRecorder) AndDo(fn MoqFrexp_genType_doFn) *MoqFrexp_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFrexp_genType_fnRecorder) DoReturnResults(fn MoqFrexp_genType_doReturnFn) *MoqFrexp_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Frac float64
			Exp  int
		}
		Sequence   uint32
		DoFn       MoqFrexp_genType_doFn
		DoReturnFn MoqFrexp_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFrexp_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFrexp_genType_resultsByParams
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
		results = &MoqFrexp_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFrexp_genType_paramsKey]*MoqFrexp_genType_results{},
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
		r.Results = &MoqFrexp_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFrexp_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFrexp_genType_fnRecorder {
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
					Frac float64
					Exp  int
				}
				Sequence   uint32
				DoFn       MoqFrexp_genType_doFn
				DoReturnFn MoqFrexp_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFrexp_genType) PrettyParams(params MoqFrexp_genType_params) string {
	return fmt.Sprintf("Frexp_genType(%#v)", params.F)
}

func (m *MoqFrexp_genType) ParamsKey(params MoqFrexp_genType_params, anyParams uint64) MoqFrexp_genType_paramsKey {
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
	return MoqFrexp_genType_paramsKey{
		Params: struct{ F float64 }{
			F: fUsed,
		},
		Hashes: struct{ F hash.Hash }{
			F: fUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqFrexp_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqFrexp_genType) AssertExpectationsMet() {
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

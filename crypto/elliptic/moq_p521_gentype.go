// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package elliptic

import (
	"crypto/elliptic"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// P521_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type P521_genType func() elliptic.Curve

// MoqP521_genType holds the state of a moq of the P521_genType type
type MoqP521_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqP521_genType_mock

	ResultsByParams []MoqP521_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct{}
	}
}

// MoqP521_genType_mock isolates the mock interface of the P521_genType type
type MoqP521_genType_mock struct {
	Moq *MoqP521_genType
}

// MoqP521_genType_params holds the params of the P521_genType type
type MoqP521_genType_params struct{}

// MoqP521_genType_paramsKey holds the map key params of the P521_genType type
type MoqP521_genType_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqP521_genType_resultsByParams contains the results for a given set of
// parameters for the P521_genType type
type MoqP521_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqP521_genType_paramsKey]*MoqP521_genType_results
}

// MoqP521_genType_doFn defines the type of function needed when calling AndDo
// for the P521_genType type
type MoqP521_genType_doFn func()

// MoqP521_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the P521_genType type
type MoqP521_genType_doReturnFn func() elliptic.Curve

// MoqP521_genType_results holds the results of the P521_genType type
type MoqP521_genType_results struct {
	Params  MoqP521_genType_params
	Results []struct {
		Values *struct {
			Result1 elliptic.Curve
		}
		Sequence   uint32
		DoFn       MoqP521_genType_doFn
		DoReturnFn MoqP521_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqP521_genType_fnRecorder routes recorded function calls to the
// MoqP521_genType moq
type MoqP521_genType_fnRecorder struct {
	Params    MoqP521_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqP521_genType_results
	Moq       *MoqP521_genType
}

// MoqP521_genType_anyParams isolates the any params functions of the
// P521_genType type
type MoqP521_genType_anyParams struct {
	Recorder *MoqP521_genType_fnRecorder
}

// NewMoqP521_genType creates a new moq of the P521_genType type
func NewMoqP521_genType(scene *moq.Scene, config *moq.Config) *MoqP521_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqP521_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqP521_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct{}
		}{ParameterIndexing: struct{}{}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the P521_genType type
func (m *MoqP521_genType) Mock() P521_genType {
	return func() elliptic.Curve { m.Scene.T.Helper(); moq := &MoqP521_genType_mock{Moq: m}; return moq.Fn() }
}

func (m *MoqP521_genType_mock) Fn() (result1 elliptic.Curve) {
	m.Moq.Scene.T.Helper()
	params := MoqP521_genType_params{}
	var results *MoqP521_genType_results
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
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn()
	}
	return
}

func (m *MoqP521_genType) OnCall() *MoqP521_genType_fnRecorder {
	return &MoqP521_genType_fnRecorder{
		Params:   MoqP521_genType_params{},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqP521_genType_fnRecorder) Any() *MoqP521_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqP521_genType_anyParams{Recorder: r}
}

func (r *MoqP521_genType_fnRecorder) Seq() *MoqP521_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqP521_genType_fnRecorder) NoSeq() *MoqP521_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqP521_genType_fnRecorder) ReturnResults(result1 elliptic.Curve) *MoqP521_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 elliptic.Curve
		}
		Sequence   uint32
		DoFn       MoqP521_genType_doFn
		DoReturnFn MoqP521_genType_doReturnFn
	}{
		Values: &struct {
			Result1 elliptic.Curve
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqP521_genType_fnRecorder) AndDo(fn MoqP521_genType_doFn) *MoqP521_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqP521_genType_fnRecorder) DoReturnResults(fn MoqP521_genType_doReturnFn) *MoqP521_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 elliptic.Curve
		}
		Sequence   uint32
		DoFn       MoqP521_genType_doFn
		DoReturnFn MoqP521_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqP521_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqP521_genType_resultsByParams
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
		results = &MoqP521_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqP521_genType_paramsKey]*MoqP521_genType_results{},
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
		r.Results = &MoqP521_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqP521_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqP521_genType_fnRecorder {
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
					Result1 elliptic.Curve
				}
				Sequence   uint32
				DoFn       MoqP521_genType_doFn
				DoReturnFn MoqP521_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqP521_genType) PrettyParams(params MoqP521_genType_params) string {
	return fmt.Sprintf("P521_genType()")
}

func (m *MoqP521_genType) ParamsKey(params MoqP521_genType_params, anyParams uint64) MoqP521_genType_paramsKey {
	m.Scene.T.Helper()
	return MoqP521_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqP521_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqP521_genType) AssertExpectationsMet() {
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

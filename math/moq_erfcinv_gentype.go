// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package math

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Erfcinv_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Erfcinv_genType func(x float64) float64

// MoqErfcinv_genType holds the state of a moq of the Erfcinv_genType type
type MoqErfcinv_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqErfcinv_genType_mock

	ResultsByParams []MoqErfcinv_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			X moq.ParamIndexing
		}
	}
}

// MoqErfcinv_genType_mock isolates the mock interface of the Erfcinv_genType
// type
type MoqErfcinv_genType_mock struct {
	Moq *MoqErfcinv_genType
}

// MoqErfcinv_genType_params holds the params of the Erfcinv_genType type
type MoqErfcinv_genType_params struct{ X float64 }

// MoqErfcinv_genType_paramsKey holds the map key params of the Erfcinv_genType
// type
type MoqErfcinv_genType_paramsKey struct {
	Params struct{ X float64 }
	Hashes struct{ X hash.Hash }
}

// MoqErfcinv_genType_resultsByParams contains the results for a given set of
// parameters for the Erfcinv_genType type
type MoqErfcinv_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqErfcinv_genType_paramsKey]*MoqErfcinv_genType_results
}

// MoqErfcinv_genType_doFn defines the type of function needed when calling
// AndDo for the Erfcinv_genType type
type MoqErfcinv_genType_doFn func(x float64)

// MoqErfcinv_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Erfcinv_genType type
type MoqErfcinv_genType_doReturnFn func(x float64) float64

// MoqErfcinv_genType_results holds the results of the Erfcinv_genType type
type MoqErfcinv_genType_results struct {
	Params  MoqErfcinv_genType_params
	Results []struct {
		Values *struct {
			Result1 float64
		}
		Sequence   uint32
		DoFn       MoqErfcinv_genType_doFn
		DoReturnFn MoqErfcinv_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqErfcinv_genType_fnRecorder routes recorded function calls to the
// MoqErfcinv_genType moq
type MoqErfcinv_genType_fnRecorder struct {
	Params    MoqErfcinv_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqErfcinv_genType_results
	Moq       *MoqErfcinv_genType
}

// MoqErfcinv_genType_anyParams isolates the any params functions of the
// Erfcinv_genType type
type MoqErfcinv_genType_anyParams struct {
	Recorder *MoqErfcinv_genType_fnRecorder
}

// NewMoqErfcinv_genType creates a new moq of the Erfcinv_genType type
func NewMoqErfcinv_genType(scene *moq.Scene, config *moq.Config) *MoqErfcinv_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqErfcinv_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqErfcinv_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				X moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			X moq.ParamIndexing
		}{
			X: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Erfcinv_genType type
func (m *MoqErfcinv_genType) Mock() Erfcinv_genType {
	return func(x float64) float64 { m.Scene.T.Helper(); moq := &MoqErfcinv_genType_mock{Moq: m}; return moq.Fn(x) }
}

func (m *MoqErfcinv_genType_mock) Fn(x float64) (result1 float64) {
	m.Moq.Scene.T.Helper()
	params := MoqErfcinv_genType_params{
		X: x,
	}
	var results *MoqErfcinv_genType_results
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
		result.DoFn(x)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(x)
	}
	return
}

func (m *MoqErfcinv_genType) OnCall(x float64) *MoqErfcinv_genType_fnRecorder {
	return &MoqErfcinv_genType_fnRecorder{
		Params: MoqErfcinv_genType_params{
			X: x,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqErfcinv_genType_fnRecorder) Any() *MoqErfcinv_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqErfcinv_genType_anyParams{Recorder: r}
}

func (a *MoqErfcinv_genType_anyParams) X() *MoqErfcinv_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqErfcinv_genType_fnRecorder) Seq() *MoqErfcinv_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqErfcinv_genType_fnRecorder) NoSeq() *MoqErfcinv_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqErfcinv_genType_fnRecorder) ReturnResults(result1 float64) *MoqErfcinv_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 float64
		}
		Sequence   uint32
		DoFn       MoqErfcinv_genType_doFn
		DoReturnFn MoqErfcinv_genType_doReturnFn
	}{
		Values: &struct {
			Result1 float64
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqErfcinv_genType_fnRecorder) AndDo(fn MoqErfcinv_genType_doFn) *MoqErfcinv_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqErfcinv_genType_fnRecorder) DoReturnResults(fn MoqErfcinv_genType_doReturnFn) *MoqErfcinv_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 float64
		}
		Sequence   uint32
		DoFn       MoqErfcinv_genType_doFn
		DoReturnFn MoqErfcinv_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqErfcinv_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqErfcinv_genType_resultsByParams
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
		results = &MoqErfcinv_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqErfcinv_genType_paramsKey]*MoqErfcinv_genType_results{},
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
		r.Results = &MoqErfcinv_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqErfcinv_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqErfcinv_genType_fnRecorder {
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
					Result1 float64
				}
				Sequence   uint32
				DoFn       MoqErfcinv_genType_doFn
				DoReturnFn MoqErfcinv_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqErfcinv_genType) PrettyParams(params MoqErfcinv_genType_params) string {
	return fmt.Sprintf("Erfcinv_genType(%#v)", params.X)
}

func (m *MoqErfcinv_genType) ParamsKey(params MoqErfcinv_genType_params, anyParams uint64) MoqErfcinv_genType_paramsKey {
	m.Scene.T.Helper()
	var xUsed float64
	var xUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.X == moq.ParamIndexByValue {
			xUsed = params.X
		} else {
			xUsedHash = hash.DeepHash(params.X)
		}
	}
	return MoqErfcinv_genType_paramsKey{
		Params: struct{ X float64 }{
			X: xUsed,
		},
		Hashes: struct{ X hash.Hash }{
			X: xUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqErfcinv_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqErfcinv_genType) AssertExpectationsMet() {
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

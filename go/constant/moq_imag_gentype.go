// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package constant

import (
	"fmt"
	"go/constant"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Imag_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Imag_genType func(x constant.Value) constant.Value

// MoqImag_genType holds the state of a moq of the Imag_genType type
type MoqImag_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqImag_genType_mock

	ResultsByParams []MoqImag_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			X moq.ParamIndexing
		}
	}
}

// MoqImag_genType_mock isolates the mock interface of the Imag_genType type
type MoqImag_genType_mock struct {
	Moq *MoqImag_genType
}

// MoqImag_genType_params holds the params of the Imag_genType type
type MoqImag_genType_params struct{ X constant.Value }

// MoqImag_genType_paramsKey holds the map key params of the Imag_genType type
type MoqImag_genType_paramsKey struct {
	Params struct{ X constant.Value }
	Hashes struct{ X hash.Hash }
}

// MoqImag_genType_resultsByParams contains the results for a given set of
// parameters for the Imag_genType type
type MoqImag_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqImag_genType_paramsKey]*MoqImag_genType_results
}

// MoqImag_genType_doFn defines the type of function needed when calling AndDo
// for the Imag_genType type
type MoqImag_genType_doFn func(x constant.Value)

// MoqImag_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Imag_genType type
type MoqImag_genType_doReturnFn func(x constant.Value) constant.Value

// MoqImag_genType_results holds the results of the Imag_genType type
type MoqImag_genType_results struct {
	Params  MoqImag_genType_params
	Results []struct {
		Values *struct {
			Result1 constant.Value
		}
		Sequence   uint32
		DoFn       MoqImag_genType_doFn
		DoReturnFn MoqImag_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqImag_genType_fnRecorder routes recorded function calls to the
// MoqImag_genType moq
type MoqImag_genType_fnRecorder struct {
	Params    MoqImag_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqImag_genType_results
	Moq       *MoqImag_genType
}

// MoqImag_genType_anyParams isolates the any params functions of the
// Imag_genType type
type MoqImag_genType_anyParams struct {
	Recorder *MoqImag_genType_fnRecorder
}

// NewMoqImag_genType creates a new moq of the Imag_genType type
func NewMoqImag_genType(scene *moq.Scene, config *moq.Config) *MoqImag_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqImag_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqImag_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				X moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			X moq.ParamIndexing
		}{
			X: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Imag_genType type
func (m *MoqImag_genType) Mock() Imag_genType {
	return func(x constant.Value) constant.Value {
		m.Scene.T.Helper()
		moq := &MoqImag_genType_mock{Moq: m}
		return moq.Fn(x)
	}
}

func (m *MoqImag_genType_mock) Fn(x constant.Value) (result1 constant.Value) {
	m.Moq.Scene.T.Helper()
	params := MoqImag_genType_params{
		X: x,
	}
	var results *MoqImag_genType_results
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

func (m *MoqImag_genType) OnCall(x constant.Value) *MoqImag_genType_fnRecorder {
	return &MoqImag_genType_fnRecorder{
		Params: MoqImag_genType_params{
			X: x,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqImag_genType_fnRecorder) Any() *MoqImag_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqImag_genType_anyParams{Recorder: r}
}

func (a *MoqImag_genType_anyParams) X() *MoqImag_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqImag_genType_fnRecorder) Seq() *MoqImag_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqImag_genType_fnRecorder) NoSeq() *MoqImag_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqImag_genType_fnRecorder) ReturnResults(result1 constant.Value) *MoqImag_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 constant.Value
		}
		Sequence   uint32
		DoFn       MoqImag_genType_doFn
		DoReturnFn MoqImag_genType_doReturnFn
	}{
		Values: &struct {
			Result1 constant.Value
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqImag_genType_fnRecorder) AndDo(fn MoqImag_genType_doFn) *MoqImag_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqImag_genType_fnRecorder) DoReturnResults(fn MoqImag_genType_doReturnFn) *MoqImag_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 constant.Value
		}
		Sequence   uint32
		DoFn       MoqImag_genType_doFn
		DoReturnFn MoqImag_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqImag_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqImag_genType_resultsByParams
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
		results = &MoqImag_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqImag_genType_paramsKey]*MoqImag_genType_results{},
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
		r.Results = &MoqImag_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqImag_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqImag_genType_fnRecorder {
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
					Result1 constant.Value
				}
				Sequence   uint32
				DoFn       MoqImag_genType_doFn
				DoReturnFn MoqImag_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqImag_genType) PrettyParams(params MoqImag_genType_params) string {
	return fmt.Sprintf("Imag_genType(%#v)", params.X)
}

func (m *MoqImag_genType) ParamsKey(params MoqImag_genType_params, anyParams uint64) MoqImag_genType_paramsKey {
	m.Scene.T.Helper()
	var xUsed constant.Value
	var xUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.X == moq.ParamIndexByValue {
			xUsed = params.X
		} else {
			xUsedHash = hash.DeepHash(params.X)
		}
	}
	return MoqImag_genType_paramsKey{
		Params: struct{ X constant.Value }{
			X: xUsed,
		},
		Hashes: struct{ X hash.Hash }{
			X: xUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqImag_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqImag_genType) AssertExpectationsMet() {
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

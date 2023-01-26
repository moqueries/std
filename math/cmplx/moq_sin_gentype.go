// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package cmplx

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Sin_genType is the fabricated implementation type of this mock (emitted when
// mocking functions directly and not from a function type)
type Sin_genType func(x complex128) complex128

// MoqSin_genType holds the state of a moq of the Sin_genType type
type MoqSin_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSin_genType_mock

	ResultsByParams []MoqSin_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			X moq.ParamIndexing
		}
	}
}

// MoqSin_genType_mock isolates the mock interface of the Sin_genType type
type MoqSin_genType_mock struct {
	Moq *MoqSin_genType
}

// MoqSin_genType_params holds the params of the Sin_genType type
type MoqSin_genType_params struct{ X complex128 }

// MoqSin_genType_paramsKey holds the map key params of the Sin_genType type
type MoqSin_genType_paramsKey struct {
	Params struct{ X complex128 }
	Hashes struct{ X hash.Hash }
}

// MoqSin_genType_resultsByParams contains the results for a given set of
// parameters for the Sin_genType type
type MoqSin_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSin_genType_paramsKey]*MoqSin_genType_results
}

// MoqSin_genType_doFn defines the type of function needed when calling AndDo
// for the Sin_genType type
type MoqSin_genType_doFn func(x complex128)

// MoqSin_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Sin_genType type
type MoqSin_genType_doReturnFn func(x complex128) complex128

// MoqSin_genType_results holds the results of the Sin_genType type
type MoqSin_genType_results struct {
	Params  MoqSin_genType_params
	Results []struct {
		Values *struct {
			Result1 complex128
		}
		Sequence   uint32
		DoFn       MoqSin_genType_doFn
		DoReturnFn MoqSin_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSin_genType_fnRecorder routes recorded function calls to the
// MoqSin_genType moq
type MoqSin_genType_fnRecorder struct {
	Params    MoqSin_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSin_genType_results
	Moq       *MoqSin_genType
}

// MoqSin_genType_anyParams isolates the any params functions of the
// Sin_genType type
type MoqSin_genType_anyParams struct {
	Recorder *MoqSin_genType_fnRecorder
}

// NewMoqSin_genType creates a new moq of the Sin_genType type
func NewMoqSin_genType(scene *moq.Scene, config *moq.Config) *MoqSin_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSin_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSin_genType_mock{},

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

// Mock returns the moq implementation of the Sin_genType type
func (m *MoqSin_genType) Mock() Sin_genType {
	return func(x complex128) complex128 {
		m.Scene.T.Helper()
		moq := &MoqSin_genType_mock{Moq: m}
		return moq.Fn(x)
	}
}

func (m *MoqSin_genType_mock) Fn(x complex128) (result1 complex128) {
	m.Moq.Scene.T.Helper()
	params := MoqSin_genType_params{
		X: x,
	}
	var results *MoqSin_genType_results
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

func (m *MoqSin_genType) OnCall(x complex128) *MoqSin_genType_fnRecorder {
	return &MoqSin_genType_fnRecorder{
		Params: MoqSin_genType_params{
			X: x,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSin_genType_fnRecorder) Any() *MoqSin_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSin_genType_anyParams{Recorder: r}
}

func (a *MoqSin_genType_anyParams) X() *MoqSin_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqSin_genType_fnRecorder) Seq() *MoqSin_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSin_genType_fnRecorder) NoSeq() *MoqSin_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSin_genType_fnRecorder) ReturnResults(result1 complex128) *MoqSin_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 complex128
		}
		Sequence   uint32
		DoFn       MoqSin_genType_doFn
		DoReturnFn MoqSin_genType_doReturnFn
	}{
		Values: &struct {
			Result1 complex128
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSin_genType_fnRecorder) AndDo(fn MoqSin_genType_doFn) *MoqSin_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSin_genType_fnRecorder) DoReturnResults(fn MoqSin_genType_doReturnFn) *MoqSin_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 complex128
		}
		Sequence   uint32
		DoFn       MoqSin_genType_doFn
		DoReturnFn MoqSin_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSin_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSin_genType_resultsByParams
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
		results = &MoqSin_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSin_genType_paramsKey]*MoqSin_genType_results{},
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
		r.Results = &MoqSin_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSin_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSin_genType_fnRecorder {
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
					Result1 complex128
				}
				Sequence   uint32
				DoFn       MoqSin_genType_doFn
				DoReturnFn MoqSin_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSin_genType) PrettyParams(params MoqSin_genType_params) string {
	return fmt.Sprintf("Sin_genType(%#v)", params.X)
}

func (m *MoqSin_genType) ParamsKey(params MoqSin_genType_params, anyParams uint64) MoqSin_genType_paramsKey {
	m.Scene.T.Helper()
	var xUsed complex128
	var xUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.X == moq.ParamIndexByValue {
			xUsed = params.X
		} else {
			xUsedHash = hash.DeepHash(params.X)
		}
	}
	return MoqSin_genType_paramsKey{
		Params: struct{ X complex128 }{
			X: xUsed,
		},
		Hashes: struct{ X hash.Hash }{
			X: xUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSin_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSin_genType) AssertExpectationsMet() {
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

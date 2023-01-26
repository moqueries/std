// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package cmplx

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Log10_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Log10_genType func(x complex128) complex128

// MoqLog10_genType holds the state of a moq of the Log10_genType type
type MoqLog10_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqLog10_genType_mock

	ResultsByParams []MoqLog10_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			X moq.ParamIndexing
		}
	}
}

// MoqLog10_genType_mock isolates the mock interface of the Log10_genType type
type MoqLog10_genType_mock struct {
	Moq *MoqLog10_genType
}

// MoqLog10_genType_params holds the params of the Log10_genType type
type MoqLog10_genType_params struct{ X complex128 }

// MoqLog10_genType_paramsKey holds the map key params of the Log10_genType
// type
type MoqLog10_genType_paramsKey struct {
	Params struct{ X complex128 }
	Hashes struct{ X hash.Hash }
}

// MoqLog10_genType_resultsByParams contains the results for a given set of
// parameters for the Log10_genType type
type MoqLog10_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqLog10_genType_paramsKey]*MoqLog10_genType_results
}

// MoqLog10_genType_doFn defines the type of function needed when calling AndDo
// for the Log10_genType type
type MoqLog10_genType_doFn func(x complex128)

// MoqLog10_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Log10_genType type
type MoqLog10_genType_doReturnFn func(x complex128) complex128

// MoqLog10_genType_results holds the results of the Log10_genType type
type MoqLog10_genType_results struct {
	Params  MoqLog10_genType_params
	Results []struct {
		Values *struct {
			Result1 complex128
		}
		Sequence   uint32
		DoFn       MoqLog10_genType_doFn
		DoReturnFn MoqLog10_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqLog10_genType_fnRecorder routes recorded function calls to the
// MoqLog10_genType moq
type MoqLog10_genType_fnRecorder struct {
	Params    MoqLog10_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqLog10_genType_results
	Moq       *MoqLog10_genType
}

// MoqLog10_genType_anyParams isolates the any params functions of the
// Log10_genType type
type MoqLog10_genType_anyParams struct {
	Recorder *MoqLog10_genType_fnRecorder
}

// NewMoqLog10_genType creates a new moq of the Log10_genType type
func NewMoqLog10_genType(scene *moq.Scene, config *moq.Config) *MoqLog10_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqLog10_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqLog10_genType_mock{},

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

// Mock returns the moq implementation of the Log10_genType type
func (m *MoqLog10_genType) Mock() Log10_genType {
	return func(x complex128) complex128 {
		m.Scene.T.Helper()
		moq := &MoqLog10_genType_mock{Moq: m}
		return moq.Fn(x)
	}
}

func (m *MoqLog10_genType_mock) Fn(x complex128) (result1 complex128) {
	m.Moq.Scene.T.Helper()
	params := MoqLog10_genType_params{
		X: x,
	}
	var results *MoqLog10_genType_results
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

func (m *MoqLog10_genType) OnCall(x complex128) *MoqLog10_genType_fnRecorder {
	return &MoqLog10_genType_fnRecorder{
		Params: MoqLog10_genType_params{
			X: x,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqLog10_genType_fnRecorder) Any() *MoqLog10_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqLog10_genType_anyParams{Recorder: r}
}

func (a *MoqLog10_genType_anyParams) X() *MoqLog10_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqLog10_genType_fnRecorder) Seq() *MoqLog10_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqLog10_genType_fnRecorder) NoSeq() *MoqLog10_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqLog10_genType_fnRecorder) ReturnResults(result1 complex128) *MoqLog10_genType_fnRecorder {
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
		DoFn       MoqLog10_genType_doFn
		DoReturnFn MoqLog10_genType_doReturnFn
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

func (r *MoqLog10_genType_fnRecorder) AndDo(fn MoqLog10_genType_doFn) *MoqLog10_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqLog10_genType_fnRecorder) DoReturnResults(fn MoqLog10_genType_doReturnFn) *MoqLog10_genType_fnRecorder {
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
		DoFn       MoqLog10_genType_doFn
		DoReturnFn MoqLog10_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqLog10_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqLog10_genType_resultsByParams
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
		results = &MoqLog10_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqLog10_genType_paramsKey]*MoqLog10_genType_results{},
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
		r.Results = &MoqLog10_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqLog10_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqLog10_genType_fnRecorder {
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
				DoFn       MoqLog10_genType_doFn
				DoReturnFn MoqLog10_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqLog10_genType) PrettyParams(params MoqLog10_genType_params) string {
	return fmt.Sprintf("Log10_genType(%#v)", params.X)
}

func (m *MoqLog10_genType) ParamsKey(params MoqLog10_genType_params, anyParams uint64) MoqLog10_genType_paramsKey {
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
	return MoqLog10_genType_paramsKey{
		Params: struct{ X complex128 }{
			X: xUsed,
		},
		Hashes: struct{ X hash.Hash }{
			X: xUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqLog10_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqLog10_genType) AssertExpectationsMet() {
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

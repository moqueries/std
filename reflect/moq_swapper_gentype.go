// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package reflect

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Swapper_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Swapper_genType func(slice any) func(i, j int)

// MoqSwapper_genType holds the state of a moq of the Swapper_genType type
type MoqSwapper_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSwapper_genType_mock

	ResultsByParams []MoqSwapper_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Slice moq.ParamIndexing
		}
	}
}

// MoqSwapper_genType_mock isolates the mock interface of the Swapper_genType
// type
type MoqSwapper_genType_mock struct {
	Moq *MoqSwapper_genType
}

// MoqSwapper_genType_params holds the params of the Swapper_genType type
type MoqSwapper_genType_params struct{ Slice any }

// MoqSwapper_genType_paramsKey holds the map key params of the Swapper_genType
// type
type MoqSwapper_genType_paramsKey struct {
	Params struct{ Slice any }
	Hashes struct{ Slice hash.Hash }
}

// MoqSwapper_genType_resultsByParams contains the results for a given set of
// parameters for the Swapper_genType type
type MoqSwapper_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSwapper_genType_paramsKey]*MoqSwapper_genType_results
}

// MoqSwapper_genType_doFn defines the type of function needed when calling
// AndDo for the Swapper_genType type
type MoqSwapper_genType_doFn func(slice any)

// MoqSwapper_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Swapper_genType type
type MoqSwapper_genType_doReturnFn func(slice any) func(i, j int)

// MoqSwapper_genType_results holds the results of the Swapper_genType type
type MoqSwapper_genType_results struct {
	Params  MoqSwapper_genType_params
	Results []struct {
		Values *struct {
			Result1 func(i, j int)
		}
		Sequence   uint32
		DoFn       MoqSwapper_genType_doFn
		DoReturnFn MoqSwapper_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSwapper_genType_fnRecorder routes recorded function calls to the
// MoqSwapper_genType moq
type MoqSwapper_genType_fnRecorder struct {
	Params    MoqSwapper_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSwapper_genType_results
	Moq       *MoqSwapper_genType
}

// MoqSwapper_genType_anyParams isolates the any params functions of the
// Swapper_genType type
type MoqSwapper_genType_anyParams struct {
	Recorder *MoqSwapper_genType_fnRecorder
}

// NewMoqSwapper_genType creates a new moq of the Swapper_genType type
func NewMoqSwapper_genType(scene *moq.Scene, config *moq.Config) *MoqSwapper_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSwapper_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSwapper_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Slice moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Slice moq.ParamIndexing
		}{
			Slice: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Swapper_genType type
func (m *MoqSwapper_genType) Mock() Swapper_genType {
	return func(slice any) func(i, j int) {
		m.Scene.T.Helper()
		moq := &MoqSwapper_genType_mock{Moq: m}
		return moq.Fn(slice)
	}
}

func (m *MoqSwapper_genType_mock) Fn(slice any) (result1 func(i, j int)) {
	m.Moq.Scene.T.Helper()
	params := MoqSwapper_genType_params{
		Slice: slice,
	}
	var results *MoqSwapper_genType_results
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
		result.DoFn(slice)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(slice)
	}
	return
}

func (m *MoqSwapper_genType) OnCall(slice any) *MoqSwapper_genType_fnRecorder {
	return &MoqSwapper_genType_fnRecorder{
		Params: MoqSwapper_genType_params{
			Slice: slice,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSwapper_genType_fnRecorder) Any() *MoqSwapper_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSwapper_genType_anyParams{Recorder: r}
}

func (a *MoqSwapper_genType_anyParams) Slice() *MoqSwapper_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqSwapper_genType_fnRecorder) Seq() *MoqSwapper_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSwapper_genType_fnRecorder) NoSeq() *MoqSwapper_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSwapper_genType_fnRecorder) ReturnResults(result1 func(i, j int)) *MoqSwapper_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 func(i, j int)
		}
		Sequence   uint32
		DoFn       MoqSwapper_genType_doFn
		DoReturnFn MoqSwapper_genType_doReturnFn
	}{
		Values: &struct {
			Result1 func(i, j int)
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSwapper_genType_fnRecorder) AndDo(fn MoqSwapper_genType_doFn) *MoqSwapper_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSwapper_genType_fnRecorder) DoReturnResults(fn MoqSwapper_genType_doReturnFn) *MoqSwapper_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 func(i, j int)
		}
		Sequence   uint32
		DoFn       MoqSwapper_genType_doFn
		DoReturnFn MoqSwapper_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSwapper_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSwapper_genType_resultsByParams
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
		results = &MoqSwapper_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSwapper_genType_paramsKey]*MoqSwapper_genType_results{},
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
		r.Results = &MoqSwapper_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSwapper_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSwapper_genType_fnRecorder {
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
					Result1 func(i, j int)
				}
				Sequence   uint32
				DoFn       MoqSwapper_genType_doFn
				DoReturnFn MoqSwapper_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSwapper_genType) PrettyParams(params MoqSwapper_genType_params) string {
	return fmt.Sprintf("Swapper_genType(%#v)", params.Slice)
}

func (m *MoqSwapper_genType) ParamsKey(params MoqSwapper_genType_params, anyParams uint64) MoqSwapper_genType_paramsKey {
	m.Scene.T.Helper()
	var sliceUsed any
	var sliceUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Slice == moq.ParamIndexByValue {
			sliceUsed = params.Slice
		} else {
			sliceUsedHash = hash.DeepHash(params.Slice)
		}
	}
	return MoqSwapper_genType_paramsKey{
		Params: struct{ Slice any }{
			Slice: sliceUsed,
		},
		Hashes: struct{ Slice hash.Hash }{
			Slice: sliceUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSwapper_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSwapper_genType) AssertExpectationsMet() {
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

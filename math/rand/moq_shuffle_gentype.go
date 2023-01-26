// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package rand

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Shuffle_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Shuffle_genType func(n int, swap func(i, j int))

// MoqShuffle_genType holds the state of a moq of the Shuffle_genType type
type MoqShuffle_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqShuffle_genType_mock

	ResultsByParams []MoqShuffle_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			N    moq.ParamIndexing
			Swap moq.ParamIndexing
		}
	}
}

// MoqShuffle_genType_mock isolates the mock interface of the Shuffle_genType
// type
type MoqShuffle_genType_mock struct {
	Moq *MoqShuffle_genType
}

// MoqShuffle_genType_params holds the params of the Shuffle_genType type
type MoqShuffle_genType_params struct {
	N    int
	Swap func(i, j int)
}

// MoqShuffle_genType_paramsKey holds the map key params of the Shuffle_genType
// type
type MoqShuffle_genType_paramsKey struct {
	Params struct{ N int }
	Hashes struct {
		N    hash.Hash
		Swap hash.Hash
	}
}

// MoqShuffle_genType_resultsByParams contains the results for a given set of
// parameters for the Shuffle_genType type
type MoqShuffle_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqShuffle_genType_paramsKey]*MoqShuffle_genType_results
}

// MoqShuffle_genType_doFn defines the type of function needed when calling
// AndDo for the Shuffle_genType type
type MoqShuffle_genType_doFn func(n int, swap func(i, j int))

// MoqShuffle_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Shuffle_genType type
type MoqShuffle_genType_doReturnFn func(n int, swap func(i, j int))

// MoqShuffle_genType_results holds the results of the Shuffle_genType type
type MoqShuffle_genType_results struct {
	Params  MoqShuffle_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqShuffle_genType_doFn
		DoReturnFn MoqShuffle_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqShuffle_genType_fnRecorder routes recorded function calls to the
// MoqShuffle_genType moq
type MoqShuffle_genType_fnRecorder struct {
	Params    MoqShuffle_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqShuffle_genType_results
	Moq       *MoqShuffle_genType
}

// MoqShuffle_genType_anyParams isolates the any params functions of the
// Shuffle_genType type
type MoqShuffle_genType_anyParams struct {
	Recorder *MoqShuffle_genType_fnRecorder
}

// NewMoqShuffle_genType creates a new moq of the Shuffle_genType type
func NewMoqShuffle_genType(scene *moq.Scene, config *moq.Config) *MoqShuffle_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqShuffle_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqShuffle_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				N    moq.ParamIndexing
				Swap moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			N    moq.ParamIndexing
			Swap moq.ParamIndexing
		}{
			N:    moq.ParamIndexByValue,
			Swap: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Shuffle_genType type
func (m *MoqShuffle_genType) Mock() Shuffle_genType {
	return func(n int, swap func(i, j int)) {
		m.Scene.T.Helper()
		moq := &MoqShuffle_genType_mock{Moq: m}
		moq.Fn(n, swap)
	}
}

func (m *MoqShuffle_genType_mock) Fn(n int, swap func(i, j int)) {
	m.Moq.Scene.T.Helper()
	params := MoqShuffle_genType_params{
		N:    n,
		Swap: swap,
	}
	var results *MoqShuffle_genType_results
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
		result.DoFn(n, swap)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(n, swap)
	}
	return
}

func (m *MoqShuffle_genType) OnCall(n int, swap func(i, j int)) *MoqShuffle_genType_fnRecorder {
	return &MoqShuffle_genType_fnRecorder{
		Params: MoqShuffle_genType_params{
			N:    n,
			Swap: swap,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqShuffle_genType_fnRecorder) Any() *MoqShuffle_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqShuffle_genType_anyParams{Recorder: r}
}

func (a *MoqShuffle_genType_anyParams) N() *MoqShuffle_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqShuffle_genType_anyParams) Swap() *MoqShuffle_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqShuffle_genType_fnRecorder) Seq() *MoqShuffle_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqShuffle_genType_fnRecorder) NoSeq() *MoqShuffle_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqShuffle_genType_fnRecorder) ReturnResults() *MoqShuffle_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqShuffle_genType_doFn
		DoReturnFn MoqShuffle_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqShuffle_genType_fnRecorder) AndDo(fn MoqShuffle_genType_doFn) *MoqShuffle_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqShuffle_genType_fnRecorder) DoReturnResults(fn MoqShuffle_genType_doReturnFn) *MoqShuffle_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqShuffle_genType_doFn
		DoReturnFn MoqShuffle_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqShuffle_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqShuffle_genType_resultsByParams
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
		results = &MoqShuffle_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqShuffle_genType_paramsKey]*MoqShuffle_genType_results{},
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
		r.Results = &MoqShuffle_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqShuffle_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqShuffle_genType_fnRecorder {
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
				DoFn       MoqShuffle_genType_doFn
				DoReturnFn MoqShuffle_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqShuffle_genType) PrettyParams(params MoqShuffle_genType_params) string {
	return fmt.Sprintf("Shuffle_genType(%#v, %#v)", params.N, params.Swap)
}

func (m *MoqShuffle_genType) ParamsKey(params MoqShuffle_genType_params, anyParams uint64) MoqShuffle_genType_paramsKey {
	m.Scene.T.Helper()
	var nUsed int
	var nUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.N == moq.ParamIndexByValue {
			nUsed = params.N
		} else {
			nUsedHash = hash.DeepHash(params.N)
		}
	}
	var swapUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Swap == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The swap parameter can't be indexed by value")
		}
		swapUsedHash = hash.DeepHash(params.Swap)
	}
	return MoqShuffle_genType_paramsKey{
		Params: struct{ N int }{
			N: nUsed,
		},
		Hashes: struct {
			N    hash.Hash
			Swap hash.Hash
		}{
			N:    nUsedHash,
			Swap: swapUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqShuffle_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqShuffle_genType) AssertExpectationsMet() {
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

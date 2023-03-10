// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package rand

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Seed_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Seed_genType func(seed int64)

// MoqSeed_genType holds the state of a moq of the Seed_genType type
type MoqSeed_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSeed_genType_mock

	ResultsByParams []MoqSeed_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Seed moq.ParamIndexing
		}
	}
}

// MoqSeed_genType_mock isolates the mock interface of the Seed_genType type
type MoqSeed_genType_mock struct {
	Moq *MoqSeed_genType
}

// MoqSeed_genType_params holds the params of the Seed_genType type
type MoqSeed_genType_params struct{ Seed int64 }

// MoqSeed_genType_paramsKey holds the map key params of the Seed_genType type
type MoqSeed_genType_paramsKey struct {
	Params struct{ Seed int64 }
	Hashes struct{ Seed hash.Hash }
}

// MoqSeed_genType_resultsByParams contains the results for a given set of
// parameters for the Seed_genType type
type MoqSeed_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSeed_genType_paramsKey]*MoqSeed_genType_results
}

// MoqSeed_genType_doFn defines the type of function needed when calling AndDo
// for the Seed_genType type
type MoqSeed_genType_doFn func(seed int64)

// MoqSeed_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Seed_genType type
type MoqSeed_genType_doReturnFn func(seed int64)

// MoqSeed_genType_results holds the results of the Seed_genType type
type MoqSeed_genType_results struct {
	Params  MoqSeed_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSeed_genType_doFn
		DoReturnFn MoqSeed_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSeed_genType_fnRecorder routes recorded function calls to the
// MoqSeed_genType moq
type MoqSeed_genType_fnRecorder struct {
	Params    MoqSeed_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSeed_genType_results
	Moq       *MoqSeed_genType
}

// MoqSeed_genType_anyParams isolates the any params functions of the
// Seed_genType type
type MoqSeed_genType_anyParams struct {
	Recorder *MoqSeed_genType_fnRecorder
}

// NewMoqSeed_genType creates a new moq of the Seed_genType type
func NewMoqSeed_genType(scene *moq.Scene, config *moq.Config) *MoqSeed_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSeed_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSeed_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Seed moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Seed moq.ParamIndexing
		}{
			Seed: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Seed_genType type
func (m *MoqSeed_genType) Mock() Seed_genType {
	return func(seed int64) { m.Scene.T.Helper(); moq := &MoqSeed_genType_mock{Moq: m}; moq.Fn(seed) }
}

func (m *MoqSeed_genType_mock) Fn(seed int64) {
	m.Moq.Scene.T.Helper()
	params := MoqSeed_genType_params{
		Seed: seed,
	}
	var results *MoqSeed_genType_results
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
		result.DoFn(seed)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(seed)
	}
	return
}

func (m *MoqSeed_genType) OnCall(seed int64) *MoqSeed_genType_fnRecorder {
	return &MoqSeed_genType_fnRecorder{
		Params: MoqSeed_genType_params{
			Seed: seed,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSeed_genType_fnRecorder) Any() *MoqSeed_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSeed_genType_anyParams{Recorder: r}
}

func (a *MoqSeed_genType_anyParams) Seed() *MoqSeed_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqSeed_genType_fnRecorder) Seq() *MoqSeed_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSeed_genType_fnRecorder) NoSeq() *MoqSeed_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSeed_genType_fnRecorder) ReturnResults() *MoqSeed_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSeed_genType_doFn
		DoReturnFn MoqSeed_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSeed_genType_fnRecorder) AndDo(fn MoqSeed_genType_doFn) *MoqSeed_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSeed_genType_fnRecorder) DoReturnResults(fn MoqSeed_genType_doReturnFn) *MoqSeed_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSeed_genType_doFn
		DoReturnFn MoqSeed_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSeed_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSeed_genType_resultsByParams
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
		results = &MoqSeed_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSeed_genType_paramsKey]*MoqSeed_genType_results{},
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
		r.Results = &MoqSeed_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSeed_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSeed_genType_fnRecorder {
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
				DoFn       MoqSeed_genType_doFn
				DoReturnFn MoqSeed_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSeed_genType) PrettyParams(params MoqSeed_genType_params) string {
	return fmt.Sprintf("Seed_genType(%#v)", params.Seed)
}

func (m *MoqSeed_genType) ParamsKey(params MoqSeed_genType_params, anyParams uint64) MoqSeed_genType_paramsKey {
	m.Scene.T.Helper()
	var seedUsed int64
	var seedUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Seed == moq.ParamIndexByValue {
			seedUsed = params.Seed
		} else {
			seedUsedHash = hash.DeepHash(params.Seed)
		}
	}
	return MoqSeed_genType_paramsKey{
		Params: struct{ Seed int64 }{
			Seed: seedUsed,
		},
		Hashes: struct{ Seed hash.Hash }{
			Seed: seedUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSeed_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSeed_genType) AssertExpectationsMet() {
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

// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package rand

import (
	"fmt"
	"math/bits"
	"math/rand"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// NewSource_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type NewSource_genType func(seed int64) rand.Source

// MoqNewSource_genType holds the state of a moq of the NewSource_genType type
type MoqNewSource_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNewSource_genType_mock

	ResultsByParams []MoqNewSource_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Seed moq.ParamIndexing
		}
	}
}

// MoqNewSource_genType_mock isolates the mock interface of the
// NewSource_genType type
type MoqNewSource_genType_mock struct {
	Moq *MoqNewSource_genType
}

// MoqNewSource_genType_params holds the params of the NewSource_genType type
type MoqNewSource_genType_params struct{ Seed int64 }

// MoqNewSource_genType_paramsKey holds the map key params of the
// NewSource_genType type
type MoqNewSource_genType_paramsKey struct {
	Params struct{ Seed int64 }
	Hashes struct{ Seed hash.Hash }
}

// MoqNewSource_genType_resultsByParams contains the results for a given set of
// parameters for the NewSource_genType type
type MoqNewSource_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNewSource_genType_paramsKey]*MoqNewSource_genType_results
}

// MoqNewSource_genType_doFn defines the type of function needed when calling
// AndDo for the NewSource_genType type
type MoqNewSource_genType_doFn func(seed int64)

// MoqNewSource_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the NewSource_genType type
type MoqNewSource_genType_doReturnFn func(seed int64) rand.Source

// MoqNewSource_genType_results holds the results of the NewSource_genType type
type MoqNewSource_genType_results struct {
	Params  MoqNewSource_genType_params
	Results []struct {
		Values *struct {
			Result1 rand.Source
		}
		Sequence   uint32
		DoFn       MoqNewSource_genType_doFn
		DoReturnFn MoqNewSource_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNewSource_genType_fnRecorder routes recorded function calls to the
// MoqNewSource_genType moq
type MoqNewSource_genType_fnRecorder struct {
	Params    MoqNewSource_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNewSource_genType_results
	Moq       *MoqNewSource_genType
}

// MoqNewSource_genType_anyParams isolates the any params functions of the
// NewSource_genType type
type MoqNewSource_genType_anyParams struct {
	Recorder *MoqNewSource_genType_fnRecorder
}

// NewMoqNewSource_genType creates a new moq of the NewSource_genType type
func NewMoqNewSource_genType(scene *moq.Scene, config *moq.Config) *MoqNewSource_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNewSource_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNewSource_genType_mock{},

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

// Mock returns the moq implementation of the NewSource_genType type
func (m *MoqNewSource_genType) Mock() NewSource_genType {
	return func(seed int64) rand.Source {
		m.Scene.T.Helper()
		moq := &MoqNewSource_genType_mock{Moq: m}
		return moq.Fn(seed)
	}
}

func (m *MoqNewSource_genType_mock) Fn(seed int64) (result1 rand.Source) {
	m.Moq.Scene.T.Helper()
	params := MoqNewSource_genType_params{
		Seed: seed,
	}
	var results *MoqNewSource_genType_results
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

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(seed)
	}
	return
}

func (m *MoqNewSource_genType) OnCall(seed int64) *MoqNewSource_genType_fnRecorder {
	return &MoqNewSource_genType_fnRecorder{
		Params: MoqNewSource_genType_params{
			Seed: seed,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNewSource_genType_fnRecorder) Any() *MoqNewSource_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNewSource_genType_anyParams{Recorder: r}
}

func (a *MoqNewSource_genType_anyParams) Seed() *MoqNewSource_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqNewSource_genType_fnRecorder) Seq() *MoqNewSource_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNewSource_genType_fnRecorder) NoSeq() *MoqNewSource_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNewSource_genType_fnRecorder) ReturnResults(result1 rand.Source) *MoqNewSource_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 rand.Source
		}
		Sequence   uint32
		DoFn       MoqNewSource_genType_doFn
		DoReturnFn MoqNewSource_genType_doReturnFn
	}{
		Values: &struct {
			Result1 rand.Source
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNewSource_genType_fnRecorder) AndDo(fn MoqNewSource_genType_doFn) *MoqNewSource_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNewSource_genType_fnRecorder) DoReturnResults(fn MoqNewSource_genType_doReturnFn) *MoqNewSource_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 rand.Source
		}
		Sequence   uint32
		DoFn       MoqNewSource_genType_doFn
		DoReturnFn MoqNewSource_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNewSource_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNewSource_genType_resultsByParams
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
		results = &MoqNewSource_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNewSource_genType_paramsKey]*MoqNewSource_genType_results{},
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
		r.Results = &MoqNewSource_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNewSource_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNewSource_genType_fnRecorder {
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
					Result1 rand.Source
				}
				Sequence   uint32
				DoFn       MoqNewSource_genType_doFn
				DoReturnFn MoqNewSource_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNewSource_genType) PrettyParams(params MoqNewSource_genType_params) string {
	return fmt.Sprintf("NewSource_genType(%#v)", params.Seed)
}

func (m *MoqNewSource_genType) ParamsKey(params MoqNewSource_genType_params, anyParams uint64) MoqNewSource_genType_paramsKey {
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
	return MoqNewSource_genType_paramsKey{
		Params: struct{ Seed int64 }{
			Seed: seedUsed,
		},
		Hashes: struct{ Seed hash.Hash }{
			Seed: seedUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNewSource_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNewSource_genType) AssertExpectationsMet() {
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
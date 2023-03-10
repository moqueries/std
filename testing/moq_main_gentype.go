// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package testing

import (
	"fmt"
	"math/bits"
	"sync/atomic"
	"testing"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Main_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Main_genType func(matchString func(pat, str string) (bool, error), tests []testing.InternalTest, benchmarks []testing.InternalBenchmark, examples []testing.InternalExample)

// MoqMain_genType holds the state of a moq of the Main_genType type
type MoqMain_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqMain_genType_mock

	ResultsByParams []MoqMain_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			MatchString moq.ParamIndexing
			Tests       moq.ParamIndexing
			Benchmarks  moq.ParamIndexing
			Examples    moq.ParamIndexing
		}
	}
}

// MoqMain_genType_mock isolates the mock interface of the Main_genType type
type MoqMain_genType_mock struct {
	Moq *MoqMain_genType
}

// MoqMain_genType_params holds the params of the Main_genType type
type MoqMain_genType_params struct {
	MatchString func(pat, str string) (bool, error)
	Tests       []testing.InternalTest
	Benchmarks  []testing.InternalBenchmark
	Examples    []testing.InternalExample
}

// MoqMain_genType_paramsKey holds the map key params of the Main_genType type
type MoqMain_genType_paramsKey struct {
	Params struct{}
	Hashes struct {
		MatchString hash.Hash
		Tests       hash.Hash
		Benchmarks  hash.Hash
		Examples    hash.Hash
	}
}

// MoqMain_genType_resultsByParams contains the results for a given set of
// parameters for the Main_genType type
type MoqMain_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqMain_genType_paramsKey]*MoqMain_genType_results
}

// MoqMain_genType_doFn defines the type of function needed when calling AndDo
// for the Main_genType type
type MoqMain_genType_doFn func(matchString func(pat, str string) (bool, error), tests []testing.InternalTest, benchmarks []testing.InternalBenchmark, examples []testing.InternalExample)

// MoqMain_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Main_genType type
type MoqMain_genType_doReturnFn func(matchString func(pat, str string) (bool, error), tests []testing.InternalTest, benchmarks []testing.InternalBenchmark, examples []testing.InternalExample)

// MoqMain_genType_results holds the results of the Main_genType type
type MoqMain_genType_results struct {
	Params  MoqMain_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqMain_genType_doFn
		DoReturnFn MoqMain_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqMain_genType_fnRecorder routes recorded function calls to the
// MoqMain_genType moq
type MoqMain_genType_fnRecorder struct {
	Params    MoqMain_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqMain_genType_results
	Moq       *MoqMain_genType
}

// MoqMain_genType_anyParams isolates the any params functions of the
// Main_genType type
type MoqMain_genType_anyParams struct {
	Recorder *MoqMain_genType_fnRecorder
}

// NewMoqMain_genType creates a new moq of the Main_genType type
func NewMoqMain_genType(scene *moq.Scene, config *moq.Config) *MoqMain_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqMain_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqMain_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				MatchString moq.ParamIndexing
				Tests       moq.ParamIndexing
				Benchmarks  moq.ParamIndexing
				Examples    moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			MatchString moq.ParamIndexing
			Tests       moq.ParamIndexing
			Benchmarks  moq.ParamIndexing
			Examples    moq.ParamIndexing
		}{
			MatchString: moq.ParamIndexByHash,
			Tests:       moq.ParamIndexByHash,
			Benchmarks:  moq.ParamIndexByHash,
			Examples:    moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Main_genType type
func (m *MoqMain_genType) Mock() Main_genType {
	return func(matchString func(pat, str string) (bool, error), tests []testing.InternalTest, benchmarks []testing.InternalBenchmark, examples []testing.InternalExample) {
		m.Scene.T.Helper()
		moq := &MoqMain_genType_mock{Moq: m}
		moq.Fn(matchString, tests, benchmarks, examples)
	}
}

func (m *MoqMain_genType_mock) Fn(matchString func(pat, str string) (bool, error), tests []testing.InternalTest, benchmarks []testing.InternalBenchmark, examples []testing.InternalExample) {
	m.Moq.Scene.T.Helper()
	params := MoqMain_genType_params{
		MatchString: matchString,
		Tests:       tests,
		Benchmarks:  benchmarks,
		Examples:    examples,
	}
	var results *MoqMain_genType_results
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
		result.DoFn(matchString, tests, benchmarks, examples)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(matchString, tests, benchmarks, examples)
	}
	return
}

func (m *MoqMain_genType) OnCall(matchString func(pat, str string) (bool, error), tests []testing.InternalTest, benchmarks []testing.InternalBenchmark, examples []testing.InternalExample) *MoqMain_genType_fnRecorder {
	return &MoqMain_genType_fnRecorder{
		Params: MoqMain_genType_params{
			MatchString: matchString,
			Tests:       tests,
			Benchmarks:  benchmarks,
			Examples:    examples,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqMain_genType_fnRecorder) Any() *MoqMain_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqMain_genType_anyParams{Recorder: r}
}

func (a *MoqMain_genType_anyParams) MatchString() *MoqMain_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqMain_genType_anyParams) Tests() *MoqMain_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqMain_genType_anyParams) Benchmarks() *MoqMain_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (a *MoqMain_genType_anyParams) Examples() *MoqMain_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 3
	return a.Recorder
}

func (r *MoqMain_genType_fnRecorder) Seq() *MoqMain_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqMain_genType_fnRecorder) NoSeq() *MoqMain_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqMain_genType_fnRecorder) ReturnResults() *MoqMain_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqMain_genType_doFn
		DoReturnFn MoqMain_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqMain_genType_fnRecorder) AndDo(fn MoqMain_genType_doFn) *MoqMain_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqMain_genType_fnRecorder) DoReturnResults(fn MoqMain_genType_doReturnFn) *MoqMain_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqMain_genType_doFn
		DoReturnFn MoqMain_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqMain_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqMain_genType_resultsByParams
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
		results = &MoqMain_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqMain_genType_paramsKey]*MoqMain_genType_results{},
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
		r.Results = &MoqMain_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqMain_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqMain_genType_fnRecorder {
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
				DoFn       MoqMain_genType_doFn
				DoReturnFn MoqMain_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqMain_genType) PrettyParams(params MoqMain_genType_params) string {
	return fmt.Sprintf("Main_genType(%#v, %#v, %#v, %#v)", moq.FnString(params.MatchString), params.Tests, params.Benchmarks, params.Examples)
}

func (m *MoqMain_genType) ParamsKey(params MoqMain_genType_params, anyParams uint64) MoqMain_genType_paramsKey {
	m.Scene.T.Helper()
	var matchStringUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.MatchString == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The matchString parameter can't be indexed by value")
		}
		matchStringUsedHash = hash.DeepHash(params.MatchString)
	}
	var testsUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Tests == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The tests parameter can't be indexed by value")
		}
		testsUsedHash = hash.DeepHash(params.Tests)
	}
	var benchmarksUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Benchmarks == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The benchmarks parameter can't be indexed by value")
		}
		benchmarksUsedHash = hash.DeepHash(params.Benchmarks)
	}
	var examplesUsedHash hash.Hash
	if anyParams&(1<<3) == 0 {
		if m.Runtime.ParameterIndexing.Examples == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The examples parameter can't be indexed by value")
		}
		examplesUsedHash = hash.DeepHash(params.Examples)
	}
	return MoqMain_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct {
			MatchString hash.Hash
			Tests       hash.Hash
			Benchmarks  hash.Hash
			Examples    hash.Hash
		}{
			MatchString: matchStringUsedHash,
			Tests:       testsUsedHash,
			Benchmarks:  benchmarksUsedHash,
			Examples:    examplesUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqMain_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqMain_genType) AssertExpectationsMet() {
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

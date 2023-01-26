// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package doc

import (
	"fmt"
	"go/ast"
	"go/doc"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Examples_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Examples_genType func(testFiles ...*ast.File) []*doc.Example

// MoqExamples_genType holds the state of a moq of the Examples_genType type
type MoqExamples_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqExamples_genType_mock

	ResultsByParams []MoqExamples_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			TestFiles moq.ParamIndexing
		}
	}
}

// MoqExamples_genType_mock isolates the mock interface of the Examples_genType
// type
type MoqExamples_genType_mock struct {
	Moq *MoqExamples_genType
}

// MoqExamples_genType_params holds the params of the Examples_genType type
type MoqExamples_genType_params struct{ TestFiles []*ast.File }

// MoqExamples_genType_paramsKey holds the map key params of the
// Examples_genType type
type MoqExamples_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ TestFiles hash.Hash }
}

// MoqExamples_genType_resultsByParams contains the results for a given set of
// parameters for the Examples_genType type
type MoqExamples_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqExamples_genType_paramsKey]*MoqExamples_genType_results
}

// MoqExamples_genType_doFn defines the type of function needed when calling
// AndDo for the Examples_genType type
type MoqExamples_genType_doFn func(testFiles ...*ast.File)

// MoqExamples_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Examples_genType type
type MoqExamples_genType_doReturnFn func(testFiles ...*ast.File) []*doc.Example

// MoqExamples_genType_results holds the results of the Examples_genType type
type MoqExamples_genType_results struct {
	Params  MoqExamples_genType_params
	Results []struct {
		Values *struct {
			Result1 []*doc.Example
		}
		Sequence   uint32
		DoFn       MoqExamples_genType_doFn
		DoReturnFn MoqExamples_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqExamples_genType_fnRecorder routes recorded function calls to the
// MoqExamples_genType moq
type MoqExamples_genType_fnRecorder struct {
	Params    MoqExamples_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqExamples_genType_results
	Moq       *MoqExamples_genType
}

// MoqExamples_genType_anyParams isolates the any params functions of the
// Examples_genType type
type MoqExamples_genType_anyParams struct {
	Recorder *MoqExamples_genType_fnRecorder
}

// NewMoqExamples_genType creates a new moq of the Examples_genType type
func NewMoqExamples_genType(scene *moq.Scene, config *moq.Config) *MoqExamples_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqExamples_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqExamples_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				TestFiles moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			TestFiles moq.ParamIndexing
		}{
			TestFiles: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Examples_genType type
func (m *MoqExamples_genType) Mock() Examples_genType {
	return func(testFiles ...*ast.File) []*doc.Example {
		m.Scene.T.Helper()
		moq := &MoqExamples_genType_mock{Moq: m}
		return moq.Fn(testFiles...)
	}
}

func (m *MoqExamples_genType_mock) Fn(testFiles ...*ast.File) (result1 []*doc.Example) {
	m.Moq.Scene.T.Helper()
	params := MoqExamples_genType_params{
		TestFiles: testFiles,
	}
	var results *MoqExamples_genType_results
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
		result.DoFn(testFiles...)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(testFiles...)
	}
	return
}

func (m *MoqExamples_genType) OnCall(testFiles ...*ast.File) *MoqExamples_genType_fnRecorder {
	return &MoqExamples_genType_fnRecorder{
		Params: MoqExamples_genType_params{
			TestFiles: testFiles,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqExamples_genType_fnRecorder) Any() *MoqExamples_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqExamples_genType_anyParams{Recorder: r}
}

func (a *MoqExamples_genType_anyParams) TestFiles() *MoqExamples_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqExamples_genType_fnRecorder) Seq() *MoqExamples_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqExamples_genType_fnRecorder) NoSeq() *MoqExamples_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqExamples_genType_fnRecorder) ReturnResults(result1 []*doc.Example) *MoqExamples_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []*doc.Example
		}
		Sequence   uint32
		DoFn       MoqExamples_genType_doFn
		DoReturnFn MoqExamples_genType_doReturnFn
	}{
		Values: &struct {
			Result1 []*doc.Example
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqExamples_genType_fnRecorder) AndDo(fn MoqExamples_genType_doFn) *MoqExamples_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqExamples_genType_fnRecorder) DoReturnResults(fn MoqExamples_genType_doReturnFn) *MoqExamples_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []*doc.Example
		}
		Sequence   uint32
		DoFn       MoqExamples_genType_doFn
		DoReturnFn MoqExamples_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqExamples_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqExamples_genType_resultsByParams
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
		results = &MoqExamples_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqExamples_genType_paramsKey]*MoqExamples_genType_results{},
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
		r.Results = &MoqExamples_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqExamples_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqExamples_genType_fnRecorder {
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
					Result1 []*doc.Example
				}
				Sequence   uint32
				DoFn       MoqExamples_genType_doFn
				DoReturnFn MoqExamples_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqExamples_genType) PrettyParams(params MoqExamples_genType_params) string {
	return fmt.Sprintf("Examples_genType(%#v)", params.TestFiles)
}

func (m *MoqExamples_genType) ParamsKey(params MoqExamples_genType_params, anyParams uint64) MoqExamples_genType_paramsKey {
	m.Scene.T.Helper()
	var testFilesUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.TestFiles == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The testFiles parameter can't be indexed by value")
		}
		testFilesUsedHash = hash.DeepHash(params.TestFiles)
	}
	return MoqExamples_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ TestFiles hash.Hash }{
			TestFiles: testFilesUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqExamples_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqExamples_genType) AssertExpectationsMet() {
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
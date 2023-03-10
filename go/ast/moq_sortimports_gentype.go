// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package ast

import (
	"fmt"
	"go/ast"
	"go/token"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// SortImports_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type SortImports_genType func(fset *token.FileSet, f *ast.File)

// MoqSortImports_genType holds the state of a moq of the SortImports_genType
// type
type MoqSortImports_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSortImports_genType_mock

	ResultsByParams []MoqSortImports_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Fset moq.ParamIndexing
			F    moq.ParamIndexing
		}
	}
}

// MoqSortImports_genType_mock isolates the mock interface of the
// SortImports_genType type
type MoqSortImports_genType_mock struct {
	Moq *MoqSortImports_genType
}

// MoqSortImports_genType_params holds the params of the SortImports_genType
// type
type MoqSortImports_genType_params struct {
	Fset *token.FileSet
	F    *ast.File
}

// MoqSortImports_genType_paramsKey holds the map key params of the
// SortImports_genType type
type MoqSortImports_genType_paramsKey struct {
	Params struct {
		Fset *token.FileSet
		F    *ast.File
	}
	Hashes struct {
		Fset hash.Hash
		F    hash.Hash
	}
}

// MoqSortImports_genType_resultsByParams contains the results for a given set
// of parameters for the SortImports_genType type
type MoqSortImports_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSortImports_genType_paramsKey]*MoqSortImports_genType_results
}

// MoqSortImports_genType_doFn defines the type of function needed when calling
// AndDo for the SortImports_genType type
type MoqSortImports_genType_doFn func(fset *token.FileSet, f *ast.File)

// MoqSortImports_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the SortImports_genType type
type MoqSortImports_genType_doReturnFn func(fset *token.FileSet, f *ast.File)

// MoqSortImports_genType_results holds the results of the SortImports_genType
// type
type MoqSortImports_genType_results struct {
	Params  MoqSortImports_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSortImports_genType_doFn
		DoReturnFn MoqSortImports_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSortImports_genType_fnRecorder routes recorded function calls to the
// MoqSortImports_genType moq
type MoqSortImports_genType_fnRecorder struct {
	Params    MoqSortImports_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSortImports_genType_results
	Moq       *MoqSortImports_genType
}

// MoqSortImports_genType_anyParams isolates the any params functions of the
// SortImports_genType type
type MoqSortImports_genType_anyParams struct {
	Recorder *MoqSortImports_genType_fnRecorder
}

// NewMoqSortImports_genType creates a new moq of the SortImports_genType type
func NewMoqSortImports_genType(scene *moq.Scene, config *moq.Config) *MoqSortImports_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSortImports_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSortImports_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Fset moq.ParamIndexing
				F    moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Fset moq.ParamIndexing
			F    moq.ParamIndexing
		}{
			Fset: moq.ParamIndexByHash,
			F:    moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the SortImports_genType type
func (m *MoqSortImports_genType) Mock() SortImports_genType {
	return func(fset *token.FileSet, f *ast.File) {
		m.Scene.T.Helper()
		moq := &MoqSortImports_genType_mock{Moq: m}
		moq.Fn(fset, f)
	}
}

func (m *MoqSortImports_genType_mock) Fn(fset *token.FileSet, f *ast.File) {
	m.Moq.Scene.T.Helper()
	params := MoqSortImports_genType_params{
		Fset: fset,
		F:    f,
	}
	var results *MoqSortImports_genType_results
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
		result.DoFn(fset, f)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(fset, f)
	}
	return
}

func (m *MoqSortImports_genType) OnCall(fset *token.FileSet, f *ast.File) *MoqSortImports_genType_fnRecorder {
	return &MoqSortImports_genType_fnRecorder{
		Params: MoqSortImports_genType_params{
			Fset: fset,
			F:    f,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSortImports_genType_fnRecorder) Any() *MoqSortImports_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSortImports_genType_anyParams{Recorder: r}
}

func (a *MoqSortImports_genType_anyParams) Fset() *MoqSortImports_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqSortImports_genType_anyParams) F() *MoqSortImports_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqSortImports_genType_fnRecorder) Seq() *MoqSortImports_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSortImports_genType_fnRecorder) NoSeq() *MoqSortImports_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSortImports_genType_fnRecorder) ReturnResults() *MoqSortImports_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSortImports_genType_doFn
		DoReturnFn MoqSortImports_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqSortImports_genType_fnRecorder) AndDo(fn MoqSortImports_genType_doFn) *MoqSortImports_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSortImports_genType_fnRecorder) DoReturnResults(fn MoqSortImports_genType_doReturnFn) *MoqSortImports_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqSortImports_genType_doFn
		DoReturnFn MoqSortImports_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSortImports_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSortImports_genType_resultsByParams
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
		results = &MoqSortImports_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSortImports_genType_paramsKey]*MoqSortImports_genType_results{},
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
		r.Results = &MoqSortImports_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSortImports_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSortImports_genType_fnRecorder {
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
				DoFn       MoqSortImports_genType_doFn
				DoReturnFn MoqSortImports_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSortImports_genType) PrettyParams(params MoqSortImports_genType_params) string {
	return fmt.Sprintf("SortImports_genType(%#v, %#v)", params.Fset, params.F)
}

func (m *MoqSortImports_genType) ParamsKey(params MoqSortImports_genType_params, anyParams uint64) MoqSortImports_genType_paramsKey {
	m.Scene.T.Helper()
	var fsetUsed *token.FileSet
	var fsetUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Fset == moq.ParamIndexByValue {
			fsetUsed = params.Fset
		} else {
			fsetUsedHash = hash.DeepHash(params.Fset)
		}
	}
	var fUsed *ast.File
	var fUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.F == moq.ParamIndexByValue {
			fUsed = params.F
		} else {
			fUsedHash = hash.DeepHash(params.F)
		}
	}
	return MoqSortImports_genType_paramsKey{
		Params: struct {
			Fset *token.FileSet
			F    *ast.File
		}{
			Fset: fsetUsed,
			F:    fUsed,
		},
		Hashes: struct {
			Fset hash.Hash
			F    hash.Hash
		}{
			Fset: fsetUsedHash,
			F:    fUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSortImports_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSortImports_genType) AssertExpectationsMet() {
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

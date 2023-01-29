// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package doc

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/token"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// NewFromFiles_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type NewFromFiles_genType func(fset *token.FileSet, files []*ast.File, importPath string, opts ...any) (*doc.Package, error)

// MoqNewFromFiles_genType holds the state of a moq of the NewFromFiles_genType
// type
type MoqNewFromFiles_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNewFromFiles_genType_mock

	ResultsByParams []MoqNewFromFiles_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Fset       moq.ParamIndexing
			Files      moq.ParamIndexing
			ImportPath moq.ParamIndexing
			Opts       moq.ParamIndexing
		}
	}
}

// MoqNewFromFiles_genType_mock isolates the mock interface of the
// NewFromFiles_genType type
type MoqNewFromFiles_genType_mock struct {
	Moq *MoqNewFromFiles_genType
}

// MoqNewFromFiles_genType_params holds the params of the NewFromFiles_genType
// type
type MoqNewFromFiles_genType_params struct {
	Fset       *token.FileSet
	Files      []*ast.File
	ImportPath string
	Opts       []any
}

// MoqNewFromFiles_genType_paramsKey holds the map key params of the
// NewFromFiles_genType type
type MoqNewFromFiles_genType_paramsKey struct {
	Params struct {
		Fset       *token.FileSet
		ImportPath string
	}
	Hashes struct {
		Fset       hash.Hash
		Files      hash.Hash
		ImportPath hash.Hash
		Opts       hash.Hash
	}
}

// MoqNewFromFiles_genType_resultsByParams contains the results for a given set
// of parameters for the NewFromFiles_genType type
type MoqNewFromFiles_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNewFromFiles_genType_paramsKey]*MoqNewFromFiles_genType_results
}

// MoqNewFromFiles_genType_doFn defines the type of function needed when
// calling AndDo for the NewFromFiles_genType type
type MoqNewFromFiles_genType_doFn func(fset *token.FileSet, files []*ast.File, importPath string, opts ...any)

// MoqNewFromFiles_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the NewFromFiles_genType type
type MoqNewFromFiles_genType_doReturnFn func(fset *token.FileSet, files []*ast.File, importPath string, opts ...any) (*doc.Package, error)

// MoqNewFromFiles_genType_results holds the results of the
// NewFromFiles_genType type
type MoqNewFromFiles_genType_results struct {
	Params  MoqNewFromFiles_genType_params
	Results []struct {
		Values *struct {
			Result1 *doc.Package
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqNewFromFiles_genType_doFn
		DoReturnFn MoqNewFromFiles_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNewFromFiles_genType_fnRecorder routes recorded function calls to the
// MoqNewFromFiles_genType moq
type MoqNewFromFiles_genType_fnRecorder struct {
	Params    MoqNewFromFiles_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNewFromFiles_genType_results
	Moq       *MoqNewFromFiles_genType
}

// MoqNewFromFiles_genType_anyParams isolates the any params functions of the
// NewFromFiles_genType type
type MoqNewFromFiles_genType_anyParams struct {
	Recorder *MoqNewFromFiles_genType_fnRecorder
}

// NewMoqNewFromFiles_genType creates a new moq of the NewFromFiles_genType
// type
func NewMoqNewFromFiles_genType(scene *moq.Scene, config *moq.Config) *MoqNewFromFiles_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNewFromFiles_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNewFromFiles_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Fset       moq.ParamIndexing
				Files      moq.ParamIndexing
				ImportPath moq.ParamIndexing
				Opts       moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Fset       moq.ParamIndexing
			Files      moq.ParamIndexing
			ImportPath moq.ParamIndexing
			Opts       moq.ParamIndexing
		}{
			Fset:       moq.ParamIndexByHash,
			Files:      moq.ParamIndexByHash,
			ImportPath: moq.ParamIndexByValue,
			Opts:       moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the NewFromFiles_genType type
func (m *MoqNewFromFiles_genType) Mock() NewFromFiles_genType {
	return func(fset *token.FileSet, files []*ast.File, importPath string, opts ...any) (*doc.Package, error) {
		m.Scene.T.Helper()
		moq := &MoqNewFromFiles_genType_mock{Moq: m}
		return moq.Fn(fset, files, importPath, opts...)
	}
}

func (m *MoqNewFromFiles_genType_mock) Fn(fset *token.FileSet, files []*ast.File, importPath string, opts ...any) (result1 *doc.Package, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqNewFromFiles_genType_params{
		Fset:       fset,
		Files:      files,
		ImportPath: importPath,
		Opts:       opts,
	}
	var results *MoqNewFromFiles_genType_results
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
		result.DoFn(fset, files, importPath, opts...)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(fset, files, importPath, opts...)
	}
	return
}

func (m *MoqNewFromFiles_genType) OnCall(fset *token.FileSet, files []*ast.File, importPath string, opts ...any) *MoqNewFromFiles_genType_fnRecorder {
	return &MoqNewFromFiles_genType_fnRecorder{
		Params: MoqNewFromFiles_genType_params{
			Fset:       fset,
			Files:      files,
			ImportPath: importPath,
			Opts:       opts,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNewFromFiles_genType_fnRecorder) Any() *MoqNewFromFiles_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNewFromFiles_genType_anyParams{Recorder: r}
}

func (a *MoqNewFromFiles_genType_anyParams) Fset() *MoqNewFromFiles_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqNewFromFiles_genType_anyParams) Files() *MoqNewFromFiles_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqNewFromFiles_genType_anyParams) ImportPath() *MoqNewFromFiles_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (a *MoqNewFromFiles_genType_anyParams) Opts() *MoqNewFromFiles_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 3
	return a.Recorder
}

func (r *MoqNewFromFiles_genType_fnRecorder) Seq() *MoqNewFromFiles_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNewFromFiles_genType_fnRecorder) NoSeq() *MoqNewFromFiles_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNewFromFiles_genType_fnRecorder) ReturnResults(result1 *doc.Package, result2 error) *MoqNewFromFiles_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *doc.Package
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqNewFromFiles_genType_doFn
		DoReturnFn MoqNewFromFiles_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *doc.Package
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNewFromFiles_genType_fnRecorder) AndDo(fn MoqNewFromFiles_genType_doFn) *MoqNewFromFiles_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNewFromFiles_genType_fnRecorder) DoReturnResults(fn MoqNewFromFiles_genType_doReturnFn) *MoqNewFromFiles_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *doc.Package
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqNewFromFiles_genType_doFn
		DoReturnFn MoqNewFromFiles_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNewFromFiles_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNewFromFiles_genType_resultsByParams
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
		results = &MoqNewFromFiles_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNewFromFiles_genType_paramsKey]*MoqNewFromFiles_genType_results{},
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
		r.Results = &MoqNewFromFiles_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNewFromFiles_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNewFromFiles_genType_fnRecorder {
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
					Result1 *doc.Package
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqNewFromFiles_genType_doFn
				DoReturnFn MoqNewFromFiles_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNewFromFiles_genType) PrettyParams(params MoqNewFromFiles_genType_params) string {
	return fmt.Sprintf("NewFromFiles_genType(%#v, %#v, %#v, %#v)", params.Fset, params.Files, params.ImportPath, params.Opts)
}

func (m *MoqNewFromFiles_genType) ParamsKey(params MoqNewFromFiles_genType_params, anyParams uint64) MoqNewFromFiles_genType_paramsKey {
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
	var filesUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Files == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The files parameter can't be indexed by value")
		}
		filesUsedHash = hash.DeepHash(params.Files)
	}
	var importPathUsed string
	var importPathUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.ImportPath == moq.ParamIndexByValue {
			importPathUsed = params.ImportPath
		} else {
			importPathUsedHash = hash.DeepHash(params.ImportPath)
		}
	}
	var optsUsedHash hash.Hash
	if anyParams&(1<<3) == 0 {
		if m.Runtime.ParameterIndexing.Opts == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The opts parameter can't be indexed by value")
		}
		optsUsedHash = hash.DeepHash(params.Opts)
	}
	return MoqNewFromFiles_genType_paramsKey{
		Params: struct {
			Fset       *token.FileSet
			ImportPath string
		}{
			Fset:       fsetUsed,
			ImportPath: importPathUsed,
		},
		Hashes: struct {
			Fset       hash.Hash
			Files      hash.Hash
			ImportPath hash.Hash
			Opts       hash.Hash
		}{
			Fset:       fsetUsedHash,
			Files:      filesUsedHash,
			ImportPath: importPathUsedHash,
			Opts:       optsUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNewFromFiles_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNewFromFiles_genType) AssertExpectationsMet() {
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

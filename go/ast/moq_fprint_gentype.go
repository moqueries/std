// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package ast

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Fprint_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Fprint_genType func(w io.Writer, fset *token.FileSet, x any, f ast.FieldFilter) error

// MoqFprint_genType holds the state of a moq of the Fprint_genType type
type MoqFprint_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqFprint_genType_mock

	ResultsByParams []MoqFprint_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			W    moq.ParamIndexing
			Fset moq.ParamIndexing
			X    moq.ParamIndexing
			F    moq.ParamIndexing
		}
	}
}

// MoqFprint_genType_mock isolates the mock interface of the Fprint_genType
// type
type MoqFprint_genType_mock struct {
	Moq *MoqFprint_genType
}

// MoqFprint_genType_params holds the params of the Fprint_genType type
type MoqFprint_genType_params struct {
	W    io.Writer
	Fset *token.FileSet
	X    any
	F    ast.FieldFilter
}

// MoqFprint_genType_paramsKey holds the map key params of the Fprint_genType
// type
type MoqFprint_genType_paramsKey struct {
	Params struct {
		W    io.Writer
		Fset *token.FileSet
		X    any
	}
	Hashes struct {
		W    hash.Hash
		Fset hash.Hash
		X    hash.Hash
		F    hash.Hash
	}
}

// MoqFprint_genType_resultsByParams contains the results for a given set of
// parameters for the Fprint_genType type
type MoqFprint_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFprint_genType_paramsKey]*MoqFprint_genType_results
}

// MoqFprint_genType_doFn defines the type of function needed when calling
// AndDo for the Fprint_genType type
type MoqFprint_genType_doFn func(w io.Writer, fset *token.FileSet, x any, f ast.FieldFilter)

// MoqFprint_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Fprint_genType type
type MoqFprint_genType_doReturnFn func(w io.Writer, fset *token.FileSet, x any, f ast.FieldFilter) error

// MoqFprint_genType_results holds the results of the Fprint_genType type
type MoqFprint_genType_results struct {
	Params  MoqFprint_genType_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqFprint_genType_doFn
		DoReturnFn MoqFprint_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFprint_genType_fnRecorder routes recorded function calls to the
// MoqFprint_genType moq
type MoqFprint_genType_fnRecorder struct {
	Params    MoqFprint_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFprint_genType_results
	Moq       *MoqFprint_genType
}

// MoqFprint_genType_anyParams isolates the any params functions of the
// Fprint_genType type
type MoqFprint_genType_anyParams struct {
	Recorder *MoqFprint_genType_fnRecorder
}

// NewMoqFprint_genType creates a new moq of the Fprint_genType type
func NewMoqFprint_genType(scene *moq.Scene, config *moq.Config) *MoqFprint_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqFprint_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqFprint_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				W    moq.ParamIndexing
				Fset moq.ParamIndexing
				X    moq.ParamIndexing
				F    moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			W    moq.ParamIndexing
			Fset moq.ParamIndexing
			X    moq.ParamIndexing
			F    moq.ParamIndexing
		}{
			W:    moq.ParamIndexByHash,
			Fset: moq.ParamIndexByHash,
			X:    moq.ParamIndexByValue,
			F:    moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Fprint_genType type
func (m *MoqFprint_genType) Mock() Fprint_genType {
	return func(w io.Writer, fset *token.FileSet, x any, f ast.FieldFilter) error {
		m.Scene.T.Helper()
		moq := &MoqFprint_genType_mock{Moq: m}
		return moq.Fn(w, fset, x, f)
	}
}

func (m *MoqFprint_genType_mock) Fn(w io.Writer, fset *token.FileSet, x any, f ast.FieldFilter) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqFprint_genType_params{
		W:    w,
		Fset: fset,
		X:    x,
		F:    f,
	}
	var results *MoqFprint_genType_results
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
		result.DoFn(w, fset, x, f)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(w, fset, x, f)
	}
	return
}

func (m *MoqFprint_genType) OnCall(w io.Writer, fset *token.FileSet, x any, f ast.FieldFilter) *MoqFprint_genType_fnRecorder {
	return &MoqFprint_genType_fnRecorder{
		Params: MoqFprint_genType_params{
			W:    w,
			Fset: fset,
			X:    x,
			F:    f,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqFprint_genType_fnRecorder) Any() *MoqFprint_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqFprint_genType_anyParams{Recorder: r}
}

func (a *MoqFprint_genType_anyParams) W() *MoqFprint_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqFprint_genType_anyParams) Fset() *MoqFprint_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqFprint_genType_anyParams) X() *MoqFprint_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (a *MoqFprint_genType_anyParams) F() *MoqFprint_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 3
	return a.Recorder
}

func (r *MoqFprint_genType_fnRecorder) Seq() *MoqFprint_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFprint_genType_fnRecorder) NoSeq() *MoqFprint_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFprint_genType_fnRecorder) ReturnResults(result1 error) *MoqFprint_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqFprint_genType_doFn
		DoReturnFn MoqFprint_genType_doReturnFn
	}{
		Values: &struct {
			Result1 error
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqFprint_genType_fnRecorder) AndDo(fn MoqFprint_genType_doFn) *MoqFprint_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFprint_genType_fnRecorder) DoReturnResults(fn MoqFprint_genType_doReturnFn) *MoqFprint_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqFprint_genType_doFn
		DoReturnFn MoqFprint_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFprint_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFprint_genType_resultsByParams
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
		results = &MoqFprint_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFprint_genType_paramsKey]*MoqFprint_genType_results{},
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
		r.Results = &MoqFprint_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFprint_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFprint_genType_fnRecorder {
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
					Result1 error
				}
				Sequence   uint32
				DoFn       MoqFprint_genType_doFn
				DoReturnFn MoqFprint_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFprint_genType) PrettyParams(params MoqFprint_genType_params) string {
	return fmt.Sprintf("Fprint_genType(%#v, %#v, %#v, %#v)", params.W, params.Fset, params.X, params.F)
}

func (m *MoqFprint_genType) ParamsKey(params MoqFprint_genType_params, anyParams uint64) MoqFprint_genType_paramsKey {
	m.Scene.T.Helper()
	var wUsed io.Writer
	var wUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.W == moq.ParamIndexByValue {
			wUsed = params.W
		} else {
			wUsedHash = hash.DeepHash(params.W)
		}
	}
	var fsetUsed *token.FileSet
	var fsetUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Fset == moq.ParamIndexByValue {
			fsetUsed = params.Fset
		} else {
			fsetUsedHash = hash.DeepHash(params.Fset)
		}
	}
	var xUsed any
	var xUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.X == moq.ParamIndexByValue {
			xUsed = params.X
		} else {
			xUsedHash = hash.DeepHash(params.X)
		}
	}
	var fUsedHash hash.Hash
	if anyParams&(1<<3) == 0 {
		if m.Runtime.ParameterIndexing.F == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The f parameter can't be indexed by value")
		}
		fUsedHash = hash.DeepHash(params.F)
	}
	return MoqFprint_genType_paramsKey{
		Params: struct {
			W    io.Writer
			Fset *token.FileSet
			X    any
		}{
			W:    wUsed,
			Fset: fsetUsed,
			X:    xUsed,
		},
		Hashes: struct {
			W    hash.Hash
			Fset hash.Hash
			X    hash.Hash
			F    hash.Hash
		}{
			W:    wUsedHash,
			Fset: fsetUsedHash,
			X:    xUsedHash,
			F:    fUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqFprint_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqFprint_genType) AssertExpectationsMet() {
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

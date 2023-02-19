// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package types

import (
	"fmt"
	"go/token"
	"go/types"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// NewChecker_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type NewChecker_genType func(conf *types.Config, fset *token.FileSet, pkg *types.Package, info *types.Info) *types.Checker

// MoqNewChecker_genType holds the state of a moq of the NewChecker_genType
// type
type MoqNewChecker_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNewChecker_genType_mock

	ResultsByParams []MoqNewChecker_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Conf moq.ParamIndexing
			Fset moq.ParamIndexing
			Pkg  moq.ParamIndexing
			Info moq.ParamIndexing
		}
	}
}

// MoqNewChecker_genType_mock isolates the mock interface of the
// NewChecker_genType type
type MoqNewChecker_genType_mock struct {
	Moq *MoqNewChecker_genType
}

// MoqNewChecker_genType_params holds the params of the NewChecker_genType type
type MoqNewChecker_genType_params struct {
	Conf *types.Config
	Fset *token.FileSet
	Pkg  *types.Package
	Info *types.Info
}

// MoqNewChecker_genType_paramsKey holds the map key params of the
// NewChecker_genType type
type MoqNewChecker_genType_paramsKey struct {
	Params struct {
		Conf *types.Config
		Fset *token.FileSet
		Pkg  *types.Package
		Info *types.Info
	}
	Hashes struct {
		Conf hash.Hash
		Fset hash.Hash
		Pkg  hash.Hash
		Info hash.Hash
	}
}

// MoqNewChecker_genType_resultsByParams contains the results for a given set
// of parameters for the NewChecker_genType type
type MoqNewChecker_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNewChecker_genType_paramsKey]*MoqNewChecker_genType_results
}

// MoqNewChecker_genType_doFn defines the type of function needed when calling
// AndDo for the NewChecker_genType type
type MoqNewChecker_genType_doFn func(conf *types.Config, fset *token.FileSet, pkg *types.Package, info *types.Info)

// MoqNewChecker_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the NewChecker_genType type
type MoqNewChecker_genType_doReturnFn func(conf *types.Config, fset *token.FileSet, pkg *types.Package, info *types.Info) *types.Checker

// MoqNewChecker_genType_results holds the results of the NewChecker_genType
// type
type MoqNewChecker_genType_results struct {
	Params  MoqNewChecker_genType_params
	Results []struct {
		Values *struct {
			Result1 *types.Checker
		}
		Sequence   uint32
		DoFn       MoqNewChecker_genType_doFn
		DoReturnFn MoqNewChecker_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNewChecker_genType_fnRecorder routes recorded function calls to the
// MoqNewChecker_genType moq
type MoqNewChecker_genType_fnRecorder struct {
	Params    MoqNewChecker_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNewChecker_genType_results
	Moq       *MoqNewChecker_genType
}

// MoqNewChecker_genType_anyParams isolates the any params functions of the
// NewChecker_genType type
type MoqNewChecker_genType_anyParams struct {
	Recorder *MoqNewChecker_genType_fnRecorder
}

// NewMoqNewChecker_genType creates a new moq of the NewChecker_genType type
func NewMoqNewChecker_genType(scene *moq.Scene, config *moq.Config) *MoqNewChecker_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNewChecker_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNewChecker_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Conf moq.ParamIndexing
				Fset moq.ParamIndexing
				Pkg  moq.ParamIndexing
				Info moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Conf moq.ParamIndexing
			Fset moq.ParamIndexing
			Pkg  moq.ParamIndexing
			Info moq.ParamIndexing
		}{
			Conf: moq.ParamIndexByHash,
			Fset: moq.ParamIndexByHash,
			Pkg:  moq.ParamIndexByHash,
			Info: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the NewChecker_genType type
func (m *MoqNewChecker_genType) Mock() NewChecker_genType {
	return func(conf *types.Config, fset *token.FileSet, pkg *types.Package, info *types.Info) *types.Checker {
		m.Scene.T.Helper()
		moq := &MoqNewChecker_genType_mock{Moq: m}
		return moq.Fn(conf, fset, pkg, info)
	}
}

func (m *MoqNewChecker_genType_mock) Fn(conf *types.Config, fset *token.FileSet, pkg *types.Package, info *types.Info) (result1 *types.Checker) {
	m.Moq.Scene.T.Helper()
	params := MoqNewChecker_genType_params{
		Conf: conf,
		Fset: fset,
		Pkg:  pkg,
		Info: info,
	}
	var results *MoqNewChecker_genType_results
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
		result.DoFn(conf, fset, pkg, info)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(conf, fset, pkg, info)
	}
	return
}

func (m *MoqNewChecker_genType) OnCall(conf *types.Config, fset *token.FileSet, pkg *types.Package, info *types.Info) *MoqNewChecker_genType_fnRecorder {
	return &MoqNewChecker_genType_fnRecorder{
		Params: MoqNewChecker_genType_params{
			Conf: conf,
			Fset: fset,
			Pkg:  pkg,
			Info: info,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNewChecker_genType_fnRecorder) Any() *MoqNewChecker_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNewChecker_genType_anyParams{Recorder: r}
}

func (a *MoqNewChecker_genType_anyParams) Conf() *MoqNewChecker_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqNewChecker_genType_anyParams) Fset() *MoqNewChecker_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqNewChecker_genType_anyParams) Pkg() *MoqNewChecker_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (a *MoqNewChecker_genType_anyParams) Info() *MoqNewChecker_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 3
	return a.Recorder
}

func (r *MoqNewChecker_genType_fnRecorder) Seq() *MoqNewChecker_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNewChecker_genType_fnRecorder) NoSeq() *MoqNewChecker_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNewChecker_genType_fnRecorder) ReturnResults(result1 *types.Checker) *MoqNewChecker_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *types.Checker
		}
		Sequence   uint32
		DoFn       MoqNewChecker_genType_doFn
		DoReturnFn MoqNewChecker_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *types.Checker
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNewChecker_genType_fnRecorder) AndDo(fn MoqNewChecker_genType_doFn) *MoqNewChecker_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNewChecker_genType_fnRecorder) DoReturnResults(fn MoqNewChecker_genType_doReturnFn) *MoqNewChecker_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *types.Checker
		}
		Sequence   uint32
		DoFn       MoqNewChecker_genType_doFn
		DoReturnFn MoqNewChecker_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNewChecker_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNewChecker_genType_resultsByParams
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
		results = &MoqNewChecker_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNewChecker_genType_paramsKey]*MoqNewChecker_genType_results{},
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
		r.Results = &MoqNewChecker_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNewChecker_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNewChecker_genType_fnRecorder {
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
					Result1 *types.Checker
				}
				Sequence   uint32
				DoFn       MoqNewChecker_genType_doFn
				DoReturnFn MoqNewChecker_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNewChecker_genType) PrettyParams(params MoqNewChecker_genType_params) string {
	return fmt.Sprintf("NewChecker_genType(%#v, %#v, %#v, %#v)", params.Conf, params.Fset, params.Pkg, params.Info)
}

func (m *MoqNewChecker_genType) ParamsKey(params MoqNewChecker_genType_params, anyParams uint64) MoqNewChecker_genType_paramsKey {
	m.Scene.T.Helper()
	var confUsed *types.Config
	var confUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Conf == moq.ParamIndexByValue {
			confUsed = params.Conf
		} else {
			confUsedHash = hash.DeepHash(params.Conf)
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
	var pkgUsed *types.Package
	var pkgUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Pkg == moq.ParamIndexByValue {
			pkgUsed = params.Pkg
		} else {
			pkgUsedHash = hash.DeepHash(params.Pkg)
		}
	}
	var infoUsed *types.Info
	var infoUsedHash hash.Hash
	if anyParams&(1<<3) == 0 {
		if m.Runtime.ParameterIndexing.Info == moq.ParamIndexByValue {
			infoUsed = params.Info
		} else {
			infoUsedHash = hash.DeepHash(params.Info)
		}
	}
	return MoqNewChecker_genType_paramsKey{
		Params: struct {
			Conf *types.Config
			Fset *token.FileSet
			Pkg  *types.Package
			Info *types.Info
		}{
			Conf: confUsed,
			Fset: fsetUsed,
			Pkg:  pkgUsed,
			Info: infoUsed,
		},
		Hashes: struct {
			Conf hash.Hash
			Fset hash.Hash
			Pkg  hash.Hash
			Info hash.Hash
		}{
			Conf: confUsedHash,
			Fset: fsetUsedHash,
			Pkg:  pkgUsedHash,
			Info: infoUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNewChecker_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNewChecker_genType) AssertExpectationsMet() {
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
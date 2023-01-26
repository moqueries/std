// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package types

import (
	"fmt"
	"go/types"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Id_genType is the fabricated implementation type of this mock (emitted when
// mocking functions directly and not from a function type)
type Id_genType func(pkg *types.Package, name string) string

// MoqId_genType holds the state of a moq of the Id_genType type
type MoqId_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqId_genType_mock

	ResultsByParams []MoqId_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Pkg  moq.ParamIndexing
			Name moq.ParamIndexing
		}
	}
}

// MoqId_genType_mock isolates the mock interface of the Id_genType type
type MoqId_genType_mock struct {
	Moq *MoqId_genType
}

// MoqId_genType_params holds the params of the Id_genType type
type MoqId_genType_params struct {
	Pkg  *types.Package
	Name string
}

// MoqId_genType_paramsKey holds the map key params of the Id_genType type
type MoqId_genType_paramsKey struct {
	Params struct {
		Pkg  *types.Package
		Name string
	}
	Hashes struct {
		Pkg  hash.Hash
		Name hash.Hash
	}
}

// MoqId_genType_resultsByParams contains the results for a given set of
// parameters for the Id_genType type
type MoqId_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqId_genType_paramsKey]*MoqId_genType_results
}

// MoqId_genType_doFn defines the type of function needed when calling AndDo
// for the Id_genType type
type MoqId_genType_doFn func(pkg *types.Package, name string)

// MoqId_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Id_genType type
type MoqId_genType_doReturnFn func(pkg *types.Package, name string) string

// MoqId_genType_results holds the results of the Id_genType type
type MoqId_genType_results struct {
	Params  MoqId_genType_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqId_genType_doFn
		DoReturnFn MoqId_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqId_genType_fnRecorder routes recorded function calls to the MoqId_genType
// moq
type MoqId_genType_fnRecorder struct {
	Params    MoqId_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqId_genType_results
	Moq       *MoqId_genType
}

// MoqId_genType_anyParams isolates the any params functions of the Id_genType
// type
type MoqId_genType_anyParams struct {
	Recorder *MoqId_genType_fnRecorder
}

// NewMoqId_genType creates a new moq of the Id_genType type
func NewMoqId_genType(scene *moq.Scene, config *moq.Config) *MoqId_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqId_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqId_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Pkg  moq.ParamIndexing
				Name moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Pkg  moq.ParamIndexing
			Name moq.ParamIndexing
		}{
			Pkg:  moq.ParamIndexByHash,
			Name: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Id_genType type
func (m *MoqId_genType) Mock() Id_genType {
	return func(pkg *types.Package, name string) string {
		m.Scene.T.Helper()
		moq := &MoqId_genType_mock{Moq: m}
		return moq.Fn(pkg, name)
	}
}

func (m *MoqId_genType_mock) Fn(pkg *types.Package, name string) (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqId_genType_params{
		Pkg:  pkg,
		Name: name,
	}
	var results *MoqId_genType_results
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
		result.DoFn(pkg, name)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(pkg, name)
	}
	return
}

func (m *MoqId_genType) OnCall(pkg *types.Package, name string) *MoqId_genType_fnRecorder {
	return &MoqId_genType_fnRecorder{
		Params: MoqId_genType_params{
			Pkg:  pkg,
			Name: name,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqId_genType_fnRecorder) Any() *MoqId_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqId_genType_anyParams{Recorder: r}
}

func (a *MoqId_genType_anyParams) Pkg() *MoqId_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqId_genType_anyParams) Name() *MoqId_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqId_genType_fnRecorder) Seq() *MoqId_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqId_genType_fnRecorder) NoSeq() *MoqId_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqId_genType_fnRecorder) ReturnResults(result1 string) *MoqId_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqId_genType_doFn
		DoReturnFn MoqId_genType_doReturnFn
	}{
		Values: &struct {
			Result1 string
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqId_genType_fnRecorder) AndDo(fn MoqId_genType_doFn) *MoqId_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqId_genType_fnRecorder) DoReturnResults(fn MoqId_genType_doReturnFn) *MoqId_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqId_genType_doFn
		DoReturnFn MoqId_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqId_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqId_genType_resultsByParams
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
		results = &MoqId_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqId_genType_paramsKey]*MoqId_genType_results{},
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
		r.Results = &MoqId_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqId_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqId_genType_fnRecorder {
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
					Result1 string
				}
				Sequence   uint32
				DoFn       MoqId_genType_doFn
				DoReturnFn MoqId_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqId_genType) PrettyParams(params MoqId_genType_params) string {
	return fmt.Sprintf("Id_genType(%#v, %#v)", params.Pkg, params.Name)
}

func (m *MoqId_genType) ParamsKey(params MoqId_genType_params, anyParams uint64) MoqId_genType_paramsKey {
	m.Scene.T.Helper()
	var pkgUsed *types.Package
	var pkgUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Pkg == moq.ParamIndexByValue {
			pkgUsed = params.Pkg
		} else {
			pkgUsedHash = hash.DeepHash(params.Pkg)
		}
	}
	var nameUsed string
	var nameUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Name == moq.ParamIndexByValue {
			nameUsed = params.Name
		} else {
			nameUsedHash = hash.DeepHash(params.Name)
		}
	}
	return MoqId_genType_paramsKey{
		Params: struct {
			Pkg  *types.Package
			Name string
		}{
			Pkg:  pkgUsed,
			Name: nameUsed,
		},
		Hashes: struct {
			Pkg  hash.Hash
			Name hash.Hash
		}{
			Pkg:  pkgUsedHash,
			Name: nameUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqId_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqId_genType) AssertExpectationsMet() {
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
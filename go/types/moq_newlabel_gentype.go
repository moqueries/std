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

// NewLabel_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type NewLabel_genType func(pos token.Pos, pkg *types.Package, name string) *types.Label

// MoqNewLabel_genType holds the state of a moq of the NewLabel_genType type
type MoqNewLabel_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNewLabel_genType_mock

	ResultsByParams []MoqNewLabel_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Pos  moq.ParamIndexing
			Pkg  moq.ParamIndexing
			Name moq.ParamIndexing
		}
	}
}

// MoqNewLabel_genType_mock isolates the mock interface of the NewLabel_genType
// type
type MoqNewLabel_genType_mock struct {
	Moq *MoqNewLabel_genType
}

// MoqNewLabel_genType_params holds the params of the NewLabel_genType type
type MoqNewLabel_genType_params struct {
	Pos  token.Pos
	Pkg  *types.Package
	Name string
}

// MoqNewLabel_genType_paramsKey holds the map key params of the
// NewLabel_genType type
type MoqNewLabel_genType_paramsKey struct {
	Params struct {
		Pos  token.Pos
		Pkg  *types.Package
		Name string
	}
	Hashes struct {
		Pos  hash.Hash
		Pkg  hash.Hash
		Name hash.Hash
	}
}

// MoqNewLabel_genType_resultsByParams contains the results for a given set of
// parameters for the NewLabel_genType type
type MoqNewLabel_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNewLabel_genType_paramsKey]*MoqNewLabel_genType_results
}

// MoqNewLabel_genType_doFn defines the type of function needed when calling
// AndDo for the NewLabel_genType type
type MoqNewLabel_genType_doFn func(pos token.Pos, pkg *types.Package, name string)

// MoqNewLabel_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the NewLabel_genType type
type MoqNewLabel_genType_doReturnFn func(pos token.Pos, pkg *types.Package, name string) *types.Label

// MoqNewLabel_genType_results holds the results of the NewLabel_genType type
type MoqNewLabel_genType_results struct {
	Params  MoqNewLabel_genType_params
	Results []struct {
		Values *struct {
			Result1 *types.Label
		}
		Sequence   uint32
		DoFn       MoqNewLabel_genType_doFn
		DoReturnFn MoqNewLabel_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNewLabel_genType_fnRecorder routes recorded function calls to the
// MoqNewLabel_genType moq
type MoqNewLabel_genType_fnRecorder struct {
	Params    MoqNewLabel_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNewLabel_genType_results
	Moq       *MoqNewLabel_genType
}

// MoqNewLabel_genType_anyParams isolates the any params functions of the
// NewLabel_genType type
type MoqNewLabel_genType_anyParams struct {
	Recorder *MoqNewLabel_genType_fnRecorder
}

// NewMoqNewLabel_genType creates a new moq of the NewLabel_genType type
func NewMoqNewLabel_genType(scene *moq.Scene, config *moq.Config) *MoqNewLabel_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNewLabel_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNewLabel_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Pos  moq.ParamIndexing
				Pkg  moq.ParamIndexing
				Name moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Pos  moq.ParamIndexing
			Pkg  moq.ParamIndexing
			Name moq.ParamIndexing
		}{
			Pos:  moq.ParamIndexByValue,
			Pkg:  moq.ParamIndexByHash,
			Name: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the NewLabel_genType type
func (m *MoqNewLabel_genType) Mock() NewLabel_genType {
	return func(pos token.Pos, pkg *types.Package, name string) *types.Label {
		m.Scene.T.Helper()
		moq := &MoqNewLabel_genType_mock{Moq: m}
		return moq.Fn(pos, pkg, name)
	}
}

func (m *MoqNewLabel_genType_mock) Fn(pos token.Pos, pkg *types.Package, name string) (result1 *types.Label) {
	m.Moq.Scene.T.Helper()
	params := MoqNewLabel_genType_params{
		Pos:  pos,
		Pkg:  pkg,
		Name: name,
	}
	var results *MoqNewLabel_genType_results
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
		result.DoFn(pos, pkg, name)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(pos, pkg, name)
	}
	return
}

func (m *MoqNewLabel_genType) OnCall(pos token.Pos, pkg *types.Package, name string) *MoqNewLabel_genType_fnRecorder {
	return &MoqNewLabel_genType_fnRecorder{
		Params: MoqNewLabel_genType_params{
			Pos:  pos,
			Pkg:  pkg,
			Name: name,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNewLabel_genType_fnRecorder) Any() *MoqNewLabel_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNewLabel_genType_anyParams{Recorder: r}
}

func (a *MoqNewLabel_genType_anyParams) Pos() *MoqNewLabel_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqNewLabel_genType_anyParams) Pkg() *MoqNewLabel_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqNewLabel_genType_anyParams) Name() *MoqNewLabel_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqNewLabel_genType_fnRecorder) Seq() *MoqNewLabel_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNewLabel_genType_fnRecorder) NoSeq() *MoqNewLabel_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNewLabel_genType_fnRecorder) ReturnResults(result1 *types.Label) *MoqNewLabel_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *types.Label
		}
		Sequence   uint32
		DoFn       MoqNewLabel_genType_doFn
		DoReturnFn MoqNewLabel_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *types.Label
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNewLabel_genType_fnRecorder) AndDo(fn MoqNewLabel_genType_doFn) *MoqNewLabel_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNewLabel_genType_fnRecorder) DoReturnResults(fn MoqNewLabel_genType_doReturnFn) *MoqNewLabel_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *types.Label
		}
		Sequence   uint32
		DoFn       MoqNewLabel_genType_doFn
		DoReturnFn MoqNewLabel_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNewLabel_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNewLabel_genType_resultsByParams
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
		results = &MoqNewLabel_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNewLabel_genType_paramsKey]*MoqNewLabel_genType_results{},
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
		r.Results = &MoqNewLabel_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNewLabel_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNewLabel_genType_fnRecorder {
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
					Result1 *types.Label
				}
				Sequence   uint32
				DoFn       MoqNewLabel_genType_doFn
				DoReturnFn MoqNewLabel_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNewLabel_genType) PrettyParams(params MoqNewLabel_genType_params) string {
	return fmt.Sprintf("NewLabel_genType(%#v, %#v, %#v)", params.Pos, params.Pkg, params.Name)
}

func (m *MoqNewLabel_genType) ParamsKey(params MoqNewLabel_genType_params, anyParams uint64) MoqNewLabel_genType_paramsKey {
	m.Scene.T.Helper()
	var posUsed token.Pos
	var posUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Pos == moq.ParamIndexByValue {
			posUsed = params.Pos
		} else {
			posUsedHash = hash.DeepHash(params.Pos)
		}
	}
	var pkgUsed *types.Package
	var pkgUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Pkg == moq.ParamIndexByValue {
			pkgUsed = params.Pkg
		} else {
			pkgUsedHash = hash.DeepHash(params.Pkg)
		}
	}
	var nameUsed string
	var nameUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Name == moq.ParamIndexByValue {
			nameUsed = params.Name
		} else {
			nameUsedHash = hash.DeepHash(params.Name)
		}
	}
	return MoqNewLabel_genType_paramsKey{
		Params: struct {
			Pos  token.Pos
			Pkg  *types.Package
			Name string
		}{
			Pos:  posUsed,
			Pkg:  pkgUsed,
			Name: nameUsed,
		},
		Hashes: struct {
			Pos  hash.Hash
			Pkg  hash.Hash
			Name hash.Hash
		}{
			Pos:  posUsedHash,
			Pkg:  pkgUsedHash,
			Name: nameUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNewLabel_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNewLabel_genType) AssertExpectationsMet() {
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

// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package ast

import (
	"fmt"
	"go/ast"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Walk_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Walk_genType func(v ast.Visitor, node ast.Node)

// MoqWalk_genType holds the state of a moq of the Walk_genType type
type MoqWalk_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqWalk_genType_mock

	ResultsByParams []MoqWalk_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			V    moq.ParamIndexing
			Node moq.ParamIndexing
		}
	}
}

// MoqWalk_genType_mock isolates the mock interface of the Walk_genType type
type MoqWalk_genType_mock struct {
	Moq *MoqWalk_genType
}

// MoqWalk_genType_params holds the params of the Walk_genType type
type MoqWalk_genType_params struct {
	V    ast.Visitor
	Node ast.Node
}

// MoqWalk_genType_paramsKey holds the map key params of the Walk_genType type
type MoqWalk_genType_paramsKey struct {
	Params struct {
		V    ast.Visitor
		Node ast.Node
	}
	Hashes struct {
		V    hash.Hash
		Node hash.Hash
	}
}

// MoqWalk_genType_resultsByParams contains the results for a given set of
// parameters for the Walk_genType type
type MoqWalk_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqWalk_genType_paramsKey]*MoqWalk_genType_results
}

// MoqWalk_genType_doFn defines the type of function needed when calling AndDo
// for the Walk_genType type
type MoqWalk_genType_doFn func(v ast.Visitor, node ast.Node)

// MoqWalk_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Walk_genType type
type MoqWalk_genType_doReturnFn func(v ast.Visitor, node ast.Node)

// MoqWalk_genType_results holds the results of the Walk_genType type
type MoqWalk_genType_results struct {
	Params  MoqWalk_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqWalk_genType_doFn
		DoReturnFn MoqWalk_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqWalk_genType_fnRecorder routes recorded function calls to the
// MoqWalk_genType moq
type MoqWalk_genType_fnRecorder struct {
	Params    MoqWalk_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqWalk_genType_results
	Moq       *MoqWalk_genType
}

// MoqWalk_genType_anyParams isolates the any params functions of the
// Walk_genType type
type MoqWalk_genType_anyParams struct {
	Recorder *MoqWalk_genType_fnRecorder
}

// NewMoqWalk_genType creates a new moq of the Walk_genType type
func NewMoqWalk_genType(scene *moq.Scene, config *moq.Config) *MoqWalk_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqWalk_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqWalk_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				V    moq.ParamIndexing
				Node moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			V    moq.ParamIndexing
			Node moq.ParamIndexing
		}{
			V:    moq.ParamIndexByHash,
			Node: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Walk_genType type
func (m *MoqWalk_genType) Mock() Walk_genType {
	return func(v ast.Visitor, node ast.Node) {
		m.Scene.T.Helper()
		moq := &MoqWalk_genType_mock{Moq: m}
		moq.Fn(v, node)
	}
}

func (m *MoqWalk_genType_mock) Fn(v ast.Visitor, node ast.Node) {
	m.Moq.Scene.T.Helper()
	params := MoqWalk_genType_params{
		V:    v,
		Node: node,
	}
	var results *MoqWalk_genType_results
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
		result.DoFn(v, node)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(v, node)
	}
	return
}

func (m *MoqWalk_genType) OnCall(v ast.Visitor, node ast.Node) *MoqWalk_genType_fnRecorder {
	return &MoqWalk_genType_fnRecorder{
		Params: MoqWalk_genType_params{
			V:    v,
			Node: node,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqWalk_genType_fnRecorder) Any() *MoqWalk_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqWalk_genType_anyParams{Recorder: r}
}

func (a *MoqWalk_genType_anyParams) V() *MoqWalk_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqWalk_genType_anyParams) Node() *MoqWalk_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqWalk_genType_fnRecorder) Seq() *MoqWalk_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqWalk_genType_fnRecorder) NoSeq() *MoqWalk_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqWalk_genType_fnRecorder) ReturnResults() *MoqWalk_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqWalk_genType_doFn
		DoReturnFn MoqWalk_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqWalk_genType_fnRecorder) AndDo(fn MoqWalk_genType_doFn) *MoqWalk_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqWalk_genType_fnRecorder) DoReturnResults(fn MoqWalk_genType_doReturnFn) *MoqWalk_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqWalk_genType_doFn
		DoReturnFn MoqWalk_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqWalk_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqWalk_genType_resultsByParams
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
		results = &MoqWalk_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqWalk_genType_paramsKey]*MoqWalk_genType_results{},
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
		r.Results = &MoqWalk_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqWalk_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqWalk_genType_fnRecorder {
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
				DoFn       MoqWalk_genType_doFn
				DoReturnFn MoqWalk_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqWalk_genType) PrettyParams(params MoqWalk_genType_params) string {
	return fmt.Sprintf("Walk_genType(%#v, %#v)", params.V, params.Node)
}

func (m *MoqWalk_genType) ParamsKey(params MoqWalk_genType_params, anyParams uint64) MoqWalk_genType_paramsKey {
	m.Scene.T.Helper()
	var vUsed ast.Visitor
	var vUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.V == moq.ParamIndexByValue {
			vUsed = params.V
		} else {
			vUsedHash = hash.DeepHash(params.V)
		}
	}
	var nodeUsed ast.Node
	var nodeUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Node == moq.ParamIndexByValue {
			nodeUsed = params.Node
		} else {
			nodeUsedHash = hash.DeepHash(params.Node)
		}
	}
	return MoqWalk_genType_paramsKey{
		Params: struct {
			V    ast.Visitor
			Node ast.Node
		}{
			V:    vUsed,
			Node: nodeUsed,
		},
		Hashes: struct {
			V    hash.Hash
			Node hash.Hash
		}{
			V:    vUsedHash,
			Node: nodeUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqWalk_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqWalk_genType) AssertExpectationsMet() {
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

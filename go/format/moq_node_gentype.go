// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package format

import (
	"fmt"
	"go/token"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Node_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Node_genType func(dst io.Writer, fset *token.FileSet, node any) error

// MoqNode_genType holds the state of a moq of the Node_genType type
type MoqNode_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNode_genType_mock

	ResultsByParams []MoqNode_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Dst  moq.ParamIndexing
			Fset moq.ParamIndexing
			Node moq.ParamIndexing
		}
	}
}

// MoqNode_genType_mock isolates the mock interface of the Node_genType type
type MoqNode_genType_mock struct {
	Moq *MoqNode_genType
}

// MoqNode_genType_params holds the params of the Node_genType type
type MoqNode_genType_params struct {
	Dst  io.Writer
	Fset *token.FileSet
	Node any
}

// MoqNode_genType_paramsKey holds the map key params of the Node_genType type
type MoqNode_genType_paramsKey struct {
	Params struct {
		Dst  io.Writer
		Fset *token.FileSet
		Node any
	}
	Hashes struct {
		Dst  hash.Hash
		Fset hash.Hash
		Node hash.Hash
	}
}

// MoqNode_genType_resultsByParams contains the results for a given set of
// parameters for the Node_genType type
type MoqNode_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNode_genType_paramsKey]*MoqNode_genType_results
}

// MoqNode_genType_doFn defines the type of function needed when calling AndDo
// for the Node_genType type
type MoqNode_genType_doFn func(dst io.Writer, fset *token.FileSet, node any)

// MoqNode_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Node_genType type
type MoqNode_genType_doReturnFn func(dst io.Writer, fset *token.FileSet, node any) error

// MoqNode_genType_results holds the results of the Node_genType type
type MoqNode_genType_results struct {
	Params  MoqNode_genType_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqNode_genType_doFn
		DoReturnFn MoqNode_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNode_genType_fnRecorder routes recorded function calls to the
// MoqNode_genType moq
type MoqNode_genType_fnRecorder struct {
	Params    MoqNode_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNode_genType_results
	Moq       *MoqNode_genType
}

// MoqNode_genType_anyParams isolates the any params functions of the
// Node_genType type
type MoqNode_genType_anyParams struct {
	Recorder *MoqNode_genType_fnRecorder
}

// NewMoqNode_genType creates a new moq of the Node_genType type
func NewMoqNode_genType(scene *moq.Scene, config *moq.Config) *MoqNode_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNode_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNode_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Dst  moq.ParamIndexing
				Fset moq.ParamIndexing
				Node moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Dst  moq.ParamIndexing
			Fset moq.ParamIndexing
			Node moq.ParamIndexing
		}{
			Dst:  moq.ParamIndexByHash,
			Fset: moq.ParamIndexByHash,
			Node: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Node_genType type
func (m *MoqNode_genType) Mock() Node_genType {
	return func(dst io.Writer, fset *token.FileSet, node any) error {
		m.Scene.T.Helper()
		moq := &MoqNode_genType_mock{Moq: m}
		return moq.Fn(dst, fset, node)
	}
}

func (m *MoqNode_genType_mock) Fn(dst io.Writer, fset *token.FileSet, node any) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqNode_genType_params{
		Dst:  dst,
		Fset: fset,
		Node: node,
	}
	var results *MoqNode_genType_results
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
		result.DoFn(dst, fset, node)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(dst, fset, node)
	}
	return
}

func (m *MoqNode_genType) OnCall(dst io.Writer, fset *token.FileSet, node any) *MoqNode_genType_fnRecorder {
	return &MoqNode_genType_fnRecorder{
		Params: MoqNode_genType_params{
			Dst:  dst,
			Fset: fset,
			Node: node,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNode_genType_fnRecorder) Any() *MoqNode_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNode_genType_anyParams{Recorder: r}
}

func (a *MoqNode_genType_anyParams) Dst() *MoqNode_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqNode_genType_anyParams) Fset() *MoqNode_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqNode_genType_anyParams) Node() *MoqNode_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqNode_genType_fnRecorder) Seq() *MoqNode_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNode_genType_fnRecorder) NoSeq() *MoqNode_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNode_genType_fnRecorder) ReturnResults(result1 error) *MoqNode_genType_fnRecorder {
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
		DoFn       MoqNode_genType_doFn
		DoReturnFn MoqNode_genType_doReturnFn
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

func (r *MoqNode_genType_fnRecorder) AndDo(fn MoqNode_genType_doFn) *MoqNode_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNode_genType_fnRecorder) DoReturnResults(fn MoqNode_genType_doReturnFn) *MoqNode_genType_fnRecorder {
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
		DoFn       MoqNode_genType_doFn
		DoReturnFn MoqNode_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNode_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNode_genType_resultsByParams
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
		results = &MoqNode_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNode_genType_paramsKey]*MoqNode_genType_results{},
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
		r.Results = &MoqNode_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNode_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNode_genType_fnRecorder {
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
				DoFn       MoqNode_genType_doFn
				DoReturnFn MoqNode_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNode_genType) PrettyParams(params MoqNode_genType_params) string {
	return fmt.Sprintf("Node_genType(%#v, %#v, %#v)", params.Dst, params.Fset, params.Node)
}

func (m *MoqNode_genType) ParamsKey(params MoqNode_genType_params, anyParams uint64) MoqNode_genType_paramsKey {
	m.Scene.T.Helper()
	var dstUsed io.Writer
	var dstUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Dst == moq.ParamIndexByValue {
			dstUsed = params.Dst
		} else {
			dstUsedHash = hash.DeepHash(params.Dst)
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
	var nodeUsed any
	var nodeUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Node == moq.ParamIndexByValue {
			nodeUsed = params.Node
		} else {
			nodeUsedHash = hash.DeepHash(params.Node)
		}
	}
	return MoqNode_genType_paramsKey{
		Params: struct {
			Dst  io.Writer
			Fset *token.FileSet
			Node any
		}{
			Dst:  dstUsed,
			Fset: fsetUsed,
			Node: nodeUsed,
		},
		Hashes: struct {
			Dst  hash.Hash
			Fset hash.Hash
			Node hash.Hash
		}{
			Dst:  dstUsedHash,
			Fset: fsetUsedHash,
			Node: nodeUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNode_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNode_genType) AssertExpectationsMet() {
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

// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package ast

import (
	"fmt"
	"go/ast"
	"go/token"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that ast.Node is mocked completely
var _ ast.Node = (*MoqNode_mock)(nil)

// MoqNode holds the state of a moq of the Node type
type MoqNode struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNode_mock

	ResultsByParams_Pos []MoqNode_Pos_resultsByParams
	ResultsByParams_End []MoqNode_End_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Pos struct{}
			End struct{}
		}
	}
}

// MoqNode_mock isolates the mock interface of the Node type
type MoqNode_mock struct {
	Moq *MoqNode
}

// MoqNode_recorder isolates the recorder interface of the Node type
type MoqNode_recorder struct {
	Moq *MoqNode
}

// MoqNode_Pos_params holds the params of the Node type
type MoqNode_Pos_params struct{}

// MoqNode_Pos_paramsKey holds the map key params of the Node type
type MoqNode_Pos_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqNode_Pos_resultsByParams contains the results for a given set of
// parameters for the Node type
type MoqNode_Pos_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNode_Pos_paramsKey]*MoqNode_Pos_results
}

// MoqNode_Pos_doFn defines the type of function needed when calling AndDo for
// the Node type
type MoqNode_Pos_doFn func()

// MoqNode_Pos_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Node type
type MoqNode_Pos_doReturnFn func() token.Pos

// MoqNode_Pos_results holds the results of the Node type
type MoqNode_Pos_results struct {
	Params  MoqNode_Pos_params
	Results []struct {
		Values *struct {
			Result1 token.Pos
		}
		Sequence   uint32
		DoFn       MoqNode_Pos_doFn
		DoReturnFn MoqNode_Pos_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNode_Pos_fnRecorder routes recorded function calls to the MoqNode moq
type MoqNode_Pos_fnRecorder struct {
	Params    MoqNode_Pos_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNode_Pos_results
	Moq       *MoqNode
}

// MoqNode_Pos_anyParams isolates the any params functions of the Node type
type MoqNode_Pos_anyParams struct {
	Recorder *MoqNode_Pos_fnRecorder
}

// MoqNode_End_params holds the params of the Node type
type MoqNode_End_params struct{}

// MoqNode_End_paramsKey holds the map key params of the Node type
type MoqNode_End_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqNode_End_resultsByParams contains the results for a given set of
// parameters for the Node type
type MoqNode_End_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNode_End_paramsKey]*MoqNode_End_results
}

// MoqNode_End_doFn defines the type of function needed when calling AndDo for
// the Node type
type MoqNode_End_doFn func()

// MoqNode_End_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Node type
type MoqNode_End_doReturnFn func() token.Pos

// MoqNode_End_results holds the results of the Node type
type MoqNode_End_results struct {
	Params  MoqNode_End_params
	Results []struct {
		Values *struct {
			Result1 token.Pos
		}
		Sequence   uint32
		DoFn       MoqNode_End_doFn
		DoReturnFn MoqNode_End_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNode_End_fnRecorder routes recorded function calls to the MoqNode moq
type MoqNode_End_fnRecorder struct {
	Params    MoqNode_End_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNode_End_results
	Moq       *MoqNode
}

// MoqNode_End_anyParams isolates the any params functions of the Node type
type MoqNode_End_anyParams struct {
	Recorder *MoqNode_End_fnRecorder
}

// NewMoqNode creates a new moq of the Node type
func NewMoqNode(scene *moq.Scene, config *moq.Config) *MoqNode {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNode{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNode_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Pos struct{}
				End struct{}
			}
		}{ParameterIndexing: struct {
			Pos struct{}
			End struct{}
		}{
			Pos: struct{}{},
			End: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Node type
func (m *MoqNode) Mock() *MoqNode_mock { return m.Moq }

func (m *MoqNode_mock) Pos() (result1 token.Pos) {
	m.Moq.Scene.T.Helper()
	params := MoqNode_Pos_params{}
	var results *MoqNode_Pos_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Pos {
		paramsKey := m.Moq.ParamsKey_Pos(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Pos(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Pos(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Pos(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn()
	}
	return
}

func (m *MoqNode_mock) End() (result1 token.Pos) {
	m.Moq.Scene.T.Helper()
	params := MoqNode_End_params{}
	var results *MoqNode_End_results
	for _, resultsByParams := range m.Moq.ResultsByParams_End {
		paramsKey := m.Moq.ParamsKey_End(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_End(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_End(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_End(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn()
	}
	return
}

// OnCall returns the recorder implementation of the Node type
func (m *MoqNode) OnCall() *MoqNode_recorder {
	return &MoqNode_recorder{
		Moq: m,
	}
}

func (m *MoqNode_recorder) Pos() *MoqNode_Pos_fnRecorder {
	return &MoqNode_Pos_fnRecorder{
		Params:   MoqNode_Pos_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqNode_Pos_fnRecorder) Any() *MoqNode_Pos_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Pos(r.Params))
		return nil
	}
	return &MoqNode_Pos_anyParams{Recorder: r}
}

func (r *MoqNode_Pos_fnRecorder) Seq() *MoqNode_Pos_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Pos(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNode_Pos_fnRecorder) NoSeq() *MoqNode_Pos_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Pos(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNode_Pos_fnRecorder) ReturnResults(result1 token.Pos) *MoqNode_Pos_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 token.Pos
		}
		Sequence   uint32
		DoFn       MoqNode_Pos_doFn
		DoReturnFn MoqNode_Pos_doReturnFn
	}{
		Values: &struct {
			Result1 token.Pos
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNode_Pos_fnRecorder) AndDo(fn MoqNode_Pos_doFn) *MoqNode_Pos_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNode_Pos_fnRecorder) DoReturnResults(fn MoqNode_Pos_doReturnFn) *MoqNode_Pos_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 token.Pos
		}
		Sequence   uint32
		DoFn       MoqNode_Pos_doFn
		DoReturnFn MoqNode_Pos_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNode_Pos_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNode_Pos_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Pos {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqNode_Pos_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNode_Pos_paramsKey]*MoqNode_Pos_results{},
		}
		r.Moq.ResultsByParams_Pos = append(r.Moq.ResultsByParams_Pos, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Pos) {
			copy(r.Moq.ResultsByParams_Pos[insertAt+1:], r.Moq.ResultsByParams_Pos[insertAt:0])
			r.Moq.ResultsByParams_Pos[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Pos(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqNode_Pos_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNode_Pos_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNode_Pos_fnRecorder {
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
					Result1 token.Pos
				}
				Sequence   uint32
				DoFn       MoqNode_Pos_doFn
				DoReturnFn MoqNode_Pos_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNode) PrettyParams_Pos(params MoqNode_Pos_params) string { return fmt.Sprintf("Pos()") }

func (m *MoqNode) ParamsKey_Pos(params MoqNode_Pos_params, anyParams uint64) MoqNode_Pos_paramsKey {
	m.Scene.T.Helper()
	return MoqNode_Pos_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqNode_recorder) End() *MoqNode_End_fnRecorder {
	return &MoqNode_End_fnRecorder{
		Params:   MoqNode_End_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqNode_End_fnRecorder) Any() *MoqNode_End_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_End(r.Params))
		return nil
	}
	return &MoqNode_End_anyParams{Recorder: r}
}

func (r *MoqNode_End_fnRecorder) Seq() *MoqNode_End_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_End(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNode_End_fnRecorder) NoSeq() *MoqNode_End_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_End(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNode_End_fnRecorder) ReturnResults(result1 token.Pos) *MoqNode_End_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 token.Pos
		}
		Sequence   uint32
		DoFn       MoqNode_End_doFn
		DoReturnFn MoqNode_End_doReturnFn
	}{
		Values: &struct {
			Result1 token.Pos
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNode_End_fnRecorder) AndDo(fn MoqNode_End_doFn) *MoqNode_End_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNode_End_fnRecorder) DoReturnResults(fn MoqNode_End_doReturnFn) *MoqNode_End_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 token.Pos
		}
		Sequence   uint32
		DoFn       MoqNode_End_doFn
		DoReturnFn MoqNode_End_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNode_End_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNode_End_resultsByParams
	for n, res := range r.Moq.ResultsByParams_End {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqNode_End_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNode_End_paramsKey]*MoqNode_End_results{},
		}
		r.Moq.ResultsByParams_End = append(r.Moq.ResultsByParams_End, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_End) {
			copy(r.Moq.ResultsByParams_End[insertAt+1:], r.Moq.ResultsByParams_End[insertAt:0])
			r.Moq.ResultsByParams_End[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_End(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqNode_End_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNode_End_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNode_End_fnRecorder {
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
					Result1 token.Pos
				}
				Sequence   uint32
				DoFn       MoqNode_End_doFn
				DoReturnFn MoqNode_End_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNode) PrettyParams_End(params MoqNode_End_params) string { return fmt.Sprintf("End()") }

func (m *MoqNode) ParamsKey_End(params MoqNode_End_params, anyParams uint64) MoqNode_End_paramsKey {
	m.Scene.T.Helper()
	return MoqNode_End_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqNode) Reset() { m.ResultsByParams_Pos = nil; m.ResultsByParams_End = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNode) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Pos {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Pos(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_End {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_End(results.Params))
			}
		}
	}
}
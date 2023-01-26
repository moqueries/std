// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package ast

import (
	"fmt"
	"go/token"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that ast.BadExpr_starGenType is mocked
// completely
var _ BadExpr_starGenType = (*MoqBadExpr_starGenType_mock)(nil)

// BadExpr_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type BadExpr_starGenType interface {
	Pos() token.Pos
	End() token.Pos
}

// MoqBadExpr_starGenType holds the state of a moq of the BadExpr_starGenType
// type
type MoqBadExpr_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqBadExpr_starGenType_mock

	ResultsByParams_Pos []MoqBadExpr_starGenType_Pos_resultsByParams
	ResultsByParams_End []MoqBadExpr_starGenType_End_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Pos struct{}
			End struct{}
		}
	}
}

// MoqBadExpr_starGenType_mock isolates the mock interface of the
// BadExpr_starGenType type
type MoqBadExpr_starGenType_mock struct {
	Moq *MoqBadExpr_starGenType
}

// MoqBadExpr_starGenType_recorder isolates the recorder interface of the
// BadExpr_starGenType type
type MoqBadExpr_starGenType_recorder struct {
	Moq *MoqBadExpr_starGenType
}

// MoqBadExpr_starGenType_Pos_params holds the params of the
// BadExpr_starGenType type
type MoqBadExpr_starGenType_Pos_params struct{}

// MoqBadExpr_starGenType_Pos_paramsKey holds the map key params of the
// BadExpr_starGenType type
type MoqBadExpr_starGenType_Pos_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqBadExpr_starGenType_Pos_resultsByParams contains the results for a given
// set of parameters for the BadExpr_starGenType type
type MoqBadExpr_starGenType_Pos_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqBadExpr_starGenType_Pos_paramsKey]*MoqBadExpr_starGenType_Pos_results
}

// MoqBadExpr_starGenType_Pos_doFn defines the type of function needed when
// calling AndDo for the BadExpr_starGenType type
type MoqBadExpr_starGenType_Pos_doFn func()

// MoqBadExpr_starGenType_Pos_doReturnFn defines the type of function needed
// when calling DoReturnResults for the BadExpr_starGenType type
type MoqBadExpr_starGenType_Pos_doReturnFn func() token.Pos

// MoqBadExpr_starGenType_Pos_results holds the results of the
// BadExpr_starGenType type
type MoqBadExpr_starGenType_Pos_results struct {
	Params  MoqBadExpr_starGenType_Pos_params
	Results []struct {
		Values *struct {
			Result1 token.Pos
		}
		Sequence   uint32
		DoFn       MoqBadExpr_starGenType_Pos_doFn
		DoReturnFn MoqBadExpr_starGenType_Pos_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqBadExpr_starGenType_Pos_fnRecorder routes recorded function calls to the
// MoqBadExpr_starGenType moq
type MoqBadExpr_starGenType_Pos_fnRecorder struct {
	Params    MoqBadExpr_starGenType_Pos_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqBadExpr_starGenType_Pos_results
	Moq       *MoqBadExpr_starGenType
}

// MoqBadExpr_starGenType_Pos_anyParams isolates the any params functions of
// the BadExpr_starGenType type
type MoqBadExpr_starGenType_Pos_anyParams struct {
	Recorder *MoqBadExpr_starGenType_Pos_fnRecorder
}

// MoqBadExpr_starGenType_End_params holds the params of the
// BadExpr_starGenType type
type MoqBadExpr_starGenType_End_params struct{}

// MoqBadExpr_starGenType_End_paramsKey holds the map key params of the
// BadExpr_starGenType type
type MoqBadExpr_starGenType_End_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqBadExpr_starGenType_End_resultsByParams contains the results for a given
// set of parameters for the BadExpr_starGenType type
type MoqBadExpr_starGenType_End_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqBadExpr_starGenType_End_paramsKey]*MoqBadExpr_starGenType_End_results
}

// MoqBadExpr_starGenType_End_doFn defines the type of function needed when
// calling AndDo for the BadExpr_starGenType type
type MoqBadExpr_starGenType_End_doFn func()

// MoqBadExpr_starGenType_End_doReturnFn defines the type of function needed
// when calling DoReturnResults for the BadExpr_starGenType type
type MoqBadExpr_starGenType_End_doReturnFn func() token.Pos

// MoqBadExpr_starGenType_End_results holds the results of the
// BadExpr_starGenType type
type MoqBadExpr_starGenType_End_results struct {
	Params  MoqBadExpr_starGenType_End_params
	Results []struct {
		Values *struct {
			Result1 token.Pos
		}
		Sequence   uint32
		DoFn       MoqBadExpr_starGenType_End_doFn
		DoReturnFn MoqBadExpr_starGenType_End_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqBadExpr_starGenType_End_fnRecorder routes recorded function calls to the
// MoqBadExpr_starGenType moq
type MoqBadExpr_starGenType_End_fnRecorder struct {
	Params    MoqBadExpr_starGenType_End_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqBadExpr_starGenType_End_results
	Moq       *MoqBadExpr_starGenType
}

// MoqBadExpr_starGenType_End_anyParams isolates the any params functions of
// the BadExpr_starGenType type
type MoqBadExpr_starGenType_End_anyParams struct {
	Recorder *MoqBadExpr_starGenType_End_fnRecorder
}

// NewMoqBadExpr_starGenType creates a new moq of the BadExpr_starGenType type
func NewMoqBadExpr_starGenType(scene *moq.Scene, config *moq.Config) *MoqBadExpr_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqBadExpr_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqBadExpr_starGenType_mock{},

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

// Mock returns the mock implementation of the BadExpr_starGenType type
func (m *MoqBadExpr_starGenType) Mock() *MoqBadExpr_starGenType_mock { return m.Moq }

func (m *MoqBadExpr_starGenType_mock) Pos() (result1 token.Pos) {
	m.Moq.Scene.T.Helper()
	params := MoqBadExpr_starGenType_Pos_params{}
	var results *MoqBadExpr_starGenType_Pos_results
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

func (m *MoqBadExpr_starGenType_mock) End() (result1 token.Pos) {
	m.Moq.Scene.T.Helper()
	params := MoqBadExpr_starGenType_End_params{}
	var results *MoqBadExpr_starGenType_End_results
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

// OnCall returns the recorder implementation of the BadExpr_starGenType type
func (m *MoqBadExpr_starGenType) OnCall() *MoqBadExpr_starGenType_recorder {
	return &MoqBadExpr_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqBadExpr_starGenType_recorder) Pos() *MoqBadExpr_starGenType_Pos_fnRecorder {
	return &MoqBadExpr_starGenType_Pos_fnRecorder{
		Params:   MoqBadExpr_starGenType_Pos_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqBadExpr_starGenType_Pos_fnRecorder) Any() *MoqBadExpr_starGenType_Pos_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Pos(r.Params))
		return nil
	}
	return &MoqBadExpr_starGenType_Pos_anyParams{Recorder: r}
}

func (r *MoqBadExpr_starGenType_Pos_fnRecorder) Seq() *MoqBadExpr_starGenType_Pos_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Pos(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqBadExpr_starGenType_Pos_fnRecorder) NoSeq() *MoqBadExpr_starGenType_Pos_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Pos(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqBadExpr_starGenType_Pos_fnRecorder) ReturnResults(result1 token.Pos) *MoqBadExpr_starGenType_Pos_fnRecorder {
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
		DoFn       MoqBadExpr_starGenType_Pos_doFn
		DoReturnFn MoqBadExpr_starGenType_Pos_doReturnFn
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

func (r *MoqBadExpr_starGenType_Pos_fnRecorder) AndDo(fn MoqBadExpr_starGenType_Pos_doFn) *MoqBadExpr_starGenType_Pos_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqBadExpr_starGenType_Pos_fnRecorder) DoReturnResults(fn MoqBadExpr_starGenType_Pos_doReturnFn) *MoqBadExpr_starGenType_Pos_fnRecorder {
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
		DoFn       MoqBadExpr_starGenType_Pos_doFn
		DoReturnFn MoqBadExpr_starGenType_Pos_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqBadExpr_starGenType_Pos_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqBadExpr_starGenType_Pos_resultsByParams
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
		results = &MoqBadExpr_starGenType_Pos_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqBadExpr_starGenType_Pos_paramsKey]*MoqBadExpr_starGenType_Pos_results{},
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
		r.Results = &MoqBadExpr_starGenType_Pos_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqBadExpr_starGenType_Pos_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqBadExpr_starGenType_Pos_fnRecorder {
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
				DoFn       MoqBadExpr_starGenType_Pos_doFn
				DoReturnFn MoqBadExpr_starGenType_Pos_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqBadExpr_starGenType) PrettyParams_Pos(params MoqBadExpr_starGenType_Pos_params) string {
	return fmt.Sprintf("Pos()")
}

func (m *MoqBadExpr_starGenType) ParamsKey_Pos(params MoqBadExpr_starGenType_Pos_params, anyParams uint64) MoqBadExpr_starGenType_Pos_paramsKey {
	m.Scene.T.Helper()
	return MoqBadExpr_starGenType_Pos_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqBadExpr_starGenType_recorder) End() *MoqBadExpr_starGenType_End_fnRecorder {
	return &MoqBadExpr_starGenType_End_fnRecorder{
		Params:   MoqBadExpr_starGenType_End_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqBadExpr_starGenType_End_fnRecorder) Any() *MoqBadExpr_starGenType_End_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_End(r.Params))
		return nil
	}
	return &MoqBadExpr_starGenType_End_anyParams{Recorder: r}
}

func (r *MoqBadExpr_starGenType_End_fnRecorder) Seq() *MoqBadExpr_starGenType_End_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_End(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqBadExpr_starGenType_End_fnRecorder) NoSeq() *MoqBadExpr_starGenType_End_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_End(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqBadExpr_starGenType_End_fnRecorder) ReturnResults(result1 token.Pos) *MoqBadExpr_starGenType_End_fnRecorder {
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
		DoFn       MoqBadExpr_starGenType_End_doFn
		DoReturnFn MoqBadExpr_starGenType_End_doReturnFn
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

func (r *MoqBadExpr_starGenType_End_fnRecorder) AndDo(fn MoqBadExpr_starGenType_End_doFn) *MoqBadExpr_starGenType_End_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqBadExpr_starGenType_End_fnRecorder) DoReturnResults(fn MoqBadExpr_starGenType_End_doReturnFn) *MoqBadExpr_starGenType_End_fnRecorder {
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
		DoFn       MoqBadExpr_starGenType_End_doFn
		DoReturnFn MoqBadExpr_starGenType_End_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqBadExpr_starGenType_End_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqBadExpr_starGenType_End_resultsByParams
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
		results = &MoqBadExpr_starGenType_End_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqBadExpr_starGenType_End_paramsKey]*MoqBadExpr_starGenType_End_results{},
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
		r.Results = &MoqBadExpr_starGenType_End_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqBadExpr_starGenType_End_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqBadExpr_starGenType_End_fnRecorder {
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
				DoFn       MoqBadExpr_starGenType_End_doFn
				DoReturnFn MoqBadExpr_starGenType_End_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqBadExpr_starGenType) PrettyParams_End(params MoqBadExpr_starGenType_End_params) string {
	return fmt.Sprintf("End()")
}

func (m *MoqBadExpr_starGenType) ParamsKey_End(params MoqBadExpr_starGenType_End_params, anyParams uint64) MoqBadExpr_starGenType_End_paramsKey {
	m.Scene.T.Helper()
	return MoqBadExpr_starGenType_End_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqBadExpr_starGenType) Reset() { m.ResultsByParams_Pos = nil; m.ResultsByParams_End = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqBadExpr_starGenType) AssertExpectationsMet() {
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

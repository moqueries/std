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

// The following type assertion assures that types.Union_starGenType is mocked
// completely
var _ Union_starGenType = (*MoqUnion_starGenType_mock)(nil)

// Union_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type Union_starGenType interface {
	Len() int
	Term(i int) *types.Term
	Underlying() types.Type
	String() string
}

// MoqUnion_starGenType holds the state of a moq of the Union_starGenType type
type MoqUnion_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqUnion_starGenType_mock

	ResultsByParams_Len        []MoqUnion_starGenType_Len_resultsByParams
	ResultsByParams_Term       []MoqUnion_starGenType_Term_resultsByParams
	ResultsByParams_Underlying []MoqUnion_starGenType_Underlying_resultsByParams
	ResultsByParams_String     []MoqUnion_starGenType_String_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Len  struct{}
			Term struct {
				Param1 moq.ParamIndexing
			}
			Underlying struct{}
			String     struct{}
		}
	}
}

// MoqUnion_starGenType_mock isolates the mock interface of the
// Union_starGenType type
type MoqUnion_starGenType_mock struct {
	Moq *MoqUnion_starGenType
}

// MoqUnion_starGenType_recorder isolates the recorder interface of the
// Union_starGenType type
type MoqUnion_starGenType_recorder struct {
	Moq *MoqUnion_starGenType
}

// MoqUnion_starGenType_Len_params holds the params of the Union_starGenType
// type
type MoqUnion_starGenType_Len_params struct{}

// MoqUnion_starGenType_Len_paramsKey holds the map key params of the
// Union_starGenType type
type MoqUnion_starGenType_Len_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqUnion_starGenType_Len_resultsByParams contains the results for a given
// set of parameters for the Union_starGenType type
type MoqUnion_starGenType_Len_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqUnion_starGenType_Len_paramsKey]*MoqUnion_starGenType_Len_results
}

// MoqUnion_starGenType_Len_doFn defines the type of function needed when
// calling AndDo for the Union_starGenType type
type MoqUnion_starGenType_Len_doFn func()

// MoqUnion_starGenType_Len_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Union_starGenType type
type MoqUnion_starGenType_Len_doReturnFn func() int

// MoqUnion_starGenType_Len_results holds the results of the Union_starGenType
// type
type MoqUnion_starGenType_Len_results struct {
	Params  MoqUnion_starGenType_Len_params
	Results []struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqUnion_starGenType_Len_doFn
		DoReturnFn MoqUnion_starGenType_Len_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqUnion_starGenType_Len_fnRecorder routes recorded function calls to the
// MoqUnion_starGenType moq
type MoqUnion_starGenType_Len_fnRecorder struct {
	Params    MoqUnion_starGenType_Len_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqUnion_starGenType_Len_results
	Moq       *MoqUnion_starGenType
}

// MoqUnion_starGenType_Len_anyParams isolates the any params functions of the
// Union_starGenType type
type MoqUnion_starGenType_Len_anyParams struct {
	Recorder *MoqUnion_starGenType_Len_fnRecorder
}

// MoqUnion_starGenType_Term_params holds the params of the Union_starGenType
// type
type MoqUnion_starGenType_Term_params struct{ Param1 int }

// MoqUnion_starGenType_Term_paramsKey holds the map key params of the
// Union_starGenType type
type MoqUnion_starGenType_Term_paramsKey struct {
	Params struct{ Param1 int }
	Hashes struct{ Param1 hash.Hash }
}

// MoqUnion_starGenType_Term_resultsByParams contains the results for a given
// set of parameters for the Union_starGenType type
type MoqUnion_starGenType_Term_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqUnion_starGenType_Term_paramsKey]*MoqUnion_starGenType_Term_results
}

// MoqUnion_starGenType_Term_doFn defines the type of function needed when
// calling AndDo for the Union_starGenType type
type MoqUnion_starGenType_Term_doFn func(i int)

// MoqUnion_starGenType_Term_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Union_starGenType type
type MoqUnion_starGenType_Term_doReturnFn func(i int) *types.Term

// MoqUnion_starGenType_Term_results holds the results of the Union_starGenType
// type
type MoqUnion_starGenType_Term_results struct {
	Params  MoqUnion_starGenType_Term_params
	Results []struct {
		Values *struct {
			Result1 *types.Term
		}
		Sequence   uint32
		DoFn       MoqUnion_starGenType_Term_doFn
		DoReturnFn MoqUnion_starGenType_Term_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqUnion_starGenType_Term_fnRecorder routes recorded function calls to the
// MoqUnion_starGenType moq
type MoqUnion_starGenType_Term_fnRecorder struct {
	Params    MoqUnion_starGenType_Term_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqUnion_starGenType_Term_results
	Moq       *MoqUnion_starGenType
}

// MoqUnion_starGenType_Term_anyParams isolates the any params functions of the
// Union_starGenType type
type MoqUnion_starGenType_Term_anyParams struct {
	Recorder *MoqUnion_starGenType_Term_fnRecorder
}

// MoqUnion_starGenType_Underlying_params holds the params of the
// Union_starGenType type
type MoqUnion_starGenType_Underlying_params struct{}

// MoqUnion_starGenType_Underlying_paramsKey holds the map key params of the
// Union_starGenType type
type MoqUnion_starGenType_Underlying_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqUnion_starGenType_Underlying_resultsByParams contains the results for a
// given set of parameters for the Union_starGenType type
type MoqUnion_starGenType_Underlying_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqUnion_starGenType_Underlying_paramsKey]*MoqUnion_starGenType_Underlying_results
}

// MoqUnion_starGenType_Underlying_doFn defines the type of function needed
// when calling AndDo for the Union_starGenType type
type MoqUnion_starGenType_Underlying_doFn func()

// MoqUnion_starGenType_Underlying_doReturnFn defines the type of function
// needed when calling DoReturnResults for the Union_starGenType type
type MoqUnion_starGenType_Underlying_doReturnFn func() types.Type

// MoqUnion_starGenType_Underlying_results holds the results of the
// Union_starGenType type
type MoqUnion_starGenType_Underlying_results struct {
	Params  MoqUnion_starGenType_Underlying_params
	Results []struct {
		Values *struct {
			Result1 types.Type
		}
		Sequence   uint32
		DoFn       MoqUnion_starGenType_Underlying_doFn
		DoReturnFn MoqUnion_starGenType_Underlying_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqUnion_starGenType_Underlying_fnRecorder routes recorded function calls to
// the MoqUnion_starGenType moq
type MoqUnion_starGenType_Underlying_fnRecorder struct {
	Params    MoqUnion_starGenType_Underlying_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqUnion_starGenType_Underlying_results
	Moq       *MoqUnion_starGenType
}

// MoqUnion_starGenType_Underlying_anyParams isolates the any params functions
// of the Union_starGenType type
type MoqUnion_starGenType_Underlying_anyParams struct {
	Recorder *MoqUnion_starGenType_Underlying_fnRecorder
}

// MoqUnion_starGenType_String_params holds the params of the Union_starGenType
// type
type MoqUnion_starGenType_String_params struct{}

// MoqUnion_starGenType_String_paramsKey holds the map key params of the
// Union_starGenType type
type MoqUnion_starGenType_String_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqUnion_starGenType_String_resultsByParams contains the results for a given
// set of parameters for the Union_starGenType type
type MoqUnion_starGenType_String_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqUnion_starGenType_String_paramsKey]*MoqUnion_starGenType_String_results
}

// MoqUnion_starGenType_String_doFn defines the type of function needed when
// calling AndDo for the Union_starGenType type
type MoqUnion_starGenType_String_doFn func()

// MoqUnion_starGenType_String_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Union_starGenType type
type MoqUnion_starGenType_String_doReturnFn func() string

// MoqUnion_starGenType_String_results holds the results of the
// Union_starGenType type
type MoqUnion_starGenType_String_results struct {
	Params  MoqUnion_starGenType_String_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqUnion_starGenType_String_doFn
		DoReturnFn MoqUnion_starGenType_String_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqUnion_starGenType_String_fnRecorder routes recorded function calls to the
// MoqUnion_starGenType moq
type MoqUnion_starGenType_String_fnRecorder struct {
	Params    MoqUnion_starGenType_String_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqUnion_starGenType_String_results
	Moq       *MoqUnion_starGenType
}

// MoqUnion_starGenType_String_anyParams isolates the any params functions of
// the Union_starGenType type
type MoqUnion_starGenType_String_anyParams struct {
	Recorder *MoqUnion_starGenType_String_fnRecorder
}

// NewMoqUnion_starGenType creates a new moq of the Union_starGenType type
func NewMoqUnion_starGenType(scene *moq.Scene, config *moq.Config) *MoqUnion_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqUnion_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqUnion_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Len  struct{}
				Term struct {
					Param1 moq.ParamIndexing
				}
				Underlying struct{}
				String     struct{}
			}
		}{ParameterIndexing: struct {
			Len  struct{}
			Term struct {
				Param1 moq.ParamIndexing
			}
			Underlying struct{}
			String     struct{}
		}{
			Len: struct{}{},
			Term: struct {
				Param1 moq.ParamIndexing
			}{
				Param1: moq.ParamIndexByValue,
			},
			Underlying: struct{}{},
			String:     struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Union_starGenType type
func (m *MoqUnion_starGenType) Mock() *MoqUnion_starGenType_mock { return m.Moq }

func (m *MoqUnion_starGenType_mock) Len() (result1 int) {
	m.Moq.Scene.T.Helper()
	params := MoqUnion_starGenType_Len_params{}
	var results *MoqUnion_starGenType_Len_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Len {
		paramsKey := m.Moq.ParamsKey_Len(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Len(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Len(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Len(params))
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

func (m *MoqUnion_starGenType_mock) Term(param1 int) (result1 *types.Term) {
	m.Moq.Scene.T.Helper()
	params := MoqUnion_starGenType_Term_params{
		Param1: param1,
	}
	var results *MoqUnion_starGenType_Term_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Term {
		paramsKey := m.Moq.ParamsKey_Term(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Term(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Term(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Term(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(param1)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(param1)
	}
	return
}

func (m *MoqUnion_starGenType_mock) Underlying() (result1 types.Type) {
	m.Moq.Scene.T.Helper()
	params := MoqUnion_starGenType_Underlying_params{}
	var results *MoqUnion_starGenType_Underlying_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Underlying {
		paramsKey := m.Moq.ParamsKey_Underlying(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Underlying(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Underlying(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Underlying(params))
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

func (m *MoqUnion_starGenType_mock) String() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqUnion_starGenType_String_params{}
	var results *MoqUnion_starGenType_String_results
	for _, resultsByParams := range m.Moq.ResultsByParams_String {
		paramsKey := m.Moq.ParamsKey_String(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_String(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_String(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_String(params))
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

// OnCall returns the recorder implementation of the Union_starGenType type
func (m *MoqUnion_starGenType) OnCall() *MoqUnion_starGenType_recorder {
	return &MoqUnion_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqUnion_starGenType_recorder) Len() *MoqUnion_starGenType_Len_fnRecorder {
	return &MoqUnion_starGenType_Len_fnRecorder{
		Params:   MoqUnion_starGenType_Len_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqUnion_starGenType_Len_fnRecorder) Any() *MoqUnion_starGenType_Len_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Len(r.Params))
		return nil
	}
	return &MoqUnion_starGenType_Len_anyParams{Recorder: r}
}

func (r *MoqUnion_starGenType_Len_fnRecorder) Seq() *MoqUnion_starGenType_Len_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Len(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqUnion_starGenType_Len_fnRecorder) NoSeq() *MoqUnion_starGenType_Len_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Len(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqUnion_starGenType_Len_fnRecorder) ReturnResults(result1 int) *MoqUnion_starGenType_Len_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqUnion_starGenType_Len_doFn
		DoReturnFn MoqUnion_starGenType_Len_doReturnFn
	}{
		Values: &struct {
			Result1 int
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqUnion_starGenType_Len_fnRecorder) AndDo(fn MoqUnion_starGenType_Len_doFn) *MoqUnion_starGenType_Len_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqUnion_starGenType_Len_fnRecorder) DoReturnResults(fn MoqUnion_starGenType_Len_doReturnFn) *MoqUnion_starGenType_Len_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqUnion_starGenType_Len_doFn
		DoReturnFn MoqUnion_starGenType_Len_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqUnion_starGenType_Len_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqUnion_starGenType_Len_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Len {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqUnion_starGenType_Len_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqUnion_starGenType_Len_paramsKey]*MoqUnion_starGenType_Len_results{},
		}
		r.Moq.ResultsByParams_Len = append(r.Moq.ResultsByParams_Len, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Len) {
			copy(r.Moq.ResultsByParams_Len[insertAt+1:], r.Moq.ResultsByParams_Len[insertAt:0])
			r.Moq.ResultsByParams_Len[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Len(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqUnion_starGenType_Len_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqUnion_starGenType_Len_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqUnion_starGenType_Len_fnRecorder {
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
					Result1 int
				}
				Sequence   uint32
				DoFn       MoqUnion_starGenType_Len_doFn
				DoReturnFn MoqUnion_starGenType_Len_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqUnion_starGenType) PrettyParams_Len(params MoqUnion_starGenType_Len_params) string {
	return fmt.Sprintf("Len()")
}

func (m *MoqUnion_starGenType) ParamsKey_Len(params MoqUnion_starGenType_Len_params, anyParams uint64) MoqUnion_starGenType_Len_paramsKey {
	m.Scene.T.Helper()
	return MoqUnion_starGenType_Len_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqUnion_starGenType_recorder) Term(param1 int) *MoqUnion_starGenType_Term_fnRecorder {
	return &MoqUnion_starGenType_Term_fnRecorder{
		Params: MoqUnion_starGenType_Term_params{
			Param1: param1,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqUnion_starGenType_Term_fnRecorder) Any() *MoqUnion_starGenType_Term_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Term(r.Params))
		return nil
	}
	return &MoqUnion_starGenType_Term_anyParams{Recorder: r}
}

func (a *MoqUnion_starGenType_Term_anyParams) Param1() *MoqUnion_starGenType_Term_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqUnion_starGenType_Term_fnRecorder) Seq() *MoqUnion_starGenType_Term_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Term(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqUnion_starGenType_Term_fnRecorder) NoSeq() *MoqUnion_starGenType_Term_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Term(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqUnion_starGenType_Term_fnRecorder) ReturnResults(result1 *types.Term) *MoqUnion_starGenType_Term_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *types.Term
		}
		Sequence   uint32
		DoFn       MoqUnion_starGenType_Term_doFn
		DoReturnFn MoqUnion_starGenType_Term_doReturnFn
	}{
		Values: &struct {
			Result1 *types.Term
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqUnion_starGenType_Term_fnRecorder) AndDo(fn MoqUnion_starGenType_Term_doFn) *MoqUnion_starGenType_Term_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqUnion_starGenType_Term_fnRecorder) DoReturnResults(fn MoqUnion_starGenType_Term_doReturnFn) *MoqUnion_starGenType_Term_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *types.Term
		}
		Sequence   uint32
		DoFn       MoqUnion_starGenType_Term_doFn
		DoReturnFn MoqUnion_starGenType_Term_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqUnion_starGenType_Term_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqUnion_starGenType_Term_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Term {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqUnion_starGenType_Term_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqUnion_starGenType_Term_paramsKey]*MoqUnion_starGenType_Term_results{},
		}
		r.Moq.ResultsByParams_Term = append(r.Moq.ResultsByParams_Term, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Term) {
			copy(r.Moq.ResultsByParams_Term[insertAt+1:], r.Moq.ResultsByParams_Term[insertAt:0])
			r.Moq.ResultsByParams_Term[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Term(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqUnion_starGenType_Term_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqUnion_starGenType_Term_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqUnion_starGenType_Term_fnRecorder {
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
					Result1 *types.Term
				}
				Sequence   uint32
				DoFn       MoqUnion_starGenType_Term_doFn
				DoReturnFn MoqUnion_starGenType_Term_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqUnion_starGenType) PrettyParams_Term(params MoqUnion_starGenType_Term_params) string {
	return fmt.Sprintf("Term(%#v)", params.Param1)
}

func (m *MoqUnion_starGenType) ParamsKey_Term(params MoqUnion_starGenType_Term_params, anyParams uint64) MoqUnion_starGenType_Term_paramsKey {
	m.Scene.T.Helper()
	var param1Used int
	var param1UsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Term.Param1 == moq.ParamIndexByValue {
			param1Used = params.Param1
		} else {
			param1UsedHash = hash.DeepHash(params.Param1)
		}
	}
	return MoqUnion_starGenType_Term_paramsKey{
		Params: struct{ Param1 int }{
			Param1: param1Used,
		},
		Hashes: struct{ Param1 hash.Hash }{
			Param1: param1UsedHash,
		},
	}
}

func (m *MoqUnion_starGenType_recorder) Underlying() *MoqUnion_starGenType_Underlying_fnRecorder {
	return &MoqUnion_starGenType_Underlying_fnRecorder{
		Params:   MoqUnion_starGenType_Underlying_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqUnion_starGenType_Underlying_fnRecorder) Any() *MoqUnion_starGenType_Underlying_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Underlying(r.Params))
		return nil
	}
	return &MoqUnion_starGenType_Underlying_anyParams{Recorder: r}
}

func (r *MoqUnion_starGenType_Underlying_fnRecorder) Seq() *MoqUnion_starGenType_Underlying_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Underlying(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqUnion_starGenType_Underlying_fnRecorder) NoSeq() *MoqUnion_starGenType_Underlying_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Underlying(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqUnion_starGenType_Underlying_fnRecorder) ReturnResults(result1 types.Type) *MoqUnion_starGenType_Underlying_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 types.Type
		}
		Sequence   uint32
		DoFn       MoqUnion_starGenType_Underlying_doFn
		DoReturnFn MoqUnion_starGenType_Underlying_doReturnFn
	}{
		Values: &struct {
			Result1 types.Type
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqUnion_starGenType_Underlying_fnRecorder) AndDo(fn MoqUnion_starGenType_Underlying_doFn) *MoqUnion_starGenType_Underlying_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqUnion_starGenType_Underlying_fnRecorder) DoReturnResults(fn MoqUnion_starGenType_Underlying_doReturnFn) *MoqUnion_starGenType_Underlying_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 types.Type
		}
		Sequence   uint32
		DoFn       MoqUnion_starGenType_Underlying_doFn
		DoReturnFn MoqUnion_starGenType_Underlying_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqUnion_starGenType_Underlying_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqUnion_starGenType_Underlying_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Underlying {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqUnion_starGenType_Underlying_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqUnion_starGenType_Underlying_paramsKey]*MoqUnion_starGenType_Underlying_results{},
		}
		r.Moq.ResultsByParams_Underlying = append(r.Moq.ResultsByParams_Underlying, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Underlying) {
			copy(r.Moq.ResultsByParams_Underlying[insertAt+1:], r.Moq.ResultsByParams_Underlying[insertAt:0])
			r.Moq.ResultsByParams_Underlying[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Underlying(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqUnion_starGenType_Underlying_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqUnion_starGenType_Underlying_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqUnion_starGenType_Underlying_fnRecorder {
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
					Result1 types.Type
				}
				Sequence   uint32
				DoFn       MoqUnion_starGenType_Underlying_doFn
				DoReturnFn MoqUnion_starGenType_Underlying_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqUnion_starGenType) PrettyParams_Underlying(params MoqUnion_starGenType_Underlying_params) string {
	return fmt.Sprintf("Underlying()")
}

func (m *MoqUnion_starGenType) ParamsKey_Underlying(params MoqUnion_starGenType_Underlying_params, anyParams uint64) MoqUnion_starGenType_Underlying_paramsKey {
	m.Scene.T.Helper()
	return MoqUnion_starGenType_Underlying_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqUnion_starGenType_recorder) String() *MoqUnion_starGenType_String_fnRecorder {
	return &MoqUnion_starGenType_String_fnRecorder{
		Params:   MoqUnion_starGenType_String_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqUnion_starGenType_String_fnRecorder) Any() *MoqUnion_starGenType_String_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	return &MoqUnion_starGenType_String_anyParams{Recorder: r}
}

func (r *MoqUnion_starGenType_String_fnRecorder) Seq() *MoqUnion_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqUnion_starGenType_String_fnRecorder) NoSeq() *MoqUnion_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqUnion_starGenType_String_fnRecorder) ReturnResults(result1 string) *MoqUnion_starGenType_String_fnRecorder {
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
		DoFn       MoqUnion_starGenType_String_doFn
		DoReturnFn MoqUnion_starGenType_String_doReturnFn
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

func (r *MoqUnion_starGenType_String_fnRecorder) AndDo(fn MoqUnion_starGenType_String_doFn) *MoqUnion_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqUnion_starGenType_String_fnRecorder) DoReturnResults(fn MoqUnion_starGenType_String_doReturnFn) *MoqUnion_starGenType_String_fnRecorder {
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
		DoFn       MoqUnion_starGenType_String_doFn
		DoReturnFn MoqUnion_starGenType_String_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqUnion_starGenType_String_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqUnion_starGenType_String_resultsByParams
	for n, res := range r.Moq.ResultsByParams_String {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqUnion_starGenType_String_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqUnion_starGenType_String_paramsKey]*MoqUnion_starGenType_String_results{},
		}
		r.Moq.ResultsByParams_String = append(r.Moq.ResultsByParams_String, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_String) {
			copy(r.Moq.ResultsByParams_String[insertAt+1:], r.Moq.ResultsByParams_String[insertAt:0])
			r.Moq.ResultsByParams_String[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_String(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqUnion_starGenType_String_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqUnion_starGenType_String_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqUnion_starGenType_String_fnRecorder {
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
				DoFn       MoqUnion_starGenType_String_doFn
				DoReturnFn MoqUnion_starGenType_String_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqUnion_starGenType) PrettyParams_String(params MoqUnion_starGenType_String_params) string {
	return fmt.Sprintf("String()")
}

func (m *MoqUnion_starGenType) ParamsKey_String(params MoqUnion_starGenType_String_params, anyParams uint64) MoqUnion_starGenType_String_paramsKey {
	m.Scene.T.Helper()
	return MoqUnion_starGenType_String_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqUnion_starGenType) Reset() {
	m.ResultsByParams_Len = nil
	m.ResultsByParams_Term = nil
	m.ResultsByParams_Underlying = nil
	m.ResultsByParams_String = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqUnion_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Len {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Len(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Term {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Term(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Underlying {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Underlying(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_String {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_String(results.Params))
			}
		}
	}
}
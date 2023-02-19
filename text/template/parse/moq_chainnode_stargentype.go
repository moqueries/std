// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package parse

import (
	"fmt"
	"math/bits"
	"sync/atomic"
	"text/template/parse"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that parse.ChainNode_starGenType is
// mocked completely
var _ ChainNode_starGenType = (*MoqChainNode_starGenType_mock)(nil)

// ChainNode_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type ChainNode_starGenType interface {
	Add(field string)
	String() string
	Copy() parse.Node
}

// MoqChainNode_starGenType holds the state of a moq of the
// ChainNode_starGenType type
type MoqChainNode_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqChainNode_starGenType_mock

	ResultsByParams_Add    []MoqChainNode_starGenType_Add_resultsByParams
	ResultsByParams_String []MoqChainNode_starGenType_String_resultsByParams
	ResultsByParams_Copy   []MoqChainNode_starGenType_Copy_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Add struct {
				Field moq.ParamIndexing
			}
			String struct{}
			Copy   struct{}
		}
	}
}

// MoqChainNode_starGenType_mock isolates the mock interface of the
// ChainNode_starGenType type
type MoqChainNode_starGenType_mock struct {
	Moq *MoqChainNode_starGenType
}

// MoqChainNode_starGenType_recorder isolates the recorder interface of the
// ChainNode_starGenType type
type MoqChainNode_starGenType_recorder struct {
	Moq *MoqChainNode_starGenType
}

// MoqChainNode_starGenType_Add_params holds the params of the
// ChainNode_starGenType type
type MoqChainNode_starGenType_Add_params struct{ Field string }

// MoqChainNode_starGenType_Add_paramsKey holds the map key params of the
// ChainNode_starGenType type
type MoqChainNode_starGenType_Add_paramsKey struct {
	Params struct{ Field string }
	Hashes struct{ Field hash.Hash }
}

// MoqChainNode_starGenType_Add_resultsByParams contains the results for a
// given set of parameters for the ChainNode_starGenType type
type MoqChainNode_starGenType_Add_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqChainNode_starGenType_Add_paramsKey]*MoqChainNode_starGenType_Add_results
}

// MoqChainNode_starGenType_Add_doFn defines the type of function needed when
// calling AndDo for the ChainNode_starGenType type
type MoqChainNode_starGenType_Add_doFn func(field string)

// MoqChainNode_starGenType_Add_doReturnFn defines the type of function needed
// when calling DoReturnResults for the ChainNode_starGenType type
type MoqChainNode_starGenType_Add_doReturnFn func(field string)

// MoqChainNode_starGenType_Add_results holds the results of the
// ChainNode_starGenType type
type MoqChainNode_starGenType_Add_results struct {
	Params  MoqChainNode_starGenType_Add_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqChainNode_starGenType_Add_doFn
		DoReturnFn MoqChainNode_starGenType_Add_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqChainNode_starGenType_Add_fnRecorder routes recorded function calls to
// the MoqChainNode_starGenType moq
type MoqChainNode_starGenType_Add_fnRecorder struct {
	Params    MoqChainNode_starGenType_Add_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqChainNode_starGenType_Add_results
	Moq       *MoqChainNode_starGenType
}

// MoqChainNode_starGenType_Add_anyParams isolates the any params functions of
// the ChainNode_starGenType type
type MoqChainNode_starGenType_Add_anyParams struct {
	Recorder *MoqChainNode_starGenType_Add_fnRecorder
}

// MoqChainNode_starGenType_String_params holds the params of the
// ChainNode_starGenType type
type MoqChainNode_starGenType_String_params struct{}

// MoqChainNode_starGenType_String_paramsKey holds the map key params of the
// ChainNode_starGenType type
type MoqChainNode_starGenType_String_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqChainNode_starGenType_String_resultsByParams contains the results for a
// given set of parameters for the ChainNode_starGenType type
type MoqChainNode_starGenType_String_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqChainNode_starGenType_String_paramsKey]*MoqChainNode_starGenType_String_results
}

// MoqChainNode_starGenType_String_doFn defines the type of function needed
// when calling AndDo for the ChainNode_starGenType type
type MoqChainNode_starGenType_String_doFn func()

// MoqChainNode_starGenType_String_doReturnFn defines the type of function
// needed when calling DoReturnResults for the ChainNode_starGenType type
type MoqChainNode_starGenType_String_doReturnFn func() string

// MoqChainNode_starGenType_String_results holds the results of the
// ChainNode_starGenType type
type MoqChainNode_starGenType_String_results struct {
	Params  MoqChainNode_starGenType_String_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqChainNode_starGenType_String_doFn
		DoReturnFn MoqChainNode_starGenType_String_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqChainNode_starGenType_String_fnRecorder routes recorded function calls to
// the MoqChainNode_starGenType moq
type MoqChainNode_starGenType_String_fnRecorder struct {
	Params    MoqChainNode_starGenType_String_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqChainNode_starGenType_String_results
	Moq       *MoqChainNode_starGenType
}

// MoqChainNode_starGenType_String_anyParams isolates the any params functions
// of the ChainNode_starGenType type
type MoqChainNode_starGenType_String_anyParams struct {
	Recorder *MoqChainNode_starGenType_String_fnRecorder
}

// MoqChainNode_starGenType_Copy_params holds the params of the
// ChainNode_starGenType type
type MoqChainNode_starGenType_Copy_params struct{}

// MoqChainNode_starGenType_Copy_paramsKey holds the map key params of the
// ChainNode_starGenType type
type MoqChainNode_starGenType_Copy_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqChainNode_starGenType_Copy_resultsByParams contains the results for a
// given set of parameters for the ChainNode_starGenType type
type MoqChainNode_starGenType_Copy_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqChainNode_starGenType_Copy_paramsKey]*MoqChainNode_starGenType_Copy_results
}

// MoqChainNode_starGenType_Copy_doFn defines the type of function needed when
// calling AndDo for the ChainNode_starGenType type
type MoqChainNode_starGenType_Copy_doFn func()

// MoqChainNode_starGenType_Copy_doReturnFn defines the type of function needed
// when calling DoReturnResults for the ChainNode_starGenType type
type MoqChainNode_starGenType_Copy_doReturnFn func() parse.Node

// MoqChainNode_starGenType_Copy_results holds the results of the
// ChainNode_starGenType type
type MoqChainNode_starGenType_Copy_results struct {
	Params  MoqChainNode_starGenType_Copy_params
	Results []struct {
		Values *struct {
			Result1 parse.Node
		}
		Sequence   uint32
		DoFn       MoqChainNode_starGenType_Copy_doFn
		DoReturnFn MoqChainNode_starGenType_Copy_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqChainNode_starGenType_Copy_fnRecorder routes recorded function calls to
// the MoqChainNode_starGenType moq
type MoqChainNode_starGenType_Copy_fnRecorder struct {
	Params    MoqChainNode_starGenType_Copy_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqChainNode_starGenType_Copy_results
	Moq       *MoqChainNode_starGenType
}

// MoqChainNode_starGenType_Copy_anyParams isolates the any params functions of
// the ChainNode_starGenType type
type MoqChainNode_starGenType_Copy_anyParams struct {
	Recorder *MoqChainNode_starGenType_Copy_fnRecorder
}

// NewMoqChainNode_starGenType creates a new moq of the ChainNode_starGenType
// type
func NewMoqChainNode_starGenType(scene *moq.Scene, config *moq.Config) *MoqChainNode_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqChainNode_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqChainNode_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Add struct {
					Field moq.ParamIndexing
				}
				String struct{}
				Copy   struct{}
			}
		}{ParameterIndexing: struct {
			Add struct {
				Field moq.ParamIndexing
			}
			String struct{}
			Copy   struct{}
		}{
			Add: struct {
				Field moq.ParamIndexing
			}{
				Field: moq.ParamIndexByValue,
			},
			String: struct{}{},
			Copy:   struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the ChainNode_starGenType type
func (m *MoqChainNode_starGenType) Mock() *MoqChainNode_starGenType_mock { return m.Moq }

func (m *MoqChainNode_starGenType_mock) Add(field string) {
	m.Moq.Scene.T.Helper()
	params := MoqChainNode_starGenType_Add_params{
		Field: field,
	}
	var results *MoqChainNode_starGenType_Add_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Add {
		paramsKey := m.Moq.ParamsKey_Add(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Add(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Add(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Add(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(field)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(field)
	}
	return
}

func (m *MoqChainNode_starGenType_mock) String() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqChainNode_starGenType_String_params{}
	var results *MoqChainNode_starGenType_String_results
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

func (m *MoqChainNode_starGenType_mock) Copy() (result1 parse.Node) {
	m.Moq.Scene.T.Helper()
	params := MoqChainNode_starGenType_Copy_params{}
	var results *MoqChainNode_starGenType_Copy_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Copy {
		paramsKey := m.Moq.ParamsKey_Copy(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Copy(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Copy(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Copy(params))
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

// OnCall returns the recorder implementation of the ChainNode_starGenType type
func (m *MoqChainNode_starGenType) OnCall() *MoqChainNode_starGenType_recorder {
	return &MoqChainNode_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqChainNode_starGenType_recorder) Add(field string) *MoqChainNode_starGenType_Add_fnRecorder {
	return &MoqChainNode_starGenType_Add_fnRecorder{
		Params: MoqChainNode_starGenType_Add_params{
			Field: field,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqChainNode_starGenType_Add_fnRecorder) Any() *MoqChainNode_starGenType_Add_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Add(r.Params))
		return nil
	}
	return &MoqChainNode_starGenType_Add_anyParams{Recorder: r}
}

func (a *MoqChainNode_starGenType_Add_anyParams) Field() *MoqChainNode_starGenType_Add_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqChainNode_starGenType_Add_fnRecorder) Seq() *MoqChainNode_starGenType_Add_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Add(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqChainNode_starGenType_Add_fnRecorder) NoSeq() *MoqChainNode_starGenType_Add_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Add(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqChainNode_starGenType_Add_fnRecorder) ReturnResults() *MoqChainNode_starGenType_Add_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqChainNode_starGenType_Add_doFn
		DoReturnFn MoqChainNode_starGenType_Add_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqChainNode_starGenType_Add_fnRecorder) AndDo(fn MoqChainNode_starGenType_Add_doFn) *MoqChainNode_starGenType_Add_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqChainNode_starGenType_Add_fnRecorder) DoReturnResults(fn MoqChainNode_starGenType_Add_doReturnFn) *MoqChainNode_starGenType_Add_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqChainNode_starGenType_Add_doFn
		DoReturnFn MoqChainNode_starGenType_Add_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqChainNode_starGenType_Add_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqChainNode_starGenType_Add_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Add {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqChainNode_starGenType_Add_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqChainNode_starGenType_Add_paramsKey]*MoqChainNode_starGenType_Add_results{},
		}
		r.Moq.ResultsByParams_Add = append(r.Moq.ResultsByParams_Add, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Add) {
			copy(r.Moq.ResultsByParams_Add[insertAt+1:], r.Moq.ResultsByParams_Add[insertAt:0])
			r.Moq.ResultsByParams_Add[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Add(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqChainNode_starGenType_Add_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqChainNode_starGenType_Add_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqChainNode_starGenType_Add_fnRecorder {
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
				DoFn       MoqChainNode_starGenType_Add_doFn
				DoReturnFn MoqChainNode_starGenType_Add_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqChainNode_starGenType) PrettyParams_Add(params MoqChainNode_starGenType_Add_params) string {
	return fmt.Sprintf("Add(%#v)", params.Field)
}

func (m *MoqChainNode_starGenType) ParamsKey_Add(params MoqChainNode_starGenType_Add_params, anyParams uint64) MoqChainNode_starGenType_Add_paramsKey {
	m.Scene.T.Helper()
	var fieldUsed string
	var fieldUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Add.Field == moq.ParamIndexByValue {
			fieldUsed = params.Field
		} else {
			fieldUsedHash = hash.DeepHash(params.Field)
		}
	}
	return MoqChainNode_starGenType_Add_paramsKey{
		Params: struct{ Field string }{
			Field: fieldUsed,
		},
		Hashes: struct{ Field hash.Hash }{
			Field: fieldUsedHash,
		},
	}
}

func (m *MoqChainNode_starGenType_recorder) String() *MoqChainNode_starGenType_String_fnRecorder {
	return &MoqChainNode_starGenType_String_fnRecorder{
		Params:   MoqChainNode_starGenType_String_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqChainNode_starGenType_String_fnRecorder) Any() *MoqChainNode_starGenType_String_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	return &MoqChainNode_starGenType_String_anyParams{Recorder: r}
}

func (r *MoqChainNode_starGenType_String_fnRecorder) Seq() *MoqChainNode_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqChainNode_starGenType_String_fnRecorder) NoSeq() *MoqChainNode_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqChainNode_starGenType_String_fnRecorder) ReturnResults(result1 string) *MoqChainNode_starGenType_String_fnRecorder {
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
		DoFn       MoqChainNode_starGenType_String_doFn
		DoReturnFn MoqChainNode_starGenType_String_doReturnFn
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

func (r *MoqChainNode_starGenType_String_fnRecorder) AndDo(fn MoqChainNode_starGenType_String_doFn) *MoqChainNode_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqChainNode_starGenType_String_fnRecorder) DoReturnResults(fn MoqChainNode_starGenType_String_doReturnFn) *MoqChainNode_starGenType_String_fnRecorder {
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
		DoFn       MoqChainNode_starGenType_String_doFn
		DoReturnFn MoqChainNode_starGenType_String_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqChainNode_starGenType_String_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqChainNode_starGenType_String_resultsByParams
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
		results = &MoqChainNode_starGenType_String_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqChainNode_starGenType_String_paramsKey]*MoqChainNode_starGenType_String_results{},
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
		r.Results = &MoqChainNode_starGenType_String_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqChainNode_starGenType_String_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqChainNode_starGenType_String_fnRecorder {
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
				DoFn       MoqChainNode_starGenType_String_doFn
				DoReturnFn MoqChainNode_starGenType_String_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqChainNode_starGenType) PrettyParams_String(params MoqChainNode_starGenType_String_params) string {
	return fmt.Sprintf("String()")
}

func (m *MoqChainNode_starGenType) ParamsKey_String(params MoqChainNode_starGenType_String_params, anyParams uint64) MoqChainNode_starGenType_String_paramsKey {
	m.Scene.T.Helper()
	return MoqChainNode_starGenType_String_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqChainNode_starGenType_recorder) Copy() *MoqChainNode_starGenType_Copy_fnRecorder {
	return &MoqChainNode_starGenType_Copy_fnRecorder{
		Params:   MoqChainNode_starGenType_Copy_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqChainNode_starGenType_Copy_fnRecorder) Any() *MoqChainNode_starGenType_Copy_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Copy(r.Params))
		return nil
	}
	return &MoqChainNode_starGenType_Copy_anyParams{Recorder: r}
}

func (r *MoqChainNode_starGenType_Copy_fnRecorder) Seq() *MoqChainNode_starGenType_Copy_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Copy(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqChainNode_starGenType_Copy_fnRecorder) NoSeq() *MoqChainNode_starGenType_Copy_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Copy(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqChainNode_starGenType_Copy_fnRecorder) ReturnResults(result1 parse.Node) *MoqChainNode_starGenType_Copy_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 parse.Node
		}
		Sequence   uint32
		DoFn       MoqChainNode_starGenType_Copy_doFn
		DoReturnFn MoqChainNode_starGenType_Copy_doReturnFn
	}{
		Values: &struct {
			Result1 parse.Node
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqChainNode_starGenType_Copy_fnRecorder) AndDo(fn MoqChainNode_starGenType_Copy_doFn) *MoqChainNode_starGenType_Copy_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqChainNode_starGenType_Copy_fnRecorder) DoReturnResults(fn MoqChainNode_starGenType_Copy_doReturnFn) *MoqChainNode_starGenType_Copy_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 parse.Node
		}
		Sequence   uint32
		DoFn       MoqChainNode_starGenType_Copy_doFn
		DoReturnFn MoqChainNode_starGenType_Copy_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqChainNode_starGenType_Copy_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqChainNode_starGenType_Copy_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Copy {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqChainNode_starGenType_Copy_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqChainNode_starGenType_Copy_paramsKey]*MoqChainNode_starGenType_Copy_results{},
		}
		r.Moq.ResultsByParams_Copy = append(r.Moq.ResultsByParams_Copy, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Copy) {
			copy(r.Moq.ResultsByParams_Copy[insertAt+1:], r.Moq.ResultsByParams_Copy[insertAt:0])
			r.Moq.ResultsByParams_Copy[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Copy(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqChainNode_starGenType_Copy_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqChainNode_starGenType_Copy_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqChainNode_starGenType_Copy_fnRecorder {
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
					Result1 parse.Node
				}
				Sequence   uint32
				DoFn       MoqChainNode_starGenType_Copy_doFn
				DoReturnFn MoqChainNode_starGenType_Copy_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqChainNode_starGenType) PrettyParams_Copy(params MoqChainNode_starGenType_Copy_params) string {
	return fmt.Sprintf("Copy()")
}

func (m *MoqChainNode_starGenType) ParamsKey_Copy(params MoqChainNode_starGenType_Copy_params, anyParams uint64) MoqChainNode_starGenType_Copy_paramsKey {
	m.Scene.T.Helper()
	return MoqChainNode_starGenType_Copy_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqChainNode_starGenType) Reset() {
	m.ResultsByParams_Add = nil
	m.ResultsByParams_String = nil
	m.ResultsByParams_Copy = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqChainNode_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Add {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Add(results.Params))
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
	for _, res := range m.ResultsByParams_Copy {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Copy(results.Params))
			}
		}
	}
}
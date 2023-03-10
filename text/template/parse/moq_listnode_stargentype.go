// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package parse

import (
	"fmt"
	"math/bits"
	"sync/atomic"
	"text/template/parse"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that parse.ListNode_starGenType is
// mocked completely
var _ ListNode_starGenType = (*MoqListNode_starGenType_mock)(nil)

// ListNode_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type ListNode_starGenType interface {
	String() string
	CopyList() *parse.ListNode
	Copy() parse.Node
}

// MoqListNode_starGenType holds the state of a moq of the ListNode_starGenType
// type
type MoqListNode_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqListNode_starGenType_mock

	ResultsByParams_String   []MoqListNode_starGenType_String_resultsByParams
	ResultsByParams_CopyList []MoqListNode_starGenType_CopyList_resultsByParams
	ResultsByParams_Copy     []MoqListNode_starGenType_Copy_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			String   struct{}
			CopyList struct{}
			Copy     struct{}
		}
	}
}

// MoqListNode_starGenType_mock isolates the mock interface of the
// ListNode_starGenType type
type MoqListNode_starGenType_mock struct {
	Moq *MoqListNode_starGenType
}

// MoqListNode_starGenType_recorder isolates the recorder interface of the
// ListNode_starGenType type
type MoqListNode_starGenType_recorder struct {
	Moq *MoqListNode_starGenType
}

// MoqListNode_starGenType_String_params holds the params of the
// ListNode_starGenType type
type MoqListNode_starGenType_String_params struct{}

// MoqListNode_starGenType_String_paramsKey holds the map key params of the
// ListNode_starGenType type
type MoqListNode_starGenType_String_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqListNode_starGenType_String_resultsByParams contains the results for a
// given set of parameters for the ListNode_starGenType type
type MoqListNode_starGenType_String_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqListNode_starGenType_String_paramsKey]*MoqListNode_starGenType_String_results
}

// MoqListNode_starGenType_String_doFn defines the type of function needed when
// calling AndDo for the ListNode_starGenType type
type MoqListNode_starGenType_String_doFn func()

// MoqListNode_starGenType_String_doReturnFn defines the type of function
// needed when calling DoReturnResults for the ListNode_starGenType type
type MoqListNode_starGenType_String_doReturnFn func() string

// MoqListNode_starGenType_String_results holds the results of the
// ListNode_starGenType type
type MoqListNode_starGenType_String_results struct {
	Params  MoqListNode_starGenType_String_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqListNode_starGenType_String_doFn
		DoReturnFn MoqListNode_starGenType_String_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqListNode_starGenType_String_fnRecorder routes recorded function calls to
// the MoqListNode_starGenType moq
type MoqListNode_starGenType_String_fnRecorder struct {
	Params    MoqListNode_starGenType_String_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqListNode_starGenType_String_results
	Moq       *MoqListNode_starGenType
}

// MoqListNode_starGenType_String_anyParams isolates the any params functions
// of the ListNode_starGenType type
type MoqListNode_starGenType_String_anyParams struct {
	Recorder *MoqListNode_starGenType_String_fnRecorder
}

// MoqListNode_starGenType_CopyList_params holds the params of the
// ListNode_starGenType type
type MoqListNode_starGenType_CopyList_params struct{}

// MoqListNode_starGenType_CopyList_paramsKey holds the map key params of the
// ListNode_starGenType type
type MoqListNode_starGenType_CopyList_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqListNode_starGenType_CopyList_resultsByParams contains the results for a
// given set of parameters for the ListNode_starGenType type
type MoqListNode_starGenType_CopyList_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqListNode_starGenType_CopyList_paramsKey]*MoqListNode_starGenType_CopyList_results
}

// MoqListNode_starGenType_CopyList_doFn defines the type of function needed
// when calling AndDo for the ListNode_starGenType type
type MoqListNode_starGenType_CopyList_doFn func()

// MoqListNode_starGenType_CopyList_doReturnFn defines the type of function
// needed when calling DoReturnResults for the ListNode_starGenType type
type MoqListNode_starGenType_CopyList_doReturnFn func() *parse.ListNode

// MoqListNode_starGenType_CopyList_results holds the results of the
// ListNode_starGenType type
type MoqListNode_starGenType_CopyList_results struct {
	Params  MoqListNode_starGenType_CopyList_params
	Results []struct {
		Values *struct {
			Result1 *parse.ListNode
		}
		Sequence   uint32
		DoFn       MoqListNode_starGenType_CopyList_doFn
		DoReturnFn MoqListNode_starGenType_CopyList_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqListNode_starGenType_CopyList_fnRecorder routes recorded function calls
// to the MoqListNode_starGenType moq
type MoqListNode_starGenType_CopyList_fnRecorder struct {
	Params    MoqListNode_starGenType_CopyList_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqListNode_starGenType_CopyList_results
	Moq       *MoqListNode_starGenType
}

// MoqListNode_starGenType_CopyList_anyParams isolates the any params functions
// of the ListNode_starGenType type
type MoqListNode_starGenType_CopyList_anyParams struct {
	Recorder *MoqListNode_starGenType_CopyList_fnRecorder
}

// MoqListNode_starGenType_Copy_params holds the params of the
// ListNode_starGenType type
type MoqListNode_starGenType_Copy_params struct{}

// MoqListNode_starGenType_Copy_paramsKey holds the map key params of the
// ListNode_starGenType type
type MoqListNode_starGenType_Copy_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqListNode_starGenType_Copy_resultsByParams contains the results for a
// given set of parameters for the ListNode_starGenType type
type MoqListNode_starGenType_Copy_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqListNode_starGenType_Copy_paramsKey]*MoqListNode_starGenType_Copy_results
}

// MoqListNode_starGenType_Copy_doFn defines the type of function needed when
// calling AndDo for the ListNode_starGenType type
type MoqListNode_starGenType_Copy_doFn func()

// MoqListNode_starGenType_Copy_doReturnFn defines the type of function needed
// when calling DoReturnResults for the ListNode_starGenType type
type MoqListNode_starGenType_Copy_doReturnFn func() parse.Node

// MoqListNode_starGenType_Copy_results holds the results of the
// ListNode_starGenType type
type MoqListNode_starGenType_Copy_results struct {
	Params  MoqListNode_starGenType_Copy_params
	Results []struct {
		Values *struct {
			Result1 parse.Node
		}
		Sequence   uint32
		DoFn       MoqListNode_starGenType_Copy_doFn
		DoReturnFn MoqListNode_starGenType_Copy_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqListNode_starGenType_Copy_fnRecorder routes recorded function calls to
// the MoqListNode_starGenType moq
type MoqListNode_starGenType_Copy_fnRecorder struct {
	Params    MoqListNode_starGenType_Copy_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqListNode_starGenType_Copy_results
	Moq       *MoqListNode_starGenType
}

// MoqListNode_starGenType_Copy_anyParams isolates the any params functions of
// the ListNode_starGenType type
type MoqListNode_starGenType_Copy_anyParams struct {
	Recorder *MoqListNode_starGenType_Copy_fnRecorder
}

// NewMoqListNode_starGenType creates a new moq of the ListNode_starGenType
// type
func NewMoqListNode_starGenType(scene *moq.Scene, config *moq.Config) *MoqListNode_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqListNode_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqListNode_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				String   struct{}
				CopyList struct{}
				Copy     struct{}
			}
		}{ParameterIndexing: struct {
			String   struct{}
			CopyList struct{}
			Copy     struct{}
		}{
			String:   struct{}{},
			CopyList: struct{}{},
			Copy:     struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the ListNode_starGenType type
func (m *MoqListNode_starGenType) Mock() *MoqListNode_starGenType_mock { return m.Moq }

func (m *MoqListNode_starGenType_mock) String() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqListNode_starGenType_String_params{}
	var results *MoqListNode_starGenType_String_results
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

func (m *MoqListNode_starGenType_mock) CopyList() (result1 *parse.ListNode) {
	m.Moq.Scene.T.Helper()
	params := MoqListNode_starGenType_CopyList_params{}
	var results *MoqListNode_starGenType_CopyList_results
	for _, resultsByParams := range m.Moq.ResultsByParams_CopyList {
		paramsKey := m.Moq.ParamsKey_CopyList(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_CopyList(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_CopyList(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_CopyList(params))
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

func (m *MoqListNode_starGenType_mock) Copy() (result1 parse.Node) {
	m.Moq.Scene.T.Helper()
	params := MoqListNode_starGenType_Copy_params{}
	var results *MoqListNode_starGenType_Copy_results
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

// OnCall returns the recorder implementation of the ListNode_starGenType type
func (m *MoqListNode_starGenType) OnCall() *MoqListNode_starGenType_recorder {
	return &MoqListNode_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqListNode_starGenType_recorder) String() *MoqListNode_starGenType_String_fnRecorder {
	return &MoqListNode_starGenType_String_fnRecorder{
		Params:   MoqListNode_starGenType_String_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqListNode_starGenType_String_fnRecorder) Any() *MoqListNode_starGenType_String_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	return &MoqListNode_starGenType_String_anyParams{Recorder: r}
}

func (r *MoqListNode_starGenType_String_fnRecorder) Seq() *MoqListNode_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqListNode_starGenType_String_fnRecorder) NoSeq() *MoqListNode_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqListNode_starGenType_String_fnRecorder) ReturnResults(result1 string) *MoqListNode_starGenType_String_fnRecorder {
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
		DoFn       MoqListNode_starGenType_String_doFn
		DoReturnFn MoqListNode_starGenType_String_doReturnFn
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

func (r *MoqListNode_starGenType_String_fnRecorder) AndDo(fn MoqListNode_starGenType_String_doFn) *MoqListNode_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqListNode_starGenType_String_fnRecorder) DoReturnResults(fn MoqListNode_starGenType_String_doReturnFn) *MoqListNode_starGenType_String_fnRecorder {
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
		DoFn       MoqListNode_starGenType_String_doFn
		DoReturnFn MoqListNode_starGenType_String_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqListNode_starGenType_String_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqListNode_starGenType_String_resultsByParams
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
		results = &MoqListNode_starGenType_String_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqListNode_starGenType_String_paramsKey]*MoqListNode_starGenType_String_results{},
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
		r.Results = &MoqListNode_starGenType_String_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqListNode_starGenType_String_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqListNode_starGenType_String_fnRecorder {
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
				DoFn       MoqListNode_starGenType_String_doFn
				DoReturnFn MoqListNode_starGenType_String_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqListNode_starGenType) PrettyParams_String(params MoqListNode_starGenType_String_params) string {
	return fmt.Sprintf("String()")
}

func (m *MoqListNode_starGenType) ParamsKey_String(params MoqListNode_starGenType_String_params, anyParams uint64) MoqListNode_starGenType_String_paramsKey {
	m.Scene.T.Helper()
	return MoqListNode_starGenType_String_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqListNode_starGenType_recorder) CopyList() *MoqListNode_starGenType_CopyList_fnRecorder {
	return &MoqListNode_starGenType_CopyList_fnRecorder{
		Params:   MoqListNode_starGenType_CopyList_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqListNode_starGenType_CopyList_fnRecorder) Any() *MoqListNode_starGenType_CopyList_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_CopyList(r.Params))
		return nil
	}
	return &MoqListNode_starGenType_CopyList_anyParams{Recorder: r}
}

func (r *MoqListNode_starGenType_CopyList_fnRecorder) Seq() *MoqListNode_starGenType_CopyList_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_CopyList(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqListNode_starGenType_CopyList_fnRecorder) NoSeq() *MoqListNode_starGenType_CopyList_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_CopyList(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqListNode_starGenType_CopyList_fnRecorder) ReturnResults(result1 *parse.ListNode) *MoqListNode_starGenType_CopyList_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *parse.ListNode
		}
		Sequence   uint32
		DoFn       MoqListNode_starGenType_CopyList_doFn
		DoReturnFn MoqListNode_starGenType_CopyList_doReturnFn
	}{
		Values: &struct {
			Result1 *parse.ListNode
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqListNode_starGenType_CopyList_fnRecorder) AndDo(fn MoqListNode_starGenType_CopyList_doFn) *MoqListNode_starGenType_CopyList_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqListNode_starGenType_CopyList_fnRecorder) DoReturnResults(fn MoqListNode_starGenType_CopyList_doReturnFn) *MoqListNode_starGenType_CopyList_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *parse.ListNode
		}
		Sequence   uint32
		DoFn       MoqListNode_starGenType_CopyList_doFn
		DoReturnFn MoqListNode_starGenType_CopyList_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqListNode_starGenType_CopyList_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqListNode_starGenType_CopyList_resultsByParams
	for n, res := range r.Moq.ResultsByParams_CopyList {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqListNode_starGenType_CopyList_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqListNode_starGenType_CopyList_paramsKey]*MoqListNode_starGenType_CopyList_results{},
		}
		r.Moq.ResultsByParams_CopyList = append(r.Moq.ResultsByParams_CopyList, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_CopyList) {
			copy(r.Moq.ResultsByParams_CopyList[insertAt+1:], r.Moq.ResultsByParams_CopyList[insertAt:0])
			r.Moq.ResultsByParams_CopyList[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_CopyList(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqListNode_starGenType_CopyList_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqListNode_starGenType_CopyList_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqListNode_starGenType_CopyList_fnRecorder {
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
					Result1 *parse.ListNode
				}
				Sequence   uint32
				DoFn       MoqListNode_starGenType_CopyList_doFn
				DoReturnFn MoqListNode_starGenType_CopyList_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqListNode_starGenType) PrettyParams_CopyList(params MoqListNode_starGenType_CopyList_params) string {
	return fmt.Sprintf("CopyList()")
}

func (m *MoqListNode_starGenType) ParamsKey_CopyList(params MoqListNode_starGenType_CopyList_params, anyParams uint64) MoqListNode_starGenType_CopyList_paramsKey {
	m.Scene.T.Helper()
	return MoqListNode_starGenType_CopyList_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqListNode_starGenType_recorder) Copy() *MoqListNode_starGenType_Copy_fnRecorder {
	return &MoqListNode_starGenType_Copy_fnRecorder{
		Params:   MoqListNode_starGenType_Copy_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqListNode_starGenType_Copy_fnRecorder) Any() *MoqListNode_starGenType_Copy_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Copy(r.Params))
		return nil
	}
	return &MoqListNode_starGenType_Copy_anyParams{Recorder: r}
}

func (r *MoqListNode_starGenType_Copy_fnRecorder) Seq() *MoqListNode_starGenType_Copy_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Copy(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqListNode_starGenType_Copy_fnRecorder) NoSeq() *MoqListNode_starGenType_Copy_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Copy(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqListNode_starGenType_Copy_fnRecorder) ReturnResults(result1 parse.Node) *MoqListNode_starGenType_Copy_fnRecorder {
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
		DoFn       MoqListNode_starGenType_Copy_doFn
		DoReturnFn MoqListNode_starGenType_Copy_doReturnFn
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

func (r *MoqListNode_starGenType_Copy_fnRecorder) AndDo(fn MoqListNode_starGenType_Copy_doFn) *MoqListNode_starGenType_Copy_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqListNode_starGenType_Copy_fnRecorder) DoReturnResults(fn MoqListNode_starGenType_Copy_doReturnFn) *MoqListNode_starGenType_Copy_fnRecorder {
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
		DoFn       MoqListNode_starGenType_Copy_doFn
		DoReturnFn MoqListNode_starGenType_Copy_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqListNode_starGenType_Copy_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqListNode_starGenType_Copy_resultsByParams
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
		results = &MoqListNode_starGenType_Copy_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqListNode_starGenType_Copy_paramsKey]*MoqListNode_starGenType_Copy_results{},
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
		r.Results = &MoqListNode_starGenType_Copy_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqListNode_starGenType_Copy_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqListNode_starGenType_Copy_fnRecorder {
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
				DoFn       MoqListNode_starGenType_Copy_doFn
				DoReturnFn MoqListNode_starGenType_Copy_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqListNode_starGenType) PrettyParams_Copy(params MoqListNode_starGenType_Copy_params) string {
	return fmt.Sprintf("Copy()")
}

func (m *MoqListNode_starGenType) ParamsKey_Copy(params MoqListNode_starGenType_Copy_params, anyParams uint64) MoqListNode_starGenType_Copy_paramsKey {
	m.Scene.T.Helper()
	return MoqListNode_starGenType_Copy_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqListNode_starGenType) Reset() {
	m.ResultsByParams_String = nil
	m.ResultsByParams_CopyList = nil
	m.ResultsByParams_Copy = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqListNode_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_String {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_String(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_CopyList {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_CopyList(results.Params))
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

// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package cookiejar

import (
	"fmt"
	"math/bits"
	"net/http/cookiejar"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that cookiejar.PublicSuffixList is
// mocked completely
var _ cookiejar.PublicSuffixList = (*MoqPublicSuffixList_mock)(nil)

// MoqPublicSuffixList holds the state of a moq of the PublicSuffixList type
type MoqPublicSuffixList struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqPublicSuffixList_mock

	ResultsByParams_PublicSuffix []MoqPublicSuffixList_PublicSuffix_resultsByParams
	ResultsByParams_String       []MoqPublicSuffixList_String_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			PublicSuffix struct {
				Domain moq.ParamIndexing
			}
			String struct{}
		}
	}
}

// MoqPublicSuffixList_mock isolates the mock interface of the PublicSuffixList
// type
type MoqPublicSuffixList_mock struct {
	Moq *MoqPublicSuffixList
}

// MoqPublicSuffixList_recorder isolates the recorder interface of the
// PublicSuffixList type
type MoqPublicSuffixList_recorder struct {
	Moq *MoqPublicSuffixList
}

// MoqPublicSuffixList_PublicSuffix_params holds the params of the
// PublicSuffixList type
type MoqPublicSuffixList_PublicSuffix_params struct{ Domain string }

// MoqPublicSuffixList_PublicSuffix_paramsKey holds the map key params of the
// PublicSuffixList type
type MoqPublicSuffixList_PublicSuffix_paramsKey struct {
	Params struct{ Domain string }
	Hashes struct{ Domain hash.Hash }
}

// MoqPublicSuffixList_PublicSuffix_resultsByParams contains the results for a
// given set of parameters for the PublicSuffixList type
type MoqPublicSuffixList_PublicSuffix_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqPublicSuffixList_PublicSuffix_paramsKey]*MoqPublicSuffixList_PublicSuffix_results
}

// MoqPublicSuffixList_PublicSuffix_doFn defines the type of function needed
// when calling AndDo for the PublicSuffixList type
type MoqPublicSuffixList_PublicSuffix_doFn func(domain string)

// MoqPublicSuffixList_PublicSuffix_doReturnFn defines the type of function
// needed when calling DoReturnResults for the PublicSuffixList type
type MoqPublicSuffixList_PublicSuffix_doReturnFn func(domain string) string

// MoqPublicSuffixList_PublicSuffix_results holds the results of the
// PublicSuffixList type
type MoqPublicSuffixList_PublicSuffix_results struct {
	Params  MoqPublicSuffixList_PublicSuffix_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqPublicSuffixList_PublicSuffix_doFn
		DoReturnFn MoqPublicSuffixList_PublicSuffix_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqPublicSuffixList_PublicSuffix_fnRecorder routes recorded function calls
// to the MoqPublicSuffixList moq
type MoqPublicSuffixList_PublicSuffix_fnRecorder struct {
	Params    MoqPublicSuffixList_PublicSuffix_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqPublicSuffixList_PublicSuffix_results
	Moq       *MoqPublicSuffixList
}

// MoqPublicSuffixList_PublicSuffix_anyParams isolates the any params functions
// of the PublicSuffixList type
type MoqPublicSuffixList_PublicSuffix_anyParams struct {
	Recorder *MoqPublicSuffixList_PublicSuffix_fnRecorder
}

// MoqPublicSuffixList_String_params holds the params of the PublicSuffixList
// type
type MoqPublicSuffixList_String_params struct{}

// MoqPublicSuffixList_String_paramsKey holds the map key params of the
// PublicSuffixList type
type MoqPublicSuffixList_String_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqPublicSuffixList_String_resultsByParams contains the results for a given
// set of parameters for the PublicSuffixList type
type MoqPublicSuffixList_String_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqPublicSuffixList_String_paramsKey]*MoqPublicSuffixList_String_results
}

// MoqPublicSuffixList_String_doFn defines the type of function needed when
// calling AndDo for the PublicSuffixList type
type MoqPublicSuffixList_String_doFn func()

// MoqPublicSuffixList_String_doReturnFn defines the type of function needed
// when calling DoReturnResults for the PublicSuffixList type
type MoqPublicSuffixList_String_doReturnFn func() string

// MoqPublicSuffixList_String_results holds the results of the PublicSuffixList
// type
type MoqPublicSuffixList_String_results struct {
	Params  MoqPublicSuffixList_String_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqPublicSuffixList_String_doFn
		DoReturnFn MoqPublicSuffixList_String_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqPublicSuffixList_String_fnRecorder routes recorded function calls to the
// MoqPublicSuffixList moq
type MoqPublicSuffixList_String_fnRecorder struct {
	Params    MoqPublicSuffixList_String_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqPublicSuffixList_String_results
	Moq       *MoqPublicSuffixList
}

// MoqPublicSuffixList_String_anyParams isolates the any params functions of
// the PublicSuffixList type
type MoqPublicSuffixList_String_anyParams struct {
	Recorder *MoqPublicSuffixList_String_fnRecorder
}

// NewMoqPublicSuffixList creates a new moq of the PublicSuffixList type
func NewMoqPublicSuffixList(scene *moq.Scene, config *moq.Config) *MoqPublicSuffixList {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqPublicSuffixList{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqPublicSuffixList_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				PublicSuffix struct {
					Domain moq.ParamIndexing
				}
				String struct{}
			}
		}{ParameterIndexing: struct {
			PublicSuffix struct {
				Domain moq.ParamIndexing
			}
			String struct{}
		}{
			PublicSuffix: struct {
				Domain moq.ParamIndexing
			}{
				Domain: moq.ParamIndexByValue,
			},
			String: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the PublicSuffixList type
func (m *MoqPublicSuffixList) Mock() *MoqPublicSuffixList_mock { return m.Moq }

func (m *MoqPublicSuffixList_mock) PublicSuffix(domain string) (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqPublicSuffixList_PublicSuffix_params{
		Domain: domain,
	}
	var results *MoqPublicSuffixList_PublicSuffix_results
	for _, resultsByParams := range m.Moq.ResultsByParams_PublicSuffix {
		paramsKey := m.Moq.ParamsKey_PublicSuffix(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_PublicSuffix(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_PublicSuffix(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_PublicSuffix(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(domain)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(domain)
	}
	return
}

func (m *MoqPublicSuffixList_mock) String() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqPublicSuffixList_String_params{}
	var results *MoqPublicSuffixList_String_results
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

// OnCall returns the recorder implementation of the PublicSuffixList type
func (m *MoqPublicSuffixList) OnCall() *MoqPublicSuffixList_recorder {
	return &MoqPublicSuffixList_recorder{
		Moq: m,
	}
}

func (m *MoqPublicSuffixList_recorder) PublicSuffix(domain string) *MoqPublicSuffixList_PublicSuffix_fnRecorder {
	return &MoqPublicSuffixList_PublicSuffix_fnRecorder{
		Params: MoqPublicSuffixList_PublicSuffix_params{
			Domain: domain,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqPublicSuffixList_PublicSuffix_fnRecorder) Any() *MoqPublicSuffixList_PublicSuffix_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_PublicSuffix(r.Params))
		return nil
	}
	return &MoqPublicSuffixList_PublicSuffix_anyParams{Recorder: r}
}

func (a *MoqPublicSuffixList_PublicSuffix_anyParams) Domain() *MoqPublicSuffixList_PublicSuffix_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqPublicSuffixList_PublicSuffix_fnRecorder) Seq() *MoqPublicSuffixList_PublicSuffix_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_PublicSuffix(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqPublicSuffixList_PublicSuffix_fnRecorder) NoSeq() *MoqPublicSuffixList_PublicSuffix_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_PublicSuffix(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqPublicSuffixList_PublicSuffix_fnRecorder) ReturnResults(result1 string) *MoqPublicSuffixList_PublicSuffix_fnRecorder {
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
		DoFn       MoqPublicSuffixList_PublicSuffix_doFn
		DoReturnFn MoqPublicSuffixList_PublicSuffix_doReturnFn
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

func (r *MoqPublicSuffixList_PublicSuffix_fnRecorder) AndDo(fn MoqPublicSuffixList_PublicSuffix_doFn) *MoqPublicSuffixList_PublicSuffix_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqPublicSuffixList_PublicSuffix_fnRecorder) DoReturnResults(fn MoqPublicSuffixList_PublicSuffix_doReturnFn) *MoqPublicSuffixList_PublicSuffix_fnRecorder {
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
		DoFn       MoqPublicSuffixList_PublicSuffix_doFn
		DoReturnFn MoqPublicSuffixList_PublicSuffix_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqPublicSuffixList_PublicSuffix_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqPublicSuffixList_PublicSuffix_resultsByParams
	for n, res := range r.Moq.ResultsByParams_PublicSuffix {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqPublicSuffixList_PublicSuffix_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqPublicSuffixList_PublicSuffix_paramsKey]*MoqPublicSuffixList_PublicSuffix_results{},
		}
		r.Moq.ResultsByParams_PublicSuffix = append(r.Moq.ResultsByParams_PublicSuffix, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_PublicSuffix) {
			copy(r.Moq.ResultsByParams_PublicSuffix[insertAt+1:], r.Moq.ResultsByParams_PublicSuffix[insertAt:0])
			r.Moq.ResultsByParams_PublicSuffix[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_PublicSuffix(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqPublicSuffixList_PublicSuffix_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqPublicSuffixList_PublicSuffix_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqPublicSuffixList_PublicSuffix_fnRecorder {
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
				DoFn       MoqPublicSuffixList_PublicSuffix_doFn
				DoReturnFn MoqPublicSuffixList_PublicSuffix_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqPublicSuffixList) PrettyParams_PublicSuffix(params MoqPublicSuffixList_PublicSuffix_params) string {
	return fmt.Sprintf("PublicSuffix(%#v)", params.Domain)
}

func (m *MoqPublicSuffixList) ParamsKey_PublicSuffix(params MoqPublicSuffixList_PublicSuffix_params, anyParams uint64) MoqPublicSuffixList_PublicSuffix_paramsKey {
	m.Scene.T.Helper()
	var domainUsed string
	var domainUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.PublicSuffix.Domain == moq.ParamIndexByValue {
			domainUsed = params.Domain
		} else {
			domainUsedHash = hash.DeepHash(params.Domain)
		}
	}
	return MoqPublicSuffixList_PublicSuffix_paramsKey{
		Params: struct{ Domain string }{
			Domain: domainUsed,
		},
		Hashes: struct{ Domain hash.Hash }{
			Domain: domainUsedHash,
		},
	}
}

func (m *MoqPublicSuffixList_recorder) String() *MoqPublicSuffixList_String_fnRecorder {
	return &MoqPublicSuffixList_String_fnRecorder{
		Params:   MoqPublicSuffixList_String_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqPublicSuffixList_String_fnRecorder) Any() *MoqPublicSuffixList_String_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	return &MoqPublicSuffixList_String_anyParams{Recorder: r}
}

func (r *MoqPublicSuffixList_String_fnRecorder) Seq() *MoqPublicSuffixList_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqPublicSuffixList_String_fnRecorder) NoSeq() *MoqPublicSuffixList_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqPublicSuffixList_String_fnRecorder) ReturnResults(result1 string) *MoqPublicSuffixList_String_fnRecorder {
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
		DoFn       MoqPublicSuffixList_String_doFn
		DoReturnFn MoqPublicSuffixList_String_doReturnFn
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

func (r *MoqPublicSuffixList_String_fnRecorder) AndDo(fn MoqPublicSuffixList_String_doFn) *MoqPublicSuffixList_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqPublicSuffixList_String_fnRecorder) DoReturnResults(fn MoqPublicSuffixList_String_doReturnFn) *MoqPublicSuffixList_String_fnRecorder {
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
		DoFn       MoqPublicSuffixList_String_doFn
		DoReturnFn MoqPublicSuffixList_String_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqPublicSuffixList_String_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqPublicSuffixList_String_resultsByParams
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
		results = &MoqPublicSuffixList_String_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqPublicSuffixList_String_paramsKey]*MoqPublicSuffixList_String_results{},
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
		r.Results = &MoqPublicSuffixList_String_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqPublicSuffixList_String_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqPublicSuffixList_String_fnRecorder {
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
				DoFn       MoqPublicSuffixList_String_doFn
				DoReturnFn MoqPublicSuffixList_String_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqPublicSuffixList) PrettyParams_String(params MoqPublicSuffixList_String_params) string {
	return fmt.Sprintf("String()")
}

func (m *MoqPublicSuffixList) ParamsKey_String(params MoqPublicSuffixList_String_params, anyParams uint64) MoqPublicSuffixList_String_paramsKey {
	m.Scene.T.Helper()
	return MoqPublicSuffixList_String_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqPublicSuffixList) Reset() {
	m.ResultsByParams_PublicSuffix = nil
	m.ResultsByParams_String = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqPublicSuffixList) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_PublicSuffix {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_PublicSuffix(results.Params))
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
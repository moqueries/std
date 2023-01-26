// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package url

import (
	"fmt"
	"math/bits"
	"net/url"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// ParseQuery_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type ParseQuery_genType func(query string) (url.Values, error)

// MoqParseQuery_genType holds the state of a moq of the ParseQuery_genType
// type
type MoqParseQuery_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqParseQuery_genType_mock

	ResultsByParams []MoqParseQuery_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Query moq.ParamIndexing
		}
	}
}

// MoqParseQuery_genType_mock isolates the mock interface of the
// ParseQuery_genType type
type MoqParseQuery_genType_mock struct {
	Moq *MoqParseQuery_genType
}

// MoqParseQuery_genType_params holds the params of the ParseQuery_genType type
type MoqParseQuery_genType_params struct{ Query string }

// MoqParseQuery_genType_paramsKey holds the map key params of the
// ParseQuery_genType type
type MoqParseQuery_genType_paramsKey struct {
	Params struct{ Query string }
	Hashes struct{ Query hash.Hash }
}

// MoqParseQuery_genType_resultsByParams contains the results for a given set
// of parameters for the ParseQuery_genType type
type MoqParseQuery_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqParseQuery_genType_paramsKey]*MoqParseQuery_genType_results
}

// MoqParseQuery_genType_doFn defines the type of function needed when calling
// AndDo for the ParseQuery_genType type
type MoqParseQuery_genType_doFn func(query string)

// MoqParseQuery_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ParseQuery_genType type
type MoqParseQuery_genType_doReturnFn func(query string) (url.Values, error)

// MoqParseQuery_genType_results holds the results of the ParseQuery_genType
// type
type MoqParseQuery_genType_results struct {
	Params  MoqParseQuery_genType_params
	Results []struct {
		Values *struct {
			Result1 url.Values
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqParseQuery_genType_doFn
		DoReturnFn MoqParseQuery_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqParseQuery_genType_fnRecorder routes recorded function calls to the
// MoqParseQuery_genType moq
type MoqParseQuery_genType_fnRecorder struct {
	Params    MoqParseQuery_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqParseQuery_genType_results
	Moq       *MoqParseQuery_genType
}

// MoqParseQuery_genType_anyParams isolates the any params functions of the
// ParseQuery_genType type
type MoqParseQuery_genType_anyParams struct {
	Recorder *MoqParseQuery_genType_fnRecorder
}

// NewMoqParseQuery_genType creates a new moq of the ParseQuery_genType type
func NewMoqParseQuery_genType(scene *moq.Scene, config *moq.Config) *MoqParseQuery_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqParseQuery_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqParseQuery_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Query moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Query moq.ParamIndexing
		}{
			Query: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ParseQuery_genType type
func (m *MoqParseQuery_genType) Mock() ParseQuery_genType {
	return func(query string) (url.Values, error) {
		m.Scene.T.Helper()
		moq := &MoqParseQuery_genType_mock{Moq: m}
		return moq.Fn(query)
	}
}

func (m *MoqParseQuery_genType_mock) Fn(query string) (result1 url.Values, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqParseQuery_genType_params{
		Query: query,
	}
	var results *MoqParseQuery_genType_results
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
		result.DoFn(query)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(query)
	}
	return
}

func (m *MoqParseQuery_genType) OnCall(query string) *MoqParseQuery_genType_fnRecorder {
	return &MoqParseQuery_genType_fnRecorder{
		Params: MoqParseQuery_genType_params{
			Query: query,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqParseQuery_genType_fnRecorder) Any() *MoqParseQuery_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqParseQuery_genType_anyParams{Recorder: r}
}

func (a *MoqParseQuery_genType_anyParams) Query() *MoqParseQuery_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqParseQuery_genType_fnRecorder) Seq() *MoqParseQuery_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqParseQuery_genType_fnRecorder) NoSeq() *MoqParseQuery_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqParseQuery_genType_fnRecorder) ReturnResults(result1 url.Values, result2 error) *MoqParseQuery_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 url.Values
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqParseQuery_genType_doFn
		DoReturnFn MoqParseQuery_genType_doReturnFn
	}{
		Values: &struct {
			Result1 url.Values
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqParseQuery_genType_fnRecorder) AndDo(fn MoqParseQuery_genType_doFn) *MoqParseQuery_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqParseQuery_genType_fnRecorder) DoReturnResults(fn MoqParseQuery_genType_doReturnFn) *MoqParseQuery_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 url.Values
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqParseQuery_genType_doFn
		DoReturnFn MoqParseQuery_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqParseQuery_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqParseQuery_genType_resultsByParams
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
		results = &MoqParseQuery_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqParseQuery_genType_paramsKey]*MoqParseQuery_genType_results{},
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
		r.Results = &MoqParseQuery_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqParseQuery_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqParseQuery_genType_fnRecorder {
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
					Result1 url.Values
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqParseQuery_genType_doFn
				DoReturnFn MoqParseQuery_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqParseQuery_genType) PrettyParams(params MoqParseQuery_genType_params) string {
	return fmt.Sprintf("ParseQuery_genType(%#v)", params.Query)
}

func (m *MoqParseQuery_genType) ParamsKey(params MoqParseQuery_genType_params, anyParams uint64) MoqParseQuery_genType_paramsKey {
	m.Scene.T.Helper()
	var queryUsed string
	var queryUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Query == moq.ParamIndexByValue {
			queryUsed = params.Query
		} else {
			queryUsedHash = hash.DeepHash(params.Query)
		}
	}
	return MoqParseQuery_genType_paramsKey{
		Params: struct{ Query string }{
			Query: queryUsed,
		},
		Hashes: struct{ Query hash.Hash }{
			Query: queryUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqParseQuery_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqParseQuery_genType) AssertExpectationsMet() {
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
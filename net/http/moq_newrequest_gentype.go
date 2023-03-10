// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package http

import (
	"fmt"
	"io"
	"math/bits"
	"net/http"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// NewRequest_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type NewRequest_genType func(method, url string, body io.Reader) (*http.Request, error)

// MoqNewRequest_genType holds the state of a moq of the NewRequest_genType
// type
type MoqNewRequest_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNewRequest_genType_mock

	ResultsByParams []MoqNewRequest_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Method moq.ParamIndexing
			Url    moq.ParamIndexing
			Body   moq.ParamIndexing
		}
	}
}

// MoqNewRequest_genType_mock isolates the mock interface of the
// NewRequest_genType type
type MoqNewRequest_genType_mock struct {
	Moq *MoqNewRequest_genType
}

// MoqNewRequest_genType_params holds the params of the NewRequest_genType type
type MoqNewRequest_genType_params struct {
	Method, Url string
	Body        io.Reader
}

// MoqNewRequest_genType_paramsKey holds the map key params of the
// NewRequest_genType type
type MoqNewRequest_genType_paramsKey struct {
	Params struct {
		Method, Url string
		Body        io.Reader
	}
	Hashes struct {
		Method, Url hash.Hash
		Body        hash.Hash
	}
}

// MoqNewRequest_genType_resultsByParams contains the results for a given set
// of parameters for the NewRequest_genType type
type MoqNewRequest_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNewRequest_genType_paramsKey]*MoqNewRequest_genType_results
}

// MoqNewRequest_genType_doFn defines the type of function needed when calling
// AndDo for the NewRequest_genType type
type MoqNewRequest_genType_doFn func(method, url string, body io.Reader)

// MoqNewRequest_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the NewRequest_genType type
type MoqNewRequest_genType_doReturnFn func(method, url string, body io.Reader) (*http.Request, error)

// MoqNewRequest_genType_results holds the results of the NewRequest_genType
// type
type MoqNewRequest_genType_results struct {
	Params  MoqNewRequest_genType_params
	Results []struct {
		Values *struct {
			Result1 *http.Request
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqNewRequest_genType_doFn
		DoReturnFn MoqNewRequest_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNewRequest_genType_fnRecorder routes recorded function calls to the
// MoqNewRequest_genType moq
type MoqNewRequest_genType_fnRecorder struct {
	Params    MoqNewRequest_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNewRequest_genType_results
	Moq       *MoqNewRequest_genType
}

// MoqNewRequest_genType_anyParams isolates the any params functions of the
// NewRequest_genType type
type MoqNewRequest_genType_anyParams struct {
	Recorder *MoqNewRequest_genType_fnRecorder
}

// NewMoqNewRequest_genType creates a new moq of the NewRequest_genType type
func NewMoqNewRequest_genType(scene *moq.Scene, config *moq.Config) *MoqNewRequest_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNewRequest_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNewRequest_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Method moq.ParamIndexing
				Url    moq.ParamIndexing
				Body   moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Method moq.ParamIndexing
			Url    moq.ParamIndexing
			Body   moq.ParamIndexing
		}{
			Method: moq.ParamIndexByValue,
			Url:    moq.ParamIndexByValue,
			Body:   moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the NewRequest_genType type
func (m *MoqNewRequest_genType) Mock() NewRequest_genType {
	return func(method, url string, body io.Reader) (*http.Request, error) {
		m.Scene.T.Helper()
		moq := &MoqNewRequest_genType_mock{Moq: m}
		return moq.Fn(method, url, body)
	}
}

func (m *MoqNewRequest_genType_mock) Fn(method, url string, body io.Reader) (result1 *http.Request, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqNewRequest_genType_params{
		Method: method,
		Url:    url,
		Body:   body,
	}
	var results *MoqNewRequest_genType_results
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
		result.DoFn(method, url, body)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(method, url, body)
	}
	return
}

func (m *MoqNewRequest_genType) OnCall(method, url string, body io.Reader) *MoqNewRequest_genType_fnRecorder {
	return &MoqNewRequest_genType_fnRecorder{
		Params: MoqNewRequest_genType_params{
			Method: method,
			Url:    url,
			Body:   body,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNewRequest_genType_fnRecorder) Any() *MoqNewRequest_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNewRequest_genType_anyParams{Recorder: r}
}

func (a *MoqNewRequest_genType_anyParams) Method() *MoqNewRequest_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqNewRequest_genType_anyParams) Url() *MoqNewRequest_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqNewRequest_genType_anyParams) Body() *MoqNewRequest_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqNewRequest_genType_fnRecorder) Seq() *MoqNewRequest_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNewRequest_genType_fnRecorder) NoSeq() *MoqNewRequest_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNewRequest_genType_fnRecorder) ReturnResults(result1 *http.Request, result2 error) *MoqNewRequest_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *http.Request
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqNewRequest_genType_doFn
		DoReturnFn MoqNewRequest_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *http.Request
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNewRequest_genType_fnRecorder) AndDo(fn MoqNewRequest_genType_doFn) *MoqNewRequest_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNewRequest_genType_fnRecorder) DoReturnResults(fn MoqNewRequest_genType_doReturnFn) *MoqNewRequest_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *http.Request
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqNewRequest_genType_doFn
		DoReturnFn MoqNewRequest_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNewRequest_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNewRequest_genType_resultsByParams
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
		results = &MoqNewRequest_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNewRequest_genType_paramsKey]*MoqNewRequest_genType_results{},
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
		r.Results = &MoqNewRequest_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNewRequest_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNewRequest_genType_fnRecorder {
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
					Result1 *http.Request
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqNewRequest_genType_doFn
				DoReturnFn MoqNewRequest_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNewRequest_genType) PrettyParams(params MoqNewRequest_genType_params) string {
	return fmt.Sprintf("NewRequest_genType(%#v, %#v, %#v)", params.Method, params.Url, params.Body)
}

func (m *MoqNewRequest_genType) ParamsKey(params MoqNewRequest_genType_params, anyParams uint64) MoqNewRequest_genType_paramsKey {
	m.Scene.T.Helper()
	var methodUsed string
	var methodUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Method == moq.ParamIndexByValue {
			methodUsed = params.Method
		} else {
			methodUsedHash = hash.DeepHash(params.Method)
		}
	}
	var urlUsed string
	var urlUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Url == moq.ParamIndexByValue {
			urlUsed = params.Url
		} else {
			urlUsedHash = hash.DeepHash(params.Url)
		}
	}
	var bodyUsed io.Reader
	var bodyUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Body == moq.ParamIndexByValue {
			bodyUsed = params.Body
		} else {
			bodyUsedHash = hash.DeepHash(params.Body)
		}
	}
	return MoqNewRequest_genType_paramsKey{
		Params: struct {
			Method, Url string
			Body        io.Reader
		}{
			Method: methodUsed,
			Url:    urlUsed,
			Body:   bodyUsed,
		},
		Hashes: struct {
			Method, Url hash.Hash
			Body        hash.Hash
		}{
			Method: methodUsedHash,
			Url:    urlUsedHash,
			Body:   bodyUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNewRequest_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNewRequest_genType) AssertExpectationsMet() {
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

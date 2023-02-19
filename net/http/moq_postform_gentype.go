// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package http

import (
	"fmt"
	"math/bits"
	"net/http"
	"net/url"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// PostForm_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type PostForm_genType func(url string, data url.Values) (resp *http.Response, err error)

// MoqPostForm_genType holds the state of a moq of the PostForm_genType type
type MoqPostForm_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqPostForm_genType_mock

	ResultsByParams []MoqPostForm_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Url  moq.ParamIndexing
			Data moq.ParamIndexing
		}
	}
}

// MoqPostForm_genType_mock isolates the mock interface of the PostForm_genType
// type
type MoqPostForm_genType_mock struct {
	Moq *MoqPostForm_genType
}

// MoqPostForm_genType_params holds the params of the PostForm_genType type
type MoqPostForm_genType_params struct {
	Url  string
	Data url.Values
}

// MoqPostForm_genType_paramsKey holds the map key params of the
// PostForm_genType type
type MoqPostForm_genType_paramsKey struct {
	Params struct{ Url string }
	Hashes struct {
		Url  hash.Hash
		Data hash.Hash
	}
}

// MoqPostForm_genType_resultsByParams contains the results for a given set of
// parameters for the PostForm_genType type
type MoqPostForm_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqPostForm_genType_paramsKey]*MoqPostForm_genType_results
}

// MoqPostForm_genType_doFn defines the type of function needed when calling
// AndDo for the PostForm_genType type
type MoqPostForm_genType_doFn func(url string, data url.Values)

// MoqPostForm_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the PostForm_genType type
type MoqPostForm_genType_doReturnFn func(url string, data url.Values) (resp *http.Response, err error)

// MoqPostForm_genType_results holds the results of the PostForm_genType type
type MoqPostForm_genType_results struct {
	Params  MoqPostForm_genType_params
	Results []struct {
		Values *struct {
			Resp *http.Response
			Err  error
		}
		Sequence   uint32
		DoFn       MoqPostForm_genType_doFn
		DoReturnFn MoqPostForm_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqPostForm_genType_fnRecorder routes recorded function calls to the
// MoqPostForm_genType moq
type MoqPostForm_genType_fnRecorder struct {
	Params    MoqPostForm_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqPostForm_genType_results
	Moq       *MoqPostForm_genType
}

// MoqPostForm_genType_anyParams isolates the any params functions of the
// PostForm_genType type
type MoqPostForm_genType_anyParams struct {
	Recorder *MoqPostForm_genType_fnRecorder
}

// NewMoqPostForm_genType creates a new moq of the PostForm_genType type
func NewMoqPostForm_genType(scene *moq.Scene, config *moq.Config) *MoqPostForm_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqPostForm_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqPostForm_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Url  moq.ParamIndexing
				Data moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Url  moq.ParamIndexing
			Data moq.ParamIndexing
		}{
			Url:  moq.ParamIndexByValue,
			Data: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the PostForm_genType type
func (m *MoqPostForm_genType) Mock() PostForm_genType {
	return func(url string, data url.Values) (_ *http.Response, _ error) {
		m.Scene.T.Helper()
		moq := &MoqPostForm_genType_mock{Moq: m}
		return moq.Fn(url, data)
	}
}

func (m *MoqPostForm_genType_mock) Fn(url string, data url.Values) (resp *http.Response, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqPostForm_genType_params{
		Url:  url,
		Data: data,
	}
	var results *MoqPostForm_genType_results
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
		result.DoFn(url, data)
	}

	if result.Values != nil {
		resp = result.Values.Resp
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		resp, err = result.DoReturnFn(url, data)
	}
	return
}

func (m *MoqPostForm_genType) OnCall(url string, data url.Values) *MoqPostForm_genType_fnRecorder {
	return &MoqPostForm_genType_fnRecorder{
		Params: MoqPostForm_genType_params{
			Url:  url,
			Data: data,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqPostForm_genType_fnRecorder) Any() *MoqPostForm_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqPostForm_genType_anyParams{Recorder: r}
}

func (a *MoqPostForm_genType_anyParams) Url() *MoqPostForm_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqPostForm_genType_anyParams) Data() *MoqPostForm_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqPostForm_genType_fnRecorder) Seq() *MoqPostForm_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqPostForm_genType_fnRecorder) NoSeq() *MoqPostForm_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqPostForm_genType_fnRecorder) ReturnResults(resp *http.Response, err error) *MoqPostForm_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Resp *http.Response
			Err  error
		}
		Sequence   uint32
		DoFn       MoqPostForm_genType_doFn
		DoReturnFn MoqPostForm_genType_doReturnFn
	}{
		Values: &struct {
			Resp *http.Response
			Err  error
		}{
			Resp: resp,
			Err:  err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqPostForm_genType_fnRecorder) AndDo(fn MoqPostForm_genType_doFn) *MoqPostForm_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqPostForm_genType_fnRecorder) DoReturnResults(fn MoqPostForm_genType_doReturnFn) *MoqPostForm_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Resp *http.Response
			Err  error
		}
		Sequence   uint32
		DoFn       MoqPostForm_genType_doFn
		DoReturnFn MoqPostForm_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqPostForm_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqPostForm_genType_resultsByParams
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
		results = &MoqPostForm_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqPostForm_genType_paramsKey]*MoqPostForm_genType_results{},
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
		r.Results = &MoqPostForm_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqPostForm_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqPostForm_genType_fnRecorder {
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
					Resp *http.Response
					Err  error
				}
				Sequence   uint32
				DoFn       MoqPostForm_genType_doFn
				DoReturnFn MoqPostForm_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqPostForm_genType) PrettyParams(params MoqPostForm_genType_params) string {
	return fmt.Sprintf("PostForm_genType(%#v, %#v)", params.Url, params.Data)
}

func (m *MoqPostForm_genType) ParamsKey(params MoqPostForm_genType_params, anyParams uint64) MoqPostForm_genType_paramsKey {
	m.Scene.T.Helper()
	var urlUsed string
	var urlUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Url == moq.ParamIndexByValue {
			urlUsed = params.Url
		} else {
			urlUsedHash = hash.DeepHash(params.Url)
		}
	}
	var dataUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Data == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The data parameter can't be indexed by value")
		}
		dataUsedHash = hash.DeepHash(params.Data)
	}
	return MoqPostForm_genType_paramsKey{
		Params: struct{ Url string }{
			Url: urlUsed,
		},
		Hashes: struct {
			Url  hash.Hash
			Data hash.Hash
		}{
			Url:  urlUsedHash,
			Data: dataUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqPostForm_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqPostForm_genType) AssertExpectationsMet() {
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

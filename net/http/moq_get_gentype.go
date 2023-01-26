// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package http

import (
	"fmt"
	"math/bits"
	"net/http"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Get_genType is the fabricated implementation type of this mock (emitted when
// mocking functions directly and not from a function type)
type Get_genType func(url string) (resp *http.Response, err error)

// MoqGet_genType holds the state of a moq of the Get_genType type
type MoqGet_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqGet_genType_mock

	ResultsByParams []MoqGet_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Url moq.ParamIndexing
		}
	}
}

// MoqGet_genType_mock isolates the mock interface of the Get_genType type
type MoqGet_genType_mock struct {
	Moq *MoqGet_genType
}

// MoqGet_genType_params holds the params of the Get_genType type
type MoqGet_genType_params struct{ Url string }

// MoqGet_genType_paramsKey holds the map key params of the Get_genType type
type MoqGet_genType_paramsKey struct {
	Params struct{ Url string }
	Hashes struct{ Url hash.Hash }
}

// MoqGet_genType_resultsByParams contains the results for a given set of
// parameters for the Get_genType type
type MoqGet_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqGet_genType_paramsKey]*MoqGet_genType_results
}

// MoqGet_genType_doFn defines the type of function needed when calling AndDo
// for the Get_genType type
type MoqGet_genType_doFn func(url string)

// MoqGet_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Get_genType type
type MoqGet_genType_doReturnFn func(url string) (resp *http.Response, err error)

// MoqGet_genType_results holds the results of the Get_genType type
type MoqGet_genType_results struct {
	Params  MoqGet_genType_params
	Results []struct {
		Values *struct {
			Resp *http.Response
			Err  error
		}
		Sequence   uint32
		DoFn       MoqGet_genType_doFn
		DoReturnFn MoqGet_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqGet_genType_fnRecorder routes recorded function calls to the
// MoqGet_genType moq
type MoqGet_genType_fnRecorder struct {
	Params    MoqGet_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqGet_genType_results
	Moq       *MoqGet_genType
}

// MoqGet_genType_anyParams isolates the any params functions of the
// Get_genType type
type MoqGet_genType_anyParams struct {
	Recorder *MoqGet_genType_fnRecorder
}

// NewMoqGet_genType creates a new moq of the Get_genType type
func NewMoqGet_genType(scene *moq.Scene, config *moq.Config) *MoqGet_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqGet_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqGet_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Url moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Url moq.ParamIndexing
		}{
			Url: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Get_genType type
func (m *MoqGet_genType) Mock() Get_genType {
	return func(url string) (_ *http.Response, _ error) {
		m.Scene.T.Helper()
		moq := &MoqGet_genType_mock{Moq: m}
		return moq.Fn(url)
	}
}

func (m *MoqGet_genType_mock) Fn(url string) (resp *http.Response, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqGet_genType_params{
		Url: url,
	}
	var results *MoqGet_genType_results
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
		result.DoFn(url)
	}

	if result.Values != nil {
		resp = result.Values.Resp
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		resp, err = result.DoReturnFn(url)
	}
	return
}

func (m *MoqGet_genType) OnCall(url string) *MoqGet_genType_fnRecorder {
	return &MoqGet_genType_fnRecorder{
		Params: MoqGet_genType_params{
			Url: url,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqGet_genType_fnRecorder) Any() *MoqGet_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqGet_genType_anyParams{Recorder: r}
}

func (a *MoqGet_genType_anyParams) Url() *MoqGet_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqGet_genType_fnRecorder) Seq() *MoqGet_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqGet_genType_fnRecorder) NoSeq() *MoqGet_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqGet_genType_fnRecorder) ReturnResults(resp *http.Response, err error) *MoqGet_genType_fnRecorder {
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
		DoFn       MoqGet_genType_doFn
		DoReturnFn MoqGet_genType_doReturnFn
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

func (r *MoqGet_genType_fnRecorder) AndDo(fn MoqGet_genType_doFn) *MoqGet_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqGet_genType_fnRecorder) DoReturnResults(fn MoqGet_genType_doReturnFn) *MoqGet_genType_fnRecorder {
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
		DoFn       MoqGet_genType_doFn
		DoReturnFn MoqGet_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqGet_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqGet_genType_resultsByParams
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
		results = &MoqGet_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqGet_genType_paramsKey]*MoqGet_genType_results{},
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
		r.Results = &MoqGet_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqGet_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqGet_genType_fnRecorder {
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
				DoFn       MoqGet_genType_doFn
				DoReturnFn MoqGet_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqGet_genType) PrettyParams(params MoqGet_genType_params) string {
	return fmt.Sprintf("Get_genType(%#v)", params.Url)
}

func (m *MoqGet_genType) ParamsKey(params MoqGet_genType_params, anyParams uint64) MoqGet_genType_paramsKey {
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
	return MoqGet_genType_paramsKey{
		Params: struct{ Url string }{
			Url: urlUsed,
		},
		Hashes: struct{ Url hash.Hash }{
			Url: urlUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqGet_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqGet_genType) AssertExpectationsMet() {
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
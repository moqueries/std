// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package httptest

import (
	"fmt"
	"math/bits"
	"net/http"
	"net/http/httptest"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// NewUnstartedServer_genType is the fabricated implementation type of this
// mock (emitted when mocking functions directly and not from a function type)
type NewUnstartedServer_genType func(handler http.Handler) *httptest.Server

// MoqNewUnstartedServer_genType holds the state of a moq of the
// NewUnstartedServer_genType type
type MoqNewUnstartedServer_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNewUnstartedServer_genType_mock

	ResultsByParams []MoqNewUnstartedServer_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Handler moq.ParamIndexing
		}
	}
}

// MoqNewUnstartedServer_genType_mock isolates the mock interface of the
// NewUnstartedServer_genType type
type MoqNewUnstartedServer_genType_mock struct {
	Moq *MoqNewUnstartedServer_genType
}

// MoqNewUnstartedServer_genType_params holds the params of the
// NewUnstartedServer_genType type
type MoqNewUnstartedServer_genType_params struct{ Handler http.Handler }

// MoqNewUnstartedServer_genType_paramsKey holds the map key params of the
// NewUnstartedServer_genType type
type MoqNewUnstartedServer_genType_paramsKey struct {
	Params struct{ Handler http.Handler }
	Hashes struct{ Handler hash.Hash }
}

// MoqNewUnstartedServer_genType_resultsByParams contains the results for a
// given set of parameters for the NewUnstartedServer_genType type
type MoqNewUnstartedServer_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNewUnstartedServer_genType_paramsKey]*MoqNewUnstartedServer_genType_results
}

// MoqNewUnstartedServer_genType_doFn defines the type of function needed when
// calling AndDo for the NewUnstartedServer_genType type
type MoqNewUnstartedServer_genType_doFn func(handler http.Handler)

// MoqNewUnstartedServer_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the NewUnstartedServer_genType type
type MoqNewUnstartedServer_genType_doReturnFn func(handler http.Handler) *httptest.Server

// MoqNewUnstartedServer_genType_results holds the results of the
// NewUnstartedServer_genType type
type MoqNewUnstartedServer_genType_results struct {
	Params  MoqNewUnstartedServer_genType_params
	Results []struct {
		Values *struct {
			Result1 *httptest.Server
		}
		Sequence   uint32
		DoFn       MoqNewUnstartedServer_genType_doFn
		DoReturnFn MoqNewUnstartedServer_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNewUnstartedServer_genType_fnRecorder routes recorded function calls to
// the MoqNewUnstartedServer_genType moq
type MoqNewUnstartedServer_genType_fnRecorder struct {
	Params    MoqNewUnstartedServer_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNewUnstartedServer_genType_results
	Moq       *MoqNewUnstartedServer_genType
}

// MoqNewUnstartedServer_genType_anyParams isolates the any params functions of
// the NewUnstartedServer_genType type
type MoqNewUnstartedServer_genType_anyParams struct {
	Recorder *MoqNewUnstartedServer_genType_fnRecorder
}

// NewMoqNewUnstartedServer_genType creates a new moq of the
// NewUnstartedServer_genType type
func NewMoqNewUnstartedServer_genType(scene *moq.Scene, config *moq.Config) *MoqNewUnstartedServer_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNewUnstartedServer_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNewUnstartedServer_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Handler moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Handler moq.ParamIndexing
		}{
			Handler: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the NewUnstartedServer_genType type
func (m *MoqNewUnstartedServer_genType) Mock() NewUnstartedServer_genType {
	return func(handler http.Handler) *httptest.Server {
		m.Scene.T.Helper()
		moq := &MoqNewUnstartedServer_genType_mock{Moq: m}
		return moq.Fn(handler)
	}
}

func (m *MoqNewUnstartedServer_genType_mock) Fn(handler http.Handler) (result1 *httptest.Server) {
	m.Moq.Scene.T.Helper()
	params := MoqNewUnstartedServer_genType_params{
		Handler: handler,
	}
	var results *MoqNewUnstartedServer_genType_results
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
		result.DoFn(handler)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(handler)
	}
	return
}

func (m *MoqNewUnstartedServer_genType) OnCall(handler http.Handler) *MoqNewUnstartedServer_genType_fnRecorder {
	return &MoqNewUnstartedServer_genType_fnRecorder{
		Params: MoqNewUnstartedServer_genType_params{
			Handler: handler,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNewUnstartedServer_genType_fnRecorder) Any() *MoqNewUnstartedServer_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNewUnstartedServer_genType_anyParams{Recorder: r}
}

func (a *MoqNewUnstartedServer_genType_anyParams) Handler() *MoqNewUnstartedServer_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqNewUnstartedServer_genType_fnRecorder) Seq() *MoqNewUnstartedServer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNewUnstartedServer_genType_fnRecorder) NoSeq() *MoqNewUnstartedServer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNewUnstartedServer_genType_fnRecorder) ReturnResults(result1 *httptest.Server) *MoqNewUnstartedServer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *httptest.Server
		}
		Sequence   uint32
		DoFn       MoqNewUnstartedServer_genType_doFn
		DoReturnFn MoqNewUnstartedServer_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *httptest.Server
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNewUnstartedServer_genType_fnRecorder) AndDo(fn MoqNewUnstartedServer_genType_doFn) *MoqNewUnstartedServer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNewUnstartedServer_genType_fnRecorder) DoReturnResults(fn MoqNewUnstartedServer_genType_doReturnFn) *MoqNewUnstartedServer_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *httptest.Server
		}
		Sequence   uint32
		DoFn       MoqNewUnstartedServer_genType_doFn
		DoReturnFn MoqNewUnstartedServer_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNewUnstartedServer_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNewUnstartedServer_genType_resultsByParams
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
		results = &MoqNewUnstartedServer_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNewUnstartedServer_genType_paramsKey]*MoqNewUnstartedServer_genType_results{},
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
		r.Results = &MoqNewUnstartedServer_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNewUnstartedServer_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNewUnstartedServer_genType_fnRecorder {
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
					Result1 *httptest.Server
				}
				Sequence   uint32
				DoFn       MoqNewUnstartedServer_genType_doFn
				DoReturnFn MoqNewUnstartedServer_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNewUnstartedServer_genType) PrettyParams(params MoqNewUnstartedServer_genType_params) string {
	return fmt.Sprintf("NewUnstartedServer_genType(%#v)", params.Handler)
}

func (m *MoqNewUnstartedServer_genType) ParamsKey(params MoqNewUnstartedServer_genType_params, anyParams uint64) MoqNewUnstartedServer_genType_paramsKey {
	m.Scene.T.Helper()
	var handlerUsed http.Handler
	var handlerUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Handler == moq.ParamIndexByValue {
			handlerUsed = params.Handler
		} else {
			handlerUsedHash = hash.DeepHash(params.Handler)
		}
	}
	return MoqNewUnstartedServer_genType_paramsKey{
		Params: struct{ Handler http.Handler }{
			Handler: handlerUsed,
		},
		Hashes: struct{ Handler hash.Hash }{
			Handler: handlerUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNewUnstartedServer_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNewUnstartedServer_genType) AssertExpectationsMet() {
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
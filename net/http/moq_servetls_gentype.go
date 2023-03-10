// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package http

import (
	"fmt"
	"math/bits"
	"net"
	"net/http"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// ServeTLS_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type ServeTLS_genType func(l net.Listener, handler http.Handler, certFile, keyFile string) error

// MoqServeTLS_genType holds the state of a moq of the ServeTLS_genType type
type MoqServeTLS_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqServeTLS_genType_mock

	ResultsByParams []MoqServeTLS_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			L        moq.ParamIndexing
			Handler  moq.ParamIndexing
			CertFile moq.ParamIndexing
			KeyFile  moq.ParamIndexing
		}
	}
}

// MoqServeTLS_genType_mock isolates the mock interface of the ServeTLS_genType
// type
type MoqServeTLS_genType_mock struct {
	Moq *MoqServeTLS_genType
}

// MoqServeTLS_genType_params holds the params of the ServeTLS_genType type
type MoqServeTLS_genType_params struct {
	L                 net.Listener
	Handler           http.Handler
	CertFile, KeyFile string
}

// MoqServeTLS_genType_paramsKey holds the map key params of the
// ServeTLS_genType type
type MoqServeTLS_genType_paramsKey struct {
	Params struct {
		L                 net.Listener
		Handler           http.Handler
		CertFile, KeyFile string
	}
	Hashes struct {
		L                 hash.Hash
		Handler           hash.Hash
		CertFile, KeyFile hash.Hash
	}
}

// MoqServeTLS_genType_resultsByParams contains the results for a given set of
// parameters for the ServeTLS_genType type
type MoqServeTLS_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqServeTLS_genType_paramsKey]*MoqServeTLS_genType_results
}

// MoqServeTLS_genType_doFn defines the type of function needed when calling
// AndDo for the ServeTLS_genType type
type MoqServeTLS_genType_doFn func(l net.Listener, handler http.Handler, certFile, keyFile string)

// MoqServeTLS_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ServeTLS_genType type
type MoqServeTLS_genType_doReturnFn func(l net.Listener, handler http.Handler, certFile, keyFile string) error

// MoqServeTLS_genType_results holds the results of the ServeTLS_genType type
type MoqServeTLS_genType_results struct {
	Params  MoqServeTLS_genType_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqServeTLS_genType_doFn
		DoReturnFn MoqServeTLS_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqServeTLS_genType_fnRecorder routes recorded function calls to the
// MoqServeTLS_genType moq
type MoqServeTLS_genType_fnRecorder struct {
	Params    MoqServeTLS_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqServeTLS_genType_results
	Moq       *MoqServeTLS_genType
}

// MoqServeTLS_genType_anyParams isolates the any params functions of the
// ServeTLS_genType type
type MoqServeTLS_genType_anyParams struct {
	Recorder *MoqServeTLS_genType_fnRecorder
}

// NewMoqServeTLS_genType creates a new moq of the ServeTLS_genType type
func NewMoqServeTLS_genType(scene *moq.Scene, config *moq.Config) *MoqServeTLS_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqServeTLS_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqServeTLS_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				L        moq.ParamIndexing
				Handler  moq.ParamIndexing
				CertFile moq.ParamIndexing
				KeyFile  moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			L        moq.ParamIndexing
			Handler  moq.ParamIndexing
			CertFile moq.ParamIndexing
			KeyFile  moq.ParamIndexing
		}{
			L:        moq.ParamIndexByHash,
			Handler:  moq.ParamIndexByHash,
			CertFile: moq.ParamIndexByValue,
			KeyFile:  moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ServeTLS_genType type
func (m *MoqServeTLS_genType) Mock() ServeTLS_genType {
	return func(l net.Listener, handler http.Handler, certFile, keyFile string) error {
		m.Scene.T.Helper()
		moq := &MoqServeTLS_genType_mock{Moq: m}
		return moq.Fn(l, handler, certFile, keyFile)
	}
}

func (m *MoqServeTLS_genType_mock) Fn(l net.Listener, handler http.Handler, certFile, keyFile string) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqServeTLS_genType_params{
		L:        l,
		Handler:  handler,
		CertFile: certFile,
		KeyFile:  keyFile,
	}
	var results *MoqServeTLS_genType_results
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
		result.DoFn(l, handler, certFile, keyFile)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(l, handler, certFile, keyFile)
	}
	return
}

func (m *MoqServeTLS_genType) OnCall(l net.Listener, handler http.Handler, certFile, keyFile string) *MoqServeTLS_genType_fnRecorder {
	return &MoqServeTLS_genType_fnRecorder{
		Params: MoqServeTLS_genType_params{
			L:        l,
			Handler:  handler,
			CertFile: certFile,
			KeyFile:  keyFile,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqServeTLS_genType_fnRecorder) Any() *MoqServeTLS_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqServeTLS_genType_anyParams{Recorder: r}
}

func (a *MoqServeTLS_genType_anyParams) L() *MoqServeTLS_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqServeTLS_genType_anyParams) Handler() *MoqServeTLS_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqServeTLS_genType_anyParams) CertFile() *MoqServeTLS_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (a *MoqServeTLS_genType_anyParams) KeyFile() *MoqServeTLS_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 3
	return a.Recorder
}

func (r *MoqServeTLS_genType_fnRecorder) Seq() *MoqServeTLS_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqServeTLS_genType_fnRecorder) NoSeq() *MoqServeTLS_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqServeTLS_genType_fnRecorder) ReturnResults(result1 error) *MoqServeTLS_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqServeTLS_genType_doFn
		DoReturnFn MoqServeTLS_genType_doReturnFn
	}{
		Values: &struct {
			Result1 error
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqServeTLS_genType_fnRecorder) AndDo(fn MoqServeTLS_genType_doFn) *MoqServeTLS_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqServeTLS_genType_fnRecorder) DoReturnResults(fn MoqServeTLS_genType_doReturnFn) *MoqServeTLS_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqServeTLS_genType_doFn
		DoReturnFn MoqServeTLS_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqServeTLS_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqServeTLS_genType_resultsByParams
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
		results = &MoqServeTLS_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqServeTLS_genType_paramsKey]*MoqServeTLS_genType_results{},
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
		r.Results = &MoqServeTLS_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqServeTLS_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqServeTLS_genType_fnRecorder {
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
					Result1 error
				}
				Sequence   uint32
				DoFn       MoqServeTLS_genType_doFn
				DoReturnFn MoqServeTLS_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqServeTLS_genType) PrettyParams(params MoqServeTLS_genType_params) string {
	return fmt.Sprintf("ServeTLS_genType(%#v, %#v, %#v, %#v)", params.L, params.Handler, params.CertFile, params.KeyFile)
}

func (m *MoqServeTLS_genType) ParamsKey(params MoqServeTLS_genType_params, anyParams uint64) MoqServeTLS_genType_paramsKey {
	m.Scene.T.Helper()
	var lUsed net.Listener
	var lUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.L == moq.ParamIndexByValue {
			lUsed = params.L
		} else {
			lUsedHash = hash.DeepHash(params.L)
		}
	}
	var handlerUsed http.Handler
	var handlerUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Handler == moq.ParamIndexByValue {
			handlerUsed = params.Handler
		} else {
			handlerUsedHash = hash.DeepHash(params.Handler)
		}
	}
	var certFileUsed string
	var certFileUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.CertFile == moq.ParamIndexByValue {
			certFileUsed = params.CertFile
		} else {
			certFileUsedHash = hash.DeepHash(params.CertFile)
		}
	}
	var keyFileUsed string
	var keyFileUsedHash hash.Hash
	if anyParams&(1<<3) == 0 {
		if m.Runtime.ParameterIndexing.KeyFile == moq.ParamIndexByValue {
			keyFileUsed = params.KeyFile
		} else {
			keyFileUsedHash = hash.DeepHash(params.KeyFile)
		}
	}
	return MoqServeTLS_genType_paramsKey{
		Params: struct {
			L                 net.Listener
			Handler           http.Handler
			CertFile, KeyFile string
		}{
			L:        lUsed,
			Handler:  handlerUsed,
			CertFile: certFileUsed,
			KeyFile:  keyFileUsed,
		},
		Hashes: struct {
			L                 hash.Hash
			Handler           hash.Hash
			CertFile, KeyFile hash.Hash
		}{
			L:        lUsedHash,
			Handler:  handlerUsedHash,
			CertFile: certFileUsedHash,
			KeyFile:  keyFileUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqServeTLS_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqServeTLS_genType) AssertExpectationsMet() {
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

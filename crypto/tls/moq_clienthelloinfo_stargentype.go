// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package tls

import (
	"context"
	"crypto/tls"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that tls.ClientHelloInfo_starGenType is
// mocked completely
var _ ClientHelloInfo_starGenType = (*MoqClientHelloInfo_starGenType_mock)(nil)

// ClientHelloInfo_starGenType is the fabricated implementation type of this
// mock (emitted when mocking a collections of methods directly and not from an
// interface type)
type ClientHelloInfo_starGenType interface {
	Context() context.Context
	SupportsCertificate(c *tls.Certificate) error
}

// MoqClientHelloInfo_starGenType holds the state of a moq of the
// ClientHelloInfo_starGenType type
type MoqClientHelloInfo_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqClientHelloInfo_starGenType_mock

	ResultsByParams_Context             []MoqClientHelloInfo_starGenType_Context_resultsByParams
	ResultsByParams_SupportsCertificate []MoqClientHelloInfo_starGenType_SupportsCertificate_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Context             struct{}
			SupportsCertificate struct {
				C moq.ParamIndexing
			}
		}
	}
	// MoqClientHelloInfo_starGenType_mock isolates the mock interface of the
}

// ClientHelloInfo_starGenType type
type MoqClientHelloInfo_starGenType_mock struct {
	Moq *MoqClientHelloInfo_starGenType
}

// MoqClientHelloInfo_starGenType_recorder isolates the recorder interface of
// the ClientHelloInfo_starGenType type
type MoqClientHelloInfo_starGenType_recorder struct {
	Moq *MoqClientHelloInfo_starGenType
}

// MoqClientHelloInfo_starGenType_Context_params holds the params of the
// ClientHelloInfo_starGenType type
type MoqClientHelloInfo_starGenType_Context_params struct{}

// MoqClientHelloInfo_starGenType_Context_paramsKey holds the map key params of
// the ClientHelloInfo_starGenType type
type MoqClientHelloInfo_starGenType_Context_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqClientHelloInfo_starGenType_Context_resultsByParams contains the results
// for a given set of parameters for the ClientHelloInfo_starGenType type
type MoqClientHelloInfo_starGenType_Context_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqClientHelloInfo_starGenType_Context_paramsKey]*MoqClientHelloInfo_starGenType_Context_results
}

// MoqClientHelloInfo_starGenType_Context_doFn defines the type of function
// needed when calling AndDo for the ClientHelloInfo_starGenType type
type MoqClientHelloInfo_starGenType_Context_doFn func()

// MoqClientHelloInfo_starGenType_Context_doReturnFn defines the type of
// function needed when calling DoReturnResults for the
// ClientHelloInfo_starGenType type
type MoqClientHelloInfo_starGenType_Context_doReturnFn func() context.Context

// MoqClientHelloInfo_starGenType_Context_results holds the results of the
// ClientHelloInfo_starGenType type
type MoqClientHelloInfo_starGenType_Context_results struct {
	Params  MoqClientHelloInfo_starGenType_Context_params
	Results []struct {
		Values *struct {
			Result1 context.Context
		}
		Sequence   uint32
		DoFn       MoqClientHelloInfo_starGenType_Context_doFn
		DoReturnFn MoqClientHelloInfo_starGenType_Context_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqClientHelloInfo_starGenType_Context_fnRecorder routes recorded function
// calls to the MoqClientHelloInfo_starGenType moq
type MoqClientHelloInfo_starGenType_Context_fnRecorder struct {
	Params    MoqClientHelloInfo_starGenType_Context_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqClientHelloInfo_starGenType_Context_results
	Moq       *MoqClientHelloInfo_starGenType
}

// MoqClientHelloInfo_starGenType_Context_anyParams isolates the any params
// functions of the ClientHelloInfo_starGenType type
type MoqClientHelloInfo_starGenType_Context_anyParams struct {
	Recorder *MoqClientHelloInfo_starGenType_Context_fnRecorder
}

// MoqClientHelloInfo_starGenType_SupportsCertificate_params holds the params
// of the ClientHelloInfo_starGenType type
type MoqClientHelloInfo_starGenType_SupportsCertificate_params struct{ C *tls.Certificate }

// MoqClientHelloInfo_starGenType_SupportsCertificate_paramsKey holds the map
// key params of the ClientHelloInfo_starGenType type
type MoqClientHelloInfo_starGenType_SupportsCertificate_paramsKey struct {
	Params struct{ C *tls.Certificate }
	Hashes struct{ C hash.Hash }
}

// MoqClientHelloInfo_starGenType_SupportsCertificate_resultsByParams contains
// the results for a given set of parameters for the
// ClientHelloInfo_starGenType type
type MoqClientHelloInfo_starGenType_SupportsCertificate_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqClientHelloInfo_starGenType_SupportsCertificate_paramsKey]*MoqClientHelloInfo_starGenType_SupportsCertificate_results
}

// MoqClientHelloInfo_starGenType_SupportsCertificate_doFn defines the type of
// function needed when calling AndDo for the ClientHelloInfo_starGenType type
type MoqClientHelloInfo_starGenType_SupportsCertificate_doFn func(c *tls.Certificate)

// MoqClientHelloInfo_starGenType_SupportsCertificate_doReturnFn defines the
// type of function needed when calling DoReturnResults for the
// ClientHelloInfo_starGenType type
type MoqClientHelloInfo_starGenType_SupportsCertificate_doReturnFn func(c *tls.Certificate) error

// MoqClientHelloInfo_starGenType_SupportsCertificate_results holds the results
// of the ClientHelloInfo_starGenType type
type MoqClientHelloInfo_starGenType_SupportsCertificate_results struct {
	Params  MoqClientHelloInfo_starGenType_SupportsCertificate_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqClientHelloInfo_starGenType_SupportsCertificate_doFn
		DoReturnFn MoqClientHelloInfo_starGenType_SupportsCertificate_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqClientHelloInfo_starGenType_SupportsCertificate_fnRecorder routes
// recorded function calls to the MoqClientHelloInfo_starGenType moq
type MoqClientHelloInfo_starGenType_SupportsCertificate_fnRecorder struct {
	Params    MoqClientHelloInfo_starGenType_SupportsCertificate_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqClientHelloInfo_starGenType_SupportsCertificate_results
	Moq       *MoqClientHelloInfo_starGenType
}

// MoqClientHelloInfo_starGenType_SupportsCertificate_anyParams isolates the
// any params functions of the ClientHelloInfo_starGenType type
type MoqClientHelloInfo_starGenType_SupportsCertificate_anyParams struct {
	Recorder *MoqClientHelloInfo_starGenType_SupportsCertificate_fnRecorder
}

// NewMoqClientHelloInfo_starGenType creates a new moq of the
// ClientHelloInfo_starGenType type
func NewMoqClientHelloInfo_starGenType(scene *moq.Scene, config *moq.Config) *MoqClientHelloInfo_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqClientHelloInfo_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqClientHelloInfo_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Context             struct{}
				SupportsCertificate struct {
					C moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Context             struct{}
			SupportsCertificate struct {
				C moq.ParamIndexing
			}
		}{
			Context: struct{}{},
			SupportsCertificate: struct {
				C moq.ParamIndexing
			}{
				C: moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the ClientHelloInfo_starGenType type
func (m *MoqClientHelloInfo_starGenType) Mock() *MoqClientHelloInfo_starGenType_mock { return m.Moq }

func (m *MoqClientHelloInfo_starGenType_mock) Context() (result1 context.Context) {
	m.Moq.Scene.T.Helper()
	params := MoqClientHelloInfo_starGenType_Context_params{}
	var results *MoqClientHelloInfo_starGenType_Context_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Context {
		paramsKey := m.Moq.ParamsKey_Context(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Context(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Context(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Context(params))
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

func (m *MoqClientHelloInfo_starGenType_mock) SupportsCertificate(c *tls.Certificate) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqClientHelloInfo_starGenType_SupportsCertificate_params{
		C: c,
	}
	var results *MoqClientHelloInfo_starGenType_SupportsCertificate_results
	for _, resultsByParams := range m.Moq.ResultsByParams_SupportsCertificate {
		paramsKey := m.Moq.ParamsKey_SupportsCertificate(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_SupportsCertificate(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_SupportsCertificate(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_SupportsCertificate(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(c)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(c)
	}
	return
}

// OnCall returns the recorder implementation of the
// ClientHelloInfo_starGenType type
func (m *MoqClientHelloInfo_starGenType) OnCall() *MoqClientHelloInfo_starGenType_recorder {
	return &MoqClientHelloInfo_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqClientHelloInfo_starGenType_recorder) Context() *MoqClientHelloInfo_starGenType_Context_fnRecorder {
	return &MoqClientHelloInfo_starGenType_Context_fnRecorder{
		Params:   MoqClientHelloInfo_starGenType_Context_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqClientHelloInfo_starGenType_Context_fnRecorder) Any() *MoqClientHelloInfo_starGenType_Context_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Context(r.Params))
		return nil
	}
	return &MoqClientHelloInfo_starGenType_Context_anyParams{Recorder: r}
}

func (r *MoqClientHelloInfo_starGenType_Context_fnRecorder) Seq() *MoqClientHelloInfo_starGenType_Context_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Context(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqClientHelloInfo_starGenType_Context_fnRecorder) NoSeq() *MoqClientHelloInfo_starGenType_Context_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Context(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqClientHelloInfo_starGenType_Context_fnRecorder) ReturnResults(result1 context.Context) *MoqClientHelloInfo_starGenType_Context_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 context.Context
		}
		Sequence   uint32
		DoFn       MoqClientHelloInfo_starGenType_Context_doFn
		DoReturnFn MoqClientHelloInfo_starGenType_Context_doReturnFn
	}{
		Values: &struct {
			Result1 context.Context
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqClientHelloInfo_starGenType_Context_fnRecorder) AndDo(fn MoqClientHelloInfo_starGenType_Context_doFn) *MoqClientHelloInfo_starGenType_Context_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqClientHelloInfo_starGenType_Context_fnRecorder) DoReturnResults(fn MoqClientHelloInfo_starGenType_Context_doReturnFn) *MoqClientHelloInfo_starGenType_Context_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 context.Context
		}
		Sequence   uint32
		DoFn       MoqClientHelloInfo_starGenType_Context_doFn
		DoReturnFn MoqClientHelloInfo_starGenType_Context_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqClientHelloInfo_starGenType_Context_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqClientHelloInfo_starGenType_Context_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Context {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqClientHelloInfo_starGenType_Context_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqClientHelloInfo_starGenType_Context_paramsKey]*MoqClientHelloInfo_starGenType_Context_results{},
		}
		r.Moq.ResultsByParams_Context = append(r.Moq.ResultsByParams_Context, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Context) {
			copy(r.Moq.ResultsByParams_Context[insertAt+1:], r.Moq.ResultsByParams_Context[insertAt:0])
			r.Moq.ResultsByParams_Context[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Context(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqClientHelloInfo_starGenType_Context_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqClientHelloInfo_starGenType_Context_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqClientHelloInfo_starGenType_Context_fnRecorder {
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
					Result1 context.Context
				}
				Sequence   uint32
				DoFn       MoqClientHelloInfo_starGenType_Context_doFn
				DoReturnFn MoqClientHelloInfo_starGenType_Context_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqClientHelloInfo_starGenType) PrettyParams_Context(params MoqClientHelloInfo_starGenType_Context_params) string {
	return fmt.Sprintf("Context()")
}

func (m *MoqClientHelloInfo_starGenType) ParamsKey_Context(params MoqClientHelloInfo_starGenType_Context_params, anyParams uint64) MoqClientHelloInfo_starGenType_Context_paramsKey {
	m.Scene.T.Helper()
	return MoqClientHelloInfo_starGenType_Context_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqClientHelloInfo_starGenType_recorder) SupportsCertificate(c *tls.Certificate) *MoqClientHelloInfo_starGenType_SupportsCertificate_fnRecorder {
	return &MoqClientHelloInfo_starGenType_SupportsCertificate_fnRecorder{
		Params: MoqClientHelloInfo_starGenType_SupportsCertificate_params{
			C: c,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqClientHelloInfo_starGenType_SupportsCertificate_fnRecorder) Any() *MoqClientHelloInfo_starGenType_SupportsCertificate_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_SupportsCertificate(r.Params))
		return nil
	}
	return &MoqClientHelloInfo_starGenType_SupportsCertificate_anyParams{Recorder: r}
}

func (a *MoqClientHelloInfo_starGenType_SupportsCertificate_anyParams) C() *MoqClientHelloInfo_starGenType_SupportsCertificate_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqClientHelloInfo_starGenType_SupportsCertificate_fnRecorder) Seq() *MoqClientHelloInfo_starGenType_SupportsCertificate_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_SupportsCertificate(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqClientHelloInfo_starGenType_SupportsCertificate_fnRecorder) NoSeq() *MoqClientHelloInfo_starGenType_SupportsCertificate_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_SupportsCertificate(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqClientHelloInfo_starGenType_SupportsCertificate_fnRecorder) ReturnResults(result1 error) *MoqClientHelloInfo_starGenType_SupportsCertificate_fnRecorder {
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
		DoFn       MoqClientHelloInfo_starGenType_SupportsCertificate_doFn
		DoReturnFn MoqClientHelloInfo_starGenType_SupportsCertificate_doReturnFn
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

func (r *MoqClientHelloInfo_starGenType_SupportsCertificate_fnRecorder) AndDo(fn MoqClientHelloInfo_starGenType_SupportsCertificate_doFn) *MoqClientHelloInfo_starGenType_SupportsCertificate_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqClientHelloInfo_starGenType_SupportsCertificate_fnRecorder) DoReturnResults(fn MoqClientHelloInfo_starGenType_SupportsCertificate_doReturnFn) *MoqClientHelloInfo_starGenType_SupportsCertificate_fnRecorder {
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
		DoFn       MoqClientHelloInfo_starGenType_SupportsCertificate_doFn
		DoReturnFn MoqClientHelloInfo_starGenType_SupportsCertificate_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqClientHelloInfo_starGenType_SupportsCertificate_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqClientHelloInfo_starGenType_SupportsCertificate_resultsByParams
	for n, res := range r.Moq.ResultsByParams_SupportsCertificate {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqClientHelloInfo_starGenType_SupportsCertificate_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqClientHelloInfo_starGenType_SupportsCertificate_paramsKey]*MoqClientHelloInfo_starGenType_SupportsCertificate_results{},
		}
		r.Moq.ResultsByParams_SupportsCertificate = append(r.Moq.ResultsByParams_SupportsCertificate, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_SupportsCertificate) {
			copy(r.Moq.ResultsByParams_SupportsCertificate[insertAt+1:], r.Moq.ResultsByParams_SupportsCertificate[insertAt:0])
			r.Moq.ResultsByParams_SupportsCertificate[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_SupportsCertificate(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqClientHelloInfo_starGenType_SupportsCertificate_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqClientHelloInfo_starGenType_SupportsCertificate_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqClientHelloInfo_starGenType_SupportsCertificate_fnRecorder {
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
				DoFn       MoqClientHelloInfo_starGenType_SupportsCertificate_doFn
				DoReturnFn MoqClientHelloInfo_starGenType_SupportsCertificate_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqClientHelloInfo_starGenType) PrettyParams_SupportsCertificate(params MoqClientHelloInfo_starGenType_SupportsCertificate_params) string {
	return fmt.Sprintf("SupportsCertificate(%#v)", params.C)
}

func (m *MoqClientHelloInfo_starGenType) ParamsKey_SupportsCertificate(params MoqClientHelloInfo_starGenType_SupportsCertificate_params, anyParams uint64) MoqClientHelloInfo_starGenType_SupportsCertificate_paramsKey {
	m.Scene.T.Helper()
	var cUsed *tls.Certificate
	var cUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.SupportsCertificate.C == moq.ParamIndexByValue {
			cUsed = params.C
		} else {
			cUsedHash = hash.DeepHash(params.C)
		}
	}
	return MoqClientHelloInfo_starGenType_SupportsCertificate_paramsKey{
		Params: struct{ C *tls.Certificate }{
			C: cUsed,
		},
		Hashes: struct{ C hash.Hash }{
			C: cUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqClientHelloInfo_starGenType) Reset() {
	m.ResultsByParams_Context = nil
	m.ResultsByParams_SupportsCertificate = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqClientHelloInfo_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Context {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Context(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_SupportsCertificate {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_SupportsCertificate(results.Params))
			}
		}
	}
}

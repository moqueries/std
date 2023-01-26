// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package x509

import (
	"crypto/x509"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// ParseCertificate_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type ParseCertificate_genType func(asn1Data []byte) (*x509.Certificate, error)

// MoqParseCertificate_genType holds the state of a moq of the
// ParseCertificate_genType type
type MoqParseCertificate_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqParseCertificate_genType_mock

	ResultsByParams []MoqParseCertificate_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Asn1Data moq.ParamIndexing
		}
	}
}

// MoqParseCertificate_genType_mock isolates the mock interface of the
// ParseCertificate_genType type
type MoqParseCertificate_genType_mock struct {
	Moq *MoqParseCertificate_genType
}

// MoqParseCertificate_genType_params holds the params of the
// ParseCertificate_genType type
type MoqParseCertificate_genType_params struct{ Asn1Data []byte }

// MoqParseCertificate_genType_paramsKey holds the map key params of the
// ParseCertificate_genType type
type MoqParseCertificate_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ Asn1Data hash.Hash }
}

// MoqParseCertificate_genType_resultsByParams contains the results for a given
// set of parameters for the ParseCertificate_genType type
type MoqParseCertificate_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqParseCertificate_genType_paramsKey]*MoqParseCertificate_genType_results
}

// MoqParseCertificate_genType_doFn defines the type of function needed when
// calling AndDo for the ParseCertificate_genType type
type MoqParseCertificate_genType_doFn func(asn1Data []byte)

// MoqParseCertificate_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the ParseCertificate_genType type
type MoqParseCertificate_genType_doReturnFn func(asn1Data []byte) (*x509.Certificate, error)

// MoqParseCertificate_genType_results holds the results of the
// ParseCertificate_genType type
type MoqParseCertificate_genType_results struct {
	Params  MoqParseCertificate_genType_params
	Results []struct {
		Values *struct {
			Result1 *x509.Certificate
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqParseCertificate_genType_doFn
		DoReturnFn MoqParseCertificate_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqParseCertificate_genType_fnRecorder routes recorded function calls to the
// MoqParseCertificate_genType moq
type MoqParseCertificate_genType_fnRecorder struct {
	Params    MoqParseCertificate_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqParseCertificate_genType_results
	Moq       *MoqParseCertificate_genType
}

// MoqParseCertificate_genType_anyParams isolates the any params functions of
// the ParseCertificate_genType type
type MoqParseCertificate_genType_anyParams struct {
	Recorder *MoqParseCertificate_genType_fnRecorder
}

// NewMoqParseCertificate_genType creates a new moq of the
// ParseCertificate_genType type
func NewMoqParseCertificate_genType(scene *moq.Scene, config *moq.Config) *MoqParseCertificate_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqParseCertificate_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqParseCertificate_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Asn1Data moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Asn1Data moq.ParamIndexing
		}{
			Asn1Data: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ParseCertificate_genType type
func (m *MoqParseCertificate_genType) Mock() ParseCertificate_genType {
	return func(asn1Data []byte) (*x509.Certificate, error) {
		m.Scene.T.Helper()
		moq := &MoqParseCertificate_genType_mock{Moq: m}
		return moq.Fn(asn1Data)
	}
}

func (m *MoqParseCertificate_genType_mock) Fn(asn1Data []byte) (result1 *x509.Certificate, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqParseCertificate_genType_params{
		Asn1Data: asn1Data,
	}
	var results *MoqParseCertificate_genType_results
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
		result.DoFn(asn1Data)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(asn1Data)
	}
	return
}

func (m *MoqParseCertificate_genType) OnCall(asn1Data []byte) *MoqParseCertificate_genType_fnRecorder {
	return &MoqParseCertificate_genType_fnRecorder{
		Params: MoqParseCertificate_genType_params{
			Asn1Data: asn1Data,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqParseCertificate_genType_fnRecorder) Any() *MoqParseCertificate_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqParseCertificate_genType_anyParams{Recorder: r}
}

func (a *MoqParseCertificate_genType_anyParams) Asn1Data() *MoqParseCertificate_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqParseCertificate_genType_fnRecorder) Seq() *MoqParseCertificate_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqParseCertificate_genType_fnRecorder) NoSeq() *MoqParseCertificate_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqParseCertificate_genType_fnRecorder) ReturnResults(result1 *x509.Certificate, result2 error) *MoqParseCertificate_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *x509.Certificate
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqParseCertificate_genType_doFn
		DoReturnFn MoqParseCertificate_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *x509.Certificate
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqParseCertificate_genType_fnRecorder) AndDo(fn MoqParseCertificate_genType_doFn) *MoqParseCertificate_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqParseCertificate_genType_fnRecorder) DoReturnResults(fn MoqParseCertificate_genType_doReturnFn) *MoqParseCertificate_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *x509.Certificate
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqParseCertificate_genType_doFn
		DoReturnFn MoqParseCertificate_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqParseCertificate_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqParseCertificate_genType_resultsByParams
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
		results = &MoqParseCertificate_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqParseCertificate_genType_paramsKey]*MoqParseCertificate_genType_results{},
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
		r.Results = &MoqParseCertificate_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqParseCertificate_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqParseCertificate_genType_fnRecorder {
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
					Result1 *x509.Certificate
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqParseCertificate_genType_doFn
				DoReturnFn MoqParseCertificate_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqParseCertificate_genType) PrettyParams(params MoqParseCertificate_genType_params) string {
	return fmt.Sprintf("ParseCertificate_genType(%#v)", params.Asn1Data)
}

func (m *MoqParseCertificate_genType) ParamsKey(params MoqParseCertificate_genType_params, anyParams uint64) MoqParseCertificate_genType_paramsKey {
	m.Scene.T.Helper()
	var asn1DataUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Asn1Data == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The asn1Data parameter can't be indexed by value")
		}
		asn1DataUsedHash = hash.DeepHash(params.Asn1Data)
	}
	return MoqParseCertificate_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ Asn1Data hash.Hash }{
			Asn1Data: asn1DataUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqParseCertificate_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqParseCertificate_genType) AssertExpectationsMet() {
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

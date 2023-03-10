// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package x509

import (
	"crypto"
	"crypto/x509"
	"fmt"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// CreateRevocationList_genType is the fabricated implementation type of this
// mock (emitted when mocking functions directly and not from a function type)
type CreateRevocationList_genType func(rand io.Reader, template *x509.RevocationList, issuer *x509.Certificate, priv crypto.Signer) ([]byte, error)

// MoqCreateRevocationList_genType holds the state of a moq of the
// CreateRevocationList_genType type
type MoqCreateRevocationList_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqCreateRevocationList_genType_mock

	ResultsByParams []MoqCreateRevocationList_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Rand     moq.ParamIndexing
			Template moq.ParamIndexing
			Issuer   moq.ParamIndexing
			Priv     moq.ParamIndexing
		}
	}
}

// MoqCreateRevocationList_genType_mock isolates the mock interface of the
// CreateRevocationList_genType type
type MoqCreateRevocationList_genType_mock struct {
	Moq *MoqCreateRevocationList_genType
}

// MoqCreateRevocationList_genType_params holds the params of the
// CreateRevocationList_genType type
type MoqCreateRevocationList_genType_params struct {
	Rand     io.Reader
	Template *x509.RevocationList
	Issuer   *x509.Certificate
	Priv     crypto.Signer
}

// MoqCreateRevocationList_genType_paramsKey holds the map key params of the
// CreateRevocationList_genType type
type MoqCreateRevocationList_genType_paramsKey struct {
	Params struct {
		Rand     io.Reader
		Template *x509.RevocationList
		Issuer   *x509.Certificate
		Priv     crypto.Signer
	}
	Hashes struct {
		Rand     hash.Hash
		Template hash.Hash
		Issuer   hash.Hash
		Priv     hash.Hash
	}
}

// MoqCreateRevocationList_genType_resultsByParams contains the results for a
// given set of parameters for the CreateRevocationList_genType type
type MoqCreateRevocationList_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqCreateRevocationList_genType_paramsKey]*MoqCreateRevocationList_genType_results
}

// MoqCreateRevocationList_genType_doFn defines the type of function needed
// when calling AndDo for the CreateRevocationList_genType type
type MoqCreateRevocationList_genType_doFn func(rand io.Reader, template *x509.RevocationList, issuer *x509.Certificate, priv crypto.Signer)

// MoqCreateRevocationList_genType_doReturnFn defines the type of function
// needed when calling DoReturnResults for the CreateRevocationList_genType
// type
type MoqCreateRevocationList_genType_doReturnFn func(rand io.Reader, template *x509.RevocationList, issuer *x509.Certificate, priv crypto.Signer) ([]byte, error)

// MoqCreateRevocationList_genType_results holds the results of the
// CreateRevocationList_genType type
type MoqCreateRevocationList_genType_results struct {
	Params  MoqCreateRevocationList_genType_params
	Results []struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqCreateRevocationList_genType_doFn
		DoReturnFn MoqCreateRevocationList_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqCreateRevocationList_genType_fnRecorder routes recorded function calls to
// the MoqCreateRevocationList_genType moq
type MoqCreateRevocationList_genType_fnRecorder struct {
	Params    MoqCreateRevocationList_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqCreateRevocationList_genType_results
	Moq       *MoqCreateRevocationList_genType
}

// MoqCreateRevocationList_genType_anyParams isolates the any params functions
// of the CreateRevocationList_genType type
type MoqCreateRevocationList_genType_anyParams struct {
	Recorder *MoqCreateRevocationList_genType_fnRecorder
}

// NewMoqCreateRevocationList_genType creates a new moq of the
// CreateRevocationList_genType type
func NewMoqCreateRevocationList_genType(scene *moq.Scene, config *moq.Config) *MoqCreateRevocationList_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqCreateRevocationList_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqCreateRevocationList_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Rand     moq.ParamIndexing
				Template moq.ParamIndexing
				Issuer   moq.ParamIndexing
				Priv     moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Rand     moq.ParamIndexing
			Template moq.ParamIndexing
			Issuer   moq.ParamIndexing
			Priv     moq.ParamIndexing
		}{
			Rand:     moq.ParamIndexByHash,
			Template: moq.ParamIndexByHash,
			Issuer:   moq.ParamIndexByHash,
			Priv:     moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the CreateRevocationList_genType type
func (m *MoqCreateRevocationList_genType) Mock() CreateRevocationList_genType {
	return func(rand io.Reader, template *x509.RevocationList, issuer *x509.Certificate, priv crypto.Signer) ([]byte, error) {
		m.Scene.T.Helper()
		moq := &MoqCreateRevocationList_genType_mock{Moq: m}
		return moq.Fn(rand, template, issuer, priv)
	}
}

func (m *MoqCreateRevocationList_genType_mock) Fn(rand io.Reader, template *x509.RevocationList, issuer *x509.Certificate, priv crypto.Signer) (result1 []byte, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqCreateRevocationList_genType_params{
		Rand:     rand,
		Template: template,
		Issuer:   issuer,
		Priv:     priv,
	}
	var results *MoqCreateRevocationList_genType_results
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
		result.DoFn(rand, template, issuer, priv)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(rand, template, issuer, priv)
	}
	return
}

func (m *MoqCreateRevocationList_genType) OnCall(rand io.Reader, template *x509.RevocationList, issuer *x509.Certificate, priv crypto.Signer) *MoqCreateRevocationList_genType_fnRecorder {
	return &MoqCreateRevocationList_genType_fnRecorder{
		Params: MoqCreateRevocationList_genType_params{
			Rand:     rand,
			Template: template,
			Issuer:   issuer,
			Priv:     priv,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqCreateRevocationList_genType_fnRecorder) Any() *MoqCreateRevocationList_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqCreateRevocationList_genType_anyParams{Recorder: r}
}

func (a *MoqCreateRevocationList_genType_anyParams) Rand() *MoqCreateRevocationList_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqCreateRevocationList_genType_anyParams) Template() *MoqCreateRevocationList_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqCreateRevocationList_genType_anyParams) Issuer() *MoqCreateRevocationList_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (a *MoqCreateRevocationList_genType_anyParams) Priv() *MoqCreateRevocationList_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 3
	return a.Recorder
}

func (r *MoqCreateRevocationList_genType_fnRecorder) Seq() *MoqCreateRevocationList_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqCreateRevocationList_genType_fnRecorder) NoSeq() *MoqCreateRevocationList_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqCreateRevocationList_genType_fnRecorder) ReturnResults(result1 []byte, result2 error) *MoqCreateRevocationList_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqCreateRevocationList_genType_doFn
		DoReturnFn MoqCreateRevocationList_genType_doReturnFn
	}{
		Values: &struct {
			Result1 []byte
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqCreateRevocationList_genType_fnRecorder) AndDo(fn MoqCreateRevocationList_genType_doFn) *MoqCreateRevocationList_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqCreateRevocationList_genType_fnRecorder) DoReturnResults(fn MoqCreateRevocationList_genType_doReturnFn) *MoqCreateRevocationList_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqCreateRevocationList_genType_doFn
		DoReturnFn MoqCreateRevocationList_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqCreateRevocationList_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqCreateRevocationList_genType_resultsByParams
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
		results = &MoqCreateRevocationList_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqCreateRevocationList_genType_paramsKey]*MoqCreateRevocationList_genType_results{},
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
		r.Results = &MoqCreateRevocationList_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqCreateRevocationList_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqCreateRevocationList_genType_fnRecorder {
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
					Result1 []byte
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqCreateRevocationList_genType_doFn
				DoReturnFn MoqCreateRevocationList_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqCreateRevocationList_genType) PrettyParams(params MoqCreateRevocationList_genType_params) string {
	return fmt.Sprintf("CreateRevocationList_genType(%#v, %#v, %#v, %#v)", params.Rand, params.Template, params.Issuer, params.Priv)
}

func (m *MoqCreateRevocationList_genType) ParamsKey(params MoqCreateRevocationList_genType_params, anyParams uint64) MoqCreateRevocationList_genType_paramsKey {
	m.Scene.T.Helper()
	var randUsed io.Reader
	var randUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Rand == moq.ParamIndexByValue {
			randUsed = params.Rand
		} else {
			randUsedHash = hash.DeepHash(params.Rand)
		}
	}
	var templateUsed *x509.RevocationList
	var templateUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Template == moq.ParamIndexByValue {
			templateUsed = params.Template
		} else {
			templateUsedHash = hash.DeepHash(params.Template)
		}
	}
	var issuerUsed *x509.Certificate
	var issuerUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Issuer == moq.ParamIndexByValue {
			issuerUsed = params.Issuer
		} else {
			issuerUsedHash = hash.DeepHash(params.Issuer)
		}
	}
	var privUsed crypto.Signer
	var privUsedHash hash.Hash
	if anyParams&(1<<3) == 0 {
		if m.Runtime.ParameterIndexing.Priv == moq.ParamIndexByValue {
			privUsed = params.Priv
		} else {
			privUsedHash = hash.DeepHash(params.Priv)
		}
	}
	return MoqCreateRevocationList_genType_paramsKey{
		Params: struct {
			Rand     io.Reader
			Template *x509.RevocationList
			Issuer   *x509.Certificate
			Priv     crypto.Signer
		}{
			Rand:     randUsed,
			Template: templateUsed,
			Issuer:   issuerUsed,
			Priv:     privUsed,
		},
		Hashes: struct {
			Rand     hash.Hash
			Template hash.Hash
			Issuer   hash.Hash
			Priv     hash.Hash
		}{
			Rand:     randUsedHash,
			Template: templateUsedHash,
			Issuer:   issuerUsedHash,
			Priv:     privUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqCreateRevocationList_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqCreateRevocationList_genType) AssertExpectationsMet() {
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

// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package tls

import (
	"crypto/tls"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// LoadX509KeyPair_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type LoadX509KeyPair_genType func(certFile, keyFile string) (tls.Certificate, error)

// MoqLoadX509KeyPair_genType holds the state of a moq of the
// LoadX509KeyPair_genType type
type MoqLoadX509KeyPair_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqLoadX509KeyPair_genType_mock

	ResultsByParams []MoqLoadX509KeyPair_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			CertFile moq.ParamIndexing
			KeyFile  moq.ParamIndexing
		}
	}
}

// MoqLoadX509KeyPair_genType_mock isolates the mock interface of the
// LoadX509KeyPair_genType type
type MoqLoadX509KeyPair_genType_mock struct {
	Moq *MoqLoadX509KeyPair_genType
}

// MoqLoadX509KeyPair_genType_params holds the params of the
// LoadX509KeyPair_genType type
type MoqLoadX509KeyPair_genType_params struct{ CertFile, KeyFile string }

// MoqLoadX509KeyPair_genType_paramsKey holds the map key params of the
// LoadX509KeyPair_genType type
type MoqLoadX509KeyPair_genType_paramsKey struct {
	Params struct{ CertFile, KeyFile string }
	Hashes struct{ CertFile, KeyFile hash.Hash }
}

// MoqLoadX509KeyPair_genType_resultsByParams contains the results for a given
// set of parameters for the LoadX509KeyPair_genType type
type MoqLoadX509KeyPair_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqLoadX509KeyPair_genType_paramsKey]*MoqLoadX509KeyPair_genType_results
}

// MoqLoadX509KeyPair_genType_doFn defines the type of function needed when
// calling AndDo for the LoadX509KeyPair_genType type
type MoqLoadX509KeyPair_genType_doFn func(certFile, keyFile string)

// MoqLoadX509KeyPair_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the LoadX509KeyPair_genType type
type MoqLoadX509KeyPair_genType_doReturnFn func(certFile, keyFile string) (tls.Certificate, error)

// MoqLoadX509KeyPair_genType_results holds the results of the
// LoadX509KeyPair_genType type
type MoqLoadX509KeyPair_genType_results struct {
	Params  MoqLoadX509KeyPair_genType_params
	Results []struct {
		Values *struct {
			Result1 tls.Certificate
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqLoadX509KeyPair_genType_doFn
		DoReturnFn MoqLoadX509KeyPair_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqLoadX509KeyPair_genType_fnRecorder routes recorded function calls to the
// MoqLoadX509KeyPair_genType moq
type MoqLoadX509KeyPair_genType_fnRecorder struct {
	Params    MoqLoadX509KeyPair_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqLoadX509KeyPair_genType_results
	Moq       *MoqLoadX509KeyPair_genType
}

// MoqLoadX509KeyPair_genType_anyParams isolates the any params functions of
// the LoadX509KeyPair_genType type
type MoqLoadX509KeyPair_genType_anyParams struct {
	Recorder *MoqLoadX509KeyPair_genType_fnRecorder
}

// NewMoqLoadX509KeyPair_genType creates a new moq of the
// LoadX509KeyPair_genType type
func NewMoqLoadX509KeyPair_genType(scene *moq.Scene, config *moq.Config) *MoqLoadX509KeyPair_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqLoadX509KeyPair_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqLoadX509KeyPair_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				CertFile moq.ParamIndexing
				KeyFile  moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			CertFile moq.ParamIndexing
			KeyFile  moq.ParamIndexing
		}{
			CertFile: moq.ParamIndexByValue,
			KeyFile:  moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the LoadX509KeyPair_genType type
func (m *MoqLoadX509KeyPair_genType) Mock() LoadX509KeyPair_genType {
	return func(certFile, keyFile string) (tls.Certificate, error) {
		m.Scene.T.Helper()
		moq := &MoqLoadX509KeyPair_genType_mock{Moq: m}
		return moq.Fn(certFile, keyFile)
	}
}

func (m *MoqLoadX509KeyPair_genType_mock) Fn(certFile, keyFile string) (result1 tls.Certificate, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqLoadX509KeyPair_genType_params{
		CertFile: certFile,
		KeyFile:  keyFile,
	}
	var results *MoqLoadX509KeyPair_genType_results
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
		result.DoFn(certFile, keyFile)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(certFile, keyFile)
	}
	return
}

func (m *MoqLoadX509KeyPair_genType) OnCall(certFile, keyFile string) *MoqLoadX509KeyPair_genType_fnRecorder {
	return &MoqLoadX509KeyPair_genType_fnRecorder{
		Params: MoqLoadX509KeyPair_genType_params{
			CertFile: certFile,
			KeyFile:  keyFile,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqLoadX509KeyPair_genType_fnRecorder) Any() *MoqLoadX509KeyPair_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqLoadX509KeyPair_genType_anyParams{Recorder: r}
}

func (a *MoqLoadX509KeyPair_genType_anyParams) CertFile() *MoqLoadX509KeyPair_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqLoadX509KeyPair_genType_anyParams) KeyFile() *MoqLoadX509KeyPair_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqLoadX509KeyPair_genType_fnRecorder) Seq() *MoqLoadX509KeyPair_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqLoadX509KeyPair_genType_fnRecorder) NoSeq() *MoqLoadX509KeyPair_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqLoadX509KeyPair_genType_fnRecorder) ReturnResults(result1 tls.Certificate, result2 error) *MoqLoadX509KeyPair_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 tls.Certificate
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqLoadX509KeyPair_genType_doFn
		DoReturnFn MoqLoadX509KeyPair_genType_doReturnFn
	}{
		Values: &struct {
			Result1 tls.Certificate
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqLoadX509KeyPair_genType_fnRecorder) AndDo(fn MoqLoadX509KeyPair_genType_doFn) *MoqLoadX509KeyPair_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqLoadX509KeyPair_genType_fnRecorder) DoReturnResults(fn MoqLoadX509KeyPair_genType_doReturnFn) *MoqLoadX509KeyPair_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 tls.Certificate
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqLoadX509KeyPair_genType_doFn
		DoReturnFn MoqLoadX509KeyPair_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqLoadX509KeyPair_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqLoadX509KeyPair_genType_resultsByParams
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
		results = &MoqLoadX509KeyPair_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqLoadX509KeyPair_genType_paramsKey]*MoqLoadX509KeyPair_genType_results{},
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
		r.Results = &MoqLoadX509KeyPair_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqLoadX509KeyPair_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqLoadX509KeyPair_genType_fnRecorder {
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
					Result1 tls.Certificate
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqLoadX509KeyPair_genType_doFn
				DoReturnFn MoqLoadX509KeyPair_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqLoadX509KeyPair_genType) PrettyParams(params MoqLoadX509KeyPair_genType_params) string {
	return fmt.Sprintf("LoadX509KeyPair_genType(%#v, %#v)", params.CertFile, params.KeyFile)
}

func (m *MoqLoadX509KeyPair_genType) ParamsKey(params MoqLoadX509KeyPair_genType_params, anyParams uint64) MoqLoadX509KeyPair_genType_paramsKey {
	m.Scene.T.Helper()
	var certFileUsed string
	var certFileUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.CertFile == moq.ParamIndexByValue {
			certFileUsed = params.CertFile
		} else {
			certFileUsedHash = hash.DeepHash(params.CertFile)
		}
	}
	var keyFileUsed string
	var keyFileUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.KeyFile == moq.ParamIndexByValue {
			keyFileUsed = params.KeyFile
		} else {
			keyFileUsedHash = hash.DeepHash(params.KeyFile)
		}
	}
	return MoqLoadX509KeyPair_genType_paramsKey{
		Params: struct{ CertFile, KeyFile string }{
			CertFile: certFileUsed,
			KeyFile:  keyFileUsed,
		},
		Hashes: struct{ CertFile, KeyFile hash.Hash }{
			CertFile: certFileUsedHash,
			KeyFile:  keyFileUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqLoadX509KeyPair_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqLoadX509KeyPair_genType) AssertExpectationsMet() {
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
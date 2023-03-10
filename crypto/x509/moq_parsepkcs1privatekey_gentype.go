// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package x509

import (
	"crypto/rsa"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// ParsePKCS1PrivateKey_genType is the fabricated implementation type of this
// mock (emitted when mocking functions directly and not from a function type)
type ParsePKCS1PrivateKey_genType func(der []byte) (*rsa.PrivateKey, error)

// MoqParsePKCS1PrivateKey_genType holds the state of a moq of the
// ParsePKCS1PrivateKey_genType type
type MoqParsePKCS1PrivateKey_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqParsePKCS1PrivateKey_genType_mock

	ResultsByParams []MoqParsePKCS1PrivateKey_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Der moq.ParamIndexing
		}
	}
}

// MoqParsePKCS1PrivateKey_genType_mock isolates the mock interface of the
// ParsePKCS1PrivateKey_genType type
type MoqParsePKCS1PrivateKey_genType_mock struct {
	Moq *MoqParsePKCS1PrivateKey_genType
}

// MoqParsePKCS1PrivateKey_genType_params holds the params of the
// ParsePKCS1PrivateKey_genType type
type MoqParsePKCS1PrivateKey_genType_params struct{ Der []byte }

// MoqParsePKCS1PrivateKey_genType_paramsKey holds the map key params of the
// ParsePKCS1PrivateKey_genType type
type MoqParsePKCS1PrivateKey_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ Der hash.Hash }
}

// MoqParsePKCS1PrivateKey_genType_resultsByParams contains the results for a
// given set of parameters for the ParsePKCS1PrivateKey_genType type
type MoqParsePKCS1PrivateKey_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqParsePKCS1PrivateKey_genType_paramsKey]*MoqParsePKCS1PrivateKey_genType_results
}

// MoqParsePKCS1PrivateKey_genType_doFn defines the type of function needed
// when calling AndDo for the ParsePKCS1PrivateKey_genType type
type MoqParsePKCS1PrivateKey_genType_doFn func(der []byte)

// MoqParsePKCS1PrivateKey_genType_doReturnFn defines the type of function
// needed when calling DoReturnResults for the ParsePKCS1PrivateKey_genType
// type
type MoqParsePKCS1PrivateKey_genType_doReturnFn func(der []byte) (*rsa.PrivateKey, error)

// MoqParsePKCS1PrivateKey_genType_results holds the results of the
// ParsePKCS1PrivateKey_genType type
type MoqParsePKCS1PrivateKey_genType_results struct {
	Params  MoqParsePKCS1PrivateKey_genType_params
	Results []struct {
		Values *struct {
			Result1 *rsa.PrivateKey
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqParsePKCS1PrivateKey_genType_doFn
		DoReturnFn MoqParsePKCS1PrivateKey_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqParsePKCS1PrivateKey_genType_fnRecorder routes recorded function calls to
// the MoqParsePKCS1PrivateKey_genType moq
type MoqParsePKCS1PrivateKey_genType_fnRecorder struct {
	Params    MoqParsePKCS1PrivateKey_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqParsePKCS1PrivateKey_genType_results
	Moq       *MoqParsePKCS1PrivateKey_genType
}

// MoqParsePKCS1PrivateKey_genType_anyParams isolates the any params functions
// of the ParsePKCS1PrivateKey_genType type
type MoqParsePKCS1PrivateKey_genType_anyParams struct {
	Recorder *MoqParsePKCS1PrivateKey_genType_fnRecorder
}

// NewMoqParsePKCS1PrivateKey_genType creates a new moq of the
// ParsePKCS1PrivateKey_genType type
func NewMoqParsePKCS1PrivateKey_genType(scene *moq.Scene, config *moq.Config) *MoqParsePKCS1PrivateKey_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqParsePKCS1PrivateKey_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqParsePKCS1PrivateKey_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Der moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Der moq.ParamIndexing
		}{
			Der: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ParsePKCS1PrivateKey_genType type
func (m *MoqParsePKCS1PrivateKey_genType) Mock() ParsePKCS1PrivateKey_genType {
	return func(der []byte) (*rsa.PrivateKey, error) {
		m.Scene.T.Helper()
		moq := &MoqParsePKCS1PrivateKey_genType_mock{Moq: m}
		return moq.Fn(der)
	}
}

func (m *MoqParsePKCS1PrivateKey_genType_mock) Fn(der []byte) (result1 *rsa.PrivateKey, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqParsePKCS1PrivateKey_genType_params{
		Der: der,
	}
	var results *MoqParsePKCS1PrivateKey_genType_results
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
		result.DoFn(der)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(der)
	}
	return
}

func (m *MoqParsePKCS1PrivateKey_genType) OnCall(der []byte) *MoqParsePKCS1PrivateKey_genType_fnRecorder {
	return &MoqParsePKCS1PrivateKey_genType_fnRecorder{
		Params: MoqParsePKCS1PrivateKey_genType_params{
			Der: der,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqParsePKCS1PrivateKey_genType_fnRecorder) Any() *MoqParsePKCS1PrivateKey_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqParsePKCS1PrivateKey_genType_anyParams{Recorder: r}
}

func (a *MoqParsePKCS1PrivateKey_genType_anyParams) Der() *MoqParsePKCS1PrivateKey_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqParsePKCS1PrivateKey_genType_fnRecorder) Seq() *MoqParsePKCS1PrivateKey_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqParsePKCS1PrivateKey_genType_fnRecorder) NoSeq() *MoqParsePKCS1PrivateKey_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqParsePKCS1PrivateKey_genType_fnRecorder) ReturnResults(result1 *rsa.PrivateKey, result2 error) *MoqParsePKCS1PrivateKey_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *rsa.PrivateKey
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqParsePKCS1PrivateKey_genType_doFn
		DoReturnFn MoqParsePKCS1PrivateKey_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *rsa.PrivateKey
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqParsePKCS1PrivateKey_genType_fnRecorder) AndDo(fn MoqParsePKCS1PrivateKey_genType_doFn) *MoqParsePKCS1PrivateKey_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqParsePKCS1PrivateKey_genType_fnRecorder) DoReturnResults(fn MoqParsePKCS1PrivateKey_genType_doReturnFn) *MoqParsePKCS1PrivateKey_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *rsa.PrivateKey
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqParsePKCS1PrivateKey_genType_doFn
		DoReturnFn MoqParsePKCS1PrivateKey_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqParsePKCS1PrivateKey_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqParsePKCS1PrivateKey_genType_resultsByParams
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
		results = &MoqParsePKCS1PrivateKey_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqParsePKCS1PrivateKey_genType_paramsKey]*MoqParsePKCS1PrivateKey_genType_results{},
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
		r.Results = &MoqParsePKCS1PrivateKey_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqParsePKCS1PrivateKey_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqParsePKCS1PrivateKey_genType_fnRecorder {
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
					Result1 *rsa.PrivateKey
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqParsePKCS1PrivateKey_genType_doFn
				DoReturnFn MoqParsePKCS1PrivateKey_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqParsePKCS1PrivateKey_genType) PrettyParams(params MoqParsePKCS1PrivateKey_genType_params) string {
	return fmt.Sprintf("ParsePKCS1PrivateKey_genType(%#v)", params.Der)
}

func (m *MoqParsePKCS1PrivateKey_genType) ParamsKey(params MoqParsePKCS1PrivateKey_genType_params, anyParams uint64) MoqParsePKCS1PrivateKey_genType_paramsKey {
	m.Scene.T.Helper()
	var derUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Der == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The der parameter can't be indexed by value")
		}
		derUsedHash = hash.DeepHash(params.Der)
	}
	return MoqParsePKCS1PrivateKey_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ Der hash.Hash }{
			Der: derUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqParsePKCS1PrivateKey_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqParsePKCS1PrivateKey_genType) AssertExpectationsMet() {
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

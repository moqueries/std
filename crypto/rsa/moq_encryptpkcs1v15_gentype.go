// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package rsa

import (
	"crypto/rsa"
	"fmt"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// EncryptPKCS1v15_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type EncryptPKCS1v15_genType func(rand io.Reader, pub *rsa.PublicKey, msg []byte) ([]byte, error)

// MoqEncryptPKCS1v15_genType holds the state of a moq of the
// EncryptPKCS1v15_genType type
type MoqEncryptPKCS1v15_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqEncryptPKCS1v15_genType_mock

	ResultsByParams []MoqEncryptPKCS1v15_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Rand moq.ParamIndexing
			Pub  moq.ParamIndexing
			Msg  moq.ParamIndexing
		}
	}
}

// MoqEncryptPKCS1v15_genType_mock isolates the mock interface of the
// EncryptPKCS1v15_genType type
type MoqEncryptPKCS1v15_genType_mock struct {
	Moq *MoqEncryptPKCS1v15_genType
}

// MoqEncryptPKCS1v15_genType_params holds the params of the
// EncryptPKCS1v15_genType type
type MoqEncryptPKCS1v15_genType_params struct {
	Rand io.Reader
	Pub  *rsa.PublicKey
	Msg  []byte
}

// MoqEncryptPKCS1v15_genType_paramsKey holds the map key params of the
// EncryptPKCS1v15_genType type
type MoqEncryptPKCS1v15_genType_paramsKey struct {
	Params struct {
		Rand io.Reader
		Pub  *rsa.PublicKey
	}
	Hashes struct {
		Rand hash.Hash
		Pub  hash.Hash
		Msg  hash.Hash
	}
}

// MoqEncryptPKCS1v15_genType_resultsByParams contains the results for a given
// set of parameters for the EncryptPKCS1v15_genType type
type MoqEncryptPKCS1v15_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqEncryptPKCS1v15_genType_paramsKey]*MoqEncryptPKCS1v15_genType_results
}

// MoqEncryptPKCS1v15_genType_doFn defines the type of function needed when
// calling AndDo for the EncryptPKCS1v15_genType type
type MoqEncryptPKCS1v15_genType_doFn func(rand io.Reader, pub *rsa.PublicKey, msg []byte)

// MoqEncryptPKCS1v15_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the EncryptPKCS1v15_genType type
type MoqEncryptPKCS1v15_genType_doReturnFn func(rand io.Reader, pub *rsa.PublicKey, msg []byte) ([]byte, error)

// MoqEncryptPKCS1v15_genType_results holds the results of the
// EncryptPKCS1v15_genType type
type MoqEncryptPKCS1v15_genType_results struct {
	Params  MoqEncryptPKCS1v15_genType_params
	Results []struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqEncryptPKCS1v15_genType_doFn
		DoReturnFn MoqEncryptPKCS1v15_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqEncryptPKCS1v15_genType_fnRecorder routes recorded function calls to the
// MoqEncryptPKCS1v15_genType moq
type MoqEncryptPKCS1v15_genType_fnRecorder struct {
	Params    MoqEncryptPKCS1v15_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqEncryptPKCS1v15_genType_results
	Moq       *MoqEncryptPKCS1v15_genType
}

// MoqEncryptPKCS1v15_genType_anyParams isolates the any params functions of
// the EncryptPKCS1v15_genType type
type MoqEncryptPKCS1v15_genType_anyParams struct {
	Recorder *MoqEncryptPKCS1v15_genType_fnRecorder
}

// NewMoqEncryptPKCS1v15_genType creates a new moq of the
// EncryptPKCS1v15_genType type
func NewMoqEncryptPKCS1v15_genType(scene *moq.Scene, config *moq.Config) *MoqEncryptPKCS1v15_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqEncryptPKCS1v15_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqEncryptPKCS1v15_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Rand moq.ParamIndexing
				Pub  moq.ParamIndexing
				Msg  moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Rand moq.ParamIndexing
			Pub  moq.ParamIndexing
			Msg  moq.ParamIndexing
		}{
			Rand: moq.ParamIndexByHash,
			Pub:  moq.ParamIndexByHash,
			Msg:  moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the EncryptPKCS1v15_genType type
func (m *MoqEncryptPKCS1v15_genType) Mock() EncryptPKCS1v15_genType {
	return func(rand io.Reader, pub *rsa.PublicKey, msg []byte) ([]byte, error) {
		m.Scene.T.Helper()
		moq := &MoqEncryptPKCS1v15_genType_mock{Moq: m}
		return moq.Fn(rand, pub, msg)
	}
}

func (m *MoqEncryptPKCS1v15_genType_mock) Fn(rand io.Reader, pub *rsa.PublicKey, msg []byte) (result1 []byte, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqEncryptPKCS1v15_genType_params{
		Rand: rand,
		Pub:  pub,
		Msg:  msg,
	}
	var results *MoqEncryptPKCS1v15_genType_results
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
		result.DoFn(rand, pub, msg)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(rand, pub, msg)
	}
	return
}

func (m *MoqEncryptPKCS1v15_genType) OnCall(rand io.Reader, pub *rsa.PublicKey, msg []byte) *MoqEncryptPKCS1v15_genType_fnRecorder {
	return &MoqEncryptPKCS1v15_genType_fnRecorder{
		Params: MoqEncryptPKCS1v15_genType_params{
			Rand: rand,
			Pub:  pub,
			Msg:  msg,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqEncryptPKCS1v15_genType_fnRecorder) Any() *MoqEncryptPKCS1v15_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqEncryptPKCS1v15_genType_anyParams{Recorder: r}
}

func (a *MoqEncryptPKCS1v15_genType_anyParams) Rand() *MoqEncryptPKCS1v15_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqEncryptPKCS1v15_genType_anyParams) Pub() *MoqEncryptPKCS1v15_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqEncryptPKCS1v15_genType_anyParams) Msg() *MoqEncryptPKCS1v15_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqEncryptPKCS1v15_genType_fnRecorder) Seq() *MoqEncryptPKCS1v15_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqEncryptPKCS1v15_genType_fnRecorder) NoSeq() *MoqEncryptPKCS1v15_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqEncryptPKCS1v15_genType_fnRecorder) ReturnResults(result1 []byte, result2 error) *MoqEncryptPKCS1v15_genType_fnRecorder {
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
		DoFn       MoqEncryptPKCS1v15_genType_doFn
		DoReturnFn MoqEncryptPKCS1v15_genType_doReturnFn
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

func (r *MoqEncryptPKCS1v15_genType_fnRecorder) AndDo(fn MoqEncryptPKCS1v15_genType_doFn) *MoqEncryptPKCS1v15_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqEncryptPKCS1v15_genType_fnRecorder) DoReturnResults(fn MoqEncryptPKCS1v15_genType_doReturnFn) *MoqEncryptPKCS1v15_genType_fnRecorder {
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
		DoFn       MoqEncryptPKCS1v15_genType_doFn
		DoReturnFn MoqEncryptPKCS1v15_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqEncryptPKCS1v15_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqEncryptPKCS1v15_genType_resultsByParams
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
		results = &MoqEncryptPKCS1v15_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqEncryptPKCS1v15_genType_paramsKey]*MoqEncryptPKCS1v15_genType_results{},
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
		r.Results = &MoqEncryptPKCS1v15_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqEncryptPKCS1v15_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqEncryptPKCS1v15_genType_fnRecorder {
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
				DoFn       MoqEncryptPKCS1v15_genType_doFn
				DoReturnFn MoqEncryptPKCS1v15_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqEncryptPKCS1v15_genType) PrettyParams(params MoqEncryptPKCS1v15_genType_params) string {
	return fmt.Sprintf("EncryptPKCS1v15_genType(%#v, %#v, %#v)", params.Rand, params.Pub, params.Msg)
}

func (m *MoqEncryptPKCS1v15_genType) ParamsKey(params MoqEncryptPKCS1v15_genType_params, anyParams uint64) MoqEncryptPKCS1v15_genType_paramsKey {
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
	var pubUsed *rsa.PublicKey
	var pubUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Pub == moq.ParamIndexByValue {
			pubUsed = params.Pub
		} else {
			pubUsedHash = hash.DeepHash(params.Pub)
		}
	}
	var msgUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Msg == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The msg parameter can't be indexed by value")
		}
		msgUsedHash = hash.DeepHash(params.Msg)
	}
	return MoqEncryptPKCS1v15_genType_paramsKey{
		Params: struct {
			Rand io.Reader
			Pub  *rsa.PublicKey
		}{
			Rand: randUsed,
			Pub:  pubUsed,
		},
		Hashes: struct {
			Rand hash.Hash
			Pub  hash.Hash
			Msg  hash.Hash
		}{
			Rand: randUsedHash,
			Pub:  pubUsedHash,
			Msg:  msgUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqEncryptPKCS1v15_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqEncryptPKCS1v15_genType) AssertExpectationsMet() {
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

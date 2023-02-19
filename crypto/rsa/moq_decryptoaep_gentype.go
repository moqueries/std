// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package rsa

import (
	"crypto/rsa"
	"fmt"
	"hash"
	"io"
	"math/bits"
	"sync/atomic"

	hash1 "moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// DecryptOAEP_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type DecryptOAEP_genType func(hash hash.Hash, random io.Reader, priv *rsa.PrivateKey, ciphertext []byte, label []byte) ([]byte, error)

// MoqDecryptOAEP_genType holds the state of a moq of the DecryptOAEP_genType
// type
type MoqDecryptOAEP_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqDecryptOAEP_genType_mock

	ResultsByParams []MoqDecryptOAEP_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Hash       moq.ParamIndexing
			Random     moq.ParamIndexing
			Priv       moq.ParamIndexing
			Ciphertext moq.ParamIndexing
			Label      moq.ParamIndexing
		}
	}
}

// MoqDecryptOAEP_genType_mock isolates the mock interface of the
// DecryptOAEP_genType type
type MoqDecryptOAEP_genType_mock struct {
	Moq *MoqDecryptOAEP_genType
}

// MoqDecryptOAEP_genType_params holds the params of the DecryptOAEP_genType
// type
type MoqDecryptOAEP_genType_params struct {
	Hash       hash.Hash
	Random     io.Reader
	Priv       *rsa.PrivateKey
	Ciphertext []byte
	Label      []byte
}

// MoqDecryptOAEP_genType_paramsKey holds the map key params of the
// DecryptOAEP_genType type
type MoqDecryptOAEP_genType_paramsKey struct {
	Params struct {
		Hash   hash.Hash
		Random io.Reader
		Priv   *rsa.PrivateKey
	}
	Hashes struct {
		Hash       hash1.Hash
		Random     hash1.Hash
		Priv       hash1.Hash
		Ciphertext hash1.Hash
		Label      hash1.Hash
	}
}

// MoqDecryptOAEP_genType_resultsByParams contains the results for a given set
// of parameters for the DecryptOAEP_genType type
type MoqDecryptOAEP_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqDecryptOAEP_genType_paramsKey]*MoqDecryptOAEP_genType_results
}

// MoqDecryptOAEP_genType_doFn defines the type of function needed when calling
// AndDo for the DecryptOAEP_genType type
type MoqDecryptOAEP_genType_doFn func(hash hash.Hash, random io.Reader, priv *rsa.PrivateKey, ciphertext []byte, label []byte)

// MoqDecryptOAEP_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the DecryptOAEP_genType type
type MoqDecryptOAEP_genType_doReturnFn func(hash hash.Hash, random io.Reader, priv *rsa.PrivateKey, ciphertext []byte, label []byte) ([]byte, error)

// MoqDecryptOAEP_genType_results holds the results of the DecryptOAEP_genType
// type
type MoqDecryptOAEP_genType_results struct {
	Params  MoqDecryptOAEP_genType_params
	Results []struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqDecryptOAEP_genType_doFn
		DoReturnFn MoqDecryptOAEP_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqDecryptOAEP_genType_fnRecorder routes recorded function calls to the
// MoqDecryptOAEP_genType moq
type MoqDecryptOAEP_genType_fnRecorder struct {
	Params    MoqDecryptOAEP_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqDecryptOAEP_genType_results
	Moq       *MoqDecryptOAEP_genType
}

// MoqDecryptOAEP_genType_anyParams isolates the any params functions of the
// DecryptOAEP_genType type
type MoqDecryptOAEP_genType_anyParams struct {
	Recorder *MoqDecryptOAEP_genType_fnRecorder
}

// NewMoqDecryptOAEP_genType creates a new moq of the DecryptOAEP_genType type
func NewMoqDecryptOAEP_genType(scene *moq.Scene, config *moq.Config) *MoqDecryptOAEP_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqDecryptOAEP_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqDecryptOAEP_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Hash       moq.ParamIndexing
				Random     moq.ParamIndexing
				Priv       moq.ParamIndexing
				Ciphertext moq.ParamIndexing
				Label      moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Hash       moq.ParamIndexing
			Random     moq.ParamIndexing
			Priv       moq.ParamIndexing
			Ciphertext moq.ParamIndexing
			Label      moq.ParamIndexing
		}{
			Hash:       moq.ParamIndexByHash,
			Random:     moq.ParamIndexByHash,
			Priv:       moq.ParamIndexByHash,
			Ciphertext: moq.ParamIndexByHash,
			Label:      moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the DecryptOAEP_genType type
func (m *MoqDecryptOAEP_genType) Mock() DecryptOAEP_genType {
	return func(hash hash.Hash, random io.Reader, priv *rsa.PrivateKey, ciphertext []byte, label []byte) ([]byte, error) {
		m.Scene.T.Helper()
		moq := &MoqDecryptOAEP_genType_mock{Moq: m}
		return moq.Fn(hash, random, priv, ciphertext, label)
	}
}

func (m *MoqDecryptOAEP_genType_mock) Fn(hash hash.Hash, random io.Reader, priv *rsa.PrivateKey, ciphertext []byte, label []byte) (result1 []byte, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqDecryptOAEP_genType_params{
		Hash:       hash,
		Random:     random,
		Priv:       priv,
		Ciphertext: ciphertext,
		Label:      label,
	}
	var results *MoqDecryptOAEP_genType_results
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
		result.DoFn(hash, random, priv, ciphertext, label)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(hash, random, priv, ciphertext, label)
	}
	return
}

func (m *MoqDecryptOAEP_genType) OnCall(hash hash.Hash, random io.Reader, priv *rsa.PrivateKey, ciphertext []byte, label []byte) *MoqDecryptOAEP_genType_fnRecorder {
	return &MoqDecryptOAEP_genType_fnRecorder{
		Params: MoqDecryptOAEP_genType_params{
			Hash:       hash,
			Random:     random,
			Priv:       priv,
			Ciphertext: ciphertext,
			Label:      label,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqDecryptOAEP_genType_fnRecorder) Any() *MoqDecryptOAEP_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqDecryptOAEP_genType_anyParams{Recorder: r}
}

func (a *MoqDecryptOAEP_genType_anyParams) Hash() *MoqDecryptOAEP_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqDecryptOAEP_genType_anyParams) Random() *MoqDecryptOAEP_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqDecryptOAEP_genType_anyParams) Priv() *MoqDecryptOAEP_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (a *MoqDecryptOAEP_genType_anyParams) Ciphertext() *MoqDecryptOAEP_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 3
	return a.Recorder
}

func (a *MoqDecryptOAEP_genType_anyParams) Label() *MoqDecryptOAEP_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 4
	return a.Recorder
}

func (r *MoqDecryptOAEP_genType_fnRecorder) Seq() *MoqDecryptOAEP_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqDecryptOAEP_genType_fnRecorder) NoSeq() *MoqDecryptOAEP_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqDecryptOAEP_genType_fnRecorder) ReturnResults(result1 []byte, result2 error) *MoqDecryptOAEP_genType_fnRecorder {
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
		DoFn       MoqDecryptOAEP_genType_doFn
		DoReturnFn MoqDecryptOAEP_genType_doReturnFn
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

func (r *MoqDecryptOAEP_genType_fnRecorder) AndDo(fn MoqDecryptOAEP_genType_doFn) *MoqDecryptOAEP_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqDecryptOAEP_genType_fnRecorder) DoReturnResults(fn MoqDecryptOAEP_genType_doReturnFn) *MoqDecryptOAEP_genType_fnRecorder {
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
		DoFn       MoqDecryptOAEP_genType_doFn
		DoReturnFn MoqDecryptOAEP_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqDecryptOAEP_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqDecryptOAEP_genType_resultsByParams
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
		results = &MoqDecryptOAEP_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqDecryptOAEP_genType_paramsKey]*MoqDecryptOAEP_genType_results{},
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
		r.Results = &MoqDecryptOAEP_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqDecryptOAEP_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqDecryptOAEP_genType_fnRecorder {
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
				DoFn       MoqDecryptOAEP_genType_doFn
				DoReturnFn MoqDecryptOAEP_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqDecryptOAEP_genType) PrettyParams(params MoqDecryptOAEP_genType_params) string {
	return fmt.Sprintf("DecryptOAEP_genType(%#v, %#v, %#v, %#v, %#v)", params.Hash, params.Random, params.Priv, params.Ciphertext, params.Label)
}

func (m *MoqDecryptOAEP_genType) ParamsKey(params MoqDecryptOAEP_genType_params, anyParams uint64) MoqDecryptOAEP_genType_paramsKey {
	m.Scene.T.Helper()
	var hashUsed hash.Hash
	var hashUsedHash hash1.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Hash == moq.ParamIndexByValue {
			hashUsed = params.Hash
		} else {
			hashUsedHash = hash1.DeepHash(params.Hash)
		}
	}
	var randomUsed io.Reader
	var randomUsedHash hash1.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Random == moq.ParamIndexByValue {
			randomUsed = params.Random
		} else {
			randomUsedHash = hash1.DeepHash(params.Random)
		}
	}
	var privUsed *rsa.PrivateKey
	var privUsedHash hash1.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Priv == moq.ParamIndexByValue {
			privUsed = params.Priv
		} else {
			privUsedHash = hash1.DeepHash(params.Priv)
		}
	}
	var ciphertextUsedHash hash1.Hash
	if anyParams&(1<<3) == 0 {
		if m.Runtime.ParameterIndexing.Ciphertext == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The ciphertext parameter can't be indexed by value")
		}
		ciphertextUsedHash = hash1.DeepHash(params.Ciphertext)
	}
	var labelUsedHash hash1.Hash
	if anyParams&(1<<4) == 0 {
		if m.Runtime.ParameterIndexing.Label == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The label parameter can't be indexed by value")
		}
		labelUsedHash = hash1.DeepHash(params.Label)
	}
	return MoqDecryptOAEP_genType_paramsKey{
		Params: struct {
			Hash   hash.Hash
			Random io.Reader
			Priv   *rsa.PrivateKey
		}{
			Hash:   hashUsed,
			Random: randomUsed,
			Priv:   privUsed,
		},
		Hashes: struct {
			Hash       hash1.Hash
			Random     hash1.Hash
			Priv       hash1.Hash
			Ciphertext hash1.Hash
			Label      hash1.Hash
		}{
			Hash:       hashUsedHash,
			Random:     randomUsedHash,
			Priv:       privUsedHash,
			Ciphertext: ciphertextUsedHash,
			Label:      labelUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqDecryptOAEP_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqDecryptOAEP_genType) AssertExpectationsMet() {
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
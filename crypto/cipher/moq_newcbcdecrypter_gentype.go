// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package cipher

import (
	"crypto/cipher"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// NewCBCDecrypter_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type NewCBCDecrypter_genType func(b cipher.Block, iv []byte) cipher.BlockMode

// MoqNewCBCDecrypter_genType holds the state of a moq of the
// NewCBCDecrypter_genType type
type MoqNewCBCDecrypter_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNewCBCDecrypter_genType_mock

	ResultsByParams []MoqNewCBCDecrypter_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			B  moq.ParamIndexing
			Iv moq.ParamIndexing
		}
	}
}

// MoqNewCBCDecrypter_genType_mock isolates the mock interface of the
// NewCBCDecrypter_genType type
type MoqNewCBCDecrypter_genType_mock struct {
	Moq *MoqNewCBCDecrypter_genType
}

// MoqNewCBCDecrypter_genType_params holds the params of the
// NewCBCDecrypter_genType type
type MoqNewCBCDecrypter_genType_params struct {
	B  cipher.Block
	Iv []byte
}

// MoqNewCBCDecrypter_genType_paramsKey holds the map key params of the
// NewCBCDecrypter_genType type
type MoqNewCBCDecrypter_genType_paramsKey struct {
	Params struct{ B cipher.Block }
	Hashes struct {
		B  hash.Hash
		Iv hash.Hash
	}
}

// MoqNewCBCDecrypter_genType_resultsByParams contains the results for a given
// set of parameters for the NewCBCDecrypter_genType type
type MoqNewCBCDecrypter_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNewCBCDecrypter_genType_paramsKey]*MoqNewCBCDecrypter_genType_results
}

// MoqNewCBCDecrypter_genType_doFn defines the type of function needed when
// calling AndDo for the NewCBCDecrypter_genType type
type MoqNewCBCDecrypter_genType_doFn func(b cipher.Block, iv []byte)

// MoqNewCBCDecrypter_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the NewCBCDecrypter_genType type
type MoqNewCBCDecrypter_genType_doReturnFn func(b cipher.Block, iv []byte) cipher.BlockMode

// MoqNewCBCDecrypter_genType_results holds the results of the
// NewCBCDecrypter_genType type
type MoqNewCBCDecrypter_genType_results struct {
	Params  MoqNewCBCDecrypter_genType_params
	Results []struct {
		Values *struct {
			Result1 cipher.BlockMode
		}
		Sequence   uint32
		DoFn       MoqNewCBCDecrypter_genType_doFn
		DoReturnFn MoqNewCBCDecrypter_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNewCBCDecrypter_genType_fnRecorder routes recorded function calls to the
// MoqNewCBCDecrypter_genType moq
type MoqNewCBCDecrypter_genType_fnRecorder struct {
	Params    MoqNewCBCDecrypter_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNewCBCDecrypter_genType_results
	Moq       *MoqNewCBCDecrypter_genType
}

// MoqNewCBCDecrypter_genType_anyParams isolates the any params functions of
// the NewCBCDecrypter_genType type
type MoqNewCBCDecrypter_genType_anyParams struct {
	Recorder *MoqNewCBCDecrypter_genType_fnRecorder
}

// NewMoqNewCBCDecrypter_genType creates a new moq of the
// NewCBCDecrypter_genType type
func NewMoqNewCBCDecrypter_genType(scene *moq.Scene, config *moq.Config) *MoqNewCBCDecrypter_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNewCBCDecrypter_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNewCBCDecrypter_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				B  moq.ParamIndexing
				Iv moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			B  moq.ParamIndexing
			Iv moq.ParamIndexing
		}{
			B:  moq.ParamIndexByHash,
			Iv: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the NewCBCDecrypter_genType type
func (m *MoqNewCBCDecrypter_genType) Mock() NewCBCDecrypter_genType {
	return func(b cipher.Block, iv []byte) cipher.BlockMode {
		m.Scene.T.Helper()
		moq := &MoqNewCBCDecrypter_genType_mock{Moq: m}
		return moq.Fn(b, iv)
	}
}

func (m *MoqNewCBCDecrypter_genType_mock) Fn(b cipher.Block, iv []byte) (result1 cipher.BlockMode) {
	m.Moq.Scene.T.Helper()
	params := MoqNewCBCDecrypter_genType_params{
		B:  b,
		Iv: iv,
	}
	var results *MoqNewCBCDecrypter_genType_results
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
		result.DoFn(b, iv)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(b, iv)
	}
	return
}

func (m *MoqNewCBCDecrypter_genType) OnCall(b cipher.Block, iv []byte) *MoqNewCBCDecrypter_genType_fnRecorder {
	return &MoqNewCBCDecrypter_genType_fnRecorder{
		Params: MoqNewCBCDecrypter_genType_params{
			B:  b,
			Iv: iv,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNewCBCDecrypter_genType_fnRecorder) Any() *MoqNewCBCDecrypter_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNewCBCDecrypter_genType_anyParams{Recorder: r}
}

func (a *MoqNewCBCDecrypter_genType_anyParams) B() *MoqNewCBCDecrypter_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqNewCBCDecrypter_genType_anyParams) Iv() *MoqNewCBCDecrypter_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqNewCBCDecrypter_genType_fnRecorder) Seq() *MoqNewCBCDecrypter_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNewCBCDecrypter_genType_fnRecorder) NoSeq() *MoqNewCBCDecrypter_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNewCBCDecrypter_genType_fnRecorder) ReturnResults(result1 cipher.BlockMode) *MoqNewCBCDecrypter_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 cipher.BlockMode
		}
		Sequence   uint32
		DoFn       MoqNewCBCDecrypter_genType_doFn
		DoReturnFn MoqNewCBCDecrypter_genType_doReturnFn
	}{
		Values: &struct {
			Result1 cipher.BlockMode
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNewCBCDecrypter_genType_fnRecorder) AndDo(fn MoqNewCBCDecrypter_genType_doFn) *MoqNewCBCDecrypter_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNewCBCDecrypter_genType_fnRecorder) DoReturnResults(fn MoqNewCBCDecrypter_genType_doReturnFn) *MoqNewCBCDecrypter_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 cipher.BlockMode
		}
		Sequence   uint32
		DoFn       MoqNewCBCDecrypter_genType_doFn
		DoReturnFn MoqNewCBCDecrypter_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNewCBCDecrypter_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNewCBCDecrypter_genType_resultsByParams
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
		results = &MoqNewCBCDecrypter_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNewCBCDecrypter_genType_paramsKey]*MoqNewCBCDecrypter_genType_results{},
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
		r.Results = &MoqNewCBCDecrypter_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNewCBCDecrypter_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNewCBCDecrypter_genType_fnRecorder {
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
					Result1 cipher.BlockMode
				}
				Sequence   uint32
				DoFn       MoqNewCBCDecrypter_genType_doFn
				DoReturnFn MoqNewCBCDecrypter_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNewCBCDecrypter_genType) PrettyParams(params MoqNewCBCDecrypter_genType_params) string {
	return fmt.Sprintf("NewCBCDecrypter_genType(%#v, %#v)", params.B, params.Iv)
}

func (m *MoqNewCBCDecrypter_genType) ParamsKey(params MoqNewCBCDecrypter_genType_params, anyParams uint64) MoqNewCBCDecrypter_genType_paramsKey {
	m.Scene.T.Helper()
	var bUsed cipher.Block
	var bUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.B == moq.ParamIndexByValue {
			bUsed = params.B
		} else {
			bUsedHash = hash.DeepHash(params.B)
		}
	}
	var ivUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Iv == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The iv parameter can't be indexed by value")
		}
		ivUsedHash = hash.DeepHash(params.Iv)
	}
	return MoqNewCBCDecrypter_genType_paramsKey{
		Params: struct{ B cipher.Block }{
			B: bUsed,
		},
		Hashes: struct {
			B  hash.Hash
			Iv hash.Hash
		}{
			B:  bUsedHash,
			Iv: ivUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNewCBCDecrypter_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNewCBCDecrypter_genType) AssertExpectationsMet() {
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

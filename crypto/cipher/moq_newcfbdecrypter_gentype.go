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

// NewCFBDecrypter_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type NewCFBDecrypter_genType func(block cipher.Block, iv []byte) cipher.Stream

// MoqNewCFBDecrypter_genType holds the state of a moq of the
// NewCFBDecrypter_genType type
type MoqNewCFBDecrypter_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNewCFBDecrypter_genType_mock

	ResultsByParams []MoqNewCFBDecrypter_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Block moq.ParamIndexing
			Iv    moq.ParamIndexing
		}
	}
}

// MoqNewCFBDecrypter_genType_mock isolates the mock interface of the
// NewCFBDecrypter_genType type
type MoqNewCFBDecrypter_genType_mock struct {
	Moq *MoqNewCFBDecrypter_genType
}

// MoqNewCFBDecrypter_genType_params holds the params of the
// NewCFBDecrypter_genType type
type MoqNewCFBDecrypter_genType_params struct {
	Block cipher.Block
	Iv    []byte
}

// MoqNewCFBDecrypter_genType_paramsKey holds the map key params of the
// NewCFBDecrypter_genType type
type MoqNewCFBDecrypter_genType_paramsKey struct {
	Params struct{ Block cipher.Block }
	Hashes struct {
		Block hash.Hash
		Iv    hash.Hash
	}
}

// MoqNewCFBDecrypter_genType_resultsByParams contains the results for a given
// set of parameters for the NewCFBDecrypter_genType type
type MoqNewCFBDecrypter_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNewCFBDecrypter_genType_paramsKey]*MoqNewCFBDecrypter_genType_results
}

// MoqNewCFBDecrypter_genType_doFn defines the type of function needed when
// calling AndDo for the NewCFBDecrypter_genType type
type MoqNewCFBDecrypter_genType_doFn func(block cipher.Block, iv []byte)

// MoqNewCFBDecrypter_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the NewCFBDecrypter_genType type
type MoqNewCFBDecrypter_genType_doReturnFn func(block cipher.Block, iv []byte) cipher.Stream

// MoqNewCFBDecrypter_genType_results holds the results of the
// NewCFBDecrypter_genType type
type MoqNewCFBDecrypter_genType_results struct {
	Params  MoqNewCFBDecrypter_genType_params
	Results []struct {
		Values *struct {
			Result1 cipher.Stream
		}
		Sequence   uint32
		DoFn       MoqNewCFBDecrypter_genType_doFn
		DoReturnFn MoqNewCFBDecrypter_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNewCFBDecrypter_genType_fnRecorder routes recorded function calls to the
// MoqNewCFBDecrypter_genType moq
type MoqNewCFBDecrypter_genType_fnRecorder struct {
	Params    MoqNewCFBDecrypter_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNewCFBDecrypter_genType_results
	Moq       *MoqNewCFBDecrypter_genType
}

// MoqNewCFBDecrypter_genType_anyParams isolates the any params functions of
// the NewCFBDecrypter_genType type
type MoqNewCFBDecrypter_genType_anyParams struct {
	Recorder *MoqNewCFBDecrypter_genType_fnRecorder
}

// NewMoqNewCFBDecrypter_genType creates a new moq of the
// NewCFBDecrypter_genType type
func NewMoqNewCFBDecrypter_genType(scene *moq.Scene, config *moq.Config) *MoqNewCFBDecrypter_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNewCFBDecrypter_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNewCFBDecrypter_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Block moq.ParamIndexing
				Iv    moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Block moq.ParamIndexing
			Iv    moq.ParamIndexing
		}{
			Block: moq.ParamIndexByHash,
			Iv:    moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the NewCFBDecrypter_genType type
func (m *MoqNewCFBDecrypter_genType) Mock() NewCFBDecrypter_genType {
	return func(block cipher.Block, iv []byte) cipher.Stream {
		m.Scene.T.Helper()
		moq := &MoqNewCFBDecrypter_genType_mock{Moq: m}
		return moq.Fn(block, iv)
	}
}

func (m *MoqNewCFBDecrypter_genType_mock) Fn(block cipher.Block, iv []byte) (result1 cipher.Stream) {
	m.Moq.Scene.T.Helper()
	params := MoqNewCFBDecrypter_genType_params{
		Block: block,
		Iv:    iv,
	}
	var results *MoqNewCFBDecrypter_genType_results
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
		result.DoFn(block, iv)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(block, iv)
	}
	return
}

func (m *MoqNewCFBDecrypter_genType) OnCall(block cipher.Block, iv []byte) *MoqNewCFBDecrypter_genType_fnRecorder {
	return &MoqNewCFBDecrypter_genType_fnRecorder{
		Params: MoqNewCFBDecrypter_genType_params{
			Block: block,
			Iv:    iv,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNewCFBDecrypter_genType_fnRecorder) Any() *MoqNewCFBDecrypter_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNewCFBDecrypter_genType_anyParams{Recorder: r}
}

func (a *MoqNewCFBDecrypter_genType_anyParams) Block() *MoqNewCFBDecrypter_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqNewCFBDecrypter_genType_anyParams) Iv() *MoqNewCFBDecrypter_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqNewCFBDecrypter_genType_fnRecorder) Seq() *MoqNewCFBDecrypter_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNewCFBDecrypter_genType_fnRecorder) NoSeq() *MoqNewCFBDecrypter_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNewCFBDecrypter_genType_fnRecorder) ReturnResults(result1 cipher.Stream) *MoqNewCFBDecrypter_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 cipher.Stream
		}
		Sequence   uint32
		DoFn       MoqNewCFBDecrypter_genType_doFn
		DoReturnFn MoqNewCFBDecrypter_genType_doReturnFn
	}{
		Values: &struct {
			Result1 cipher.Stream
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNewCFBDecrypter_genType_fnRecorder) AndDo(fn MoqNewCFBDecrypter_genType_doFn) *MoqNewCFBDecrypter_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNewCFBDecrypter_genType_fnRecorder) DoReturnResults(fn MoqNewCFBDecrypter_genType_doReturnFn) *MoqNewCFBDecrypter_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 cipher.Stream
		}
		Sequence   uint32
		DoFn       MoqNewCFBDecrypter_genType_doFn
		DoReturnFn MoqNewCFBDecrypter_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNewCFBDecrypter_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNewCFBDecrypter_genType_resultsByParams
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
		results = &MoqNewCFBDecrypter_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNewCFBDecrypter_genType_paramsKey]*MoqNewCFBDecrypter_genType_results{},
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
		r.Results = &MoqNewCFBDecrypter_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNewCFBDecrypter_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNewCFBDecrypter_genType_fnRecorder {
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
					Result1 cipher.Stream
				}
				Sequence   uint32
				DoFn       MoqNewCFBDecrypter_genType_doFn
				DoReturnFn MoqNewCFBDecrypter_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNewCFBDecrypter_genType) PrettyParams(params MoqNewCFBDecrypter_genType_params) string {
	return fmt.Sprintf("NewCFBDecrypter_genType(%#v, %#v)", params.Block, params.Iv)
}

func (m *MoqNewCFBDecrypter_genType) ParamsKey(params MoqNewCFBDecrypter_genType_params, anyParams uint64) MoqNewCFBDecrypter_genType_paramsKey {
	m.Scene.T.Helper()
	var blockUsed cipher.Block
	var blockUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Block == moq.ParamIndexByValue {
			blockUsed = params.Block
		} else {
			blockUsedHash = hash.DeepHash(params.Block)
		}
	}
	var ivUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Iv == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The iv parameter can't be indexed by value")
		}
		ivUsedHash = hash.DeepHash(params.Iv)
	}
	return MoqNewCFBDecrypter_genType_paramsKey{
		Params: struct{ Block cipher.Block }{
			Block: blockUsed,
		},
		Hashes: struct {
			Block hash.Hash
			Iv    hash.Hash
		}{
			Block: blockUsedHash,
			Iv:    ivUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNewCFBDecrypter_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNewCFBDecrypter_genType) AssertExpectationsMet() {
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

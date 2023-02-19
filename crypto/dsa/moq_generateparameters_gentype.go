// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package dsa

import (
	"crypto/dsa"
	"fmt"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// GenerateParameters_genType is the fabricated implementation type of this
// mock (emitted when mocking functions directly and not from a function type)
type GenerateParameters_genType func(params *dsa.Parameters, rand io.Reader, sizes dsa.ParameterSizes) error

// MoqGenerateParameters_genType holds the state of a moq of the
// GenerateParameters_genType type
type MoqGenerateParameters_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqGenerateParameters_genType_mock

	ResultsByParams []MoqGenerateParameters_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Param1 moq.ParamIndexing
			Rand   moq.ParamIndexing
			Sizes  moq.ParamIndexing
		}
	}
}

// MoqGenerateParameters_genType_mock isolates the mock interface of the
// GenerateParameters_genType type
type MoqGenerateParameters_genType_mock struct {
	Moq *MoqGenerateParameters_genType
}

// MoqGenerateParameters_genType_params holds the params of the
// GenerateParameters_genType type
type MoqGenerateParameters_genType_params struct {
	Param1 *dsa.Parameters
	Rand   io.Reader
	Sizes  dsa.ParameterSizes
}

// MoqGenerateParameters_genType_paramsKey holds the map key params of the
// GenerateParameters_genType type
type MoqGenerateParameters_genType_paramsKey struct {
	Params struct {
		Param1 *dsa.Parameters
		Rand   io.Reader
		Sizes  dsa.ParameterSizes
	}
	Hashes struct {
		Param1 hash.Hash
		Rand   hash.Hash
		Sizes  hash.Hash
	}
}

// MoqGenerateParameters_genType_resultsByParams contains the results for a
// given set of parameters for the GenerateParameters_genType type
type MoqGenerateParameters_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqGenerateParameters_genType_paramsKey]*MoqGenerateParameters_genType_results
}

// MoqGenerateParameters_genType_doFn defines the type of function needed when
// calling AndDo for the GenerateParameters_genType type
type MoqGenerateParameters_genType_doFn func(params *dsa.Parameters, rand io.Reader, sizes dsa.ParameterSizes)

// MoqGenerateParameters_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the GenerateParameters_genType type
type MoqGenerateParameters_genType_doReturnFn func(params *dsa.Parameters, rand io.Reader, sizes dsa.ParameterSizes) error

// MoqGenerateParameters_genType_results holds the results of the
// GenerateParameters_genType type
type MoqGenerateParameters_genType_results struct {
	Params  MoqGenerateParameters_genType_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqGenerateParameters_genType_doFn
		DoReturnFn MoqGenerateParameters_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqGenerateParameters_genType_fnRecorder routes recorded function calls to
// the MoqGenerateParameters_genType moq
type MoqGenerateParameters_genType_fnRecorder struct {
	Params    MoqGenerateParameters_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqGenerateParameters_genType_results
	Moq       *MoqGenerateParameters_genType
}

// MoqGenerateParameters_genType_anyParams isolates the any params functions of
// the GenerateParameters_genType type
type MoqGenerateParameters_genType_anyParams struct {
	Recorder *MoqGenerateParameters_genType_fnRecorder
}

// NewMoqGenerateParameters_genType creates a new moq of the
// GenerateParameters_genType type
func NewMoqGenerateParameters_genType(scene *moq.Scene, config *moq.Config) *MoqGenerateParameters_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqGenerateParameters_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqGenerateParameters_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Param1 moq.ParamIndexing
				Rand   moq.ParamIndexing
				Sizes  moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Param1 moq.ParamIndexing
			Rand   moq.ParamIndexing
			Sizes  moq.ParamIndexing
		}{
			Param1: moq.ParamIndexByHash,
			Rand:   moq.ParamIndexByHash,
			Sizes:  moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the GenerateParameters_genType type
func (m *MoqGenerateParameters_genType) Mock() GenerateParameters_genType {
	return func(param1 *dsa.Parameters, rand io.Reader, sizes dsa.ParameterSizes) error {
		m.Scene.T.Helper()
		moq := &MoqGenerateParameters_genType_mock{Moq: m}
		return moq.Fn(param1, rand, sizes)
	}
}

func (m *MoqGenerateParameters_genType_mock) Fn(param1 *dsa.Parameters, rand io.Reader, sizes dsa.ParameterSizes) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqGenerateParameters_genType_params{
		Param1: param1,
		Rand:   rand,
		Sizes:  sizes,
	}
	var results *MoqGenerateParameters_genType_results
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
		result.DoFn(param1, rand, sizes)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(param1, rand, sizes)
	}
	return
}

func (m *MoqGenerateParameters_genType) OnCall(param1 *dsa.Parameters, rand io.Reader, sizes dsa.ParameterSizes) *MoqGenerateParameters_genType_fnRecorder {
	return &MoqGenerateParameters_genType_fnRecorder{
		Params: MoqGenerateParameters_genType_params{
			Param1: param1,
			Rand:   rand,
			Sizes:  sizes,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqGenerateParameters_genType_fnRecorder) Any() *MoqGenerateParameters_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqGenerateParameters_genType_anyParams{Recorder: r}
}

func (a *MoqGenerateParameters_genType_anyParams) Param1() *MoqGenerateParameters_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqGenerateParameters_genType_anyParams) Rand() *MoqGenerateParameters_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqGenerateParameters_genType_anyParams) Sizes() *MoqGenerateParameters_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqGenerateParameters_genType_fnRecorder) Seq() *MoqGenerateParameters_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqGenerateParameters_genType_fnRecorder) NoSeq() *MoqGenerateParameters_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqGenerateParameters_genType_fnRecorder) ReturnResults(result1 error) *MoqGenerateParameters_genType_fnRecorder {
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
		DoFn       MoqGenerateParameters_genType_doFn
		DoReturnFn MoqGenerateParameters_genType_doReturnFn
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

func (r *MoqGenerateParameters_genType_fnRecorder) AndDo(fn MoqGenerateParameters_genType_doFn) *MoqGenerateParameters_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqGenerateParameters_genType_fnRecorder) DoReturnResults(fn MoqGenerateParameters_genType_doReturnFn) *MoqGenerateParameters_genType_fnRecorder {
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
		DoFn       MoqGenerateParameters_genType_doFn
		DoReturnFn MoqGenerateParameters_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqGenerateParameters_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqGenerateParameters_genType_resultsByParams
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
		results = &MoqGenerateParameters_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqGenerateParameters_genType_paramsKey]*MoqGenerateParameters_genType_results{},
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
		r.Results = &MoqGenerateParameters_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqGenerateParameters_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqGenerateParameters_genType_fnRecorder {
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
				DoFn       MoqGenerateParameters_genType_doFn
				DoReturnFn MoqGenerateParameters_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqGenerateParameters_genType) PrettyParams(params MoqGenerateParameters_genType_params) string {
	return fmt.Sprintf("GenerateParameters_genType(%#v, %#v, %#v)", params.Param1, params.Rand, params.Sizes)
}

func (m *MoqGenerateParameters_genType) ParamsKey(params MoqGenerateParameters_genType_params, anyParams uint64) MoqGenerateParameters_genType_paramsKey {
	m.Scene.T.Helper()
	var param1Used *dsa.Parameters
	var param1UsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Param1 == moq.ParamIndexByValue {
			param1Used = params.Param1
		} else {
			param1UsedHash = hash.DeepHash(params.Param1)
		}
	}
	var randUsed io.Reader
	var randUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Rand == moq.ParamIndexByValue {
			randUsed = params.Rand
		} else {
			randUsedHash = hash.DeepHash(params.Rand)
		}
	}
	var sizesUsed dsa.ParameterSizes
	var sizesUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Sizes == moq.ParamIndexByValue {
			sizesUsed = params.Sizes
		} else {
			sizesUsedHash = hash.DeepHash(params.Sizes)
		}
	}
	return MoqGenerateParameters_genType_paramsKey{
		Params: struct {
			Param1 *dsa.Parameters
			Rand   io.Reader
			Sizes  dsa.ParameterSizes
		}{
			Param1: param1Used,
			Rand:   randUsed,
			Sizes:  sizesUsed,
		},
		Hashes: struct {
			Param1 hash.Hash
			Rand   hash.Hash
			Sizes  hash.Hash
		}{
			Param1: param1UsedHash,
			Rand:   randUsedHash,
			Sizes:  sizesUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqGenerateParameters_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqGenerateParameters_genType) AssertExpectationsMet() {
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

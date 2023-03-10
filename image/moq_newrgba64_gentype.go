// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package image

import (
	"fmt"
	"image"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// NewRGBA64_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type NewRGBA64_genType func(r image.Rectangle) *image.RGBA64

// MoqNewRGBA64_genType holds the state of a moq of the NewRGBA64_genType type
type MoqNewRGBA64_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNewRGBA64_genType_mock

	ResultsByParams []MoqNewRGBA64_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Param1 moq.ParamIndexing
		}
	}
}

// MoqNewRGBA64_genType_mock isolates the mock interface of the
// NewRGBA64_genType type
type MoqNewRGBA64_genType_mock struct {
	Moq *MoqNewRGBA64_genType
}

// MoqNewRGBA64_genType_params holds the params of the NewRGBA64_genType type
type MoqNewRGBA64_genType_params struct{ Param1 image.Rectangle }

// MoqNewRGBA64_genType_paramsKey holds the map key params of the
// NewRGBA64_genType type
type MoqNewRGBA64_genType_paramsKey struct {
	Params struct{ Param1 image.Rectangle }
	Hashes struct{ Param1 hash.Hash }
}

// MoqNewRGBA64_genType_resultsByParams contains the results for a given set of
// parameters for the NewRGBA64_genType type
type MoqNewRGBA64_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNewRGBA64_genType_paramsKey]*MoqNewRGBA64_genType_results
}

// MoqNewRGBA64_genType_doFn defines the type of function needed when calling
// AndDo for the NewRGBA64_genType type
type MoqNewRGBA64_genType_doFn func(r image.Rectangle)

// MoqNewRGBA64_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the NewRGBA64_genType type
type MoqNewRGBA64_genType_doReturnFn func(r image.Rectangle) *image.RGBA64

// MoqNewRGBA64_genType_results holds the results of the NewRGBA64_genType type
type MoqNewRGBA64_genType_results struct {
	Params  MoqNewRGBA64_genType_params
	Results []struct {
		Values *struct {
			Result1 *image.RGBA64
		}
		Sequence   uint32
		DoFn       MoqNewRGBA64_genType_doFn
		DoReturnFn MoqNewRGBA64_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNewRGBA64_genType_fnRecorder routes recorded function calls to the
// MoqNewRGBA64_genType moq
type MoqNewRGBA64_genType_fnRecorder struct {
	Params    MoqNewRGBA64_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNewRGBA64_genType_results
	Moq       *MoqNewRGBA64_genType
}

// MoqNewRGBA64_genType_anyParams isolates the any params functions of the
// NewRGBA64_genType type
type MoqNewRGBA64_genType_anyParams struct {
	Recorder *MoqNewRGBA64_genType_fnRecorder
}

// NewMoqNewRGBA64_genType creates a new moq of the NewRGBA64_genType type
func NewMoqNewRGBA64_genType(scene *moq.Scene, config *moq.Config) *MoqNewRGBA64_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNewRGBA64_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNewRGBA64_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Param1 moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Param1 moq.ParamIndexing
		}{
			Param1: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the NewRGBA64_genType type
func (m *MoqNewRGBA64_genType) Mock() NewRGBA64_genType {
	return func(param1 image.Rectangle) *image.RGBA64 {
		m.Scene.T.Helper()
		moq := &MoqNewRGBA64_genType_mock{Moq: m}
		return moq.Fn(param1)
	}
}

func (m *MoqNewRGBA64_genType_mock) Fn(param1 image.Rectangle) (result1 *image.RGBA64) {
	m.Moq.Scene.T.Helper()
	params := MoqNewRGBA64_genType_params{
		Param1: param1,
	}
	var results *MoqNewRGBA64_genType_results
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
		result.DoFn(param1)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(param1)
	}
	return
}

func (m *MoqNewRGBA64_genType) OnCall(param1 image.Rectangle) *MoqNewRGBA64_genType_fnRecorder {
	return &MoqNewRGBA64_genType_fnRecorder{
		Params: MoqNewRGBA64_genType_params{
			Param1: param1,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNewRGBA64_genType_fnRecorder) Any() *MoqNewRGBA64_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNewRGBA64_genType_anyParams{Recorder: r}
}

func (a *MoqNewRGBA64_genType_anyParams) Param1() *MoqNewRGBA64_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqNewRGBA64_genType_fnRecorder) Seq() *MoqNewRGBA64_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNewRGBA64_genType_fnRecorder) NoSeq() *MoqNewRGBA64_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNewRGBA64_genType_fnRecorder) ReturnResults(result1 *image.RGBA64) *MoqNewRGBA64_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *image.RGBA64
		}
		Sequence   uint32
		DoFn       MoqNewRGBA64_genType_doFn
		DoReturnFn MoqNewRGBA64_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *image.RGBA64
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNewRGBA64_genType_fnRecorder) AndDo(fn MoqNewRGBA64_genType_doFn) *MoqNewRGBA64_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNewRGBA64_genType_fnRecorder) DoReturnResults(fn MoqNewRGBA64_genType_doReturnFn) *MoqNewRGBA64_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *image.RGBA64
		}
		Sequence   uint32
		DoFn       MoqNewRGBA64_genType_doFn
		DoReturnFn MoqNewRGBA64_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNewRGBA64_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNewRGBA64_genType_resultsByParams
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
		results = &MoqNewRGBA64_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNewRGBA64_genType_paramsKey]*MoqNewRGBA64_genType_results{},
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
		r.Results = &MoqNewRGBA64_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNewRGBA64_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNewRGBA64_genType_fnRecorder {
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
					Result1 *image.RGBA64
				}
				Sequence   uint32
				DoFn       MoqNewRGBA64_genType_doFn
				DoReturnFn MoqNewRGBA64_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNewRGBA64_genType) PrettyParams(params MoqNewRGBA64_genType_params) string {
	return fmt.Sprintf("NewRGBA64_genType(%#v)", params.Param1)
}

func (m *MoqNewRGBA64_genType) ParamsKey(params MoqNewRGBA64_genType_params, anyParams uint64) MoqNewRGBA64_genType_paramsKey {
	m.Scene.T.Helper()
	var param1Used image.Rectangle
	var param1UsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Param1 == moq.ParamIndexByValue {
			param1Used = params.Param1
		} else {
			param1UsedHash = hash.DeepHash(params.Param1)
		}
	}
	return MoqNewRGBA64_genType_paramsKey{
		Params: struct{ Param1 image.Rectangle }{
			Param1: param1Used,
		},
		Hashes: struct{ Param1 hash.Hash }{
			Param1: param1UsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNewRGBA64_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNewRGBA64_genType) AssertExpectationsMet() {
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

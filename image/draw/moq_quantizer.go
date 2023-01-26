// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package draw

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that draw.Quantizer is mocked
// completely
var _ draw.Quantizer = (*MoqQuantizer_mock)(nil)

// MoqQuantizer holds the state of a moq of the Quantizer type
type MoqQuantizer struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqQuantizer_mock

	ResultsByParams_Quantize []MoqQuantizer_Quantize_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Quantize struct {
				P      moq.ParamIndexing
				Param2 moq.ParamIndexing
			}
		}
	}
	// MoqQuantizer_mock isolates the mock interface of the Quantizer type
}

type MoqQuantizer_mock struct {
	Moq *MoqQuantizer
}

// MoqQuantizer_recorder isolates the recorder interface of the Quantizer type
type MoqQuantizer_recorder struct {
	Moq *MoqQuantizer
}

// MoqQuantizer_Quantize_params holds the params of the Quantizer type
type MoqQuantizer_Quantize_params struct {
	P      color.Palette
	Param2 image.Image
}

// MoqQuantizer_Quantize_paramsKey holds the map key params of the Quantizer
// type
type MoqQuantizer_Quantize_paramsKey struct {
	Params struct{ Param2 image.Image }
	Hashes struct {
		P      hash.Hash
		Param2 hash.Hash
	}
}

// MoqQuantizer_Quantize_resultsByParams contains the results for a given set
// of parameters for the Quantizer type
type MoqQuantizer_Quantize_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqQuantizer_Quantize_paramsKey]*MoqQuantizer_Quantize_results
}

// MoqQuantizer_Quantize_doFn defines the type of function needed when calling
// AndDo for the Quantizer type
type MoqQuantizer_Quantize_doFn func(p color.Palette, m image.Image)

// MoqQuantizer_Quantize_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Quantizer type
type MoqQuantizer_Quantize_doReturnFn func(p color.Palette, m image.Image) color.Palette

// MoqQuantizer_Quantize_results holds the results of the Quantizer type
type MoqQuantizer_Quantize_results struct {
	Params  MoqQuantizer_Quantize_params
	Results []struct {
		Values *struct {
			Result1 color.Palette
		}
		Sequence   uint32
		DoFn       MoqQuantizer_Quantize_doFn
		DoReturnFn MoqQuantizer_Quantize_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqQuantizer_Quantize_fnRecorder routes recorded function calls to the
// MoqQuantizer moq
type MoqQuantizer_Quantize_fnRecorder struct {
	Params    MoqQuantizer_Quantize_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqQuantizer_Quantize_results
	Moq       *MoqQuantizer
}

// MoqQuantizer_Quantize_anyParams isolates the any params functions of the
// Quantizer type
type MoqQuantizer_Quantize_anyParams struct {
	Recorder *MoqQuantizer_Quantize_fnRecorder
}

// NewMoqQuantizer creates a new moq of the Quantizer type
func NewMoqQuantizer(scene *moq.Scene, config *moq.Config) *MoqQuantizer {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqQuantizer{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqQuantizer_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Quantize struct {
					P      moq.ParamIndexing
					Param2 moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Quantize struct {
				P      moq.ParamIndexing
				Param2 moq.ParamIndexing
			}
		}{
			Quantize: struct {
				P      moq.ParamIndexing
				Param2 moq.ParamIndexing
			}{
				P:      moq.ParamIndexByHash,
				Param2: moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Quantizer type
func (m *MoqQuantizer) Mock() *MoqQuantizer_mock { return m.Moq }

func (m *MoqQuantizer_mock) Quantize(p color.Palette, param2 image.Image) (result1 color.Palette) {
	m.Moq.Scene.T.Helper()
	params := MoqQuantizer_Quantize_params{
		P:      p,
		Param2: param2,
	}
	var results *MoqQuantizer_Quantize_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Quantize {
		paramsKey := m.Moq.ParamsKey_Quantize(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Quantize(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Quantize(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Quantize(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(p, param2)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(p, param2)
	}
	return
}

// OnCall returns the recorder implementation of the Quantizer type
func (m *MoqQuantizer) OnCall() *MoqQuantizer_recorder {
	return &MoqQuantizer_recorder{
		Moq: m,
	}
}

func (m *MoqQuantizer_recorder) Quantize(p color.Palette, param2 image.Image) *MoqQuantizer_Quantize_fnRecorder {
	return &MoqQuantizer_Quantize_fnRecorder{
		Params: MoqQuantizer_Quantize_params{
			P:      p,
			Param2: param2,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqQuantizer_Quantize_fnRecorder) Any() *MoqQuantizer_Quantize_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Quantize(r.Params))
		return nil
	}
	return &MoqQuantizer_Quantize_anyParams{Recorder: r}
}

func (a *MoqQuantizer_Quantize_anyParams) P() *MoqQuantizer_Quantize_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqQuantizer_Quantize_anyParams) Param2() *MoqQuantizer_Quantize_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqQuantizer_Quantize_fnRecorder) Seq() *MoqQuantizer_Quantize_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Quantize(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqQuantizer_Quantize_fnRecorder) NoSeq() *MoqQuantizer_Quantize_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Quantize(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqQuantizer_Quantize_fnRecorder) ReturnResults(result1 color.Palette) *MoqQuantizer_Quantize_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 color.Palette
		}
		Sequence   uint32
		DoFn       MoqQuantizer_Quantize_doFn
		DoReturnFn MoqQuantizer_Quantize_doReturnFn
	}{
		Values: &struct {
			Result1 color.Palette
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqQuantizer_Quantize_fnRecorder) AndDo(fn MoqQuantizer_Quantize_doFn) *MoqQuantizer_Quantize_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqQuantizer_Quantize_fnRecorder) DoReturnResults(fn MoqQuantizer_Quantize_doReturnFn) *MoqQuantizer_Quantize_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 color.Palette
		}
		Sequence   uint32
		DoFn       MoqQuantizer_Quantize_doFn
		DoReturnFn MoqQuantizer_Quantize_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqQuantizer_Quantize_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqQuantizer_Quantize_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Quantize {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqQuantizer_Quantize_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqQuantizer_Quantize_paramsKey]*MoqQuantizer_Quantize_results{},
		}
		r.Moq.ResultsByParams_Quantize = append(r.Moq.ResultsByParams_Quantize, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Quantize) {
			copy(r.Moq.ResultsByParams_Quantize[insertAt+1:], r.Moq.ResultsByParams_Quantize[insertAt:0])
			r.Moq.ResultsByParams_Quantize[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Quantize(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqQuantizer_Quantize_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqQuantizer_Quantize_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqQuantizer_Quantize_fnRecorder {
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
					Result1 color.Palette
				}
				Sequence   uint32
				DoFn       MoqQuantizer_Quantize_doFn
				DoReturnFn MoqQuantizer_Quantize_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqQuantizer) PrettyParams_Quantize(params MoqQuantizer_Quantize_params) string {
	return fmt.Sprintf("Quantize(%#v, %#v)", params.P, params.Param2)
}

func (m *MoqQuantizer) ParamsKey_Quantize(params MoqQuantizer_Quantize_params, anyParams uint64) MoqQuantizer_Quantize_paramsKey {
	m.Scene.T.Helper()
	var pUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Quantize.P == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The p parameter of the Quantize function can't be indexed by value")
		}
		pUsedHash = hash.DeepHash(params.P)
	}
	var param2Used image.Image
	var param2UsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Quantize.Param2 == moq.ParamIndexByValue {
			param2Used = params.Param2
		} else {
			param2UsedHash = hash.DeepHash(params.Param2)
		}
	}
	return MoqQuantizer_Quantize_paramsKey{
		Params: struct{ Param2 image.Image }{
			Param2: param2Used,
		},
		Hashes: struct {
			P      hash.Hash
			Param2 hash.Hash
		}{
			P:      pUsedHash,
			Param2: param2UsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqQuantizer) Reset() { m.ResultsByParams_Quantize = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqQuantizer) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Quantize {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Quantize(results.Params))
			}
		}
	}
}

// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package jpeg

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Encode_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Encode_genType func(w io.Writer, m image.Image, o *jpeg.Options) error

// MoqEncode_genType holds the state of a moq of the Encode_genType type
type MoqEncode_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqEncode_genType_mock

	ResultsByParams []MoqEncode_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			W      moq.ParamIndexing
			Param2 moq.ParamIndexing
			O      moq.ParamIndexing
		}
	}
}

// MoqEncode_genType_mock isolates the mock interface of the Encode_genType
// type
type MoqEncode_genType_mock struct {
	Moq *MoqEncode_genType
}

// MoqEncode_genType_params holds the params of the Encode_genType type
type MoqEncode_genType_params struct {
	W      io.Writer
	Param2 image.Image
	O      *jpeg.Options
}

// MoqEncode_genType_paramsKey holds the map key params of the Encode_genType
// type
type MoqEncode_genType_paramsKey struct {
	Params struct {
		W      io.Writer
		Param2 image.Image
		O      *jpeg.Options
	}
	Hashes struct {
		W      hash.Hash
		Param2 hash.Hash
		O      hash.Hash
	}
}

// MoqEncode_genType_resultsByParams contains the results for a given set of
// parameters for the Encode_genType type
type MoqEncode_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqEncode_genType_paramsKey]*MoqEncode_genType_results
}

// MoqEncode_genType_doFn defines the type of function needed when calling
// AndDo for the Encode_genType type
type MoqEncode_genType_doFn func(w io.Writer, m image.Image, o *jpeg.Options)

// MoqEncode_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Encode_genType type
type MoqEncode_genType_doReturnFn func(w io.Writer, m image.Image, o *jpeg.Options) error

// MoqEncode_genType_results holds the results of the Encode_genType type
type MoqEncode_genType_results struct {
	Params  MoqEncode_genType_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqEncode_genType_doFn
		DoReturnFn MoqEncode_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqEncode_genType_fnRecorder routes recorded function calls to the
// MoqEncode_genType moq
type MoqEncode_genType_fnRecorder struct {
	Params    MoqEncode_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqEncode_genType_results
	Moq       *MoqEncode_genType
}

// MoqEncode_genType_anyParams isolates the any params functions of the
// Encode_genType type
type MoqEncode_genType_anyParams struct {
	Recorder *MoqEncode_genType_fnRecorder
}

// NewMoqEncode_genType creates a new moq of the Encode_genType type
func NewMoqEncode_genType(scene *moq.Scene, config *moq.Config) *MoqEncode_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqEncode_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqEncode_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				W      moq.ParamIndexing
				Param2 moq.ParamIndexing
				O      moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			W      moq.ParamIndexing
			Param2 moq.ParamIndexing
			O      moq.ParamIndexing
		}{
			W:      moq.ParamIndexByHash,
			Param2: moq.ParamIndexByHash,
			O:      moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Encode_genType type
func (m *MoqEncode_genType) Mock() Encode_genType {
	return func(w io.Writer, param2 image.Image, o *jpeg.Options) error {
		m.Scene.T.Helper()
		moq := &MoqEncode_genType_mock{Moq: m}
		return moq.Fn(w, param2, o)
	}
}

func (m *MoqEncode_genType_mock) Fn(w io.Writer, param2 image.Image, o *jpeg.Options) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqEncode_genType_params{
		W:      w,
		Param2: param2,
		O:      o,
	}
	var results *MoqEncode_genType_results
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
		result.DoFn(w, param2, o)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(w, param2, o)
	}
	return
}

func (m *MoqEncode_genType) OnCall(w io.Writer, param2 image.Image, o *jpeg.Options) *MoqEncode_genType_fnRecorder {
	return &MoqEncode_genType_fnRecorder{
		Params: MoqEncode_genType_params{
			W:      w,
			Param2: param2,
			O:      o,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqEncode_genType_fnRecorder) Any() *MoqEncode_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqEncode_genType_anyParams{Recorder: r}
}

func (a *MoqEncode_genType_anyParams) W() *MoqEncode_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqEncode_genType_anyParams) Param2() *MoqEncode_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqEncode_genType_anyParams) O() *MoqEncode_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqEncode_genType_fnRecorder) Seq() *MoqEncode_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqEncode_genType_fnRecorder) NoSeq() *MoqEncode_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqEncode_genType_fnRecorder) ReturnResults(result1 error) *MoqEncode_genType_fnRecorder {
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
		DoFn       MoqEncode_genType_doFn
		DoReturnFn MoqEncode_genType_doReturnFn
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

func (r *MoqEncode_genType_fnRecorder) AndDo(fn MoqEncode_genType_doFn) *MoqEncode_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqEncode_genType_fnRecorder) DoReturnResults(fn MoqEncode_genType_doReturnFn) *MoqEncode_genType_fnRecorder {
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
		DoFn       MoqEncode_genType_doFn
		DoReturnFn MoqEncode_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqEncode_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqEncode_genType_resultsByParams
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
		results = &MoqEncode_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqEncode_genType_paramsKey]*MoqEncode_genType_results{},
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
		r.Results = &MoqEncode_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqEncode_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqEncode_genType_fnRecorder {
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
				DoFn       MoqEncode_genType_doFn
				DoReturnFn MoqEncode_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqEncode_genType) PrettyParams(params MoqEncode_genType_params) string {
	return fmt.Sprintf("Encode_genType(%#v, %#v, %#v)", params.W, params.Param2, params.O)
}

func (m *MoqEncode_genType) ParamsKey(params MoqEncode_genType_params, anyParams uint64) MoqEncode_genType_paramsKey {
	m.Scene.T.Helper()
	var wUsed io.Writer
	var wUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.W == moq.ParamIndexByValue {
			wUsed = params.W
		} else {
			wUsedHash = hash.DeepHash(params.W)
		}
	}
	var param2Used image.Image
	var param2UsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Param2 == moq.ParamIndexByValue {
			param2Used = params.Param2
		} else {
			param2UsedHash = hash.DeepHash(params.Param2)
		}
	}
	var oUsed *jpeg.Options
	var oUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.O == moq.ParamIndexByValue {
			oUsed = params.O
		} else {
			oUsedHash = hash.DeepHash(params.O)
		}
	}
	return MoqEncode_genType_paramsKey{
		Params: struct {
			W      io.Writer
			Param2 image.Image
			O      *jpeg.Options
		}{
			W:      wUsed,
			Param2: param2Used,
			O:      oUsed,
		},
		Hashes: struct {
			W      hash.Hash
			Param2 hash.Hash
			O      hash.Hash
		}{
			W:      wUsedHash,
			Param2: param2UsedHash,
			O:      oUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqEncode_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqEncode_genType) AssertExpectationsMet() {
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
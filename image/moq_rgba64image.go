// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package image

import (
	"fmt"
	"image"
	"image/color"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that image.RGBA64Image is mocked
// completely
var _ image.RGBA64Image = (*MoqRGBA64Image_mock)(nil)

// MoqRGBA64Image holds the state of a moq of the RGBA64Image type
type MoqRGBA64Image struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqRGBA64Image_mock

	ResultsByParams_RGBA64At   []MoqRGBA64Image_RGBA64At_resultsByParams
	ResultsByParams_ColorModel []MoqRGBA64Image_ColorModel_resultsByParams
	ResultsByParams_Bounds     []MoqRGBA64Image_Bounds_resultsByParams
	ResultsByParams_At         []MoqRGBA64Image_At_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			RGBA64At struct {
				X moq.ParamIndexing
				Y moq.ParamIndexing
			}
			ColorModel struct{}
			Bounds     struct{}
			At         struct {
				X moq.ParamIndexing
				Y moq.ParamIndexing
			}
		}
	}
	// MoqRGBA64Image_mock isolates the mock interface of the RGBA64Image type
}

type MoqRGBA64Image_mock struct {
	Moq *MoqRGBA64Image
}

// MoqRGBA64Image_recorder isolates the recorder interface of the RGBA64Image
// type
type MoqRGBA64Image_recorder struct {
	Moq *MoqRGBA64Image
}

// MoqRGBA64Image_RGBA64At_params holds the params of the RGBA64Image type
type MoqRGBA64Image_RGBA64At_params struct{ X, Y int }

// MoqRGBA64Image_RGBA64At_paramsKey holds the map key params of the
// RGBA64Image type
type MoqRGBA64Image_RGBA64At_paramsKey struct {
	Params struct{ X, Y int }
	Hashes struct{ X, Y hash.Hash }
}

// MoqRGBA64Image_RGBA64At_resultsByParams contains the results for a given set
// of parameters for the RGBA64Image type
type MoqRGBA64Image_RGBA64At_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqRGBA64Image_RGBA64At_paramsKey]*MoqRGBA64Image_RGBA64At_results
}

// MoqRGBA64Image_RGBA64At_doFn defines the type of function needed when
// calling AndDo for the RGBA64Image type
type MoqRGBA64Image_RGBA64At_doFn func(x, y int)

// MoqRGBA64Image_RGBA64At_doReturnFn defines the type of function needed when
// calling DoReturnResults for the RGBA64Image type
type MoqRGBA64Image_RGBA64At_doReturnFn func(x, y int) color.RGBA64

// MoqRGBA64Image_RGBA64At_results holds the results of the RGBA64Image type
type MoqRGBA64Image_RGBA64At_results struct {
	Params  MoqRGBA64Image_RGBA64At_params
	Results []struct {
		Values *struct {
			Result1 color.RGBA64
		}
		Sequence   uint32
		DoFn       MoqRGBA64Image_RGBA64At_doFn
		DoReturnFn MoqRGBA64Image_RGBA64At_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqRGBA64Image_RGBA64At_fnRecorder routes recorded function calls to the
// MoqRGBA64Image moq
type MoqRGBA64Image_RGBA64At_fnRecorder struct {
	Params    MoqRGBA64Image_RGBA64At_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqRGBA64Image_RGBA64At_results
	Moq       *MoqRGBA64Image
}

// MoqRGBA64Image_RGBA64At_anyParams isolates the any params functions of the
// RGBA64Image type
type MoqRGBA64Image_RGBA64At_anyParams struct {
	Recorder *MoqRGBA64Image_RGBA64At_fnRecorder
}

// MoqRGBA64Image_ColorModel_params holds the params of the RGBA64Image type
type MoqRGBA64Image_ColorModel_params struct{}

// MoqRGBA64Image_ColorModel_paramsKey holds the map key params of the
// RGBA64Image type
type MoqRGBA64Image_ColorModel_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqRGBA64Image_ColorModel_resultsByParams contains the results for a given
// set of parameters for the RGBA64Image type
type MoqRGBA64Image_ColorModel_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqRGBA64Image_ColorModel_paramsKey]*MoqRGBA64Image_ColorModel_results
}

// MoqRGBA64Image_ColorModel_doFn defines the type of function needed when
// calling AndDo for the RGBA64Image type
type MoqRGBA64Image_ColorModel_doFn func()

// MoqRGBA64Image_ColorModel_doReturnFn defines the type of function needed
// when calling DoReturnResults for the RGBA64Image type
type MoqRGBA64Image_ColorModel_doReturnFn func() color.Model

// MoqRGBA64Image_ColorModel_results holds the results of the RGBA64Image type
type MoqRGBA64Image_ColorModel_results struct {
	Params  MoqRGBA64Image_ColorModel_params
	Results []struct {
		Values *struct {
			Result1 color.Model
		}
		Sequence   uint32
		DoFn       MoqRGBA64Image_ColorModel_doFn
		DoReturnFn MoqRGBA64Image_ColorModel_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqRGBA64Image_ColorModel_fnRecorder routes recorded function calls to the
// MoqRGBA64Image moq
type MoqRGBA64Image_ColorModel_fnRecorder struct {
	Params    MoqRGBA64Image_ColorModel_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqRGBA64Image_ColorModel_results
	Moq       *MoqRGBA64Image
}

// MoqRGBA64Image_ColorModel_anyParams isolates the any params functions of the
// RGBA64Image type
type MoqRGBA64Image_ColorModel_anyParams struct {
	Recorder *MoqRGBA64Image_ColorModel_fnRecorder
}

// MoqRGBA64Image_Bounds_params holds the params of the RGBA64Image type
type MoqRGBA64Image_Bounds_params struct{}

// MoqRGBA64Image_Bounds_paramsKey holds the map key params of the RGBA64Image
// type
type MoqRGBA64Image_Bounds_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqRGBA64Image_Bounds_resultsByParams contains the results for a given set
// of parameters for the RGBA64Image type
type MoqRGBA64Image_Bounds_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqRGBA64Image_Bounds_paramsKey]*MoqRGBA64Image_Bounds_results
}

// MoqRGBA64Image_Bounds_doFn defines the type of function needed when calling
// AndDo for the RGBA64Image type
type MoqRGBA64Image_Bounds_doFn func()

// MoqRGBA64Image_Bounds_doReturnFn defines the type of function needed when
// calling DoReturnResults for the RGBA64Image type
type MoqRGBA64Image_Bounds_doReturnFn func() image.Rectangle

// MoqRGBA64Image_Bounds_results holds the results of the RGBA64Image type
type MoqRGBA64Image_Bounds_results struct {
	Params  MoqRGBA64Image_Bounds_params
	Results []struct {
		Values *struct {
			Result1 image.Rectangle
		}
		Sequence   uint32
		DoFn       MoqRGBA64Image_Bounds_doFn
		DoReturnFn MoqRGBA64Image_Bounds_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqRGBA64Image_Bounds_fnRecorder routes recorded function calls to the
// MoqRGBA64Image moq
type MoqRGBA64Image_Bounds_fnRecorder struct {
	Params    MoqRGBA64Image_Bounds_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqRGBA64Image_Bounds_results
	Moq       *MoqRGBA64Image
}

// MoqRGBA64Image_Bounds_anyParams isolates the any params functions of the
// RGBA64Image type
type MoqRGBA64Image_Bounds_anyParams struct {
	Recorder *MoqRGBA64Image_Bounds_fnRecorder
}

// MoqRGBA64Image_At_params holds the params of the RGBA64Image type
type MoqRGBA64Image_At_params struct{ X, Y int }

// MoqRGBA64Image_At_paramsKey holds the map key params of the RGBA64Image type
type MoqRGBA64Image_At_paramsKey struct {
	Params struct{ X, Y int }
	Hashes struct{ X, Y hash.Hash }
}

// MoqRGBA64Image_At_resultsByParams contains the results for a given set of
// parameters for the RGBA64Image type
type MoqRGBA64Image_At_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqRGBA64Image_At_paramsKey]*MoqRGBA64Image_At_results
}

// MoqRGBA64Image_At_doFn defines the type of function needed when calling
// AndDo for the RGBA64Image type
type MoqRGBA64Image_At_doFn func(x, y int)

// MoqRGBA64Image_At_doReturnFn defines the type of function needed when
// calling DoReturnResults for the RGBA64Image type
type MoqRGBA64Image_At_doReturnFn func(x, y int) color.Color

// MoqRGBA64Image_At_results holds the results of the RGBA64Image type
type MoqRGBA64Image_At_results struct {
	Params  MoqRGBA64Image_At_params
	Results []struct {
		Values *struct {
			Result1 color.Color
		}
		Sequence   uint32
		DoFn       MoqRGBA64Image_At_doFn
		DoReturnFn MoqRGBA64Image_At_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqRGBA64Image_At_fnRecorder routes recorded function calls to the
// MoqRGBA64Image moq
type MoqRGBA64Image_At_fnRecorder struct {
	Params    MoqRGBA64Image_At_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqRGBA64Image_At_results
	Moq       *MoqRGBA64Image
}

// MoqRGBA64Image_At_anyParams isolates the any params functions of the
// RGBA64Image type
type MoqRGBA64Image_At_anyParams struct {
	Recorder *MoqRGBA64Image_At_fnRecorder
}

// NewMoqRGBA64Image creates a new moq of the RGBA64Image type
func NewMoqRGBA64Image(scene *moq.Scene, config *moq.Config) *MoqRGBA64Image {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqRGBA64Image{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqRGBA64Image_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				RGBA64At struct {
					X moq.ParamIndexing
					Y moq.ParamIndexing
				}
				ColorModel struct{}
				Bounds     struct{}
				At         struct {
					X moq.ParamIndexing
					Y moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			RGBA64At struct {
				X moq.ParamIndexing
				Y moq.ParamIndexing
			}
			ColorModel struct{}
			Bounds     struct{}
			At         struct {
				X moq.ParamIndexing
				Y moq.ParamIndexing
			}
		}{
			RGBA64At: struct {
				X moq.ParamIndexing
				Y moq.ParamIndexing
			}{
				X: moq.ParamIndexByValue,
				Y: moq.ParamIndexByValue,
			},
			ColorModel: struct{}{},
			Bounds:     struct{}{},
			At: struct {
				X moq.ParamIndexing
				Y moq.ParamIndexing
			}{
				X: moq.ParamIndexByValue,
				Y: moq.ParamIndexByValue,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the RGBA64Image type
func (m *MoqRGBA64Image) Mock() *MoqRGBA64Image_mock { return m.Moq }

func (m *MoqRGBA64Image_mock) RGBA64At(x, y int) (result1 color.RGBA64) {
	m.Moq.Scene.T.Helper()
	params := MoqRGBA64Image_RGBA64At_params{
		X: x,
		Y: y,
	}
	var results *MoqRGBA64Image_RGBA64At_results
	for _, resultsByParams := range m.Moq.ResultsByParams_RGBA64At {
		paramsKey := m.Moq.ParamsKey_RGBA64At(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_RGBA64At(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_RGBA64At(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_RGBA64At(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(x, y)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(x, y)
	}
	return
}

func (m *MoqRGBA64Image_mock) ColorModel() (result1 color.Model) {
	m.Moq.Scene.T.Helper()
	params := MoqRGBA64Image_ColorModel_params{}
	var results *MoqRGBA64Image_ColorModel_results
	for _, resultsByParams := range m.Moq.ResultsByParams_ColorModel {
		paramsKey := m.Moq.ParamsKey_ColorModel(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_ColorModel(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_ColorModel(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_ColorModel(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn()
	}
	return
}

func (m *MoqRGBA64Image_mock) Bounds() (result1 image.Rectangle) {
	m.Moq.Scene.T.Helper()
	params := MoqRGBA64Image_Bounds_params{}
	var results *MoqRGBA64Image_Bounds_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Bounds {
		paramsKey := m.Moq.ParamsKey_Bounds(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Bounds(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Bounds(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Bounds(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn()
	}
	return
}

func (m *MoqRGBA64Image_mock) At(x, y int) (result1 color.Color) {
	m.Moq.Scene.T.Helper()
	params := MoqRGBA64Image_At_params{
		X: x,
		Y: y,
	}
	var results *MoqRGBA64Image_At_results
	for _, resultsByParams := range m.Moq.ResultsByParams_At {
		paramsKey := m.Moq.ParamsKey_At(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_At(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_At(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_At(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(x, y)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(x, y)
	}
	return
}

// OnCall returns the recorder implementation of the RGBA64Image type
func (m *MoqRGBA64Image) OnCall() *MoqRGBA64Image_recorder {
	return &MoqRGBA64Image_recorder{
		Moq: m,
	}
}

func (m *MoqRGBA64Image_recorder) RGBA64At(x, y int) *MoqRGBA64Image_RGBA64At_fnRecorder {
	return &MoqRGBA64Image_RGBA64At_fnRecorder{
		Params: MoqRGBA64Image_RGBA64At_params{
			X: x,
			Y: y,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqRGBA64Image_RGBA64At_fnRecorder) Any() *MoqRGBA64Image_RGBA64At_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_RGBA64At(r.Params))
		return nil
	}
	return &MoqRGBA64Image_RGBA64At_anyParams{Recorder: r}
}

func (a *MoqRGBA64Image_RGBA64At_anyParams) X() *MoqRGBA64Image_RGBA64At_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqRGBA64Image_RGBA64At_anyParams) Y() *MoqRGBA64Image_RGBA64At_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqRGBA64Image_RGBA64At_fnRecorder) Seq() *MoqRGBA64Image_RGBA64At_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_RGBA64At(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqRGBA64Image_RGBA64At_fnRecorder) NoSeq() *MoqRGBA64Image_RGBA64At_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_RGBA64At(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqRGBA64Image_RGBA64At_fnRecorder) ReturnResults(result1 color.RGBA64) *MoqRGBA64Image_RGBA64At_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 color.RGBA64
		}
		Sequence   uint32
		DoFn       MoqRGBA64Image_RGBA64At_doFn
		DoReturnFn MoqRGBA64Image_RGBA64At_doReturnFn
	}{
		Values: &struct {
			Result1 color.RGBA64
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqRGBA64Image_RGBA64At_fnRecorder) AndDo(fn MoqRGBA64Image_RGBA64At_doFn) *MoqRGBA64Image_RGBA64At_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqRGBA64Image_RGBA64At_fnRecorder) DoReturnResults(fn MoqRGBA64Image_RGBA64At_doReturnFn) *MoqRGBA64Image_RGBA64At_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 color.RGBA64
		}
		Sequence   uint32
		DoFn       MoqRGBA64Image_RGBA64At_doFn
		DoReturnFn MoqRGBA64Image_RGBA64At_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqRGBA64Image_RGBA64At_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqRGBA64Image_RGBA64At_resultsByParams
	for n, res := range r.Moq.ResultsByParams_RGBA64At {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqRGBA64Image_RGBA64At_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqRGBA64Image_RGBA64At_paramsKey]*MoqRGBA64Image_RGBA64At_results{},
		}
		r.Moq.ResultsByParams_RGBA64At = append(r.Moq.ResultsByParams_RGBA64At, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_RGBA64At) {
			copy(r.Moq.ResultsByParams_RGBA64At[insertAt+1:], r.Moq.ResultsByParams_RGBA64At[insertAt:0])
			r.Moq.ResultsByParams_RGBA64At[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_RGBA64At(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqRGBA64Image_RGBA64At_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqRGBA64Image_RGBA64At_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqRGBA64Image_RGBA64At_fnRecorder {
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
					Result1 color.RGBA64
				}
				Sequence   uint32
				DoFn       MoqRGBA64Image_RGBA64At_doFn
				DoReturnFn MoqRGBA64Image_RGBA64At_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqRGBA64Image) PrettyParams_RGBA64At(params MoqRGBA64Image_RGBA64At_params) string {
	return fmt.Sprintf("RGBA64At(%#v, %#v)", params.X, params.Y)
}

func (m *MoqRGBA64Image) ParamsKey_RGBA64At(params MoqRGBA64Image_RGBA64At_params, anyParams uint64) MoqRGBA64Image_RGBA64At_paramsKey {
	m.Scene.T.Helper()
	var xUsed int
	var xUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.RGBA64At.X == moq.ParamIndexByValue {
			xUsed = params.X
		} else {
			xUsedHash = hash.DeepHash(params.X)
		}
	}
	var yUsed int
	var yUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.RGBA64At.Y == moq.ParamIndexByValue {
			yUsed = params.Y
		} else {
			yUsedHash = hash.DeepHash(params.Y)
		}
	}
	return MoqRGBA64Image_RGBA64At_paramsKey{
		Params: struct{ X, Y int }{
			X: xUsed,
			Y: yUsed,
		},
		Hashes: struct{ X, Y hash.Hash }{
			X: xUsedHash,
			Y: yUsedHash,
		},
	}
}

func (m *MoqRGBA64Image_recorder) ColorModel() *MoqRGBA64Image_ColorModel_fnRecorder {
	return &MoqRGBA64Image_ColorModel_fnRecorder{
		Params:   MoqRGBA64Image_ColorModel_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqRGBA64Image_ColorModel_fnRecorder) Any() *MoqRGBA64Image_ColorModel_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ColorModel(r.Params))
		return nil
	}
	return &MoqRGBA64Image_ColorModel_anyParams{Recorder: r}
}

func (r *MoqRGBA64Image_ColorModel_fnRecorder) Seq() *MoqRGBA64Image_ColorModel_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ColorModel(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqRGBA64Image_ColorModel_fnRecorder) NoSeq() *MoqRGBA64Image_ColorModel_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ColorModel(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqRGBA64Image_ColorModel_fnRecorder) ReturnResults(result1 color.Model) *MoqRGBA64Image_ColorModel_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 color.Model
		}
		Sequence   uint32
		DoFn       MoqRGBA64Image_ColorModel_doFn
		DoReturnFn MoqRGBA64Image_ColorModel_doReturnFn
	}{
		Values: &struct {
			Result1 color.Model
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqRGBA64Image_ColorModel_fnRecorder) AndDo(fn MoqRGBA64Image_ColorModel_doFn) *MoqRGBA64Image_ColorModel_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqRGBA64Image_ColorModel_fnRecorder) DoReturnResults(fn MoqRGBA64Image_ColorModel_doReturnFn) *MoqRGBA64Image_ColorModel_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 color.Model
		}
		Sequence   uint32
		DoFn       MoqRGBA64Image_ColorModel_doFn
		DoReturnFn MoqRGBA64Image_ColorModel_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqRGBA64Image_ColorModel_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqRGBA64Image_ColorModel_resultsByParams
	for n, res := range r.Moq.ResultsByParams_ColorModel {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqRGBA64Image_ColorModel_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqRGBA64Image_ColorModel_paramsKey]*MoqRGBA64Image_ColorModel_results{},
		}
		r.Moq.ResultsByParams_ColorModel = append(r.Moq.ResultsByParams_ColorModel, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_ColorModel) {
			copy(r.Moq.ResultsByParams_ColorModel[insertAt+1:], r.Moq.ResultsByParams_ColorModel[insertAt:0])
			r.Moq.ResultsByParams_ColorModel[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_ColorModel(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqRGBA64Image_ColorModel_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqRGBA64Image_ColorModel_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqRGBA64Image_ColorModel_fnRecorder {
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
					Result1 color.Model
				}
				Sequence   uint32
				DoFn       MoqRGBA64Image_ColorModel_doFn
				DoReturnFn MoqRGBA64Image_ColorModel_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqRGBA64Image) PrettyParams_ColorModel(params MoqRGBA64Image_ColorModel_params) string {
	return fmt.Sprintf("ColorModel()")
}

func (m *MoqRGBA64Image) ParamsKey_ColorModel(params MoqRGBA64Image_ColorModel_params, anyParams uint64) MoqRGBA64Image_ColorModel_paramsKey {
	m.Scene.T.Helper()
	return MoqRGBA64Image_ColorModel_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqRGBA64Image_recorder) Bounds() *MoqRGBA64Image_Bounds_fnRecorder {
	return &MoqRGBA64Image_Bounds_fnRecorder{
		Params:   MoqRGBA64Image_Bounds_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqRGBA64Image_Bounds_fnRecorder) Any() *MoqRGBA64Image_Bounds_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Bounds(r.Params))
		return nil
	}
	return &MoqRGBA64Image_Bounds_anyParams{Recorder: r}
}

func (r *MoqRGBA64Image_Bounds_fnRecorder) Seq() *MoqRGBA64Image_Bounds_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Bounds(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqRGBA64Image_Bounds_fnRecorder) NoSeq() *MoqRGBA64Image_Bounds_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Bounds(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqRGBA64Image_Bounds_fnRecorder) ReturnResults(result1 image.Rectangle) *MoqRGBA64Image_Bounds_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 image.Rectangle
		}
		Sequence   uint32
		DoFn       MoqRGBA64Image_Bounds_doFn
		DoReturnFn MoqRGBA64Image_Bounds_doReturnFn
	}{
		Values: &struct {
			Result1 image.Rectangle
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqRGBA64Image_Bounds_fnRecorder) AndDo(fn MoqRGBA64Image_Bounds_doFn) *MoqRGBA64Image_Bounds_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqRGBA64Image_Bounds_fnRecorder) DoReturnResults(fn MoqRGBA64Image_Bounds_doReturnFn) *MoqRGBA64Image_Bounds_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 image.Rectangle
		}
		Sequence   uint32
		DoFn       MoqRGBA64Image_Bounds_doFn
		DoReturnFn MoqRGBA64Image_Bounds_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqRGBA64Image_Bounds_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqRGBA64Image_Bounds_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Bounds {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqRGBA64Image_Bounds_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqRGBA64Image_Bounds_paramsKey]*MoqRGBA64Image_Bounds_results{},
		}
		r.Moq.ResultsByParams_Bounds = append(r.Moq.ResultsByParams_Bounds, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Bounds) {
			copy(r.Moq.ResultsByParams_Bounds[insertAt+1:], r.Moq.ResultsByParams_Bounds[insertAt:0])
			r.Moq.ResultsByParams_Bounds[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Bounds(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqRGBA64Image_Bounds_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqRGBA64Image_Bounds_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqRGBA64Image_Bounds_fnRecorder {
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
					Result1 image.Rectangle
				}
				Sequence   uint32
				DoFn       MoqRGBA64Image_Bounds_doFn
				DoReturnFn MoqRGBA64Image_Bounds_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqRGBA64Image) PrettyParams_Bounds(params MoqRGBA64Image_Bounds_params) string {
	return fmt.Sprintf("Bounds()")
}

func (m *MoqRGBA64Image) ParamsKey_Bounds(params MoqRGBA64Image_Bounds_params, anyParams uint64) MoqRGBA64Image_Bounds_paramsKey {
	m.Scene.T.Helper()
	return MoqRGBA64Image_Bounds_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqRGBA64Image_recorder) At(x, y int) *MoqRGBA64Image_At_fnRecorder {
	return &MoqRGBA64Image_At_fnRecorder{
		Params: MoqRGBA64Image_At_params{
			X: x,
			Y: y,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqRGBA64Image_At_fnRecorder) Any() *MoqRGBA64Image_At_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_At(r.Params))
		return nil
	}
	return &MoqRGBA64Image_At_anyParams{Recorder: r}
}

func (a *MoqRGBA64Image_At_anyParams) X() *MoqRGBA64Image_At_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqRGBA64Image_At_anyParams) Y() *MoqRGBA64Image_At_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqRGBA64Image_At_fnRecorder) Seq() *MoqRGBA64Image_At_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_At(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqRGBA64Image_At_fnRecorder) NoSeq() *MoqRGBA64Image_At_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_At(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqRGBA64Image_At_fnRecorder) ReturnResults(result1 color.Color) *MoqRGBA64Image_At_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 color.Color
		}
		Sequence   uint32
		DoFn       MoqRGBA64Image_At_doFn
		DoReturnFn MoqRGBA64Image_At_doReturnFn
	}{
		Values: &struct {
			Result1 color.Color
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqRGBA64Image_At_fnRecorder) AndDo(fn MoqRGBA64Image_At_doFn) *MoqRGBA64Image_At_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqRGBA64Image_At_fnRecorder) DoReturnResults(fn MoqRGBA64Image_At_doReturnFn) *MoqRGBA64Image_At_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 color.Color
		}
		Sequence   uint32
		DoFn       MoqRGBA64Image_At_doFn
		DoReturnFn MoqRGBA64Image_At_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqRGBA64Image_At_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqRGBA64Image_At_resultsByParams
	for n, res := range r.Moq.ResultsByParams_At {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqRGBA64Image_At_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqRGBA64Image_At_paramsKey]*MoqRGBA64Image_At_results{},
		}
		r.Moq.ResultsByParams_At = append(r.Moq.ResultsByParams_At, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_At) {
			copy(r.Moq.ResultsByParams_At[insertAt+1:], r.Moq.ResultsByParams_At[insertAt:0])
			r.Moq.ResultsByParams_At[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_At(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqRGBA64Image_At_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqRGBA64Image_At_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqRGBA64Image_At_fnRecorder {
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
					Result1 color.Color
				}
				Sequence   uint32
				DoFn       MoqRGBA64Image_At_doFn
				DoReturnFn MoqRGBA64Image_At_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqRGBA64Image) PrettyParams_At(params MoqRGBA64Image_At_params) string {
	return fmt.Sprintf("At(%#v, %#v)", params.X, params.Y)
}

func (m *MoqRGBA64Image) ParamsKey_At(params MoqRGBA64Image_At_params, anyParams uint64) MoqRGBA64Image_At_paramsKey {
	m.Scene.T.Helper()
	var xUsed int
	var xUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.At.X == moq.ParamIndexByValue {
			xUsed = params.X
		} else {
			xUsedHash = hash.DeepHash(params.X)
		}
	}
	var yUsed int
	var yUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.At.Y == moq.ParamIndexByValue {
			yUsed = params.Y
		} else {
			yUsedHash = hash.DeepHash(params.Y)
		}
	}
	return MoqRGBA64Image_At_paramsKey{
		Params: struct{ X, Y int }{
			X: xUsed,
			Y: yUsed,
		},
		Hashes: struct{ X, Y hash.Hash }{
			X: xUsedHash,
			Y: yUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqRGBA64Image) Reset() {
	m.ResultsByParams_RGBA64At = nil
	m.ResultsByParams_ColorModel = nil
	m.ResultsByParams_Bounds = nil
	m.ResultsByParams_At = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqRGBA64Image) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_RGBA64At {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_RGBA64At(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_ColorModel {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_ColorModel(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Bounds {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Bounds(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_At {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_At(results.Params))
			}
		}
	}
}

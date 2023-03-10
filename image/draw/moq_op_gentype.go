// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package draw

import (
	"fmt"
	"image"
	"image/draw"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that draw.Op_genType is mocked
// completely
var _ Op_genType = (*MoqOp_genType_mock)(nil)

// Op_genType is the fabricated implementation type of this mock (emitted when
// mocking a collections of methods directly and not from an interface type)
type Op_genType interface {
	Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point)
}

// MoqOp_genType holds the state of a moq of the Op_genType type
type MoqOp_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqOp_genType_mock

	ResultsByParams_Draw []MoqOp_genType_Draw_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Draw struct {
				Dst    moq.ParamIndexing
				Param2 moq.ParamIndexing
				Src    moq.ParamIndexing
				Sp     moq.ParamIndexing
			}
		}
	}
	// MoqOp_genType_mock isolates the mock interface of the Op_genType type
}

type MoqOp_genType_mock struct {
	Moq *MoqOp_genType
}

// MoqOp_genType_recorder isolates the recorder interface of the Op_genType
// type
type MoqOp_genType_recorder struct {
	Moq *MoqOp_genType
}

// MoqOp_genType_Draw_params holds the params of the Op_genType type
type MoqOp_genType_Draw_params struct {
	Dst    draw.Image
	Param2 image.Rectangle
	Src    image.Image
	Sp     image.Point
}

// MoqOp_genType_Draw_paramsKey holds the map key params of the Op_genType type
type MoqOp_genType_Draw_paramsKey struct {
	Params struct {
		Dst    draw.Image
		Param2 image.Rectangle
		Src    image.Image
		Sp     image.Point
	}
	Hashes struct {
		Dst    hash.Hash
		Param2 hash.Hash
		Src    hash.Hash
		Sp     hash.Hash
	}
}

// MoqOp_genType_Draw_resultsByParams contains the results for a given set of
// parameters for the Op_genType type
type MoqOp_genType_Draw_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqOp_genType_Draw_paramsKey]*MoqOp_genType_Draw_results
}

// MoqOp_genType_Draw_doFn defines the type of function needed when calling
// AndDo for the Op_genType type
type MoqOp_genType_Draw_doFn func(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point)

// MoqOp_genType_Draw_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Op_genType type
type MoqOp_genType_Draw_doReturnFn func(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point)

// MoqOp_genType_Draw_results holds the results of the Op_genType type
type MoqOp_genType_Draw_results struct {
	Params  MoqOp_genType_Draw_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqOp_genType_Draw_doFn
		DoReturnFn MoqOp_genType_Draw_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqOp_genType_Draw_fnRecorder routes recorded function calls to the
// MoqOp_genType moq
type MoqOp_genType_Draw_fnRecorder struct {
	Params    MoqOp_genType_Draw_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqOp_genType_Draw_results
	Moq       *MoqOp_genType
}

// MoqOp_genType_Draw_anyParams isolates the any params functions of the
// Op_genType type
type MoqOp_genType_Draw_anyParams struct {
	Recorder *MoqOp_genType_Draw_fnRecorder
}

// NewMoqOp_genType creates a new moq of the Op_genType type
func NewMoqOp_genType(scene *moq.Scene, config *moq.Config) *MoqOp_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqOp_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqOp_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Draw struct {
					Dst    moq.ParamIndexing
					Param2 moq.ParamIndexing
					Src    moq.ParamIndexing
					Sp     moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Draw struct {
				Dst    moq.ParamIndexing
				Param2 moq.ParamIndexing
				Src    moq.ParamIndexing
				Sp     moq.ParamIndexing
			}
		}{
			Draw: struct {
				Dst    moq.ParamIndexing
				Param2 moq.ParamIndexing
				Src    moq.ParamIndexing
				Sp     moq.ParamIndexing
			}{
				Dst:    moq.ParamIndexByHash,
				Param2: moq.ParamIndexByValue,
				Src:    moq.ParamIndexByHash,
				Sp:     moq.ParamIndexByValue,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Op_genType type
func (m *MoqOp_genType) Mock() *MoqOp_genType_mock { return m.Moq }

func (m *MoqOp_genType_mock) Draw(dst draw.Image, param2 image.Rectangle, src image.Image, sp image.Point) {
	m.Moq.Scene.T.Helper()
	params := MoqOp_genType_Draw_params{
		Dst:    dst,
		Param2: param2,
		Src:    src,
		Sp:     sp,
	}
	var results *MoqOp_genType_Draw_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Draw {
		paramsKey := m.Moq.ParamsKey_Draw(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Draw(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Draw(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Draw(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(dst, param2, src, sp)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(dst, param2, src, sp)
	}
	return
}

// OnCall returns the recorder implementation of the Op_genType type
func (m *MoqOp_genType) OnCall() *MoqOp_genType_recorder {
	return &MoqOp_genType_recorder{
		Moq: m,
	}
}

func (m *MoqOp_genType_recorder) Draw(dst draw.Image, param2 image.Rectangle, src image.Image, sp image.Point) *MoqOp_genType_Draw_fnRecorder {
	return &MoqOp_genType_Draw_fnRecorder{
		Params: MoqOp_genType_Draw_params{
			Dst:    dst,
			Param2: param2,
			Src:    src,
			Sp:     sp,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqOp_genType_Draw_fnRecorder) Any() *MoqOp_genType_Draw_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Draw(r.Params))
		return nil
	}
	return &MoqOp_genType_Draw_anyParams{Recorder: r}
}

func (a *MoqOp_genType_Draw_anyParams) Dst() *MoqOp_genType_Draw_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqOp_genType_Draw_anyParams) Param2() *MoqOp_genType_Draw_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqOp_genType_Draw_anyParams) Src() *MoqOp_genType_Draw_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (a *MoqOp_genType_Draw_anyParams) Sp() *MoqOp_genType_Draw_fnRecorder {
	a.Recorder.AnyParams |= 1 << 3
	return a.Recorder
}

func (r *MoqOp_genType_Draw_fnRecorder) Seq() *MoqOp_genType_Draw_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Draw(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqOp_genType_Draw_fnRecorder) NoSeq() *MoqOp_genType_Draw_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Draw(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqOp_genType_Draw_fnRecorder) ReturnResults() *MoqOp_genType_Draw_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqOp_genType_Draw_doFn
		DoReturnFn MoqOp_genType_Draw_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqOp_genType_Draw_fnRecorder) AndDo(fn MoqOp_genType_Draw_doFn) *MoqOp_genType_Draw_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqOp_genType_Draw_fnRecorder) DoReturnResults(fn MoqOp_genType_Draw_doReturnFn) *MoqOp_genType_Draw_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqOp_genType_Draw_doFn
		DoReturnFn MoqOp_genType_Draw_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqOp_genType_Draw_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqOp_genType_Draw_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Draw {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqOp_genType_Draw_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqOp_genType_Draw_paramsKey]*MoqOp_genType_Draw_results{},
		}
		r.Moq.ResultsByParams_Draw = append(r.Moq.ResultsByParams_Draw, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Draw) {
			copy(r.Moq.ResultsByParams_Draw[insertAt+1:], r.Moq.ResultsByParams_Draw[insertAt:0])
			r.Moq.ResultsByParams_Draw[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Draw(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqOp_genType_Draw_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqOp_genType_Draw_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqOp_genType_Draw_fnRecorder {
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
				Values     *struct{}
				Sequence   uint32
				DoFn       MoqOp_genType_Draw_doFn
				DoReturnFn MoqOp_genType_Draw_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqOp_genType) PrettyParams_Draw(params MoqOp_genType_Draw_params) string {
	return fmt.Sprintf("Draw(%#v, %#v, %#v, %#v)", params.Dst, params.Param2, params.Src, params.Sp)
}

func (m *MoqOp_genType) ParamsKey_Draw(params MoqOp_genType_Draw_params, anyParams uint64) MoqOp_genType_Draw_paramsKey {
	m.Scene.T.Helper()
	var dstUsed draw.Image
	var dstUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Draw.Dst == moq.ParamIndexByValue {
			dstUsed = params.Dst
		} else {
			dstUsedHash = hash.DeepHash(params.Dst)
		}
	}
	var param2Used image.Rectangle
	var param2UsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Draw.Param2 == moq.ParamIndexByValue {
			param2Used = params.Param2
		} else {
			param2UsedHash = hash.DeepHash(params.Param2)
		}
	}
	var srcUsed image.Image
	var srcUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Draw.Src == moq.ParamIndexByValue {
			srcUsed = params.Src
		} else {
			srcUsedHash = hash.DeepHash(params.Src)
		}
	}
	var spUsed image.Point
	var spUsedHash hash.Hash
	if anyParams&(1<<3) == 0 {
		if m.Runtime.ParameterIndexing.Draw.Sp == moq.ParamIndexByValue {
			spUsed = params.Sp
		} else {
			spUsedHash = hash.DeepHash(params.Sp)
		}
	}
	return MoqOp_genType_Draw_paramsKey{
		Params: struct {
			Dst    draw.Image
			Param2 image.Rectangle
			Src    image.Image
			Sp     image.Point
		}{
			Dst:    dstUsed,
			Param2: param2Used,
			Src:    srcUsed,
			Sp:     spUsed,
		},
		Hashes: struct {
			Dst    hash.Hash
			Param2 hash.Hash
			Src    hash.Hash
			Sp     hash.Hash
		}{
			Dst:    dstUsedHash,
			Param2: param2UsedHash,
			Src:    srcUsedHash,
			Sp:     spUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqOp_genType) Reset() { m.ResultsByParams_Draw = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqOp_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Draw {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Draw(results.Params))
			}
		}
	}
}

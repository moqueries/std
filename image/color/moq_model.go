// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package color

import (
	"fmt"
	"image/color"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that color.Model is mocked completely
var _ color.Model = (*MoqModel_mock)(nil)

// MoqModel holds the state of a moq of the Model type
type MoqModel struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqModel_mock

	ResultsByParams_Convert []MoqModel_Convert_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Convert struct {
				C moq.ParamIndexing
			}
		}
	}
	// MoqModel_mock isolates the mock interface of the Model type
}

type MoqModel_mock struct {
	Moq *MoqModel
}

// MoqModel_recorder isolates the recorder interface of the Model type
type MoqModel_recorder struct {
	Moq *MoqModel
}

// MoqModel_Convert_params holds the params of the Model type
type MoqModel_Convert_params struct{ C color.Color }

// MoqModel_Convert_paramsKey holds the map key params of the Model type
type MoqModel_Convert_paramsKey struct {
	Params struct{ C color.Color }
	Hashes struct{ C hash.Hash }
}

// MoqModel_Convert_resultsByParams contains the results for a given set of
// parameters for the Model type
type MoqModel_Convert_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqModel_Convert_paramsKey]*MoqModel_Convert_results
}

// MoqModel_Convert_doFn defines the type of function needed when calling AndDo
// for the Model type
type MoqModel_Convert_doFn func(c color.Color)

// MoqModel_Convert_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Model type
type MoqModel_Convert_doReturnFn func(c color.Color) color.Color

// MoqModel_Convert_results holds the results of the Model type
type MoqModel_Convert_results struct {
	Params  MoqModel_Convert_params
	Results []struct {
		Values *struct {
			Result1 color.Color
		}
		Sequence   uint32
		DoFn       MoqModel_Convert_doFn
		DoReturnFn MoqModel_Convert_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqModel_Convert_fnRecorder routes recorded function calls to the MoqModel
// moq
type MoqModel_Convert_fnRecorder struct {
	Params    MoqModel_Convert_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqModel_Convert_results
	Moq       *MoqModel
}

// MoqModel_Convert_anyParams isolates the any params functions of the Model
// type
type MoqModel_Convert_anyParams struct {
	Recorder *MoqModel_Convert_fnRecorder
}

// NewMoqModel creates a new moq of the Model type
func NewMoqModel(scene *moq.Scene, config *moq.Config) *MoqModel {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqModel{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqModel_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Convert struct {
					C moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Convert struct {
				C moq.ParamIndexing
			}
		}{
			Convert: struct {
				C moq.ParamIndexing
			}{
				C: moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Model type
func (m *MoqModel) Mock() *MoqModel_mock { return m.Moq }

func (m *MoqModel_mock) Convert(c color.Color) (result1 color.Color) {
	m.Moq.Scene.T.Helper()
	params := MoqModel_Convert_params{
		C: c,
	}
	var results *MoqModel_Convert_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Convert {
		paramsKey := m.Moq.ParamsKey_Convert(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Convert(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Convert(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Convert(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(c)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(c)
	}
	return
}

// OnCall returns the recorder implementation of the Model type
func (m *MoqModel) OnCall() *MoqModel_recorder {
	return &MoqModel_recorder{
		Moq: m,
	}
}

func (m *MoqModel_recorder) Convert(c color.Color) *MoqModel_Convert_fnRecorder {
	return &MoqModel_Convert_fnRecorder{
		Params: MoqModel_Convert_params{
			C: c,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqModel_Convert_fnRecorder) Any() *MoqModel_Convert_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Convert(r.Params))
		return nil
	}
	return &MoqModel_Convert_anyParams{Recorder: r}
}

func (a *MoqModel_Convert_anyParams) C() *MoqModel_Convert_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqModel_Convert_fnRecorder) Seq() *MoqModel_Convert_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Convert(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqModel_Convert_fnRecorder) NoSeq() *MoqModel_Convert_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Convert(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqModel_Convert_fnRecorder) ReturnResults(result1 color.Color) *MoqModel_Convert_fnRecorder {
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
		DoFn       MoqModel_Convert_doFn
		DoReturnFn MoqModel_Convert_doReturnFn
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

func (r *MoqModel_Convert_fnRecorder) AndDo(fn MoqModel_Convert_doFn) *MoqModel_Convert_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqModel_Convert_fnRecorder) DoReturnResults(fn MoqModel_Convert_doReturnFn) *MoqModel_Convert_fnRecorder {
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
		DoFn       MoqModel_Convert_doFn
		DoReturnFn MoqModel_Convert_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqModel_Convert_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqModel_Convert_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Convert {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqModel_Convert_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqModel_Convert_paramsKey]*MoqModel_Convert_results{},
		}
		r.Moq.ResultsByParams_Convert = append(r.Moq.ResultsByParams_Convert, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Convert) {
			copy(r.Moq.ResultsByParams_Convert[insertAt+1:], r.Moq.ResultsByParams_Convert[insertAt:0])
			r.Moq.ResultsByParams_Convert[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Convert(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqModel_Convert_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqModel_Convert_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqModel_Convert_fnRecorder {
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
				DoFn       MoqModel_Convert_doFn
				DoReturnFn MoqModel_Convert_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqModel) PrettyParams_Convert(params MoqModel_Convert_params) string {
	return fmt.Sprintf("Convert(%#v)", params.C)
}

func (m *MoqModel) ParamsKey_Convert(params MoqModel_Convert_params, anyParams uint64) MoqModel_Convert_paramsKey {
	m.Scene.T.Helper()
	var cUsed color.Color
	var cUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Convert.C == moq.ParamIndexByValue {
			cUsed = params.C
		} else {
			cUsedHash = hash.DeepHash(params.C)
		}
	}
	return MoqModel_Convert_paramsKey{
		Params: struct{ C color.Color }{
			C: cUsed,
		},
		Hashes: struct{ C hash.Hash }{
			C: cUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqModel) Reset() { m.ResultsByParams_Convert = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqModel) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Convert {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Convert(results.Params))
			}
		}
	}
}
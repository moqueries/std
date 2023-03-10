// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package gif

import (
	"fmt"
	"image/gif"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// EncodeAll_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type EncodeAll_genType func(w io.Writer, g *gif.GIF) error

// MoqEncodeAll_genType holds the state of a moq of the EncodeAll_genType type
type MoqEncodeAll_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqEncodeAll_genType_mock

	ResultsByParams []MoqEncodeAll_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			W moq.ParamIndexing
			G moq.ParamIndexing
		}
	}
}

// MoqEncodeAll_genType_mock isolates the mock interface of the
// EncodeAll_genType type
type MoqEncodeAll_genType_mock struct {
	Moq *MoqEncodeAll_genType
}

// MoqEncodeAll_genType_params holds the params of the EncodeAll_genType type
type MoqEncodeAll_genType_params struct {
	W io.Writer
	G *gif.GIF
}

// MoqEncodeAll_genType_paramsKey holds the map key params of the
// EncodeAll_genType type
type MoqEncodeAll_genType_paramsKey struct {
	Params struct {
		W io.Writer
		G *gif.GIF
	}
	Hashes struct {
		W hash.Hash
		G hash.Hash
	}
}

// MoqEncodeAll_genType_resultsByParams contains the results for a given set of
// parameters for the EncodeAll_genType type
type MoqEncodeAll_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqEncodeAll_genType_paramsKey]*MoqEncodeAll_genType_results
}

// MoqEncodeAll_genType_doFn defines the type of function needed when calling
// AndDo for the EncodeAll_genType type
type MoqEncodeAll_genType_doFn func(w io.Writer, g *gif.GIF)

// MoqEncodeAll_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the EncodeAll_genType type
type MoqEncodeAll_genType_doReturnFn func(w io.Writer, g *gif.GIF) error

// MoqEncodeAll_genType_results holds the results of the EncodeAll_genType type
type MoqEncodeAll_genType_results struct {
	Params  MoqEncodeAll_genType_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqEncodeAll_genType_doFn
		DoReturnFn MoqEncodeAll_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqEncodeAll_genType_fnRecorder routes recorded function calls to the
// MoqEncodeAll_genType moq
type MoqEncodeAll_genType_fnRecorder struct {
	Params    MoqEncodeAll_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqEncodeAll_genType_results
	Moq       *MoqEncodeAll_genType
}

// MoqEncodeAll_genType_anyParams isolates the any params functions of the
// EncodeAll_genType type
type MoqEncodeAll_genType_anyParams struct {
	Recorder *MoqEncodeAll_genType_fnRecorder
}

// NewMoqEncodeAll_genType creates a new moq of the EncodeAll_genType type
func NewMoqEncodeAll_genType(scene *moq.Scene, config *moq.Config) *MoqEncodeAll_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqEncodeAll_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqEncodeAll_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				W moq.ParamIndexing
				G moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			W moq.ParamIndexing
			G moq.ParamIndexing
		}{
			W: moq.ParamIndexByHash,
			G: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the EncodeAll_genType type
func (m *MoqEncodeAll_genType) Mock() EncodeAll_genType {
	return func(w io.Writer, g *gif.GIF) error {
		m.Scene.T.Helper()
		moq := &MoqEncodeAll_genType_mock{Moq: m}
		return moq.Fn(w, g)
	}
}

func (m *MoqEncodeAll_genType_mock) Fn(w io.Writer, g *gif.GIF) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqEncodeAll_genType_params{
		W: w,
		G: g,
	}
	var results *MoqEncodeAll_genType_results
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
		result.DoFn(w, g)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(w, g)
	}
	return
}

func (m *MoqEncodeAll_genType) OnCall(w io.Writer, g *gif.GIF) *MoqEncodeAll_genType_fnRecorder {
	return &MoqEncodeAll_genType_fnRecorder{
		Params: MoqEncodeAll_genType_params{
			W: w,
			G: g,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqEncodeAll_genType_fnRecorder) Any() *MoqEncodeAll_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqEncodeAll_genType_anyParams{Recorder: r}
}

func (a *MoqEncodeAll_genType_anyParams) W() *MoqEncodeAll_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqEncodeAll_genType_anyParams) G() *MoqEncodeAll_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqEncodeAll_genType_fnRecorder) Seq() *MoqEncodeAll_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqEncodeAll_genType_fnRecorder) NoSeq() *MoqEncodeAll_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqEncodeAll_genType_fnRecorder) ReturnResults(result1 error) *MoqEncodeAll_genType_fnRecorder {
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
		DoFn       MoqEncodeAll_genType_doFn
		DoReturnFn MoqEncodeAll_genType_doReturnFn
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

func (r *MoqEncodeAll_genType_fnRecorder) AndDo(fn MoqEncodeAll_genType_doFn) *MoqEncodeAll_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqEncodeAll_genType_fnRecorder) DoReturnResults(fn MoqEncodeAll_genType_doReturnFn) *MoqEncodeAll_genType_fnRecorder {
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
		DoFn       MoqEncodeAll_genType_doFn
		DoReturnFn MoqEncodeAll_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqEncodeAll_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqEncodeAll_genType_resultsByParams
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
		results = &MoqEncodeAll_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqEncodeAll_genType_paramsKey]*MoqEncodeAll_genType_results{},
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
		r.Results = &MoqEncodeAll_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqEncodeAll_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqEncodeAll_genType_fnRecorder {
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
				DoFn       MoqEncodeAll_genType_doFn
				DoReturnFn MoqEncodeAll_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqEncodeAll_genType) PrettyParams(params MoqEncodeAll_genType_params) string {
	return fmt.Sprintf("EncodeAll_genType(%#v, %#v)", params.W, params.G)
}

func (m *MoqEncodeAll_genType) ParamsKey(params MoqEncodeAll_genType_params, anyParams uint64) MoqEncodeAll_genType_paramsKey {
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
	var gUsed *gif.GIF
	var gUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.G == moq.ParamIndexByValue {
			gUsed = params.G
		} else {
			gUsedHash = hash.DeepHash(params.G)
		}
	}
	return MoqEncodeAll_genType_paramsKey{
		Params: struct {
			W io.Writer
			G *gif.GIF
		}{
			W: wUsed,
			G: gUsed,
		},
		Hashes: struct {
			W hash.Hash
			G hash.Hash
		}{
			W: wUsedHash,
			G: gUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqEncodeAll_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqEncodeAll_genType) AssertExpectationsMet() {
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

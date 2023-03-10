// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package zlib

import (
	"compress/zlib"
	"fmt"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// NewWriterLevel_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type NewWriterLevel_genType func(w io.Writer, level int) (*zlib.Writer, error)

// MoqNewWriterLevel_genType holds the state of a moq of the
// NewWriterLevel_genType type
type MoqNewWriterLevel_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNewWriterLevel_genType_mock

	ResultsByParams []MoqNewWriterLevel_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			W     moq.ParamIndexing
			Level moq.ParamIndexing
		}
	}
}

// MoqNewWriterLevel_genType_mock isolates the mock interface of the
// NewWriterLevel_genType type
type MoqNewWriterLevel_genType_mock struct {
	Moq *MoqNewWriterLevel_genType
}

// MoqNewWriterLevel_genType_params holds the params of the
// NewWriterLevel_genType type
type MoqNewWriterLevel_genType_params struct {
	W     io.Writer
	Level int
}

// MoqNewWriterLevel_genType_paramsKey holds the map key params of the
// NewWriterLevel_genType type
type MoqNewWriterLevel_genType_paramsKey struct {
	Params struct {
		W     io.Writer
		Level int
	}
	Hashes struct {
		W     hash.Hash
		Level hash.Hash
	}
}

// MoqNewWriterLevel_genType_resultsByParams contains the results for a given
// set of parameters for the NewWriterLevel_genType type
type MoqNewWriterLevel_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNewWriterLevel_genType_paramsKey]*MoqNewWriterLevel_genType_results
}

// MoqNewWriterLevel_genType_doFn defines the type of function needed when
// calling AndDo for the NewWriterLevel_genType type
type MoqNewWriterLevel_genType_doFn func(w io.Writer, level int)

// MoqNewWriterLevel_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the NewWriterLevel_genType type
type MoqNewWriterLevel_genType_doReturnFn func(w io.Writer, level int) (*zlib.Writer, error)

// MoqNewWriterLevel_genType_results holds the results of the
// NewWriterLevel_genType type
type MoqNewWriterLevel_genType_results struct {
	Params  MoqNewWriterLevel_genType_params
	Results []struct {
		Values *struct {
			Result1 *zlib.Writer
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqNewWriterLevel_genType_doFn
		DoReturnFn MoqNewWriterLevel_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNewWriterLevel_genType_fnRecorder routes recorded function calls to the
// MoqNewWriterLevel_genType moq
type MoqNewWriterLevel_genType_fnRecorder struct {
	Params    MoqNewWriterLevel_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNewWriterLevel_genType_results
	Moq       *MoqNewWriterLevel_genType
}

// MoqNewWriterLevel_genType_anyParams isolates the any params functions of the
// NewWriterLevel_genType type
type MoqNewWriterLevel_genType_anyParams struct {
	Recorder *MoqNewWriterLevel_genType_fnRecorder
}

// NewMoqNewWriterLevel_genType creates a new moq of the NewWriterLevel_genType
// type
func NewMoqNewWriterLevel_genType(scene *moq.Scene, config *moq.Config) *MoqNewWriterLevel_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNewWriterLevel_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNewWriterLevel_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				W     moq.ParamIndexing
				Level moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			W     moq.ParamIndexing
			Level moq.ParamIndexing
		}{
			W:     moq.ParamIndexByHash,
			Level: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the NewWriterLevel_genType type
func (m *MoqNewWriterLevel_genType) Mock() NewWriterLevel_genType {
	return func(w io.Writer, level int) (*zlib.Writer, error) {
		m.Scene.T.Helper()
		moq := &MoqNewWriterLevel_genType_mock{Moq: m}
		return moq.Fn(w, level)
	}
}

func (m *MoqNewWriterLevel_genType_mock) Fn(w io.Writer, level int) (result1 *zlib.Writer, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqNewWriterLevel_genType_params{
		W:     w,
		Level: level,
	}
	var results *MoqNewWriterLevel_genType_results
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
		result.DoFn(w, level)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(w, level)
	}
	return
}

func (m *MoqNewWriterLevel_genType) OnCall(w io.Writer, level int) *MoqNewWriterLevel_genType_fnRecorder {
	return &MoqNewWriterLevel_genType_fnRecorder{
		Params: MoqNewWriterLevel_genType_params{
			W:     w,
			Level: level,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNewWriterLevel_genType_fnRecorder) Any() *MoqNewWriterLevel_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNewWriterLevel_genType_anyParams{Recorder: r}
}

func (a *MoqNewWriterLevel_genType_anyParams) W() *MoqNewWriterLevel_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqNewWriterLevel_genType_anyParams) Level() *MoqNewWriterLevel_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqNewWriterLevel_genType_fnRecorder) Seq() *MoqNewWriterLevel_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNewWriterLevel_genType_fnRecorder) NoSeq() *MoqNewWriterLevel_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNewWriterLevel_genType_fnRecorder) ReturnResults(result1 *zlib.Writer, result2 error) *MoqNewWriterLevel_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *zlib.Writer
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqNewWriterLevel_genType_doFn
		DoReturnFn MoqNewWriterLevel_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *zlib.Writer
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNewWriterLevel_genType_fnRecorder) AndDo(fn MoqNewWriterLevel_genType_doFn) *MoqNewWriterLevel_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNewWriterLevel_genType_fnRecorder) DoReturnResults(fn MoqNewWriterLevel_genType_doReturnFn) *MoqNewWriterLevel_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *zlib.Writer
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqNewWriterLevel_genType_doFn
		DoReturnFn MoqNewWriterLevel_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNewWriterLevel_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNewWriterLevel_genType_resultsByParams
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
		results = &MoqNewWriterLevel_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNewWriterLevel_genType_paramsKey]*MoqNewWriterLevel_genType_results{},
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
		r.Results = &MoqNewWriterLevel_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNewWriterLevel_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNewWriterLevel_genType_fnRecorder {
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
					Result1 *zlib.Writer
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqNewWriterLevel_genType_doFn
				DoReturnFn MoqNewWriterLevel_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNewWriterLevel_genType) PrettyParams(params MoqNewWriterLevel_genType_params) string {
	return fmt.Sprintf("NewWriterLevel_genType(%#v, %#v)", params.W, params.Level)
}

func (m *MoqNewWriterLevel_genType) ParamsKey(params MoqNewWriterLevel_genType_params, anyParams uint64) MoqNewWriterLevel_genType_paramsKey {
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
	var levelUsed int
	var levelUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Level == moq.ParamIndexByValue {
			levelUsed = params.Level
		} else {
			levelUsedHash = hash.DeepHash(params.Level)
		}
	}
	return MoqNewWriterLevel_genType_paramsKey{
		Params: struct {
			W     io.Writer
			Level int
		}{
			W:     wUsed,
			Level: levelUsed,
		},
		Hashes: struct {
			W     hash.Hash
			Level hash.Hash
		}{
			W:     wUsedHash,
			Level: levelUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNewWriterLevel_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNewWriterLevel_genType) AssertExpectationsMet() {
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

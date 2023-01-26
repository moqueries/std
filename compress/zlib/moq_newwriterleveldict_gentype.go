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

// NewWriterLevelDict_genType is the fabricated implementation type of this
// mock (emitted when mocking functions directly and not from a function type)
type NewWriterLevelDict_genType func(w io.Writer, level int, dict []byte) (*zlib.Writer, error)

// MoqNewWriterLevelDict_genType holds the state of a moq of the
// NewWriterLevelDict_genType type
type MoqNewWriterLevelDict_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNewWriterLevelDict_genType_mock

	ResultsByParams []MoqNewWriterLevelDict_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			W     moq.ParamIndexing
			Level moq.ParamIndexing
			Dict  moq.ParamIndexing
		}
	}
}

// MoqNewWriterLevelDict_genType_mock isolates the mock interface of the
// NewWriterLevelDict_genType type
type MoqNewWriterLevelDict_genType_mock struct {
	Moq *MoqNewWriterLevelDict_genType
}

// MoqNewWriterLevelDict_genType_params holds the params of the
// NewWriterLevelDict_genType type
type MoqNewWriterLevelDict_genType_params struct {
	W     io.Writer
	Level int
	Dict  []byte
}

// MoqNewWriterLevelDict_genType_paramsKey holds the map key params of the
// NewWriterLevelDict_genType type
type MoqNewWriterLevelDict_genType_paramsKey struct {
	Params struct {
		W     io.Writer
		Level int
	}
	Hashes struct {
		W     hash.Hash
		Level hash.Hash
		Dict  hash.Hash
	}
}

// MoqNewWriterLevelDict_genType_resultsByParams contains the results for a
// given set of parameters for the NewWriterLevelDict_genType type
type MoqNewWriterLevelDict_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNewWriterLevelDict_genType_paramsKey]*MoqNewWriterLevelDict_genType_results
}

// MoqNewWriterLevelDict_genType_doFn defines the type of function needed when
// calling AndDo for the NewWriterLevelDict_genType type
type MoqNewWriterLevelDict_genType_doFn func(w io.Writer, level int, dict []byte)

// MoqNewWriterLevelDict_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the NewWriterLevelDict_genType type
type MoqNewWriterLevelDict_genType_doReturnFn func(w io.Writer, level int, dict []byte) (*zlib.Writer, error)

// MoqNewWriterLevelDict_genType_results holds the results of the
// NewWriterLevelDict_genType type
type MoqNewWriterLevelDict_genType_results struct {
	Params  MoqNewWriterLevelDict_genType_params
	Results []struct {
		Values *struct {
			Result1 *zlib.Writer
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqNewWriterLevelDict_genType_doFn
		DoReturnFn MoqNewWriterLevelDict_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNewWriterLevelDict_genType_fnRecorder routes recorded function calls to
// the MoqNewWriterLevelDict_genType moq
type MoqNewWriterLevelDict_genType_fnRecorder struct {
	Params    MoqNewWriterLevelDict_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNewWriterLevelDict_genType_results
	Moq       *MoqNewWriterLevelDict_genType
}

// MoqNewWriterLevelDict_genType_anyParams isolates the any params functions of
// the NewWriterLevelDict_genType type
type MoqNewWriterLevelDict_genType_anyParams struct {
	Recorder *MoqNewWriterLevelDict_genType_fnRecorder
}

// NewMoqNewWriterLevelDict_genType creates a new moq of the
// NewWriterLevelDict_genType type
func NewMoqNewWriterLevelDict_genType(scene *moq.Scene, config *moq.Config) *MoqNewWriterLevelDict_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNewWriterLevelDict_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNewWriterLevelDict_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				W     moq.ParamIndexing
				Level moq.ParamIndexing
				Dict  moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			W     moq.ParamIndexing
			Level moq.ParamIndexing
			Dict  moq.ParamIndexing
		}{
			W:     moq.ParamIndexByHash,
			Level: moq.ParamIndexByValue,
			Dict:  moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the NewWriterLevelDict_genType type
func (m *MoqNewWriterLevelDict_genType) Mock() NewWriterLevelDict_genType {
	return func(w io.Writer, level int, dict []byte) (*zlib.Writer, error) {
		m.Scene.T.Helper()
		moq := &MoqNewWriterLevelDict_genType_mock{Moq: m}
		return moq.Fn(w, level, dict)
	}
}

func (m *MoqNewWriterLevelDict_genType_mock) Fn(w io.Writer, level int, dict []byte) (result1 *zlib.Writer, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqNewWriterLevelDict_genType_params{
		W:     w,
		Level: level,
		Dict:  dict,
	}
	var results *MoqNewWriterLevelDict_genType_results
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
		result.DoFn(w, level, dict)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(w, level, dict)
	}
	return
}

func (m *MoqNewWriterLevelDict_genType) OnCall(w io.Writer, level int, dict []byte) *MoqNewWriterLevelDict_genType_fnRecorder {
	return &MoqNewWriterLevelDict_genType_fnRecorder{
		Params: MoqNewWriterLevelDict_genType_params{
			W:     w,
			Level: level,
			Dict:  dict,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNewWriterLevelDict_genType_fnRecorder) Any() *MoqNewWriterLevelDict_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNewWriterLevelDict_genType_anyParams{Recorder: r}
}

func (a *MoqNewWriterLevelDict_genType_anyParams) W() *MoqNewWriterLevelDict_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqNewWriterLevelDict_genType_anyParams) Level() *MoqNewWriterLevelDict_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqNewWriterLevelDict_genType_anyParams) Dict() *MoqNewWriterLevelDict_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqNewWriterLevelDict_genType_fnRecorder) Seq() *MoqNewWriterLevelDict_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNewWriterLevelDict_genType_fnRecorder) NoSeq() *MoqNewWriterLevelDict_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNewWriterLevelDict_genType_fnRecorder) ReturnResults(result1 *zlib.Writer, result2 error) *MoqNewWriterLevelDict_genType_fnRecorder {
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
		DoFn       MoqNewWriterLevelDict_genType_doFn
		DoReturnFn MoqNewWriterLevelDict_genType_doReturnFn
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

func (r *MoqNewWriterLevelDict_genType_fnRecorder) AndDo(fn MoqNewWriterLevelDict_genType_doFn) *MoqNewWriterLevelDict_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNewWriterLevelDict_genType_fnRecorder) DoReturnResults(fn MoqNewWriterLevelDict_genType_doReturnFn) *MoqNewWriterLevelDict_genType_fnRecorder {
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
		DoFn       MoqNewWriterLevelDict_genType_doFn
		DoReturnFn MoqNewWriterLevelDict_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNewWriterLevelDict_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNewWriterLevelDict_genType_resultsByParams
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
		results = &MoqNewWriterLevelDict_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNewWriterLevelDict_genType_paramsKey]*MoqNewWriterLevelDict_genType_results{},
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
		r.Results = &MoqNewWriterLevelDict_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNewWriterLevelDict_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNewWriterLevelDict_genType_fnRecorder {
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
				DoFn       MoqNewWriterLevelDict_genType_doFn
				DoReturnFn MoqNewWriterLevelDict_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNewWriterLevelDict_genType) PrettyParams(params MoqNewWriterLevelDict_genType_params) string {
	return fmt.Sprintf("NewWriterLevelDict_genType(%#v, %#v, %#v)", params.W, params.Level, params.Dict)
}

func (m *MoqNewWriterLevelDict_genType) ParamsKey(params MoqNewWriterLevelDict_genType_params, anyParams uint64) MoqNewWriterLevelDict_genType_paramsKey {
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
	var dictUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Dict == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The dict parameter can't be indexed by value")
		}
		dictUsedHash = hash.DeepHash(params.Dict)
	}
	return MoqNewWriterLevelDict_genType_paramsKey{
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
			Dict  hash.Hash
		}{
			W:     wUsedHash,
			Level: levelUsedHash,
			Dict:  dictUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNewWriterLevelDict_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNewWriterLevelDict_genType) AssertExpectationsMet() {
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
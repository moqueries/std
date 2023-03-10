// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package flate

import (
	"fmt"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// NewReaderDict_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type NewReaderDict_genType func(r io.Reader, dict []byte) io.ReadCloser

// MoqNewReaderDict_genType holds the state of a moq of the
// NewReaderDict_genType type
type MoqNewReaderDict_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNewReaderDict_genType_mock

	ResultsByParams []MoqNewReaderDict_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Param1 moq.ParamIndexing
			Dict   moq.ParamIndexing
		}
	}
}

// MoqNewReaderDict_genType_mock isolates the mock interface of the
// NewReaderDict_genType type
type MoqNewReaderDict_genType_mock struct {
	Moq *MoqNewReaderDict_genType
}

// MoqNewReaderDict_genType_params holds the params of the
// NewReaderDict_genType type
type MoqNewReaderDict_genType_params struct {
	Param1 io.Reader
	Dict   []byte
}

// MoqNewReaderDict_genType_paramsKey holds the map key params of the
// NewReaderDict_genType type
type MoqNewReaderDict_genType_paramsKey struct {
	Params struct{ Param1 io.Reader }
	Hashes struct {
		Param1 hash.Hash
		Dict   hash.Hash
	}
}

// MoqNewReaderDict_genType_resultsByParams contains the results for a given
// set of parameters for the NewReaderDict_genType type
type MoqNewReaderDict_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNewReaderDict_genType_paramsKey]*MoqNewReaderDict_genType_results
}

// MoqNewReaderDict_genType_doFn defines the type of function needed when
// calling AndDo for the NewReaderDict_genType type
type MoqNewReaderDict_genType_doFn func(r io.Reader, dict []byte)

// MoqNewReaderDict_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the NewReaderDict_genType type
type MoqNewReaderDict_genType_doReturnFn func(r io.Reader, dict []byte) io.ReadCloser

// MoqNewReaderDict_genType_results holds the results of the
// NewReaderDict_genType type
type MoqNewReaderDict_genType_results struct {
	Params  MoqNewReaderDict_genType_params
	Results []struct {
		Values *struct {
			Result1 io.ReadCloser
		}
		Sequence   uint32
		DoFn       MoqNewReaderDict_genType_doFn
		DoReturnFn MoqNewReaderDict_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNewReaderDict_genType_fnRecorder routes recorded function calls to the
// MoqNewReaderDict_genType moq
type MoqNewReaderDict_genType_fnRecorder struct {
	Params    MoqNewReaderDict_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNewReaderDict_genType_results
	Moq       *MoqNewReaderDict_genType
}

// MoqNewReaderDict_genType_anyParams isolates the any params functions of the
// NewReaderDict_genType type
type MoqNewReaderDict_genType_anyParams struct {
	Recorder *MoqNewReaderDict_genType_fnRecorder
}

// NewMoqNewReaderDict_genType creates a new moq of the NewReaderDict_genType
// type
func NewMoqNewReaderDict_genType(scene *moq.Scene, config *moq.Config) *MoqNewReaderDict_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNewReaderDict_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNewReaderDict_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Param1 moq.ParamIndexing
				Dict   moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Param1 moq.ParamIndexing
			Dict   moq.ParamIndexing
		}{
			Param1: moq.ParamIndexByHash,
			Dict:   moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the NewReaderDict_genType type
func (m *MoqNewReaderDict_genType) Mock() NewReaderDict_genType {
	return func(param1 io.Reader, dict []byte) io.ReadCloser {
		m.Scene.T.Helper()
		moq := &MoqNewReaderDict_genType_mock{Moq: m}
		return moq.Fn(param1, dict)
	}
}

func (m *MoqNewReaderDict_genType_mock) Fn(param1 io.Reader, dict []byte) (result1 io.ReadCloser) {
	m.Moq.Scene.T.Helper()
	params := MoqNewReaderDict_genType_params{
		Param1: param1,
		Dict:   dict,
	}
	var results *MoqNewReaderDict_genType_results
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
		result.DoFn(param1, dict)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(param1, dict)
	}
	return
}

func (m *MoqNewReaderDict_genType) OnCall(param1 io.Reader, dict []byte) *MoqNewReaderDict_genType_fnRecorder {
	return &MoqNewReaderDict_genType_fnRecorder{
		Params: MoqNewReaderDict_genType_params{
			Param1: param1,
			Dict:   dict,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNewReaderDict_genType_fnRecorder) Any() *MoqNewReaderDict_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNewReaderDict_genType_anyParams{Recorder: r}
}

func (a *MoqNewReaderDict_genType_anyParams) Param1() *MoqNewReaderDict_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqNewReaderDict_genType_anyParams) Dict() *MoqNewReaderDict_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqNewReaderDict_genType_fnRecorder) Seq() *MoqNewReaderDict_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNewReaderDict_genType_fnRecorder) NoSeq() *MoqNewReaderDict_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNewReaderDict_genType_fnRecorder) ReturnResults(result1 io.ReadCloser) *MoqNewReaderDict_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 io.ReadCloser
		}
		Sequence   uint32
		DoFn       MoqNewReaderDict_genType_doFn
		DoReturnFn MoqNewReaderDict_genType_doReturnFn
	}{
		Values: &struct {
			Result1 io.ReadCloser
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNewReaderDict_genType_fnRecorder) AndDo(fn MoqNewReaderDict_genType_doFn) *MoqNewReaderDict_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNewReaderDict_genType_fnRecorder) DoReturnResults(fn MoqNewReaderDict_genType_doReturnFn) *MoqNewReaderDict_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 io.ReadCloser
		}
		Sequence   uint32
		DoFn       MoqNewReaderDict_genType_doFn
		DoReturnFn MoqNewReaderDict_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNewReaderDict_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNewReaderDict_genType_resultsByParams
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
		results = &MoqNewReaderDict_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNewReaderDict_genType_paramsKey]*MoqNewReaderDict_genType_results{},
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
		r.Results = &MoqNewReaderDict_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNewReaderDict_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNewReaderDict_genType_fnRecorder {
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
					Result1 io.ReadCloser
				}
				Sequence   uint32
				DoFn       MoqNewReaderDict_genType_doFn
				DoReturnFn MoqNewReaderDict_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNewReaderDict_genType) PrettyParams(params MoqNewReaderDict_genType_params) string {
	return fmt.Sprintf("NewReaderDict_genType(%#v, %#v)", params.Param1, params.Dict)
}

func (m *MoqNewReaderDict_genType) ParamsKey(params MoqNewReaderDict_genType_params, anyParams uint64) MoqNewReaderDict_genType_paramsKey {
	m.Scene.T.Helper()
	var param1Used io.Reader
	var param1UsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Param1 == moq.ParamIndexByValue {
			param1Used = params.Param1
		} else {
			param1UsedHash = hash.DeepHash(params.Param1)
		}
	}
	var dictUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Dict == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The dict parameter can't be indexed by value")
		}
		dictUsedHash = hash.DeepHash(params.Dict)
	}
	return MoqNewReaderDict_genType_paramsKey{
		Params: struct{ Param1 io.Reader }{
			Param1: param1Used,
		},
		Hashes: struct {
			Param1 hash.Hash
			Dict   hash.Hash
		}{
			Param1: param1UsedHash,
			Dict:   dictUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNewReaderDict_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNewReaderDict_genType) AssertExpectationsMet() {
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

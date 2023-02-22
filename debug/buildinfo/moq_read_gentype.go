// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package buildinfo

import (
	"debug/buildinfo"
	"fmt"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Read_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Read_genType func(r io.ReaderAt) (*buildinfo.BuildInfo, error)

// MoqRead_genType holds the state of a moq of the Read_genType type
type MoqRead_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqRead_genType_mock

	ResultsByParams []MoqRead_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Param1 moq.ParamIndexing
		}
	}
}

// MoqRead_genType_mock isolates the mock interface of the Read_genType type
type MoqRead_genType_mock struct {
	Moq *MoqRead_genType
}

// MoqRead_genType_params holds the params of the Read_genType type
type MoqRead_genType_params struct{ Param1 io.ReaderAt }

// MoqRead_genType_paramsKey holds the map key params of the Read_genType type
type MoqRead_genType_paramsKey struct {
	Params struct{ Param1 io.ReaderAt }
	Hashes struct{ Param1 hash.Hash }
}

// MoqRead_genType_resultsByParams contains the results for a given set of
// parameters for the Read_genType type
type MoqRead_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqRead_genType_paramsKey]*MoqRead_genType_results
}

// MoqRead_genType_doFn defines the type of function needed when calling AndDo
// for the Read_genType type
type MoqRead_genType_doFn func(r io.ReaderAt)

// MoqRead_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Read_genType type
type MoqRead_genType_doReturnFn func(r io.ReaderAt) (*buildinfo.BuildInfo, error)

// MoqRead_genType_results holds the results of the Read_genType type
type MoqRead_genType_results struct {
	Params  MoqRead_genType_params
	Results []struct {
		Values *struct {
			Result1 *buildinfo.BuildInfo
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqRead_genType_doFn
		DoReturnFn MoqRead_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqRead_genType_fnRecorder routes recorded function calls to the
// MoqRead_genType moq
type MoqRead_genType_fnRecorder struct {
	Params    MoqRead_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqRead_genType_results
	Moq       *MoqRead_genType
}

// MoqRead_genType_anyParams isolates the any params functions of the
// Read_genType type
type MoqRead_genType_anyParams struct {
	Recorder *MoqRead_genType_fnRecorder
}

// NewMoqRead_genType creates a new moq of the Read_genType type
func NewMoqRead_genType(scene *moq.Scene, config *moq.Config) *MoqRead_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqRead_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqRead_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Param1 moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Param1 moq.ParamIndexing
		}{
			Param1: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Read_genType type
func (m *MoqRead_genType) Mock() Read_genType {
	return func(param1 io.ReaderAt) (*buildinfo.BuildInfo, error) {
		m.Scene.T.Helper()
		moq := &MoqRead_genType_mock{Moq: m}
		return moq.Fn(param1)
	}
}

func (m *MoqRead_genType_mock) Fn(param1 io.ReaderAt) (result1 *buildinfo.BuildInfo, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqRead_genType_params{
		Param1: param1,
	}
	var results *MoqRead_genType_results
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
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(param1)
	}
	return
}

func (m *MoqRead_genType) OnCall(param1 io.ReaderAt) *MoqRead_genType_fnRecorder {
	return &MoqRead_genType_fnRecorder{
		Params: MoqRead_genType_params{
			Param1: param1,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqRead_genType_fnRecorder) Any() *MoqRead_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqRead_genType_anyParams{Recorder: r}
}

func (a *MoqRead_genType_anyParams) Param1() *MoqRead_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqRead_genType_fnRecorder) Seq() *MoqRead_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqRead_genType_fnRecorder) NoSeq() *MoqRead_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqRead_genType_fnRecorder) ReturnResults(result1 *buildinfo.BuildInfo, result2 error) *MoqRead_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *buildinfo.BuildInfo
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqRead_genType_doFn
		DoReturnFn MoqRead_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *buildinfo.BuildInfo
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqRead_genType_fnRecorder) AndDo(fn MoqRead_genType_doFn) *MoqRead_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqRead_genType_fnRecorder) DoReturnResults(fn MoqRead_genType_doReturnFn) *MoqRead_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *buildinfo.BuildInfo
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqRead_genType_doFn
		DoReturnFn MoqRead_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqRead_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqRead_genType_resultsByParams
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
		results = &MoqRead_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqRead_genType_paramsKey]*MoqRead_genType_results{},
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
		r.Results = &MoqRead_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqRead_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqRead_genType_fnRecorder {
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
					Result1 *buildinfo.BuildInfo
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqRead_genType_doFn
				DoReturnFn MoqRead_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqRead_genType) PrettyParams(params MoqRead_genType_params) string {
	return fmt.Sprintf("Read_genType(%#v)", params.Param1)
}

func (m *MoqRead_genType) ParamsKey(params MoqRead_genType_params, anyParams uint64) MoqRead_genType_paramsKey {
	m.Scene.T.Helper()
	var param1Used io.ReaderAt
	var param1UsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Param1 == moq.ParamIndexByValue {
			param1Used = params.Param1
		} else {
			param1UsedHash = hash.DeepHash(params.Param1)
		}
	}
	return MoqRead_genType_paramsKey{
		Params: struct{ Param1 io.ReaderAt }{
			Param1: param1Used,
		},
		Hashes: struct{ Param1 hash.Hash }{
			Param1: param1UsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqRead_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqRead_genType) AssertExpectationsMet() {
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

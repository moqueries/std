// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package debug

import (
	"fmt"
	"math/bits"
	"runtime/debug"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// ParseBuildInfo_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type ParseBuildInfo_genType func(data string) (bi *debug.BuildInfo, err error)

// MoqParseBuildInfo_genType holds the state of a moq of the
// ParseBuildInfo_genType type
type MoqParseBuildInfo_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqParseBuildInfo_genType_mock

	ResultsByParams []MoqParseBuildInfo_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Data moq.ParamIndexing
		}
	}
}

// MoqParseBuildInfo_genType_mock isolates the mock interface of the
// ParseBuildInfo_genType type
type MoqParseBuildInfo_genType_mock struct {
	Moq *MoqParseBuildInfo_genType
}

// MoqParseBuildInfo_genType_params holds the params of the
// ParseBuildInfo_genType type
type MoqParseBuildInfo_genType_params struct{ Data string }

// MoqParseBuildInfo_genType_paramsKey holds the map key params of the
// ParseBuildInfo_genType type
type MoqParseBuildInfo_genType_paramsKey struct {
	Params struct{ Data string }
	Hashes struct{ Data hash.Hash }
}

// MoqParseBuildInfo_genType_resultsByParams contains the results for a given
// set of parameters for the ParseBuildInfo_genType type
type MoqParseBuildInfo_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqParseBuildInfo_genType_paramsKey]*MoqParseBuildInfo_genType_results
}

// MoqParseBuildInfo_genType_doFn defines the type of function needed when
// calling AndDo for the ParseBuildInfo_genType type
type MoqParseBuildInfo_genType_doFn func(data string)

// MoqParseBuildInfo_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the ParseBuildInfo_genType type
type MoqParseBuildInfo_genType_doReturnFn func(data string) (bi *debug.BuildInfo, err error)

// MoqParseBuildInfo_genType_results holds the results of the
// ParseBuildInfo_genType type
type MoqParseBuildInfo_genType_results struct {
	Params  MoqParseBuildInfo_genType_params
	Results []struct {
		Values *struct {
			Bi  *debug.BuildInfo
			Err error
		}
		Sequence   uint32
		DoFn       MoqParseBuildInfo_genType_doFn
		DoReturnFn MoqParseBuildInfo_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqParseBuildInfo_genType_fnRecorder routes recorded function calls to the
// MoqParseBuildInfo_genType moq
type MoqParseBuildInfo_genType_fnRecorder struct {
	Params    MoqParseBuildInfo_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqParseBuildInfo_genType_results
	Moq       *MoqParseBuildInfo_genType
}

// MoqParseBuildInfo_genType_anyParams isolates the any params functions of the
// ParseBuildInfo_genType type
type MoqParseBuildInfo_genType_anyParams struct {
	Recorder *MoqParseBuildInfo_genType_fnRecorder
}

// NewMoqParseBuildInfo_genType creates a new moq of the ParseBuildInfo_genType
// type
func NewMoqParseBuildInfo_genType(scene *moq.Scene, config *moq.Config) *MoqParseBuildInfo_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqParseBuildInfo_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqParseBuildInfo_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Data moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Data moq.ParamIndexing
		}{
			Data: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ParseBuildInfo_genType type
func (m *MoqParseBuildInfo_genType) Mock() ParseBuildInfo_genType {
	return func(data string) (_ *debug.BuildInfo, _ error) {
		m.Scene.T.Helper()
		moq := &MoqParseBuildInfo_genType_mock{Moq: m}
		return moq.Fn(data)
	}
}

func (m *MoqParseBuildInfo_genType_mock) Fn(data string) (bi *debug.BuildInfo, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqParseBuildInfo_genType_params{
		Data: data,
	}
	var results *MoqParseBuildInfo_genType_results
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
		result.DoFn(data)
	}

	if result.Values != nil {
		bi = result.Values.Bi
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		bi, err = result.DoReturnFn(data)
	}
	return
}

func (m *MoqParseBuildInfo_genType) OnCall(data string) *MoqParseBuildInfo_genType_fnRecorder {
	return &MoqParseBuildInfo_genType_fnRecorder{
		Params: MoqParseBuildInfo_genType_params{
			Data: data,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqParseBuildInfo_genType_fnRecorder) Any() *MoqParseBuildInfo_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqParseBuildInfo_genType_anyParams{Recorder: r}
}

func (a *MoqParseBuildInfo_genType_anyParams) Data() *MoqParseBuildInfo_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqParseBuildInfo_genType_fnRecorder) Seq() *MoqParseBuildInfo_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqParseBuildInfo_genType_fnRecorder) NoSeq() *MoqParseBuildInfo_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqParseBuildInfo_genType_fnRecorder) ReturnResults(bi *debug.BuildInfo, err error) *MoqParseBuildInfo_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Bi  *debug.BuildInfo
			Err error
		}
		Sequence   uint32
		DoFn       MoqParseBuildInfo_genType_doFn
		DoReturnFn MoqParseBuildInfo_genType_doReturnFn
	}{
		Values: &struct {
			Bi  *debug.BuildInfo
			Err error
		}{
			Bi:  bi,
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqParseBuildInfo_genType_fnRecorder) AndDo(fn MoqParseBuildInfo_genType_doFn) *MoqParseBuildInfo_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqParseBuildInfo_genType_fnRecorder) DoReturnResults(fn MoqParseBuildInfo_genType_doReturnFn) *MoqParseBuildInfo_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Bi  *debug.BuildInfo
			Err error
		}
		Sequence   uint32
		DoFn       MoqParseBuildInfo_genType_doFn
		DoReturnFn MoqParseBuildInfo_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqParseBuildInfo_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqParseBuildInfo_genType_resultsByParams
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
		results = &MoqParseBuildInfo_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqParseBuildInfo_genType_paramsKey]*MoqParseBuildInfo_genType_results{},
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
		r.Results = &MoqParseBuildInfo_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqParseBuildInfo_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqParseBuildInfo_genType_fnRecorder {
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
					Bi  *debug.BuildInfo
					Err error
				}
				Sequence   uint32
				DoFn       MoqParseBuildInfo_genType_doFn
				DoReturnFn MoqParseBuildInfo_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqParseBuildInfo_genType) PrettyParams(params MoqParseBuildInfo_genType_params) string {
	return fmt.Sprintf("ParseBuildInfo_genType(%#v)", params.Data)
}

func (m *MoqParseBuildInfo_genType) ParamsKey(params MoqParseBuildInfo_genType_params, anyParams uint64) MoqParseBuildInfo_genType_paramsKey {
	m.Scene.T.Helper()
	var dataUsed string
	var dataUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Data == moq.ParamIndexByValue {
			dataUsed = params.Data
		} else {
			dataUsedHash = hash.DeepHash(params.Data)
		}
	}
	return MoqParseBuildInfo_genType_paramsKey{
		Params: struct{ Data string }{
			Data: dataUsed,
		},
		Hashes: struct{ Data hash.Hash }{
			Data: dataUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqParseBuildInfo_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqParseBuildInfo_genType) AssertExpectationsMet() {
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
// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package regexp

import (
	"fmt"
	"math/bits"
	"regexp"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// MustCompilePOSIX_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type MustCompilePOSIX_genType func(str string) *regexp.Regexp

// MoqMustCompilePOSIX_genType holds the state of a moq of the
// MustCompilePOSIX_genType type
type MoqMustCompilePOSIX_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqMustCompilePOSIX_genType_mock

	ResultsByParams []MoqMustCompilePOSIX_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Str moq.ParamIndexing
		}
	}
}

// MoqMustCompilePOSIX_genType_mock isolates the mock interface of the
// MustCompilePOSIX_genType type
type MoqMustCompilePOSIX_genType_mock struct {
	Moq *MoqMustCompilePOSIX_genType
}

// MoqMustCompilePOSIX_genType_params holds the params of the
// MustCompilePOSIX_genType type
type MoqMustCompilePOSIX_genType_params struct{ Str string }

// MoqMustCompilePOSIX_genType_paramsKey holds the map key params of the
// MustCompilePOSIX_genType type
type MoqMustCompilePOSIX_genType_paramsKey struct {
	Params struct{ Str string }
	Hashes struct{ Str hash.Hash }
}

// MoqMustCompilePOSIX_genType_resultsByParams contains the results for a given
// set of parameters for the MustCompilePOSIX_genType type
type MoqMustCompilePOSIX_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqMustCompilePOSIX_genType_paramsKey]*MoqMustCompilePOSIX_genType_results
}

// MoqMustCompilePOSIX_genType_doFn defines the type of function needed when
// calling AndDo for the MustCompilePOSIX_genType type
type MoqMustCompilePOSIX_genType_doFn func(str string)

// MoqMustCompilePOSIX_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the MustCompilePOSIX_genType type
type MoqMustCompilePOSIX_genType_doReturnFn func(str string) *regexp.Regexp

// MoqMustCompilePOSIX_genType_results holds the results of the
// MustCompilePOSIX_genType type
type MoqMustCompilePOSIX_genType_results struct {
	Params  MoqMustCompilePOSIX_genType_params
	Results []struct {
		Values *struct {
			Result1 *regexp.Regexp
		}
		Sequence   uint32
		DoFn       MoqMustCompilePOSIX_genType_doFn
		DoReturnFn MoqMustCompilePOSIX_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqMustCompilePOSIX_genType_fnRecorder routes recorded function calls to the
// MoqMustCompilePOSIX_genType moq
type MoqMustCompilePOSIX_genType_fnRecorder struct {
	Params    MoqMustCompilePOSIX_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqMustCompilePOSIX_genType_results
	Moq       *MoqMustCompilePOSIX_genType
}

// MoqMustCompilePOSIX_genType_anyParams isolates the any params functions of
// the MustCompilePOSIX_genType type
type MoqMustCompilePOSIX_genType_anyParams struct {
	Recorder *MoqMustCompilePOSIX_genType_fnRecorder
}

// NewMoqMustCompilePOSIX_genType creates a new moq of the
// MustCompilePOSIX_genType type
func NewMoqMustCompilePOSIX_genType(scene *moq.Scene, config *moq.Config) *MoqMustCompilePOSIX_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqMustCompilePOSIX_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqMustCompilePOSIX_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Str moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Str moq.ParamIndexing
		}{
			Str: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the MustCompilePOSIX_genType type
func (m *MoqMustCompilePOSIX_genType) Mock() MustCompilePOSIX_genType {
	return func(str string) *regexp.Regexp {
		m.Scene.T.Helper()
		moq := &MoqMustCompilePOSIX_genType_mock{Moq: m}
		return moq.Fn(str)
	}
}

func (m *MoqMustCompilePOSIX_genType_mock) Fn(str string) (result1 *regexp.Regexp) {
	m.Moq.Scene.T.Helper()
	params := MoqMustCompilePOSIX_genType_params{
		Str: str,
	}
	var results *MoqMustCompilePOSIX_genType_results
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
		result.DoFn(str)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(str)
	}
	return
}

func (m *MoqMustCompilePOSIX_genType) OnCall(str string) *MoqMustCompilePOSIX_genType_fnRecorder {
	return &MoqMustCompilePOSIX_genType_fnRecorder{
		Params: MoqMustCompilePOSIX_genType_params{
			Str: str,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqMustCompilePOSIX_genType_fnRecorder) Any() *MoqMustCompilePOSIX_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqMustCompilePOSIX_genType_anyParams{Recorder: r}
}

func (a *MoqMustCompilePOSIX_genType_anyParams) Str() *MoqMustCompilePOSIX_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqMustCompilePOSIX_genType_fnRecorder) Seq() *MoqMustCompilePOSIX_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqMustCompilePOSIX_genType_fnRecorder) NoSeq() *MoqMustCompilePOSIX_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqMustCompilePOSIX_genType_fnRecorder) ReturnResults(result1 *regexp.Regexp) *MoqMustCompilePOSIX_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *regexp.Regexp
		}
		Sequence   uint32
		DoFn       MoqMustCompilePOSIX_genType_doFn
		DoReturnFn MoqMustCompilePOSIX_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *regexp.Regexp
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqMustCompilePOSIX_genType_fnRecorder) AndDo(fn MoqMustCompilePOSIX_genType_doFn) *MoqMustCompilePOSIX_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqMustCompilePOSIX_genType_fnRecorder) DoReturnResults(fn MoqMustCompilePOSIX_genType_doReturnFn) *MoqMustCompilePOSIX_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *regexp.Regexp
		}
		Sequence   uint32
		DoFn       MoqMustCompilePOSIX_genType_doFn
		DoReturnFn MoqMustCompilePOSIX_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqMustCompilePOSIX_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqMustCompilePOSIX_genType_resultsByParams
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
		results = &MoqMustCompilePOSIX_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqMustCompilePOSIX_genType_paramsKey]*MoqMustCompilePOSIX_genType_results{},
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
		r.Results = &MoqMustCompilePOSIX_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqMustCompilePOSIX_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqMustCompilePOSIX_genType_fnRecorder {
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
					Result1 *regexp.Regexp
				}
				Sequence   uint32
				DoFn       MoqMustCompilePOSIX_genType_doFn
				DoReturnFn MoqMustCompilePOSIX_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqMustCompilePOSIX_genType) PrettyParams(params MoqMustCompilePOSIX_genType_params) string {
	return fmt.Sprintf("MustCompilePOSIX_genType(%#v)", params.Str)
}

func (m *MoqMustCompilePOSIX_genType) ParamsKey(params MoqMustCompilePOSIX_genType_params, anyParams uint64) MoqMustCompilePOSIX_genType_paramsKey {
	m.Scene.T.Helper()
	var strUsed string
	var strUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Str == moq.ParamIndexByValue {
			strUsed = params.Str
		} else {
			strUsedHash = hash.DeepHash(params.Str)
		}
	}
	return MoqMustCompilePOSIX_genType_paramsKey{
		Params: struct{ Str string }{
			Str: strUsed,
		},
		Hashes: struct{ Str hash.Hash }{
			Str: strUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqMustCompilePOSIX_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqMustCompilePOSIX_genType) AssertExpectationsMet() {
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
// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package mime

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// ParseMediaType_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type ParseMediaType_genType func(v string) (mediatype string, params map[string]string, err error)

// MoqParseMediaType_genType holds the state of a moq of the
// ParseMediaType_genType type
type MoqParseMediaType_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqParseMediaType_genType_mock

	ResultsByParams []MoqParseMediaType_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			V moq.ParamIndexing
		}
	}
}

// MoqParseMediaType_genType_mock isolates the mock interface of the
// ParseMediaType_genType type
type MoqParseMediaType_genType_mock struct {
	Moq *MoqParseMediaType_genType
}

// MoqParseMediaType_genType_params holds the params of the
// ParseMediaType_genType type
type MoqParseMediaType_genType_params struct{ V string }

// MoqParseMediaType_genType_paramsKey holds the map key params of the
// ParseMediaType_genType type
type MoqParseMediaType_genType_paramsKey struct {
	Params struct{ V string }
	Hashes struct{ V hash.Hash }
}

// MoqParseMediaType_genType_resultsByParams contains the results for a given
// set of parameters for the ParseMediaType_genType type
type MoqParseMediaType_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqParseMediaType_genType_paramsKey]*MoqParseMediaType_genType_results
}

// MoqParseMediaType_genType_doFn defines the type of function needed when
// calling AndDo for the ParseMediaType_genType type
type MoqParseMediaType_genType_doFn func(v string)

// MoqParseMediaType_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the ParseMediaType_genType type
type MoqParseMediaType_genType_doReturnFn func(v string) (mediatype string, params map[string]string, err error)

// MoqParseMediaType_genType_results holds the results of the
// ParseMediaType_genType type
type MoqParseMediaType_genType_results struct {
	Params  MoqParseMediaType_genType_params
	Results []struct {
		Values *struct {
			Mediatype string
			Result2   map[string]string
			Err       error
		}
		Sequence   uint32
		DoFn       MoqParseMediaType_genType_doFn
		DoReturnFn MoqParseMediaType_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqParseMediaType_genType_fnRecorder routes recorded function calls to the
// MoqParseMediaType_genType moq
type MoqParseMediaType_genType_fnRecorder struct {
	Params    MoqParseMediaType_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqParseMediaType_genType_results
	Moq       *MoqParseMediaType_genType
}

// MoqParseMediaType_genType_anyParams isolates the any params functions of the
// ParseMediaType_genType type
type MoqParseMediaType_genType_anyParams struct {
	Recorder *MoqParseMediaType_genType_fnRecorder
}

// NewMoqParseMediaType_genType creates a new moq of the ParseMediaType_genType
// type
func NewMoqParseMediaType_genType(scene *moq.Scene, config *moq.Config) *MoqParseMediaType_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqParseMediaType_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqParseMediaType_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				V moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			V moq.ParamIndexing
		}{
			V: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ParseMediaType_genType type
func (m *MoqParseMediaType_genType) Mock() ParseMediaType_genType {
	return func(v string) (_ string, _ map[string]string, _ error) {
		m.Scene.T.Helper()
		moq := &MoqParseMediaType_genType_mock{Moq: m}
		return moq.Fn(v)
	}
}

func (m *MoqParseMediaType_genType_mock) Fn(v string) (mediatype string, result2 map[string]string, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqParseMediaType_genType_params{
		V: v,
	}
	var results *MoqParseMediaType_genType_results
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
		result.DoFn(v)
	}

	if result.Values != nil {
		mediatype = result.Values.Mediatype
		result2 = result.Values.Result2
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		mediatype, result2, err = result.DoReturnFn(v)
	}
	return
}

func (m *MoqParseMediaType_genType) OnCall(v string) *MoqParseMediaType_genType_fnRecorder {
	return &MoqParseMediaType_genType_fnRecorder{
		Params: MoqParseMediaType_genType_params{
			V: v,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqParseMediaType_genType_fnRecorder) Any() *MoqParseMediaType_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqParseMediaType_genType_anyParams{Recorder: r}
}

func (a *MoqParseMediaType_genType_anyParams) V() *MoqParseMediaType_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqParseMediaType_genType_fnRecorder) Seq() *MoqParseMediaType_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqParseMediaType_genType_fnRecorder) NoSeq() *MoqParseMediaType_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqParseMediaType_genType_fnRecorder) ReturnResults(mediatype string, result2 map[string]string, err error) *MoqParseMediaType_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Mediatype string
			Result2   map[string]string
			Err       error
		}
		Sequence   uint32
		DoFn       MoqParseMediaType_genType_doFn
		DoReturnFn MoqParseMediaType_genType_doReturnFn
	}{
		Values: &struct {
			Mediatype string
			Result2   map[string]string
			Err       error
		}{
			Mediatype: mediatype,
			Result2:   result2,
			Err:       err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqParseMediaType_genType_fnRecorder) AndDo(fn MoqParseMediaType_genType_doFn) *MoqParseMediaType_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqParseMediaType_genType_fnRecorder) DoReturnResults(fn MoqParseMediaType_genType_doReturnFn) *MoqParseMediaType_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Mediatype string
			Result2   map[string]string
			Err       error
		}
		Sequence   uint32
		DoFn       MoqParseMediaType_genType_doFn
		DoReturnFn MoqParseMediaType_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqParseMediaType_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqParseMediaType_genType_resultsByParams
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
		results = &MoqParseMediaType_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqParseMediaType_genType_paramsKey]*MoqParseMediaType_genType_results{},
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
		r.Results = &MoqParseMediaType_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqParseMediaType_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqParseMediaType_genType_fnRecorder {
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
					Mediatype string
					Result2   map[string]string
					Err       error
				}
				Sequence   uint32
				DoFn       MoqParseMediaType_genType_doFn
				DoReturnFn MoqParseMediaType_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqParseMediaType_genType) PrettyParams(params MoqParseMediaType_genType_params) string {
	return fmt.Sprintf("ParseMediaType_genType(%#v)", params.V)
}

func (m *MoqParseMediaType_genType) ParamsKey(params MoqParseMediaType_genType_params, anyParams uint64) MoqParseMediaType_genType_paramsKey {
	m.Scene.T.Helper()
	var vUsed string
	var vUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.V == moq.ParamIndexByValue {
			vUsed = params.V
		} else {
			vUsedHash = hash.DeepHash(params.V)
		}
	}
	return MoqParseMediaType_genType_paramsKey{
		Params: struct{ V string }{
			V: vUsed,
		},
		Hashes: struct{ V hash.Hash }{
			V: vUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqParseMediaType_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqParseMediaType_genType) AssertExpectationsMet() {
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

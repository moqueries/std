// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package json

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that json.Encoder_starGenType is mocked
// completely
var _ Encoder_starGenType = (*MoqEncoder_starGenType_mock)(nil)

// Encoder_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type Encoder_starGenType interface {
	Encode(v any) error
	SetIndent(prefix, indent string)
	SetEscapeHTML(on bool)
}

// MoqEncoder_starGenType holds the state of a moq of the Encoder_starGenType
// type
type MoqEncoder_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqEncoder_starGenType_mock

	ResultsByParams_Encode        []MoqEncoder_starGenType_Encode_resultsByParams
	ResultsByParams_SetIndent     []MoqEncoder_starGenType_SetIndent_resultsByParams
	ResultsByParams_SetEscapeHTML []MoqEncoder_starGenType_SetEscapeHTML_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Encode struct {
				V moq.ParamIndexing
			}
			SetIndent struct {
				Prefix moq.ParamIndexing
				Indent moq.ParamIndexing
			}
			SetEscapeHTML struct {
				On moq.ParamIndexing
			}
		}
	}
	// MoqEncoder_starGenType_mock isolates the mock interface of the
}

// Encoder_starGenType type
type MoqEncoder_starGenType_mock struct {
	Moq *MoqEncoder_starGenType
}

// MoqEncoder_starGenType_recorder isolates the recorder interface of the
// Encoder_starGenType type
type MoqEncoder_starGenType_recorder struct {
	Moq *MoqEncoder_starGenType
}

// MoqEncoder_starGenType_Encode_params holds the params of the
// Encoder_starGenType type
type MoqEncoder_starGenType_Encode_params struct{ V any }

// MoqEncoder_starGenType_Encode_paramsKey holds the map key params of the
// Encoder_starGenType type
type MoqEncoder_starGenType_Encode_paramsKey struct {
	Params struct{ V any }
	Hashes struct{ V hash.Hash }
}

// MoqEncoder_starGenType_Encode_resultsByParams contains the results for a
// given set of parameters for the Encoder_starGenType type
type MoqEncoder_starGenType_Encode_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqEncoder_starGenType_Encode_paramsKey]*MoqEncoder_starGenType_Encode_results
}

// MoqEncoder_starGenType_Encode_doFn defines the type of function needed when
// calling AndDo for the Encoder_starGenType type
type MoqEncoder_starGenType_Encode_doFn func(v any)

// MoqEncoder_starGenType_Encode_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Encoder_starGenType type
type MoqEncoder_starGenType_Encode_doReturnFn func(v any) error

// MoqEncoder_starGenType_Encode_results holds the results of the
// Encoder_starGenType type
type MoqEncoder_starGenType_Encode_results struct {
	Params  MoqEncoder_starGenType_Encode_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqEncoder_starGenType_Encode_doFn
		DoReturnFn MoqEncoder_starGenType_Encode_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqEncoder_starGenType_Encode_fnRecorder routes recorded function calls to
// the MoqEncoder_starGenType moq
type MoqEncoder_starGenType_Encode_fnRecorder struct {
	Params    MoqEncoder_starGenType_Encode_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqEncoder_starGenType_Encode_results
	Moq       *MoqEncoder_starGenType
}

// MoqEncoder_starGenType_Encode_anyParams isolates the any params functions of
// the Encoder_starGenType type
type MoqEncoder_starGenType_Encode_anyParams struct {
	Recorder *MoqEncoder_starGenType_Encode_fnRecorder
}

// MoqEncoder_starGenType_SetIndent_params holds the params of the
// Encoder_starGenType type
type MoqEncoder_starGenType_SetIndent_params struct{ Prefix, Indent string }

// MoqEncoder_starGenType_SetIndent_paramsKey holds the map key params of the
// Encoder_starGenType type
type MoqEncoder_starGenType_SetIndent_paramsKey struct {
	Params struct{ Prefix, Indent string }
	Hashes struct{ Prefix, Indent hash.Hash }
}

// MoqEncoder_starGenType_SetIndent_resultsByParams contains the results for a
// given set of parameters for the Encoder_starGenType type
type MoqEncoder_starGenType_SetIndent_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqEncoder_starGenType_SetIndent_paramsKey]*MoqEncoder_starGenType_SetIndent_results
}

// MoqEncoder_starGenType_SetIndent_doFn defines the type of function needed
// when calling AndDo for the Encoder_starGenType type
type MoqEncoder_starGenType_SetIndent_doFn func(prefix, indent string)

// MoqEncoder_starGenType_SetIndent_doReturnFn defines the type of function
// needed when calling DoReturnResults for the Encoder_starGenType type
type MoqEncoder_starGenType_SetIndent_doReturnFn func(prefix, indent string)

// MoqEncoder_starGenType_SetIndent_results holds the results of the
// Encoder_starGenType type
type MoqEncoder_starGenType_SetIndent_results struct {
	Params  MoqEncoder_starGenType_SetIndent_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqEncoder_starGenType_SetIndent_doFn
		DoReturnFn MoqEncoder_starGenType_SetIndent_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqEncoder_starGenType_SetIndent_fnRecorder routes recorded function calls
// to the MoqEncoder_starGenType moq
type MoqEncoder_starGenType_SetIndent_fnRecorder struct {
	Params    MoqEncoder_starGenType_SetIndent_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqEncoder_starGenType_SetIndent_results
	Moq       *MoqEncoder_starGenType
}

// MoqEncoder_starGenType_SetIndent_anyParams isolates the any params functions
// of the Encoder_starGenType type
type MoqEncoder_starGenType_SetIndent_anyParams struct {
	Recorder *MoqEncoder_starGenType_SetIndent_fnRecorder
}

// MoqEncoder_starGenType_SetEscapeHTML_params holds the params of the
// Encoder_starGenType type
type MoqEncoder_starGenType_SetEscapeHTML_params struct{ On bool }

// MoqEncoder_starGenType_SetEscapeHTML_paramsKey holds the map key params of
// the Encoder_starGenType type
type MoqEncoder_starGenType_SetEscapeHTML_paramsKey struct {
	Params struct{ On bool }
	Hashes struct{ On hash.Hash }
}

// MoqEncoder_starGenType_SetEscapeHTML_resultsByParams contains the results
// for a given set of parameters for the Encoder_starGenType type
type MoqEncoder_starGenType_SetEscapeHTML_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqEncoder_starGenType_SetEscapeHTML_paramsKey]*MoqEncoder_starGenType_SetEscapeHTML_results
}

// MoqEncoder_starGenType_SetEscapeHTML_doFn defines the type of function
// needed when calling AndDo for the Encoder_starGenType type
type MoqEncoder_starGenType_SetEscapeHTML_doFn func(on bool)

// MoqEncoder_starGenType_SetEscapeHTML_doReturnFn defines the type of function
// needed when calling DoReturnResults for the Encoder_starGenType type
type MoqEncoder_starGenType_SetEscapeHTML_doReturnFn func(on bool)

// MoqEncoder_starGenType_SetEscapeHTML_results holds the results of the
// Encoder_starGenType type
type MoqEncoder_starGenType_SetEscapeHTML_results struct {
	Params  MoqEncoder_starGenType_SetEscapeHTML_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqEncoder_starGenType_SetEscapeHTML_doFn
		DoReturnFn MoqEncoder_starGenType_SetEscapeHTML_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqEncoder_starGenType_SetEscapeHTML_fnRecorder routes recorded function
// calls to the MoqEncoder_starGenType moq
type MoqEncoder_starGenType_SetEscapeHTML_fnRecorder struct {
	Params    MoqEncoder_starGenType_SetEscapeHTML_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqEncoder_starGenType_SetEscapeHTML_results
	Moq       *MoqEncoder_starGenType
}

// MoqEncoder_starGenType_SetEscapeHTML_anyParams isolates the any params
// functions of the Encoder_starGenType type
type MoqEncoder_starGenType_SetEscapeHTML_anyParams struct {
	Recorder *MoqEncoder_starGenType_SetEscapeHTML_fnRecorder
}

// NewMoqEncoder_starGenType creates a new moq of the Encoder_starGenType type
func NewMoqEncoder_starGenType(scene *moq.Scene, config *moq.Config) *MoqEncoder_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqEncoder_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqEncoder_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Encode struct {
					V moq.ParamIndexing
				}
				SetIndent struct {
					Prefix moq.ParamIndexing
					Indent moq.ParamIndexing
				}
				SetEscapeHTML struct {
					On moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Encode struct {
				V moq.ParamIndexing
			}
			SetIndent struct {
				Prefix moq.ParamIndexing
				Indent moq.ParamIndexing
			}
			SetEscapeHTML struct {
				On moq.ParamIndexing
			}
		}{
			Encode: struct {
				V moq.ParamIndexing
			}{
				V: moq.ParamIndexByValue,
			},
			SetIndent: struct {
				Prefix moq.ParamIndexing
				Indent moq.ParamIndexing
			}{
				Prefix: moq.ParamIndexByValue,
				Indent: moq.ParamIndexByValue,
			},
			SetEscapeHTML: struct {
				On moq.ParamIndexing
			}{
				On: moq.ParamIndexByValue,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Encoder_starGenType type
func (m *MoqEncoder_starGenType) Mock() *MoqEncoder_starGenType_mock { return m.Moq }

func (m *MoqEncoder_starGenType_mock) Encode(v any) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqEncoder_starGenType_Encode_params{
		V: v,
	}
	var results *MoqEncoder_starGenType_Encode_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Encode {
		paramsKey := m.Moq.ParamsKey_Encode(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Encode(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Encode(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Encode(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(v)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(v)
	}
	return
}

func (m *MoqEncoder_starGenType_mock) SetIndent(prefix, indent string) {
	m.Moq.Scene.T.Helper()
	params := MoqEncoder_starGenType_SetIndent_params{
		Prefix: prefix,
		Indent: indent,
	}
	var results *MoqEncoder_starGenType_SetIndent_results
	for _, resultsByParams := range m.Moq.ResultsByParams_SetIndent {
		paramsKey := m.Moq.ParamsKey_SetIndent(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_SetIndent(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_SetIndent(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_SetIndent(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(prefix, indent)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(prefix, indent)
	}
	return
}

func (m *MoqEncoder_starGenType_mock) SetEscapeHTML(on bool) {
	m.Moq.Scene.T.Helper()
	params := MoqEncoder_starGenType_SetEscapeHTML_params{
		On: on,
	}
	var results *MoqEncoder_starGenType_SetEscapeHTML_results
	for _, resultsByParams := range m.Moq.ResultsByParams_SetEscapeHTML {
		paramsKey := m.Moq.ParamsKey_SetEscapeHTML(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_SetEscapeHTML(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_SetEscapeHTML(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_SetEscapeHTML(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(on)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(on)
	}
	return
}

// OnCall returns the recorder implementation of the Encoder_starGenType type
func (m *MoqEncoder_starGenType) OnCall() *MoqEncoder_starGenType_recorder {
	return &MoqEncoder_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqEncoder_starGenType_recorder) Encode(v any) *MoqEncoder_starGenType_Encode_fnRecorder {
	return &MoqEncoder_starGenType_Encode_fnRecorder{
		Params: MoqEncoder_starGenType_Encode_params{
			V: v,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqEncoder_starGenType_Encode_fnRecorder) Any() *MoqEncoder_starGenType_Encode_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Encode(r.Params))
		return nil
	}
	return &MoqEncoder_starGenType_Encode_anyParams{Recorder: r}
}

func (a *MoqEncoder_starGenType_Encode_anyParams) V() *MoqEncoder_starGenType_Encode_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqEncoder_starGenType_Encode_fnRecorder) Seq() *MoqEncoder_starGenType_Encode_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Encode(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqEncoder_starGenType_Encode_fnRecorder) NoSeq() *MoqEncoder_starGenType_Encode_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Encode(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqEncoder_starGenType_Encode_fnRecorder) ReturnResults(result1 error) *MoqEncoder_starGenType_Encode_fnRecorder {
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
		DoFn       MoqEncoder_starGenType_Encode_doFn
		DoReturnFn MoqEncoder_starGenType_Encode_doReturnFn
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

func (r *MoqEncoder_starGenType_Encode_fnRecorder) AndDo(fn MoqEncoder_starGenType_Encode_doFn) *MoqEncoder_starGenType_Encode_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqEncoder_starGenType_Encode_fnRecorder) DoReturnResults(fn MoqEncoder_starGenType_Encode_doReturnFn) *MoqEncoder_starGenType_Encode_fnRecorder {
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
		DoFn       MoqEncoder_starGenType_Encode_doFn
		DoReturnFn MoqEncoder_starGenType_Encode_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqEncoder_starGenType_Encode_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqEncoder_starGenType_Encode_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Encode {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqEncoder_starGenType_Encode_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqEncoder_starGenType_Encode_paramsKey]*MoqEncoder_starGenType_Encode_results{},
		}
		r.Moq.ResultsByParams_Encode = append(r.Moq.ResultsByParams_Encode, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Encode) {
			copy(r.Moq.ResultsByParams_Encode[insertAt+1:], r.Moq.ResultsByParams_Encode[insertAt:0])
			r.Moq.ResultsByParams_Encode[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Encode(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqEncoder_starGenType_Encode_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqEncoder_starGenType_Encode_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqEncoder_starGenType_Encode_fnRecorder {
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
				DoFn       MoqEncoder_starGenType_Encode_doFn
				DoReturnFn MoqEncoder_starGenType_Encode_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqEncoder_starGenType) PrettyParams_Encode(params MoqEncoder_starGenType_Encode_params) string {
	return fmt.Sprintf("Encode(%#v)", params.V)
}

func (m *MoqEncoder_starGenType) ParamsKey_Encode(params MoqEncoder_starGenType_Encode_params, anyParams uint64) MoqEncoder_starGenType_Encode_paramsKey {
	m.Scene.T.Helper()
	var vUsed any
	var vUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Encode.V == moq.ParamIndexByValue {
			vUsed = params.V
		} else {
			vUsedHash = hash.DeepHash(params.V)
		}
	}
	return MoqEncoder_starGenType_Encode_paramsKey{
		Params: struct{ V any }{
			V: vUsed,
		},
		Hashes: struct{ V hash.Hash }{
			V: vUsedHash,
		},
	}
}

func (m *MoqEncoder_starGenType_recorder) SetIndent(prefix, indent string) *MoqEncoder_starGenType_SetIndent_fnRecorder {
	return &MoqEncoder_starGenType_SetIndent_fnRecorder{
		Params: MoqEncoder_starGenType_SetIndent_params{
			Prefix: prefix,
			Indent: indent,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqEncoder_starGenType_SetIndent_fnRecorder) Any() *MoqEncoder_starGenType_SetIndent_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_SetIndent(r.Params))
		return nil
	}
	return &MoqEncoder_starGenType_SetIndent_anyParams{Recorder: r}
}

func (a *MoqEncoder_starGenType_SetIndent_anyParams) Prefix() *MoqEncoder_starGenType_SetIndent_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqEncoder_starGenType_SetIndent_anyParams) Indent() *MoqEncoder_starGenType_SetIndent_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqEncoder_starGenType_SetIndent_fnRecorder) Seq() *MoqEncoder_starGenType_SetIndent_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_SetIndent(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqEncoder_starGenType_SetIndent_fnRecorder) NoSeq() *MoqEncoder_starGenType_SetIndent_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_SetIndent(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqEncoder_starGenType_SetIndent_fnRecorder) ReturnResults() *MoqEncoder_starGenType_SetIndent_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqEncoder_starGenType_SetIndent_doFn
		DoReturnFn MoqEncoder_starGenType_SetIndent_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqEncoder_starGenType_SetIndent_fnRecorder) AndDo(fn MoqEncoder_starGenType_SetIndent_doFn) *MoqEncoder_starGenType_SetIndent_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqEncoder_starGenType_SetIndent_fnRecorder) DoReturnResults(fn MoqEncoder_starGenType_SetIndent_doReturnFn) *MoqEncoder_starGenType_SetIndent_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqEncoder_starGenType_SetIndent_doFn
		DoReturnFn MoqEncoder_starGenType_SetIndent_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqEncoder_starGenType_SetIndent_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqEncoder_starGenType_SetIndent_resultsByParams
	for n, res := range r.Moq.ResultsByParams_SetIndent {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqEncoder_starGenType_SetIndent_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqEncoder_starGenType_SetIndent_paramsKey]*MoqEncoder_starGenType_SetIndent_results{},
		}
		r.Moq.ResultsByParams_SetIndent = append(r.Moq.ResultsByParams_SetIndent, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_SetIndent) {
			copy(r.Moq.ResultsByParams_SetIndent[insertAt+1:], r.Moq.ResultsByParams_SetIndent[insertAt:0])
			r.Moq.ResultsByParams_SetIndent[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_SetIndent(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqEncoder_starGenType_SetIndent_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqEncoder_starGenType_SetIndent_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqEncoder_starGenType_SetIndent_fnRecorder {
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
				DoFn       MoqEncoder_starGenType_SetIndent_doFn
				DoReturnFn MoqEncoder_starGenType_SetIndent_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqEncoder_starGenType) PrettyParams_SetIndent(params MoqEncoder_starGenType_SetIndent_params) string {
	return fmt.Sprintf("SetIndent(%#v, %#v)", params.Prefix, params.Indent)
}

func (m *MoqEncoder_starGenType) ParamsKey_SetIndent(params MoqEncoder_starGenType_SetIndent_params, anyParams uint64) MoqEncoder_starGenType_SetIndent_paramsKey {
	m.Scene.T.Helper()
	var prefixUsed string
	var prefixUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.SetIndent.Prefix == moq.ParamIndexByValue {
			prefixUsed = params.Prefix
		} else {
			prefixUsedHash = hash.DeepHash(params.Prefix)
		}
	}
	var indentUsed string
	var indentUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.SetIndent.Indent == moq.ParamIndexByValue {
			indentUsed = params.Indent
		} else {
			indentUsedHash = hash.DeepHash(params.Indent)
		}
	}
	return MoqEncoder_starGenType_SetIndent_paramsKey{
		Params: struct{ Prefix, Indent string }{
			Prefix: prefixUsed,
			Indent: indentUsed,
		},
		Hashes: struct{ Prefix, Indent hash.Hash }{
			Prefix: prefixUsedHash,
			Indent: indentUsedHash,
		},
	}
}

func (m *MoqEncoder_starGenType_recorder) SetEscapeHTML(on bool) *MoqEncoder_starGenType_SetEscapeHTML_fnRecorder {
	return &MoqEncoder_starGenType_SetEscapeHTML_fnRecorder{
		Params: MoqEncoder_starGenType_SetEscapeHTML_params{
			On: on,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqEncoder_starGenType_SetEscapeHTML_fnRecorder) Any() *MoqEncoder_starGenType_SetEscapeHTML_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_SetEscapeHTML(r.Params))
		return nil
	}
	return &MoqEncoder_starGenType_SetEscapeHTML_anyParams{Recorder: r}
}

func (a *MoqEncoder_starGenType_SetEscapeHTML_anyParams) On() *MoqEncoder_starGenType_SetEscapeHTML_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqEncoder_starGenType_SetEscapeHTML_fnRecorder) Seq() *MoqEncoder_starGenType_SetEscapeHTML_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_SetEscapeHTML(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqEncoder_starGenType_SetEscapeHTML_fnRecorder) NoSeq() *MoqEncoder_starGenType_SetEscapeHTML_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_SetEscapeHTML(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqEncoder_starGenType_SetEscapeHTML_fnRecorder) ReturnResults() *MoqEncoder_starGenType_SetEscapeHTML_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqEncoder_starGenType_SetEscapeHTML_doFn
		DoReturnFn MoqEncoder_starGenType_SetEscapeHTML_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqEncoder_starGenType_SetEscapeHTML_fnRecorder) AndDo(fn MoqEncoder_starGenType_SetEscapeHTML_doFn) *MoqEncoder_starGenType_SetEscapeHTML_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqEncoder_starGenType_SetEscapeHTML_fnRecorder) DoReturnResults(fn MoqEncoder_starGenType_SetEscapeHTML_doReturnFn) *MoqEncoder_starGenType_SetEscapeHTML_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqEncoder_starGenType_SetEscapeHTML_doFn
		DoReturnFn MoqEncoder_starGenType_SetEscapeHTML_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqEncoder_starGenType_SetEscapeHTML_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqEncoder_starGenType_SetEscapeHTML_resultsByParams
	for n, res := range r.Moq.ResultsByParams_SetEscapeHTML {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqEncoder_starGenType_SetEscapeHTML_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqEncoder_starGenType_SetEscapeHTML_paramsKey]*MoqEncoder_starGenType_SetEscapeHTML_results{},
		}
		r.Moq.ResultsByParams_SetEscapeHTML = append(r.Moq.ResultsByParams_SetEscapeHTML, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_SetEscapeHTML) {
			copy(r.Moq.ResultsByParams_SetEscapeHTML[insertAt+1:], r.Moq.ResultsByParams_SetEscapeHTML[insertAt:0])
			r.Moq.ResultsByParams_SetEscapeHTML[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_SetEscapeHTML(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqEncoder_starGenType_SetEscapeHTML_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqEncoder_starGenType_SetEscapeHTML_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqEncoder_starGenType_SetEscapeHTML_fnRecorder {
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
				DoFn       MoqEncoder_starGenType_SetEscapeHTML_doFn
				DoReturnFn MoqEncoder_starGenType_SetEscapeHTML_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqEncoder_starGenType) PrettyParams_SetEscapeHTML(params MoqEncoder_starGenType_SetEscapeHTML_params) string {
	return fmt.Sprintf("SetEscapeHTML(%#v)", params.On)
}

func (m *MoqEncoder_starGenType) ParamsKey_SetEscapeHTML(params MoqEncoder_starGenType_SetEscapeHTML_params, anyParams uint64) MoqEncoder_starGenType_SetEscapeHTML_paramsKey {
	m.Scene.T.Helper()
	var onUsed bool
	var onUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.SetEscapeHTML.On == moq.ParamIndexByValue {
			onUsed = params.On
		} else {
			onUsedHash = hash.DeepHash(params.On)
		}
	}
	return MoqEncoder_starGenType_SetEscapeHTML_paramsKey{
		Params: struct{ On bool }{
			On: onUsed,
		},
		Hashes: struct{ On hash.Hash }{
			On: onUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqEncoder_starGenType) Reset() {
	m.ResultsByParams_Encode = nil
	m.ResultsByParams_SetIndent = nil
	m.ResultsByParams_SetEscapeHTML = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqEncoder_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Encode {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Encode(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_SetIndent {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_SetIndent(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_SetEscapeHTML {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_SetEscapeHTML(results.Params))
			}
		}
	}
}

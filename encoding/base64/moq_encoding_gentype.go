// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package base64

import (
	"encoding/base64"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that base64.Encoding_genType is mocked
// completely
var _ Encoding_genType = (*MoqEncoding_genType_mock)(nil)

// Encoding_genType is the fabricated implementation type of this mock (emitted
// when mocking a collections of methods directly and not from an interface
// type)
type Encoding_genType interface {
	WithPadding(padding rune) *base64.Encoding
	Strict() *base64.Encoding
}

// MoqEncoding_genType holds the state of a moq of the Encoding_genType type
type MoqEncoding_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqEncoding_genType_mock

	ResultsByParams_WithPadding []MoqEncoding_genType_WithPadding_resultsByParams
	ResultsByParams_Strict      []MoqEncoding_genType_Strict_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			WithPadding struct {
				Padding moq.ParamIndexing
			}
			Strict struct{}
		}
	}
}

// MoqEncoding_genType_mock isolates the mock interface of the Encoding_genType
// type
type MoqEncoding_genType_mock struct {
	Moq *MoqEncoding_genType
}

// MoqEncoding_genType_recorder isolates the recorder interface of the
// Encoding_genType type
type MoqEncoding_genType_recorder struct {
	Moq *MoqEncoding_genType
}

// MoqEncoding_genType_WithPadding_params holds the params of the
// Encoding_genType type
type MoqEncoding_genType_WithPadding_params struct{ Padding rune }

// MoqEncoding_genType_WithPadding_paramsKey holds the map key params of the
// Encoding_genType type
type MoqEncoding_genType_WithPadding_paramsKey struct {
	Params struct{ Padding rune }
	Hashes struct{ Padding hash.Hash }
}

// MoqEncoding_genType_WithPadding_resultsByParams contains the results for a
// given set of parameters for the Encoding_genType type
type MoqEncoding_genType_WithPadding_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqEncoding_genType_WithPadding_paramsKey]*MoqEncoding_genType_WithPadding_results
}

// MoqEncoding_genType_WithPadding_doFn defines the type of function needed
// when calling AndDo for the Encoding_genType type
type MoqEncoding_genType_WithPadding_doFn func(padding rune)

// MoqEncoding_genType_WithPadding_doReturnFn defines the type of function
// needed when calling DoReturnResults for the Encoding_genType type
type MoqEncoding_genType_WithPadding_doReturnFn func(padding rune) *base64.Encoding

// MoqEncoding_genType_WithPadding_results holds the results of the
// Encoding_genType type
type MoqEncoding_genType_WithPadding_results struct {
	Params  MoqEncoding_genType_WithPadding_params
	Results []struct {
		Values *struct {
			Result1 *base64.Encoding
		}
		Sequence   uint32
		DoFn       MoqEncoding_genType_WithPadding_doFn
		DoReturnFn MoqEncoding_genType_WithPadding_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqEncoding_genType_WithPadding_fnRecorder routes recorded function calls to
// the MoqEncoding_genType moq
type MoqEncoding_genType_WithPadding_fnRecorder struct {
	Params    MoqEncoding_genType_WithPadding_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqEncoding_genType_WithPadding_results
	Moq       *MoqEncoding_genType
}

// MoqEncoding_genType_WithPadding_anyParams isolates the any params functions
// of the Encoding_genType type
type MoqEncoding_genType_WithPadding_anyParams struct {
	Recorder *MoqEncoding_genType_WithPadding_fnRecorder
}

// MoqEncoding_genType_Strict_params holds the params of the Encoding_genType
// type
type MoqEncoding_genType_Strict_params struct{}

// MoqEncoding_genType_Strict_paramsKey holds the map key params of the
// Encoding_genType type
type MoqEncoding_genType_Strict_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqEncoding_genType_Strict_resultsByParams contains the results for a given
// set of parameters for the Encoding_genType type
type MoqEncoding_genType_Strict_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqEncoding_genType_Strict_paramsKey]*MoqEncoding_genType_Strict_results
}

// MoqEncoding_genType_Strict_doFn defines the type of function needed when
// calling AndDo for the Encoding_genType type
type MoqEncoding_genType_Strict_doFn func()

// MoqEncoding_genType_Strict_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Encoding_genType type
type MoqEncoding_genType_Strict_doReturnFn func() *base64.Encoding

// MoqEncoding_genType_Strict_results holds the results of the Encoding_genType
// type
type MoqEncoding_genType_Strict_results struct {
	Params  MoqEncoding_genType_Strict_params
	Results []struct {
		Values *struct {
			Result1 *base64.Encoding
		}
		Sequence   uint32
		DoFn       MoqEncoding_genType_Strict_doFn
		DoReturnFn MoqEncoding_genType_Strict_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqEncoding_genType_Strict_fnRecorder routes recorded function calls to the
// MoqEncoding_genType moq
type MoqEncoding_genType_Strict_fnRecorder struct {
	Params    MoqEncoding_genType_Strict_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqEncoding_genType_Strict_results
	Moq       *MoqEncoding_genType
}

// MoqEncoding_genType_Strict_anyParams isolates the any params functions of
// the Encoding_genType type
type MoqEncoding_genType_Strict_anyParams struct {
	Recorder *MoqEncoding_genType_Strict_fnRecorder
}

// NewMoqEncoding_genType creates a new moq of the Encoding_genType type
func NewMoqEncoding_genType(scene *moq.Scene, config *moq.Config) *MoqEncoding_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqEncoding_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqEncoding_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				WithPadding struct {
					Padding moq.ParamIndexing
				}
				Strict struct{}
			}
		}{ParameterIndexing: struct {
			WithPadding struct {
				Padding moq.ParamIndexing
			}
			Strict struct{}
		}{
			WithPadding: struct {
				Padding moq.ParamIndexing
			}{
				Padding: moq.ParamIndexByValue,
			},
			Strict: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Encoding_genType type
func (m *MoqEncoding_genType) Mock() *MoqEncoding_genType_mock { return m.Moq }

func (m *MoqEncoding_genType_mock) WithPadding(padding rune) (result1 *base64.Encoding) {
	m.Moq.Scene.T.Helper()
	params := MoqEncoding_genType_WithPadding_params{
		Padding: padding,
	}
	var results *MoqEncoding_genType_WithPadding_results
	for _, resultsByParams := range m.Moq.ResultsByParams_WithPadding {
		paramsKey := m.Moq.ParamsKey_WithPadding(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_WithPadding(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_WithPadding(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_WithPadding(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(padding)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(padding)
	}
	return
}

func (m *MoqEncoding_genType_mock) Strict() (result1 *base64.Encoding) {
	m.Moq.Scene.T.Helper()
	params := MoqEncoding_genType_Strict_params{}
	var results *MoqEncoding_genType_Strict_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Strict {
		paramsKey := m.Moq.ParamsKey_Strict(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Strict(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Strict(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Strict(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn()
	}
	return
}

// OnCall returns the recorder implementation of the Encoding_genType type
func (m *MoqEncoding_genType) OnCall() *MoqEncoding_genType_recorder {
	return &MoqEncoding_genType_recorder{
		Moq: m,
	}
}

func (m *MoqEncoding_genType_recorder) WithPadding(padding rune) *MoqEncoding_genType_WithPadding_fnRecorder {
	return &MoqEncoding_genType_WithPadding_fnRecorder{
		Params: MoqEncoding_genType_WithPadding_params{
			Padding: padding,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqEncoding_genType_WithPadding_fnRecorder) Any() *MoqEncoding_genType_WithPadding_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_WithPadding(r.Params))
		return nil
	}
	return &MoqEncoding_genType_WithPadding_anyParams{Recorder: r}
}

func (a *MoqEncoding_genType_WithPadding_anyParams) Padding() *MoqEncoding_genType_WithPadding_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqEncoding_genType_WithPadding_fnRecorder) Seq() *MoqEncoding_genType_WithPadding_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_WithPadding(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqEncoding_genType_WithPadding_fnRecorder) NoSeq() *MoqEncoding_genType_WithPadding_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_WithPadding(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqEncoding_genType_WithPadding_fnRecorder) ReturnResults(result1 *base64.Encoding) *MoqEncoding_genType_WithPadding_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *base64.Encoding
		}
		Sequence   uint32
		DoFn       MoqEncoding_genType_WithPadding_doFn
		DoReturnFn MoqEncoding_genType_WithPadding_doReturnFn
	}{
		Values: &struct {
			Result1 *base64.Encoding
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqEncoding_genType_WithPadding_fnRecorder) AndDo(fn MoqEncoding_genType_WithPadding_doFn) *MoqEncoding_genType_WithPadding_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqEncoding_genType_WithPadding_fnRecorder) DoReturnResults(fn MoqEncoding_genType_WithPadding_doReturnFn) *MoqEncoding_genType_WithPadding_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *base64.Encoding
		}
		Sequence   uint32
		DoFn       MoqEncoding_genType_WithPadding_doFn
		DoReturnFn MoqEncoding_genType_WithPadding_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqEncoding_genType_WithPadding_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqEncoding_genType_WithPadding_resultsByParams
	for n, res := range r.Moq.ResultsByParams_WithPadding {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqEncoding_genType_WithPadding_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqEncoding_genType_WithPadding_paramsKey]*MoqEncoding_genType_WithPadding_results{},
		}
		r.Moq.ResultsByParams_WithPadding = append(r.Moq.ResultsByParams_WithPadding, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_WithPadding) {
			copy(r.Moq.ResultsByParams_WithPadding[insertAt+1:], r.Moq.ResultsByParams_WithPadding[insertAt:0])
			r.Moq.ResultsByParams_WithPadding[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_WithPadding(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqEncoding_genType_WithPadding_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqEncoding_genType_WithPadding_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqEncoding_genType_WithPadding_fnRecorder {
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
					Result1 *base64.Encoding
				}
				Sequence   uint32
				DoFn       MoqEncoding_genType_WithPadding_doFn
				DoReturnFn MoqEncoding_genType_WithPadding_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqEncoding_genType) PrettyParams_WithPadding(params MoqEncoding_genType_WithPadding_params) string {
	return fmt.Sprintf("WithPadding(%#v)", params.Padding)
}

func (m *MoqEncoding_genType) ParamsKey_WithPadding(params MoqEncoding_genType_WithPadding_params, anyParams uint64) MoqEncoding_genType_WithPadding_paramsKey {
	m.Scene.T.Helper()
	var paddingUsed rune
	var paddingUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.WithPadding.Padding == moq.ParamIndexByValue {
			paddingUsed = params.Padding
		} else {
			paddingUsedHash = hash.DeepHash(params.Padding)
		}
	}
	return MoqEncoding_genType_WithPadding_paramsKey{
		Params: struct{ Padding rune }{
			Padding: paddingUsed,
		},
		Hashes: struct{ Padding hash.Hash }{
			Padding: paddingUsedHash,
		},
	}
}

func (m *MoqEncoding_genType_recorder) Strict() *MoqEncoding_genType_Strict_fnRecorder {
	return &MoqEncoding_genType_Strict_fnRecorder{
		Params:   MoqEncoding_genType_Strict_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqEncoding_genType_Strict_fnRecorder) Any() *MoqEncoding_genType_Strict_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Strict(r.Params))
		return nil
	}
	return &MoqEncoding_genType_Strict_anyParams{Recorder: r}
}

func (r *MoqEncoding_genType_Strict_fnRecorder) Seq() *MoqEncoding_genType_Strict_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Strict(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqEncoding_genType_Strict_fnRecorder) NoSeq() *MoqEncoding_genType_Strict_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Strict(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqEncoding_genType_Strict_fnRecorder) ReturnResults(result1 *base64.Encoding) *MoqEncoding_genType_Strict_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *base64.Encoding
		}
		Sequence   uint32
		DoFn       MoqEncoding_genType_Strict_doFn
		DoReturnFn MoqEncoding_genType_Strict_doReturnFn
	}{
		Values: &struct {
			Result1 *base64.Encoding
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqEncoding_genType_Strict_fnRecorder) AndDo(fn MoqEncoding_genType_Strict_doFn) *MoqEncoding_genType_Strict_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqEncoding_genType_Strict_fnRecorder) DoReturnResults(fn MoqEncoding_genType_Strict_doReturnFn) *MoqEncoding_genType_Strict_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *base64.Encoding
		}
		Sequence   uint32
		DoFn       MoqEncoding_genType_Strict_doFn
		DoReturnFn MoqEncoding_genType_Strict_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqEncoding_genType_Strict_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqEncoding_genType_Strict_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Strict {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqEncoding_genType_Strict_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqEncoding_genType_Strict_paramsKey]*MoqEncoding_genType_Strict_results{},
		}
		r.Moq.ResultsByParams_Strict = append(r.Moq.ResultsByParams_Strict, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Strict) {
			copy(r.Moq.ResultsByParams_Strict[insertAt+1:], r.Moq.ResultsByParams_Strict[insertAt:0])
			r.Moq.ResultsByParams_Strict[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Strict(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqEncoding_genType_Strict_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqEncoding_genType_Strict_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqEncoding_genType_Strict_fnRecorder {
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
					Result1 *base64.Encoding
				}
				Sequence   uint32
				DoFn       MoqEncoding_genType_Strict_doFn
				DoReturnFn MoqEncoding_genType_Strict_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqEncoding_genType) PrettyParams_Strict(params MoqEncoding_genType_Strict_params) string {
	return fmt.Sprintf("Strict()")
}

func (m *MoqEncoding_genType) ParamsKey_Strict(params MoqEncoding_genType_Strict_params, anyParams uint64) MoqEncoding_genType_Strict_paramsKey {
	m.Scene.T.Helper()
	return MoqEncoding_genType_Strict_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqEncoding_genType) Reset() {
	m.ResultsByParams_WithPadding = nil
	m.ResultsByParams_Strict = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqEncoding_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_WithPadding {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_WithPadding(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Strict {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Strict(results.Params))
			}
		}
	}
}

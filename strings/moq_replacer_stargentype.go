// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package strings

import (
	"fmt"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that strings.Replacer_starGenType is
// mocked completely
var _ Replacer_starGenType = (*MoqReplacer_starGenType_mock)(nil)

// Replacer_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type Replacer_starGenType interface {
	Replace(s string) string
	WriteString(w io.Writer, s string) (n int, err error)
}

// MoqReplacer_starGenType holds the state of a moq of the Replacer_starGenType
// type
type MoqReplacer_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqReplacer_starGenType_mock

	ResultsByParams_Replace     []MoqReplacer_starGenType_Replace_resultsByParams
	ResultsByParams_WriteString []MoqReplacer_starGenType_WriteString_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Replace struct {
				S moq.ParamIndexing
			}
			WriteString struct {
				W moq.ParamIndexing
				S moq.ParamIndexing
			}
		}
	}
	// MoqReplacer_starGenType_mock isolates the mock interface of the
}

// Replacer_starGenType type
type MoqReplacer_starGenType_mock struct {
	Moq *MoqReplacer_starGenType
}

// MoqReplacer_starGenType_recorder isolates the recorder interface of the
// Replacer_starGenType type
type MoqReplacer_starGenType_recorder struct {
	Moq *MoqReplacer_starGenType
}

// MoqReplacer_starGenType_Replace_params holds the params of the
// Replacer_starGenType type
type MoqReplacer_starGenType_Replace_params struct{ S string }

// MoqReplacer_starGenType_Replace_paramsKey holds the map key params of the
// Replacer_starGenType type
type MoqReplacer_starGenType_Replace_paramsKey struct {
	Params struct{ S string }
	Hashes struct{ S hash.Hash }
}

// MoqReplacer_starGenType_Replace_resultsByParams contains the results for a
// given set of parameters for the Replacer_starGenType type
type MoqReplacer_starGenType_Replace_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqReplacer_starGenType_Replace_paramsKey]*MoqReplacer_starGenType_Replace_results
}

// MoqReplacer_starGenType_Replace_doFn defines the type of function needed
// when calling AndDo for the Replacer_starGenType type
type MoqReplacer_starGenType_Replace_doFn func(s string)

// MoqReplacer_starGenType_Replace_doReturnFn defines the type of function
// needed when calling DoReturnResults for the Replacer_starGenType type
type MoqReplacer_starGenType_Replace_doReturnFn func(s string) string

// MoqReplacer_starGenType_Replace_results holds the results of the
// Replacer_starGenType type
type MoqReplacer_starGenType_Replace_results struct {
	Params  MoqReplacer_starGenType_Replace_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqReplacer_starGenType_Replace_doFn
		DoReturnFn MoqReplacer_starGenType_Replace_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqReplacer_starGenType_Replace_fnRecorder routes recorded function calls to
// the MoqReplacer_starGenType moq
type MoqReplacer_starGenType_Replace_fnRecorder struct {
	Params    MoqReplacer_starGenType_Replace_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqReplacer_starGenType_Replace_results
	Moq       *MoqReplacer_starGenType
}

// MoqReplacer_starGenType_Replace_anyParams isolates the any params functions
// of the Replacer_starGenType type
type MoqReplacer_starGenType_Replace_anyParams struct {
	Recorder *MoqReplacer_starGenType_Replace_fnRecorder
}

// MoqReplacer_starGenType_WriteString_params holds the params of the
// Replacer_starGenType type
type MoqReplacer_starGenType_WriteString_params struct {
	W io.Writer
	S string
}

// MoqReplacer_starGenType_WriteString_paramsKey holds the map key params of
// the Replacer_starGenType type
type MoqReplacer_starGenType_WriteString_paramsKey struct {
	Params struct {
		W io.Writer
		S string
	}
	Hashes struct {
		W hash.Hash
		S hash.Hash
	}
}

// MoqReplacer_starGenType_WriteString_resultsByParams contains the results for
// a given set of parameters for the Replacer_starGenType type
type MoqReplacer_starGenType_WriteString_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqReplacer_starGenType_WriteString_paramsKey]*MoqReplacer_starGenType_WriteString_results
}

// MoqReplacer_starGenType_WriteString_doFn defines the type of function needed
// when calling AndDo for the Replacer_starGenType type
type MoqReplacer_starGenType_WriteString_doFn func(w io.Writer, s string)

// MoqReplacer_starGenType_WriteString_doReturnFn defines the type of function
// needed when calling DoReturnResults for the Replacer_starGenType type
type MoqReplacer_starGenType_WriteString_doReturnFn func(w io.Writer, s string) (n int, err error)

// MoqReplacer_starGenType_WriteString_results holds the results of the
// Replacer_starGenType type
type MoqReplacer_starGenType_WriteString_results struct {
	Params  MoqReplacer_starGenType_WriteString_params
	Results []struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqReplacer_starGenType_WriteString_doFn
		DoReturnFn MoqReplacer_starGenType_WriteString_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqReplacer_starGenType_WriteString_fnRecorder routes recorded function
// calls to the MoqReplacer_starGenType moq
type MoqReplacer_starGenType_WriteString_fnRecorder struct {
	Params    MoqReplacer_starGenType_WriteString_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqReplacer_starGenType_WriteString_results
	Moq       *MoqReplacer_starGenType
}

// MoqReplacer_starGenType_WriteString_anyParams isolates the any params
// functions of the Replacer_starGenType type
type MoqReplacer_starGenType_WriteString_anyParams struct {
	Recorder *MoqReplacer_starGenType_WriteString_fnRecorder
}

// NewMoqReplacer_starGenType creates a new moq of the Replacer_starGenType
// type
func NewMoqReplacer_starGenType(scene *moq.Scene, config *moq.Config) *MoqReplacer_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqReplacer_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqReplacer_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Replace struct {
					S moq.ParamIndexing
				}
				WriteString struct {
					W moq.ParamIndexing
					S moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Replace struct {
				S moq.ParamIndexing
			}
			WriteString struct {
				W moq.ParamIndexing
				S moq.ParamIndexing
			}
		}{
			Replace: struct {
				S moq.ParamIndexing
			}{
				S: moq.ParamIndexByValue,
			},
			WriteString: struct {
				W moq.ParamIndexing
				S moq.ParamIndexing
			}{
				W: moq.ParamIndexByHash,
				S: moq.ParamIndexByValue,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Replacer_starGenType type
func (m *MoqReplacer_starGenType) Mock() *MoqReplacer_starGenType_mock { return m.Moq }

func (m *MoqReplacer_starGenType_mock) Replace(s string) (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqReplacer_starGenType_Replace_params{
		S: s,
	}
	var results *MoqReplacer_starGenType_Replace_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Replace {
		paramsKey := m.Moq.ParamsKey_Replace(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Replace(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Replace(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Replace(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(s)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(s)
	}
	return
}

func (m *MoqReplacer_starGenType_mock) WriteString(w io.Writer, s string) (n int, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqReplacer_starGenType_WriteString_params{
		W: w,
		S: s,
	}
	var results *MoqReplacer_starGenType_WriteString_results
	for _, resultsByParams := range m.Moq.ResultsByParams_WriteString {
		paramsKey := m.Moq.ParamsKey_WriteString(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_WriteString(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_WriteString(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_WriteString(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(w, s)
	}

	if result.Values != nil {
		n = result.Values.N
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		n, err = result.DoReturnFn(w, s)
	}
	return
}

// OnCall returns the recorder implementation of the Replacer_starGenType type
func (m *MoqReplacer_starGenType) OnCall() *MoqReplacer_starGenType_recorder {
	return &MoqReplacer_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqReplacer_starGenType_recorder) Replace(s string) *MoqReplacer_starGenType_Replace_fnRecorder {
	return &MoqReplacer_starGenType_Replace_fnRecorder{
		Params: MoqReplacer_starGenType_Replace_params{
			S: s,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqReplacer_starGenType_Replace_fnRecorder) Any() *MoqReplacer_starGenType_Replace_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Replace(r.Params))
		return nil
	}
	return &MoqReplacer_starGenType_Replace_anyParams{Recorder: r}
}

func (a *MoqReplacer_starGenType_Replace_anyParams) S() *MoqReplacer_starGenType_Replace_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqReplacer_starGenType_Replace_fnRecorder) Seq() *MoqReplacer_starGenType_Replace_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Replace(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqReplacer_starGenType_Replace_fnRecorder) NoSeq() *MoqReplacer_starGenType_Replace_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Replace(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqReplacer_starGenType_Replace_fnRecorder) ReturnResults(result1 string) *MoqReplacer_starGenType_Replace_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqReplacer_starGenType_Replace_doFn
		DoReturnFn MoqReplacer_starGenType_Replace_doReturnFn
	}{
		Values: &struct {
			Result1 string
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqReplacer_starGenType_Replace_fnRecorder) AndDo(fn MoqReplacer_starGenType_Replace_doFn) *MoqReplacer_starGenType_Replace_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqReplacer_starGenType_Replace_fnRecorder) DoReturnResults(fn MoqReplacer_starGenType_Replace_doReturnFn) *MoqReplacer_starGenType_Replace_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqReplacer_starGenType_Replace_doFn
		DoReturnFn MoqReplacer_starGenType_Replace_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqReplacer_starGenType_Replace_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqReplacer_starGenType_Replace_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Replace {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqReplacer_starGenType_Replace_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqReplacer_starGenType_Replace_paramsKey]*MoqReplacer_starGenType_Replace_results{},
		}
		r.Moq.ResultsByParams_Replace = append(r.Moq.ResultsByParams_Replace, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Replace) {
			copy(r.Moq.ResultsByParams_Replace[insertAt+1:], r.Moq.ResultsByParams_Replace[insertAt:0])
			r.Moq.ResultsByParams_Replace[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Replace(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqReplacer_starGenType_Replace_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqReplacer_starGenType_Replace_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqReplacer_starGenType_Replace_fnRecorder {
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
					Result1 string
				}
				Sequence   uint32
				DoFn       MoqReplacer_starGenType_Replace_doFn
				DoReturnFn MoqReplacer_starGenType_Replace_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqReplacer_starGenType) PrettyParams_Replace(params MoqReplacer_starGenType_Replace_params) string {
	return fmt.Sprintf("Replace(%#v)", params.S)
}

func (m *MoqReplacer_starGenType) ParamsKey_Replace(params MoqReplacer_starGenType_Replace_params, anyParams uint64) MoqReplacer_starGenType_Replace_paramsKey {
	m.Scene.T.Helper()
	var sUsed string
	var sUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Replace.S == moq.ParamIndexByValue {
			sUsed = params.S
		} else {
			sUsedHash = hash.DeepHash(params.S)
		}
	}
	return MoqReplacer_starGenType_Replace_paramsKey{
		Params: struct{ S string }{
			S: sUsed,
		},
		Hashes: struct{ S hash.Hash }{
			S: sUsedHash,
		},
	}
}

func (m *MoqReplacer_starGenType_recorder) WriteString(w io.Writer, s string) *MoqReplacer_starGenType_WriteString_fnRecorder {
	return &MoqReplacer_starGenType_WriteString_fnRecorder{
		Params: MoqReplacer_starGenType_WriteString_params{
			W: w,
			S: s,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqReplacer_starGenType_WriteString_fnRecorder) Any() *MoqReplacer_starGenType_WriteString_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_WriteString(r.Params))
		return nil
	}
	return &MoqReplacer_starGenType_WriteString_anyParams{Recorder: r}
}

func (a *MoqReplacer_starGenType_WriteString_anyParams) W() *MoqReplacer_starGenType_WriteString_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqReplacer_starGenType_WriteString_anyParams) S() *MoqReplacer_starGenType_WriteString_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqReplacer_starGenType_WriteString_fnRecorder) Seq() *MoqReplacer_starGenType_WriteString_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_WriteString(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqReplacer_starGenType_WriteString_fnRecorder) NoSeq() *MoqReplacer_starGenType_WriteString_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_WriteString(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqReplacer_starGenType_WriteString_fnRecorder) ReturnResults(n int, err error) *MoqReplacer_starGenType_WriteString_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqReplacer_starGenType_WriteString_doFn
		DoReturnFn MoqReplacer_starGenType_WriteString_doReturnFn
	}{
		Values: &struct {
			N   int
			Err error
		}{
			N:   n,
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqReplacer_starGenType_WriteString_fnRecorder) AndDo(fn MoqReplacer_starGenType_WriteString_doFn) *MoqReplacer_starGenType_WriteString_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqReplacer_starGenType_WriteString_fnRecorder) DoReturnResults(fn MoqReplacer_starGenType_WriteString_doReturnFn) *MoqReplacer_starGenType_WriteString_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqReplacer_starGenType_WriteString_doFn
		DoReturnFn MoqReplacer_starGenType_WriteString_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqReplacer_starGenType_WriteString_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqReplacer_starGenType_WriteString_resultsByParams
	for n, res := range r.Moq.ResultsByParams_WriteString {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqReplacer_starGenType_WriteString_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqReplacer_starGenType_WriteString_paramsKey]*MoqReplacer_starGenType_WriteString_results{},
		}
		r.Moq.ResultsByParams_WriteString = append(r.Moq.ResultsByParams_WriteString, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_WriteString) {
			copy(r.Moq.ResultsByParams_WriteString[insertAt+1:], r.Moq.ResultsByParams_WriteString[insertAt:0])
			r.Moq.ResultsByParams_WriteString[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_WriteString(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqReplacer_starGenType_WriteString_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqReplacer_starGenType_WriteString_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqReplacer_starGenType_WriteString_fnRecorder {
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
					N   int
					Err error
				}
				Sequence   uint32
				DoFn       MoqReplacer_starGenType_WriteString_doFn
				DoReturnFn MoqReplacer_starGenType_WriteString_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqReplacer_starGenType) PrettyParams_WriteString(params MoqReplacer_starGenType_WriteString_params) string {
	return fmt.Sprintf("WriteString(%#v, %#v)", params.W, params.S)
}

func (m *MoqReplacer_starGenType) ParamsKey_WriteString(params MoqReplacer_starGenType_WriteString_params, anyParams uint64) MoqReplacer_starGenType_WriteString_paramsKey {
	m.Scene.T.Helper()
	var wUsed io.Writer
	var wUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.WriteString.W == moq.ParamIndexByValue {
			wUsed = params.W
		} else {
			wUsedHash = hash.DeepHash(params.W)
		}
	}
	var sUsed string
	var sUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.WriteString.S == moq.ParamIndexByValue {
			sUsed = params.S
		} else {
			sUsedHash = hash.DeepHash(params.S)
		}
	}
	return MoqReplacer_starGenType_WriteString_paramsKey{
		Params: struct {
			W io.Writer
			S string
		}{
			W: wUsed,
			S: sUsed,
		},
		Hashes: struct {
			W hash.Hash
			S hash.Hash
		}{
			W: wUsedHash,
			S: sUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqReplacer_starGenType) Reset() {
	m.ResultsByParams_Replace = nil
	m.ResultsByParams_WriteString = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqReplacer_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Replace {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Replace(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_WriteString {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_WriteString(results.Params))
			}
		}
	}
}
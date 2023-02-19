// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package textproto

import (
	"fmt"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that textproto.Writer_starGenType is
// mocked completely
var _ Writer_starGenType = (*MoqWriter_starGenType_mock)(nil)

// Writer_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type Writer_starGenType interface {
	PrintfLine(format string, args ...interface{}) error
	DotWriter() io.WriteCloser
}

// MoqWriter_starGenType holds the state of a moq of the Writer_starGenType
// type
type MoqWriter_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqWriter_starGenType_mock

	ResultsByParams_PrintfLine []MoqWriter_starGenType_PrintfLine_resultsByParams
	ResultsByParams_DotWriter  []MoqWriter_starGenType_DotWriter_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			PrintfLine struct {
				Format moq.ParamIndexing
				Args   moq.ParamIndexing
			}
			DotWriter struct{}
		}
	}
}

// MoqWriter_starGenType_mock isolates the mock interface of the
// Writer_starGenType type
type MoqWriter_starGenType_mock struct {
	Moq *MoqWriter_starGenType
}

// MoqWriter_starGenType_recorder isolates the recorder interface of the
// Writer_starGenType type
type MoqWriter_starGenType_recorder struct {
	Moq *MoqWriter_starGenType
}

// MoqWriter_starGenType_PrintfLine_params holds the params of the
// Writer_starGenType type
type MoqWriter_starGenType_PrintfLine_params struct {
	Format string
	Args   []interface{}
}

// MoqWriter_starGenType_PrintfLine_paramsKey holds the map key params of the
// Writer_starGenType type
type MoqWriter_starGenType_PrintfLine_paramsKey struct {
	Params struct{ Format string }
	Hashes struct {
		Format hash.Hash
		Args   hash.Hash
	}
}

// MoqWriter_starGenType_PrintfLine_resultsByParams contains the results for a
// given set of parameters for the Writer_starGenType type
type MoqWriter_starGenType_PrintfLine_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqWriter_starGenType_PrintfLine_paramsKey]*MoqWriter_starGenType_PrintfLine_results
}

// MoqWriter_starGenType_PrintfLine_doFn defines the type of function needed
// when calling AndDo for the Writer_starGenType type
type MoqWriter_starGenType_PrintfLine_doFn func(format string, args ...interface{})

// MoqWriter_starGenType_PrintfLine_doReturnFn defines the type of function
// needed when calling DoReturnResults for the Writer_starGenType type
type MoqWriter_starGenType_PrintfLine_doReturnFn func(format string, args ...interface{}) error

// MoqWriter_starGenType_PrintfLine_results holds the results of the
// Writer_starGenType type
type MoqWriter_starGenType_PrintfLine_results struct {
	Params  MoqWriter_starGenType_PrintfLine_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqWriter_starGenType_PrintfLine_doFn
		DoReturnFn MoqWriter_starGenType_PrintfLine_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqWriter_starGenType_PrintfLine_fnRecorder routes recorded function calls
// to the MoqWriter_starGenType moq
type MoqWriter_starGenType_PrintfLine_fnRecorder struct {
	Params    MoqWriter_starGenType_PrintfLine_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqWriter_starGenType_PrintfLine_results
	Moq       *MoqWriter_starGenType
}

// MoqWriter_starGenType_PrintfLine_anyParams isolates the any params functions
// of the Writer_starGenType type
type MoqWriter_starGenType_PrintfLine_anyParams struct {
	Recorder *MoqWriter_starGenType_PrintfLine_fnRecorder
}

// MoqWriter_starGenType_DotWriter_params holds the params of the
// Writer_starGenType type
type MoqWriter_starGenType_DotWriter_params struct{}

// MoqWriter_starGenType_DotWriter_paramsKey holds the map key params of the
// Writer_starGenType type
type MoqWriter_starGenType_DotWriter_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqWriter_starGenType_DotWriter_resultsByParams contains the results for a
// given set of parameters for the Writer_starGenType type
type MoqWriter_starGenType_DotWriter_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqWriter_starGenType_DotWriter_paramsKey]*MoqWriter_starGenType_DotWriter_results
}

// MoqWriter_starGenType_DotWriter_doFn defines the type of function needed
// when calling AndDo for the Writer_starGenType type
type MoqWriter_starGenType_DotWriter_doFn func()

// MoqWriter_starGenType_DotWriter_doReturnFn defines the type of function
// needed when calling DoReturnResults for the Writer_starGenType type
type MoqWriter_starGenType_DotWriter_doReturnFn func() io.WriteCloser

// MoqWriter_starGenType_DotWriter_results holds the results of the
// Writer_starGenType type
type MoqWriter_starGenType_DotWriter_results struct {
	Params  MoqWriter_starGenType_DotWriter_params
	Results []struct {
		Values *struct {
			Result1 io.WriteCloser
		}
		Sequence   uint32
		DoFn       MoqWriter_starGenType_DotWriter_doFn
		DoReturnFn MoqWriter_starGenType_DotWriter_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqWriter_starGenType_DotWriter_fnRecorder routes recorded function calls to
// the MoqWriter_starGenType moq
type MoqWriter_starGenType_DotWriter_fnRecorder struct {
	Params    MoqWriter_starGenType_DotWriter_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqWriter_starGenType_DotWriter_results
	Moq       *MoqWriter_starGenType
}

// MoqWriter_starGenType_DotWriter_anyParams isolates the any params functions
// of the Writer_starGenType type
type MoqWriter_starGenType_DotWriter_anyParams struct {
	Recorder *MoqWriter_starGenType_DotWriter_fnRecorder
}

// NewMoqWriter_starGenType creates a new moq of the Writer_starGenType type
func NewMoqWriter_starGenType(scene *moq.Scene, config *moq.Config) *MoqWriter_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqWriter_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqWriter_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				PrintfLine struct {
					Format moq.ParamIndexing
					Args   moq.ParamIndexing
				}
				DotWriter struct{}
			}
		}{ParameterIndexing: struct {
			PrintfLine struct {
				Format moq.ParamIndexing
				Args   moq.ParamIndexing
			}
			DotWriter struct{}
		}{
			PrintfLine: struct {
				Format moq.ParamIndexing
				Args   moq.ParamIndexing
			}{
				Format: moq.ParamIndexByValue,
				Args:   moq.ParamIndexByHash,
			},
			DotWriter: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Writer_starGenType type
func (m *MoqWriter_starGenType) Mock() *MoqWriter_starGenType_mock { return m.Moq }

func (m *MoqWriter_starGenType_mock) PrintfLine(format string, args ...interface{}) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqWriter_starGenType_PrintfLine_params{
		Format: format,
		Args:   args,
	}
	var results *MoqWriter_starGenType_PrintfLine_results
	for _, resultsByParams := range m.Moq.ResultsByParams_PrintfLine {
		paramsKey := m.Moq.ParamsKey_PrintfLine(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_PrintfLine(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_PrintfLine(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_PrintfLine(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(format, args...)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(format, args...)
	}
	return
}

func (m *MoqWriter_starGenType_mock) DotWriter() (result1 io.WriteCloser) {
	m.Moq.Scene.T.Helper()
	params := MoqWriter_starGenType_DotWriter_params{}
	var results *MoqWriter_starGenType_DotWriter_results
	for _, resultsByParams := range m.Moq.ResultsByParams_DotWriter {
		paramsKey := m.Moq.ParamsKey_DotWriter(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_DotWriter(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_DotWriter(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_DotWriter(params))
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

// OnCall returns the recorder implementation of the Writer_starGenType type
func (m *MoqWriter_starGenType) OnCall() *MoqWriter_starGenType_recorder {
	return &MoqWriter_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqWriter_starGenType_recorder) PrintfLine(format string, args ...interface{}) *MoqWriter_starGenType_PrintfLine_fnRecorder {
	return &MoqWriter_starGenType_PrintfLine_fnRecorder{
		Params: MoqWriter_starGenType_PrintfLine_params{
			Format: format,
			Args:   args,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqWriter_starGenType_PrintfLine_fnRecorder) Any() *MoqWriter_starGenType_PrintfLine_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_PrintfLine(r.Params))
		return nil
	}
	return &MoqWriter_starGenType_PrintfLine_anyParams{Recorder: r}
}

func (a *MoqWriter_starGenType_PrintfLine_anyParams) Format() *MoqWriter_starGenType_PrintfLine_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqWriter_starGenType_PrintfLine_anyParams) Args() *MoqWriter_starGenType_PrintfLine_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqWriter_starGenType_PrintfLine_fnRecorder) Seq() *MoqWriter_starGenType_PrintfLine_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_PrintfLine(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqWriter_starGenType_PrintfLine_fnRecorder) NoSeq() *MoqWriter_starGenType_PrintfLine_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_PrintfLine(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqWriter_starGenType_PrintfLine_fnRecorder) ReturnResults(result1 error) *MoqWriter_starGenType_PrintfLine_fnRecorder {
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
		DoFn       MoqWriter_starGenType_PrintfLine_doFn
		DoReturnFn MoqWriter_starGenType_PrintfLine_doReturnFn
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

func (r *MoqWriter_starGenType_PrintfLine_fnRecorder) AndDo(fn MoqWriter_starGenType_PrintfLine_doFn) *MoqWriter_starGenType_PrintfLine_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqWriter_starGenType_PrintfLine_fnRecorder) DoReturnResults(fn MoqWriter_starGenType_PrintfLine_doReturnFn) *MoqWriter_starGenType_PrintfLine_fnRecorder {
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
		DoFn       MoqWriter_starGenType_PrintfLine_doFn
		DoReturnFn MoqWriter_starGenType_PrintfLine_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqWriter_starGenType_PrintfLine_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqWriter_starGenType_PrintfLine_resultsByParams
	for n, res := range r.Moq.ResultsByParams_PrintfLine {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqWriter_starGenType_PrintfLine_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqWriter_starGenType_PrintfLine_paramsKey]*MoqWriter_starGenType_PrintfLine_results{},
		}
		r.Moq.ResultsByParams_PrintfLine = append(r.Moq.ResultsByParams_PrintfLine, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_PrintfLine) {
			copy(r.Moq.ResultsByParams_PrintfLine[insertAt+1:], r.Moq.ResultsByParams_PrintfLine[insertAt:0])
			r.Moq.ResultsByParams_PrintfLine[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_PrintfLine(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqWriter_starGenType_PrintfLine_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqWriter_starGenType_PrintfLine_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqWriter_starGenType_PrintfLine_fnRecorder {
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
				DoFn       MoqWriter_starGenType_PrintfLine_doFn
				DoReturnFn MoqWriter_starGenType_PrintfLine_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqWriter_starGenType) PrettyParams_PrintfLine(params MoqWriter_starGenType_PrintfLine_params) string {
	return fmt.Sprintf("PrintfLine(%#v, %#v)", params.Format, params.Args)
}

func (m *MoqWriter_starGenType) ParamsKey_PrintfLine(params MoqWriter_starGenType_PrintfLine_params, anyParams uint64) MoqWriter_starGenType_PrintfLine_paramsKey {
	m.Scene.T.Helper()
	var formatUsed string
	var formatUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.PrintfLine.Format == moq.ParamIndexByValue {
			formatUsed = params.Format
		} else {
			formatUsedHash = hash.DeepHash(params.Format)
		}
	}
	var argsUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.PrintfLine.Args == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The args parameter of the PrintfLine function can't be indexed by value")
		}
		argsUsedHash = hash.DeepHash(params.Args)
	}
	return MoqWriter_starGenType_PrintfLine_paramsKey{
		Params: struct{ Format string }{
			Format: formatUsed,
		},
		Hashes: struct {
			Format hash.Hash
			Args   hash.Hash
		}{
			Format: formatUsedHash,
			Args:   argsUsedHash,
		},
	}
}

func (m *MoqWriter_starGenType_recorder) DotWriter() *MoqWriter_starGenType_DotWriter_fnRecorder {
	return &MoqWriter_starGenType_DotWriter_fnRecorder{
		Params:   MoqWriter_starGenType_DotWriter_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqWriter_starGenType_DotWriter_fnRecorder) Any() *MoqWriter_starGenType_DotWriter_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_DotWriter(r.Params))
		return nil
	}
	return &MoqWriter_starGenType_DotWriter_anyParams{Recorder: r}
}

func (r *MoqWriter_starGenType_DotWriter_fnRecorder) Seq() *MoqWriter_starGenType_DotWriter_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_DotWriter(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqWriter_starGenType_DotWriter_fnRecorder) NoSeq() *MoqWriter_starGenType_DotWriter_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_DotWriter(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqWriter_starGenType_DotWriter_fnRecorder) ReturnResults(result1 io.WriteCloser) *MoqWriter_starGenType_DotWriter_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 io.WriteCloser
		}
		Sequence   uint32
		DoFn       MoqWriter_starGenType_DotWriter_doFn
		DoReturnFn MoqWriter_starGenType_DotWriter_doReturnFn
	}{
		Values: &struct {
			Result1 io.WriteCloser
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqWriter_starGenType_DotWriter_fnRecorder) AndDo(fn MoqWriter_starGenType_DotWriter_doFn) *MoqWriter_starGenType_DotWriter_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqWriter_starGenType_DotWriter_fnRecorder) DoReturnResults(fn MoqWriter_starGenType_DotWriter_doReturnFn) *MoqWriter_starGenType_DotWriter_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 io.WriteCloser
		}
		Sequence   uint32
		DoFn       MoqWriter_starGenType_DotWriter_doFn
		DoReturnFn MoqWriter_starGenType_DotWriter_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqWriter_starGenType_DotWriter_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqWriter_starGenType_DotWriter_resultsByParams
	for n, res := range r.Moq.ResultsByParams_DotWriter {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqWriter_starGenType_DotWriter_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqWriter_starGenType_DotWriter_paramsKey]*MoqWriter_starGenType_DotWriter_results{},
		}
		r.Moq.ResultsByParams_DotWriter = append(r.Moq.ResultsByParams_DotWriter, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_DotWriter) {
			copy(r.Moq.ResultsByParams_DotWriter[insertAt+1:], r.Moq.ResultsByParams_DotWriter[insertAt:0])
			r.Moq.ResultsByParams_DotWriter[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_DotWriter(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqWriter_starGenType_DotWriter_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqWriter_starGenType_DotWriter_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqWriter_starGenType_DotWriter_fnRecorder {
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
					Result1 io.WriteCloser
				}
				Sequence   uint32
				DoFn       MoqWriter_starGenType_DotWriter_doFn
				DoReturnFn MoqWriter_starGenType_DotWriter_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqWriter_starGenType) PrettyParams_DotWriter(params MoqWriter_starGenType_DotWriter_params) string {
	return fmt.Sprintf("DotWriter()")
}

func (m *MoqWriter_starGenType) ParamsKey_DotWriter(params MoqWriter_starGenType_DotWriter_params, anyParams uint64) MoqWriter_starGenType_DotWriter_paramsKey {
	m.Scene.T.Helper()
	return MoqWriter_starGenType_DotWriter_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqWriter_starGenType) Reset() {
	m.ResultsByParams_PrintfLine = nil
	m.ResultsByParams_DotWriter = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqWriter_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_PrintfLine {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_PrintfLine(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_DotWriter {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_DotWriter(results.Params))
			}
		}
	}
}
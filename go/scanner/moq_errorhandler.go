// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package scanner

import (
	"fmt"
	"go/scanner"
	"go/token"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// MoqErrorHandler holds the state of a moq of the ErrorHandler type
type MoqErrorHandler struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqErrorHandler_mock

	ResultsByParams []MoqErrorHandler_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Pos moq.ParamIndexing
			Msg moq.ParamIndexing
		}
	}
}

// MoqErrorHandler_mock isolates the mock interface of the ErrorHandler type
type MoqErrorHandler_mock struct {
	Moq *MoqErrorHandler
}

// MoqErrorHandler_params holds the params of the ErrorHandler type
type MoqErrorHandler_params struct {
	Pos token.Position
	Msg string
}

// MoqErrorHandler_paramsKey holds the map key params of the ErrorHandler type
type MoqErrorHandler_paramsKey struct {
	Params struct {
		Pos token.Position
		Msg string
	}
	Hashes struct {
		Pos hash.Hash
		Msg hash.Hash
	}
}

// MoqErrorHandler_resultsByParams contains the results for a given set of
// parameters for the ErrorHandler type
type MoqErrorHandler_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqErrorHandler_paramsKey]*MoqErrorHandler_results
}

// MoqErrorHandler_doFn defines the type of function needed when calling AndDo
// for the ErrorHandler type
type MoqErrorHandler_doFn func(pos token.Position, msg string)

// MoqErrorHandler_doReturnFn defines the type of function needed when calling
// DoReturnResults for the ErrorHandler type
type MoqErrorHandler_doReturnFn func(pos token.Position, msg string)

// MoqErrorHandler_results holds the results of the ErrorHandler type
type MoqErrorHandler_results struct {
	Params  MoqErrorHandler_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqErrorHandler_doFn
		DoReturnFn MoqErrorHandler_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqErrorHandler_fnRecorder routes recorded function calls to the
// MoqErrorHandler moq
type MoqErrorHandler_fnRecorder struct {
	Params    MoqErrorHandler_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqErrorHandler_results
	Moq       *MoqErrorHandler
}

// MoqErrorHandler_anyParams isolates the any params functions of the
// ErrorHandler type
type MoqErrorHandler_anyParams struct {
	Recorder *MoqErrorHandler_fnRecorder
}

// NewMoqErrorHandler creates a new moq of the ErrorHandler type
func NewMoqErrorHandler(scene *moq.Scene, config *moq.Config) *MoqErrorHandler {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqErrorHandler{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqErrorHandler_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Pos moq.ParamIndexing
				Msg moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Pos moq.ParamIndexing
			Msg moq.ParamIndexing
		}{
			Pos: moq.ParamIndexByValue,
			Msg: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ErrorHandler type
func (m *MoqErrorHandler) Mock() scanner.ErrorHandler {
	return func(pos token.Position, msg string) {
		m.Scene.T.Helper()
		moq := &MoqErrorHandler_mock{Moq: m}
		moq.Fn(pos, msg)
	}
}

func (m *MoqErrorHandler_mock) Fn(pos token.Position, msg string) {
	m.Moq.Scene.T.Helper()
	params := MoqErrorHandler_params{
		Pos: pos,
		Msg: msg,
	}
	var results *MoqErrorHandler_results
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
		result.DoFn(pos, msg)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(pos, msg)
	}
	return
}

func (m *MoqErrorHandler) OnCall(pos token.Position, msg string) *MoqErrorHandler_fnRecorder {
	return &MoqErrorHandler_fnRecorder{
		Params: MoqErrorHandler_params{
			Pos: pos,
			Msg: msg,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqErrorHandler_fnRecorder) Any() *MoqErrorHandler_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqErrorHandler_anyParams{Recorder: r}
}

func (a *MoqErrorHandler_anyParams) Pos() *MoqErrorHandler_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqErrorHandler_anyParams) Msg() *MoqErrorHandler_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqErrorHandler_fnRecorder) Seq() *MoqErrorHandler_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqErrorHandler_fnRecorder) NoSeq() *MoqErrorHandler_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqErrorHandler_fnRecorder) ReturnResults() *MoqErrorHandler_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqErrorHandler_doFn
		DoReturnFn MoqErrorHandler_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqErrorHandler_fnRecorder) AndDo(fn MoqErrorHandler_doFn) *MoqErrorHandler_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqErrorHandler_fnRecorder) DoReturnResults(fn MoqErrorHandler_doReturnFn) *MoqErrorHandler_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqErrorHandler_doFn
		DoReturnFn MoqErrorHandler_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqErrorHandler_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqErrorHandler_resultsByParams
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
		results = &MoqErrorHandler_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqErrorHandler_paramsKey]*MoqErrorHandler_results{},
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
		r.Results = &MoqErrorHandler_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqErrorHandler_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqErrorHandler_fnRecorder {
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
				DoFn       MoqErrorHandler_doFn
				DoReturnFn MoqErrorHandler_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqErrorHandler) PrettyParams(params MoqErrorHandler_params) string {
	return fmt.Sprintf("ErrorHandler(%#v, %#v)", params.Pos, params.Msg)
}

func (m *MoqErrorHandler) ParamsKey(params MoqErrorHandler_params, anyParams uint64) MoqErrorHandler_paramsKey {
	m.Scene.T.Helper()
	var posUsed token.Position
	var posUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Pos == moq.ParamIndexByValue {
			posUsed = params.Pos
		} else {
			posUsedHash = hash.DeepHash(params.Pos)
		}
	}
	var msgUsed string
	var msgUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Msg == moq.ParamIndexByValue {
			msgUsed = params.Msg
		} else {
			msgUsedHash = hash.DeepHash(params.Msg)
		}
	}
	return MoqErrorHandler_paramsKey{
		Params: struct {
			Pos token.Position
			Msg string
		}{
			Pos: posUsed,
			Msg: msgUsed,
		},
		Hashes: struct {
			Pos hash.Hash
			Msg hash.Hash
		}{
			Pos: posUsedHash,
			Msg: msgUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqErrorHandler) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqErrorHandler) AssertExpectationsMet() {
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

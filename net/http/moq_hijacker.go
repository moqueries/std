// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package http

import (
	"bufio"
	"fmt"
	"math/bits"
	"net"
	"net/http"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that http.Hijacker is mocked completely
var _ http.Hijacker = (*MoqHijacker_mock)(nil)

// MoqHijacker holds the state of a moq of the Hijacker type
type MoqHijacker struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqHijacker_mock

	ResultsByParams_Hijack []MoqHijacker_Hijack_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Hijack struct{}
		}
	}
}

// MoqHijacker_mock isolates the mock interface of the Hijacker type
type MoqHijacker_mock struct {
	Moq *MoqHijacker
}

// MoqHijacker_recorder isolates the recorder interface of the Hijacker type
type MoqHijacker_recorder struct {
	Moq *MoqHijacker
}

// MoqHijacker_Hijack_params holds the params of the Hijacker type
type MoqHijacker_Hijack_params struct{}

// MoqHijacker_Hijack_paramsKey holds the map key params of the Hijacker type
type MoqHijacker_Hijack_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqHijacker_Hijack_resultsByParams contains the results for a given set of
// parameters for the Hijacker type
type MoqHijacker_Hijack_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqHijacker_Hijack_paramsKey]*MoqHijacker_Hijack_results
}

// MoqHijacker_Hijack_doFn defines the type of function needed when calling
// AndDo for the Hijacker type
type MoqHijacker_Hijack_doFn func()

// MoqHijacker_Hijack_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Hijacker type
type MoqHijacker_Hijack_doReturnFn func() (net.Conn, *bufio.ReadWriter, error)

// MoqHijacker_Hijack_results holds the results of the Hijacker type
type MoqHijacker_Hijack_results struct {
	Params  MoqHijacker_Hijack_params
	Results []struct {
		Values *struct {
			Result1 net.Conn
			Result2 *bufio.ReadWriter
			Result3 error
		}
		Sequence   uint32
		DoFn       MoqHijacker_Hijack_doFn
		DoReturnFn MoqHijacker_Hijack_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqHijacker_Hijack_fnRecorder routes recorded function calls to the
// MoqHijacker moq
type MoqHijacker_Hijack_fnRecorder struct {
	Params    MoqHijacker_Hijack_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqHijacker_Hijack_results
	Moq       *MoqHijacker
}

// MoqHijacker_Hijack_anyParams isolates the any params functions of the
// Hijacker type
type MoqHijacker_Hijack_anyParams struct {
	Recorder *MoqHijacker_Hijack_fnRecorder
}

// NewMoqHijacker creates a new moq of the Hijacker type
func NewMoqHijacker(scene *moq.Scene, config *moq.Config) *MoqHijacker {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqHijacker{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqHijacker_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Hijack struct{}
			}
		}{ParameterIndexing: struct {
			Hijack struct{}
		}{
			Hijack: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Hijacker type
func (m *MoqHijacker) Mock() *MoqHijacker_mock { return m.Moq }

func (m *MoqHijacker_mock) Hijack() (result1 net.Conn, result2 *bufio.ReadWriter, result3 error) {
	m.Moq.Scene.T.Helper()
	params := MoqHijacker_Hijack_params{}
	var results *MoqHijacker_Hijack_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Hijack {
		paramsKey := m.Moq.ParamsKey_Hijack(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Hijack(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Hijack(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Hijack(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
		result3 = result.Values.Result3
	}
	if result.DoReturnFn != nil {
		result1, result2, result3 = result.DoReturnFn()
	}
	return
}

// OnCall returns the recorder implementation of the Hijacker type
func (m *MoqHijacker) OnCall() *MoqHijacker_recorder {
	return &MoqHijacker_recorder{
		Moq: m,
	}
}

func (m *MoqHijacker_recorder) Hijack() *MoqHijacker_Hijack_fnRecorder {
	return &MoqHijacker_Hijack_fnRecorder{
		Params:   MoqHijacker_Hijack_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqHijacker_Hijack_fnRecorder) Any() *MoqHijacker_Hijack_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Hijack(r.Params))
		return nil
	}
	return &MoqHijacker_Hijack_anyParams{Recorder: r}
}

func (r *MoqHijacker_Hijack_fnRecorder) Seq() *MoqHijacker_Hijack_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Hijack(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqHijacker_Hijack_fnRecorder) NoSeq() *MoqHijacker_Hijack_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Hijack(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqHijacker_Hijack_fnRecorder) ReturnResults(result1 net.Conn, result2 *bufio.ReadWriter, result3 error) *MoqHijacker_Hijack_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 net.Conn
			Result2 *bufio.ReadWriter
			Result3 error
		}
		Sequence   uint32
		DoFn       MoqHijacker_Hijack_doFn
		DoReturnFn MoqHijacker_Hijack_doReturnFn
	}{
		Values: &struct {
			Result1 net.Conn
			Result2 *bufio.ReadWriter
			Result3 error
		}{
			Result1: result1,
			Result2: result2,
			Result3: result3,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqHijacker_Hijack_fnRecorder) AndDo(fn MoqHijacker_Hijack_doFn) *MoqHijacker_Hijack_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqHijacker_Hijack_fnRecorder) DoReturnResults(fn MoqHijacker_Hijack_doReturnFn) *MoqHijacker_Hijack_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 net.Conn
			Result2 *bufio.ReadWriter
			Result3 error
		}
		Sequence   uint32
		DoFn       MoqHijacker_Hijack_doFn
		DoReturnFn MoqHijacker_Hijack_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqHijacker_Hijack_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqHijacker_Hijack_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Hijack {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqHijacker_Hijack_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqHijacker_Hijack_paramsKey]*MoqHijacker_Hijack_results{},
		}
		r.Moq.ResultsByParams_Hijack = append(r.Moq.ResultsByParams_Hijack, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Hijack) {
			copy(r.Moq.ResultsByParams_Hijack[insertAt+1:], r.Moq.ResultsByParams_Hijack[insertAt:0])
			r.Moq.ResultsByParams_Hijack[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Hijack(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqHijacker_Hijack_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqHijacker_Hijack_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqHijacker_Hijack_fnRecorder {
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
					Result1 net.Conn
					Result2 *bufio.ReadWriter
					Result3 error
				}
				Sequence   uint32
				DoFn       MoqHijacker_Hijack_doFn
				DoReturnFn MoqHijacker_Hijack_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqHijacker) PrettyParams_Hijack(params MoqHijacker_Hijack_params) string {
	return fmt.Sprintf("Hijack()")
}

func (m *MoqHijacker) ParamsKey_Hijack(params MoqHijacker_Hijack_params, anyParams uint64) MoqHijacker_Hijack_paramsKey {
	m.Scene.T.Helper()
	return MoqHijacker_Hijack_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqHijacker) Reset() { m.ResultsByParams_Hijack = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqHijacker) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Hijack {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Hijack(results.Params))
			}
		}
	}
}
// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package fmt

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that fmt.Scanner is mocked completely
var _ fmt.Scanner = (*MoqScanner_mock)(nil)

// MoqScanner holds the state of a moq of the Scanner type
type MoqScanner struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqScanner_mock

	ResultsByParams_Scan []MoqScanner_Scan_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Scan struct {
				State moq.ParamIndexing
				Verb  moq.ParamIndexing
			}
		}
	}
	// MoqScanner_mock isolates the mock interface of the Scanner type
}

type MoqScanner_mock struct {
	Moq *MoqScanner
}

// MoqScanner_recorder isolates the recorder interface of the Scanner type
type MoqScanner_recorder struct {
	Moq *MoqScanner
}

// MoqScanner_Scan_params holds the params of the Scanner type
type MoqScanner_Scan_params struct {
	State fmt.ScanState
	Verb  rune
}

// MoqScanner_Scan_paramsKey holds the map key params of the Scanner type
type MoqScanner_Scan_paramsKey struct {
	Params struct {
		State fmt.ScanState
		Verb  rune
	}
	Hashes struct {
		State hash.Hash
		Verb  hash.Hash
	}
}

// MoqScanner_Scan_resultsByParams contains the results for a given set of
// parameters for the Scanner type
type MoqScanner_Scan_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqScanner_Scan_paramsKey]*MoqScanner_Scan_results
}

// MoqScanner_Scan_doFn defines the type of function needed when calling AndDo
// for the Scanner type
type MoqScanner_Scan_doFn func(state fmt.ScanState, verb rune)

// MoqScanner_Scan_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Scanner type
type MoqScanner_Scan_doReturnFn func(state fmt.ScanState, verb rune) error

// MoqScanner_Scan_results holds the results of the Scanner type
type MoqScanner_Scan_results struct {
	Params  MoqScanner_Scan_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqScanner_Scan_doFn
		DoReturnFn MoqScanner_Scan_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqScanner_Scan_fnRecorder routes recorded function calls to the MoqScanner
// moq
type MoqScanner_Scan_fnRecorder struct {
	Params    MoqScanner_Scan_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqScanner_Scan_results
	Moq       *MoqScanner
}

// MoqScanner_Scan_anyParams isolates the any params functions of the Scanner
// type
type MoqScanner_Scan_anyParams struct {
	Recorder *MoqScanner_Scan_fnRecorder
}

// NewMoqScanner creates a new moq of the Scanner type
func NewMoqScanner(scene *moq.Scene, config *moq.Config) *MoqScanner {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqScanner{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqScanner_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Scan struct {
					State moq.ParamIndexing
					Verb  moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Scan struct {
				State moq.ParamIndexing
				Verb  moq.ParamIndexing
			}
		}{
			Scan: struct {
				State moq.ParamIndexing
				Verb  moq.ParamIndexing
			}{
				State: moq.ParamIndexByHash,
				Verb:  moq.ParamIndexByValue,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Scanner type
func (m *MoqScanner) Mock() *MoqScanner_mock { return m.Moq }

func (m *MoqScanner_mock) Scan(state fmt.ScanState, verb rune) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqScanner_Scan_params{
		State: state,
		Verb:  verb,
	}
	var results *MoqScanner_Scan_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Scan {
		paramsKey := m.Moq.ParamsKey_Scan(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Scan(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Scan(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Scan(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(state, verb)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(state, verb)
	}
	return
}

// OnCall returns the recorder implementation of the Scanner type
func (m *MoqScanner) OnCall() *MoqScanner_recorder {
	return &MoqScanner_recorder{
		Moq: m,
	}
}

func (m *MoqScanner_recorder) Scan(state fmt.ScanState, verb rune) *MoqScanner_Scan_fnRecorder {
	return &MoqScanner_Scan_fnRecorder{
		Params: MoqScanner_Scan_params{
			State: state,
			Verb:  verb,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqScanner_Scan_fnRecorder) Any() *MoqScanner_Scan_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Scan(r.Params))
		return nil
	}
	return &MoqScanner_Scan_anyParams{Recorder: r}
}

func (a *MoqScanner_Scan_anyParams) State() *MoqScanner_Scan_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqScanner_Scan_anyParams) Verb() *MoqScanner_Scan_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqScanner_Scan_fnRecorder) Seq() *MoqScanner_Scan_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Scan(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqScanner_Scan_fnRecorder) NoSeq() *MoqScanner_Scan_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Scan(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqScanner_Scan_fnRecorder) ReturnResults(result1 error) *MoqScanner_Scan_fnRecorder {
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
		DoFn       MoqScanner_Scan_doFn
		DoReturnFn MoqScanner_Scan_doReturnFn
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

func (r *MoqScanner_Scan_fnRecorder) AndDo(fn MoqScanner_Scan_doFn) *MoqScanner_Scan_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqScanner_Scan_fnRecorder) DoReturnResults(fn MoqScanner_Scan_doReturnFn) *MoqScanner_Scan_fnRecorder {
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
		DoFn       MoqScanner_Scan_doFn
		DoReturnFn MoqScanner_Scan_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqScanner_Scan_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqScanner_Scan_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Scan {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqScanner_Scan_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqScanner_Scan_paramsKey]*MoqScanner_Scan_results{},
		}
		r.Moq.ResultsByParams_Scan = append(r.Moq.ResultsByParams_Scan, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Scan) {
			copy(r.Moq.ResultsByParams_Scan[insertAt+1:], r.Moq.ResultsByParams_Scan[insertAt:0])
			r.Moq.ResultsByParams_Scan[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Scan(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqScanner_Scan_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqScanner_Scan_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqScanner_Scan_fnRecorder {
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
				DoFn       MoqScanner_Scan_doFn
				DoReturnFn MoqScanner_Scan_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqScanner) PrettyParams_Scan(params MoqScanner_Scan_params) string {
	return fmt.Sprintf("Scan(%#v, %#v)", params.State, params.Verb)
}

func (m *MoqScanner) ParamsKey_Scan(params MoqScanner_Scan_params, anyParams uint64) MoqScanner_Scan_paramsKey {
	m.Scene.T.Helper()
	var stateUsed fmt.ScanState
	var stateUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Scan.State == moq.ParamIndexByValue {
			stateUsed = params.State
		} else {
			stateUsedHash = hash.DeepHash(params.State)
		}
	}
	var verbUsed rune
	var verbUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Scan.Verb == moq.ParamIndexByValue {
			verbUsed = params.Verb
		} else {
			verbUsedHash = hash.DeepHash(params.Verb)
		}
	}
	return MoqScanner_Scan_paramsKey{
		Params: struct {
			State fmt.ScanState
			Verb  rune
		}{
			State: stateUsed,
			Verb:  verbUsed,
		},
		Hashes: struct {
			State hash.Hash
			Verb  hash.Hash
		}{
			State: stateUsedHash,
			Verb:  verbUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqScanner) Reset() { m.ResultsByParams_Scan = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqScanner) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Scan {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Scan(results.Params))
			}
		}
	}
}

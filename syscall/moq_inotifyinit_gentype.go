// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// InotifyInit_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type InotifyInit_genType func() (fd int, err error)

// MoqInotifyInit_genType holds the state of a moq of the InotifyInit_genType
// type
type MoqInotifyInit_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqInotifyInit_genType_mock

	ResultsByParams []MoqInotifyInit_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct{}
	}
}

// MoqInotifyInit_genType_mock isolates the mock interface of the
// InotifyInit_genType type
type MoqInotifyInit_genType_mock struct {
	Moq *MoqInotifyInit_genType
}

// MoqInotifyInit_genType_params holds the params of the InotifyInit_genType
// type
type MoqInotifyInit_genType_params struct{}

// MoqInotifyInit_genType_paramsKey holds the map key params of the
// InotifyInit_genType type
type MoqInotifyInit_genType_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqInotifyInit_genType_resultsByParams contains the results for a given set
// of parameters for the InotifyInit_genType type
type MoqInotifyInit_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqInotifyInit_genType_paramsKey]*MoqInotifyInit_genType_results
}

// MoqInotifyInit_genType_doFn defines the type of function needed when calling
// AndDo for the InotifyInit_genType type
type MoqInotifyInit_genType_doFn func()

// MoqInotifyInit_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the InotifyInit_genType type
type MoqInotifyInit_genType_doReturnFn func() (fd int, err error)

// MoqInotifyInit_genType_results holds the results of the InotifyInit_genType
// type
type MoqInotifyInit_genType_results struct {
	Params  MoqInotifyInit_genType_params
	Results []struct {
		Values *struct {
			Fd  int
			Err error
		}
		Sequence   uint32
		DoFn       MoqInotifyInit_genType_doFn
		DoReturnFn MoqInotifyInit_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqInotifyInit_genType_fnRecorder routes recorded function calls to the
// MoqInotifyInit_genType moq
type MoqInotifyInit_genType_fnRecorder struct {
	Params    MoqInotifyInit_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqInotifyInit_genType_results
	Moq       *MoqInotifyInit_genType
}

// MoqInotifyInit_genType_anyParams isolates the any params functions of the
// InotifyInit_genType type
type MoqInotifyInit_genType_anyParams struct {
	Recorder *MoqInotifyInit_genType_fnRecorder
}

// NewMoqInotifyInit_genType creates a new moq of the InotifyInit_genType type
func NewMoqInotifyInit_genType(scene *moq.Scene, config *moq.Config) *MoqInotifyInit_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqInotifyInit_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqInotifyInit_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct{}
		}{ParameterIndexing: struct{}{}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the InotifyInit_genType type
func (m *MoqInotifyInit_genType) Mock() InotifyInit_genType {
	return func() (_ int, _ error) {
		m.Scene.T.Helper()
		moq := &MoqInotifyInit_genType_mock{Moq: m}
		return moq.Fn()
	}
}

func (m *MoqInotifyInit_genType_mock) Fn() (fd int, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqInotifyInit_genType_params{}
	var results *MoqInotifyInit_genType_results
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
		result.DoFn()
	}

	if result.Values != nil {
		fd = result.Values.Fd
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		fd, err = result.DoReturnFn()
	}
	return
}

func (m *MoqInotifyInit_genType) OnCall() *MoqInotifyInit_genType_fnRecorder {
	return &MoqInotifyInit_genType_fnRecorder{
		Params:   MoqInotifyInit_genType_params{},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqInotifyInit_genType_fnRecorder) Any() *MoqInotifyInit_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqInotifyInit_genType_anyParams{Recorder: r}
}

func (r *MoqInotifyInit_genType_fnRecorder) Seq() *MoqInotifyInit_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqInotifyInit_genType_fnRecorder) NoSeq() *MoqInotifyInit_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqInotifyInit_genType_fnRecorder) ReturnResults(fd int, err error) *MoqInotifyInit_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Fd  int
			Err error
		}
		Sequence   uint32
		DoFn       MoqInotifyInit_genType_doFn
		DoReturnFn MoqInotifyInit_genType_doReturnFn
	}{
		Values: &struct {
			Fd  int
			Err error
		}{
			Fd:  fd,
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqInotifyInit_genType_fnRecorder) AndDo(fn MoqInotifyInit_genType_doFn) *MoqInotifyInit_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqInotifyInit_genType_fnRecorder) DoReturnResults(fn MoqInotifyInit_genType_doReturnFn) *MoqInotifyInit_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Fd  int
			Err error
		}
		Sequence   uint32
		DoFn       MoqInotifyInit_genType_doFn
		DoReturnFn MoqInotifyInit_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqInotifyInit_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqInotifyInit_genType_resultsByParams
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
		results = &MoqInotifyInit_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqInotifyInit_genType_paramsKey]*MoqInotifyInit_genType_results{},
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
		r.Results = &MoqInotifyInit_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqInotifyInit_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqInotifyInit_genType_fnRecorder {
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
					Fd  int
					Err error
				}
				Sequence   uint32
				DoFn       MoqInotifyInit_genType_doFn
				DoReturnFn MoqInotifyInit_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqInotifyInit_genType) PrettyParams(params MoqInotifyInit_genType_params) string {
	return fmt.Sprintf("InotifyInit_genType()")
}

func (m *MoqInotifyInit_genType) ParamsKey(params MoqInotifyInit_genType_params, anyParams uint64) MoqInotifyInit_genType_paramsKey {
	m.Scene.T.Helper()
	return MoqInotifyInit_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqInotifyInit_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqInotifyInit_genType) AssertExpectationsMet() {
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

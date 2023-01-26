// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// PtracePeekText_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type PtracePeekText_genType func(pid int, addr uintptr, out []byte) (count int, err error)

// MoqPtracePeekText_genType holds the state of a moq of the
// PtracePeekText_genType type
type MoqPtracePeekText_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqPtracePeekText_genType_mock

	ResultsByParams []MoqPtracePeekText_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Pid  moq.ParamIndexing
			Addr moq.ParamIndexing
			Out  moq.ParamIndexing
		}
	}
}

// MoqPtracePeekText_genType_mock isolates the mock interface of the
// PtracePeekText_genType type
type MoqPtracePeekText_genType_mock struct {
	Moq *MoqPtracePeekText_genType
}

// MoqPtracePeekText_genType_params holds the params of the
// PtracePeekText_genType type
type MoqPtracePeekText_genType_params struct {
	Pid  int
	Addr uintptr
	Out  []byte
}

// MoqPtracePeekText_genType_paramsKey holds the map key params of the
// PtracePeekText_genType type
type MoqPtracePeekText_genType_paramsKey struct {
	Params struct {
		Pid  int
		Addr uintptr
	}
	Hashes struct {
		Pid  hash.Hash
		Addr hash.Hash
		Out  hash.Hash
	}
}

// MoqPtracePeekText_genType_resultsByParams contains the results for a given
// set of parameters for the PtracePeekText_genType type
type MoqPtracePeekText_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqPtracePeekText_genType_paramsKey]*MoqPtracePeekText_genType_results
}

// MoqPtracePeekText_genType_doFn defines the type of function needed when
// calling AndDo for the PtracePeekText_genType type
type MoqPtracePeekText_genType_doFn func(pid int, addr uintptr, out []byte)

// MoqPtracePeekText_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the PtracePeekText_genType type
type MoqPtracePeekText_genType_doReturnFn func(pid int, addr uintptr, out []byte) (count int, err error)

// MoqPtracePeekText_genType_results holds the results of the
// PtracePeekText_genType type
type MoqPtracePeekText_genType_results struct {
	Params  MoqPtracePeekText_genType_params
	Results []struct {
		Values *struct {
			Count int
			Err   error
		}
		Sequence   uint32
		DoFn       MoqPtracePeekText_genType_doFn
		DoReturnFn MoqPtracePeekText_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqPtracePeekText_genType_fnRecorder routes recorded function calls to the
// MoqPtracePeekText_genType moq
type MoqPtracePeekText_genType_fnRecorder struct {
	Params    MoqPtracePeekText_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqPtracePeekText_genType_results
	Moq       *MoqPtracePeekText_genType
}

// MoqPtracePeekText_genType_anyParams isolates the any params functions of the
// PtracePeekText_genType type
type MoqPtracePeekText_genType_anyParams struct {
	Recorder *MoqPtracePeekText_genType_fnRecorder
}

// NewMoqPtracePeekText_genType creates a new moq of the PtracePeekText_genType
// type
func NewMoqPtracePeekText_genType(scene *moq.Scene, config *moq.Config) *MoqPtracePeekText_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqPtracePeekText_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqPtracePeekText_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Pid  moq.ParamIndexing
				Addr moq.ParamIndexing
				Out  moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Pid  moq.ParamIndexing
			Addr moq.ParamIndexing
			Out  moq.ParamIndexing
		}{
			Pid:  moq.ParamIndexByValue,
			Addr: moq.ParamIndexByValue,
			Out:  moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the PtracePeekText_genType type
func (m *MoqPtracePeekText_genType) Mock() PtracePeekText_genType {
	return func(pid int, addr uintptr, out []byte) (_ int, _ error) {
		m.Scene.T.Helper()
		moq := &MoqPtracePeekText_genType_mock{Moq: m}
		return moq.Fn(pid, addr, out)
	}
}

func (m *MoqPtracePeekText_genType_mock) Fn(pid int, addr uintptr, out []byte) (count int, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqPtracePeekText_genType_params{
		Pid:  pid,
		Addr: addr,
		Out:  out,
	}
	var results *MoqPtracePeekText_genType_results
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
		result.DoFn(pid, addr, out)
	}

	if result.Values != nil {
		count = result.Values.Count
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		count, err = result.DoReturnFn(pid, addr, out)
	}
	return
}

func (m *MoqPtracePeekText_genType) OnCall(pid int, addr uintptr, out []byte) *MoqPtracePeekText_genType_fnRecorder {
	return &MoqPtracePeekText_genType_fnRecorder{
		Params: MoqPtracePeekText_genType_params{
			Pid:  pid,
			Addr: addr,
			Out:  out,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqPtracePeekText_genType_fnRecorder) Any() *MoqPtracePeekText_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqPtracePeekText_genType_anyParams{Recorder: r}
}

func (a *MoqPtracePeekText_genType_anyParams) Pid() *MoqPtracePeekText_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqPtracePeekText_genType_anyParams) Addr() *MoqPtracePeekText_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqPtracePeekText_genType_anyParams) Out() *MoqPtracePeekText_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqPtracePeekText_genType_fnRecorder) Seq() *MoqPtracePeekText_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqPtracePeekText_genType_fnRecorder) NoSeq() *MoqPtracePeekText_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqPtracePeekText_genType_fnRecorder) ReturnResults(count int, err error) *MoqPtracePeekText_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Count int
			Err   error
		}
		Sequence   uint32
		DoFn       MoqPtracePeekText_genType_doFn
		DoReturnFn MoqPtracePeekText_genType_doReturnFn
	}{
		Values: &struct {
			Count int
			Err   error
		}{
			Count: count,
			Err:   err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqPtracePeekText_genType_fnRecorder) AndDo(fn MoqPtracePeekText_genType_doFn) *MoqPtracePeekText_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqPtracePeekText_genType_fnRecorder) DoReturnResults(fn MoqPtracePeekText_genType_doReturnFn) *MoqPtracePeekText_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Count int
			Err   error
		}
		Sequence   uint32
		DoFn       MoqPtracePeekText_genType_doFn
		DoReturnFn MoqPtracePeekText_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqPtracePeekText_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqPtracePeekText_genType_resultsByParams
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
		results = &MoqPtracePeekText_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqPtracePeekText_genType_paramsKey]*MoqPtracePeekText_genType_results{},
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
		r.Results = &MoqPtracePeekText_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqPtracePeekText_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqPtracePeekText_genType_fnRecorder {
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
					Count int
					Err   error
				}
				Sequence   uint32
				DoFn       MoqPtracePeekText_genType_doFn
				DoReturnFn MoqPtracePeekText_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqPtracePeekText_genType) PrettyParams(params MoqPtracePeekText_genType_params) string {
	return fmt.Sprintf("PtracePeekText_genType(%#v, %#v, %#v)", params.Pid, params.Addr, params.Out)
}

func (m *MoqPtracePeekText_genType) ParamsKey(params MoqPtracePeekText_genType_params, anyParams uint64) MoqPtracePeekText_genType_paramsKey {
	m.Scene.T.Helper()
	var pidUsed int
	var pidUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Pid == moq.ParamIndexByValue {
			pidUsed = params.Pid
		} else {
			pidUsedHash = hash.DeepHash(params.Pid)
		}
	}
	var addrUsed uintptr
	var addrUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Addr == moq.ParamIndexByValue {
			addrUsed = params.Addr
		} else {
			addrUsedHash = hash.DeepHash(params.Addr)
		}
	}
	var outUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Out == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The out parameter can't be indexed by value")
		}
		outUsedHash = hash.DeepHash(params.Out)
	}
	return MoqPtracePeekText_genType_paramsKey{
		Params: struct {
			Pid  int
			Addr uintptr
		}{
			Pid:  pidUsed,
			Addr: addrUsed,
		},
		Hashes: struct {
			Pid  hash.Hash
			Addr hash.Hash
			Out  hash.Hash
		}{
			Pid:  pidUsedHash,
			Addr: addrUsedHash,
			Out:  outUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqPtracePeekText_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqPtracePeekText_genType) AssertExpectationsMet() {
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
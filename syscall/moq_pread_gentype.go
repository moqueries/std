// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Pread_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Pread_genType func(fd int, p []byte, offset int64) (n int, err error)

// MoqPread_genType holds the state of a moq of the Pread_genType type
type MoqPread_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqPread_genType_mock

	ResultsByParams []MoqPread_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Fd     moq.ParamIndexing
			P      moq.ParamIndexing
			Offset moq.ParamIndexing
		}
	}
}

// MoqPread_genType_mock isolates the mock interface of the Pread_genType type
type MoqPread_genType_mock struct {
	Moq *MoqPread_genType
}

// MoqPread_genType_params holds the params of the Pread_genType type
type MoqPread_genType_params struct {
	Fd     int
	P      []byte
	Offset int64
}

// MoqPread_genType_paramsKey holds the map key params of the Pread_genType
// type
type MoqPread_genType_paramsKey struct {
	Params struct {
		Fd     int
		Offset int64
	}
	Hashes struct {
		Fd     hash.Hash
		P      hash.Hash
		Offset hash.Hash
	}
}

// MoqPread_genType_resultsByParams contains the results for a given set of
// parameters for the Pread_genType type
type MoqPread_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqPread_genType_paramsKey]*MoqPread_genType_results
}

// MoqPread_genType_doFn defines the type of function needed when calling AndDo
// for the Pread_genType type
type MoqPread_genType_doFn func(fd int, p []byte, offset int64)

// MoqPread_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Pread_genType type
type MoqPread_genType_doReturnFn func(fd int, p []byte, offset int64) (n int, err error)

// MoqPread_genType_results holds the results of the Pread_genType type
type MoqPread_genType_results struct {
	Params  MoqPread_genType_params
	Results []struct {
		Values *struct {
			N   int
			Err error
		}
		Sequence   uint32
		DoFn       MoqPread_genType_doFn
		DoReturnFn MoqPread_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqPread_genType_fnRecorder routes recorded function calls to the
// MoqPread_genType moq
type MoqPread_genType_fnRecorder struct {
	Params    MoqPread_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqPread_genType_results
	Moq       *MoqPread_genType
}

// MoqPread_genType_anyParams isolates the any params functions of the
// Pread_genType type
type MoqPread_genType_anyParams struct {
	Recorder *MoqPread_genType_fnRecorder
}

// NewMoqPread_genType creates a new moq of the Pread_genType type
func NewMoqPread_genType(scene *moq.Scene, config *moq.Config) *MoqPread_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqPread_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqPread_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Fd     moq.ParamIndexing
				P      moq.ParamIndexing
				Offset moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Fd     moq.ParamIndexing
			P      moq.ParamIndexing
			Offset moq.ParamIndexing
		}{
			Fd:     moq.ParamIndexByValue,
			P:      moq.ParamIndexByHash,
			Offset: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Pread_genType type
func (m *MoqPread_genType) Mock() Pread_genType {
	return func(fd int, p []byte, offset int64) (_ int, _ error) {
		m.Scene.T.Helper()
		moq := &MoqPread_genType_mock{Moq: m}
		return moq.Fn(fd, p, offset)
	}
}

func (m *MoqPread_genType_mock) Fn(fd int, p []byte, offset int64) (n int, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqPread_genType_params{
		Fd:     fd,
		P:      p,
		Offset: offset,
	}
	var results *MoqPread_genType_results
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
		result.DoFn(fd, p, offset)
	}

	if result.Values != nil {
		n = result.Values.N
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		n, err = result.DoReturnFn(fd, p, offset)
	}
	return
}

func (m *MoqPread_genType) OnCall(fd int, p []byte, offset int64) *MoqPread_genType_fnRecorder {
	return &MoqPread_genType_fnRecorder{
		Params: MoqPread_genType_params{
			Fd:     fd,
			P:      p,
			Offset: offset,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqPread_genType_fnRecorder) Any() *MoqPread_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqPread_genType_anyParams{Recorder: r}
}

func (a *MoqPread_genType_anyParams) Fd() *MoqPread_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqPread_genType_anyParams) P() *MoqPread_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqPread_genType_anyParams) Offset() *MoqPread_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqPread_genType_fnRecorder) Seq() *MoqPread_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqPread_genType_fnRecorder) NoSeq() *MoqPread_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqPread_genType_fnRecorder) ReturnResults(n int, err error) *MoqPread_genType_fnRecorder {
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
		DoFn       MoqPread_genType_doFn
		DoReturnFn MoqPread_genType_doReturnFn
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

func (r *MoqPread_genType_fnRecorder) AndDo(fn MoqPread_genType_doFn) *MoqPread_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqPread_genType_fnRecorder) DoReturnResults(fn MoqPread_genType_doReturnFn) *MoqPread_genType_fnRecorder {
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
		DoFn       MoqPread_genType_doFn
		DoReturnFn MoqPread_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqPread_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqPread_genType_resultsByParams
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
		results = &MoqPread_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqPread_genType_paramsKey]*MoqPread_genType_results{},
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
		r.Results = &MoqPread_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqPread_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqPread_genType_fnRecorder {
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
				DoFn       MoqPread_genType_doFn
				DoReturnFn MoqPread_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqPread_genType) PrettyParams(params MoqPread_genType_params) string {
	return fmt.Sprintf("Pread_genType(%#v, %#v, %#v)", params.Fd, params.P, params.Offset)
}

func (m *MoqPread_genType) ParamsKey(params MoqPread_genType_params, anyParams uint64) MoqPread_genType_paramsKey {
	m.Scene.T.Helper()
	var fdUsed int
	var fdUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Fd == moq.ParamIndexByValue {
			fdUsed = params.Fd
		} else {
			fdUsedHash = hash.DeepHash(params.Fd)
		}
	}
	var pUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.P == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The p parameter can't be indexed by value")
		}
		pUsedHash = hash.DeepHash(params.P)
	}
	var offsetUsed int64
	var offsetUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Offset == moq.ParamIndexByValue {
			offsetUsed = params.Offset
		} else {
			offsetUsedHash = hash.DeepHash(params.Offset)
		}
	}
	return MoqPread_genType_paramsKey{
		Params: struct {
			Fd     int
			Offset int64
		}{
			Fd:     fdUsed,
			Offset: offsetUsed,
		},
		Hashes: struct {
			Fd     hash.Hash
			P      hash.Hash
			Offset hash.Hash
		}{
			Fd:     fdUsedHash,
			P:      pUsedHash,
			Offset: offsetUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqPread_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqPread_genType) AssertExpectationsMet() {
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

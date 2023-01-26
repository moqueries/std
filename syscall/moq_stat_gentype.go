// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"
	"syscall"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Stat_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Stat_genType func(path string, stat *syscall.Stat_t) (err error)

// MoqStat_genType holds the state of a moq of the Stat_genType type
type MoqStat_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqStat_genType_mock

	ResultsByParams []MoqStat_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Path moq.ParamIndexing
			Stat moq.ParamIndexing
		}
	}
}

// MoqStat_genType_mock isolates the mock interface of the Stat_genType type
type MoqStat_genType_mock struct {
	Moq *MoqStat_genType
}

// MoqStat_genType_params holds the params of the Stat_genType type
type MoqStat_genType_params struct {
	Path string
	Stat *syscall.Stat_t
}

// MoqStat_genType_paramsKey holds the map key params of the Stat_genType type
type MoqStat_genType_paramsKey struct {
	Params struct {
		Path string
		Stat *syscall.Stat_t
	}
	Hashes struct {
		Path hash.Hash
		Stat hash.Hash
	}
}

// MoqStat_genType_resultsByParams contains the results for a given set of
// parameters for the Stat_genType type
type MoqStat_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqStat_genType_paramsKey]*MoqStat_genType_results
}

// MoqStat_genType_doFn defines the type of function needed when calling AndDo
// for the Stat_genType type
type MoqStat_genType_doFn func(path string, stat *syscall.Stat_t)

// MoqStat_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Stat_genType type
type MoqStat_genType_doReturnFn func(path string, stat *syscall.Stat_t) (err error)

// MoqStat_genType_results holds the results of the Stat_genType type
type MoqStat_genType_results struct {
	Params  MoqStat_genType_params
	Results []struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqStat_genType_doFn
		DoReturnFn MoqStat_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqStat_genType_fnRecorder routes recorded function calls to the
// MoqStat_genType moq
type MoqStat_genType_fnRecorder struct {
	Params    MoqStat_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqStat_genType_results
	Moq       *MoqStat_genType
}

// MoqStat_genType_anyParams isolates the any params functions of the
// Stat_genType type
type MoqStat_genType_anyParams struct {
	Recorder *MoqStat_genType_fnRecorder
}

// NewMoqStat_genType creates a new moq of the Stat_genType type
func NewMoqStat_genType(scene *moq.Scene, config *moq.Config) *MoqStat_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqStat_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqStat_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Path moq.ParamIndexing
				Stat moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Path moq.ParamIndexing
			Stat moq.ParamIndexing
		}{
			Path: moq.ParamIndexByValue,
			Stat: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Stat_genType type
func (m *MoqStat_genType) Mock() Stat_genType {
	return func(path string, stat *syscall.Stat_t) (_ error) {
		m.Scene.T.Helper()
		moq := &MoqStat_genType_mock{Moq: m}
		return moq.Fn(path, stat)
	}
}

func (m *MoqStat_genType_mock) Fn(path string, stat *syscall.Stat_t) (err error) {
	m.Moq.Scene.T.Helper()
	params := MoqStat_genType_params{
		Path: path,
		Stat: stat,
	}
	var results *MoqStat_genType_results
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
		result.DoFn(path, stat)
	}

	if result.Values != nil {
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		err = result.DoReturnFn(path, stat)
	}
	return
}

func (m *MoqStat_genType) OnCall(path string, stat *syscall.Stat_t) *MoqStat_genType_fnRecorder {
	return &MoqStat_genType_fnRecorder{
		Params: MoqStat_genType_params{
			Path: path,
			Stat: stat,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqStat_genType_fnRecorder) Any() *MoqStat_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqStat_genType_anyParams{Recorder: r}
}

func (a *MoqStat_genType_anyParams) Path() *MoqStat_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqStat_genType_anyParams) Stat() *MoqStat_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqStat_genType_fnRecorder) Seq() *MoqStat_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqStat_genType_fnRecorder) NoSeq() *MoqStat_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqStat_genType_fnRecorder) ReturnResults(err error) *MoqStat_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqStat_genType_doFn
		DoReturnFn MoqStat_genType_doReturnFn
	}{
		Values: &struct{ Err error }{
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqStat_genType_fnRecorder) AndDo(fn MoqStat_genType_doFn) *MoqStat_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqStat_genType_fnRecorder) DoReturnResults(fn MoqStat_genType_doReturnFn) *MoqStat_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqStat_genType_doFn
		DoReturnFn MoqStat_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqStat_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqStat_genType_resultsByParams
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
		results = &MoqStat_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqStat_genType_paramsKey]*MoqStat_genType_results{},
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
		r.Results = &MoqStat_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqStat_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqStat_genType_fnRecorder {
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
				Values     *struct{ Err error }
				Sequence   uint32
				DoFn       MoqStat_genType_doFn
				DoReturnFn MoqStat_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqStat_genType) PrettyParams(params MoqStat_genType_params) string {
	return fmt.Sprintf("Stat_genType(%#v, %#v)", params.Path, params.Stat)
}

func (m *MoqStat_genType) ParamsKey(params MoqStat_genType_params, anyParams uint64) MoqStat_genType_paramsKey {
	m.Scene.T.Helper()
	var pathUsed string
	var pathUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Path == moq.ParamIndexByValue {
			pathUsed = params.Path
		} else {
			pathUsedHash = hash.DeepHash(params.Path)
		}
	}
	var statUsed *syscall.Stat_t
	var statUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Stat == moq.ParamIndexByValue {
			statUsed = params.Stat
		} else {
			statUsedHash = hash.DeepHash(params.Stat)
		}
	}
	return MoqStat_genType_paramsKey{
		Params: struct {
			Path string
			Stat *syscall.Stat_t
		}{
			Path: pathUsed,
			Stat: statUsed,
		},
		Hashes: struct {
			Path hash.Hash
			Stat hash.Hash
		}{
			Path: pathUsedHash,
			Stat: statUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqStat_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqStat_genType) AssertExpectationsMet() {
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

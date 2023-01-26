// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Lchown_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Lchown_genType func(path string, uid int, gid int) (err error)

// MoqLchown_genType holds the state of a moq of the Lchown_genType type
type MoqLchown_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqLchown_genType_mock

	ResultsByParams []MoqLchown_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Path moq.ParamIndexing
			Uid  moq.ParamIndexing
			Gid  moq.ParamIndexing
		}
	}
}

// MoqLchown_genType_mock isolates the mock interface of the Lchown_genType
// type
type MoqLchown_genType_mock struct {
	Moq *MoqLchown_genType
}

// MoqLchown_genType_params holds the params of the Lchown_genType type
type MoqLchown_genType_params struct {
	Path string
	Uid  int
	Gid  int
}

// MoqLchown_genType_paramsKey holds the map key params of the Lchown_genType
// type
type MoqLchown_genType_paramsKey struct {
	Params struct {
		Path string
		Uid  int
		Gid  int
	}
	Hashes struct {
		Path hash.Hash
		Uid  hash.Hash
		Gid  hash.Hash
	}
}

// MoqLchown_genType_resultsByParams contains the results for a given set of
// parameters for the Lchown_genType type
type MoqLchown_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqLchown_genType_paramsKey]*MoqLchown_genType_results
}

// MoqLchown_genType_doFn defines the type of function needed when calling
// AndDo for the Lchown_genType type
type MoqLchown_genType_doFn func(path string, uid int, gid int)

// MoqLchown_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Lchown_genType type
type MoqLchown_genType_doReturnFn func(path string, uid int, gid int) (err error)

// MoqLchown_genType_results holds the results of the Lchown_genType type
type MoqLchown_genType_results struct {
	Params  MoqLchown_genType_params
	Results []struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqLchown_genType_doFn
		DoReturnFn MoqLchown_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqLchown_genType_fnRecorder routes recorded function calls to the
// MoqLchown_genType moq
type MoqLchown_genType_fnRecorder struct {
	Params    MoqLchown_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqLchown_genType_results
	Moq       *MoqLchown_genType
}

// MoqLchown_genType_anyParams isolates the any params functions of the
// Lchown_genType type
type MoqLchown_genType_anyParams struct {
	Recorder *MoqLchown_genType_fnRecorder
}

// NewMoqLchown_genType creates a new moq of the Lchown_genType type
func NewMoqLchown_genType(scene *moq.Scene, config *moq.Config) *MoqLchown_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqLchown_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqLchown_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Path moq.ParamIndexing
				Uid  moq.ParamIndexing
				Gid  moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Path moq.ParamIndexing
			Uid  moq.ParamIndexing
			Gid  moq.ParamIndexing
		}{
			Path: moq.ParamIndexByValue,
			Uid:  moq.ParamIndexByValue,
			Gid:  moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Lchown_genType type
func (m *MoqLchown_genType) Mock() Lchown_genType {
	return func(path string, uid int, gid int) (_ error) {
		m.Scene.T.Helper()
		moq := &MoqLchown_genType_mock{Moq: m}
		return moq.Fn(path, uid, gid)
	}
}

func (m *MoqLchown_genType_mock) Fn(path string, uid int, gid int) (err error) {
	m.Moq.Scene.T.Helper()
	params := MoqLchown_genType_params{
		Path: path,
		Uid:  uid,
		Gid:  gid,
	}
	var results *MoqLchown_genType_results
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
		result.DoFn(path, uid, gid)
	}

	if result.Values != nil {
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		err = result.DoReturnFn(path, uid, gid)
	}
	return
}

func (m *MoqLchown_genType) OnCall(path string, uid int, gid int) *MoqLchown_genType_fnRecorder {
	return &MoqLchown_genType_fnRecorder{
		Params: MoqLchown_genType_params{
			Path: path,
			Uid:  uid,
			Gid:  gid,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqLchown_genType_fnRecorder) Any() *MoqLchown_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqLchown_genType_anyParams{Recorder: r}
}

func (a *MoqLchown_genType_anyParams) Path() *MoqLchown_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqLchown_genType_anyParams) Uid() *MoqLchown_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqLchown_genType_anyParams) Gid() *MoqLchown_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqLchown_genType_fnRecorder) Seq() *MoqLchown_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqLchown_genType_fnRecorder) NoSeq() *MoqLchown_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqLchown_genType_fnRecorder) ReturnResults(err error) *MoqLchown_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqLchown_genType_doFn
		DoReturnFn MoqLchown_genType_doReturnFn
	}{
		Values: &struct{ Err error }{
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqLchown_genType_fnRecorder) AndDo(fn MoqLchown_genType_doFn) *MoqLchown_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqLchown_genType_fnRecorder) DoReturnResults(fn MoqLchown_genType_doReturnFn) *MoqLchown_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqLchown_genType_doFn
		DoReturnFn MoqLchown_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqLchown_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqLchown_genType_resultsByParams
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
		results = &MoqLchown_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqLchown_genType_paramsKey]*MoqLchown_genType_results{},
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
		r.Results = &MoqLchown_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqLchown_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqLchown_genType_fnRecorder {
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
				DoFn       MoqLchown_genType_doFn
				DoReturnFn MoqLchown_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqLchown_genType) PrettyParams(params MoqLchown_genType_params) string {
	return fmt.Sprintf("Lchown_genType(%#v, %#v, %#v)", params.Path, params.Uid, params.Gid)
}

func (m *MoqLchown_genType) ParamsKey(params MoqLchown_genType_params, anyParams uint64) MoqLchown_genType_paramsKey {
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
	var uidUsed int
	var uidUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Uid == moq.ParamIndexByValue {
			uidUsed = params.Uid
		} else {
			uidUsedHash = hash.DeepHash(params.Uid)
		}
	}
	var gidUsed int
	var gidUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Gid == moq.ParamIndexByValue {
			gidUsed = params.Gid
		} else {
			gidUsedHash = hash.DeepHash(params.Gid)
		}
	}
	return MoqLchown_genType_paramsKey{
		Params: struct {
			Path string
			Uid  int
			Gid  int
		}{
			Path: pathUsed,
			Uid:  uidUsed,
			Gid:  gidUsed,
		},
		Hashes: struct {
			Path hash.Hash
			Uid  hash.Hash
			Gid  hash.Hash
		}{
			Path: pathUsedHash,
			Uid:  uidUsedHash,
			Gid:  gidUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqLchown_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqLchown_genType) AssertExpectationsMet() {
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
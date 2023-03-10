// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package pkix

import (
	"crypto/x509/pkix"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that pkix.Name_starGenType is mocked
// completely
var _ Name_starGenType = (*MoqName_starGenType_mock)(nil)

// Name_starGenType is the fabricated implementation type of this mock (emitted
// when mocking a collections of methods directly and not from an interface
// type)
type Name_starGenType interface {
	FillFromRDNSequence(rdns *pkix.RDNSequence)
}

// MoqName_starGenType holds the state of a moq of the Name_starGenType type
type MoqName_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqName_starGenType_mock

	ResultsByParams_FillFromRDNSequence []MoqName_starGenType_FillFromRDNSequence_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			FillFromRDNSequence struct {
				Rdns moq.ParamIndexing
			}
		}
	}
	// MoqName_starGenType_mock isolates the mock interface of the Name_starGenType
}

// type
type MoqName_starGenType_mock struct {
	Moq *MoqName_starGenType
}

// MoqName_starGenType_recorder isolates the recorder interface of the
// Name_starGenType type
type MoqName_starGenType_recorder struct {
	Moq *MoqName_starGenType
}

// MoqName_starGenType_FillFromRDNSequence_params holds the params of the
// Name_starGenType type
type MoqName_starGenType_FillFromRDNSequence_params struct{ Rdns *pkix.RDNSequence }

// MoqName_starGenType_FillFromRDNSequence_paramsKey holds the map key params
// of the Name_starGenType type
type MoqName_starGenType_FillFromRDNSequence_paramsKey struct {
	Params struct{ Rdns *pkix.RDNSequence }
	Hashes struct{ Rdns hash.Hash }
}

// MoqName_starGenType_FillFromRDNSequence_resultsByParams contains the results
// for a given set of parameters for the Name_starGenType type
type MoqName_starGenType_FillFromRDNSequence_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqName_starGenType_FillFromRDNSequence_paramsKey]*MoqName_starGenType_FillFromRDNSequence_results
}

// MoqName_starGenType_FillFromRDNSequence_doFn defines the type of function
// needed when calling AndDo for the Name_starGenType type
type MoqName_starGenType_FillFromRDNSequence_doFn func(rdns *pkix.RDNSequence)

// MoqName_starGenType_FillFromRDNSequence_doReturnFn defines the type of
// function needed when calling DoReturnResults for the Name_starGenType type
type MoqName_starGenType_FillFromRDNSequence_doReturnFn func(rdns *pkix.RDNSequence)

// MoqName_starGenType_FillFromRDNSequence_results holds the results of the
// Name_starGenType type
type MoqName_starGenType_FillFromRDNSequence_results struct {
	Params  MoqName_starGenType_FillFromRDNSequence_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqName_starGenType_FillFromRDNSequence_doFn
		DoReturnFn MoqName_starGenType_FillFromRDNSequence_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqName_starGenType_FillFromRDNSequence_fnRecorder routes recorded function
// calls to the MoqName_starGenType moq
type MoqName_starGenType_FillFromRDNSequence_fnRecorder struct {
	Params    MoqName_starGenType_FillFromRDNSequence_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqName_starGenType_FillFromRDNSequence_results
	Moq       *MoqName_starGenType
}

// MoqName_starGenType_FillFromRDNSequence_anyParams isolates the any params
// functions of the Name_starGenType type
type MoqName_starGenType_FillFromRDNSequence_anyParams struct {
	Recorder *MoqName_starGenType_FillFromRDNSequence_fnRecorder
}

// NewMoqName_starGenType creates a new moq of the Name_starGenType type
func NewMoqName_starGenType(scene *moq.Scene, config *moq.Config) *MoqName_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqName_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqName_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				FillFromRDNSequence struct {
					Rdns moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			FillFromRDNSequence struct {
				Rdns moq.ParamIndexing
			}
		}{
			FillFromRDNSequence: struct {
				Rdns moq.ParamIndexing
			}{
				Rdns: moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Name_starGenType type
func (m *MoqName_starGenType) Mock() *MoqName_starGenType_mock { return m.Moq }

func (m *MoqName_starGenType_mock) FillFromRDNSequence(rdns *pkix.RDNSequence) {
	m.Moq.Scene.T.Helper()
	params := MoqName_starGenType_FillFromRDNSequence_params{
		Rdns: rdns,
	}
	var results *MoqName_starGenType_FillFromRDNSequence_results
	for _, resultsByParams := range m.Moq.ResultsByParams_FillFromRDNSequence {
		paramsKey := m.Moq.ParamsKey_FillFromRDNSequence(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_FillFromRDNSequence(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_FillFromRDNSequence(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_FillFromRDNSequence(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(rdns)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(rdns)
	}
	return
}

// OnCall returns the recorder implementation of the Name_starGenType type
func (m *MoqName_starGenType) OnCall() *MoqName_starGenType_recorder {
	return &MoqName_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqName_starGenType_recorder) FillFromRDNSequence(rdns *pkix.RDNSequence) *MoqName_starGenType_FillFromRDNSequence_fnRecorder {
	return &MoqName_starGenType_FillFromRDNSequence_fnRecorder{
		Params: MoqName_starGenType_FillFromRDNSequence_params{
			Rdns: rdns,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqName_starGenType_FillFromRDNSequence_fnRecorder) Any() *MoqName_starGenType_FillFromRDNSequence_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_FillFromRDNSequence(r.Params))
		return nil
	}
	return &MoqName_starGenType_FillFromRDNSequence_anyParams{Recorder: r}
}

func (a *MoqName_starGenType_FillFromRDNSequence_anyParams) Rdns() *MoqName_starGenType_FillFromRDNSequence_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqName_starGenType_FillFromRDNSequence_fnRecorder) Seq() *MoqName_starGenType_FillFromRDNSequence_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_FillFromRDNSequence(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqName_starGenType_FillFromRDNSequence_fnRecorder) NoSeq() *MoqName_starGenType_FillFromRDNSequence_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_FillFromRDNSequence(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqName_starGenType_FillFromRDNSequence_fnRecorder) ReturnResults() *MoqName_starGenType_FillFromRDNSequence_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqName_starGenType_FillFromRDNSequence_doFn
		DoReturnFn MoqName_starGenType_FillFromRDNSequence_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqName_starGenType_FillFromRDNSequence_fnRecorder) AndDo(fn MoqName_starGenType_FillFromRDNSequence_doFn) *MoqName_starGenType_FillFromRDNSequence_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqName_starGenType_FillFromRDNSequence_fnRecorder) DoReturnResults(fn MoqName_starGenType_FillFromRDNSequence_doReturnFn) *MoqName_starGenType_FillFromRDNSequence_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqName_starGenType_FillFromRDNSequence_doFn
		DoReturnFn MoqName_starGenType_FillFromRDNSequence_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqName_starGenType_FillFromRDNSequence_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqName_starGenType_FillFromRDNSequence_resultsByParams
	for n, res := range r.Moq.ResultsByParams_FillFromRDNSequence {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqName_starGenType_FillFromRDNSequence_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqName_starGenType_FillFromRDNSequence_paramsKey]*MoqName_starGenType_FillFromRDNSequence_results{},
		}
		r.Moq.ResultsByParams_FillFromRDNSequence = append(r.Moq.ResultsByParams_FillFromRDNSequence, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_FillFromRDNSequence) {
			copy(r.Moq.ResultsByParams_FillFromRDNSequence[insertAt+1:], r.Moq.ResultsByParams_FillFromRDNSequence[insertAt:0])
			r.Moq.ResultsByParams_FillFromRDNSequence[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_FillFromRDNSequence(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqName_starGenType_FillFromRDNSequence_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqName_starGenType_FillFromRDNSequence_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqName_starGenType_FillFromRDNSequence_fnRecorder {
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
				DoFn       MoqName_starGenType_FillFromRDNSequence_doFn
				DoReturnFn MoqName_starGenType_FillFromRDNSequence_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqName_starGenType) PrettyParams_FillFromRDNSequence(params MoqName_starGenType_FillFromRDNSequence_params) string {
	return fmt.Sprintf("FillFromRDNSequence(%#v)", params.Rdns)
}

func (m *MoqName_starGenType) ParamsKey_FillFromRDNSequence(params MoqName_starGenType_FillFromRDNSequence_params, anyParams uint64) MoqName_starGenType_FillFromRDNSequence_paramsKey {
	m.Scene.T.Helper()
	var rdnsUsed *pkix.RDNSequence
	var rdnsUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.FillFromRDNSequence.Rdns == moq.ParamIndexByValue {
			rdnsUsed = params.Rdns
		} else {
			rdnsUsedHash = hash.DeepHash(params.Rdns)
		}
	}
	return MoqName_starGenType_FillFromRDNSequence_paramsKey{
		Params: struct{ Rdns *pkix.RDNSequence }{
			Rdns: rdnsUsed,
		},
		Hashes: struct{ Rdns hash.Hash }{
			Rdns: rdnsUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqName_starGenType) Reset() { m.ResultsByParams_FillFromRDNSequence = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqName_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_FillFromRDNSequence {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_FillFromRDNSequence(results.Params))
			}
		}
	}
}

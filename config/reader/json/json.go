package json

import (
	"errors"
	"time"

	"dario.cat/mergo"
	"github.com/tripleear/triear-go-core/config/encoder"
	"github.com/tripleear/triear-go-core/config/encoder/json"
	"github.com/tripleear/triear-go-core/config/reader"
	"github.com/tripleear/triear-go-core/config/source"
)

const readerTyp = "json"

type jsonReader struct {
	opts reader.Options
	json encoder.Encoder
}

func (j *jsonReader) Merge(changes ...*source.ChangeSet) (*source.ChangeSet, error) {
	var merged map[string]interface{}

	for _, m := range changes {
		if m == nil {
			continue
		}

		if len(m.Data) == 0 {
			continue
		}

		codec, ok := j.opts.Encoding[m.Format]
		if !ok {
			// fallback
			codec = j.json
		}

		var data map[string]interface{}
		if err := codec.Decode(m.Data, &data); err != nil {
			return nil, err
		}
		if err := mergo.Map(&merged, data, mergo.WithOverride); err != nil {
			return nil, err
		}
	}

	b, err := j.json.Encode(merged)
	if err != nil {
		return nil, err
	}

	cs := &source.ChangeSet{
		Timestamp: time.Now(),
		Data:      b,
		Source:    "json",
		Format:    j.json.String(),
	}
	cs.Checksum = cs.Sum()

	return cs, nil
}

func (j *jsonReader) Values(ch *source.ChangeSet) (reader.Values, error) {
	if ch == nil {
		return nil, errors.New("changeset is nil")
	}
	if ch.Format != "json" {
		return nil, errors.New("unsupported format")
	}
	return newValues(ch)
}

func (j *jsonReader) String() string {
	return "json"
}

// NewReader creates a json reader
func NewReader(opts ...reader.Option) reader.Reader {
	options := reader.NewOptions(opts...)
	return &jsonReader{
		json: json.NewEncoder(),
		opts: options,
	}
}

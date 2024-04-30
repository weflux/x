package structs

import "github.com/mitchellh/mapstructure"

func Convert(input interface{}, output interface{}, tagName ...string) error {
	tag := "redis"
	if len(tagName) > 1 {
		tag = tagName[0]
	}
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           output,
		TagName:          tag,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

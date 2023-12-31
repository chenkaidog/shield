package sensitive

import (
	"encoding/json"
)

type SensitiveMarshal struct {
	sensitiveSet map[string]bool
}

func NewSensitiveMarshal(words ...string) *SensitiveMarshal {
	set := make(map[string]bool)
	for _, word := range words {
		set[word] = true
	}

	return &SensitiveMarshal{
		sensitiveSet: set,
	}
}

func (sm *SensitiveMarshal) AddSensitiveWord(words ...string) {
	for _, word := range words {
		sm.sensitiveSet[word] = true
	}
}

func (sm *SensitiveMarshal) SafeMarshal(obj interface{}) string {
	return SafeJson(sm.sensitiveMarshal(obj))
}

func (sm *SensitiveMarshal) sensitiveMarshal(obj interface{}) interface{} {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil
	}

	sensitiveMapper := make(map[string]interface{})

	if err = json.Unmarshal(data, &sensitiveMapper); err != nil {
		return nil
	}

	return sm.maskData(sensitiveMapper)
}

func (sm *SensitiveMarshal) maskData(rawMapper map[string]interface{}) map[string]interface{} {
	for k, v := range rawMapper {
		if sm.sensitiveSet[k] {
			rawMapper[k] = "******"
		}
		if m, ok := v.(map[string]interface{}); ok {
			rawMapper[k] = sm.maskData(m)
		}
	}

	return rawMapper
}

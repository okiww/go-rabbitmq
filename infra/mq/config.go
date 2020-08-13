package mq

// ConfigQOS for define config QOS
type ConfigQOS struct {
	Count  int
	Size   int
	Global bool
}

// ConfigFunc is signature function for Configuration QOS
type ConfigFunc func(*ConfigQOS)

// SetQOSCount for set value of QOS Count
func (m *mqService) SetQOSCount(count int) ConfigFunc {
	return func(qos *ConfigQOS) {
		qos.Count = count
	}
}

// SetQOSCount for set value of QOS Size
func (m *mqService) SetQOSSize(size int) ConfigFunc {
	return func(qos *ConfigQOS) {
		qos.Size = size
	}
}

// SetQOSCount for set value of QOS Global
func (m *mqService) SetQOSGlobal(global bool) ConfigFunc {
	return func(qos *ConfigQOS) {
		qos.Global = global
	}
}

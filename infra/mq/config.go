package mq

type ConfigQOS struct {
	Count  int
	Size   int
	Global bool
}

type ConfigFunc func(*ConfigQOS)

func SetQOSCount(count int) ConfigFunc {
	return func(qos *ConfigQOS) {
		qos.Count = count
	}
}

func SetQOSSize(size int) ConfigFunc {
	return func(qos *ConfigQOS) {
		qos.Size = size
	}
}

func SetQOSGlobal(global bool) ConfigFunc {
	return func(qos *ConfigQOS) {
		qos.Global = global
	}
}

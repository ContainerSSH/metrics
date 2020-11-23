package metrics

type gaugeImpl struct {
	name      string
	collector *collector
}

func (g *gaugeImpl) Increment() {
	g.IncrementBy(1)
}

func (g *gaugeImpl) IncrementBy(by float64) {
	g.collector.mutex.Lock()
	defer g.collector.mutex.Unlock()

	value := g.collector.get(g.name, map[string]string{})
	g.collector.set(g.name, map[string]string{}, value+by)
}

func (g *gaugeImpl) Decrement() {
	g.DecrementBy(1)
}

func (g *gaugeImpl) DecrementBy(by float64) {
	g.collector.mutex.Lock()
	defer g.collector.mutex.Unlock()

	value := g.collector.get(g.name, map[string]string{})
	g.collector.set(g.name, map[string]string{}, value-by)
}

func (g *gaugeImpl) Set(value float64) {
	g.collector.mutex.Lock()
	defer g.collector.mutex.Unlock()

	g.collector.set(g.name, map[string]string{}, value)
}

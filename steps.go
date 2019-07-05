// Package steps это фреймворк для написания demo сценариев
package steps

// Attempt это название попытки выполнения запроса
type Attempt string

const (
	// StepNotFound означает что шаг с указанным названием не был найден
	StepNotFound Value = "STEP_NOT_FOUND"

	// AttemptNotFound означает что попытка с таким названием не была найдена
	// для данного получателя
	AttemptNotFound Value = "ATTEMPT_NOT_FOUND"
)

// Step это название шага тестового сценария
type Step string

// Value это значение шага
type Value string

// Steps это набор шагов тестового сценария
type Steps map[Step]Value

// NewSteps возвращает новый набор шагов тестового сценария
func NewSteps() Steps {
	return make(Steps)
}

// Set устанавливает значение шага тестового сценария
func (s Steps) Set(step Step, value Value) Steps {
	s[step] = value
	return s
}

// Get возвращает значение шага тестового сценария
func (s Steps) Get(step Step) Value {
	v, ok := s[step]
	if ok {
		return v
	}
	return StepNotFound
}

// Test это тестовый сценарий
type Test struct {
	// data это перечень наборов шагов
	data map[Attempt]Steps
}

// NewTest возвращает новый тестовый сценарий
func NewTest() *Test {
	return &Test{
		data: make(map[Attempt]Steps),
	}
}

// Set устанавливает значение в тестовый сценарий
func (t *Test) Set(attempt Attempt) Steps {
	steps, ok := t.data[attempt]
	if !ok {
		steps = NewSteps()
		t.data[attempt] = steps
	}
	return steps
}

// Get возвращает значение шага по определённой попытке
func (t *Test) Get(attempt Attempt, step Step) Value {
	steps, ok := t.data[attempt]
	if ok {
		return steps.Get(step)
	}
	return AttemptNotFound
}

// Tests это набор тестовых сценариев
type Tests struct {
	data map[interface{}]*Test
}

// NewTests возвращает новый набор тестовых сценариев
func NewTests() *Tests {
	return &Tests{
		data: make(map[interface{}]*Test),
	}
}

// Get возвращает тестовый сценарий по получателю платежа
func (t *Tests) Get(recipient interface{}) *Test {
	test, ok := t.data[recipient]
	if !ok {
		test = NewTest()
		t.data[recipient] = test
	}
	return test
}

// Set устанавливает тестовый сценарий по получателю платежа
func (t *Tests) Set(recipient interface{}, test *Test) {
	t.data[recipient] = test
}

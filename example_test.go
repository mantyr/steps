package steps

import (
	"fmt"
)

// Перечень попыток
// Используются для хранения нескольких наборов шагов в одном тестовом сценарии
const (
	First Attempt = "first"
	Second Attempt = "second"
	Other Attempt = "other"
)

// Перечень названий тестовых сценариев
const (
	Recipient1 TestName = "recipient1"
	Recipient2 TestName = "recipient2"
)

// TestName это название тестового сценария
type TestName string

// Перечень шагов которые можно использовать в тестовом сценарии
const (
	Step1 Step = "step1"
	Step2 Step = "step2"
)

// Перечень значений которые можно использовать в шагах
const (
	StatusOK Value = "OK"
	StatusError Value = "ERROR"
	StatusWait Value = "WAIT"
)

func ExampleSteps() {
	tests := NewTests()

	tests.Get(Recipient1).Set(First).
		Set(Step1, StatusOK).
		Set(Step2, StatusWait)
	tests.Get(Recipient1).Set(Second).
		Set(Step1, StatusError)

	tests.Get(Recipient2).Set(First).
		Set(Step1, StatusError)

	fmt.Println(tests.Get(Recipient1).Get(First, Step1))
	fmt.Println(tests.Get(Recipient1).Get(First, Step2))
	fmt.Println(tests.Get(Recipient1).Get(Second, Step1))
	fmt.Println(tests.Get(Recipient2).Get(First, Step1))
	fmt.Println(tests.Get(Recipient1).Get(Second, Step2))
	fmt.Println(tests.Get(Recipient1).Get(Other, Step2))
	// Output:
	// OK
	// WAIT
	// ERROR
	// ERROR
	// STEP_NOT_FOUND
	// ATTEMPT_NOT_FOUND
}

package steps_test

import (
	"fmt"

	"github.com/mantyr/steps"
)

// Перечень попыток
// Используются для хранения нескольких наборов шагов в одном тестовом сценарии
const (
	First steps.Attempt = "first"
	Second steps.Attempt = "second"
	Other steps.Attempt = "other"
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
	Step1 steps.Step = "step1"
	Step2 steps.Step = "step2"
)

// Перечень значений которые можно использовать в шагах
const (
	StatusOK steps.Value = "OK"
	StatusError steps.Value = "ERROR"
	StatusWait steps.Value = "WAIT"
)

func ExampleNewSteps() {
	tests := steps.NewTests()

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

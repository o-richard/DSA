package stack_test

import (
	"testing"

	"github.com/o-richard/DSA/stack"
)

func TestInit(t *testing.T) {
	tests := map[string]struct {
		size int
		isSuccess bool
	}{
		"Test 1": {
			size: 8,
			isSuccess: true,
		},
		"Test 2": {
		},
		"Test 3": {
			size: -8,
		},
	}

	for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            _, err := stack.Init(test.size)
			if (err == nil) != test.isSuccess {
				t.Errorf("expected %v in stack initialization, actual result is %v", test.isSuccess, !test.isSuccess)
			}
        })
    }
}

func TestStack_Push(t *testing.T) {
	size := 2
	myStack, err := stack.Init(size)
	if err != nil {
		t.Fatalf("the stack was not iniatilized successfully")
	}

	tests := map[string]struct {
		data int
		isSuccess bool
	}{
		"Test 1": {
			data: 1,
			isSuccess: true,
		},
		"Test 2": {
			data: 2,
			isSuccess: true,
		},
		"Test 3": {
			data: 3,
		},
	}

	for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            err := myStack.Push(test.data)
			if (err == nil) != test.isSuccess {
				t.Errorf("expected %v while pushing %v to the stack, actual result is %v", test.isSuccess, test.data, !test.isSuccess)
			}
        })
    }
}

func TestStack_Pop(t *testing.T) {
	size := 3
	myStack, err := stack.Init(size)
	if err != nil {
		t.Fatalf("the stack was not iniatilized successfully")
	}

	data := []int {1,2,3}
	for _, v := range data {
		err := myStack.Push(v)
		if err != nil {
			t.Fatalf("data was not pushed to the stack successfully")
		}
	}

	tests := map[string]struct {
		data int
	}{
		"Test 1": {
			data: 3,
		},
		"Test 2": {
			data: 2,
		},
		"Test 3": {
			data: 1,
		},
	}

	for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            val := myStack.Pop()
			if val != test.data {
				t.Errorf("expected %v while popping from the stack, actual result is %v", test.data, val)
			}
        })
    }
}

func TestStack_Peek(t *testing.T) {
	size := 1
	myStack, err := stack.Init(size)
	if err != nil {
		t.Fatalf("the stack was not iniatilized successfully")
	}

	testData := 3
	err = myStack.Push(testData)
	if err != nil {
		t.Fatalf("data was not pushed to the stack successfully")
	}

	val := myStack.Peek()
	if val != testData {
		t.Errorf("expected %v while peeking from the stack, actual result is %v", testData, val)
	}
}

func TestStack_IsEmpty(t *testing.T) {
	size := 1
	myStack, err := stack.Init(size)
	if err != nil {
		t.Fatalf("the stack was not iniatilized successfully")
	}

	isEmpty := myStack.IsEmpty()
	if isEmpty != true {
		t.Errorf("the stack was expected to be empty")
	}

	testData := 3
	err = myStack.Push(testData)
	if err != nil {
		t.Fatalf("data was not pushed to the stack successfully")
	}

	isEmpty = myStack.IsEmpty()
	if isEmpty != false {
		t.Errorf("the stack was expected not to be empty")
	}
}

func TestStack_IsFull(t *testing.T) {
	size := 1
	myStack, err := stack.Init(size)
	if err != nil {
		t.Fatalf("the stack was not iniatilized successfully")
	}

	isFull := myStack.IsFull()
	if isFull != false {
		t.Errorf("the stack was expected not to be full")
	}

	testData := 3
	err = myStack.Push(testData)
	if err != nil {
		t.Fatalf("data was not pushed to the stack successfully")
	}

	isFull = myStack.IsFull()
	if isFull != true {
		t.Errorf("the stack was expected to be full")
	}
}
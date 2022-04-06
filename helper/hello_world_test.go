package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("before testing")

	m.Run()

	fmt.Println("after testing")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Fauzan")
	if result != "Hello Fauzan" {
		fmt.Println("value is not valid")
	}

	fmt.Println(HelloWorld("Fauzan"))

}

func TestFaileHelloWorld(t *testing.T) {
	result := HelloWorld("Ady")
	if result != "Hello Ady" {
		t.Fail()
	}

	fmt.Println("Test is done")
}

func TestFailNowHelloWorld(t *testing.T) {
	result := HelloWorld("Ady")
	if result != "Hello Ady" {
		t.FailNow()
	}

	fmt.Println("Test is done")
}

func TestErrorHelloWorld(t *testing.T) {
	result := HelloWorld("Ady")
	if result != "Hello Ady" {
		t.Error("value is not valid")
	}

	fmt.Println("Test is done")
}

func TestFatalHelloWorld(t *testing.T) {
	result := HelloWorld("Ady")
	if result != "Hello Ady" {
		t.Fatal()
	}

	fmt.Println("Test is done")
}

func TestAssert(t *testing.T) {
	res := HelloWorld("Fauzan")
	assert.Equal(t, "Hello Fauzan", res, "result is not valid")
	fmt.Println("test assert done")
}

func TestRequeire(t *testing.T) {
	res := HelloWorld("Fauzan")
	require.Equal(t, "Hello Fauzan", res, "result is not valid")
	fmt.Println("test assert done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("this statment just for windows")
	}

	res := HelloWorld("Fauzan")
	require.Equal(t, "Hello Fauzan", res, "this not fauzan")
}

func TestSubTest(t *testing.T) {
	t.Run("Fauzan", func(t *testing.T) {
		res := HelloWorld("Fauzan")
		require.Equal(t, "Hello Fauzan", res, "this not fauzan")
	})
	t.Run("Ady", func(t *testing.T) {
		res := HelloWorld("Ady")
		require.Equal(t, "Hello Ady", res, "this not ady")
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name   string
		req    string
		expect string
	}{
		// TODO: Add test cases.
		{
			name:   "Fauzan",
			req:    "Fauzan",
			expect: "Hello Fauzan",
		},
		{
			name:   "Ady",
			req:    "Ady",
			expect: "Hello Ady",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := HelloWorld(test.req)
			require.Equal(t, test.expect, res)
		})
	}
}

func BenchmarkHelloWorldTable(b *testing.B) {
	// TODO: Initialize
	benchmarks := []struct {
		name string
		req  string
	}{
		{
			name: "Fauzan",
			req:  "Fauzan",
		},
		{
			name: "Ady",
			req:  "Ady",
		},
	}
	for _, benchmarks := range benchmarks {
		b.Run(benchmarks.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmarks.name)
			}
		})
	}

}

func BenchmarkHelloWorldSub(b *testing.B) {
	// TODO: Initialize
	b.Run("Fauzan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// TODO: Your Code Here
			HelloWorld("Fauzan")
		}
	})

	b.Run("Ady", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ady")
		}
	})

}

func BenchmarkHelloWorld(b *testing.B) {
	// TODO: Initialize
	for i := 0; i < b.N; i++ {
		// TODO: Your Code Here
		HelloWorld("Fauzan")
	}
}

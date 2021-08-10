package helpers

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// sebelum unit test
	fmt.Println("sebelum unit test..")

	m.Run()

	// setelah unit test
	fmt.Println("setelah unit test..")
}
func TestBasicFunc(t *testing.T) {
	result := BasicFunc("Depapepe")

	if result != "Hello Depapepe" {
		t.Error("Harusnya Hello Depapepe")
	}

	fmt.Println("TestBasicFunc Done")
}

func TestBasicFuncSayHi(t *testing.T) {
	result := BasicFuncSayHi("Novan")

	if result != "Hi Novan" {
		t.Error("Harusnya Hello Novan")
	}

	fmt.Println("TestBasicFuncSayHi Done")
}

func TestBasicFuncAssert(t *testing.T) {
	result := BasicFunc("Depapepe")
	assert.Equal(t, "Hello Depapepe", result, "Harusnya Hello Depapepe")

	fmt.Println("TestBasicFunc with testify assert Done!")
}

func TestBasicFuncSayHiAssert(t *testing.T) {
	result := BasicFuncSayHi("Novan")
	assert.Equal(t, "Hi Novan", result, "Harusnya Hi Novan")

	fmt.Println("TestBasicFuncSayHi with testifyassert Done!")
}

func TestBasicFuncRequire(t *testing.T) {
	result := BasicFunc("Depapepe")
	require.Equal(t, "Hello Depapepe", result, "Harusnya Hello Depapepe")

	fmt.Println("TestBasicFunc with testify require Done!")
}

func TestBasicFuncSayHiRequire(t *testing.T) {
	result := BasicFuncSayHi("Novan")
	require.Equal(t, "Hi Novan", result, "Harusnya Hi Novan")

	fmt.Println("TestBasicFuncSayHi with testify require Done!")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("hanya bisa berjalan di linux")
	}

	result := BasicFunc("Depapepe")
	assert.Equal(t, "Hello Depapepe", result, "Harusnya Hello Depapepe")

	fmt.Println("TestBasicFunc with testify assert Done!")
}

func TestSubTest(t *testing.T) {
	t.Run("BasicFunc", func(t *testing.T) {
		result := BasicFunc("Depapepe")
		assert.Equal(t, "Hello Depapepe", result, "Harusnya Hello Depapepe")

		fmt.Println("TestBasicFunc with testify assert Done!")
	})

	t.Run("BasicFuncSayHi", func(t *testing.T) {
		result := BasicFuncSayHi("Novan")
		assert.Equal(t, "Hi Novan", result, "Harusnya Hi Novan")

		fmt.Println("TestBasicFuncSayHi with testify assert Done!")
	})
}

func TestBasicFuncTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Depapepe",
			request:  "Depapepe",
			expected: "Hello Depapepe",
		},
		{
			name:     "Novan",
			request:  "Novan",
			expected: "Hello Novan",
		},
		{
			name:     "Kartika",
			request:  "Kartika",
			expected: "Hello Kartika",
		},

		{
			name:     "Wulandhari",
			request:  "Wulandhari",
			expected: "Hello Wulandhari",
		},

		{
			name:     "SedekahCode",
			request:  "SedekahCode",
			expected: "Hello SedekahCode",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := BasicFunc("Depapepe")
			assert.Equal(t, "Hello Depapepe", result)
		})
	}
}

// each benchmark
func BenchmarkBasicFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BasicFunc("Depapepe")
	}
}

// each benchmark
func BenchmarkBasicFuncSayHi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BasicFuncSayHi("Depapepe")
	}
}

// Sub Benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("BenchmarkBasicFun", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BasicFunc("Depapepe")
		}
	})

	b.Run("BenchmarkBasicFuncSayHi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			BasicFuncSayHi("Novan")
		}
	})
}

// table benchmark
func BenchmarkBasicFuncTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Depapepe",
			request: "Depapepe",
		},
		{
			name:    "Novan",
			request: "Novan",
		},
		{
			name:    "Kartika",
			request: "Kartika",
		},

		{
			name:    "Wulandhari",
			request: "Wulandhari",
		},

		{
			name:    "SedekahCodeCommunityCimahiClub",
			request: "SedekahCodeCommunityCimahiClub",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				BasicFunc(benchmark.request)
			}
		})
	}
}

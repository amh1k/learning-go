package sum

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t * testing.T) {
	t.Run("Collection of 5 numbers", func(t* testing.T) {
		numbers := []int{1,2,3,4,5}
		got := Sum(numbers)
		want := 15
		if got != want{
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	t.Run("Collection of any size", func(t* testing.T) {
		numbers := []int{1,2,3}
		got := Sum(numbers)
		want := 6
		if got != want{
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}
func TestSumAll(t* testing.T) {
	got := SumAll([] int {1,2}, []int{0,9});
	want := []int{3,9}
	if !slices.Equal(got,want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t* testing.T) {
	checkSums := func(t testing.TB, got,want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got,want)
	})
	t.Run("Safety sum empty slices", func(t*testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t,got,want)
	})
}


func TestSumReduce(t *testing.T) {
	t.Run("Mulitplication of al elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}
		AssertEqual(t, Reduce([]int {1,2,3}, multiply, 1), 6)
	})
	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func TestBadBank( t *testing.T) {
	var(
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(adil, chris, 25),
		}
	)
	newBalanceFor := func(account Account)float64 {
		return NewBalanceFor(account, transactions).Balance

	}
	AssertEqual(t, newBalanceFor(riya), 200)
	AssertEqual(t, newBalanceFor(chris), 0)
	AssertEqual(t, newBalanceFor(adil), 175)

}

func TestFind(t *testing.T) {
	t.Run("find first even numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})
		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})
}

func AssertEqual[T any](t testing.TB, got  T, want T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

} 
func AssertTrue(t testing.TB, got bool){
	t.Helper()
	if (got != true) {
		t.Fatal("Error")

	}
}
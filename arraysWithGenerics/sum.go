package sum


type Transaction struct {
	From string
	To string
	Sum float64
}
type Account struct {
	Name    string
	Balance float64
}
func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}

// func BalanceFor(transactions []Transaction, name string) float64 {
// 	adjustBalance := func(currentBalance float64, t Transaction) float64 {
// 		if t.From == name {
// 			currentBalance -= t.Sum
// 		}
// 		if t.To == name {
// 			currentBalance += t.Sum
// 		}
// 		return currentBalance
// 	}
// 	return Reduce(transactions, adjustBalance, 0.0)

// }
func Reduce[A,B any](collection [] A, f func (B, A) B, initialValue B) B {
	var result = initialValue
	for _ ,x := range collection {
		result = f(result, x)
	}
	return result
}

func Find[T any](numbers []T, f func(T)(bool) )(T, bool) {
	
	for _, x := range numbers {
		result := f(x)
		if result == true {
			return x, true
		}

	}
	
	var zero T
	return zero, false


}

func Sum(numbers []int) int{
	add := func(acc, x int) int {return acc + x}
	return Reduce(numbers, add, 0)
}
func SumAll(numbersToSum ...[]int )[]int {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)
	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }
	sum := func(acc, x[]int)[]int {
		
		return append(acc, Sum(x))
	}
	return Reduce(numbersToSum, sum, []int{})

}
func SumAllTails(numbersToSum...[]int)[]int {
	sumTail := func(acc, x[]int)[]int {
		if len(x) == 0{
			return append(acc, 0)
		}else {
			tail := x[1 :]
			return append(acc, Sum(tail))


		}
	}
	return Reduce(numbersToSum, sumTail, []int{})
}
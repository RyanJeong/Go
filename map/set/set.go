// 빈 구조체를 값으로 사용해 메모리 낭비를 방지
// struct{}처럼 아무 필드가 없는 구조체는 값이 아닌 자료형
// 해당 구조체에 아무 내용이 없는 상태의 값을 표현하려면?
// struct{}{}
package set

func hasDupeRune(s string) bool {
	runeSet := map[rune]struct{}{}
	for _, r := range s {
		if _, exists := runeSet[r]; exists {

			return true
		}
		runeSet[r] = struct{}{}
	}

	return false
}

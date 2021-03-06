//TODO: 테스트 후 적절히 수정할 것

* 반복자 예제

** 반복자 벤치마크

컨테이너를 순회할 수 있는 객체를 반복자라고 합니다. 반복자 패턴을
이용하면 컨테이너의 종류에 상관없이 자료를 순회할 수 있습니다.

Go 언어에서 정해둔 반복자 패턴은 따로 없습니다. 그러나 몇 가지
방식들이 있을 수 있는데 성능상에 어떤 차이가 있는지 벤치마크 해
보았습니다.

아래와 같은 명령을 내리면 벤치마크를 할 수 있습니다.

: go test -bench . github.com/jaeyeom/gogo/examples/iterator

빠른 것이 위로 오도록 정렬했습니다. 윈도우10 머신에서 실행한
것입니다. 뒤에 -2는 프로세서 2개를 이용했다는 뜻입니다.

|                                       |     |           <r> |
| BenchmarkCallback-2                   | 500 | 3324446 ns/op |
| BenchmarkFunc-2                       | 500 | 3302264 ns/op |
| BenchmarkInterface-2                  | 500 | 3331370 ns/op |
| BenchmarkBufferedChannel-2            | 300 | 4551638 ns/op |
| BenchmarkChannel-2                    | 300 | 5676556 ns/op |
| BenchmarkBufferedChannelWithContext-2 | 200 | 6507083 ns/op |
| BenchmarkChannelWithContext-2         | 200 | 7594841 ns/op |

iterator_test.go 파일에는 2개의 상수가 있습니다. 하나는 반복할 횟수를
정하는 size이고, 다른 하나는 하나의 컨테이너 값을 뽑아낼 때 시간을
지체시키기 위하여 몇 번 반복할지를 정하는 eachTask 입니다.

이 eachTask가 0이 되면 거의 iterator의 오버헤드만 측정이 되게 됩니다.

|                                       |       |           <r> |
| BenchmarkCallback-2                   | 50000 |   34381 ns/op |
| BenchmarkFunc-2                       | 30000 |   54028 ns/op |
| BenchmarkInterface-2                  | 30000 |   56265 ns/op |
| BenchmarkBufferedChannel-2            |  2000 | 1102318 ns/op |
| BenchmarkChannel-2                    |  1000 | 2433695 ns/op |
| BenchmarkBufferedChannelWithContext-2 |  1000 | 2673675 ns/op |
| BenchmarkChannelWithContext-2         |   300 | 4187138 ns/op |


이 경우에는 콜백 구현과 컨텍스트를 이용한 채널 구현이 약 100배 이상
차이가 납니다.

재미있는 점은 프로세서 개수를 1개로 만들면 오히려 채널 성능이
올라간다는 것입니다. 더욱 놀라운 점은 버퍼 있는 채널의 성능 향상이
프로세서 개수 1개일 때 더 크다는 점입니다.

아래는 두 표는 리눅스에서 eachTask를 0으로 두고 측정한 것입니다. 먼저
코어를 2개 사용하여 측정해 보았습니다.

|                                       |       |           <r> |
| BenchmarkCallback-2                   | 50000 |   35658 ns/op |
| BenchmarkFunc-2                       | 30000 |   47043 ns/op |
| BenchmarkInterface-2                  | 30000 |   57108 ns/op |
| BenchmarkBufferedChannel-2            |  1000 | 1640416 ns/op |
| BenchmarkBufferedChannelWithContext-2 |   500 | 4009550 ns/op |
| BenchmarkChannel-2                    |   500 | 4101341 ns/op |
| BenchmarkChannelWithContext-2         |   200 | 7390680 ns/op |


이제 싱글 코어만 사용하여 측정했습니다. 채널 성능이 모두 2배로
향상되었음을 알 수 있습니다.

|                                     |       |           <r> |
| BenchmarkCallback                   | 50000 |   35480 ns/op |
| BenchmarkFunc                       | 30000 |   47029 ns/op |
| BenchmarkInterface                  | 30000 |   57291 ns/op |
| BenchmarkBufferedChannel            |  2000 |  866433 ns/op |
| BenchmarkBufferedChannelWithContext |  1000 | 1932076 ns/op |
| BenchmarkChannel                    |  1000 | 2077974 ns/op |
| BenchmarkChannelWithContext         |   500 | 3506129 ns/op |

이제 사용법을 살펴보겠습니다. 밑줄 문자로 시작하는 줄은 별 의미가 없고
단지 우변에 있는 변수를 사용하는 척 하고 버리기 위한 것입니다.

*** 콜백

사용법을 보면, 콜백은 for 반복문을 사용하지 않습니다. 콜백 함수가
참/거짓 값으로 계속 진행하고 싶은지 중간에 중단하고 싶은지를 정할 수
있습니다. 일반적으로 콜백은 보기 어려운 경우가 많지만 함수 리터럴을
바로 이용할 수 있는 장점 때문에, 큰 문제는 없어 보입니다.

: iter(func(num int) bool {
: 	_ = num
: 	return true  // Continue
: })

*** 채널

채널은 for range를 이용합니다. 이용 방법은 굉장히 자연스럽습니다만,
중간에 중지할 수 있는 방법이 없습니다. 중지할 수 있는 방법이 있으려면
컨텍스트가 있는 채널을 이용해야 합니다.

: for num := range iter() {
: 	_ = num
: }

*** 컨텍스트가 있는 채널

마찬가지로 for range를 이용합니다. 중간에 중지하려면 cancel()을 부르면
됩니다.

: ctx, cancel := context.WithCancel(context.Background())
: for num := range iter() {
: 	_ = num
: 	// cancel()
: }

*** 버퍼가 있는 채널

위의 채널 구현에서 버퍼 있는 채널을 사용한 경우인데, 사용법은
동일합니다.

*** 함수

함수를 연속적으로 불러서 계속 값을 꺼내옵니다. 두 번째 결과값을
이용하여 계속 호출을 할지 말지를 알려줍니다. 인터페이스와 비슷하지만,
사용법이 직관적이지 않은 단점이 있습니다.

사용하는 쪽에서 중간에 중단하고 싶으면 대부분의 경우에 함수를 호출하지
않으면 됩니다.

: itr := iter()
: for num, ok := itr(); ok; num, ok = itr() {
: 	_ = num
: }

그러나 어떤 자원에 접근하고 있는 경우에 제대로 해제되지 않을 수 있기
때문에 함수에 더이상 값을 받고 싶지 않다는 값을 하나 함수에 넘겨주는
식으로 응수할 수 있습니다. 별로 깔끔하지는 않은 인터페이스라서 별로
내키지는 않습니다. 아래 코드에는 그나마 가독성을 주기 위하여 NEXT와
STOP과 같은 상수를 이용했습니다.

: itr := iter()
: for num, ok := itr(NEXT); ok; num, ok = itr(NEXT) {
: 	_ = num
: 	// itr(STOP); break
: }

*** 인터페이스

Java 스타일과 흡사합니다. 사용하는 쪽에서 중간에 중단하고 싶으면
더이상 Next()를 호출하지 않으면 됩니다. 그러나 자원의 해제가 필요한
경우라면, Close() 같은 메서드를 추가해야 할 수 있습니다.

: itr := iterator(0)
: for num := itr.Next(); !itr.Done(); num = itr.Next() {
: 	_ = num
: 	// itr.Close(); break
: }


출처: https://github.com/jaeyeom/gogo

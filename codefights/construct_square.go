import "math"
import "sort"

func isPerfectSquareInt(candidate int) bool {
    return isPerfectSquare(float64(1.0 * candidate))
}

func isPerfectSquare(candidate float64) bool {
    return math.Sqrt(candidate) == math.Floor(math.Sqrt(candidate))
}

func reverseArray(arr []int) []int {
    for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }

    return arr
}

func generateCandidateNumber(frequencies []int, digits []int) int {
    var num int = 0

    for i := 0; i < len(digits); i++ {
        var digit int = digits[i]
        var frequency int = frequencies[i]
        for (frequency > 0) {
            num = num * 10
            num = num + digit
            frequency = frequency - 1
        }
    }
    return num
}

func sliceEquals(sliceA []int, sliceB []int) bool {
    var sliceALen int = len(sliceA)
    var sliceBLen int = len(sliceB)

    if (sliceALen != sliceBLen) {
        return false
    }

    for i := 0; i < sliceALen; i++ {
        if (sliceA[i] != sliceB[i]) {
            return false
        }
    }

    return true
}

func findNextBiggest(digits []int) int {
    var counts [10]int
    for i := 0; i < len(digits); i++ {
        counts[digits[i]] = 1
    }

    for i := 9; i >= 0; i-- {
        if counts[i] == 0 {
            return i
        }
    }

    return 0
}

func findNextSmallest(digits []int) int {
    var counts [10]int
    for i := 0; i < len(digits); i++ {
        counts[digits[i]] = 1
    }

    for i := 0; i <= 9; i++ {
        if counts[i] == 0 {
            return i
        }
    }

    return 0
}

func allUniqueDigits(num int) bool {
    var current int = num
    var digits map[int]bool = make(map[int]bool)
    for (current > 0) {
        var nextDigit int = current % 10
        if _, ok := digits[nextDigit]; ok {
            return false
        } else {
            digits[nextDigit] = true
            current = current / 10
        }
    }

    return true
}

func numToDigits(num int) []int {
    var current int = num
    var digits []int
    for (current > 0) {
        var nextDigit int = current % 10
        digits = append(digits, nextDigit)
        current = current / 10
    }
    return reverseArray(digits)
}

func digitsToNum(digits []int) int {
    var num int = 0
    var digitsLen int = len(digits)
    for i := 0; i < digitsLen; i++ {
        num = num * 10
        num = num + digits[i]
    }
    return num
}

func getNextDigits(digits []int) []int {
    var num int = digitsToNum(digits) - 1
    var condition bool = true
    for (condition) {
        if allUniqueDigits(num) {
            condition = false
        } else {
            num = num - 1
        }
    }

    return numToDigits(num)
}

func intInSlice(cand int, items []int) bool {
    var itemsLen int = len(items)
    for i := 0; i < itemsLen; i++ {
        if items[i] == cand {
            return true
        }
    }

    return false
}

func genUniqueNums(current []int) []int {
    var uniqueNumbers []int;
    var currentLen int = len(current)
    for i := 0; i < currentLen; i++ {
        if (!intInSlice(current[i], uniqueNumbers)) {
            uniqueNumbers = append(uniqueNumbers, current[i])
        }
    }

    return uniqueNumbers
}

func getNextPermutationDigitCandidates(items []int, idx int) []int {
    var candidates []int
    var item int = items[idx]
    var itemsLen int = len(items)

    for i := idx; i < itemsLen; i++ {
        if (items[i] < item) {
            if (!intInSlice(items[i], candidates)) {
                candidates = append(candidates, items[i])
            }
        }
    }

    sort.Ints(candidates)
    candidates = reverseArray(candidates)
    return candidates
}

func getLastDecreasingIdx(current []int) int {
    var currentLen = len(current)
    var lastIdx = -1
    for i := 0; i < currentLen - 1; i++ {
        if (current[i] > current[i + 1]) {
            lastIdx = i
        }
    }

    return lastIdx
}

func generateRemainingDigitsToAssign(current []int, idx int, itemToSkipOnce int) []int {
    var remainingDigitsToAssign []int
    var itemToSkipOnceFound bool = false
    var currentLen int = len(current)
    remainingDigitsToAssign = append(remainingDigitsToAssign, current[idx])


    var currentIdx int = idx + 1
    for (currentIdx < currentLen) {
        var currentItem int = current[currentIdx]
        if (currentItem == itemToSkipOnce) {
            if (!itemToSkipOnceFound) {
                itemToSkipOnceFound = true
            } else {
                remainingDigitsToAssign = append(remainingDigitsToAssign, currentItem)
            }
        } else {
            remainingDigitsToAssign = append( remainingDigitsToAssign, currentItem)
        }
        currentIdx = currentIdx + 1
    }

    sort.Ints(remainingDigitsToAssign)
    remainingDigitsToAssign = reverseArray(remainingDigitsToAssign)

    return remainingDigitsToAssign
}

func genNextPermutation(current []int, sortedUnique []int) []int {
    var currentLen int = len(current)
    var nextPermutation = make([]int, currentLen)
    var lastDecreasingIdx = getLastDecreasingIdx(current)

    copy(nextPermutation, current)

    if (lastDecreasingIdx > -1) {
        var candidates []int = getNextPermutationDigitCandidates(current, lastDecreasingIdx)
        if (len(candidates) > 0) {
            var candidate int = candidates[0]
            var remainingDigitsToAssign []int = generateRemainingDigitsToAssign(current, lastDecreasingIdx, candidate)
            nextPermutation[lastDecreasingIdx] = candidate
            lastDecreasingIdx = lastDecreasingIdx + 1

            var remainingDigitsIdx = 0

            for (lastDecreasingIdx < currentLen) {
                nextPermutation[lastDecreasingIdx] = remainingDigitsToAssign[remainingDigitsIdx]
                lastDecreasingIdx = lastDecreasingIdx + 1
                remainingDigitsIdx = remainingDigitsIdx + 1
            }
        }
    }
    return nextPermutation
}

func constructSquare(s string) int {
    var charCounts [26]int
    var idx int
    var largestSquare int = -1
    var charCountsCondensed []int
    var charCountsCondensedCopy []int
    var charCountsCondensedTotalLen int
    var charCountsCondensedLen int
    var digits []int
    var uniqueCharCounts []int

    for _, nextRune := range s {
        idx = int(nextRune - 'a')
        charCounts[idx] = charCounts[idx] + 1
    }

    charCountsCondensedTotalLen = 0
    for idx = 0; idx < 26; idx++ {
        if charCounts[idx] > 0 {
            charCountsCondensed = append(charCountsCondensed, charCounts[idx])
            charCountsCondensedCopy = append(charCountsCondensedCopy, charCounts[idx])
            charCountsCondensedTotalLen = charCountsCondensedTotalLen + charCounts[idx]
            if (!intInSlice(charCounts[idx], uniqueCharCounts)) {
                uniqueCharCounts = append(uniqueCharCounts, charCounts[idx])
            }
        }
    }

    sort.Ints(charCountsCondensed)
    sort.Ints(uniqueCharCounts)
    charCountsCondensed = reverseArray(charCountsCondensed)
    uniqueCharCounts = reverseArray(uniqueCharCounts)

    if charCountsCondensedTotalLen >= 10 {
        return -1
    }

    var largest int = 9
    charCountsCondensedLen = len(charCountsCondensed)
    for idx = 0; idx < charCountsCondensedLen; idx++ {
        digits = append(digits, largest)
        largest = largest - 1
    }

    //fmt.Println(charCountsCondensedCopy)

    var condition bool = true
    for (condition) {
        var innerCondition bool = true
        copy(charCountsCondensedCopy, charCountsCondensed)
        for (innerCondition) {
            charCountsCondensedCopy = genNextPermutation(charCountsCondensedCopy, uniqueCharCounts)

            var candidateNumber = generateCandidateNumber(charCountsCondensedCopy, digits)


            if (isPerfectSquareInt(candidateNumber)) {
                return candidateNumber
            }

            if (getLastDecreasingIdx(charCountsCondensedCopy) == -1) {
                innerCondition = false
            }
        }

        digits = getNextDigits(digits)
        if (len(digits) != charCountsCondensedLen) {
            condition = false
        }
    }

    //fmt.Println(charCountsCondensedCopy)
    //fmt.Println(digits)

    return largestSquare
}

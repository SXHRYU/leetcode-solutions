package leetcode

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	result, wordMap := make([][]string, 0), make(map[string]bool)
	for _, w := range wordList {
		wordMap[w] = true
	}
	if !wordMap[endWord] {
		return result
	}

	queue := make([][]string, 0)
	queue = append(queue, []string{beginWord})

	queueLen := 1

	levelMap := make(map[string]bool)
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		lastWord := path[len(path)-1]
		for i := 0; i < len(lastWord); i++ {
			for c := 'a'; c <= 'z'; c++ {
				nextWord := lastWord[:i] + string(c) + lastWord[i+1:]
				if nextWord == endWord {
					path = append(path, endWord)
					result = append(result, path)
					continue
				}
				if wordMap[nextWord] {

					levelMap[nextWord] = true
					newPath := make([]string, len(path))
					copy(newPath, path)
					newPath = append(newPath, nextWord)
					queue = append(queue, newPath)
				}
			}
		}
		queueLen--

		if queueLen == 0 {
			if len(result) > 0 {
				return result
			}
			for k := range levelMap {
				delete(wordMap, k)
			}

			levelMap = make(map[string]bool)
			queueLen = len(queue)
		}
	}
	return result
}

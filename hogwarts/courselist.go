//go:build !solution

package hogwarts

/*func GetCourseList(prereqs map[string][]string) []string {
	basicCouses := make(map[string]interface{})
	result := []string{}
	for k, v := range prereqs {
		if len(v) == 0 {
			result = append(result, k)
			basicCouses[k] = nil
		}
	}
	fmt.Println(result)

	var OKFlag bool
	for k, v := range prereqs {
		OKFlag = true
		if len(v) != 0 {
			for _, course := range v {
				if _, ok := basicCouses[course]; !ok {
					OKFlag = false
					break
				}
			}
			if OKFlag {
				result = append(result, k)
			}

		}
	}
	fmt.Println(result)
	return result
}*/

func GetCourseList(prerequisites map[string][]string) []string {
	// Создаем граф и степени входа
	graph := make(map[string][]string)
	inDegree := make(map[string]int)

	// Заполняем граф и степени входа
	for course, prereqs := range prerequisites {
		inDegree[course] = 0 // Инициализируем степень входа для каждого курса
		for _, prereq := range prereqs {
			graph[prereq] = append(graph[prereq], course)
			inDegree[course]++ // Увеличиваем степень входа для курса
		}
	}

	// Находим курсы без пререквизитов
	var zeroInDegree []string
	for course := range inDegree {
		if inDegree[course] == 0 {
			zeroInDegree = append(zeroInDegree, course)
		}
	}

	var courseOrder []string

	// Обработка очереди
	for len(zeroInDegree) > 0 {
		course := zeroInDegree[0]
		zeroInDegree = zeroInDegree[1:]
		courseOrder = append(courseOrder, course)

		// Уменьшаем степень входа для зависимых курсов
		for _, neighbor := range graph[course] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				zeroInDegree = append(zeroInDegree, neighbor)
			}
		}
	}

	// Проверка на наличие циклов
	if len(courseOrder) != len(prerequisites) {
		panic("Cource cycle")
	}

	return courseOrder
}

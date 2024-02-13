package handlers

import (
	"fmt"
	"math"
	"sort"
	"todo_planning/model"
)

type Developer struct {
	ID     string
	Cap    int // 1 Hour work capacity
	Rating int // Work capacity rating
}

func calculateWorkload(tasks []model.Provider1Model) (totalDifficulty float64, totalDuration int) {
	for _, task := range tasks {
		totalDifficulty += float64(task.Zorluk)
		// totalDuration += task.Duration // Eğer Duration bilgisi alınacaksa buraya eklenmeli
	}
	return totalDifficulty, totalDuration
}

func distributeTasks(tasks []model.Provider1Model, developers []Developer) map[string][]model.Provider1Model {
	distributedTasks := make(map[string][]model.Provider1Model)
	devWork := make(map[string]float64)

	// Developer'ların çalışma kapasitesini tutan bir map oluşturuyoruz
	for _, dev := range developers {
		devWork[dev.ID] = 0
	}

	// Görevleri zorluklarına göre sıralıyoruz
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Zorluk > tasks[j].Zorluk
	})

	// Her bir görevi en uygun geliştiriciye atıyoruz
	for _, task := range tasks {
		minWork := 1e9 // Yüksek bir başlangıç değeri
		var selectedDev string

		// Her bir geliştirici için iş yükünü hesaplayıp en az olanı seçiyoruz
		for _, dev := range developers {
			work := devWork[dev.ID] + float64(task.Zorluk)/float64(dev.Rating)
			if work < minWork && work <= float64(dev.Cap) {
				minWork = work
				selectedDev = dev.ID
			}
		}

		// Eğer bir geliştiriciye atanabilen bir görev varsa, ona atıyoruz ve iş yükünü güncelliyoruz
		if selectedDev != "" {
			distributedTasks[selectedDev] = append(distributedTasks[selectedDev], task)
			devWork[selectedDev] += float64(task.Zorluk) / float64(developers[0].Rating) // Developer'ların rating'lerini referans alıyoruz
		}
	}

	return distributedTasks
}

// TODO Tasklar zorluk derecesine ve developer çözme kapasitesine göre dağıtılacak
func WeeklyWorkPlan(provider1Tasks []model.Provider1Model, developers []Developer, weeklyHours int) {
	_, totalDuration1 := calculateWorkload(provider1Tasks)
	// totalDifficulty2, totalDuration2 := calculateWorkload(provider2Tasks)
	// totalDifficulty := totalDifficulty1 + totalDifficulty2
	// totalDuration := totalDuration1 + totalDuration2

	numWeeks := int(math.Ceil(float64(totalDuration1) / float64(weeklyHours)))
	fmt.Printf("Minimum %d haftada iş bitirilecek.\n", numWeeks)

	distributedTasks := distributeTasks(provider1Tasks, developers)

	for devID, tasks := range distributedTasks {
		fmt.Printf("Developer %s: %d görevi üstlendi:\n", devID, len(tasks))
		for _, task := range tasks {
			fmt.Printf("\t%s\n", task.Id)
		}
	}
}

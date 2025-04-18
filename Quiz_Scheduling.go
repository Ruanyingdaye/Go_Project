package main

import (
	"log"
	"sort"
)

// for output
type TimeSlice struct {
	indexSlice      []int
	remainTimeSlice []int
}

// input given
type TaskTimeSlice struct {
	index     int
	task_time int
}

const exectime = 5

func ExecQuestion(task_time_slice []*TaskTimeSlice) []*TimeSlice {
	ts := []*TimeSlice{}
	last_time_slice := 0
	for _, task := range task_time_slice {
		ts, last_time_slice = DoCheck(task.index, task.task_time, last_time_slice, ts)
	}
	return ts
}

func DoCheck(index int, indexNeedTime int, last_time_slice int, ts []*TimeSlice) ([]*TimeSlice, int) {
	//return the last_time_slice
	//last_time_slice: used in this condition: if last time slice didn't consume, then need consume last_time_slice first
	//indexNeedTime: index need timeSlice
	if len(ts) == 0 && last_time_slice != 0 {
		log.Panic("Invalid condition")
	}
	var lastTS, nowTS *TimeSlice
	if len(ts) != 0 {
		lastTS = ts[len(ts)-1]
		nowTS = ts[len(ts)-1]
	}

	for indexNeedTime != 0 {
		if last_time_slice != 0 {
			indexNeedTime = indexNeedTime - last_time_slice
			last_time_slice = 0
			lastTS.indexSlice = append(lastTS.indexSlice, index)
			if indexNeedTime > 0 {
				lastTS.remainTimeSlice = append(lastTS.remainTimeSlice, indexNeedTime)
			} else {
				lastTS.remainTimeSlice = append(lastTS.remainTimeSlice, 0)
				return ts, -indexNeedTime
			}
		} else {
			// if last_time_slice == 0, then create new nowTS
			nowTS = &TimeSlice{}
			ts = append(ts, nowTS)
			indexNeedTime = indexNeedTime - exectime
			nowTS.indexSlice = append(nowTS.indexSlice, index)
			if indexNeedTime > 0 {
				nowTS.remainTimeSlice = append(nowTS.remainTimeSlice, indexNeedTime)
			} else {
				nowTS.remainTimeSlice = append(nowTS.remainTimeSlice, 0)
				return ts, -indexNeedTime
			}
		}
	}

	//indexNeedTime = 0, invalid
	return ts, last_time_slice
}

func CheckExecQuestion() {
	tts := []*TaskTimeSlice{
		&TaskTimeSlice{
			index:     0,
			task_time: 2,
		},
		&TaskTimeSlice{
			index:     1,
			task_time: 4,
		},
		&TaskTimeSlice{
			index:     2,
			task_time: 100,
		},
		&TaskTimeSlice{
			index:     3,
			task_time: 2,
		},
		&TaskTimeSlice{
			index:     4,
			task_time: 1,
		},
		&TaskTimeSlice{
			index:     5,
			task_time: 8,
		},
	}

	resTS := ExecQuestion(tts)
	for i, item := range resTS {
		log.Printf("time: %d, index: %v, remain_time_slice: %v", i, item.indexSlice, item.remainTimeSlice)
	}
}

func ExecAdvancedQuestion(task_time_slice []*TaskTimeSlice) []*TimeSlice {
	// we use the SJF, it will change the index after sort
	sort.Slice(task_time_slice, func(i, j int) bool {
		return task_time_slice[i].task_time < task_time_slice[j].task_time
	})
	// for _, item := range task_time_slice {
	// 	fmt.Println(item.index, item.task_time)
	// }
	ts := []*TimeSlice{}
	last_time_slice := 0
	for _, task := range task_time_slice {
		ts, last_time_slice = DoCheck(task.index, task.task_time, last_time_slice, ts)
	}
	return ts
}

func CheckExecAdvancedQuestion() {
	//SJF
	tts := []*TaskTimeSlice{
		&TaskTimeSlice{
			index:     0,
			task_time: 2,
		},
		&TaskTimeSlice{
			index:     1,
			task_time: 4,
		},
		&TaskTimeSlice{
			index:     2,
			task_time: 100,
		},
		&TaskTimeSlice{
			index:     3,
			task_time: 2,
		},
		&TaskTimeSlice{
			index:     4,
			task_time: 1,
		},
		&TaskTimeSlice{
			index:     5,
			task_time: 8,
		},
	}

	resTS := ExecAdvancedQuestion(tts)
	for i, item := range resTS {
		log.Printf("time: %d, index: %v, remain_time_slice: %v", i, item.indexSlice, item.remainTimeSlice)
	}
}

func main() {
	CheckExecQuestion()
	CheckExecAdvancedQuestion()
}

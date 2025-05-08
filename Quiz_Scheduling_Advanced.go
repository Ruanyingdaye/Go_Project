package main

import (
	"log"
	"time"
	// 没有container/heap包，使用container/heap包
)

/*
1. 还需要一个update才能实现priorityQueue？
2. 协程安全，是否需要添加锁/Preemption？
如果实现Preemption，那么可以在EnqueueTask的时候添加一个事件驱动机制，给原来的Task添加一个中断，保留执行完该时间片后的信息？(假设每一个时间片很长)
或者如果时间片足够短，那就等下一个时间片到了之后，会自动触发重新入队，则不需要显式的进行Preemption了。
*/
const exectime = 5

// for output
type TimeSlice struct {
	indexSlice      []int
	remainTimeSlice []int
}

// input given， also producer/consumer task
type TaskTimeSlice struct {
	index     int
	task_time int
}

/*
1. 有一个消费者始终拿取container/heap中pop出来的数据，相当于之前的for循环，每次循环都会执行DoCheck的工作。
   协程会每隔5s中从这个heap中Pop一个元素
2. 有一个生产者（每次）都会开一个协程，往heap容器中添加元素
3. 1和2两个是异步并行协程进行执行，作为生产者消费者队列，将结果输出到PriorityQueue中。然后在每次PriorityQueue进行时间片消耗的时候
同时也进行channel的执行操作~
4. 还会有一个协程
5. 输入是小顶堆，输出就是按顺序输出

*/
// Multipley-level Queue
//**********************************construct the container/heap**********************
type PriorityQueue []*TaskTimeSlice

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].task_time < pq[j].task_time
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

//实现pq的push和pop需要指针

func (pq *PriorityQueue) Push(x interface{}) {
	newTask := x.(*TaskTimeSlice)
	for i, item := range *pq {
		if newTask.task_time <= item.task_time {
			//insert newTask into the i-1~ i
			*pq = append(*pq, nil)       // 先扩容一个元素
			copy((*pq)[i+1:], (*pq)[i:]) // 将i到len(pq)的元素向后移动一位
			(*pq)[i] = newTask           // 将新元素放入到i的位置
		}
	}
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	old[n-1] = nil //避免内存泄漏
	// 这里需要注意的是，pop出来的元素是一个指针，不能直接使用pq[len(pq)-1]，因为会导致内存泄漏
	return x
}

// 返回any，强调泛化适配
// 涉及到具体的要用*
// 作用就是这个queue会一遍进行时间片消耗，一边将消耗的时间片提取出来，并重新入队，期间另一个携程会每个x秒钟也进行同样的入队操作

func (pq PriorityQueue) ExecuteTask(leftTimeSlice int) *TimeSlice {

	// 执行一个时间片，然后一定会return一个time
	if pq.Len() == 0 {
		log.Panic("Invalid condition")
	}
	i := 0
	time.Sleep(5 * time.Second)

	// 需要判断该task是否执行完，如果没执行完重新入队, 执行EnqueueTask
	nowts := &TimeSlice{}
	for leftTimeSlice > 0 {
		task := pq[i]
		if task.task_time == leftTimeSlice {
			leftTimeSlice = 0
			nowts.indexSlice = append(nowts.indexSlice, task.index)
			nowts.remainTimeSlice = append(nowts.remainTimeSlice, 0)
			return nowts
		} else if task.task_time > leftTimeSlice {
			//表示执行不完，剩下的重新入队
			//已经执行的进行操作
			leftTimeSlice = 0
			nowts.indexSlice = append(nowts.indexSlice, task.index)
			nowts.remainTimeSlice = append(nowts.remainTimeSlice, 0)
			//重新入队
			leftTimeSlice = task.task_time - leftTimeSlice
			newTaskSlice := &TaskTimeSlice{
				index:     task.index,
				task_time: leftTimeSlice,
			}
			pq.Push(newTaskSlice)
			return nowts
		} else {
			// leftTimeSlice还有富裕，可以执行下一个task
			// 将该task入队，task->i ++
			leftTimeSlice = leftTimeSlice - task.task_time
			nowts.indexSlice = append(nowts.indexSlice, task.index)
			nowts.remainTimeSlice = append(nowts.remainTimeSlice, leftTimeSlice)
			i++
		}
	}
	return nowts
}

//************************************************************************************

func ExecAdvancedQuestion(task_time_slice []*TaskTimeSlice) []*TimeSlice {

	//启动三个协程，生产者消费者以及ExecuteTask
	//消费者的time.Sleep是一个random(1~10)
	//Execute的协程每5s执行一个ExecuteTask()
	//最后执行x个数后结束，去看最终结果
	return []*TimeSlice{}
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
	CheckExecAdvancedQuestion()
}

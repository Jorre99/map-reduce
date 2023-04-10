package mr

import "container/list"

type Task struct {
	Phase         Phase
	InputFilepath string
	TaskId        int
	WorkerId      int
}

type Tasks struct {
	Queue    map[Progression]*list.List
	Node     map[int]*list.Element
	State    map[int]Progression
	Capacity int
}

func allocateTasks(Capacity int) *Tasks {
	tasks := &Tasks{
		Queue:    make(map[Progression]*list.List, TotalState),
		Node:     make(map[int]*list.Element),
		State:    make(map[int]Progression),
		Capacity: Capacity,
	}

	tasks.Queue[Idle] = list.New()
	tasks.Queue[InProgress] = list.New()
	tasks.Queue[Completed] = list.New()

	return tasks
}

func GenerateTasks(files []string, nReduce int) (mapTasks, reduceTasks *Tasks) {
	nMap := len(files)
	mapTasks = allocateTasks(nMap)
	reduceTasks = allocateTasks(nReduce)

	for i, file := range files {
		task := &Task{Phase: MapTask, InputFilepath: file, TaskId: i}
		mapTasks.Node[task.TaskId] = mapTasks.Queue[Idle].PushBack(task)
		mapTasks.State[task.TaskId] = Idle
	}

	for i := 0; i < nReduce; i++ {
		task := &Task{Phase: ReduceTask, TaskId: i}
		reduceTasks.Node[task.TaskId] = reduceTasks.Queue[Idle].PushBack(task)
		reduceTasks.State[task.TaskId] = Idle
	}

	return
}

func (tasks *Tasks) GetIdleTask() *Task {
	idleTasks := tasks.Queue[Idle]
	if idleTasks.Len() > 0 {
		taskId := idleTasks.Front().Value.(*Task).TaskId
		tasks.UpdateTaskState(taskId, InProgress)
		task, _ := tasks.findTask(taskId)
		return task
	}
	return &Task{Phase: VoidTask}
}

func (tasks *Tasks) findTask(taskId int) (*Task, bool) {
	if val, ok := tasks.Node[taskId]; ok {
		return val.Value.(*Task), ok
	}
	return nil, false
}

func (tasks *Tasks) UpdateTaskState(taskId int, newState Progression) {
	if _, ok := tasks.findTask(taskId); !ok {
		return
	}
	state := tasks.State[taskId]
	task := tasks.Queue[state].Remove(tasks.Node[taskId]).(*Task)
	node := tasks.Queue[newState].PushBack(task)
	tasks.Node[taskId] = node
	tasks.State[taskId] = newState
}

func (tasks *Tasks) GetWorker(taskId int) (int, bool) {
	if task, ok := tasks.findTask(taskId); ok {
		return task.WorkerId, ok
	}
	return 0, false
}

func (tasks *Tasks) SetWorker(taskId int, workerId int) {
	if task, ok := tasks.findTask(taskId); !ok {
		task.WorkerId = workerId
	}
}

func (tasks *Tasks) Done() bool {
	return tasks.Queue[Completed].Len() == tasks.Capacity
}

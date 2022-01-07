package mr

import "log"
import "net"
import "os"
import "net/rpc"
import "net/http"
import "sync"
import "time"

type MasterTaskStatus int

//3 status for each task(Map & Reduce)
const(
	Idle MasterTaskStatus = iota
	In_Progess
	Completed
)

//States for Task and Master
type State int

const(
	Map State = iota
	Reduce
	Exit
	Wait
)

type Task struct {
	Input string
	Output string
	TaskState State
	TaskNumber int
	NReducer int
	Intermediates []string
}

type Master struct {
	// Your definitions here.
	TaskQueue chan *Task
	TaskMeta map[int]*MasterTask
	MasterPhase States
	NReduce int
	InputFiles []string
	Intermediates [][]string
}

type MasterTask struct {
	TaskStatus MasterTaskStatus
	StartTime time.Time
	TaskReference *Task
}

var mu sync.Mutex

// Your code here -- RPC handlers for the worker to call.

//
// an example RPC handler.
//
// the RPC argument and reply types are defined in rpc.go.
//
func (m *Master) Example(args *ExampleArgs, reply *ExampleReply) error {
	reply.Y = args.X + 1
	return nil
}


//
// start a thread that listens for RPCs from worker.go
//
func (m *Master) server() {
	rpc.Register(m)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := masterSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

//
// main/mrmaster.go calls Done() periodically to find out
// if the entire job has finished.
//
func (m *Master) Done() bool {
	ret := false

	// Your code here.


	return ret
}

//
// create a Master.
// main/mrmaster.go calls this function.
// nReduce is the number of reduce tasks to use.
//
func MakeMaster(files []string, nReduce int) *Master {
	m := Master{}

	// Your code here.


	m.server()
	return &m
}

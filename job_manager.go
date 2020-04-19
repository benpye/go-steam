package steam

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/benpye/go-steam/protocol"
	"github.com/benpye/go-steam/protocol/steamlang"
)

// TODO: Make this timeout configurable
const jobTimeoutDuration = time.Second * 10

// ErrJobTimedOut is returned when the job times out due to no response.
var ErrJobTimedOut = errors.New("job has timed out")

// ErrJobFailed is returned when the job fails remotely.
var ErrJobFailed = errors.New("job has failed")

// ErrJobCancelled is returned when the job must be cancelled - ie. disconnect.
var ErrJobCancelled = errors.New("job has been cancelled")

// AsyncJob represents an asynchronous Steam job.
type AsyncJob struct {
	completed uint32
	jm        *JobManager
	JobID     protocol.JobID
	timer     *time.Timer
	channel   chan interface{}
}

func newJob(jm *JobManager) *AsyncJob {
	job := &AsyncJob{
		jm:      jm,
		JobID:   jm.NextJobID(),
		timer:   time.NewTimer(jobTimeoutDuration),
		channel: make(chan interface{}, 1),
	}

	go job.timerThread()

	return job
}

func (j *AsyncJob) timerThread() {
	<-j.timer.C
	j.complete(ErrJobTimedOut)
	// On a timeout we must remove ourselves from the job manager.
	_ = j.jm.getJob(j.JobID, true)
}

func (j *AsyncJob) fail() {
	j.complete(ErrJobFailed)
}

func (j *AsyncJob) cancel() {
	j.complete(ErrJobCancelled)
}

func (j *AsyncJob) heartbeat() {
	// If the timer has already fired Stop() returns false, this means our job already
	// timed out.
	if j.timer.Stop() {
		j.timer.Reset(jobTimeoutDuration)
	}
}

// Wait for the job to finish synchronously.
func (j *AsyncJob) Wait(ctx context.Context) (CallbackEvent, error) {
	select {
	case ev := <-j.channel:
		switch ev.(type) {
		case error:
			return nil, ev.(error)
		default:
			return ev.(CallbackEvent), nil
		}
	case <-ctx.Done():
		// If the context is cancelled remove the job from the job
		// manager.
		_ = j.jm.getJob(j.JobID, true)
		return nil, ctx.Err()
	}

}

// complete the job with the given result, if the job has already been completed
// this will have no effect.
func (j *AsyncJob) complete(event interface{}) {
	if atomic.CompareAndSwapUint32(&j.completed, 0, 1) {
		j.channel <- event
		close(j.channel)
	}
}

// JobManager manages the collection of asynchronous jobs.
type JobManager struct {
	currentJobID uint64
	lock         sync.RWMutex
	currentJobs  map[protocol.JobID]*AsyncJob
}

func newJobManager(client *Client) *JobManager {
	return &JobManager{
		currentJobs: make(map[protocol.JobID]*AsyncJob),
	}
}

// NewJob creates a new AsyncJob and adds it to the set of managed jobs.
func (jm *JobManager) NewJob() *AsyncJob {
	// TODO: Do we need to split the timer start so the timeout
	//       doesn't start until the message is sent?
	job := newJob(jm)

	jm.lock.Lock()
	defer jm.lock.Unlock()

	jm.currentJobs[job.JobID] = job

	return job
}

// HandlePacket recieves all packets from the Steam3 network.
func (jm *JobManager) HandlePacket(packet *protocol.Packet) {
	switch packet.EMsg {
	case steamlang.EMsg_JobHeartbeat:
		jm.handleJobHeartbeat(packet.TargetJobID)
	case steamlang.EMsg_DestJobFailed:
		jm.handleJobFailed(packet.TargetJobID)
	}
}

func (jm *JobManager) handleJobHeartbeat(jobID protocol.JobID) {
	// On a heartbeat we extend the expiration timer as Steam is indicating a response will follow
	if job := jm.getJob(jobID, false); job != nil {
		job.heartbeat()
	}
}

func (jm *JobManager) handleJobFailed(jobID protocol.JobID) {
	// Job has failed - cancel it
	if job := jm.getJob(jobID, true); job != nil {
		job.fail()
	}
}

func (jm *JobManager) getJob(jobID protocol.JobID, remove bool) *AsyncJob {
	// JobID 0 will never be used, we bail quickly to avoid taking the lock
	if jobID == 0 {
		return nil
	}

	if remove {
		jm.lock.Lock()
		defer jm.lock.Unlock()
	} else {
		jm.lock.RLock()
		defer jm.lock.RUnlock()
	}

	if job, ok := jm.currentJobs[jobID]; ok {
		if remove {
			delete(jm.currentJobs, jobID)
		}

		return job
	}

	return nil
}

// NextJobID returns the next sequentially available job ID atomically.
func (jm *JobManager) NextJobID() protocol.JobID {
	return protocol.JobID(atomic.AddUint64(&jm.currentJobID, 1))
}

// CancelAllJobs marks all in progress jobs as cancelled - this is called upon disconnect.
func (jm *JobManager) CancelAllJobs() {
	jm.lock.Lock()
	defer jm.lock.Unlock()

	for _, job := range jm.currentJobs {
		job.cancel()
	}

	// Replace the job list with a new empty list
	jm.currentJobs = make(map[protocol.JobID]*AsyncJob)
}

// CompleteJob completes a specific job with the given event.
func (jm *JobManager) CompleteJob(event CallbackEvent) {
	if job := jm.getJob(event.GetJobID(), true); job != nil {
		job.complete(event)
	}
}

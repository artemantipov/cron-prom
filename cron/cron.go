package cron

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"../metrics"

	"gopkg.in/robfig/cron.v3"
)

// Cron jobs
type Cron struct {
	Name    string
	Time    string
	Command []string
}

//Start cron with jobs from env with prefixes CRON_JOB_
func Start() {

	cronJobs := []Cron{}

	for _, env := range os.Environ() {
		pair := strings.Split(env, "=")
		if strings.HasPrefix(pair[0], "CRON_JOB_") {
			var job Cron
			job.Name = strings.ReplaceAll(pair[0], "CRON_JOB_", "")
			s := strings.Fields(pair[1])
			if len(s) < 6 {
				log.Printf("Wrong cron format for %v\n", job.Name)
				continue
			} else {
				job.Time = fmt.Sprintf("%v", strings.Join(s[:5], " "))
				job.Command = s[5:]
				cronJobs = append(cronJobs, job)
			}
		}
	}

	// Create cron jobs for all metrics
	c := cron.New()
	for _, job := range cronJobs {
		c.AddFunc(job.Time, runCommand(job.Name, job.Command))
		log.Printf("Job %v added", job.Name)
	}
	c.Start()

}

func runCommand(job string, command []string) func() {
	return func() {
		cmd := exec.Command(command[0], command[1:]...)
		stdout, _ := cmd.StdoutPipe()
		cmd.Start()
		log.Printf("Running: %v", job)
		scanner := bufio.NewScanner(stdout)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			m := scanner.Text()
			fmt.Println(m)
		}
		err := cmd.Wait()
		if err != nil {
			log.Printf("Finished with error: %v", err)
			metrics.JobsFailed.Inc()
		} else {
			log.Printf("Finished: %v", job)
		}
	}
}

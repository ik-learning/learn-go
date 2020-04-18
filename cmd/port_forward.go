import (
	"fmt"
	"os"
	"os/exec"
	. "integration/utils"
	"time"
	"syscall"
	"regexp"
	"strconv"
	"io"
  
	"net/http"
	"net/url"

	"k8s.io/apimachinery/pkg/util/wait"
	"github.com/DATA-DOG/godog"
)

var (
	portForwardRegexp = regexp.MustCompile("Forwarding from (127.0.0.1|\\[::1\\]):([0-9]+)")
)


type portForwardCommand struct {
	cmd  *exec.Cmd
	port int
}

func (c *portForwardCommand) Stop() {
	if err := c.cmd.Process.Signal(syscall.SIGINT); err != nil {
		panic(fmt.Sprintf("error sending SIGINT to kubectl port-forward: %v", err))
	}

	// try to wait for a clean exit
	done := make(chan error)
	go func() {
		done <- c.cmd.Wait()
	}()

	expired := time.NewTimer(wait.ForeverTestTimeout)
	defer expired.Stop()

	select {
	case err := <-done:
		if err == nil {
			// success
			return
		}
		panic(fmt.Sprintf("error waiting for kubectl port-forward to exit: %v", err))
	case <-expired.C:
		panic(fmt.Sprintf("timed out waiting for kubectl port-forward to exit"))
	}

	fmt.Println("trying to forcibly kill kubectl")
	TryKill(c.cmd)
}

func runPortForward(name string, args ...string) *portForwardCommand {
	cmd := exec.Command(name, args...)
	cmd.Env = os.Environ()

	portOutput, _, err := StartCmdAndStreamOutput(cmd)
	if err != nil {
		panic(fmt.Sprintf("Failed to start port-forward command: %v", err))
	}

	buf := make([]byte, 128)

	var n int
	if n, err = portOutput.Read(buf); err != nil {
		panic(fmt.Sprintf("Failed to read from kubectl port-forward stdout: %v", err))
	}
	portForwardOutput := string(buf[:n])
	match := portForwardRegexp.FindStringSubmatch(portForwardOutput)
	if len(match) != 3 {
		panic(fmt.Sprintf("Failed to parse kubectl port-forward output: %s", portForwardOutput))
	}

	listenPort, err := strconv.Atoi(match[2])
	if err != nil {
		fmt.Println("Error converting %s to an int: %v", match[2], err)
	}

	return &portForwardCommand{
		cmd:  cmd,
		port: listenPort,
	}
}

func TryKill(cmd *exec.Cmd) {
	if err := cmd.Process.Kill(); err != nil {
		fmt.Println("ERROR failed to kill command %v! The process may leak", cmd)
	}
}

func StartCmdAndStreamOutput(cmd *exec.Cmd) (stdout, stderr io.ReadCloser, err error) {
	stdout, err = cmd.StdoutPipe()
	if err != nil {
		return
	}
	stderr, err = cmd.StderrPipe()
	if err != nil {
		return
	}
	err = cmd.Start()
	return
}

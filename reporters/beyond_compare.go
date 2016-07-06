package reporters

import (
	"fmt"
	"os/exec"

	"github.com/Approvals/ApprovalTests_go/utils"
)

type beyondCompare struct {
}

func NewBeyondCompareReporter() Reporter {
	return &beyondCompare{}
}

func (s *beyondCompare) Report(approved, received string) bool {
	programName := "C:/Program Files/Beyond Compare 4/BComp.exe"

	if !utils.DoesFileExist(programName) {
		return false
	}

	utils.EnsureExists(approved)

	cmd := exec.Command(programName, received, approved)
	cmd.Start()

	err := cmd.Wait()
	if err != nil {
		panic(fmt.Sprintf("err=%s", err))
	}

	return true
}
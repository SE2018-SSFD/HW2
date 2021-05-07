package tests

import (
	"backend/dao"
	"backend/model"

	// "backend/model"
	"backend/service"
	"backend/utils"
	"testing"

	"github.com/prashantv/gostub"
)

func TestNewProject(t *testing.T) {
	//s1->c1->s2->s4
	stub := gostub.StubFunc(&service.CheckToken, uint(1))
	stub.StubFunc(&dao.CreateProject, uint(1))
	stub.StubFunc(&dao.AddProjectToUser)
	success, msg, data := service.NewProject(utils.NewProjectParams{})
	t.Log(success, msg, data)
	stub.Reset()

	//s1->c1->s3->s4
	stub = gostub.StubFunc(&service.CheckToken, uint(0))
	stub.StubFunc(&dao.CreateProject, uint(1))
	stub.StubFunc(&dao.AddProjectToUser)
	success, msg, data = service.NewProject(utils.NewProjectParams{})
	t.Log(success, msg, data)
	stub.Reset()
}

func TestNewRoomProjectAndFile(t *testing.T) {
	//s1->s2->s3
	stub := gostub.StubFunc(&dao.CreateProject, uint(1))
	stub.StubFunc(&dao.AddProjectToUser)
	stub.StubFunc(&dao.CreateFile, uint(1))
	service.NewRoomProjectAndFile("project", "filename", 1, 1, "file")
	stub.Reset()
}

func TestModifyProject(t *testing.T) {

	//s1->c1->s2->c2->s3->s6
	stub := gostub.StubFunc(&service.CheckToken, uint(1))
	stub.StubFunc(&dao.GetPidsByUid, []uint{1})
	stub.StubFunc(&utils.UintListContains, bool(false))

	service.ModifyProject(utils.ModifyProjectParams{})
	stub.Reset()

	//s1->c1->s2->c2->s4->s6
	stub = gostub.StubFunc(&service.CheckToken, uint(1))
	stub.StubFunc(&dao.GetPidsByUid, []uint{1})
	stub.StubFunc(&utils.UintListContains, bool(true))
	stub.StubFunc(&dao.GetProjectByPid, model.Project{Pid: 1})
	stub.StubFunc(&dao.SetProject)

	service.ModifyProject(utils.ModifyProjectParams{})
	stub.Reset()

	//s1->c1->s5->s6
	stub = gostub.StubFunc(&service.CheckToken, uint(0))

	service.ModifyProject(utils.ModifyProjectParams{})
	stub.Reset()
}

func TestGetProject(t *testing.T) {
	//s1->c1->s2->c2->s3->s6
	stub := gostub.StubFunc(&service.CheckToken, uint(1))
	stub.StubFunc(&dao.GetPidsByUid, []uint{1})
	stub.StubFunc(&utils.UintListContains, bool(false))

	service.GetProject(utils.GetProjectParams{})
	stub.Reset()

	//s1->c1->s2->c2->s4->s6
	stub = gostub.StubFunc(&service.CheckToken, uint(1))
	stub.StubFunc(&dao.GetPidsByUid, []uint{1})
	stub.StubFunc(&utils.UintListContains, bool(true))
	stub.StubFunc(&dao.GetProjectWithFilesByPid, model.Project{Pid: 1})
	service.GetProject(utils.GetProjectParams{})
	stub.Reset()

	//s1->c1->s5->s6
	stub = gostub.StubFunc(&service.CheckToken, uint(0))

	service.GetProject(utils.GetProjectParams{})
	stub.Reset()
}

func TestDeleteProject(t *testing.T) {
	//s1->c1->s2->c2->s3->s6
	stub := gostub.StubFunc(&service.CheckToken, uint(1))
	stub.StubFunc(&dao.GetPidsByUid, []uint{1})
	stub.StubFunc(&utils.UintListContains, bool(false))

	service.DeleteProject(utils.DeleteProjectParams{})
	stub.Reset()

	//s1->c1->s2->c2->s4->s6
	stub = gostub.StubFunc(&service.CheckToken, uint(1))
	stub.StubFunc(&dao.GetPidsByUid, []uint{1})
	stub.StubFunc(&utils.UintListContains, bool(true))
	stub.StubFunc(&dao.DeleteProject)

	service.DeleteProject(utils.DeleteProjectParams{})
	stub.Reset()

	// //s1->c1->s5->s6
	stub = gostub.StubFunc(&service.CheckToken, uint(0))

	service.DeleteProject(utils.DeleteProjectParams{})
	stub.Reset()
}

func TestNewFile(t *testing.T) {

	//s1->c1->s2->c2->s3->s6
	stub := gostub.StubFunc(&service.CheckToken, uint(1))
	stub.StubFunc(&dao.GetPidsByUid, []uint{1})
	stub.StubFunc(&utils.UintListContains, bool(false))

	service.NewFile(utils.NewFileParams{})
	stub.Reset()

	//s1->c1->s2->c2->s4->s6
	stub = gostub.StubFunc(&service.CheckToken, uint(1))
	stub.StubFunc(&dao.GetPidsByUid, []uint{1})
	stub.StubFunc(&utils.UintListContains, bool(true))
	stub.StubFunc(&dao.CreateFile, uint(1))
	service.NewFile(utils.NewFileParams{})
	stub.Reset()

	//s1->c1->s5->s6
	stub = gostub.StubFunc(&service.CheckToken, uint(0))

	service.NewFile(utils.NewFileParams{})
	stub.Reset()
}

func TestModifyFile(t *testing.T) {

	//s1->c1->s2->c2->s3->s6
	stub := gostub.StubFunc(&service.CheckToken, uint(1))
	stub.StubFunc(&dao.GetPidsByUid, []uint{1})
	stub.StubFunc(&dao.GetFileByFid, model.File{Fid: 1})
	stub.StubFunc(&utils.UintListContains, bool(false))

	service.ModifyFile(utils.ModifyFileParams{})
	stub.Reset()

	//s1->c1->s2->c2->s4->s6
	stub = gostub.StubFunc(&service.CheckToken, uint(1))
	stub.StubFunc(&dao.GetPidsByUid, []uint{1})
	stub.StubFunc(&dao.GetFileByFid, model.File{Fid: 1})
	stub.StubFunc(&utils.UintListContains, bool(true))
	stub.StubFunc(&dao.SetFile)

	service.ModifyFile(utils.ModifyFileParams{})
	stub.Reset()

	// //s1->c1->s5->s6
	stub = gostub.StubFunc(&service.CheckToken, uint(0))

	service.ModifyFile(utils.ModifyFileParams{})
	stub.Reset()
}

func TestDeleteFile(t *testing.T) {

	//s1->c1->s2->c2->s3->s6
	stub := gostub.StubFunc(&service.CheckToken, uint(1))
	stub.StubFunc(&dao.GetPidsByUid, []uint{1})
	stub.StubFunc(&dao.GetFileByFid, model.File{Fid: 1})
	stub.StubFunc(&utils.UintListContains, bool(false))

	service.DeleteFile(utils.DeleteFileParams{})
	stub.Reset()

	//s1->c1->s2->c2->s4->s6
	stub = gostub.StubFunc(&service.CheckToken, uint(1))
	stub.StubFunc(&dao.GetPidsByUid, []uint{1})
	stub.StubFunc(&dao.GetFileByFid, model.File{Fid: 1})
	stub.StubFunc(&utils.UintListContains, bool(true))
	stub.StubFunc(&dao.DeleteFile)

	service.DeleteFile(utils.DeleteFileParams{})
	stub.Reset()

	// //s1->c1->s5->s6
	stub = gostub.StubFunc(&service.CheckToken, uint(0))

	service.DeleteFile(utils.DeleteFileParams{})
	stub.Reset()
}

package supply

import (
	"github.com/cloudfoundry/libbuildpack"
	"io"
)

type Stager interface {
	//TODO: See more options at https://github.com/cloudfoundry/libbuildpack/blob/master/stager.go
	BuildDir() string
	DepDir() string
	DepsIdx() string
	DepsDir() string
}

type Manifest interface {
	//TODO: See more options at https://github.com/cloudfoundry/libbuildpack/blob/master/manifest.go
	AllDependencyVersions(string) []string
	DefaultVersion(string) (libbuildpack.Dependency, error)
}

type Installer interface {
	//TODO: See more options at https://github.com/cloudfoundry/libbuildpack/blob/master/installer.go
	InstallDependency(libbuildpack.Dependency, string) error
	InstallOnlyVersion(string, string) error
}

type Command interface {
	//TODO: See more options at https://github.com/cloudfoundry/libbuildpack/blob/master/command.go
	Execute(string, io.Writer, io.Writer, string, ...string) error
	Output(dir string, program string, args ...string) (string, error)
}

type Supplier struct {
	Manifest  Manifest
	Installer Installer
	Stager    Stager
	Command   Command
	Log       *libbuildpack.Logger
}

func (s *Supplier) Run() error {
	s.Log.BeginStep("Supplying owasp-scanner")

	dep := libbuildpack.Dependency{Name: "java", Version: "8.222"}
	if err := s.Installer.InstallDependency(dep, s.Stager.DepsDir()); err != nil {
		return err
	}

	dep = libbuildpack.Dependency{Name: "dependency-check", Version: "5.2.1"}
	if err := s.Installer.InstallDependency(dep, s.Stager.DepsDir()); err != nil {
		return err
	}

	return nil
}

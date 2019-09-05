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
	CacheDir() string
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

	s.Log.Info("Cache Directory Path: %v", s.Stager.CacheDir())
	s.Log.Info("Build Directory Path: %v", s.Stager.BuildDir())
	s.Log.Info("Deps Directory Path: %v", s.Stager.DepsDir())
	s.Log.Info("Dep Directory Path: %v", s.Stager.DepDir())

	dep := libbuildpack.Dependency{Name: "java", Version: "0.0.0"}
	if err := s.Installer.InstallDependency(dep, s.Stager.CacheDir()); err != nil {
		return err
	}

	dep2 := libbuildpack.Dependency{Name: "dependency-check", Version: "0.0.0"}
	if err := s.Installer.InstallDependency(dep2, s.Stager.CacheDir()); err != nil {
		return err
	}

	return nil
}

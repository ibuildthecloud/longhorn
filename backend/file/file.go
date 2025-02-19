package file

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/rancher/longhorn/types"
)

func New() types.BackendFactory {
	return &Factory{}
}

type Factory struct {
}

type Wrapper struct {
	*os.File
}

func (f *Wrapper) Close() error {
	logrus.Infof("Closing: %s", f.Name())
	return f.File.Close()
}

func (f *Wrapper) Snapshot(name string) error {
	return nil
}

func (f *Wrapper) Size() (int64, error) {
	stat, err := f.Stat()
	if err != nil {
		return 0, err
	}
	return stat.Size(), nil
}

func (ff *Factory) Create(address string) (types.Backend, error) {
	logrus.Infof("Creating file: %s", address)
	file, err := os.OpenFile(address, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		logrus.Infof("Failed to create file %s: %v", address, err)
		return nil, err
	}

	return &Wrapper{file}, nil
}

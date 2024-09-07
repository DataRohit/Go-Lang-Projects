package database

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/datarohit/go-database/pkg/schemas"
	"github.com/jcelliott/lumber"
)

type Driver struct {
	Mutex   sync.Mutex
	Mutexes map[string]*sync.Mutex
	Dir     string
	Log     schemas.Logger
}

func NewDatabase(dir string, options *schemas.Options) (*Driver, error) {
	dir = filepath.Clean(dir)

	opts := schemas.Options{}
	if options != nil {
		opts = *options
	}

	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger(lumber.INFO)
	}

	driver := Driver{
		Dir:     dir,
		Mutexes: make(map[string]*sync.Mutex),
		Log:     opts.Logger,
	}

	if _, err := os.Stat(dir); err == nil {
		opts.Logger.Debug("Database already exists")
		return &driver, nil
	}

	opts.Logger.Debug("Creating the database")
	return &driver, os.MkdirAll(dir, 0755)
}

func (driver *Driver) GetOrCreateMutex(collection string) *sync.Mutex {
	driver.Mutex.Lock()
	defer driver.Mutex.Unlock()

	mutex, ok := driver.Mutexes[collection]
	if !ok {
		mutex = &sync.Mutex{}
		driver.Mutexes[collection] = mutex
	}

	return mutex
}

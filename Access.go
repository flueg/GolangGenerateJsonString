package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	log "github.com/cihub/seelog"
)

type Access struct {
	DBPath  string //文件路径
	OutFile string // result file output

	Result []string // save the resutl in array firstly. then output to file
	// at the end.
	rwlock sync.RWMutex // 读写锁保证协程安全
	flock  sync.Mutex
}

func NewAccess() *Access {
	return &Access{}
}

func (m *Access) Init(dbConf string) error {
	log.Tracef("Initialazing db %s", dbConf)
	m.DBPath = dbConf
	m.OutFile = dbConf + ".output"
	fd, err := os.Stat(m.DBPath)
	if err != nil {
		log.Errorf("Failed to load file:%s\n", dbConf)
		return fmt.Errorf("门票数据文件%s出错:%s", dbConf, err)
	}

	if fd.IsDir() {
		log.Errorf("Directory is not expected:%s\n", dbConf)
		return fmt.Errorf("Directory is not expected:%s", dbConf)
	}

	return nil
}

// parse ticket info from file. File contains line format:
// CodeID	Type	IsCheckin	Price	Desc
func _parseLine(line string) (string, error) {
	separator := "    "
	values := strings.Split(string(line), separator)

	opt := NewOption(values[0], strings.Split(values[1], ";"))
	if opt != nil {
		result, err := opt.GetJsonString()
		if err != nil {
			return "", fmt.Errorf("Failed to parse options to json.")
		}
		return result, nil
	}
	return "", fmt.Errorf("Failed to parse options.")
}

func (m *Access) Load() (rnt_err error) {
	log.Tracef("Loading tickets from dbfile: %s", m.DBPath)
	m.rwlock.RLock()
	m.flock.Lock()
	defer m.rwlock.RUnlock()
	defer m.flock.Unlock()

	log.Tracef("Try to open file:%s", m.DBPath)
	fd, err := os.Open(m.DBPath)
	if err != nil {
		log.Errorf("Failed to open db file:%s", m.DBPath)
		return fmt.Errorf("%s", err)
	}
	defer func() {
		if err := recover(); err != nil {
			//log.Errorf("Panic when proccess data. err:%v", err)
			fmt.Printf("Panic when proccess data. err:%v\n", err)
			rnt_err = fmt.Errorf("Server panics")
		}
		fd.Close()
	}()

	reader := bufio.NewReader(fd)
	lineNum := 0
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			log.Trace("End of files")
			break
		}
		//log.Tracef("Read line-%d: %s", lineNum, string(line))
		//fmt.Printf("Read line-%d: %s\n", lineNum, string(line))
		lineNum = lineNum + 1
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		myTicket, err := _parseLine(string(line))
		if err != nil {
			log.Errorf("Invalid ticket info found. [%s]", line)
			continue
		}
		//log.Tracef("Fetch ticket code:%s", myTicket)
		fmt.Printf("%s\n", myTicket)
		m.Result = append(m.Result, myTicket)
	}

	return nil
}

func (m *Access) Dump() (rnt_err error) {
	//log.Debug("Try to dump tickets changes into file.")
	m.flock.Lock()
	defer m.flock.Unlock()

	fileName := m.OutFile + ".tmp"
	data := strings.Join(m.Result, "\n")
	err := ioutil.WriteFile(m.OutFile, []byte(data), 0755)
	if err != nil {
		log.Errorf("write data into %s error:%s", fileName, err)
		return err
	}
	log.Tracef("Save options json string into %s successfully.", fileName)

	err = os.Rename(fileName, m.OutFile)
	if err != nil {
		log.Errorf("Failed to save options json string. err:%s", err)
		return err
	}
	return nil
}

package repository

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"

	"github.com/916RasnitaRadu/todo-cli/file"
	"github.com/916RasnitaRadu/todo-cli/types"
)

type FileRepository struct {
	path string
}

func NewFileRepository(path string) Repository {
	return &FileRepository{path}
}

func (r *FileRepository) GetTasks() ([]types.Task, error) {
	var tasks []types.Task

	f, err := file.LoadFile(r.path)
	if err != nil {
		return nil, err
	}
	defer file.CloseFile(f)

	reader := csv.NewReader(f)

	// reading the header
	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("error reading the header")
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error reading row:", err)
			continue
		}

		tasks = append(tasks, types.TaskFromCSV(record))
	}

	return tasks, nil
}

func (r *FileRepository) Create(task types.Task) error {
	f, err := file.LoadFile(r.path)
	if err != nil {
		return err
	}
	defer file.CloseFile(f)

	writer := csv.NewWriter(f)

	row := []string{strconv.FormatInt(int64(task.ID), 10), task.Description, task.CreatedAt, strconv.FormatBool(task.Done)}
	if err := writer.Write(row); err != nil {
		return fmt.Errorf("error creating new row")
	}

	writer.Flush()
	if err = writer.Error(); err != nil {
		return fmt.Errorf("error at flushing writer")
	}

	return nil
}

func (r *FileRepository) Delete(id int) error {
	f, err := file.LoadFile(r.path)
	if err != nil {
		return err
	}
	defer file.CloseFile(f)

	reader := csv.NewReader(f)

	// reading the header
	if _, err := reader.Read(); err != nil {
		return fmt.Errorf("error reading the header")
	}
	var newRecords [][]string
	for {
		task, err := reader.Read()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			fmt.Println("error reading row:", err)
			continue
		}

		taskId, _ := strconv.Atoi(task[0])
		if id != taskId {
			newRecords = append(newRecords, task)
		}

	}

	// TODO: finish delete function: truncate + append

	return nil
}

func (r *FileRepository) ChangeStatus(id int) error {
	return nil
}

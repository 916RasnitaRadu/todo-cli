package repository

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"

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

	row := []string{
		strconv.Itoa(task.ID),
		task.Description,
		task.CreatedAt.Format(time.DateTime),
		strconv.FormatBool(task.Done),
		"\n",
	}
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

	var newRecords [][]string
	for {
		task, err := reader.Read()
		if err == io.EOF {
			break
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

	if err = f.Truncate(0); err != nil {
		return fmt.Errorf("failed to truncate file: %w", err)
	}
	if _, err = f.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to move cursor at beginning of file: %w", err)
	}

	writer := csv.NewWriter(f)

	if err := writer.WriteAll(newRecords); err != nil {
		return fmt.Errorf("failed writing new records: %w", err)
	}
	writer.Flush()

	return nil
}

func (r *FileRepository) ChangeStatus(id int) error {
	f, err := file.LoadFile(r.path)
	if err != nil {
		return err
	}
	defer file.CloseFile(f)

	reader := csv.NewReader(f)

	var newRecords [][]string

	for {
		task, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error reading row:", err)
			continue
		}

		var newTask []string = task
		taskId, _ := strconv.Atoi(task[0])
		if taskId == id {
			status, _ := strconv.ParseBool(task[3])
			newTask = []string{task[0], task[1], task[2], strconv.FormatBool(!status), "\n"}
		}
		newRecords = append(newRecords, newTask)
	}

	if err = f.Truncate(0); err != nil {
		return fmt.Errorf("failed to truncate file: %w", err)
	}
	if _, err = f.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to move cursor at beginning of file: %w", err)
	}

	writer := csv.NewWriter(f)

	if err := writer.WriteAll(newRecords); err != nil {
		return fmt.Errorf("failed writing new records: %w", err)
	}
	writer.Flush()

	return nil
}

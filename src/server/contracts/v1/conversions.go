package v1

import (
	"../../../storage/models"
)

// CreateContractTask func
func CreateContractTask(m *models.Task) Task {
	logEntries := []LogEntry{}
	for _, le := range m.Logs {
		logEntries = append(logEntries, CreateContractLogEntry(&le))
	}
	return Task{
		Name:       m.Name,
		LogEntries: logEntries,
	}
}

// CreateContractLogEntry func
func CreateContractLogEntry(m *models.LogEntry) LogEntry {
	return LogEntry{
		Start: m.Start,
		End:   m.End,
	}
}

// CreateModelsTask func
func CreateModelsTask(id string, m *Task) models.Task {
	logEntries := []models.LogEntry{}
	for _, le := range m.LogEntries {
		logEntries = append(logEntries, CreateModelsLogEntry(&le))
	}

	return models.Task{
		Name: m.Name,
		Logs: logEntries,
	}
}

// CreateModelsLogEntry func
func CreateModelsLogEntry(m *LogEntry) models.LogEntry {
	return models.LogEntry{
		Start: m.Start,
		End:   m.End,
	}
}

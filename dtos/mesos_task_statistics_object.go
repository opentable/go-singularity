package dtos

import "io"

type MesosTaskStatisticsObject struct {
	CpusLimit             int32   `json:"cpusLimit"`
	CpusNrPeriods         int64   `json:"cpusNrPeriods"`
	CpusNrThrottled       int64   `json:"cpusNrThrottled"`
	CpusSystemTimeSecs    float64 `json:"cpusSystemTimeSecs"`
	CpusThrottledTimeSecs float64 `json:"cpusThrottledTimeSecs"`
	CpusUserTimeSecs      float64 `json:"cpusUserTimeSecs"`
	MemAnonBytes          int64   `json:"memAnonBytes"`
	MemFileBytes          int64   `json:"memFileBytes"`
	MemLimitBytes         int64   `json:"memLimitBytes"`
	MemMappedFileBytes    int64   `json:"memMappedFileBytes"`
	MemRssBytes           int64   `json:"memRssBytes"`
	Timestamp             float64 `json:"timestamp"`
}

func (self *MesosTaskStatisticsObject) Populate(jsonReader io.ReadCloser) (err error) {
	err = ReadPopulate(jsonReader, self)
	return
}

func (self *MesosTaskStatisticsObject) FormatText() string {
	return FormatText(self)
}

func (self *MesosTaskStatisticsObject) FormatJSON() string {
	return FormatJSON(self)
}

type MesosTaskStatisticsObjectList []*MesosTaskStatisticsObject

func (list *MesosTaskStatisticsObjectList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list *MesosTaskStatisticsObjectList) FormatText() string {
	text := []byte{}
	for _, dto := range *list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list *MesosTaskStatisticsObjectList) FormatJSON() string {
	return FormatJSON(list)
}

package dtos

import "io"

type MesosTaskStatisticsObject struct {
	CpusLimit             int32
	CpusNrPeriods         int64
	CpusNrThrottled       int64
	CpusSystemTimeSecs    float64
	CpusThrottledTimeSecs float64
	CpusUserTimeSecs      float64
	MemAnonBytes          int64
	MemFileBytes          int64
	MemLimitBytes         int64
	MemMappedFileBytes    int64
	MemRssBytes           int64
	Timestamp             float64
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

func (list MesosTaskStatisticsObjectList) Populate(jsonReader io.ReadCloser) (err error) {
	return ReadPopulate(jsonReader, list)
}

func (list MesosTaskStatisticsObjectList) FormatText() string {
	text := []byte{}
	for _, dto := range list {
		text = append(text, (*dto).FormatText()...)
		text = append(text, "\n"...)
	}
	return string(text)
}

func (list MesosTaskStatisticsObjectList) FormatJSON() string {
	return FormatJSON(list)
}

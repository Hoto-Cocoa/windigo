package dshow

import (
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/ns-strmif-am_media_type
type AM_MEDIA_TYPE struct {
	Majortype            win.GUID
	Subtype              win.GUID
	bFixedSizeSamples    win.BOOL
	bTemporalCompression win.BOOL
	LSampleSize          uint32
	Formattype           win.GUID
	IUnknown             win.IUnknown
	CbFormat             uint32
	PbFormat             *byte
}

func (mt *AM_MEDIA_TYPE) BFixedSizeSamples() bool { return mt.bFixedSizeSamples != 0 }
func (mt *AM_MEDIA_TYPE) SetBFixedSizeSamples(val bool) {
	mt.bFixedSizeSamples = win.BOOL(util.BoolToUintptr(val))
}

func (mt *AM_MEDIA_TYPE) BTemporalCompression() bool { return mt.bTemporalCompression != 0 }
func (mt *AM_MEDIA_TYPE) SetBTemporalCompression(val bool) {
	mt.bTemporalCompression = win.BOOL(util.BoolToUintptr(val))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/ns-strmif-filter_info
type FILTER_INFO struct {
	achName [128]uint16
	PGraph  IFilterGraph
}

func (fi *FILTER_INFO) AchName() string { return win.Str.FromUint16Slice(fi.achName[:]) }
func (fi *FILTER_INFO) SetAchName(val string) {
	copy(fi.achName[:], win.Str.ToUint16Slice(win.Str.Substr(val, 0, len(fi.achName)-1)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/ns-evr-mfvideonormalizedrect
type MFVideoNormalizedRect struct {
	Left, Top, Right, Bottom float32
}

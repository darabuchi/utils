package utils

import (
	"io"
	"net/http"
	"strings"
)

// https://github.com/gabriel-vasile/mimetype

func GetContentType(out io.Reader) (*ContentType, error) {
	buffer := make([]byte, 512)
	_, err := out.Read(buffer)
	if err != nil {
		return nil, err
	}

	return NewContentType(http.DetectContentType(buffer)), nil
}

type ContentType struct {
	raw       string
	mediaTYpe MediaType
}

func NewContentType(raw string) *ContentType {
	p := &ContentType{
		raw: raw,
	}

	items := strings.Split(p.raw, "/")
	if len(items) > 0 {
		switch items[0] {
		case "application":
			if len(items) > 1 && strings.HasPrefix(items[1], "vnd") {
				p.mediaTYpe = MediaVND
			} else {
				p.mediaTYpe = MediaApplication
			}
		case "audio":
			p.mediaTYpe = MediaAudio
		case "image":
			p.mediaTYpe = MediaImage
		case "multipart":
			p.mediaTYpe = MediaMultipart
		case "text":
			p.mediaTYpe = MediaText
		case "video":
			p.mediaTYpe = MediaVideo
		default:
			p.mediaTYpe = MediaTypeUnknown
		}
	}

	return p
}

func (p *ContentType) MediaType() MediaType {
	return p.mediaTYpe
}

func (p *ContentType) String() string {
	return p.raw
}

type MediaType uint8

const (
	MediaTypeUnknown MediaType = iota
	MediaApplication
	MediaAudio
	MediaImage
	MediaMultipart
	MediaText
	MediaVideo
	MediaVND
)

func (p MediaType) String() string {
	switch p {
	case MediaApplication:
		return "application"
	case MediaAudio:
		return "audio"
	case MediaImage:
		return "image"
	case MediaMultipart:
		return "multipart"
	case MediaText:
		return "text"
	case MediaVideo:
		return "video"
	case MediaVND:
		return "application/vnd"
	case MediaTypeUnknown:
		fallthrough
	default:
		return "unknown"
	}
}

const (
	MimeJpeg      = "image/jpeg"
	MimeJpg       = "image/jpg"
	MimeGif       = "image/gif"
	MimePng       = "image/png"
	MimeHtml      = "text/html"
	MimePdf       = "application/pdf"
	MimeMpeg      = "audio/mpeg"
	MimeQuicktime = "video/quicktime"
	MimeMp4       = "video/mp4"
	MimeWebp      = "image/webp"
	MimeYml       = "application/x-yaml"
	MimeXlsx      = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	MimeMsExec    = "application/x-ms-dos-executable"
	MimeText      = "text/plain"
)

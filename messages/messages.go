package messages

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Liriker/YaMa/types"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"
)

const (
	sendMessageUrl   = "https://botapi.messenger.yandex.net/bot/v1/messages/sendText/"
	sendFileUrl      = "https://botapi.messenger.yandex.net/bot/v1/messages/sendFile/"
	getFileUrl       = "https://botapi.messenger.yandex.net/bot/v1/messages/getFile/"
	sendImageUrl     = "https://botapi.messenger.yandex.net/bot/v1/messages/sendImage/"
	sendGalleryUrl   = "https://botapi.messenger.yandex.net/bot/v1/messages/sendGallery/"
	deleteMessageUrl = "https://botapi.messenger.yandex.net/bot/v1/messages/delete"

	documentFiledName = "Document"
	imageFieldName    = "Image"
	boundary          = "-----FormAaB03xBoundary-----"
)

type Client struct {
	client  *http.Client
	headers http.Header
}

func NewClient(cl *http.Client, h http.Header) *Client {
	return &Client{
		client:  cl,
		headers: h,
	}
}

func (cl *Client) Send(message types.NewMessage) (int64, error) {
	data, err := json.Marshal(message)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(http.MethodPost, sendMessageUrl, bytes.NewBuffer(data))
	if err != nil {
		return 0, err
	}
	req.Header = cl.headers

	resp, err := cl.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, errors.New(fmt.Sprint(resp))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	result := response{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, err
	}
	if !result.Ok {
		return 0, errors.New(result.Description)
	}
	return result.MessageID, nil
}

func (cl *Client) SendFile(message types.NewFileMessage, filename string) (int64, error) {
	data, err := structToMultipartForm(message, filename)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(http.MethodPost, sendFileUrl, data)
	if err != nil {
		return 0, err
	}
	headers := cl.headers.Clone()
	headers.Set("Content-Disposition", "document; filename=\""+filename+"\"")
	headers.Set("Content-Type", "multipart/form-data; boundary="+boundary)
	req.Header = headers

	resp, err := cl.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return 0, err
		}
		return 0, errors.New(string(data))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, errors.New(fmt.Sprint(resp))
	}

	result := response{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, errors.New(fmt.Sprint(resp))
	}
	if !result.Ok {
		return 0, errors.New(result.Description)
	}
	return result.MessageID, nil
}

func (cl *Client) GetFile(id int64) (io.ReadCloser, error) {
	js := getFileRequest{FileID: id}
	data, err := json.Marshal(&js.FileID)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, getFileUrl, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header = cl.headers

	resp, err := cl.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprint(resp))
	}
	return resp.Body, nil
}

func (cl *Client) SendImage(message types.NewImageMessage, filename string) (int64, error) {
	data, err := structToMultipartForm(message, filename)
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequest(http.MethodPost, sendImageUrl, data)
	if err != nil {
		return 0, err
	}

	headers := cl.headers.Clone()
	headers.Set("Content-Type", "multipart/form-data; boundary="+boundary)
	req.Header = headers

	resp, err := cl.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return 0, errors.New(string(body))
	}

	if err != nil {
		return 0, err
	}
	result := response{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, err
	}
	if !result.Ok {
		return 0, errors.New(result.Description)
	}
	return result.MessageID, nil
}

func (cl *Client) SendGallery(message types.NewGalleryMessage, filenames ...string) (int64, error) {
	data, err := structToMultipartForm(message, filenames...)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(http.MethodPost, sendGalleryUrl, data)
	if err != nil {
		return 0, err
	}

	headers := cl.headers.Clone()
	headers.Set("Content-Type", "multipart/form-data; boundary="+boundary)
	req.Header = headers

	resp, err := cl.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return 0, errors.New(string(body))
	}
	if err != nil {
		return 0, err
	}
	result := response{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, err
	}
	if !result.Ok {
		return 0, errors.New(result.Description)
	}
	return result.MessageID, nil
}

func (cl *Client) Delete(request types.NewDeleteMessageRequest) (int64, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(http.MethodPost, deleteMessageUrl, bytes.NewBuffer(data))
	if err != nil {
		return 0, err
	}
	req.Header = cl.headers
	resp, err := cl.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, errors.New(string(body))
	}
	result := response{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, err
	}
	if !result.Ok {
		return 0, errors.New(result.Description)
	}
	return result.MessageID, nil

}

func structToMultipartForm(value any, filenames ...string) (io.Reader, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	err := w.SetBoundary(boundary)
	if err != nil {
		return nil, err
	}
	defer w.Close()
	structure := reflect.ValueOf(value)
	valueTypes := structure.Type()
	for i := 0; i < structure.NumField(); i++ {
		f := structure.Field(i)
		typeF := valueTypes.Field(i)
		tag := typeF.Tag.Get("json")
		tag = strings.ReplaceAll(tag, ",omitempty", "")
		if typeF.Name != documentFiledName && typeF.Name != imageFieldName && typeF.Name != imageFieldName+"s" {
			err = w.WriteField(tag, fmt.Sprintf("%v", f.Interface()))
			if err != nil {
				return nil, err
			}
		} else {
			if typeF.Type.Elem().Kind() == reflect.Slice {
				slice := f
				for j := 0; j < slice.Len(); j++ {
					wr, err := w.CreateFormFile(tag, filenames[j])
					if err != nil {
						return nil, err
					}
					_, err = wr.Write(slice.Index(j).Bytes())
					if err != nil {
						return nil, err
					}
				}
			} else {
				wr, err := w.CreateFormFile(tag, filenames[0])
				if err != nil {
					return nil, err
				}
				_, err = wr.Write(f.Bytes())
				if err != nil {
					return nil, err
				}
			}

		}
		err := w.WriteField(valueTypes.Field(i).Name, f.String())
		if err != nil {
			return nil, err
		}

	}
	return &b, nil
}

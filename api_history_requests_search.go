package singularity

import "bytes"

func (client *Client) getRequestHistoryForRequestLike(requestIdLike string, count int32, page int32) (response string, err error) {
	pathParamMap := map[string]interface{}{}
	queryParamMap := map[string]interface{}{
		"requestIdLike": requestIdLike, "count": count, "page": page,
	}

	resBody, err := client.Request("GET", "/api/history/requests/search", pathParamMap, queryParamMap)
	readBuf := bytes.Buffer{}
	readBuf.ReadFrom(resBody)
	response = string(readBuf.Bytes())
	return
}

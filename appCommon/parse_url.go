package appCommon

import "net/url"

func GetPathFromUrl(urlString string) (string, error) {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return "", ErrInvalidRequest(err)
	}
	res := parsedURL.Path

	res = res[1:]
	return res, nil
}

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type getBooksmarksResponse struct {
	Data   []Bookmark `json:"data"`
	Cursor string     `json:"cursor"`
}

type Bookmark struct {
	BookmarkID string    `json:"bookmarkId"`
	UserID     string    `json:"userId"`
	PostID     string    `json:"postId,omitempty"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"createdAt"`
	ReplyID    string    `json:"replyId,omitempty"`
}

func (c *APIClient) GetBookmarks(limit int, cursor string) (posts []Bookmark, newCursor string, err error) {
	url := makeGetUrl(c.ApiUrl+"/bookmarks", limit, cursor)

	req, err := makeRequest("GET", url, c.Tokens, nil)
	if err != nil {
		return []Bookmark{}, cursor, fmt.Errorf("Error forming request: %s", err)
	}

	res, err := c.sendRequest(req)
	if err != nil {
		return nil, cursor, fmt.Errorf("Error retrieving Bookmarks: %s", err)
	}

	var getBooksmarksResponse getBooksmarksResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&getBooksmarksResponse)
	if err != nil {
		panic(err)
	}
	//fmt.Print(getNotificationsReply)
	c.Cursors["bookmarks_standard"] = getBooksmarksResponse.Cursor
	return getBooksmarksResponse.Data, getBooksmarksResponse.Cursor, nil

}

type CreateBookmarkInput struct {
	ReplyID string `json:"replyId"`
	Type    string `json:"type"`
}

func (c *APIClient) CreateBookmark(id, bookmarkType string) error {
	bookmarkInput := CreateBookmarkInput{ReplyID: id, Type: bookmarkType}
	postJson, err := json.Marshal(bookmarkInput)
	if err != nil {
		panic(err)
	}
	req, err := makeRequest("POST", c.ApiUrl+"/bookmarks", c.Tokens, bytes.NewBuffer(postJson))
	if err != nil {
		return fmt.Errorf("Error making post request:%s", err)
	}
	_, err = c.sendRequest(req)
	if err != nil {
		return fmt.Errorf("Error sending post request:%s", err)
	}
	return nil

}

func (c *APIClient) DeleteBookmark(bookmarkId string) error {

	req, err := makeRequest("DELETE", c.ApiUrl+"/bookmarks/"+bookmarkId, c.Tokens, nil)
	if err != nil {
		return fmt.Errorf("Error forming delete request: %s", err)
	}
	_, err = c.sendRequest(req)
	if err != nil {
		return fmt.Errorf("Error during the request process: %s", err)
	}
	fmt.Print("Successfully deleted bookmark")
	return nil
}

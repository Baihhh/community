/*
 * Copyright (c) 2023 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package core

import (
	"context"
	"testing"
	"time"
)

func TestPutArticle(t *testing.T) {
	conf := &Config{}
	todo := context.TODO()
	community, err := New(todo, conf)

	if err != nil {
		t.Skip(err)
	}

	// test data
	article := &Article{
		ArticleEntry: ArticleEntry{
			Title: "Test Article",
			UId:   "1",
			Cover: "cover1",
			Tags:  "tag1",
			Ctime: time.Now(),
			Mtime: time.Now(),
		},
		Content: "This is a test article.",
	}
	articleUpdate := &Article{
		ArticleEntry: ArticleEntry{
			ID:    "1",
			Title: "Test Article",
			UId:   "1",
			Cover: "cover1",
			Tags:  "tag1",
			Ctime: time.Now(),
			Mtime: time.Now(),
		},
		Content: "This is a test article.",
	}

	tests := []struct {
		uid        string
		article    *Article
		expectedID string
	}{
		{"1", article, "1"},       // insert
		{"1", articleUpdate, "1"}, // update
	}

	for _, tt := range tests {
		id, _ := community.PutArticle(todo, tt.uid, tt.article)

		if id != tt.expectedID {
			t.Errorf("PutArticle(%s, %+v) returned ID %s, expected: %s", tt.uid, tt.article, id, tt.expectedID)
		}
	}
}

func TestCanEditable(t *testing.T) {
	conf := &Config{}
	todo := context.TODO()
	community, err := New(todo, conf)

	if err != nil {
		t.Skip(err)
	}

	// test data
	tests := []struct {
		uid           string
		articleID     string
		expectedEdit  bool
		expectedError error
	}{
		{"1", "1", true, nil},
		{"2", "1", false, ErrPermission},
	}

	for _, tt := range tests {
		canEdit, _ := community.CanEditable(todo, tt.uid, tt.articleID)

		if canEdit != tt.expectedEdit {
			t.Errorf("CanEditable(%s, %s) returned %t, expected: %t", tt.uid, tt.articleID, canEdit, tt.expectedEdit)
		}
	}
}

func TestArticle(t *testing.T) {
	conf := &Config{}
	todo := context.TODO()
	community, err := New(todo, conf)

	if err != nil {
		t.Skip(err)
	}

	// test data
	tests := []struct {
		id            string
		expectedID    string
		expectedError error
	}{
		{"1", "1", nil},
		{"10", "", ErrNotExist},
	}

	for _, tt := range tests {
		article, err := community.Article(todo, tt.id)

		if article.ID != tt.expectedID {
			t.Errorf("Article(%s) returned id is %s, expected: %s", tt.id, article.ID, tt.expectedID)
		}
		if err != tt.expectedError {
			t.Errorf("Article(%s) returned err is %s, expected: %s", tt.id, err, tt.expectedError)
		}
	}
}

func TestListArticle(t *testing.T) {
	conf := &Config{}
	todo := context.TODO()
	community, err := New(todo, conf)

	if err != nil {
		t.Skip(err)
	}

	tests := []struct {
		from         string
		limit        int
		expectedLen  int
		expectedNext string
	}{
		{MarkBegin, 5, 5, "5"},
		{"5", 5, 1, "6"},
		{"6", 5, 0, MarkEnd},
	}

	for _, tt := range tests {
		items, next, err := community.ListArticle(todo, tt.from, tt.limit)

		if err != nil {
			t.Errorf("ListArticle(%s, %d) returned error: %v", tt.from, tt.limit, err)
		}

		if len(items) != tt.expectedLen {
			t.Errorf("ListArticle(%s, %d) returned %d items, expected %d", tt.from, tt.limit, len(items), tt.expectedLen)
		}

		if next != tt.expectedNext {
			t.Errorf("ListArticle(%s, %d) returned next %s, expected %s", tt.from, tt.limit, next, tt.expectedNext)
		}
	}
}

func TestDeleteArticle(t *testing.T) {
	conf := &Config{}
	todo := context.TODO()
	community, err := New(todo, conf)

	if err != nil {
		t.Skip(err)
	}

	// test data
	tests := []struct {
		uid         string
		articleID   string
		expectedErr error
	}{
		{"22", "2", ErrPermission}, // no permission
		{"1", "1", nil},
	}

	for _, tt := range tests {
		err := community.DeleteArticle(todo, tt.uid, tt.articleID)

		if err != tt.expectedErr {
			t.Errorf("DeleteArticle(%s, %s) returned error: %v, expected: %v", tt.uid, tt.articleID, err, tt.expectedErr)
		}
	}
}

package repositories

import (
	"go-todo-app/models"
	"log"
	"testing"
)

func TestGetTodo(t *testing.T) {
	// DB接続
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// テスト名とテーブルの中身の構造体を定義
	testCases := []struct {
		name     string
		expected models.Todo
	}{
		{
			name: "test1",
			expected: models.Todo{
				ID:      1,
				Title:   "todo1 title",
				Content: "todo1 content",
			},
		},
		{
			name: "test2",
			expected: models.Todo{
				ID:      2,
				Title:   "todo2 title",
				Content: "todo2 content",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) { // サブテストを実行
			got, err := GetTodo(db, tc.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			// IDのテスト
			if got.ID != tc.expected.ID {
				t.Errorf("got %v, expected %v", got.ID, tc.expected.ID)
			}

			// Titleのテスト
			if got.Title != tc.expected.Title {
				t.Errorf("got %v, expected %v", got.Title, tc.expected.Title)
			}

			// Contentのテスト
			if got.Content != tc.expected.Content {
				t.Errorf("got %v, expected %v", got.Content, tc.expected.Content)
			}
		})
	}
}

func TestGetTodos(t *testing.T) {
	// DB接続
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	todos, err := GetTodos(db)
	if err != nil {
		log.Fatal(err)
	}

	if len(todos) < 2 {
		t.Errorf("Expected minimum todos: 2, got: %d", len(todos))
	}

	expectedTodos := []models.Todo{
		{
			ID:      1,
			Title:   "todo1 title",
			Content: "todo1 content",
		},
		{
			ID:      2,
			Title:   "todo2 title",
			Content: "todo2 content",
		},
	}

	for i, expected := range expectedTodos {
		if i >= len(todos) {
			t.Fatalf("Not enough todos returned")
		}

		// IDのテスト
		if todos[i].ID != expected.ID {
			t.Errorf("ID: got %v, expected %v", todos[i].ID, expected.ID)
		}

		// Titleのテスト
		if todos[i].Title != expected.Title {
			t.Errorf("Title: got %v, expected %v", todos[i].Title, expected.Title)
		}

		// Contentのテスト
		if todos[i].Content != expected.Content {
			t.Errorf("Content: got %v, expected %v", todos[i].Content, expected.Content)
		}
	}
}

func TestCreateTodo(t *testing.T) {
	// DB接続
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	newTodo := models.Todo{
		Title:   "newTodo title",
		Content: "newTodo content",
	}

	createdTodo, err := CreateTodo(db, newTodo)
	if err != nil {
		t.Fatal(err)
	}

	// IDが割り振られていることを確認
	if createdTodo.ID == 0 {
		t.Error("Created todo has no ID assigned")
	}

	// 作成されたTODOを取得して確認
	fetchedTodo, err := GetTodo(db, createdTodo.ID)
	if err != nil {
		t.Fatal(err)
	}

	if fetchedTodo.Title != createdTodo.Title {
		t.Errorf("Title: got %v, expected %v", fetchedTodo.Title, newTodo.Title)
	}

	if fetchedTodo.Content != createdTodo.Content {
		t.Errorf("Content: got %v, expected %v", fetchedTodo.Content, newTodo.Content)
	}

	// テスト終了のデータ削除
	err = DeleteTodo(db, createdTodo.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateTodo(t *testing.T) {
	// DB接続
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//　更新用のTodoを作成
	newTodo := models.Todo{
		Title:   "updateTodo title",
		Content: "updateTodo content",
	}

	createdTodo, err := CreateTodo(db, newTodo)
	if err != nil {
		t.Fatal(err)
	}

	// 作成したTodoを更新
	updatedTodo := models.Todo{
		ID:      createdTodo.ID,
		Title:   "updatedTodo title",
		Content: "updatedTodo content",
	}

	_, err = UpdateTodo(db, updatedTodo)
	if err != nil {
		t.Fatal(err)
	}

	// 更新したTodoを取得
	fetchedTodo, err := GetTodo(db, createdTodo.ID)
	if err != nil {
		t.Fatal(err)
	}

	if fetchedTodo.Title != updatedTodo.Title {
		t.Errorf("Title: got %v, expected %v", fetchedTodo.Title, updatedTodo.Title)
	}

	if fetchedTodo.Content != updatedTodo.Content {
		t.Errorf("Content: got %v, expected %v", fetchedTodo.Content, updatedTodo.Content)
	}

	// 更新したTodoを削除
	err = DeleteTodo(db, createdTodo.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteTodo(t *testing.T) {
	// DB接続
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	newTodo := models.Todo{
		Title:   "deleteTodo title",
		Content: "deleteTodo content",
	}

	createdTodo, err := CreateTodo(db, newTodo)
	if err != nil {
		t.Fatal(err)
	}

	err = DeleteTodo(db, createdTodo.ID)
	if err != nil {
		t.Fatal(err)
	}

	_, err = GetTodo(db, createdTodo.ID)
	if err == nil {
		t.Error("Todo was not deleted")
	}
}

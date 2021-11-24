package entity

type Todo struct {
	Id          uint64
	Title       string
	Description string
}

type TodoResponse struct {
	Id          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoCreateRequest struct {
	Title       string `json:"title" validate:"required,min=0,max=255"`
	Description string `json:"description" validate:"required,min=0,max=255"`
}

type TodoUpdateRequest struct {
	Id          uint64 `json:"id"`
	Title       string `json:"title" validate:"required,min=0,max=255"`
	Description string `json:"description" validate:"required,min=0,max=255"`
}

func ToTodoResponse(todo *Todo) TodoResponse {
	return TodoResponse{
		Id:          todo.Id,
		Title:       todo.Title,
		Description: todo.Description,
	}
}

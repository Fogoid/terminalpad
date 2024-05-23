package entities

type Note struct {
    Id string
    Title string
    Content string
}

func NewNote(id string, title string, content string) *Note {
    return &Note{
        Id: id,
        Title: title,
        Content: content,
    }
}

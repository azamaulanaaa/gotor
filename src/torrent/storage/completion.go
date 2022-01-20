package storage

type Completion struct {
    database *Database
    index int64
}

func NewCompletion(database *Database, index int) Completion {
    return Completion{
        database: database,
        index: int64(index),
    }
}

func (completion Completion) Set(value bool) {
    completion.database.Set(completion.index, value)
}

func (completion Completion) Get() bool {
    return completion.database.Get(completion.index)
}

